## Go testing utilities

[![GoDoc](https://godoc.org/github.com/umk/go-testutil?status.svg)](https://godoc.org/github.com/umk/go-testutil)
[![Go Report Card](https://goreportcard.com/badge/github.com/umk/go-testutil)](https://goreportcard.com/report/github.com/umk/go-testutil)

This package provides the helper functions for testing the code, written in Go.

### DoFix

```go
func DoFix() bool
```

DoFix gets a value indicating whether the xxx_test.go files must not only perform the tests, but also update the files these tests depend on. The value returned by this function is declared via the command line:
```
go test ... -ldflags="-X 'github.com/umk/go-testutil.fix=fix'"
```
When the `fix` variable value equals to "fix", the function returns `true`, and otherwise `false`. Use a makefile or a script to run the test command with this parameter for better experience.

After the files have been updated, check with a source control by diff'ing with the current version, that your xxx_test.go have made the proper changes to the files, and run the test again without the fix flag specified.

### EqualDiff

```go
func EqualDiff(t *testing.T, expected, actual string, fn string)
```

EqualDiff checks specified two strings for equality. If they are not equal, its diff is printed to the standard output with the expected value identified by `fn`, which is supposed to be a file name, where the expected value has been loaded from.

### Lfsr28

Lfsr28 is a 28-bit linear-feedback shift register which generates an m-sequence. From [Wikipedia](https://en.wikipedia.org/wiki/Maximum_length_sequence):

> They are bit sequences generated using maximal linear feedback shift registers and are so called because they are periodic and reproduce every binary sequence (except the zero vector) that can be represented by the shift registers (i.e., for length-m registers they produce a sequence of length 2<sup>m</sup> − 1). 

Use this to produce a pseudo-random sequence of numbers to implement reproducible tests.
