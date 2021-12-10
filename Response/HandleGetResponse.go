package Response

type HandleGetResponse struct {
	Success bool                      `json:"success"`
	Error   string                    `json:"error"`
	Payload *HandleGetResponsePayload `json:"payload"`
}

type HandleGetResponsePayload struct {
	Text    string            `json:"text"`
	Headers map[string]string `json:"headers"`
	Status  int               `json:"status"`
}
