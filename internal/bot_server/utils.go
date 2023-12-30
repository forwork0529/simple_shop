package bot_server

import (
	"io/ioutil"
	"log"
	"strings"
)

// Функция возвращающая список именён файлов в папке и значение говорящее пуст ли список
func returnFilesInPath(path string) ([]string, bool) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Printf("Cant read direcory: %v\n", err.Error())
	}
	var fileNames []string
	for _, fileInfo := range files {
		fileNames = append(fileNames, fileInfo.Name())
	}
	if len(fileNames) < 1 {
		return fileNames, true
	}
	return fileNames, false

}

func TrimSuffix(toPrintProductName string) string {
	if idx := strings.IndexByte(toPrintProductName, '.'); idx >= 0 {
		toPrintProductName = toPrintProductName[:idx]
	}
	return toPrintProductName
}
