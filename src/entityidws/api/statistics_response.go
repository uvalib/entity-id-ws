package api

//
// StatisticsResponse -- response to the statistics request
//
type StatisticsResponse struct {
   Status  int        `json:"status"`
   Message string     `json:"message"`
   Details Statistics `json:"statistics"`
}

//
// end of file
//
