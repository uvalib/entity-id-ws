package main

type Response struct {
   Status        int     `json:"status"`
   Message       string  `json:"message"`
   Details       Entity  `json:"details,omitempty"`
}

