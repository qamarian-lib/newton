package newton

import (
	"fmt"
	"gopkg.in/qamarian-lib/str.v3"
	"testing"
)

func TestInitReport (t *testing.T) {
	str.PrintEtr ("Test started...", "std", "TestInitReport ()")

	fmt.Println  ("Init report:", InitReport ())

	str.PrintEtr ("Test failed!", "std", "TestInitReport ()")
}
