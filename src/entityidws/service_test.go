package main

import (
	"entityidws/api"
	"entityidws/client"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"testing"
	//"fmt"
)

type TestConfig struct {
	Endpoint string
	Token    string
}

var cfg = loadConfig()

var plausableDoi = "doi:10.5072/FK2QJ7DN9V"
var badDoi = "badness"
var goodShoulder = "doi:10.5072/FK2"
var badShoulder = "abc:/blablabla"
var goodToken = cfg.Token
var badToken = "badness"
var empty = " "
var crossrefSchema = "crossref"
var dataciteSchema = "datacite"
var badSchema = "badschema"

//
// healthcheck tests
//

func TestHealthCheck(t *testing.T) {
	expected := http.StatusOK
	status := client.HealthCheck(cfg.Endpoint)
	if status != expected {
		t.Fatalf("Expected %v, got %v\n", expected, status)
	}
}

//
// version tests
//

func TestVersionCheck(t *testing.T) {
	expected := http.StatusOK
	status, version := client.VersionCheck(cfg.Endpoint)
	if status != expected {
		t.Fatalf("Expected %v, got %v\n", expected, status)
	}

	if len(version) == 0 {
		t.Fatalf("Expected non-zero length version string\n")
	}
}

//
// runtime tests
//

func TestRuntimeCheck(t *testing.T) {
	expected := http.StatusOK
	status, runtime := client.RuntimeCheck(cfg.Endpoint)
	if status != expected {
		t.Fatalf("Expected %v, got %v\n", expected, status)
	}

	if runtime == nil {
		t.Fatalf("Expected non-nil runtime info\n")
	}

	if runtime.AllocatedMemory == 0 ||
		runtime.CpuCount == 0 ||
		runtime.GoRoutineCount == 0 ||
		runtime.ObjectCount == 0 {
		t.Fatalf("Expected non-zero value in runtime info but one is zero\n")
	}
}

//
// statistics tests
//

func TestStatistics(t *testing.T) {
	expected := http.StatusOK
	status, _ := client.Statistics(cfg.Endpoint)
	if status != expected {
		t.Fatalf("Expected %v, got %v\n", expected, status)
	}

	//if len( version ) == 0 {
	//    t.Fatalf( "Expected non-zero length version string\n" )
	//}
}

//
// DOI get tests
//

func TestGetCrossRef(t *testing.T) {

	doi := createGoodDoi(crossrefSchema, t)
	expected := http.StatusOK
	status, response := client.Get(cfg.Endpoint, doi, goodToken)
	if status != expected {
		t.Fatalf("Expected %v, got %v\n", expected, status)
	}

	if response == nil {
		t.Fatalf("Expected to find entity %v and did not\n", doi)
	}

	if response.Schema != crossrefSchema {
		t.Fatalf("Received unexpected schema in response\n")
	}
	if emptyField(response.Id) {
		t.Fatalf("Received blank ID in response\n")
	}
	verifyCrossRefSchema(response.CrossRef, t)
}

func TestGetDataCite(t *testing.T) {

	doi := createGoodDoi(dataciteSchema, t)
	expected := http.StatusOK
	status, response := client.Get(cfg.Endpoint, doi, goodToken)
	if status != expected {
		t.Fatalf("Expected %v, got %v\n", expected, status)
	}

	if response == nil {
		t.Fatalf("Expected to find entity %v and did not\n", doi)
	}

	if response.Schema != dataciteSchema {
		t.Fatalf("Received unexpected schema in response\n")
	}
	if emptyField(response.Id) {
		t.Fatalf("Received blank ID in response\n")
	}
	verifyDataCiteSchema(response.DataCite, t)
}

func TestGetEmptyId(t *testing.T) {
	expected := http.StatusBadRequest
	status, _ := client.Get(cfg.Endpoint, empty, goodToken)
	if status != expected {
		t.Fatalf("Expected %v, got %v\n", expected, status)
	}
}

func TestGetBadId(t *testing.T) {
	expected := http.StatusBadRequest
	status, _ := client.Get(cfg.Endpoint, badDoi, goodToken)
	if status != expected {
		t.Fatalf("Expected %v, got %v\n", expected, status)
	}
}

func TestGetEmptyToken(t *testing.T) {
	expected := http.StatusBadRequest
	status, _ := client.Get(cfg.Endpoint, plausableDoi, empty)
	if status != expected {
		t.Fatalf("Expected %v, got %v\n", expected, status)
	}
}

func TestGetBadToken(t *testing.T) {
	expected := http.StatusForbidden
	status, _ := client.Get(cfg.Endpoint, plausableDoi, badToken)
	if status != expected {
		t.Fatalf("Expected %v, got %v\n", expected, status)
	}
}

//
// DOI create tests
//

func TestCreateCrossRef(t *testing.T) {
	expected := http.StatusOK

	status, entity := client.Create(cfg.Endpoint, goodShoulder, createTestRequest(crossrefSchema), goodToken)
	if status != expected {
		t.Fatalf("Expected %v, got %v\n", expected, status)
	}

	if entity == nil {
		t.Fatalf("Expected to create entity successfully and did not\n")
	}

	if emptyField(entity.Id) {
		t.Fatalf("Expected non-empty ID field but it is empty\n")
	}
}

func TestCreateDataCite(t *testing.T) {
	expected := http.StatusOK

	status, entity := client.Create(cfg.Endpoint, goodShoulder, createTestRequest(dataciteSchema), goodToken)
	if status != expected {
		t.Fatalf("Expected %v, got %v\n", expected, status)
	}

	if entity == nil {
		t.Fatalf("Expected to create entity successfully and did not\n")
	}

	if emptyField(entity.Id) {
		t.Fatalf("Expected non-empty ID field but it is empty\n")
	}
}

func TestCreateBadSchema(t *testing.T) {
	expected := http.StatusBadRequest

	status, _ := client.Create(cfg.Endpoint, goodShoulder, createTestRequest(badSchema), goodToken)
	if status != expected {
		t.Fatalf("Expected %v, got %v\n", expected, status)
	}
}

func TestCreateEmptyToken(t *testing.T) {
	expected := http.StatusBadRequest

	status, _ := client.Create(cfg.Endpoint, goodShoulder, createTestRequest(crossrefSchema), empty)
	if status != expected {
		t.Fatalf("Expected %v, got %v\n", expected, status)
	}
}

func TestCreateBadToken(t *testing.T) {
	expected := http.StatusForbidden

	status, _ := client.Create(cfg.Endpoint, goodShoulder, createTestRequest(crossrefSchema), badToken)
	if status != expected {
		t.Fatalf("Expected %v, got %v\n", expected, status)
	}
}

//
// DOI update tests
//

func TestUpdateCrossRef(t *testing.T) {

	doi := createGoodDoi(crossrefSchema, t)
	entity := createTestRequest(crossrefSchema)
	entity.Id = doi

	expected := http.StatusOK
	status := client.Update(cfg.Endpoint, entity, goodToken)
	if status != expected {
		t.Fatalf("Expected %v, got %v\n", expected, status)
	}
}

func TestUpdateDataCite(t *testing.T) {

	doi := createGoodDoi(dataciteSchema, t)
	entity := createTestRequest(dataciteSchema)
	entity.Id = doi

	expected := http.StatusOK
	status := client.Update(cfg.Endpoint, entity, goodToken)
	if status != expected {
		t.Fatalf("Expected %v, got %v\n", expected, status)
	}
}

func TestUpdateBadSchema(t *testing.T) {

	doi := createGoodDoi(crossrefSchema, t)
	entity := createTestRequest(badSchema)
	entity.Id = doi

	expected := http.StatusBadRequest
	status := client.Update(cfg.Endpoint, entity, goodToken)
	if status != expected {
		t.Fatalf("Expected %v, got %v\n", expected, status)
	}
}

func TestUpdateEmptyId(t *testing.T) {
	entity := createTestRequest(crossrefSchema)
	entity.Id = empty
	expected := http.StatusBadRequest
	status := client.Update(cfg.Endpoint, entity, goodToken)
	if status != expected {
		t.Fatalf("Expected %v, got %v\n", expected, status)
	}
}

func TestUpdateBadId(t *testing.T) {
	entity := createTestRequest(crossrefSchema)
	entity.Id = badDoi
	expected := http.StatusBadRequest
	status := client.Update(cfg.Endpoint, entity, goodToken)
	if status != expected {
		t.Fatalf("Expected %v, got %v\n", expected, status)
	}
}

func TestUpdateEmptyToken(t *testing.T) {
	entity := createTestRequest(crossrefSchema)
	entity.Id = plausableDoi
	expected := http.StatusBadRequest
	status := client.Update(cfg.Endpoint, entity, empty)
	if status != expected {
		t.Fatalf("Expected %v, got %v\n", expected, status)
	}
}

func TestUpdateBadToken(t *testing.T) {
	entity := createTestRequest(crossrefSchema)
	entity.Id = plausableDoi
	expected := http.StatusForbidden
	status := client.Update(cfg.Endpoint, entity, badToken)
	if status != expected {
		t.Fatalf("Expected %v, got %v\n", expected, status)
	}
}

//
// DOI delete tests
//

func TestDeleteCrossRef(t *testing.T) {
	expected := http.StatusOK
	doi := createGoodDoi(crossrefSchema, t)
	status := client.Delete(cfg.Endpoint, doi, goodToken)
	if status != expected {
		t.Fatalf("Expected %v, got %v\n", expected, status)
	}
}

func TestDeleteDataCite(t *testing.T) {
	expected := http.StatusOK
	doi := createGoodDoi(dataciteSchema, t)
	status := client.Delete(cfg.Endpoint, doi, goodToken)
	if status != expected {
		t.Fatalf("Expected %v, got %v\n", expected, status)
	}
}

func TestDeleteEmptyId(t *testing.T) {
	expected := http.StatusBadRequest
	status := client.Delete(cfg.Endpoint, empty, goodToken)
	if status != expected {
		t.Fatalf("Expected %v, got %v\n", expected, status)
	}
}

func TestDeleteBadId(t *testing.T) {
	expected := http.StatusBadRequest
	status := client.Delete(cfg.Endpoint, badDoi, goodToken)
	if status != expected {
		t.Fatalf("Expected %v, got %v\n", expected, status)
	}
}

func TestDeleteEmptyToken(t *testing.T) {
	expected := http.StatusBadRequest
	status := client.Delete(cfg.Endpoint, plausableDoi, empty)
	if status != expected {
		t.Fatalf("Expected %v, got %v\n", expected, status)
	}
}

func TestDeleteBadToken(t *testing.T) {
	expected := http.StatusForbidden
	status := client.Delete(cfg.Endpoint, plausableDoi, badToken)
	if status != expected {
		t.Fatalf("Expected %v, got %v\n", expected, status)
	}
}

//
// DOI revoke tests
//

func TestRevokeCrossRef(t *testing.T) {

	expected := http.StatusOK
	doi := createGoodDoi(crossrefSchema, t)
	entity := createTestRequest(crossrefSchema)
	entity.Id = doi

	status := client.Update(cfg.Endpoint, entity, goodToken)
	if status != expected {
		t.Fatalf("Expected %v, got %v\n", expected, status)
	}

	status = client.Revoke(cfg.Endpoint, entity.Id, goodToken)
	if status != expected {
		t.Fatalf("Expected %v, got %v\n", expected, status)
	}
}

func TestRevokeDataSite(t *testing.T) {

	expected := http.StatusOK
	doi := createGoodDoi(dataciteSchema, t)
	entity := createTestRequest(dataciteSchema)
	entity.Id = doi

	status := client.Update(cfg.Endpoint, entity, goodToken)
	if status != expected {
		t.Fatalf("Expected %v, got %v\n", expected, status)
	}

	status = client.Revoke(cfg.Endpoint, entity.Id, goodToken)
	if status != expected {
		t.Fatalf("Expected %v, got %v\n", expected, status)
	}
}

func TestRevokeBadSchema(t *testing.T) {

	expected := http.StatusBadRequest
	doi := createGoodDoi(crossrefSchema, t)
	entity := createTestRequest(badSchema)
	entity.Id = doi

	status := client.Update(cfg.Endpoint, entity, goodToken)
	if status != expected {
		t.Fatalf("Expected %v, got %v\n", expected, status)
	}

	status = client.Revoke(cfg.Endpoint, entity.Id, goodToken)
	if status != expected {
		t.Fatalf("Expected %v, got %v\n", expected, status)
	}
}

func TestRevokeEmptyId(t *testing.T) {
	expected := http.StatusBadRequest
	status := client.Revoke(cfg.Endpoint, empty, goodToken)
	if status != expected {
		t.Fatalf("Expected %v, got %v\n", expected, status)
	}
}

func TestRevokeBadId(t *testing.T) {
	expected := http.StatusBadRequest
	status := client.Revoke(cfg.Endpoint, badDoi, goodToken)
	if status != expected {
		t.Fatalf("Expected %v, got %v\n", expected, status)
	}
}

func TestRevokeEmptyToken(t *testing.T) {
	expected := http.StatusBadRequest
	status := client.Revoke(cfg.Endpoint, plausableDoi, empty)
	if status != expected {
		t.Fatalf("Expected %v, got %v\n", expected, status)
	}
}

func TestRevokeBadToken(t *testing.T) {
	expected := http.StatusForbidden
	status := client.Revoke(cfg.Endpoint, plausableDoi, badToken)
	if status != expected {
		t.Fatalf("Expected %v, got %v\n", expected, status)
	}
}

//
// helpers
//

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
		Url:                "http://google.com",
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
		Url:             "http://google.com",
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
		//emptyField( schema.Url ) ||
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
		//emptyField( schema.Url ) ||
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
		if len(entity.Id) == 0 {
			t.Fatalf("Create reported success but returned a blank DOI\n")
		}
		return entity.Id
	}

	t.Fatalf("Unable to create new DOI\n")
	return ""
}

func loadConfig() TestConfig {

	data, err := ioutil.ReadFile("service_test.yml")
	if err != nil {
		log.Fatal(err)
	}

	var c TestConfig
	if err := yaml.Unmarshal(data, &c); err != nil {
		log.Fatal(err)
	}

	log.Printf("endpoint [%s]\n", c.Endpoint)
	log.Printf("token    [%s]\n", c.Token)

	return c
}
