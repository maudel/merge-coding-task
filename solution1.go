package main

import "sort"


func merge(intervals [][]int) [][]int {
	/*
		Wenn die Größe 1 oder weniger beträgt, geben Sie sie einfach zurück
	*/
	if len(intervals) < 2 {
		return intervals
	}

	/*
		Sort the intervals in increasing order
	*/
	sort.SliceStable(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	/*
	 Erstellen Sie einen leeren Stapel von Intervallen
	*/
	var res [][]int
	res = append(res,intervals[0])
	for i := 1 ; i < len(intervals); i++{
		n := len(res) - 1
		// Intervall vom obersten Stapel holen
		top := res[n]
		// wenn das aktuelle Intervall sich nicht mit der Stapeloberseite überlappt,
		// schieben Sie es auf den Stapel
		if top[1] < intervals[i][0] {
			res = append(res, intervals[i])
		} else
		// Andernfalls aktualisieren Sie die Endzeit von oben, wenn das Ende der aktuellen
		// Intervall ist mehr
		if top[1] < intervals[i][1] {
			top[1] = intervals[i][1]
			res = res[:n]
			res = append(res, top)
		}
	}
	return res
}