package go_web_service

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}
