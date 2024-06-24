package utils

func RemoveLastRune(s string, count int) string {
	r := []rune(s)
	return string(r[:len(r)-count])
}
