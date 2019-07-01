package main

import (
	"flag"
	"html/template"
	"log"
	"os"
)

func main() {
	tag := flag.String("tag", "", "used to select which tag to render out a kubectl deployment file from")
	flag.Parse()

	if *tag == "" {
		log.Fatal("please supply a value to the flag --tag")
	}

	deployFile, err := os.Create("build/package/kubernetes/deployment.yaml")
	if err != nil {
		log.Fatal(err)
	}

	tmpl, err := template.ParseFiles("build/package/kubernetes/templates/deployment.gotemplate.yaml")
	if err != nil {
		log.Fatal(err)
	}

	if err := tmpl.Execute(deployFile, struct{ Tag string }{Tag: *tag}); err != nil {
		log.Fatal(err)
	}
}
