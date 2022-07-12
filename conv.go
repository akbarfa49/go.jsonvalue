package jsonvalue

import (
	"bytes"
	"fmt"
	"reflect"
)

func formatBool(b bool) string {
	if b {
		return "true"
	}
	return "false"
}

// reference:
// - [UTF-16](https://zh.wikipedia.org/zh-cn/UTF-16)
// - [JavaScript has a Unicode problem](https://mathiasbynens.be/notes/javascript-unicode)
// - [Meaning of escaped unicode characters in JSON](https://stackoverflow.com/questions/21995410/meaning-of-escaped-unicode-characters-in-json)
func escapeGreaterUnicodeToBuffByUTF16(r rune, buf *bytes.Buffer) {
	if r <= '\uffff' {
		buf.WriteString(fmt.Sprintf("\\u%04X", r))
		return
	}
	// if r > 0x10FFFF {
	// 	// invalid unicode
	// 	buf.WriteRune(r)
	// 	return
	// }

	r = r - 0x10000
	lo := r & 0x003FF
	hi := (r & 0xFFC00) >> 10
	buf.WriteString(fmt.Sprintf("\\u%04X", hi+0xD800))
	buf.WriteString(fmt.Sprintf("\\u%04X", lo+0xDC00))
}

func escapeGreaterUnicodeToBuffByUTF8(r rune, buf *bytes.Buffer) {
	buf.WriteRune(r)
}

func escapeNothing(b byte, buf *bytes.Buffer) {
	buf.WriteByte(b)
}

func escDoubleQuote(_ byte, buf *bytes.Buffer) {
	buf.Write([]byte{'\\', '"'})
}

func escSlash(_ byte, buf *bytes.Buffer) {
	buf.Write([]byte{'\\', '/'})
}

func escBaskslash(_ byte, buf *bytes.Buffer) {
	buf.Write([]byte{'\\', '\\'})
}

func escBaskspace(_ byte, buf *bytes.Buffer) {
	buf.Write([]byte{'\\', 'b'})
}

func escVertTab(_ byte, buf *bytes.Buffer) {
	buf.Write([]byte{'\\', 'f'})
}

func escTab(_ byte, buf *bytes.Buffer) {
	buf.Write([]byte{'\\', 't'})
}

func escNewLine(_ byte, buf *bytes.Buffer) {
	buf.Write([]byte{'\\', 'n'})
}

func escReturn(_ byte, buf *bytes.Buffer) {
	buf.Write([]byte{'\\', 'r'})
}

func escLeftAngle(_ byte, buf *bytes.Buffer) {
	buf.Write([]byte{'\\', 'u', '0', '0', '3', 'C'})
}

func escRightAngle(_ byte, buf *bytes.Buffer) {
	buf.Write([]byte{'\\', 'u', '0', '0', '3', 'E'})
}

func escAnd(_ byte, buf *bytes.Buffer) {
	buf.Write([]byte{'\\', 'u', '0', '0', '2', '6'})
}

// func escPercent(_ byte, buf *bytes.Buffer) {
// 	buf.Write([]byte{'\\', 'u', '0', '0', '2', '5'})
// }

func escapeStringToBuff(s string, buf *bytes.Buffer, opt *Opt) {
	for _, r := range s {
		if r <= 0x7F {
			b := byte(r)
			opt.asciiCharEscapingFunc[b](b, buf)
		} else {
			opt.unicodeEscapingFunc(r, buf)
		}
	}
}

func intfToInt(v interface{}) (u int, err error) {
	switch v := v.(type) {
	case int:
		u = v
	case uint:
		u = int(v)
	case int64:
		u = int(v)
	case uint64:
		u = int(v)
	case int32:
		u = int(v)
	case uint32:
		u = int(v)
	case int16:
		u = int(v)
	case uint16:
		u = int(v)
	case int8:
		u = int(v)
	case uint8:
		u = int(v)
	default:
		err = fmt.Errorf("%s is not a number", reflect.TypeOf(v).String())
	}

	return
}

// func intfToInt64(v interface{}) (i int64, err error) {
// 	switch v.(type) {
// 	case int:
// 		i = int64(v.(int))
// 	case uint:
// 		i = int64(v.(uint))
// 	case int64:
// 		i = int64(v.(int64))
// 	case uint64:
// 		i = int64(v.(uint64))
// 	case int32:
// 		i = int64(v.(int32))
// 	case uint32:
// 		i = int64(v.(uint32))
// 	case int16:
// 		i = int64(v.(int16))
// 	case uint16:
// 		i = int64(v.(uint16))
// 	case int8:
// 		i = int64(v.(int8))
// 	case uint8:
// 		i = int64(v.(uint8))
// 	default:
// 		err = fmt.Errorf("%s is not a number", reflect.TypeOf(v).String())
// 	}

// 	return
// }

func intfToString(v interface{}) (s string, err error) {
	switch str := v.(type) {
	case string:
		return str, nil
	default:
		return "", fmt.Errorf("%s is not a string", reflect.TypeOf(v).String())
	}
}

// func intfToJsonvalue(v interface{}) (j *V, err error) {
// 	switch v.(type) {
// 	case *V:
// 		j = v.(*V)
// 	default:
// 		err = fmt.Errorf("%s is not a *jsonvalue.V type", reflect.TypeOf(v).String())
// 	}

// 	return
// }
