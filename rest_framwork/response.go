package rest_framwork

/*
	
*/

type ResultData struct {
	Code    int         `json:"code"`    // 业务code
	Message string      `json:"message"` //返回的提示消息
	Data    interface{} `json:"data"`    // json响应数据
}

// 封装http请求统一返回的数据结构
func ResponseResult(data interface{}, code int, msg string) ResultData {
	return ResultData{
		Code:    code,
		Message: msg,
		Data:    data,
	}
}
