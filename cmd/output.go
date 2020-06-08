package cmd

import (
	"container/ring"
	"errors"
)

func generateScaleChart(noteRing *ring.Ring, notesInScale []string, guitarSettings *GuitarSettings, fileName string) error {
	/*<style>
		table {
			border: 1px solid black;
			border-collapse: collapse;
		}
	td {
		width: 75px;
		height: 75px;
		border: 1px solid black;
		text-align: center;
		font-weight: bold;
	}
	</style>

	<table style="border: 1px solid black;">
	<tr>
	<td> </td>
	<td>hi</td>
	<td style="background-color:#90C978;">A#</td>
	<td style="background-color:#5DB1D1;">B</td>
	</tr>
	</table>*/

	return errors.New("NYI")
}