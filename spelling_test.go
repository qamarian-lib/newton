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
		n, errX := Name ("ecvcv")
		if errX != nil {
			errY := err.New ("Broken dependency. Ref: 0", nil, nil, errX)
			str.PrintEtr (errLib.Fup (errY), "err", "TestSpelling ()")
			t.FailNow ()
		}

		spellings, errZ := Spelling (n, 16)
		if errZ != nil {
			errA := err.New ("Broken dependency. Ref: 1", nil, nil, errZ)
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
