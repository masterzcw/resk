package main

import (
	"net/rpc"

	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"
	"resk.com/services"
)

func main() {
	c, err := rpc.Dial("tcp", ":18082")
	if err != nil {
		logrus.Panic(err)
	}
	sendout(c)
	receive(c)

}

func receive(c *rpc.Client) {
	in := services.RedEnvelopeReceiveDTO{
		EnvelopeNo:   "",
		RecvUserId:   "",
		RecvUsername: "",
		AccountNo:    "",
	}
	out := &services.RedEnvelopeItemDTO{}
	err := c.Call("Envelope.Receive", in, out)
	if err != nil {
		logrus.Panic(err)
	}
	logrus.Infof("%+v", out)
}
func sendout(c *rpc.Client) {
	in := services.RedEnvelopeSendingDTO{
		Amount:       decimal.NewFromFloat(1),
		UserId:       "47692588035919872",
		Username:     "测试用户",
		EnvelopeType: services.GeneralEnvelopeType,
		Quantity:     2,
		Blessing:     "",
	}
	out := &services.RedEnvelopeActivity{}
	err := c.Call("EnvelopeRpc.SendOut", in, &out)
	if err != nil {
		logrus.Panic(err)
	}
	logrus.Infof("%+v", out)
}

/*
<masterzcw@/users/crr/golang/resk.com/example/rpcclient>
$ go run rpc_client.go
[2020-12-08 12:01:05]  INFO &{RedEnvelopeGoodsDTO:{EnvelopeNo: EnvelopeType:0 Username: UserId: Blessing: Amount:0 AmountOne:0 Quantity:0 RemainAmount:0 RemainQuantity:0 ExpiredAt:0001-01-01 00:00:00 +0000 UTC Status:0 OrderType:0 PayStatus:0 CreatedAt:0001-01-01 00:00:00 +0000 UTC UpdatedAt:0001-01-01 00:00:00 +0000 UTC AccountNo: OriginEnvelopeNo:} Link:}
*/
