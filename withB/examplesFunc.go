package main

import "strconv"

func getRightTime(sec int) string {
	var hh = sec / 3600
	sec = sec - hh*3600
	var mm = sec / 60
	sec = sec - mm*60
	var s = ""
	if hh < 10 {
		s += "0"
	}
	s += strconv.Itoa(hh) + ":"
	if mm < 10 {
		s += "0"
	}
	s += strconv.Itoa(mm) + ":"
	if sec < 10 {
		s += "0"
	}
	s += strconv.Itoa(sec)
	return s
}

func main() {
	time := getRightTime(43253)
	println(time)
}
