package error

type XKSJHTError struct {
	msg string
}

func (e *XKSJHTError) Error() string {
	return e.msg
}

func New(msg string) *XKSJHTError {
	return &XKSJHTError{msg: msg}
}

func DBError() *XKSJHTError {
	return New("数据库错误")
}

func AuthError() *XKSJHTError {
	return New("认证失败")
}

func RestError() *XKSJHTError {
	return New("REST交互失败")
}

func HandleRequestError() *XKSJHTError {
	return New("服务异常")
}
