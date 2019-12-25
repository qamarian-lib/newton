package newton

import (
	"fmt"
	"github.com/qamarian-dtp/err"
	errLib "github.com/qamarian-lib/err"
	"gopkg.in/qamarian-lib/str.v3"
	"testing"
)

func TestSpelling (t *testing.T) {
	str.PrintEtr ("Test started...", "std", "TestSpelling ()")
	fmt.Println ()

	for i := 1; i <= 8; i ++ {
		n, errX := Name ("cvvccvc")
		if errX != nil {
			errY := err.New ("Unable to obtain a name.", nil, nil, errX)
			str.PrintEtr (errLib.Fup (errY), "err", "TestSpelling ()")
			t.FailNow ()
		}

		spellings, errZ := Spelling (n, 16)
		if errZ != nil {
			errA := err.New ("Unable to obtain spellings for obtained name.", nil, nil, errZ)
			str.PrintEtr (errLib.Fup (errA), "err", "TestSpelling ()")
			t.FailNow ()
		}

		o := fmt.Sprintf ("%s: %v", n, spellings)
		fmt.Println (o)
		fmt.Println ()
	}

	fmt.Println ()
	str.PrintEtr ("Test passed!", "std", "TestSpelling ()")
}
