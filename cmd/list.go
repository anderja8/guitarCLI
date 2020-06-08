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
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"sort"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all scales that can be generated",
	Long: `This command will list all scales that can currently be generated.
New scales can be added to the config file.`,
	Run: func(cmd *cobra.Command, args []string) {
		scaleList := viper.GetStringMapString("scales")
		fmt.Printf("Key List:\n")
		var scaleSlice []string
		for key, _ := range scaleList {
			scaleSlice = append(scaleSlice, key)
		}
		sort.Strings(scaleSlice)
		for _, scale := range scaleSlice {
			fmt.Printf("\t%v\n", scale)
		}
	},
}

func init() {
	RootCmd.AddCommand(listCmd)
}
