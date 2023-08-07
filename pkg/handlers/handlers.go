package handlers

import "github.com/omarabdul3ziz/resper/pkg/codec"

func Handle(msg string) string {
	if msg == "ping" {
		return handlePing()
	} else if msg == "count" {
		return handleCount()
	} else {
		return handleDefault()
	}
}

func handlePing() string {
	msg := codec.Msg{
		Type:  "str",
		Value: "pong",
	}
	res := codec.Encode(msg)
	return res
}

func handleCount() string {
	msg := codec.Msg{
		Type:  "int",
		Value: "123",
	}
	res := codec.Encode(msg)
	return res
}

func handleDefault() string {
	msg := codec.Msg{
		Type:  "err",
		Value: "I don't understand what you mean",
	}
	res := codec.Encode(msg)
	return res
}
