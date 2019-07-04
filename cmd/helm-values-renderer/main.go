package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"os"
)

func main() {
	commit := flag.String("commit", "", "used to tag the values.yaml file")
	tag := flag.String("tag", "", "used to select which tag to use when creating a values.yaml file")
	image := flag.String("image", "", "used to select which image to use when creating a values.yaml file")
	flag.Parse()

	if *commit == "" {
		log.Fatal("please supply a value to the flag --commit")
	}
	if *image == "" {
		log.Fatal("please supply a value to the flag --image")
	}
	if *tag == "" {
		log.Fatal("please supply a value to the tag --image")
	}

	deployFile, err := os.Create(fmt.Sprintf("build/package/helm/values-%s.yaml", *commit))
	if err != nil {
		log.Fatal(err)
	}

	tmpl, err := template.ParseFiles("build/package/helm/templates/values.gotemplate.yaml")
	if err != nil {
		log.Fatal(err)
	}

	if err := tmpl.Execute(deployFile, struct {
		Image string
		Tag   string
	}{Image: *image, Tag: *tag}); err != nil {
		log.Fatal(err)
	}
}
