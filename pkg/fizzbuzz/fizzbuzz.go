package fizzbuzz

import "strconv"

func UpToLength(length int) []string {
	out := []string{}
	for i := 1; i <= length; i++ {
		out = append(out, digitToFizzBuzz(i))
	}
	return out
}

func digitToFizzBuzz(input int) string {
	if input%3 == 0 && input%5 == 0 {
		return "FizzBuzz"
	} else if input%3 == 0 {
		return "Fizz"
	} else if input%5 == 0 {
		return "Buzz"
	} else {
		return strconv.Itoa(input)
	}
}
