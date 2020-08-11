package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"text/template"
	"time"
)

const (
	fileName    = "README.md"
	name        = "Lee Sander"
	description = "Check out my website and get to know more about me and the cool stuff I am working on or connect with me!"
)

var (
	intros = []string{
		"Hey I am",
		"Hey there I'm",
		"Hi my name is",
		"Hi there I'm",
		"Yo I'm",
		"What's up! I'm",
		"What's up! My name is",
		"Howdy, I am",
		"Howdy, my name is",
		"Hello there, I am",
		"Hello there, my name is",
		"Bonjour, my name is",
		"What's happening! I am",
		"What's happening! my name is",
		"Hola! my name is",
		"How's it going? My name is",
		"How's it going? I am",
	}
)

type Title struct {
	Intro string
	Name  string
}

type Readme struct {
	Title       Title
	Description string
	UpdatedAt   string
}

func main() {
	tmpl := fmt.Sprintf("./%s.tmpl", fileName)
	t, err := template.ParseFiles(tmpl)
	if err != nil {
		log.Fatal(err)
	}

	f, err := os.Create(fileName)
	if err != nil {
		log.Fatal(err)
	}
	rand.Seed(time.Now().Unix())

	readme := Readme{
		Title: Title{
			Intro: intros[rand.Intn(len(intros))],
			Name:  name,
		},
		Description: description,
		UpdatedAt:   time.Now().UTC().Format("Mon Jan _2 15:04:05 2006"),
	}

	err = t.Execute(f, readme)
	if err != nil {
		log.Fatal(err)
	}

}
