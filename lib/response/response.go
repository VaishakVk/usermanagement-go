package response

// APIResponse struct
type APIResponse struct {
	Status   bool        `json:"status"`
	Response interface{} `json:"response"`
}
