package main

//import "time"

type Entity struct {
    Id             string   `json:"id"`
    Url            string   `json:"url"`
    Title          string   `json:"title"`
    Publisher      string   `json:"publisher"`
    Creator        string   `json:"creator"`
    PubYear        string   `json:"publication_year"`
    ResourceType   string   `json:"type"`
}

