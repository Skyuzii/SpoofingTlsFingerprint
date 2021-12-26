package Request

import "github.com/Skyuzii/CycleTLS/cycletls"

type HandleRequest struct {
	Cookies            []cycletls.Cookie `json:"cookies"`
	Method             string            `json:"method"`
	Body               string            `json:"body"`
	Proxy              string            `json:"proxy"`
	Timeout            int               `json:"timeout"`
	Url                string            `json:"url"`
	UserAgent          string            `json:"userAgent"`
	Ja3                string            `json:"ja3"`
	Headers            map[string]string `json:"headers"`
	InsecureSkipVerify bool              `json:"insecureSkipVerify"`
}
