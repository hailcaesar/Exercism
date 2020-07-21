package raindrops

import (
	"bytes"
	"strconv"
)

//Convert integer to raindrop-like string
func Convert(input int) string {
	var res bytes.Buffer
	if input%3 == 0 {
		res.WriteString("Pling")
	}
	if input%5 == 0 {
		res.WriteString("Plang")
	}
	if input%7 == 0 {
		res.WriteString("Plong")
	}

	if len(res.String()) == 0 {
		r := strconv.Itoa(input)
		res.WriteString(r)
	}

	return res.String()
}
