package model

import (
	"log"
	"os"
	"encoding/json"
	"path/filepath"
)

type Compose struct {
	Path string `json:"path"`
	Meta Meta `json:"meta"`
	Decors []Decor `json:"decors"`
}
type Meta struct {
	Module string `json:"module"`
	Name   string `json:"name"`
	Memo   string `json:"memo"`
	Uri    Uri `json:"uri"`
}

type Uri struct {
	Parameters []Parameter
}
type Parameter struct {
	Field string `json:"field"`
	Name  string `json:"name"`
	Type  string `json:"type"`
	Memo  string `json:"memo"`
}
type Decor struct {
	Field string `json:"field"`
	Source Source `json:"source"`
	Decors []Decor `json:"decors"`

}
type Source struct {
	Base string `json:"base"`
	Uri string `json:"uri"`
	OnError OnError `json:"decors"`
}
type OnError struct {
	Type string `json:"type"`
}
func ComposeLoad(rootPath string) []Compose {
	var composes []Compose
	filePaths := LoadComposeFiles(rootPath)
	for _, filePath := range filePaths {
		var compose Compose
		log.Println(filePath)
		file,err := os.Open(filePath)
		if err != nil {
			log.Fatalln("open file failed")
		}
		jsonParser := json.NewDecoder(file)
		jsonParser.Decode(&compose)
		composes = append(composes,compose)
	}
	return composes
}
func LoadComposeFiles(rootPath string) ([]string){
	var files []string
	rootPath,err := filepath.Abs(rootPath)
	if err != nil {
		log.Panic(err)
		return nil
	}
	filepath.Walk(rootPath, visit(&files))
	return files
}
func visit(files *[]string) filepath.WalkFunc {
	return func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Panic(err)
			return err
		}
		fileInfo, err := os.Stat(path)
		if err != nil {
			log.Panic(err)
			return err
		}
		if !fileInfo.IsDir() {
			*files = append(*files, path)
		}
		return nil
	}
}
