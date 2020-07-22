// Copyright © 2020 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"container/ring"
	"errors"
	"fmt"
	"github.com/spf13/cobra"
)

// chordCmd represents the chord command
var chordCmd = &cobra.Command{
	Use:   "chord",
	Short: "generate valid chord shapes for the configured instrument and desired chord",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString("name")
		if name == "" {
			fmt.Println("chord name must be specified")
		}
		chordData, err := parseChord(name)

	},
}

func init() {
	RootCmd.AddCommand(chordCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// chordCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	chordCmd.Flags().StringP("name", "n", "", "Name of the chord voicing to generate.")
}

// type to hold information on the chord. The chord types parsed so far are: major, minor, augmented, and diminished
type chord struct {
	rootNote string
	chordType string
	chordNotes []string
}

func parseChord(name string) (*chord, error) {
	chordData := chord{}
	noteRing := initNoteRing()

	err := parseRoot(name, &chordData, noteRing)
	if err != nil {
		return &chordData, err
	}

	err = parseType(name, &chordData)

	return nil, nil
}

func parseRoot(name string, chordData *chord, noteRing *ring.Ring) error {
	rootValid := validateNote(string(name[0]), noteRing)
	if !rootValid {
		return errors.New("error: the root note of the chord is not a valid note")
	}
	chordData.rootNote = string(name[0])
	return nil
}

func parseType(name string, chordData *chord) error {
	//1 char length chords
	if len(name) == 1 {
		chordData.chordType = "MAJ"
		return nil
	}

	//Check for letters (indicates
	containsChar := false
	for char := range name[1:] {
		if (char > "A" && char < "Z") || (char > "a" && char < "z") {
			containsChar = true
		}
	}


}