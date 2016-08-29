package api

type Entity struct {
    Id                    string   `json:"id,omitempty"`
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