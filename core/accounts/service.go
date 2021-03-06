package accounts

import (
	"fmt"
	"sync"

	"github.com/masterzcw/resk.wen/base"
	"resk.com/services"

	"github.com/kataras/iris/core/errors"
	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"
	"gopkg.in/go-playground/validator.v9"
)

var _ services.AccountService = new(accountService)
var once sync.Once

func init() {
	fmt.Println("/Users/crr/golang/resk.com/core/accounts/service.go->init()")
	/*
		once.Do无论这个函数被调用多少次,只执行一次
		初始化IAccountService
	*/
	once.Do(func() {
		services.IAccountService = new(accountService)
	})
}

type accountService struct {
}

func (a *accountService) CreateAccount(dto services.AccountCreatedDTO) (*services.AccountDTO, error) {
	domain := accountDomain{}
	//验证输入参数
	err := base.Validate().Struct(&dto)
	if err != nil {
		_, ok := err.(*validator.InvalidValidationError)
		if ok {
			logrus.Error("验证错误", err)
		}
		errs, ok := err.(validator.ValidationErrors)
		if ok {
			for _, e := range errs {
				logrus.Error(e.Translate(base.Transtate()))
			}
		}
		return nil, err
	}
	//验证账户是否存在以保证本函数的幂等性
	acc := domain.GetAccountByUserIdAndType(dto.UserId, services.AccountType(dto.AccountType))
	if acc != nil {
		return acc, errors.New(
			fmt.Sprintf("用户的该类型账户已经存在：username=%s[%s],账户类型=%d",
				dto.Username, dto.UserId, dto.AccountType))

	}
	//执行账户创建的业务逻辑
	amount, err := decimal.NewFromString(dto.Amount)
	if err != nil {
		return nil, err
	}
	account := services.AccountDTO{
		UserId:       dto.UserId,
		Username:     dto.Username,
		AccountType:  dto.AccountType,
		AccountName:  dto.AccountName,
		CurrencyCode: dto.CurrencyCode,
		Status:       1,
		Balance:      amount,
	}
	rdto, err := domain.Create(account)
	return rdto, err
}

// 转账
func (a *accountService) Transfer(dto services.AccountTransferDTO) (services.TransferedStatus, error) {
	//验证参数
	domain := accountDomain{}
	err := base.Validate().Struct(&dto)
	if err != nil {
		_, ok := err.(*validator.InvalidValidationError)
		if ok {
			logrus.Error("验证错误", err)
		}
		errs, ok := err.(validator.ValidationErrors)
		if ok {
			for _, e := range errs {
				logrus.Error(e.Translate(base.Transtate()))
			}
		}
		return services.TransferedStatusFailure, err
	}
	//执行转账逻辑
	amount, err := decimal.NewFromString(dto.AmountStr)
	if err != nil {
		return services.TransferedStatusFailure, err
	}
	dto.Amount = amount
	if dto.ChangeFlag == services.FlagTransferOut {
		if dto.ChangeType > 0 {
			return services.TransferedStatusFailure,
				errors.New("如果changeFlag为支出，那么changeType必须小于0")
		}
	} else {
		if dto.ChangeType < 0 {
			return services.TransferedStatusFailure,
				errors.New("如果changeFlag为收入,那么changeType必须大于0")
		}
	}

	status, err := domain.Transfer(dto)
	// if status == services.TransferedStatusSuccess { // 倒退的
	// 	backwardDto := dto
	// 	backwardDto.TradeBody = dto.TradeTarget
	// 	backwardDto.TradeTarget = dto.TradeBody
	// 	backwardDto.ChangeType = -dto.ChangeType
	// 	backwardDto.ChangeFlag = -dto.ChangeFlag
	// 	status, err := domain.Transfer(backwardDto)
	// 	return status, err
	// }
	return status, err
}

// 储值
func (a *accountService) StoreValue(dto services.AccountTransferDTO) (services.TransferedStatus, error) {
	dto.TradeTarget = dto.TradeBody
	dto.ChangeFlag = services.FlagTransferIn
	dto.ChangeType = services.AccountStoreValue
	return a.Transfer(dto)
}

func (a *accountService) GetEnvelopeAccountByUserId(userId string) *services.AccountDTO {
	domain := accountDomain{}
	account := domain.GetEnvelopeAccountByUserId(userId)
	return account
}

func (a *accountService) GetAccount(accountNo string) *services.AccountDTO {
	domain := accountDomain{}
	return domain.GetAccount(accountNo)
}
