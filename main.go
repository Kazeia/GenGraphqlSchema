package main

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

func getFilesPath(ext string, directoryPaths ...string) []string {
	var files []string

	for i := 0; i < len(directoryPaths); i++ {
		err := filepath.Walk(directoryPaths[i], func(path string, f os.FileInfo, err error) error {
			if filepath.Ext(path) == ext {
				files = append(files, path)
			}
			return nil
		})
		if err != nil {
			log.Fatal("ERROR-FATAL: main-getFilesPath-Walk, " + err.Error())
		}
	}

	return files
}

func main() {
	schemaPath:=os.Args[1]
	filePaths := getFilesPath(".graphql", schemaPath)
	schema := "package main\nconst gqlRawSchema=`"

	for i := 0; i < len(filePaths); i++ {
		content, err := ioutil.ReadFile(filePaths[i])
		if err != nil {
			log.Fatal("ERROR-FATAL: main-ReadFile, " + err.Error())
		}
		schema += string(content) + "\n"
	}
	schema += "`"

	ioutil.WriteFile("./schema.go", []byte(schema), 0644)
}
