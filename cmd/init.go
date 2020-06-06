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
	noteRing.Next()
	noteRing.Value = "A#"
	noteRing.Next()
	noteRing.Value = "B"
	noteRing.Next()
	noteRing.Value = "C"
	noteRing.Next()
	noteRing.Value = "C#"
	noteRing.Next()
	noteRing.Value = "D"
	noteRing.Next()
	noteRing.Value = "D#"
	noteRing.Next()
	noteRing.Value = "E"
	noteRing.Next()
	noteRing.Value = "F"
	noteRing.Next()
	noteRing.Value = "F#"
	noteRing.Next()
	noteRing.Value = "G"
	noteRing.Next()
	noteRing.Value = "G#"
	noteRing.Next()

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
		err := errors.New("config file is invalid. Must contain the number of strings, number of frets, and a slice" +
			"of valid string tuning values equal to the number of strings")
		return settings, err
	}

	//Validate string tunings contains only valid notes
	for _, v := range settings.stringTunings {
		foundNote := false
		for i := 0; i < 11; i++ {
			if v == noteRing.Value {
				foundNote = true
				break
			}
			noteRing.Next()
		}

		if foundNote == false {
			err := errors.New("config file is invalid. Must contain the number of strings, number of frets, and a slice" +
				"of valid string tuning values equal to the number of strings")
			return settings, err
		}
	}

	return settings, nil
}
