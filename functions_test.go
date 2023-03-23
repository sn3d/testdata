package testdata_test

import (
	"fmt"

	"github.com/sn3d/testdata"
)

func ExampleCompareFiles() {
	testdata.Setup()

	fmt.Println(testdata.CompareFiles("compare/f1.txt", "compare/f2.txt"))
	fmt.Println(testdata.CompareFiles("compare/f1.txt", "compare/f3.txt"))

	// Output:
	// true
	// false
}
