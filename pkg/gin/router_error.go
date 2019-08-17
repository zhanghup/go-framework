package gin

type exception struct {
	error
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func (this exception) Error() string {
	return this.Message
}
func NewErr(msg string, res ...interface{}) error {
	e := exception{
		Code:    9998,
		Message: msg,
	}
	if len(res) == 0 {
		e.Data = map[string]string{}
	} else if len(res) == 1 {
		e.Data = res[0]
	} else {
		e.Data = map[string]interface{}{
			"result": res,
		}
	}
	return e
}
func NewError(code int, msg string, res ...interface{}) error {
	e := exception{
		Code:    code,
		Message: msg,
	}
	if len(res) == 0 {
		e.Data = map[string]string{}
	} else if len(res) == 1 {
		e.Data = res[0]
	} else {
		e.Data = map[string]interface{}{
			"result": res,
		}
	}
	return e
}
