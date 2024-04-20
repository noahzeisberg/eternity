package api

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

func Handler(w http.ResponseWriter, r *http.Request) {
	handler := ResponseHandler{w: w, r: r}
	client := http.Client{Transport: &http.Transport{}}

	t, err := strconv.Atoi(r.FormValue("t"))
	class := r.FormValue("class")
	if err != nil {
		return
	}

	datetime := time.Now()
	datetime.AddDate(0, 0, t)

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
		handler.RespondError(err)
		return
	}

	request, err := http.NewRequest("POST", "https://ajax.webuntis.com/WebUntis/monitor/substitution/data?school=hbs-hattersheim", bytes.NewBuffer(requestBody))
	if err != nil {
		handler.RespondError(err)
		return
	}

	request.Header.Set("Content-Type", "application/json")
	request.Close = true

	res, err := client.Do(request)
	if err != nil {
		handler.RespondError(err)
		return
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		handler.RespondError(err)
		return
	}

	payload := UntisPayload{}
	err = sonic.Unmarshal(body, &payload)
	if err != nil {
		handler.RespondError(err)
		return
	}

	date, err := time.Parse("20060102", strconv.Itoa(payload.Payload.Date))
	lastUpdate, err := time.Parse("02.01.2006 15:04:05", payload.Payload.LastUpdate)

	var substitutions []Substitution

	for _, row := range payload.Payload.Rows {
		data := row.Data

		classes := strings.Split(SanitizeHTML(data[1]), ", ")

		if !slices.Contains(classes, class) {
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
		Affected:      slices.Contains(payload.Payload.AffectedElements.List, class),
		Substitutions: substitutions,
	}

	handler.Respond(substitutionPlan)
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

type ResponseHandler struct {
	w http.ResponseWriter
	r *http.Request
}

func (h ResponseHandler) Respond(object any) {
	b, err := json.Marshal(object)
	if err != nil {
		return
	}
	_, err = h.w.Write(b)
	if err != nil {
		return
	}
}

func (h ResponseHandler) RespondError(err error) {
	b, err := json.Marshal(map[string]string{
		"error": err.Error(),
	})
	if err != nil {
		return
	}
	_, err = h.w.Write(b)
	if err != nil {
		return
	}
}
