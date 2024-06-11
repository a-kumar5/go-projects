package encoder

const (
	base         uint64 = 62
	characterSet        = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
)

func ToNumber(s string) uint64 {
	var number uint64
	for _, char := range s {
		number = number*256 + uint64(char)
	}
	return number
}

func ToBase62(num uint64) string {
	encoded := ""
	for num > 0 {
		r := num % base
		num /= base
		encoded = string(characterSet[r]) + encoded
	}
	return encoded
}
