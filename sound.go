package newton

import (
	"crypto/rand"
	"github.com/qamarian-dtp/err"
	"regexp"
)

// Function Sound_Rand () provides a random sound. The type of sound you want can be
// optionally supplied as an argument. For a consonant sound: use 'c' as your argument,
// and for a vowel sound: use 'v'.
//
//	sound, err := Sound_Rand ('c')
//
// If you are not bothered about the sound being consonant or vowel, do not supply any
// argument.
//
// 	sound, err := Sound_Rand ()
//
func Sound_Rand (soundType ... rune) (sound Sound, e error) {
	// Input validation. ..1.. {
	if len (soundType) > 1 {
		e = err.New ("More than one sound type specified.", nil, nil)
		return
	}

	if len (soundType) == 1 && ((soundType [0] != 'c') && (soundType [0] != 'v')) {
		e = err.New ("Invalid sound type.", nil, nil)
		return
	}
	// ..1.. }

	// ..1.. {
	sounds := []rune {}
	if len (soundType) == 1 && soundType [0] == 'c' {
		sounds = sound_cons
	} else if len (soundType) == 1 && soundType [0] == 'v' {
		sounds = sound_vowl
	} else {
		sounds = sound_all
	}
	// ..1.. }
	
	randByte := make ([]byte, 1)
	_, errX := rand.Read (randByte)
	if errX != nil {
		e = err.New ("Unable to obtain a random byte: random byte need for " +
			"deciding what sound to create.", nil, nil, errX)
		return
	}
	chosenSound := int (randByte [0]) % len (sounds)

	var errY error
	sound, errY = Sound_New (sounds [chosenSound])
	if errY != nil {
		e = err.New ("Unable to obtain a new sound.", nil, nil, errX)
		return
	}

	return
}

func Sound_New (sound rune) (s Sound, e error) {
	if ! sound_soundPattern.MatchString (string (sound)) {
		e = err.New ("Invalid sound.", nil, nil)
		return
	}

	s = Sound {string (sound), Sound_TypeCons}

	if ! sound_consPattern.MatchString (string (sound)) {
		s.soundType = Sound_TypeVowl
	}

	return
}

type Sound struct {
	sound string
	soundType byte
}; var (
	sound_cons = []rune ("bcdf*jklmnprtvwxy")
	sound_vowl = []rune ("@a!eo_#")
	sound_all  = append (sound_cons, sound_vowl...)

	Sound_TypeCons byte = 0
	Sound_TypeVowl byte = 1

	sound_soundPattern *regexp.Regexp
	sound_consPattern *regexp.Regexp

); func init () {

	if initReport != nil {
		return
	}

	var errX error
	sound_soundPattern, errX = regexp.Compile (`^[bcdf\*jklmnprtvwxy@a!eo_#]$`)
	if errX != nil {
		initReport = err.New ("Sound () initialization failed: sound regular " +
			"expression compilation failed.", nil, nil, errX)
		return
	}

	var errY error
	sound_consPattern, errY = regexp.Compile (`^[bcdf\*jklmnprtvwxy]$`)
	if errY != nil {
		initReport = err.New ("Sound () initialization failed: consonant " +
			"regular expression compilation failed.", nil, nil, errY)
		return
	}
}

func (s *Sound) String () (string) {
	return s.sound
}

// Method Type () tells if the sound is a consonant (Sound_TypeCons) or vowel
// (Sound_TypeCons).
func (s *Sound) Type () (byte) {
	return s.soundType
}
