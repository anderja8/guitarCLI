package cmd

import (
	"container/ring"
	"fmt"
	"strconv"
)

func generateScaleHTML(noteRing *ring.Ring, notesInScale []string, guitarSettings *GuitarSettings, scaleRoot string, intervals NoteIntervals) string {

	notesInScaleMap := make(map[string]int)
	for _, note := range notesInScale {
		notesInScaleMap[note] = 0
	}

	scaleHTML := `
<style>
	table {
		border-collapse: collapse;
	}

	tr.strikeout td:before {
	content: " ";
	position: absolute;
	top: 50%;
	left: 0;
	border-bottom: 1px solid #111;
	width: 100%;
}

	td {
		position: relative;
		padding: 5px 10px;
		border-left: 1px solid #111;
		border-right: 1px solid #111;
		text-align: center;
		font-weight: bold;
	}

	td.upper {
		border-left: none;
		border-right: none;
	}
</style>

<table>
`
	var stringCurrNote string

	for i := 0; i <= guitarSettings.numStrings; i++ {
		if i == 0 {
			scaleHTML += `<tr style="border-bottom: 1px solid #111;">
				`
			for j := 0; j <= guitarSettings.numFrets; j++ {
				scaleHTML += `<td class="upper">` + strconv.Itoa(j) + `</td>
					`
			}
		} else {
			if i == guitarSettings.numStrings {
				scaleHTML += `<tr class="strikeout"; style="border-bottom: 1px solid #111;">
				`
			} else {
				scaleHTML += `<tr class="strikeout";>
				`
			}

			stringCurrNote = guitarSettings.stringTunings[guitarSettings.numStrings-i]
			for stringCurrNote != noteRing.Value {
				noteRing = noteRing.Next()
			}

			for j := 0; j <= guitarSettings.numFrets; j++ {
				_, inScale := notesInScaleMap[stringCurrNote]
				if inScale { //stringCurrNote in notesInScale
					if stringCurrNote == scaleRoot {
						scaleHTML += `<td><span style="background-color:#5DB1D1; padding: 0px 5px;">` + stringCurrNote + `</span></td>
						`
					} else {
						scaleHTML += `<td><span style="background-color:#90C978; padding: 0px 5px;">` + stringCurrNote + `</span></td>
						`
					}
				} else {
					scaleHTML += `<td> </td>
					`
				}
				noteRing = noteRing.Next()
				stringCurrNote = fmt.Sprintf("%v", noteRing.Value)
			}
		}

		scaleHTML += `</tr>
`
	}
	scaleHTML += `</table>`

	return scaleHTML
}
