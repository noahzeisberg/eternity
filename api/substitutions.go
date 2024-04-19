package main

import (
	"bytes"
	"encoding/json"
	"github.com/bytedance/sonic"
	"html"
	"io"
	"net/http"
	"regexp"
	"slices"
	"strconv"
	"strings"
	"time"
)

func GetSubstitutionPlan(datetime time.Time) (SubstitutionPlan, error) {
	requestBody, err := json.Marshal(map[string]any{
		"formatName":                 "Vetr_Kla",
		"date":                       datetime.Format("20060102"),
		"dateOffset":                 0,
		"strikethrough":              false,
		"mergeBlocks":                false,
		"showOnlyFutureSub":          false,
		"showBreakSupervisions":      false,
		"showTeacher":                true,
		"showClass":                  true,
		"showHour":                   true,
		"showInfo":                   true,
		"showRoom":                   true,
		"showSubject":                true,
		"groupBy":                    1,
		"hideAbsent":                 false,
		"departmentIds":              []int{},
		"departmentElementType":      -1,
		"hideCancelWithSubstitution": true,
		"hideCancelCausedByEvent":    true,
		"showTime":                   false,
		"showSubstText":              false,
		"showAbsentElements":         []int{},
		"showAffectedElements":       []int{1},
		"showUnitTime":               true,
		"showMessages":               true,
		"showStudentgroup":           false,
		"enableSubstitutionFrom":     false,
		"showSubstitutionFrom":       1800,
		"showTeacherOnEvent":         false,
		"showAbsentTeacher":          false,
		"strikethroughAbsentTeacher": false,
		"activityTypeIds":            []int{2, 3, 4, 8},
		"showEvent":                  true,
		"showCancel":                 true,
		"showOnlyCancel":             false,
		"showSubstTypeColor":         false,
		"showExamSupervision":        false,
		"showUnheraldedExams":        false,
	})
	if err != nil {
		return SubstitutionPlan{}, err
	}
	Logger.Info("Prepared body.")

	request, err := http.NewRequest("POST", "https://ajax.webuntis.com/WebUntis/monitor/substitution/data?school="+School, bytes.NewBuffer(requestBody))
	if err != nil {
		return SubstitutionPlan{}, err
	}

	request.Header.Set("Content-Type", "application/json")
	request.Close = true

	res, err := Client.Do(request)
	if err != nil {
		return SubstitutionPlan{}, err
	}

	Logger.Info("Retrieved substitutions.")
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return SubstitutionPlan{}, err
	}
	Logger.Info("Read body.")

	payload := UntisPayload{}
	err = sonic.Unmarshal(body, &payload)
	if err != nil {
		return SubstitutionPlan{}, err
	}
	Logger.Info("Parsed body.")

	date, err := time.Parse("20060102", strconv.Itoa(payload.Payload.Date))
	lastUpdate, err := time.Parse("02.01.2006 15:04:05", payload.Payload.LastUpdate)

	var substitutions []Substitution

	for _, row := range payload.Payload.Rows {
		data := row.Data

		classes := strings.Split(SanitizeHTML(data[1]), ", ")

		if !slices.Contains(classes, Class) {
			continue
		}

		var substitutionType string
		if data[5] == "" {
			substitutionType = "Vertretung"
		} else {
			substitutionType = data[5]
		}

		substitutions = append(substitutions, Substitution{
			Lesson:  SanitizeHTML(data[0]),
			Classes: classes,
			Subject: SanitizeHTML(data[2]),
			Rooms:   strings.Split(SanitizeHTML(data[3]), ", "),
			Teacher: strings.Split(SanitizeHTML(data[4]), ", "),
			Type:    SanitizeHTML(substitutionType),
		})
	}

	substitutionPlan := SubstitutionPlan{
		Date:          date,
		LastUpdate:    lastUpdate,
		Affected:      slices.Contains(payload.Payload.AffectedElements.List, Class),
		Substitutions: substitutions,
	}

	return substitutionPlan, nil
}

func SanitizeHTML(str string) string {
	return html.UnescapeString(regexp.MustCompile("<[^>]*>").ReplaceAllString(str, ""))
}

type UntisPayload struct {
	Payload struct {
		Date       int    `json:"date"`
		NextDate   int    `json:"nextDate"`
		LastUpdate string `json:"lastUpdate"`

		Rows []struct {
			Data []string `json:"data"`
		} `json:"rows"`

		AffectedElements struct {
			List []string `json:"1"`
		} `json:"affectedElements"`
	} `json:"payload"`
}
