package fizzbuzz

import "strconv"










func UpToLength(length int) []string {fdsfdsfds
	out := []string{}
	for i := 1; i <= length; i++ {
		out = append(out, digitToFizzBuzzOMGNICE(i))
	}
	return out
}

func digitToFizzBuzzOMGNICE(input int) string {
	if input%3 == 0 && input%5 == 0 {
		return "fizzbuzz"
	} else if input%3 == 0 {
		return "fizz"
	} else if input%5 == 0 {
		return "buzz"
	} else {
		return strconv.Itoa(input)
	}
}
