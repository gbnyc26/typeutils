package typeutils

func BoolPtr(b bool) *bool {
	return &b
}

func BoolYNString(b bool) string {
	if !b {
		return "N"
	}
	return "Y"
}
