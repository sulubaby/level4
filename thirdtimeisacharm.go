package piscine

func ThirdTimeIsACharm(str string) string {
	res := ""
	if len(str) < 3 {
		return "\n"
	}
	for i := 2; i < len(str); i = i + 3 {
		res += string(str[i])
	}
	return res + "\n"
}
