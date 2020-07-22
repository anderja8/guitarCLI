package cmd

import (
	"container/ring"
	"errors"
	"github.com/spf13/viper"
)

// initNoteRing will initialize a ring data structure with the 12 notes
// in modern Western music.
func initNoteRing() *ring.Ring {
	noteRing := ring.New(12)
	noteRing.Value = "A"
	noteRing = noteRing.Next()
	noteRing.Value = "A#"
	noteRing = noteRing.Next()
	noteRing.Value = "B"
	noteRing = noteRing.Next()
	noteRing.Value = "C"
	noteRing = noteRing.Next()
	noteRing.Value = "C#"
	noteRing = noteRing.Next()
	noteRing.Value = "D"
	noteRing = noteRing.Next()
	noteRing.Value = "D#"
	noteRing = noteRing.Next()
	noteRing.Value = "E"
	noteRing = noteRing.Next()
	noteRing.Value = "F"
	noteRing = noteRing.Next()
	noteRing.Value = "F#"
	noteRing = noteRing.Next()
	noteRing.Value = "G"
	noteRing = noteRing.Next()
	noteRing.Value = "G#"
	noteRing = noteRing.Next()

	return noteRing
}

// guitarSettings type will hold information on what type of guitar
// to do calculations on
type GuitarSettings struct {
	numStrings    int      `yml:num_strings`
	numFrets      int      `yml:num_frets`
	stringTunings []string `yml:string_tunings`
}

// initGuitarSettings will populate the guitar  settings type as specified in
// config.yml in the working directory
func initGuitarSettings(noteRing *ring.Ring) (*GuitarSettings, error) {
	var settings = &GuitarSettings{-1, -1, nil}

	//Read in the settings from viper
	settings.numStrings = viper.GetInt("num_strings")
	settings.numFrets = viper.GetInt("num_frets")
	settings.stringTunings = viper.GetStringSlice("string_tunings")

	//Validate the settings were found in viper and string tunings is the right length
	if settings.numStrings == -1 || settings.numFrets == -1 || settings.stringTunings == nil ||
		len(settings.stringTunings) != settings.numStrings {
		err := errors.New("config file is invalid - must contain the number of strings, number of frets, and a slice" +
			" of valid string tuning values equal to the number of strings")
		return settings, err
	}

	//Validate string tunings contains only valid notes
	for _, v := range settings.stringTunings {
		foundNote := false
		for i := 0; i < noteRing.Len(); i++ {
			if v == noteRing.Value {
				foundNote = true
				break
			}
			noteRing = noteRing.Next()
		}

		if foundNote == false {
			err := errors.New("config file is invalid - must contain the number of strings, number of frets, and a slice" +
				" of valid string tuning values equal to the number of strings")
			return settings, err
		}
	}

	return settings, nil
}

// NoteIntervals
type NoteIntervals struct {
	root, min2, maj2, min3, maj3, p4, p5, min6, maj6, min7, maj7 string
}

func initNoteIntervals(rootNote string, noteRing *ring.Ring) (*NoteIntervals, error) {
	foundNote := false
	for i := 0; i < noteRing.Len(); i++ {
		if rootNote == noteRing.Value {
			foundNote = true
			break
		}
		noteRing = noteRing.Next()
	}
	if foundNote == false {
		err := errors.New("error the root note could not be found in the noteRing")
		return nil, err
	}

	intervals := NoteIntervals{}

	noteRing = noteRing.Next()
	intervals.min2 = noteRing.Value.(string)

	noteRing = noteRing.Next()
	intervals.maj2 = noteRing.Value.(string)

	noteRing = noteRing.Next()
	intervals.min3 = noteRing.Value.(string)

	noteRing = noteRing.Next()
	intervals.maj3 = noteRing.Value.(string)

	noteRing = noteRing.Next()
	intervals.p4 = noteRing.Value.(string)

	noteRing = noteRing.Next()
	intervals.p5 = noteRing.Value.(string)

	noteRing = noteRing.Next()
	intervals.min6 = noteRing.Value.(string)

	noteRing = noteRing.Next()
	intervals.maj6 = noteRing.Value.(string)

	noteRing = noteRing.Next()
	intervals.min7 = noteRing.Value.(string)

	noteRing = noteRing.Next()
	intervals.maj7 = noteRing.Value.(string)

	return &intervals, nil
}