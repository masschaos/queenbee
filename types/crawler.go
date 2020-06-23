package types

import "net/http"

// RequestTpl is Request Job for job
type RequestTpl struct {
	URL       string `json:"url"`                // url template, including fixed query params
	URLVarCnt int    `json:"url_variable_count"` // for job validation

	Method  string       `json:"method"`  // GET/PUT/POST
	Headers *http.Header `json:"headers"` // fixed headers

	Body       string `json:"body"`                // body template
	BodyVarCnt int    `json:"body_variable_count"` // for job validation
}

// ResponseTpl is Response Job for job
type ResponseTpl struct {
	ContentType string      `json:"content_type"` // html/json/xml
	Processors  []Processor `json:"processors"`   // all response process rules
}

// Processor can grep data from response content
type Processor struct {
	ID       string `json:"id"`       // use xid for privacy protection
	Name     string `json:"name"`     // TODO: hide name to worker
	Selector string `json:"selector"` // css selector for html, path for json and xml
	Attr     string `json:"attr"`     // if attr exists,grep the value of it,otherwise grep dom text content
}
