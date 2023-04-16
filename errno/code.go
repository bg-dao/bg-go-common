package errno

var (
	OK                = &Errno{Code: 0, Msg: "成功"}
	ParamError        = &Errno{Code: 1, Msg: "请求参数错误"}
	AuthError         = &Errno{Code: 2, Msg: "登陆状态已失效"}
	AccountOrPwdError = &Errno{Code: 3, Msg: "账户或密码错误"}
	HandleError       = &Errno{Code: 4, Msg: "处理错误"}
	OtherError        = &Errno{Code: 5, Msg: "其他错误"}
)
