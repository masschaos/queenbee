package types

import (
	"net/http"
	"time"
)

// Template is a job description without request parameters.
type Template struct {
	ID       string       `json:"id" gorm:"type:varchar(20);primary_key"` // use xid for privacy protection
	Name     string       `json:"name" gorm:"type:varchar(255)`           // user defined name
	OrgID    string       `json:"org_id" gorm:"type:varchar(255)`         // org id
	Version  int          `json:"version"`                                // start from 0, +1 after each update
	IsSimple bool         `json:"is_simple"`                              // simple job can use any http client, otherwise should use a browser.
	Request  *RequestTpl  `json:"request" gorm:"type:json"`               // store as json
	Response *ResponseTpl `json:"response" gorm:"type:json"`              // store as json
	// TODO: Worker side hooks, OnStart OnFinished OnError
	// TODO: Server side hooks
	Retries   int       `json:"retries" gorm:"type:int"` // tell worker retry times
	CreatedAt time.Time // for gorm
	UpdatedAt time.Time // for gorm
}

// RequestTpl is Request Template for job
type RequestTpl struct {
	URL       string `json:"url"`                // url template, including fixed query params
	URLVarCnt int    `json:"url_variable_count"` // for job validation

	Method  string       `json:"method"`  // GET/PUT/POST
	Headers *http.Header `json:"headers"` // fixed headers

	Body       string `json:"body"`                // body template
	BodyVarCnt int    `json:"body_variable_count"` // for job validation
}

// ResponseTpl is Response Template for job
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
