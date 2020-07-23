package raindrops

import (
	"bytes"
	"strconv"
)

//Convert integer to raindrop-like string
func Convert(input int) string {
	var res string
	if input%3 == 0 {
		res += "Pling"
	}
	if input%5 == 0 {
		res += "Plang"
	}
	if input%7 == 0 {
		res += "Plong"
	}

	if len(res) == 0 {
		res += strconv.Itoa(input)
	}

	return res
}

//Convert integer to raindrop-like string
func convert2(input int) string {
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
