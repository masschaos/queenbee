package types

import (
	"net/http"
	"time"
)

// Template is a job description without request parameters.
type Template struct {
	ID       string        `json:"id" gorm:"type:varchar(20);primary_key"` // use xid for privacy protection
	Name     string        // user defined name
	OID      string        // org id
	Version  int           // start from 0, +1 after each update
	IsSimple bool          // simple job can use any http client, otherwise should use a browser.
	Request  *RequestTemp  `json:"request" gorm:"type:json"`  // store as json
	Response *ResponseTemp `json:"response" gorm:"type:json"` // store as json
	// TODO: Worker side hooks, OnStart OnFinished OnError
	// TODO: Server side hooks
	Retries   int       // tell worker retry times
	CreatedAt time.Time // for gorm
	UpdatedAt time.Time // for gorm
}

// RequestTemp is Request Template for job
type RequestTemp struct {
	URL               string       // url template, including fixed query params
	URLVariableCount  int          // for job validation
	Method            string       // GET/PUT/POST
	Headers           *http.Header // fixed headers
	Body              string       // body template
	BodyVariableCount int          // for job validation
}

// ResponseTemp is Response Template for job
type ResponseTemp struct {
	ContentType string      // html/json/xml
	Processors  []Processor // all response process rules
}

// Processor can grep data from response content
type Processor struct {
	ID       string // use xid for privacy protection
	Name     string // TODO: hide name to worker
	Selector string // css selector for html, path for json and xml
	Attr     string // if attr exists,grep the value of it,otherwise grep dom text content
}
