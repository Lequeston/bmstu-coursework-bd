package results

import (
	"encoding/csv"
	"os"
	"reflect"

	log "github.com/sirupsen/logrus"
)

type Question struct {
	Question    string
	Answer      string
	RightAnswer string
}

type Information struct {
	Surname     string
	Name        string
	ID          string
	Department  string
	Institution string
	Address     string
	Email       string
	Username    string
	State       string
	Started     string
	Completed   string
	Time        string
	Grade       string
}

type Result struct {
	Information Information
	Questions   []Question
}

func ParseResults(fileName string) ([]Result, error) {
	file, err := os.Open(fileName)
	if err != nil {
		log.WithField("filename", fileName).Errorf("Failed to open result file: %s", err)
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)

	records, err := reader.ReadAll()
	if err != nil {
		log.Errorf("Failed to create result file reader: %s", err)
		return nil, err
	}

	// В первой записи содержится только информация о полях, поэтому она нам не нужна
	records = records[1:]

	results := make([]Result, 0, len(records))

	infoFieldAmount := reflect.TypeOf(Information{}).NumField()
	questionFieldAmount := reflect.TypeOf(Question{}).NumField()

	for _, record := range records {
		info := record[:infoFieldAmount]
		questions := record[infoFieldAmount:]

		questionsAmount := len(questions) / questionFieldAmount

		var result Result
		var resultInformation Information
		resultQuestions := make([]Question, 0, questionsAmount)

		resultInformationReflect := reflect.ValueOf(&resultInformation).Elem()
		for i, infoField := range info {
			resultInformationReflect.Field(i).SetString(infoField)
		}

		for i := 0; i < questionsAmount; i++ {
			var question Question
			questionReflect := reflect.ValueOf(&question).Elem()
			for j := 0; j < questionFieldAmount; j++ {
				questionReflect.Field(j).SetString(questions[j+(i*questionFieldAmount)])
			}
			resultQuestions = append(resultQuestions, question)
		}

		result.Information = resultInformation
		result.Questions = resultQuestions
		results = append(results, result)
	}

	return results, nil
}
