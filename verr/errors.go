package verr

var (
	ErrInternal = New(10000, "内部错误")
	ErrUpstream = New(10001, "上游服务错误")
)