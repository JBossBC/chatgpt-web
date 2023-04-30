package controller

type View struct {
	Code int
	Data map[string]interface{}
	Info string
}

func GetBadResponse(code int, info string) (view View) {
	view = View{}
	view.Code = code
	view.Info = info
	return view
}
func GetSuccessResponse(data map[string]interface{}) (view View) {
	view = View{}
	view.Data = data
	return view
}
