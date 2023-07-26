package helper

import "time"

func CalculateFee(age int) float64 {
	fee := 0.0

	if age < 25 {
		fee = 50.0
	} else if age >= 25 && age < 40 {
		fee = 100.0
	} else {
		fee = 150.0
	}

	return fee
}

func CalculateAge(birthdateString string) int {
	layout := "2006-01-02"
	birthdate, _ := time.Parse(layout, birthdateString)

	now := time.Now()
	age := now.Year() - birthdate.Year()

	if now.YearDay() < birthdate.YearDay() {
		age--
	}

	return age
}
