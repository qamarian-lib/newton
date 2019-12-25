package newton

import (
	"github.com/kensharval/etc"
	"github.com/qamarian-dtp/err"
	"regexp"
)

// Function Name () provides a random name (i.e. a name's sound). Argument should be the
// desired pattern for the name's sound.
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
func Name (pattern string) (name string, e error) {
	if name_inputPattern.MatchString (pattern) == false {
		e = err.New ("Invalid input", nil, nil)
		return
	}

	soundPattern, errV := etc.Decompress (pattern)
	if errV != nil {
		e = err.New ("Unable to decompress pattern.", nil, nil, errV)
		return
	}

	for _, s := range soundPattern {
		var sound rune
		if s != 'e' {
			var errX error
			sound, errX = Sound (s)
			if errX != nil {
				e = err.New ("Unable to obtain a sound.", nil, nil, errX)
				return
			}
		} else  {
			var errX error
			sound, errX = Sound ()
			if errX != nil {
				e = err.New ("Unable to obtain a sound.", nil, nil, errX)
				return
			}
		}

		name = name + string (sound)
	}

	return
}
var (
	name_inputPattern *regexp.Regexp
)
func init () {
	if initReport != nil {
		return
	}

	var errX error
	name_inputPattern, errX = regexp.Compile (`^([cve]([2-9]|[1-9]\d+)?)+$`)
	if errX != nil {
		initReport = err.New ("Regular expression compilation failed. {Name ()}.",
			nil, nil)
	}
}
