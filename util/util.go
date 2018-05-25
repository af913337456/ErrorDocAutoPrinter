package util

func GetCommonErr(errorCode int,info interface{}) map[string]interface{} {
	d := map[string]interface{}{
		"ret": "failed",
		"code":errorCode,
		"msg" : info,
		"data":nil,
	}
	return d
}

func GetErrWithTips(errorCode int,info interface{},tips string) map[string]interface{} {
	d := map[string]interface{}{
		"ret": "failed",
		"code":errorCode,
		"msg" : info,
		"data":nil,
	}
	return d
}