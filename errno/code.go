package errno

var (
	OK                = &Errno{Code: 0, Msg: "success"}
	ParamError        = &Errno{Code: 1, Msg: "param error"}
	AuthError         = &Errno{Code: 2, Msg: "auth error"}
	AccountOrPwdError = &Errno{Code: 3, Msg: "account or pwd error"}
	HandleError       = &Errno{Code: 4, Msg: "handle error"}
	OtherError        = &Errno{Code: 5, Msg: "other error"}
)
