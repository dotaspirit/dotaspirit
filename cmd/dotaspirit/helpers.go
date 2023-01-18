package main

func secondsToTime(timer int) (hours, minutes, seconds int) {
	minutes = timer / 60
	seconds = timer % 60
	hours = minutes / 60
	minutes = minutes % 60
	return hours, minutes, seconds
}
