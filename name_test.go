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

	testPattern := []string {"cvcvcvcv", "vcvcvc", "ecvc", ""}
	for _, pattern := range testPattern {
		name, errX := Name (pattern)
		if errX != nil {
			e := err.New ("Unable to create a new name.", nil, nil, errX)
			str.PrintEtr (errLib.Fup (e), "std", "TestName ()")
			continue
		}

		o := fmt.Sprintf ("Pattern %s; Name: %s", pattern, name)
		str.PrintEtr (o, "std", "TestName ()")
	}

	fmt.Print ("\n\n")
	str.PrintEtr ("Test passed!", "std", "TestName ()")
}
