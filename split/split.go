package piscine

func Split(s, sep string) (res []string) {
	slen := len(s)
	seplen := len(sep)
	word := ""
	for i := 0; i < slen; i++ {
		if i+seplen <= slen && s[i:i+seplen] == sep {
			res = append(res, word)
			i += seplen-1
			word = ""
		} else {
			word += string(s[i])
		}
	}
	if word != "" {
		res = append(res, word)
	}
	return res
}
