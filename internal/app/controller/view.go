package controller

type View struct {
	Result string
	Data   map[string]interface{}
	Info   string
}

func GetBadResponse(info string) (view View) {
	view = View{}
	view.Result = "failed"
	view.Info = info
	return view
}
func GetSuccessResponse(data map[string]interface{}) (view View) {
	view.Result = "success"
	view = View{}
	view.Data = data
	return view
}
