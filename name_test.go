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
	fmt.Print ("\n")

	patterns := []string {"cvvcvvcc", "eeccvvcc", "cvvcvvcc", "ccccvvvv"}
	for _, pattern := range patterns {
		name, errX := Name_Rand (pattern)
		if errX != nil {
			e := err.New ("Unable to obtain a random name.", nil, nil, errX)
			str.PrintEtr (errLib.Fup (e), "err", "TestName ()")
			t.FailNow ()
		}
		
		nameBeforePolish := name.String ()

		spellingBeforePolish, errY := name.Spelling (8)
		if errY != nil {
			e := err.New ("Unable to obtain spellings.", nil, nil, errY)
			str.PrintEtr (errLib.Fup (e), "err", "TestName ()")
			t.FailNow ()
		}

		name.Polish ()

		spelling, errZ := name.Spelling (8)
		if errZ != nil {
			e := err.New ("Unable to obtain spellings.", nil, nil, errZ)
			str.PrintEtr (errLib.Fup (e), "err", "TestName ()")
			t.FailNow ()
		}

		o := fmt.Sprintf ("pattern: %s;        nameBeforePolish: %s; " +
			"spellingBeforePolish: %v\n" +
			"        name: %s; spelling: %v", pattern, nameBeforePolish,
			spellingBeforePolish, name.String (), spelling)
			
		fmt.Println (o)
		fmt.Println ()
	}

	fmt.Print ("\n\n")
}
