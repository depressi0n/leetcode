package question

func addBinary(a string, b string) string {
	var c, d []byte
	//看哪一个长
	if len(a) < len(b) {
		c = []byte(b)
		d = []byte(a)
	} else {
		c = []byte(a)
		d = []byte(b)
	}
	b = ""
	for i := 0; i < len(c); i++ {
		if i < len(c)-len(d) {
			b += "0"
		} else {
			b += string(d[i-len(c)+len(d)])
		}
	}
	d = []byte(b)
	carry := 0
	for i := len(c) - 1; i >= 0; i-- {
		switch {
		case c[i] == '1' && d[i] == '1':
			if carry == 1 {
				c[i] = '1'
			} else {
				c[i] = '0'
			}
			carry = 1
		case c[i] == '1' || d[i] == '1':
			if carry == 1 {
				c[i] = '0'
			} else {
				carry = 0
				c[i] = '1'
			}
		default:
			if carry == 1 {
				c[i] = '1'
			} else {
				c[i] = '0'
			}
			carry = 0
		}
	}
	if carry == 1 {
		c = append([]byte{'1'}, c...)
	}
	return string(c)
}
