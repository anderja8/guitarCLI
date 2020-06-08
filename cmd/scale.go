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
	"github.com/spf13/viper"
	"strconv"
	"strings"
)

// scaleCmd represents the scale command
var scaleCmd = &cobra.Command{
	Use:   "scale",
	Short: "Generate a scale pattern html document",
	Long: `Will use the config file to compute the desired scale pattern for your
instrument. The flag for name and root are required, while the flag for degree will allow
you to adjust the mode.`,
	Run: func(cmd *cobra.Command, args []string) {
		//Init the config data
		noteRing := initNoteRing()
		guitarSettings, err := initGuitarSettings(noteRing)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		validFlags := true
		name, _ := cmd.Flags().GetString("name")
		name = strings.ToLower(name)
		validFlags = validateName(name)
		if !validFlags {
			fmt.Printf("The specified scale name, %s, could not be resolved. Is it in the config file?\n", name)
			return
		}

		root, _ := cmd.Flags().GetString("root")
		root = strings.ToUpper(root)
		validFlags = validateNote(root, noteRing)
		if !validFlags {
			fmt.Printf("The specified root note, %s, is not a valid note.\n", root)
			return
		}

		degree, _ := cmd.Flags().GetInt("degree")
		validFlags = validateDegree(degree, name)
		if !validFlags {
			fmt.Printf("The specified degree, %v, is not valid.\n", degree)
			return
		}

		notesInScale, err := computeNotes(name, root, noteRing)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		fileName, _ := cmd.Flags().GetString("out")

		err = generateScaleChart(noteRing, notesInScale, guitarSettings, fileName)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
	},
}

func init() {
	RootCmd.AddCommand(scaleCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// scaleCmd.PersistentFlags().String("foo", "", "A help for foo")

	scaleCmd.Flags().StringP("name", "n", "", "Name of the scale to generate.")
	scaleCmd.Flags().StringP("root", "r", "", "Root note of the scale pattern.")
	scaleCmd.Flags().IntP("degree", "d", 1, "Degree of the root note.")
	scaleCmd.Flags().StringP("out", "o", "", "Filename to save output to.")
}

func validateName(name string) bool {
	scaleList := viper.GetStringMapString("scales")
	for scaleName, _ := range scaleList {
		if scaleName == name {
			return true
		}
	}
	return false
}

func validateNote(note string, noteRing *ring.Ring) bool {
	for i := 0; i < noteRing.Len(); i++ {
		if note == noteRing.Value {
			return true
		}
		noteRing = noteRing.Next()
	}
	return false
}

func validateDegree(degree int, scaleName string) bool {
	scaleList := viper.GetStringMapStringSlice("scales")
	scale := scaleList[scaleName]
	if degree > 0 && degree <= len(scale) {
		return true
	}
	return false
}

func computeNotes(scaleName string, root string, noteRing *ring.Ring) ([]string, error) {
	for noteRing.Value != root {
		noteRing = noteRing.Next()
	}

	scaleList := viper.GetStringMapStringSlice("scales")
	scale := scaleList[scaleName]

	var notes []string
	totalSemitones := 0

	for i := 0; i < len(scale); i++ {
		notes = append(notes, fmt.Sprintf("%v", noteRing.Value))

		interval, err := strconv.Atoi(scale[i])
		totalSemitones += interval
		if err != nil {
			err := errors.New("one of the scale integers in the config file could not be cast to an integer")
			return nil, err
		}

		if totalSemitones < 12 {
			for j := 0; j < interval; j++ {
				noteRing = noteRing.Next()
			}
		}
	}

	if totalSemitones != 12 {
		err := errors.New("the scale integers in the config file do not sum to 12")
		return nil, err
	}

	return notes, nil
}