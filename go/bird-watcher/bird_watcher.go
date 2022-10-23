package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
	result := 0

	for _, value := range birdsPerDay {
		result += value
	}

	return result
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
	dayStart := (week - 1) * 7
	dayEnd := week * 7

	result := 0
	for day, value := range birdsPerDay {
		if day < dayStart || day > dayEnd {
			continue
		}
		result += value
	}

	return result
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	for day, _ := range birdsPerDay {
		if day%2 != 0 {
			continue
		}
		birdsPerDay[day] += 1
	}

	return birdsPerDay
}
