package testutil

import (
	"os"
	"testing"

	"github.com/pmezard/go-difflib/difflib"
)

// Must be set via the command line.
var fix string

// DoFix gets a value indicating whether the xxx_test.go files must not only
// perform the tests, but also update the files these tests depend on. The value
// returned by this function is declared via the command line:
//	go test <...> -ldflags="-X 'github.com/umk/go-testutil.fix=fix'"
// When the `fix` variable value equals to "fix", the function returns `true`,
// and otherwise `false`. Use a makefile or a script to run the test command
// with this parameter for better experience.
//
// After the files have been updated, check with a source control by diff'ing
// with the current version, that your xxx_test.go have made the proper changes
// to the files, and run the test again without the fix flag specified.
func DoFix() bool { return fix == "fix" }

// EqualDiff checks specified two strings for equality. If they are not equal,
// its diff is printed to the standard output with the expected value
// identified by `fn`, which is supposed to be a file name, where the expected
// value has been loaded from.
func EqualDiff(t *testing.T, expected, actual string, fn string) {
	d := difflib.UnifiedDiff{
		A:        difflib.SplitLines(expected),
		B:        difflib.SplitLines(actual),
		FromFile: fn,
		ToFile:   "generated by test",
		Context:  3,
	}
	m := difflib.NewMatcher(d.A, d.B)
	for _, c := range m.GetOpCodes() {
		if c.Tag != 'e' {
			if err := difflib.WriteUnifiedDiff(os.Stdout, d); err != nil {
				panic(err)
			}
			t.Fail()
			return
		}
	}
}
