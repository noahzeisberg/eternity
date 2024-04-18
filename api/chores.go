package main

import "time"

func GetChores() (string, string) {
	students := []string{
		"Dominik",
		"Hussain",
		"Halil",
		"Simon",
		"Daniel",
		"Massih",
		"Lea",
		"Manuel",
		"Mikail",
		"Asia",
		"Elina",
		"Carolina",
		"Lana",
		"Luca",
		"Lorena",
		"Linus",
		"Enrico",
		"Felix",
		"Laurin",
		"Emina",
		"Naz Fatma",
		"Jannik",
		"Christian",
		"Elias",
		"Jeremy",
		"Mia",
		"Noah",
	}
	var cycle int
	weekOfYear := WeekOfYear()
	var groups [][]string

	if len(students)%2 != 0 {
		students = append(students, students...)
	}

	var even []string
	var odd []string

	for i, student := range students {
		if i%2 == 0 {
			even = append(even, student)
		} else {
			odd = append(odd, student)
		}
	}

	for i := range len(even) {
		groups = append(groups, []string{even[i], odd[i]})
	}
	cycle = weekOfYear % len(groups)

	return groups[cycle][0], groups[cycle][1]
}

func WeekOfYear() int {
	now := time.Now()
	startOfYear := time.Date(now.Year(), 0, 0, 0, 0, 0, 0, now.Location())
	diff := now.Sub(startOfYear)
	oneWeek := time.Duration(1000*60*60*24*7) * time.Millisecond
	return int(diff / oneWeek)
}
