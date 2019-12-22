package newton

import (
	"github.com/qamarian-dtp/err"
	"strings"
)

// Function Name () provides a random name (i.e. a name's sound). Argument should be the
// length of the name (i.e. the length of the sound of the name).
func Name (length int) (name string, e error) {
	// Input validation. ..1.. {
	if length < 1 {
		e = err.New ("Invalid length.", nil, nil)
		return
	}
	// ..1.. }

	for i := 1; i <= length; i ++ {
		if i == 1 || i == 2 {
			sound, errX := Sound ()
			if errX != nil {
				e = err.New ("Unable to source a new sound.", nil, nil,
					errX)
				return
			}
			name = name + string (sound)
			continue
		}

		var nextSoundType rune
		if strings.Contains (Cons, string (name [i - 3])) &&
			strings.Contains (Cons, string (name [i - 2])) {
			nextSoundType = 'v'
		}
		if strings.Contains (Vowl, string (name [i - 3])) &&
			strings.Contains (Vowl, string (name [i - 2])) {
			nextSoundType = 'c'
		}

		if nextSoundType == 'c' || nextSoundType == 'v' {
			sound, errX := Sound (nextSoundType)
			if errX != nil {
				e = err.New ("Unable to source a new sound.", nil, nil,
					errX)
				return
			}
			name = name + string (sound)
			continue
		}
		sound, errX := Sound ()
		if errX != nil {
			e = err.New ("Unable to source a new sound.", nil, nil, errX)
			return
		}
		name = name + string (sound)
	}

	return
}
