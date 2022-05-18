package results

import (
	"fmt"
	"os"
	"path"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

var records = [][]string{
	{
		"Сазонов",
		"Даниил",
		"18У667",
		"ИУ6-43Б",
		"Бакалавриат",
		"ИУ6",
		"sazonovdd@student.bmstu.ru",
		"sdd18u667",
		"Завершено",
		"18 марта 2022  16:12",
		"18 марта 2022  16:13",
		"18 сек.",
		"Еще не оценено",
		"Верно?",
		"Неверно",
		"Неверно",
		"Дайте определение БД и СУБД",
		"фффф",
		"Таблицы и связи между ними",
		"Вручную",
		"фффф",
		"-",
	},
	{
		"Логачев",
		"Кирилл",
		"18У354",
		"ИУ6-43Б",
		"Бакалавриат",
		"ИУ6",
		"logachevka@student.bmstu.ru",
		"lka18u354",
		"Завершено",
		"18 марта 2022  16:15",
		"18 марта 2022  16:15",
		"25 сек.",
		"Еще не оценено",
		"Верно?",
		"Неверно",
		"Неверно",
		"Дайте определение БД и СУБД",
		"ыфвф",
		"Таблицы и связи между ними",
		"Вручную",
		"фыфвффывф",
		"-",
	},
}

var results = []Result{
	{
		Information{
			"Сазонов",
			"Даниил",
			"18У667",
			"ИУ6-43Б",
			"Бакалавриат",
			"ИУ6",
			"sazonovdd@student.bmstu.ru",
			"sdd18u667",
			"Завершено",
			"18 марта 2022  16:12",
			"18 марта 2022  16:13",
			"18 сек.",
			"Еще не оценено",
		},
		[]Question{
			{
				"Верно?",
				"Неверно",
				"Неверно",
			},
			{
				"Дайте определение БД и СУБД",
				"фффф",
				"Таблицы и связи между ними",
			},
			{
				"Вручную",
				"фффф",
				"-",
			},
		},
	},
	{
		Information{
			"Логачев",
			"Кирилл",
			"18У354",
			"ИУ6-43Б",
			"Бакалавриат",
			"ИУ6",
			"logachevka@student.bmstu.ru",
			"lka18u354",
			"Завершено",
			"18 марта 2022  16:15",
			"18 марта 2022  16:15",
			"25 сек.",
			"Еще не оценено",
		},
		[]Question{
			{
				"Верно?",
				"Неверно",
				"Неверно",
			},
			{
				"Дайте определение БД и СУБД",
				"ыфвф",
				"Таблицы и связи между ними",
			},
			{
				"Вручную",
				"фыфвффывф",
				"-",
			},
		},
	},
}

func TestGetRecords(t *testing.T) {

	rootDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	fileName := path.Join(rootDir, "test", "response.csv")
	fmt.Println(fileName)

	Convey("While reading from file ", t, func() {
		testRecords, err := GetRecords(fileName)

		Convey("error should be nil", func() {
			So(err, ShouldBeNil)
		})

		Convey("file should be readed into [][]string", func() {
			So(testRecords, ShouldResemble, records)
		})
	})
}

func TestParseRecords(t *testing.T) {
	Convey("ParseRecords should parse records into []Result structure", t, func() {
		So(ParseRecords(records), ShouldResemble, results)
	})
}
