package strings

const Mask = "******"

func MaskNotEmpty(s string) string {
	if s != "" {
		s = Mask
	}

	return s
}
