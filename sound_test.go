package newton

import (
	"fmt"
	"github.com/qamarian-dtp/err"
	errLib "github.com/qamarian-lib/err"
	"gopkg.in/qamarian-lib/str.v3"
	"testing"
)

func TestSound (t *testing.T) {
	str.PrintEtr ("Test started...", "std", "TestSound ()")

	for i := 1; i <= 8; i ++ {
		// ..1.. {
		s, errX := Sound ()
		if errX != nil {
			e := err.New ("Test failed. [Unable to obtain a sound.]", nil, nil,
				errX)
			str.PrintEtr (errLib.Fup (e), "err", "TestSound ()")
			t.FailNow ()
		}
		str.PrintEtr ("Some sound: " + string (s), "std", "TestSound ()")	
		// ..1.. }

		// ..1.. {
		c, errY := Sound ('c')
		if errY != nil {
			e := err.New ("Test failed. [Unable to obtain a sound.]", nil, nil,
				errY)
			str.PrintEtr (errLib.Fup (e), "err", "TestSound ()")
			t.FailNow ()
		}
		str.PrintEtr ("Cons sound: " + string (c), "std", "TestSound ()")	
		// ..1.. }

		// ..1.. {
		v, errZ := Sound ('v')
		if errZ != nil {
			e := err.New ("Test failed. [Unable to obtain a sound.]", nil, nil,
				errZ)
			str.PrintEtr (errLib.Fup (e), "err", "TestSound ()")
			t.FailNow ()
		}
		str.PrintEtr ("Vowl sound: " + string (v), "std", "TestSound ()")	
		// ..1.. }
		
		fmt.Print ("\n\n")
	}

	str.PrintEtr ("Test passed!", "std", "TestSound ()")
}
