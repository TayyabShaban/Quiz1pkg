package Quiz1pkg

func PrintCoronaDetails(no int) string {
	if no == 0 {
		return "break"
	} else if no == 1 {
		return "Corona cases in Pakistan = 98,943"
	} else if no == 2 {
		return "Corona cases in South Korea = 11,776"
	} else if no == 3 {
		return "Corona cases in France = 154,000"
	} else if no == 4 {
		return "Hello, Corona!"
	}
	return ""
}
