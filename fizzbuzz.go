package fizzbuzz

//Fizzbuzz takes a number of args, div by 3, prints fizz
// div by 5 prints buzz
func Fizzbuzz(x int) (string, bool) {
	if x%3 == 0 && x%5 == 0 {
		return "Fizz Buzz", true
	}
	if x%3 == 0 {
		return "Fizz", true
	}
	if x%5 == 0 {
		return "Buzz", true
	}
	return "", false
}
