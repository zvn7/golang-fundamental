package library

const secondsInMinute int = 60
const minutesInHour int = 60
const HourInADay int = 24

func daysToHours(day int) int {
	return day * HourInADay
}

func DaystoMinutes(day int) int {
	return day * minutesInHour * secondsInMinute
}