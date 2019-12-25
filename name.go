package newton

import (
	"github.com/kensharval/etc"
	"github.com/qamarian-dtp/err"
	"github.com/qamarian-dtp/unket"
	"math/bits"
	"regexp"
	"strconv"
	"strings"
)

// Function Name_Rand () provides a random name (i.e. a name's sound). Argument should be
// the desired pattern for the name's sound.
//
// Example:
//
// 	cevecv
// 	cve
//
// 'c' stands for consonant sound;
// 'v' stands for vowel sound;
// 'e' stands for either 'c' or 'v'
//
// In other words, pattern 'ecv' can yield 'et@'.
//
func Name_Rand (pattern string) (name Name, e error) {

	if ! name_new_inputPattern.MatchString (pattern) {
		e = err.New ("Invalid input", nil, nil)
		return
	}

	soundPattern, errV := etc.Decompress (pattern)
	if errV != nil {
		e = err.New ("Unable to decompress pattern.", nil, nil, errV)
		return
	}

	soundArray := []Sound {}
	for _, s := range soundPattern {
		var sound Sound

		if s == 'c' || s == 'v' {
			var errX error
			sound, errX = Sound_Rand (s)
			if errX != nil {
				e = err.New ("Unable to obtain a sound.", nil, nil, errX)
				return
			}
		} else  {
			var errX error
			sound, errX = Sound_Rand ()
			if errX != nil {
				e = err.New ("Unable to obtain a sound.", nil, nil, errX)
				return
			}
		}

		soundArray = append (soundArray, sound)
	}

	name = Name_New (soundArray)

	return
}; var (
	
	name_new_inputPattern *regexp.Regexp

); func init () {
	
	if initReport != nil {
		return
	}

	var errX error
	name_new_inputPattern, errX = regexp.Compile (`^([cve]([2-9]|[1-9]\d+)?)+$`)
	if errX != nil {
		initReport = err.New ("Regular expression compilation failed. {Name ()}.",
			nil, nil)
		return
	}
}

func Name_New (sound []Sound) (Name) {
	return Name {sound}
}

type Name struct {
	name []Sound
}

// Method Polish () modifies a name such that it becomes easier to pronounce.
func (n *Name) Polish () {
	newSound := []Sound {}
	ascendSound := n.name [0]
	lastAscendType := byte (0) // default value

	for x := 2; x <= len (n.name); x ++ {
		if len (newSound) > 0 && lastAscendType == ascendSound.Type () &&
			ascendSound.Type () == n.name [x - 1].Type () {

			continue
		}

		if (ascendSound.Type () == Sound_TypeCons &&
			n.name [x - 1].Type () == Sound_TypeCons) ||
			(ascendSound.Type () == Sound_TypeVowl &&
			n.name [x - 1].Type () == Sound_TypeVowl) {

			if ! strings.Contains (name_polish_goodCombo, ascendSound.String () +
				n.name [x - 1].String ()) {

				continue
			}
		}

		newSound = append (newSound, ascendSound)
		lastAscendType = ascendSound.Type ()
		ascendSound = n.name [x - 1]
	}
	newSound = append (newSound, ascendSound)

	n.name = newSound

	return
}; var (
	name_polish_goodCombo = `
ck
cl
cm
cn
ct
cp
cw
==

kl
kn
kr
==

pl
pr
==




@a
@!
@e
@o
@_
@#
==

a@
a!
ae
ao
a_
a#
==

!@
!a
!e
!o
!_
!#
==

e@
ea
e!
ee
eo
e_
e#
==

o@
oa
o!
oe
o_
o#
==

_@
_a
_!
_e
_u
==

#@
#a
#!
#e
#o
#_
==
`
)

// Method Spelling () provides English spelling for the name. Its argument should be the
// number of possible spellings you would like.
//
// If there are as much spelling as you requested, you would get that number of spellings.
// If there are not as much spelling as you requested, all the ones available would be
// returned.
//
func (n Name) Spelling (no int) (o []string, e error) {
	for i, aSound := range n.name {
		if i == 0 {
			o = name_spelling_spelling [aSound.String ()]["first"]
			continue
		}

		setToUse := "other"
		if (n.name [i - 1].Type () == Sound_TypeCons &&
			aSound.Type () == Sound_TypeCons) ||
			(n.name [i - 1].Type () == Sound_TypeVowl &&
			aSound.Type () == Sound_TypeVowl) {
			
			setToUse = "sType"
		}

		if i < (len (n.name) - 1) {
			if (aSound.Type () == Sound_TypeCons &&
				n.name [i + 1].Type () == Sound_TypeCons) ||
				(aSound.Type () == Sound_TypeVowl &&
				n.name [i + 1].Type () == Sound_TypeVowl) {
			
				setToUse = "sType"
			}
		}

		oCopy := o
		o = []string {}

		for _, elem := range name_spelling_spelling [aSound.String ()][setToUse] {
			for _, elemY := range oCopy {
				spelling := elemY + elem
				if len (o) == name_spelling_MaxInt {
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
}; var (
	name_spelling_spelling = map[string] map[string] []string {

// Consonant sounds spelling
	"b": map[string] []string {
		"first": []string {"b"},
		"sType": []string {"b"},
		"other": []string {"b"},
	},
	"c": map[string] []string {
		"first": []string {"c", "s"},
		"sType": []string {"s"},
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
		"sType": []string {"a"},
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

	name_spelling_MaxInt = (1 << (bits.UintSize - 1)) - 1
)

func (n Name) String () (o string) {
	for _, sound := range n.name {
		o += sound.String ()
	}

	return
}
