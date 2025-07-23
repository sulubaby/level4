package piscine

func ZipString(s string) string {
	result := ""
	for i := 0; i < len(s); i++ {
		count := 1
		if s[i] == '0' || (s[i] < 'A') || (s[i] > 'z') || (s[i] > 'Z' && s[i] < 'a') {
			continue
		}
		for j := i + 1; j < len(s); j++ {
			if s[i] == s[j] {
				count++
				s = s[:j] + "0" + s[j+1:] // Mark duplicate as '0'
			}
		}
		result += string('0'+count) + string(s[i])
	}
	return result
}
