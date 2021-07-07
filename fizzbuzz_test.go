package fizzbuzz_test

import (
	"fmt"
	"testing"

	"github.com/rajeev1976/fizzbuzz"
)

func TestFizz(t *testing.T) {
	t.Log("starting fizzbuzztest")

	_, ok := fizzbuzz.Fizzbuzz(1)
	if ok {
		t.Logf("The value %v should not have retruned true", 1)
		t.Fail()
	}

	result, ok := fizzbuzz.Fizzbuzz(3)
	if !ok {
		t.Logf("The value %v should have retruned", 3)
		t.Fail()
	}

	if result != "fizz" {
		t.Log("result should have been fizz")
		t.Fail()
	}
	t.Log("Ending fizzbuzztest")
}
func ExampleFizzbuzz() {
	result, _ := fizzbuzz.Fizzbuzz(3)
	fmt.Println(result)
}
