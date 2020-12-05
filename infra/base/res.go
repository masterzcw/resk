package base

type ResCode int

const (
	ResCodeOk                   ResCode = 1000 // 正常
	ResCodeValidationError      ResCode = 2000 // 验证错误
	ResCodeRequestParamsError   ResCode = 2100 // 请求参数错误
	ResCodeInnerServerError     ResCode = 5000 // 服务器错误
	ResCodeBizError             ResCode = 6000 // 业务异常
	ResCodeBizTransferedFailure ResCode = 6010 // 转账失败
)

type Res struct {
	Code    ResCode     `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
