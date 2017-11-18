package model

import (
	"testing"
	"log"
)

func TestComposeLoad(t *testing.T) {
	conposes := ComposeLoad("../conf")
	log.Println(conposes)
}
func TestLoadComposeFiles(t *testing.T) {
	files := LoadComposeFiles("../conf")
	for _, file := range files {
		log.Println(file)
	}
}