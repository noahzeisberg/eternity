package main

import (
	"time"
)

type SubstitutionPlan struct {
	Date          time.Time      `json:"date"`
	LastUpdate    time.Time      `json:"last_update"`
	Affected      bool           `json:"affected"`
	Substitutions []Substitution `json:"substitutions"`
}

type Substitution struct {
	Lesson  string   `json:"lesson"`
	Classes []string `json:"classes"`
	Subject string   `json:"subject"`
	Rooms   []string `json:"room"`
	Teacher []string `json:"teacher"`
	Type    string   `json:"type"`
}

type ChoreList struct {
	Cycle    int      `json:"cycle"`
	Students []string `json:"students"`
}
