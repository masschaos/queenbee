package types

import "net/http"

// Job will dispatch to worker
type Job struct {
	ID               string   `json:"id"`              // use xid for privacy protection
	TemplateID       string   `json:"templateID"`      // template reference
	TemplateVersion  int      `json:"templateVersion"` // for worker side cache
	URLVariables     []string `json:"urlVariables"`
	BodyVariables    []string `json:"bodyVariables"`
	ExtraQueryParams map[string]string
	ExtraHeaders     *http.Header // fixed headers
	// Context map[string]string // extra data for worker side hooks
}

// JobLog TODO: we'll save job log
type JobLog struct {
}
