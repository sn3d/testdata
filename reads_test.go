package testdata_test

import (
	"fmt"

	"github.com/sn3d/testdata"
)

func ExampleReadStr() {
	testdata.Setup()

	helloworld := testdata.ReadStr("folder/subfolder/helloworld.txt")
	fmt.Println(helloworld)
	// Output: Testdata
}

func ExampleReadYAML() {
	testdata.Setup()

	book := struct {
		Title string `yaml:"title"`
		Pages int    `yaml:"pages"`
	}{}

	testdata.ReadYAML("folder/subfolder/book.yaml", &book)

	fmt.Println(book.Title)
	fmt.Println(book.Pages)
	// Output:
	// The Mythical Man-Month
	// 272
}

func ExampleReadJSON() {
	testdata.Setup()

	book := struct {
		Title string `json:"title"`
		Pages int    `json:"pages"`
	}{}

	testdata.ReadJSON("folder/subfolder/book.json", &book)

	fmt.Println(book.Title)
	fmt.Println(book.Pages)
	// Output:
	// The Mythical Man-Month
	// 272
}
