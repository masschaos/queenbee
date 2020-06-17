package types

import (
	"net/http"
	"time"

	"github.com/masschaos/x/xtype"
)

// WorkerTask defines a task which can run on workers.
type WorkerTask struct {
	ID       string       `json:"id" gorm:"type:varchar(20);primary_key"` // use xid for privacy protection
	Version  int          `json:"version"`                                // start from 0, +1 after each update
	Request  *RequestTpl  `json:"request" gorm:"type:json"`               // store as json
	Response *ResponseTpl `json:"response" gorm:"type:json"`              // store as json
	Retries  int          `json:"retries" gorm:"type:int"`                // tell worker the max retry times
}

// Task struct is the complete definition of task.
type Task struct {
	Name      string        `json:"name" gorm:"type:varchar(50)"` // user defined name
	Scope     xtype.Strings `json:"scope" gorm:"type:json"`       // required scope
	CreatedAt time.Time     `json:"created_at"`                   // for gorm
	UpdatedAt time.Time     `json:"updated_at"`                   // for gorm
}

// RequestTpl is Request Task for job
type RequestTpl struct {
	URL       string `json:"url"`                // url template, including fixed query params
	URLVarCnt int    `json:"url_variable_count"` // for job validation

	Method  string       `json:"method"`  // GET/PUT/POST
	Headers *http.Header `json:"headers"` // fixed headers

	Body       string `json:"body"`                // body template
	BodyVarCnt int    `json:"body_variable_count"` // for job validation
}

// ResponseTpl is Response Task for job
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
