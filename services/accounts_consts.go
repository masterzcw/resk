package services

// 转账状态
type TransferedStatus int8

const (
	// 转账失败
	TransferedStatusFailure TransferedStatus = -1
	// 余额不足
	TransferedStatusSufficientFunds TransferedStatus = 0
	// 转账成功
	TransferedStatusSuccess TransferedStatus = 1
)

// 转账的类型 0-创建; 正数-进账; 负数-支出
type ChangeType int8

const (
	// 账户创建
	AccountCreated ChangeType = 0
	// 储值
	AccountStorValue ChangeType = 1
	// 红包资金的支出
	EnvelopeOutgoing ChangeType = -2
	// 红包资金的收入
	EnvelopeIncoming ChangeType = 2
	// 红包过期退款
	EnvelopeExpiredRefund ChangeType = 3
)

// 资金交易变化标识
type ChangeFlag int8

const (
	// 创建账户
	FlagAccountCreated ChangeFlag = 0
	// 支出
	FlagTransferOut ChangeFlag = -1
	// 收入
	FlagTransferIn ChangeFlag = 1
)