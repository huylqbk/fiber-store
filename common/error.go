package common

func CheckError(err error) {
	if err == nil {
		return
	}
	NewResponse(400, err.Error(), nil, 0)
}
