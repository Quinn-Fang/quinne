package main

func main() {
	a, b := 5, "hello world"
	if len(b) < 10 {
		bodyLessFunction_1()
	} else if len(b) > 6 || 7 > 8 {
		bodyLessFunction_2()
	}
}
