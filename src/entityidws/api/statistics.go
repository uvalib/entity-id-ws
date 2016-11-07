package api

type Statistics struct {
    RequestCount    int   `json:"request_count"`
    CreateCount     int   `json:"create_count"`
    UpdateCount     int   `json:"update_count"`
    LookupCount     int   `json:"lookup_count"`
    DeleteCount     int   `json:"delete_count"`
    RevokeCount     int   `json:"revoke_count"`
    HeartbeatCount  int   `json:"heartbeat_count"`
}

