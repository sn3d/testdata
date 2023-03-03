# Testdata

This little framework helps you deal with `testdata` folder for Go unit tests.

Time to time I need to write unit test that manipulate with some files and folders.
Good practice is putting all files needed for testing into `testdata` folder. 
That's fine if you need to read files and you're not modify it. But I have tests,
they're mutating files. I need some way how to ensure idempotency of tests.

This framework creates copy of your `testdata` folder in your $TEMPDIR, 
for every test run. 

## Example

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
