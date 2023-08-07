package codec

import (
	"fmt"
	"strconv"
	"strings"
)

// Decode should take a RESP string and return a normal string without RESP encoding.
func Decode(msg string) string {
	switch msg[0] {
	case '+':
		return decodeString(msg[1:])
	case '-':
		return decodeError(msg[1:])
	case ':':
		return decodeInteger(msg[1:])
	case '$':
		return decodeBulk(msg[1:])
	case '*':
		return decodeArray(msg[1:])
	default:
		return "Error: Unknown type"
	}
}

func untilCLRF(msg string) string {
	idx := strings.Index(msg, CRLF)
	if idx == -1 {
		return ""
	}
	return msg[:idx]
}

func decodeString(msg string) string {
	return untilCLRF(msg)
}

func decodeError(msg string) string {
	return untilCLRF(msg)
}

func decodeInteger(msg string) string {
	return untilCLRF(msg)
}

func decodeBulk(msg string) string {
	length := untilCLRF(msg)
	if length == "-1" {
		return `\0\0`
	} else {
		res := msg[len(length)+len(CRLF) : len(msg)-len(CRLF)]
		return res
	}
}

func decodeArray(msg string) string {
	length := untilCLRF(msg)
	if length == "-1" {
		return `\0\0`
	}

	leng, err := strconv.Atoi(length)
	if err != nil {
		return "Error: Invalid length"
	}

	arr := make([]string, leng)

	arrStr := msg[len(length)+len(CRLF) : len(msg)-len(CRLF)]

	for i := 0; i < leng; i++ {
		value := Decode(arrStr)
		arr = append(arr, value)
	}

	fmt.Println(arr)
	return ""

}
