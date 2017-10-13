package api

import (
   "sort"
)

//
// Request -- the request structure
//
type Request struct {
   Schema   string         `json:"schema,omitempty"`
   ID       string         `json:"id,omitempty"`
   CrossRef CrossRefSchema `json:"crossref,omitempty"`
   DataCite DataCiteSchema `json:"datacite,omitempty"`
}

//
// CrossRefSchema -- the schema used for CrossRef requests
//
type CrossRefSchema struct {
   URL                  string `json:"url,omitempty"`
   Title                string `json:"title,omitempty"`
   Publisher            string `json:"publisher,omitempty"`
   CreatorFirstName     string `json:"creator_firstname,omitempty"`
   CreatorLastName      string `json:"creator_lastname,omitempty"`
   CreatorDepartment    string `json:"creator_department,omitempty"`
   CreatorInstitution   string `json:"creator_institution,omitempty"`
   PublicationDate      string `json:"publication_date,omitempty"`
   PublicationMilestone string `json:"publication_milestone,omitempty"`
   ResourceType         string `json:"type,omitempty"`
}

//
// DataCiteSchema -- the schema used for DataCite requests
//
type DataCiteSchema struct {
   URL             string   `json:"url,omitempty"`
   Title           string   `json:"title,omitempty"`
   Abstract        string   `json:"abstract,omitempty"`
   Creators        []Person `json:"creators,omitempty"`
   Contributors    []Person `json:"contributors,omitempty"`
   Rights          string   `json:"rights,omitempty"`
   Keywords        []string `json:"keywords,omitempty"`
   Sponsors        []string `json:"sponsors,omitempty"`
   Publisher       string   `json:"publisher,omitempty"`
   PublicationDate string   `json:"publication_date,omitempty"`
   GeneralType     string   `json:"general_type,omitempty"`
   ResourceType    string   `json:"resource_type,omitempty"`
}

//
// Person -- the basic person details used for datacite creators and contributors
//
type Person struct {
   Index       int    `json:"index"`
   FirstName   string `json:"first_name,omitempty"`
   LastName    string `json:"last_name,omitempty"`
   Department  string `json:"department,omitempty"`
   Institution string `json:"institution,omitempty"`
}

//
// helpers to sort the people lists
//

// SortPeople -- sort a person list
func SortPeople(people []Person) []Person {
   sortedPeople := make([]Person, len(people))
   copy(sortedPeople, people)
   sort.Sort(PeopleSorter(sortedPeople))
   return sortedPeople
}

// PeopleSorter -- sorts people by index
type PeopleSorter []Person

func (people PeopleSorter) Len() int           { return len(people) }
func (people PeopleSorter) Swap(i, j int)      { people[i], people[j] = people[j], people[i] }
func (people PeopleSorter) Less(i, j int) bool { return people[i].Index < people[j].Index }

//
// end of file
//
