package accounts

import (
	"github.com/sirupsen/logrus"
	"github.com/tietang/dbx"
)

type AccountLogDao struct {
	runner *dbx.TxRunner
}

// 通过流水编号查询流水记录
func (dao *AccountLogDao) GetOne(logNo string) *AccountLog {
	a := &AccountLog{
		LogNo: logNo,
	}
	ok, err := dao.runner.GetOne(a) // 通过唯一索引查询
	if err != nil {
		logrus.Error(err)
		return nil
	}
	if !ok {
		return nil
	}
	return a
}

// 通过交易编号来查询流水记录
func (dao *AccountLogDao) GetByTradeNo(TradeNo string) *AccountLog {
	al := &AccountLog{}
	sql := `select * from account_log where trade_no=?`
	ok, err := dao.runner.Get(al, sql, TradeNo)
	if err != nil {
		logrus.Error(err)
		return nil
	}
	if !ok {
		return nil
	}
	return al
}

// 流水记录的写入
func (dao *AccountLogDao) Insert(l *AccountLog) (id int64, err error) {
	rs, err := dao.runner.Insert(l)
	if err != nil {
		return 0, err
	}
	return rs.LastInsertId()
}
