package main

func main() {
	day := 3

	switch day {
	case 1:
		println("Monday")
	case 2:
		println("Tuesday")
	case 3:
		println("Wednesday")
	case 4:
		println("Thursday")
	case 5:
		println("Friday")
	case 6:
		println("Saturday")
	case 7:
		println("Sunday")
	default:
		println("Invalid day")
	}

	days := "Monday"

	switch days {
	case "Monday", "Tuesday", "Wednesday", "Thursday", "Friday":
		println("Weekday")
	case "Saturday", "Sunday":
		println("Weekend")
	}

	temperature := 32

	switch {
	case temperature < 0:
		println("Freezing")
	case temperature >= 0 && temperature < 10:
		println("Very cold")
	case temperature >= 10 && temperature < 20:
		println("Cold")
	case temperature >= 20 && temperature < 30:
		println("Normal")
	case temperature >= 30 && temperature < 40:
		println("Hot")
	case temperature >= 40:
		println("Very hot")
	default:
		println("Invalid temperature")
	}
}
