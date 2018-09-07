package main

import (
	"github.com/uvalib/entity-id-ws/entityidws/api"
	"github.com/uvalib/entity-id-ws/entityidws/client"
	"flag"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
)

type testConfig struct {
	Endpoint string
	Token    string
	Shoulder string
}

var cfg = loadConfig()

func main() {

	if len(os.Args) == 1 {
		fmt.Printf("Mint a new block of DOI's\n")
		fmt.Printf("use: %s [-ignore] <count>\n", os.Args[0])
		os.Exit(0)
	}

	var ignoreError bool
	flag.BoolVar(&ignoreError, "ignore", false, "Ignore errors")
	flag.Parse()

	count, _ := strconv.Atoi(os.Args[len(os.Args)-1])
	for current := 0; current < count; current++ {

		expected := http.StatusOK

		request := api.Request{Schema: "datacite", DataCite: makeMintPayload()}
		status, entity := client.Create(cfg.Endpoint, cfg.Shoulder, request, cfg.Token)
		if status == expected {
			fmt.Printf("%03d -> %s\n", current+1, entity.ID)
		} else {
			fmt.Printf("ERROR minting. Expected %v, got %v\n", expected, status)
			if ignoreError == false {
				os.Exit(status)
			}
		}
	}
	os.Exit(0)
}

func makeMintPayload() api.DataCiteSchema {

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

	data, err := ioutil.ReadFile("src/entityidws/tools/bulk-mint/config.yml")
	if err != nil {
		log.Fatal(err)
	}

	var c testConfig
	if err := yaml.Unmarshal(data, &c); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("endpoint [%s]\n", c.Endpoint)
	//fmt.Printf("token    [%s]\n", c.Token )
	fmt.Printf("shoulder [%s]\n", c.Shoulder)

	return c
}

//
// end of file
//
