# SpoofingTlsFingerprint
Прокси сервер для обхода TLS Fingerprint

## Зависимости
```
golang ^v1.17
```

## Запуск
```
go run main.go
```

## Сборка в exe
```
go build -o main.exe main.go
```

## Основные структры SpoofingTlsFingerprint
<a href="https://github.com/Skyuzii/SpoofingTlsFingerprint/blob/main/Request/HandleRequest.go">HandleRequest</a>
```GO
Cookies   []cycletls.Cookie `json:"cookies"`
Method    string            `json:"method"`
Body      string            `json:"body"`
Proxy     string            `json:"proxy"`
Timeout   int               `json:"timeout"`
Url       string            `json:"url"`
UserAgent string            `json:"userAgent"`
Ja3       string            `json:"ja3"`
Headers   map[string]string `json:"headers"`
```

<a href="https://github.com/Skyuzii/SpoofingTlsFingerprint/blob/main/Response/HandleResponse.go">HandleResponse</a>
```GO
Success bool                   `json:"success"`
Error   string                 `json:"error"`
Payload *HandleResponsePayload `json:"payload"`
```

<a href="https://github.com/Skyuzii/SpoofingTlsFingerprint/blob/main/Response/HandleResponse.go">HandleResponsePayload</a>
```GO
Text    string             `json:"text"`
Headers map[string]string  `json:"headers"`
Status  int                `json:"status"`
Url     string             `json:"url"`
Cookies []*cycletls.Cookie `json:"cookies"`
```
