package logger

import (
	"log"
	"os"
	"path"

	"github.com/Lequeston/bmstu-coursework-bd/config/env"
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

func LoggerInit(conf env.ApplicationConfig) {
	createLogFolder(path.Join(conf.RootDir, folderName))
	pathToRootFile := path.Join(conf.RootDir, folderName, rootLogFile)
	createLogFile(pathToRootFile)
	pathToTestFile := path.Join(conf.RootDir, folderName, testLogFile)
	createLogFile(pathToTestFile)

	fileWrite := pathToRootFile
	if conf.Mode == env.TEST_MODE {
		fileWrite = pathToTestFile
	}
	rootFile, err := os.OpenFile(fileWrite, os.O_WRONLY|os.O_CREATE, 0755)
	if err != nil {
		log.Fatal(err)
	}

	logrus.SetOutput(rootFile)
}
