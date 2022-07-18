package main

const REPEAT_COUNT = 5

func Repeat(str string) string {
	var newStr string
	for i := 0; i < REPEAT_COUNT; i++ {
		newStr += str
	}
	return newStr
}
