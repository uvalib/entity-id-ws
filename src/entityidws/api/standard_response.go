package api

type StandardResponse struct {
   Status        int      `json:"status"`
   Message       string   `json:"message"`
   Details       *Entity  `json:"details,omitempty"`
}

