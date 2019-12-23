package newton

import (
	"github.com/kensharval/etc"
	"github.com/qamarian-dtp/err"
	"regexp"
)

// Function Name () provides a random name (i.e. a name's sound). Argument should be the
// length of the name (i.e. the length of the sound of the name).
func Name (pattern string) (name string, e error) {
	if name_InputPattern.MatchString (pattern) == false {
		e = err.New ("Invalid input", nil, nil)
		return
	}

	soundPattern, errV := etc.Decompress (pattern)
	if errV != nil {
		e = err.New ("Bug detected: possibly due to broken dependency; ref 0.", nil, nil, errV)
		return
	}

	for _, s := range soundPattern {
		sound, errX := Sound ()
		if errX != nil {
			e = err.New ("Bug detected: possibly due to broken dependency; ref 1.", nil, nil, errX)
			return
		}
		if s != 'e' {
			sound, errX = Sound (s)
			if errX != nil {
				e = err.New ("Bug detected: possibly due to broken dependency; ref 2.", nil, nil, errX)
				return
			}
		}
		name = name + string (sound)
	}

	return
}
var (
	name_InputPattern *regexp.Regexp
)
func init () {
	if initReport != nil {
		return
	}

	var errX error
	name_InputPattern, errX = regexp.Compile (`^([cve]([2-9]|[1-9]\d+)?)+$`)
	if errX != nil {
		initReport = err.New ("Buggy package: regular expression compilation failed; ref: name-1.", nil, nil)
	}
}
