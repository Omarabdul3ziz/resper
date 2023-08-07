package codec

import "fmt"

const (
	// it is important to use backticks here. to not escape the \r and \n
	CRLF = `\r\n`
)

type Msg struct {
	Type  string
	Value string
}

func Encode(msg Msg) string {
	switch msg.Type {
	case "str":
		return encodeString(msg)
	case "err":
		return encodeError(msg)
	case "int":
		return encodeInteger(msg)
	case "bulk":
		return encodeBulk(msg)
	case "arr":
		return encodeArray(msg)
	default:
		return "Error: Unknown type"
	}
}

func encodeString(msg Msg) string {
	return fmt.Sprintf("+%s%s", msg.Value, CRLF)
}

func encodeError(msg Msg) string {
	return fmt.Sprintf("-%s%s", msg.Value, CRLF)
}

func encodeInteger(msg Msg) string {
	return fmt.Sprintf(":%s%s", msg.Value, CRLF)
}

func encodeBulk(msg Msg) string {
	return fmt.Sprintf(`$%d%s%s%s`, len(msg.Value), CRLF, msg.Value, CRLF)
}

func encodeArray(mag Msg) string {
	res := fmt.Sprintf("*%d%s", len(mag.Value), CRLF)
	for _, v := range mag.Value {
		res += string(v)
	}
	return res
}
