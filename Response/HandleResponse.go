package Response

type HandleResponse struct {
	Success bool                   `json:"success"`
	Error   string                 `json:"error"`
	Payload *HandleResponsePayload `json:"payload"`
}

type HandleResponsePayload struct {
	Text    string            `json:"text"`
	Headers map[string]string `json:"headers"`
	Status  int               `json:"status"`
}
