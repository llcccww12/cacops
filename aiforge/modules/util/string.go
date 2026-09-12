package util

func TruncateString(msg string, maxLength int) string {
	if msg == "" {
		return ""
	}
	if len(msg) < maxLength {
		maxLength = len(msg)
	}
	return msg[0:maxLength]
}

func UniqueStr(strs []string) []string {
	appeared := make(map[string]bool)
	var result []string
	for _, str := range strs {
		if !appeared[str] {
			appeared[str] = true
			result = append(result, str)
		}
	}
	return result
}
