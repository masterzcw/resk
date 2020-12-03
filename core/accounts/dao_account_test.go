package accounts

import (
	"database/sql"
	"testing"

	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"
	"github.com/tietang/dbx"

	_ "resk/testx"

	"resk/infra/base"

	"github.com/segmentio/ksuid"
	"github.com/smartystreets/goconvey/convey"
	. "github.com/smartystreets/goconvey/convey"
)

func TestAccountDao_GetOne(t *testing.T) {
	err := base.Tx(func(runner *dbx.TxRunner) error {
		dao := &AccountDao{
			runner: runner,
		}
		Convey("通过编号查询账户数据", t, func() {
			a := &Account{
				Balance:     decimal.NewFromFloat(100),
				Status:      1,
				AccountNo:   ksuid.New().Next().String(),
				AccountName: "测试资金账户",
				UserId:      ksuid.New().Next().String(),
				Username:    sql.NullString{String: "测试用户", Valid: true},
			}
			id, err := dao.Insert(a)
			So(err, ShouldBeNil)
			So(id, ShouldBeGreaterThan, 0)
			na := dao.GetOne(a.AccountNo)
			So(na, ShouldNotBeNil)
			So(na.Balance.String(), ShouldEqual, a.Balance.String())
			So(na.CreatedAt, ShouldNotBeNil)
			So(na.UpdatedAt, ShouldNotBeNil)
		})
		return nil
	})
	if err != nil {
		logrus.Error(err)
	}
}

func TestAccountDao_GetByUserId(t *testing.T) {
	err := base.Tx(func(runner *dbx.TxRunner) error {
		dao := &AccountDao{
			runner: runner,
		}
		convey.Convey("通过用户ID和账户类型查询账户数据", t, func() {
			a := &Account{
				Balance:     decimal.NewFromFloat(100),
				Status:      1,
				AccountNo:   ksuid.New().Next().String(),
				AccountName: "测试资金账户1",
				UserId:      ksuid.New().Next().String(),
				Username:    sql.NullString{String: "测试用户", Valid: true},
				AccountType: 2,
			}
			id, err := dao.Insert(a)
			So(err, ShouldBeNil)
			So(id, ShouldBeGreaterThan, 0)
			na := dao.GetByUserId(a.UserId, a.AccountType)
			So(na, ShouldNotBeNil)
			So(na.Balance.String(), ShouldEqual, a.Balance.String())
			So(na.CreatedAt, ShouldNotBeNil)
			So(na.UpdatedAt, ShouldNotBeNil)
		})
		return nil
	})
	if err != nil {
		logrus.Error(err)
	}
}

func TestAccountDao_UpdateBalance(t *testing.T) {
	err := base.Tx(func(runner *dbx.TxRunner) error {
		dao := &AccountDao{
			runner: runner,
		}
		balance := decimal.NewFromFloat(100)
		convey.Convey("更新账户余额", t, func() {
			a := &Account{
				Balance:     balance,
				Status:      1,
				AccountNo:   ksuid.New().Next().String(),
				AccountName: "测试资金账户1",
				UserId:      ksuid.New().Next().String(),
				Username:    sql.NullString{String: "测试用户", Valid: true},
				AccountType: 2,
			}
			id, err := dao.Insert(a)
			So(err, ShouldBeNil)
			So(id, ShouldBeGreaterThan, 0)

			// 增加
			Convey("增加余额", func() {
				amount := decimal.NewFromFloat(10)
				rows, err := dao.UpdateBalance(a.AccountNo, amount)
				So(err, ShouldBeNil)     // 报错为空
				So(rows, ShouldEqual, 1) // 影响行数应该是一行

				na := dao.GetOne(a.AccountNo)
				newBalance := balance.Add(amount) // 新余额应该等于原有余额+增加余额
				So(na, ShouldNotBeNil)
				So(na.Balance.String(), ShouldEqual, newBalance.String())
			})

			// 扣减, 余额足够
			Convey("扣减, 余额足够", func() {
				amount := decimal.NewFromFloat(-10)
				rows, err := dao.UpdateBalance(a.AccountNo, amount)
				So(err, ShouldBeNil)     // 报错为空
				So(rows, ShouldEqual, 1) // 影响行数应该是一行

				na := dao.GetOne(a.AccountNo)
				newBalance := balance.Add(amount) // 新余额应该等于原有余额+增加余额
				So(na, ShouldNotBeNil)
				So(na.Balance.String(), ShouldEqual, newBalance.String())
			})

			// 扣减, 余额不够
			Convey("扣减, 余额不够", func() {

				amount := decimal.NewFromFloat(-110)
				naBefor := dao.GetOne(a.AccountNo)
				rows, err := dao.UpdateBalance(a.AccountNo, amount)
				naAfter := dao.GetOne(a.AccountNo)
				So(err, ShouldBeNil)     // 报错为空
				So(rows, ShouldEqual, 0) // 影响行数应该是零行
				So(naBefor, ShouldNotBeNil)
				So(naAfter, ShouldNotBeNil)
				So(naBefor.Balance.String(), ShouldEqual, naAfter.Balance.String())
			})
		})
		return nil
	})
	if err != nil {
		logrus.Error(err)
	}
}
