package newton

import (
	"github.com/qamarian-dtp/err"
	"github.com/qamarian-dtp/unket"
	"math/bits"
	"strconv"
	"strings"
)

// Function Spelling () provides English spelling for a sound. The first argument should be
// the sound, while the second argument should be the number of possible spellings you
// want.
//
// If there are as much spelling as you requested, you would get that number of spellings.
// If there are not as much spelling as you requested, all the ones available would be
// returned.
//
func Spelling (sound string, no int) (o []string, e error) {
	for i, aSound := range sound {
		if i == 0 {
			o = spelling_spelling [string (aSound)]["first"]
			continue
		}

		setToUse := "other"
		if (strings.Contains (Cons, string (aSound)) && strings.Contains (Cons,
			string (sound [i - 1]))) || (strings.Contains (Vowl,
			string (aSound)) && strings.Contains (Vowl,
			string (sound [i - 1]))) {
			
			setToUse = "sType"
		}

		if i < (len (sound) - 1) {
		if (strings.Contains (Cons, string (aSound)) && strings.Contains (Cons,
			string (sound [i + 1]))) || (strings.Contains (Vowl,
			string (aSound)) && strings.Contains (Vowl,
			string (sound [i + 1]))) {
			
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
		e = err.New ("Unable to randomly select from spellings available: " +
			"unable to create gene need for the random selection.", nil, nil,
			errT)
		return
	}

	for i := 1; i <= no; i ++ {
		n, errX := u.Pick ()
		if errX != nil {
			e = err.New ("Unable to randomly select from spellings " +
				"available: unable to obtain a random number that " +
				"would be used for selection.", nil, nil, errX)
			return
		}
		if n == nil {
			e = err.New ("Unable to randomly select from spellings " +
				"available: no random number was returned by dependency.",
				nil, nil)
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

// Consonant sounds spelling
	"b": map[string] []string {
		"first": []string {"b"},
		"sType": []string {"b"},
		"other": []string {"b"},
	},
	"c": map[string] []string {
		"first": []string {"c", "s"},
		"sType": []string {"c", "s"},
		"other": []string {"c", "s"},
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
	"*": map[string] []string {
		"first": []string {"g"},
		"sType": []string {"g"},
		"other": []string {"g"},
	},
	"j": map[string] []string {
		"first": []string {"j"},
		"sType": []string {"j"},
		"other": []string {"j"},
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
	"x": map[string] []string {
		"first": []string {"x", "z"},
		"sType": []string {"x", "z"},
		"other": []string {"x", "z"},
	},
	"y": map[string] []string {
		"first": []string {"y"},
		"sType": []string {"y"},
		"other": []string {"y"},
	},

// Vowel sounds spellings
	"@": map[string] []string {
		"first": []string {"a", "ar", "er", "ha", "har", "her"},
		"sType": []string {"a", "ar", "er"},
		"other": []string {"a", "ar", "er"},
	},
	"a": map[string] []string {
		"first": []string {"a"},
		"sType": []string {"a"},
		"other": []string {"a"},
	},
	"!": map[string] []string {
		"first": []string {"e", "he"},
		"sType": []string {"e"},
		"other": []string {"e"},
	},
	"e": map[string] []string {
		"first": []string {"i", "ea", "he", "hi", "hea"},
		"sType": []string {"i"},
		"other": []string {"i"},
	},
	"o": map[string] []string {
		"first": []string {"o", "ho"},
		"sType": []string {"o"},
		"other": []string {"o"},
	},
	"_": map[string] []string {
		"first": []string {"ho", "hu", "or", "ur", "hor", "hur"},
		"sType": []string {"u"},
		"other": []string {"or", "ur"},
	},
	"#": map[string] []string {
		"first": []string {"u", "hu", "oo", "hoo"},
		"sType": []string {"u"},
		"other": []string {"u"},
	},
}

var spelling_MaxInt = (1 << (bits.UintSize - 1)) - 1
