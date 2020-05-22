package cadenas

func Insertln(s0, s1 string, pos int) string {
	return s0[:pos] + s1 + s0[pos:]
}
