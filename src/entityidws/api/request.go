package api

type Request struct {
    Schema                string           `json:"schema,omitempty"`
    Id                    string           // placeholder, all requests have an ID
    CrossRef              CrossRefSchema   `json:"crossref,omitempty"`
    DataCite              DataCiteSchema   `json:"datacite,omitempty"`
}

//
// the schema used for CrossRef requests
//
type CrossRefSchema struct {
    //Id                    string   // placeholder, all requests have an ID
    Url                   string   `json:"url,omitempty"`
    Title                 string   `json:"title,omitempty"`
    Publisher             string   `json:"publisher,omitempty"`
    CreatorFirstName      string   `json:"creator_firstname,omitempty"`
    CreatorLastName       string   `json:"creator_lastname,omitempty"`
    CreatorDepartment     string   `json:"creator_department,omitempty"`
    CreatorInstitution    string   `json:"creator_institution,omitempty"`
    PublicationDate       string   `json:"publication_date,omitempty"`
    PublicationMilestone  string   `json:"publication_milestone,omitempty"`
    ResourceType          string   `json:"type,omitempty"`
}

//
// the schema used for DataCite requests
//
type DataCiteSchema struct {
    //Id                    string   // placeholder, all requests have an ID
    Url                   string   `json:"url,omitempty"`
    Title                 string   `json:"title,omitempty"`
    Abstract              string   `json:"abstract,omitempty"`
    Creators           [] Person   `json:"creators,omitempty"`
    Contributors       [] Person   `json:"contributors,omitempty"`
    Rights                string   `json:"rights,omitempty"`
    Keywords           [] string   `json:"keywords,omitempty"`
    Sponsors           [] string   `json:"sponsors,omitempty"`
    Publisher             string   `json:"publisher,omitempty"`
    PublicationDate       string   `json:"publication_date,omitempty"`
    ResourceType          string   `json:"type,omitempty"`
}

//
// the basic person details used for datacite creators and contributors
//
type Person struct {
    Index                 int      `json:"index,omitempty"`
    FirstName             string   `json:"first_name,omitempty"`
    LastName              string   `json:"last_name,omitempty"`
    Department            string   `json:"department,omitempty"`
    Institution           string   `json:"institution,omitempty"`
}