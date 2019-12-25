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
	fmt.Print ("\n")
	baseChar := byte ('b')

	for x := 1; x <= 4; x ++ {
		soundX, errX := Sound_New (rune (baseChar + byte (x)))
		if errX != nil {
			e := err.New ("Unable to create sound.", nil, nil, errX)
			str.PrintEtr (errLib.Fup (e), "err", "TestSound ()")
			t.FailNow ()
		}
		oX := fmt.Sprintf ("input sound: %s;        sound; %s; type: %d",
			string (baseChar + byte (x)), soundX.String (), soundX.Type ())
		fmt.Println (oX)

		soundY, errY := Sound_Rand ('c')
		if errY != nil {
			e := err.New ("Unable to obtain random sound.", nil, nil, errY)
			str.PrintEtr (errLib.Fup (e), "err", "TestSound ()")
			t.FailNow ()
		}
		oY := fmt.Sprintf ("input sound: ?;        sound; %s; type: %d",
			soundY.String (), soundY.Type ())
		fmt.Println (oY)

		soundZ, errZ := Sound_Rand ('v')
		if errZ != nil {
			e := err.New ("Unable to obtain random sound.", nil, nil, errZ)
			str.PrintEtr (errLib.Fup (e), "err", "TestSound ()")
			t.FailNow ()
		}
		oZ := fmt.Sprintf ("input sound: ?;        sound; %s; type: %d",
			soundZ.String (), soundZ.Type ())
		fmt.Println (oZ)

		soundA, errA := Sound_Rand ()
		if errA != nil {
			e := err.New ("Unable to obtain random sound.", nil, nil, errA)
			str.PrintEtr (errLib.Fup (e), "err", "TestSound ()")
			t.FailNow ()
		}
		oA := fmt.Sprintf ("input sound: ?;        sound; %s; type: %d",
			soundA.String (), soundA.Type ())
		fmt.Println (oA)

		fmt.Println ()
	}

	str.PrintEtr ("Test passed.", "std", "TestSound()")
}
