package main

type Statistics struct {
    RequestCount  int   `json:"request_count"`
    CreateCount   int   `json:"create_count"`
    UpdateCount   int   `json:"update_count"`
    LookupCount   int   `json:"lookup_count"`
    DeleteCount   int   `json:"delete_count"`
}

type StatisticsResponse struct {
    Status        int        `json:"status"`
    Message       string     `json:"message"`
    Details       Statistics `json:"statistics"`
}


