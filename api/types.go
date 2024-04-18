package main

import (
	"slices"
	"time"
)

type SubstitutionPlan struct {
	Date          time.Time      `json:"date"`
	LastUpdate    time.Time      `json:"last_update"`
	Affected      []string       `json:"affected"`
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

func (substitutionPlan SubstitutionPlan) IsAffected(class string) bool {
	return slices.Contains(substitutionPlan.Affected, class)
}

func (substitutionPlan SubstitutionPlan) GetSubstitutions(class string) []Substitution {
	var substitutions []Substitution
	for _, substitution := range substitutionPlan.Substitutions {
		if slices.Contains(substitution.Classes, class) {
			substitutions = append(substitutions, substitution)
		}
	}
	return substitutions
}
