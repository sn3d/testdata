package testdata_test

import (
	"fmt"

	"github.com/sn3d/testdata"
)

func ExampleReadAsStr() {
	testdata.Setup()

	helloworld := testdata.ReadAsStr("helloworld.txt")
	fmt.Println(helloworld)
	// Output: Testdata
}
