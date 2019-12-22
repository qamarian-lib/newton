package newton

import (
	"fmt"
	"github.com/qamarian-dtp/err"
	errLib "github.com/qamarian-lib/err"
	"gopkg.in/qamarian-lib/str.v3"
	"testing"
)

func TestName (t *testing.T) {
	str.PrintEtr ("Test started...", "std", "TestName ()")
	fmt.Print ("\n\n")

	for i := 1; i <= 16; i ++ {
		name, errX := Name (i)
		if errX != nil {
			e := err.New ("Unable to create a new name.", nil, nil, errX)
			str.PrintEtr (errLib.Fup (e), "err", "TestName ()")
			return
		}

		o := fmt.Sprintf ("Name %d: %s", i, name)
		str.PrintEtr (o, "std", "TestName ()")
	}

	fmt.Print ("\n\n")
	str.PrintEtr ("Test passed!", "std", "TestName ()")
}
