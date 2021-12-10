package Request

type HandleGetRequest struct {
	Proxy     string            `json:"proxy"`
	Timeout   int               `json:"timeout"`
	Url       string            `json:"url"`
	UserAgent string            `json:"userAgent"`
	Ja3       string            `json:"ja3"`
	Headers   map[string]string `json:"headers"`
}
