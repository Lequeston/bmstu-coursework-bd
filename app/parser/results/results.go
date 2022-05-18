package results

import (
	"encoding/csv"
	"os"
	"reflect"
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

func ParseResults(fileName string) []Result {
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)

	records, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}

	// В первой записи содержится только информация о полях, поэтому она нам не нужна
	records = records[1:]

	results := make([]Result, 0)

	infoFieldAmount := reflect.TypeOf(Information{}).NumField()
	questionFieldAmount := reflect.TypeOf(Question{}).NumField()

	for _, record := range records {
		info := record[:infoFieldAmount]
		questions := record[infoFieldAmount:]

		var result Result
		var resultInformation Information
		resultQuestions := make([]Question, 0)

		resultInformationReflect := reflect.ValueOf(&resultInformation).Elem()
		for i, infoField := range info {
			resultInformationReflect.Field(i).SetString(infoField)
		}

		for i := 0; i < questionFieldAmount; i++ {
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

	return results
}
