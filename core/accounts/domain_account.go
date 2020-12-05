package accounts

import (
	"errors"

	"resk.com/infra/base"
	"resk.com/services"

	"github.com/segmentio/ksuid"
	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"
	"github.com/tietang/dbx"
)

// 有状态的, 每次使用时都要实例化
type accountDomain struct {
	account    Account
	accountLog AccountLog
}

// 创建logNo的逻辑
func (domain *accountDomain) createAccountLogNo() {
	domain.accountLog.LogNo = ksuid.New().Next().String()
}

// 生成accountNo的逻辑
func (domain *accountDomain) createAccountNo() {
	domain.account.AccountNo = ksuid.New().Next().String()
}

// 创建账户的流水的记录
func (domain *accountDomain) createAccountLog() {
	domain.accountLog = AccountLog{}
	domain.createAccountLogNo()
	domain.accountLog.TradeNo = domain.accountLog.LogNo // 使用同一个值, emmm
	// 流水中的交易主体信息
	domain.accountLog.AccountNo = domain.account.AccountNo
	domain.accountLog.UserId = domain.account.UserId
	domain.accountLog.Username = domain.account.Username.String
	// 交易对象信息
	domain.accountLog.TargetAccountNo = domain.account.AccountNo
	domain.accountLog.TargetUserId = domain.account.UserId
	domain.accountLog.TargetUsername = domain.account.Username.String
	// 交易金额
	domain.accountLog.Amount = domain.account.Balance  // 交易金额
	domain.accountLog.Balance = domain.account.Balance // 交易后余额
	// 交易变化属性
	domain.accountLog.Decs = "账户创建"
	domain.accountLog.ChangeType = services.AccountCreated
	domain.accountLog.ChangeFlag = services.FlagAccountCreated
}

// 账户创建的业务逻辑
func (domain *accountDomain) Create(dto services.AccountDTO) (*services.AccountDTO, error) {
	domain.account = Account{}
	domain.account.FromDTO(&dto)         // 把"数据传输对象"转为"持久化对象"
	domain.createAccountNo()             // 生成账户编号
	domain.account.Username.Valid = true // true表示写入数据库
	domain.createAccountLog()            // 创建账户流水"持久化对象"

	accountDao := AccountDao{} // 创建"数据访问对象"
	accountLogDao := AccountLogDao{}
	var rdto *services.AccountDTO
	err := base.Tx(func(runner *dbx.TxRunner) error {
		accountDao.runner = runner // 赋予数据访问能力
		accountLogDao.runner = runner
		// 插入账户数据
		id, err := accountDao.Insert(&domain.account)
		if err != nil {
			return err
		}
		if id <= 0 {
			return errors.New("创建账户失败")
		}
		// 如果插入成功, 就插入流水数据
		id, err = accountLogDao.Insert(&domain.accountLog)
		if err != nil {
			return err
		}
		if id <= 0 {
			return errors.New("创建账户流水失败")
		}
		domain.account = *accountDao.GetOne(domain.account.AccountNo) // 获取创建账户的数据
		return nil
	})
	rdto = domain.account.ToDTO()
	return rdto, err
}

// 账户交易
func (a *accountDomain) Transfer(dto services.AccountTransferDTO) (status services.TransferedStatus, err error) {
	// 如果是支出, 修正amount
	if dto.ChangeFlag == services.FlagTransferOut {
		dto.Amount = dto.Amount.Mul(decimal.NewFromFloat(-1)) // 修正为负数
	}
	// 创建账户流水记录
	a.accountLog = AccountLog{}        // 创建持久化对象
	a.accountLog.FromTransferDTO(&dto) // 从数据传输对象赋值到持久化对象
	a.createAccountLogNo()             // 创建流水编号
	// 检查余额是否足够和更新余额: 通过乐观锁来验证, 更新余额的同事来验证余额是否足够
	// 更新成功后, 写入流水记录
	err = base.Tx(func(runner *dbx.TxRunner) error {
		accountDao := AccountDao{runner: runner}
		accountLogDao := AccountLogDao{runner: runner}
		rows, err := accountDao.UpdateBalance(dto.TradeBody.AccountNo, dto.Amount)
		if err != nil {
			// 执行失败 => 事务回滚
			status = services.TransferedStatusFailure
			return err
		}
		if rows <= 0 && dto.ChangeFlag == services.FlagTransferOut {
			// 没有更新任何数据 & 交易类型为出账 => 余额不足
			status = services.TransferedStatusSufficientFunds
			return errors.New("余额不足")
		}
		a.account = *accountDao.GetOne(dto.TradeBody.AccountNo) // 通过交易主体的账号, 获得交易主体账号信息
		a.accountLog.Balance = a.account.Balance                // 交易流水的账户余额
		id, err := accountLogDao.Insert(&a.accountLog)          // 记录交易流水
		if err != nil || id <= 0 {
			status = services.TransferedStatusFailure
			return errors.New("账户流水创建失败")
		}
		return nil
	})
	if err != nil {
		logrus.Error(err)
		status = services.TransferedStatusFailure
	} else {
		status = services.TransferedStatusSuccess
	}

	return status, err
}

//根据账户编号来查询账户信息
func (a *accountDomain) GetAccount(accountNo string) *services.AccountDTO {
	accountDao := AccountDao{}
	var account *Account

	err := base.Tx(func(runner *dbx.TxRunner) error {
		accountDao.runner = runner
		account = accountDao.GetOne(accountNo)
		return nil
	})
	if err != nil {
		return nil
	}
	if account == nil {
		return nil
	}
	return account.ToDTO()
}

//根据用户ID来查询红包账户信息
func (a *accountDomain) GetEnvelopeAccountByUserId(userId string) *services.AccountDTO {
	accountDao := AccountDao{}
	var account *Account

	err := base.Tx(func(runner *dbx.TxRunner) error {
		accountDao.runner = runner
		account = accountDao.GetByUserId(userId, int(services.EnvelopeAccountType))
		return nil
	})
	if err != nil {
		return nil
	}
	if account == nil {
		return nil
	}
	return account.ToDTO()

}

//根据用户ID和账户类型来查询账户信息
func (a *accountDomain) GetAccountByUserIdAndType(userId string, accountType services.AccountType) *services.AccountDTO {
	accountDao := AccountDao{}
	var account *Account

	err := base.Tx(func(runner *dbx.TxRunner) error {
		accountDao.runner = runner
		account = accountDao.GetByUserId(userId, int(accountType))
		return nil
	})
	if err != nil {
		return nil
	}
	if account == nil {
		return nil
	}
	return account.ToDTO()

}

//根据流水ID来查询账户流水
func (a *accountDomain) GetAccountLog(logNo string) *services.AccountLogDTO {
	dao := AccountLogDao{}
	var log *AccountLog
	err := base.Tx(func(runner *dbx.TxRunner) error {
		dao.runner = runner
		log = dao.GetOne(logNo)
		return nil
	})
	if err != nil {
		logrus.Error(err)
		return nil
	}
	if log == nil {
		return nil
	}
	return log.ToDTO()
}

//根据交易编号来查询账户流水
func (a *accountDomain) GetAccountLogByTradeNo(tradeNo string) *services.AccountLogDTO {
	dao := AccountLogDao{}
	var log *AccountLog
	err := base.Tx(func(runner *dbx.TxRunner) error {
		dao.runner = runner
		log = dao.GetByTradeNo(tradeNo)
		return nil
	})
	if err != nil {
		logrus.Error(err)
		return nil
	}
	if log == nil {
		return nil
	}
	return log.ToDTO()
}
