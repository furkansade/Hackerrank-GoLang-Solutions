package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	y := ScanInt(1700, 2700)
	mlen := []int{0, 31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	if y == 1918 {
		mlen[2] = 15
	} else if y%4 == 0 {
		if y < 1918 || y%100 != 0 || y%400 == 0 {
			mlen[2] = 29
		}
	}
	m := 1
	d := 256
	for mlen[m] < d {
		d -= mlen[m]
		m++
	}
	fmt.Printf("%02d.%02d.%4d\n", d, m, y)
}

// Boilerplate

func Assert(condition bool, items ...interface{}) {
	if !condition {
		panic("assertion failed: " + fmt.Sprint(items...))
	}
}

func Log(items ...interface{}) {
	fmt.Println(items...)
}

var Input = bufio.NewReader(os.Stdin)

func ReadByte() byte {
	b, e := Input.ReadByte()
	if e != nil {
		panic(e)
	}
	return b
}

func MaybeReadByte() (byte, bool) {
	b, e := Input.ReadByte()
	if e != nil {
		if e == io.EOF {
			return 0, false
		}
		panic(e)
	}
	return b, true
}

func UnreadByte() {
	e := Input.UnreadByte()
	if e != nil {
		panic(e)
	}
}

func SkipWhitespace() byte {
	for {
		b := ReadByte()
		if b != ' ' && b != '\t' && b != '\r' {
			return b
		}
	}
}

func NewLine() {
	for {
		b := ReadByte()
		switch b {
		case ' ', '\t', '\r':
			// keep looking
		case '\n':
			return
		default:
			panic(fmt.Sprintf("expecting newline, but found character <%c>", b))
		}
	}
}

func ScanInt(low, high int) int {
	return int(ScanInt64(int64(low), int64(high)))
}

func ScanUint(low, high uint) uint {
	return uint(ScanUint64(uint64(low), uint64(high)))
}

func ScanInt64(low, high int64) int64 {
	Assert(low <= high)
	b := SkipWhitespace()
	switch b {
	case '\n':
		panic(fmt.Sprintf(
			"unexpected newline; expecting range %d..%d", low, high))
	case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		if high < 0 {
			panic(fmt.Sprintf(
				"found <%c> but expecting range %d..%d", b, low, high))
		}
		lw := low
		if lw < 0 {
			lw = 0
		}
		UnreadByte()
		x, e := _scanu64(uint64(lw), uint64(high))
		if e != "" {
			panic(fmt.Sprintf("%s %d..%d", e, low, high))
		}
		return int64(x)
	case '-':
		if low > 0 {
			panic(fmt.Sprintf(
				"found minus sign but expecting range %d..%d", low, high))
		}
		h := high
		if h > 0 {
			h = 0
		}
		x, e := _scanu64(uint64(-h), uint64(-low))
		if e != "" {
			panic(fmt.Sprintf("-%s %d..%d", e, low, high))
		}
		return -int64(x)
	default:
		panic(fmt.Sprintf(
			"unexpected character <%c>; expecting range %d..%d", b, low, high))
	}
}

func ScanUint64(low, high uint64) uint64 {
	Assert(low <= high)
	b := SkipWhitespace()
	switch b {
	case '\n':
		panic(fmt.Sprintf(
			"unexpected newline; expecting range %d..%d", low, high))
	case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		UnreadByte()
		x, e := _scanu64(low, high)
		if e != "" {
			panic(fmt.Sprintf("%s %d..%d", e, low, high))
		}
		return x
	default:
		panic(fmt.Sprintf(
			"unexpected character <%c>; expecting range %d..%d", b, low, high))
	}
}

func _scanu64(low, high uint64) (result uint64, err string) {
	x := uint64(0)
buildnumber:
	for {
		b, ok := MaybeReadByte()
		if !ok {
			break buildnumber
		}
		switch b {
		case ' ', '\t', '\r':
			break buildnumber
		case '\n':
			UnreadByte()
			break buildnumber
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			d := uint64(b - '0')
			if (high-d)/10 < x {
				return x, fmt.Sprintf("%d%c... not in range", x, b)
			}
			x = x*10 + d
		default:
			return x, fmt.Sprintf("%d%c found; expecting range", x, b)
		}
	}
	if x < low || x > high {
		return x, fmt.Sprintf("%d not in range", x)
	}
	return x, ""
}

func ScanBuf(buf []byte) []byte {
	b := SkipWhitespace()
	if b == '\n' {
		panic(fmt.Sprintf("unexpected newline; expecting string"))
	}
	buf[0] = b
	length := 1
buildstring:
	for {
		var ok bool
		b, ok = MaybeReadByte()
		if !ok {
			break buildstring
		}
		switch b {
		case ' ', '\t', '\r':
			break buildstring
		case '\n':
			UnreadByte()
			break buildstring
		default:
			if length >= len(buf) {
				panic(fmt.Sprintf("string longer than %d bytes", len(buf)))
			}
			buf[length] = b
			length++
		}
	}
	return buf[:length]
}

func ScanBytes(short, long int) []byte {
	Assert(1 <= short && short <= long)
	buf := make([]byte, long)
	buf = ScanBuf(buf)
	if len(buf) < short {
		panic(fmt.Sprintf("string shorter than %d bytes", short))
	}
	return buf
}

func ScanString(short, long int) string {
	return string(ScanBytes(short, long))
}
