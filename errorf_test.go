package errorf_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/Konstantin8105/errorf"
)

func testfunction() {
	_ = fmt.Errorf("Wrong Errorf")
}

func Test(t *testing.T) {
	t.Run("not filename", func(t *testing.T) {
		err := errorf.Test("./some not exist file.go")
		if err == nil {
			t.Error("error not found")
		}
	})

	t.Run("not valid expressions", func(t *testing.T) {
		err := errorf.Test("./errorf_test.go")
		if err == nil {
			t.Errorf("error not found")
		}
	})

	t.Run("acceptable test", func(t *testing.T) {
		err := errorf.Test("./errorf.go")
		if err != nil {
			t.Error("error not found")
		}
	})
}

func Example() {
	err := errorf.Test("./errorf_test.go")
	fmt.Fprintf(os.Stdout, "%v", err)
	// Output:
	// ./errorf_test.go
	// └──./errorf_test.go:12:6:	not acceprable first letter: "Wrong Errorf"
}
