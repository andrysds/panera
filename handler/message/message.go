package message

// Message is router function for others message handlers
func Message(command string) string {
	switch command {
	case "/birthdays":
		return Birthdays()
	}
	return ""
}
