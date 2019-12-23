package newton

import (
	"github.com/qamarian-dtp/err"
	"github.com/qamarian-dtp/unket"
	"math/bits"
	"strconv"
	"strings"
)

func Spelling (sound string, no int) (o []string, e error) {
	for i, aSound := range sound {
		if i == 0 {
			o = spelling_spelling [string (aSound)]["first"]
			continue
		}

		setToUse := "other"
		if (strings.Contains (Cons, string (aSound)) && strings.Contains (Cons, string (sound [i - 1]))) ||
			(strings.Contains (Vowl, string (aSound)) && strings.Contains (Vowl, string (sound [i - 1]))) {
			setToUse = "sType"
		}

		if i < (len (sound) - 1) {
		if (strings.Contains (Cons, string (aSound)) && strings.Contains (Cons, string (sound [i + 1]))) ||
			(strings.Contains (Vowl, string (aSound)) && strings.Contains (Vowl, string (sound [i + 1]))) {
			setToUse = "sType"
		}
		}

		oCopy := o
		o = []string {}

		for _, elem := range spelling_spelling [string (aSound)][setToUse] {
			for _, elemY := range oCopy {
				spelling := elemY + elem
				if len (o) == spelling_MaxInt {
					goto selection
				}
				o = append (o, spelling)
			}
		}
	}

	selection:

	if len (o) <= no {
		return
	}

	allSpelling := o
	o = []string {}
	u, errT := unket.New (int64 (len (allSpelling)))
	if errT != nil {
		e = err.New ("Broken dependency.", nil, nil, errT)
		return
	}

	for i := 1; i <= no; i ++ {
		n, errX := u.Pick ()
		if errX != nil {
			e = err.New ("Error occured while trying to pick number from unket.", nil, nil, errX)
			return
		}
		if n == nil {
			e = err.New ("Bug detected. Ref: 0.", nil, nil)
			return
		}

		noAsInt, errY := strconv.Atoi (n.String ())
		if errY != nil {
			e = err.New ("Bug detected. Ref: 1.", nil, nil)
			return
		}

		o = append (o, allSpelling [noAsInt - 1])
	}

	return
}

var spelling_spelling = map[string] map[string] []string {

	// Consonants

	"b": map[string] []string {
		"first": []string {"b"},
		"sType": []string {"b"},
		"other": []string {"b"},
	},
	"d": map[string] []string {
		"first": []string {"d"},
		"sType": []string {"d"},
		"other": []string {"d"},
	},
	"f": map[string] []string {
		"first": []string {"f", "ph"},
		"sType": []string {"f", "ph"},
		"other": []string {"f", "ph"},
	},
	"g": map[string] []string {
		"first": []string {"g"},
		"sType": []string {"g"},
		"other": []string {"g"},
	},
	"j": map[string] []string {
		"first": []string {"j", "g", "dj", "dg"},
		"sType": []string {"j", "g"},
		"other": []string {"j", "g", "dj", "dg"},
	},
	"k": map[string] []string {
		"first": []string {"k", "q"},
		"sType": []string {"k", "q"},
		"other": []string {"k", "q"},
	},
	"l": map[string] []string {
		"first": []string {"l"},
		"sType": []string {"l"},
		"other": []string {"l"},
	},
	"m": map[string] []string {
		"first": []string {"m"},
		"sType": []string {"m"},
		"other": []string {"m"},
	},
	"n": map[string] []string {
		"first": []string {"n"},
		"sType": []string {"n"},
		"other": []string {"n"},
	},
	"p": map[string] []string {
		"first": []string {"p"},
		"sType": []string {"p"},
		"other": []string {"p"},
	},
	"r": map[string] []string {
		"first": []string {"r"},
		"sType": []string {"r"},
		"other": []string {"r"},
	},
	"s": map[string] []string {
		"first": []string {"s", "c"},
		"sType": []string {"s", "c"},
		"other": []string {"s", "c"},
	},
	"t": map[string] []string {
		"first": []string {"t"},
		"sType": []string {"t"},
		"other": []string {"t"},
	},
	"v": map[string] []string {
		"first": []string {"v"},
		"sType": []string {"v"},
		"other": []string {"v"},
	},
	"w": map[string] []string {
		"first": []string {"w"},
		"sType": []string {"w"},
		"other": []string {"w"},
	},
	"y": map[string] []string {
		"first": []string {"y"},
		"sType": []string {"y"},
		"other": []string {"y"},
	},
	"z": map[string] []string {
		"first": []string {"x", "z"},
		"sType": []string {"x", "z"},
		"other": []string {"x", "z"},
	},

	// Vowels

	"a": map[string] []string {
		"first": []string {"a", "ar", "er", "ha", "har", "her"},
		"sType": []string {"a", "ar", "er"},
		"other": []string {"a", "ar", "er"},
	},
	"A": map[string] []string {
		"first": []string {"a"},
		"sType": []string {"a"},
		"other": []string {"a"},
	},
	"e": map[string] []string {
		"first": []string {"e", "he"},
		"sType": []string {"e"},
		"other": []string {"e"},
	},
	"i": map[string] []string {
		"first": []string {"i", "e", "y", "hi", "he"},
		"sType": []string {"i", "e", "y"},
		"other": []string {"i", "e", "y"},
	},
	"o": map[string] []string {
		"first": []string {"o", "ho"},
		"sType": []string {"o"},
		"other": []string {"o"},
	},
	"|": map[string] []string {
		"first": []string {"ho", "hu", "or", "ur", "hor", "hur"},
		"sType": []string {"u"},
		"other": []string {"or", "ur"},
	},
	"u": map[string] []string {
		"first": []string {"u", "hu", "oo", "hoo"},
		"sType": []string {"u"},
		"other": []string {"u"},
	},
}

var spelling_MaxInt = (1 << (bits.UintSize - 1)) - 1
