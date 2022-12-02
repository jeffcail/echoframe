package global

// CheckErr
func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}
