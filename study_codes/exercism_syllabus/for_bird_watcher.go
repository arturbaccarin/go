package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
	var sumBirds int = 0
	for _, birdsCount := range birdsPerDay {
		sumBirds += birdsCount
	}
	return sumBirds
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
	var weekFirstDay int
	var weekLastDay int

	weekLastDay = week * 7
	weekFirstDay = weekLastDay - 7

	var birdsWeek []int = birdsPerDay[weekFirstDay:weekLastDay]
	var sumBirds int = 0

	for _, birdsCount := range birdsWeek {
		sumBirds += birdsCount
	}

	return sumBirds
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	var birdsLength int = len(birdsPerDay)
	for i := 0; i < birdsLength; i++ {
		if i%2 == 0 {
			birdsPerDay[i] += 1
		}
	}

	return birdsPerDay
}
