# Testdata

This little framework helps you deal with `testdata` folder for Go unit tests.

Time to time I need to write unit test that manipulate with some files and folders.
Good practice is putting all files needed for testing into `testdata` folder. 
That's fine if you need to read files and you're not modify it. But I have tests,
they're mutating files. I need some way how to ensure idempotency of tests.

This framework creates copy of your `testdata` folder in your $TEMPDIR, 
for every test run. 

## Example


### Basic usage

Let's assume we have `helloworld_test.go` and `testdata`, where is `helloworld.txt`. 
The unit test will load the text from file, append new text and save it. 


```go
func Test_HelloWorld(t *testing.T) {
   testdata.Setup()

   content, err := ioutil.ReadFile(testdata.Abs("helloworld.txt")) 
   if err != nil {
      t.FailNow()
   }

   content := fmt.Sprintf("%s hello world\n", content)

   err := ioutil.WriteFile(testdata.Abs("helloworld.txt"), []byte(content), 0644)
   if err != nil {
      t.FailNow()
   }
}
```

The above example is example of idempotent file test. First, the `Setup()` will create
copy of your `testdata` folder in $TMPDIR. The `Abs()` function will return you
absolute path to `hellowold.txt` file in $TMPDIR.


### Get file in tests without go:embed

Another common problem with tests with files is you cannot use `go embed`
direcly in your test.

You can use hi-level ReadXXXX() functions for loading files as strings, 
JSONs, YAMLs etc. These functions suppose to not fail. If there is problem 
with file, functions will give you no data.

Example how to read file as string
```go
func Test_ReadString(t *testing.T) {
   testdata.Setup()

   var text string = testdata.ReadStr("helloworld.txt")

   ...
}
```

Example how to read YAML file into structure:

```go
type Book struct {
   Title string `yaml:"title"`
   Pages int    `yaml:"pages"`
}

func Test_ReadString(t *testing.T) {
   testdata.Setup()

	book := new(Book)
	testdata.ReadYAML("folder/subfolder/book.yaml", book)

   ...
}

```
