package common

func CheckError(err error) {
	if err == nil {
		return
	}
	panic(err)
}
