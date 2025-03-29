package modpow

// Modpow calculate a^b % mod.
func Modpow(a, b, mod int) int {
	ret := 1
	i := a
	for b > 0 {
		if b&1 == 1 {
			ret = (ret * i) % mod
		}
		b = b >> 1
		i = (i * i) % mod
	}
	return ret
}
