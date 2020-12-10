package services

const (
	DefaultBlessing   = "恭喜发财！"
	DefaultTimeFormat = "2006-01-02.15:04:05"
)

//订单类型：发布单、退款单
type OrderType int

const (
	OrderTypeSending OrderType = 1 // 发红包
	OrderTypeRefund  OrderType = 2 // 收红包
)

//支付状态：(未)支付，支付(中)，(已)支付，支付(失败)
//退款：(未)退款，退款(中)，(已)退款，退款(失败)
type PayStatus int

const (
	PayNothing PayStatus = 1 // 未
	Paying     PayStatus = 2 // 中
	Payed      PayStatus = 3 // 已
	PayFailure PayStatus = 4 // 失败
	//
	RefundNothing PayStatus = 61
	Refunding     PayStatus = 62
	Refunded      PayStatus = 63
	RefundFailure PayStatus = 64
)

//红包订单状态：创建、发布、过期、失效
type OrderStatus int

const (
	OrderCreate                  OrderStatus = 1 // 创建
	OrderSending                 OrderStatus = 2 // 发布
	OrderExpired                 OrderStatus = 3 // 过期
	OrderDisabled                OrderStatus = 4 // 失效
	OrderExpiredRefundSuccessful OrderStatus = 5 // 退款成功
	OrderExpiredRefundFalured    OrderStatus = 6 // 退款失败
)

//红包类型：普通红包，碰运气红包
type EnvelopeType int

const (
	GeneralEnvelopeType = 1 // 普通红包
	LuckyEnvelopeType   = 2 // 运气红包
)

var EnvelopeTypes = map[EnvelopeType]string{
	GeneralEnvelopeType: "普通红包",
	LuckyEnvelopeType:   "碰运气红包",
}
