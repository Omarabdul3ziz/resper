package handlers

func Handle(msg string) string {
	if msg == "ping" {
		return handlePing()
	} else if msg == "hello" {
		return handleHello()
	} else {
		return handleDefault()
	}
}

func handlePing() string {
	return "pong"
}

func handleHello() string {
	return "hello there"
}

func handleDefault() string {
	return "I don't understand what you mean"
}
