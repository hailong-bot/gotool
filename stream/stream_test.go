package stream

import (
	"fmt"
	"testing"
)

func Test1(t *testing.T) {
	t.Parallel()
	type Person struct {
		Name string `json:"name"`
		Age  int64  `json:"age"`
	}

	needs := []Person{
		{Name: "Tom", Age: 18},
		{Name: "Mike", Age: 19},
	}
	res := Of(needs...).Filter(func(person Person) bool {
		return person.Age > 18
	}).Collect(&ToCollection{})
	fmt.Println("res", res)
}