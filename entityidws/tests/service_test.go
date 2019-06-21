package tests

import (
	"github.com/uvalib/entity-id-ws/entityidws/api"
	"github.com/uvalib/entity-id-ws/entityidws/client"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"testing"
)

type testConfig struct {
	Endpoint string
	Token    string
}

var cfg = loadConfig()

var plausableDoi = "doi:10.70020J7DN9V"
var badDoi = "badness"
var goodShoulder = "doi:10.70020"
var badShoulder = "abc:/blablabla"
var goodToken = cfg.Token
var badToken = "badness"
var empty = " "
var crossrefSchema = "crossref"
var dataciteSchema = "datacite"
var badSchema = "badschema"

func emptyField(field string) bool {
	return len(strings.TrimSpace(field)) == 0
}

func emptyPersonArray(array []api.Person) bool {
	return len(array) == 0
}

func emptyStringArray(array []string) bool {
	return len(array) == 0
}

//
// create a request based on the requested schema type
//
func createTestRequest(schema string) api.Request {

	request := api.Request{Schema: schema}

	if schema == dataciteSchema {
		request.DataCite = buildDataCiteSchema()
	}

	if schema == crossrefSchema {
		request.CrossRef = buildCrossRefSchema()
	}

	return request
}

//
// populate a crossref schema to use in a request
//
func buildCrossRefSchema() api.CrossRefSchema {

	return api.CrossRefSchema{
		Title:              "my crossref title",
		URL:                "http://google.com",
		Publisher:          "UVa Press",
		CreatorFirstName:   "Joe",
		CreatorLastName:    "Blow",
		CreatorDepartment:  "Math",
		CreatorInstitution: "UVa",
		PublicationDate:    "2001-01-01",
		ResourceType:       "Text",
	}
}

//
// populate a crossref schema to use in a request
//
func buildDataCiteSchema() api.DataCiteSchema {

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

func verifyCrossRefSchema(schema api.CrossRefSchema, t *testing.T) {

	if emptyField(schema.Title) ||
		//emptyField( schema.URL ) ||
		//emptyField( schema.Publisher ) ||
		emptyField(schema.CreatorFirstName) ||
		emptyField(schema.CreatorLastName) ||
		emptyField(schema.CreatorDepartment) ||
		emptyField(schema.CreatorInstitution) ||
		emptyField(schema.PublicationDate) {
		//emptyField( schema.ResourceType ) {

		t.Fatalf("Received incorrectly blank field in %v\n", schema)
	}
}

func verifyDataCiteSchema(schema api.DataCiteSchema, t *testing.T) {

	if emptyField(schema.Title) ||
		//emptyField( schema.URL ) ||
		emptyField(schema.Publisher) ||
		emptyField(schema.Abstract) ||
		//emptyPersonArray( schema.Creators ) ||
		//emptyPersonArray( schema.Contributors ) ||
		emptyField(schema.Rights) ||
		emptyStringArray(schema.Keywords) ||
		emptyStringArray(schema.Sponsors) ||
		emptyField(schema.Publisher) ||
		emptyField(schema.PublicationDate) ||
		emptyField(schema.GeneralType) ||
		emptyField(schema.ResourceType) {

		t.Fatalf("Received incorrectly blank field in %v\n", schema)
	}
}

func createGoodDoi(schema string, t *testing.T) string {

	status, entity := client.Create(cfg.Endpoint, goodShoulder, createTestRequest(schema), goodToken)
	if status == http.StatusOK {
		if len(entity.ID) == 0 {
			t.Fatalf("Create reported success but returned a blank DOI\n")
		}
		return entity.ID
	}

	t.Fatalf("Unable to create new DOI\n")
	return ""
}

func loadConfig() testConfig {

	data, err := ioutil.ReadFile("service_test.yml")
	if err != nil {
		log.Fatal(err)
	}

	var c testConfig
	if err := yaml.Unmarshal(data, &c); err != nil {
		log.Fatal(err)
	}

	log.Printf("endpoint [%s]\n", c.Endpoint)
	log.Printf("token    [%s]\n", c.Token)

	return c
}

//
// end of file
//
