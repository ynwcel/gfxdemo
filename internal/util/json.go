package util

const (
	JSON_OK_CODE   = 200
	JSON_FAIL_CODE = 0
)

func JsonOk(data any, msg ...string) map[string]any {
	json_msg := "ok"
	if len(msg) > 0 && len(msg[0]) > 0 {
		json_msg = msg[0]
	}
	return Json(JSON_OK_CODE, data, json_msg)
}

func JsonFail(data any, msg ...string) map[string]any {
	json_msg := "ok"
	if len(msg) > 0 && len(msg[0]) > 0 {
		json_msg = msg[0]
	}
	return Json(JSON_FAIL_CODE, data, json_msg)
}

func Json(code int, data any, msg string) map[string]any {
	return map[string]any{
		"code": code,
		"data": data,
		"msg":  msg,
	}
}
