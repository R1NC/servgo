package go_web_api

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}
