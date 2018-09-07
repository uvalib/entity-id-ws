package main

import (
	"bufio"
	"github.com/uvalib/entity-id-ws/entityidws/client"
	"flag"
	"fmt"
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
		fmt.Printf("Get metadata for a set of DOI's\n")
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
		status, entity := client.Get(cfg.Endpoint, doi, cfg.Token)
		if status == expected {
			fmt.Printf("DOI: %s\n", entity.ID)
			fmt.Printf(" Schema:          %s\n", entity.Schema)
			fmt.Printf(" URL:             %s\n", entity.DataCite.URL)
			fmt.Printf(" Title:           %s\n", entity.DataCite.Title)
			fmt.Printf(" Abstract:        %s\n", entity.DataCite.Abstract)
			fmt.Printf(" Creators:        %t\n", entity.DataCite.Creators)
			fmt.Printf(" Contributors:    %t\n", entity.DataCite.Contributors)
			fmt.Printf(" Rights:          %s\n", entity.DataCite.Rights)
			fmt.Printf(" Keywords:        %t\n", entity.DataCite.Keywords)
			fmt.Printf(" Sponsors:        %t\n", entity.DataCite.Sponsors)
			fmt.Printf(" Publisher:       %s\n", entity.DataCite.Publisher)
			fmt.Printf(" PublicationDate: %s\n", entity.DataCite.PublicationDate)
			fmt.Printf(" GeneralType:     %s\n", entity.DataCite.GeneralType)
			fmt.Printf(" ResourceType:    %s\n", entity.DataCite.ResourceType)

		} else {
			fmt.Printf("ERROR: getting %s. Expected %v, got %v\n", doi, expected, status)
			if ignoreError == false {
				os.Exit(status)
			}
		}
	}
	os.Exit(0)
}

func loadConfig() testConfig {

	data, err := ioutil.ReadFile("src/entityidws/tools/bulk-get/config.yml")
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
