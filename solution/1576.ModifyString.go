package solution

func switch2abc(l, r byte) byte {
	if l == r {
		if l == 'a' {
			return 'b'
		} else {
			return 'a'
		}
	} else {
		switch l {
		case 'a':
			switch r {
			case 'b':
				return 'c'
			default:
				return 'b'
			}
		case 'b':
			switch r {
			case 'a':
				return 'c'
			default:
				return 'a'
			}
		default:
			switch r {
			case 'a':
				return 'b'
			default:
				return 'a'
			}
		}
	}
}

func ModifyString(s string) string {
	if len(s) == 1 {
		if s[0] == '?' {
			return "a"
		} else {
			return s
		}
	}
	res := []byte(s)
	l := len(s) - 1
	for i := 0; i <= l; i++ {
		if res[i] != '?' {
			continue
		}
		if i == 0 {
			res[i] = switch2abc('0', res[i+1])
		} else if i == l {
			res[i] = switch2abc('0', res[i-1])
		} else {
			res[i] = switch2abc(res[i-1], res[i+1])
		}
	}
	return string(res)
}
