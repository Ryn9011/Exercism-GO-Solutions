package secret

func Handshake(i uint) (code []string) {

	if i&1 == 1 {
		code = append(code, "wink")
	}
	if i&2 == 2 {
		code = append(code, "double blink")
	}
	if i&4 == 4 {
		code = append(code, "close your eyes")
	}
	if i&8 == 8 {
		code = append(code, "jump")
	}

	if i&16 == 16 {
		for i := 0; i < len(code)/2; i++ {
			j := len(code) - i - 1
			code[i], code[j] = code[j], code[i]
		}
	}

	return

}
