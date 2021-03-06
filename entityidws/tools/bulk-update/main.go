package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/uvalib/entity-id-ws/entityidws/api"
	"github.com/uvalib/entity-id-ws/entityidws/client"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type testConfig struct {
	Endpoint string
	Token    string
}

var cfg = loadConfig()

func main() {

	if len(os.Args) == 1 {
		fmt.Printf("Update metadata for a set of DOI's\n")
		fmt.Printf("use: %s [-ignore] <file>\n", os.Args[0])
		os.Exit(0)
	}

	var ignoreError bool
	flag.BoolVar(&ignoreError, "ignore", false, "Ignore errors")
	flag.Parse()

	file, err := os.Open(os.Args[len(os.Args)-1])
	if err != nil {
		fmt.Printf("%s\n", err)
		os.Exit(1)
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		expected := http.StatusOK

		doi := scanner.Text()
		entity := api.Request{ID: doi, Schema: "datacite", DataCite: makeUpdatePayload()}
		status := client.Update(cfg.Endpoint, entity, cfg.Token)
		if status == expected {
			fmt.Printf("Updated %s\n", doi)
		} else {
			fmt.Printf("ERROR: updating %s. Expected %v, got %v\n", doi, expected, status)
			if ignoreError == false {
				os.Exit(status)
			}
		}
	}
	os.Exit(0)
}

func makeUpdatePayload() api.DataCiteSchema {

	person1 := api.Person{FirstName: "John", LastName: "Smith", Department: "Biology", Institution: "UVa"}
	person2 := api.Person{FirstName: "Joe", LastName: "Blow", Department: "History", Institution: "UVa"}

	return api.DataCiteSchema{
		Title:           "my datacite title",
		URL:             "http://google.com",
		Abstract:        "my interesting abstract",
		Creators:        []api.Person{person1},
		Contributors:    []api.Person{person2},
		Rights:          "All rights reserved",
		Keywords:        []string{"keyword1", "keyword2"},
		Sponsors:        []string{"sponsor1", "sponsor2"},
		Publisher:       "UVa Press",
		PublicationDate: "2002-02-02",
		GeneralType:     "Sound",
		ResourceType:    "Audio",
	}
}

func loadConfig() testConfig {

	data, err := ioutil.ReadFile("src/entityidws/tools/bulk-update/config.yml")
	if err != nil {
		log.Fatal(err)
	}

	var c testConfig
	if err := yaml.Unmarshal(data, &c); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("endpoint [%s]\n", c.Endpoint)
	//fmt.Printf("token    [%s]\n", c.Token )

	return c
}

//
// end of file
//
