package main

type Entity struct {
    Id             string   `json:"id,omitempty"`
    Url            string   `json:"url,omitempty"`
    Title          string   `json:"title,omitempty"`
    Publisher      string   `json:"publisher,omitempty"`
    Creator        string   `json:"creator,omitempty"`
    PubYear        string   `json:"publication_year,omitempty"`
    ResourceType   string   `json:"type,omitempty"`
}