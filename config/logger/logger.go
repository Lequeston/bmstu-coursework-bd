package logger

import (
	"fmt"
	"log"
	"os"

	"github.com/sirupsen/logrus"
)

func createLogFolder(folderName string) {
	err := os.Mkdir(folderName, 0755)
	if err != nil {
		log.Println(err)
	}
}

func createLogFile(filename string) {
	file, err := os.Create(filename)
	if err != nil {
		log.Println(err)
	}
	defer file.Close()
}

func LoggerInit() {
	rootDir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	pathToRootFile := fmt.Sprintf("%s/%s/%s", rootDir, folderName, rootLogFile)
	log.Printf("Path to root log file: %s", pathToRootFile)
	createLogFolder(folderName)
	createLogFile(pathToRootFile)

	rootFile, err := os.OpenFile(pathToRootFile, os.O_WRONLY|os.O_CREATE, 0755)
	if err != nil {
		log.Fatal(err)
	}

	logrus.SetOutput(rootFile)
}
