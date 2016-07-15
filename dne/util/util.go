package util

//Error Util
func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}

//String Util
func FillBefore(inicialStr, strToFillWith string, lenght int) (finalStr string) {
	finalStr = inicialStr
	for len(finalStr) < lenght {
		finalStr = strToFillWith + finalStr
	}
	return
}
