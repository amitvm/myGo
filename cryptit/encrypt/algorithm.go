package encrypt

func Nimbus(str string) string {
	encyptedStr := ""
	for _, c := range str {
		asciiCode := int(c)
		character := string(asciiCode + 3)
		encyptedStr += character
	}
	return encyptedStr
}
