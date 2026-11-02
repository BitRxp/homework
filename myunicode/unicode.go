package myunicode

func IsDigit(r rune) bool {
	if r >= 127 {
		return false
	}
	if r >= '0' && r <= '9' {
		return true
	}
	return false
}

func IsLetter(r rune) bool {
	if r >= 127 {
		return false
	}
	if (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') {
		return true
	}
	return false
}

func IsSpace(r rune) bool {
	if r >= 127 {
		return false
	}
	if r == ' ' {
		return true
	}
	return false
}

func IsUpper(r rune) bool {
	if r >= 127 {
		return false
	}
	if r >= 'A' && r <= 'Z' {
		return true
	}
	return false
}

func IsLower(r rune) bool {
	if r >= 127 {
		return false
	}
	if r >= 'a' && r <= 'z' {
		return true
	}
	return false
}

func IsPunct(r rune) bool {
	if r >= 127 {
		return false
	}
	if (r >= 'A' && r <= 'Z') || (r >= 'a' && r <= 'z') || (r >= '0' && r <= '9') || (r == ' ') {
		return false
	}
	return true
}

func IsControl(r rune) bool {
	if r > 127 {
		return false
	}
	if (r >= 0 && r <= 31) || r == 127 {
		return true
	}
	return false
}

func ToUpper(r rune) rune {
	if r >= 'a' || r <= 'z' {
		return r - 32
	}
	return r
}

func ToLower(r rune) rune {
	if r >= 'A' || r <= 'Z' {
		return r + 32
	}
	return r
}
