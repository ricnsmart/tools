package util

import "crypto/rand"

var AsciiChars = []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789!@#$%^&*()-_=+,.?/:;{}[]`~")

func NewLen(length int) string {
	return NewLenChars(length, AsciiChars)
}

// NewLenChars returns a new random string of the provided length, consisting of the provided byte slice of allowed characters(maximum 256).
func NewLenChars(length int, chars []byte) string {
	if length == 0 {
		return ""
	}
	clen := len(chars)
	if clen < 2 || clen > 256 {
		panic("Wrong charset length for NewLenChars()")
	}
	maxrb := 255 - (256 % clen)
	b := make([]byte, length)
	r := make([]byte, length+(length/4)) // storage for random bytes.
	i := 0
	for {
		if _, err := rand.Read(r); err != nil {
			panic("Error reading random bytes: " + err.Error())
		}
		for _, rb := range r {
			c := int(rb)
			if c > maxrb {
				continue // Skip this number to avoid modulo bias.
			}
			b[i] = chars[c%clen]
			i++
			if i == length {
				return string(b)
			}
		}
	}
}
