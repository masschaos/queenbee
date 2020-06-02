package types

import "net/http"

// Job will dispatch to worker
type Job struct {
	ID               string            `json:"id"`                                      // use xid for privacy protection
	TemplateID       string            `json:"template_id" gorm:"template_id"`          // template reference
	TemplateVersion  int               `json:"template_version gorm:"template_version"` // for worker side cache
	URLVars          []string          `json:"url_variables"`
	BodyVars         []string          `json:"body_variables"`
	ExtraQueryParams map[string]string `json:"extra_query_params" gorm:"extra_query_params"`
	ExtraHeaders     *http.Header      `json:"extra_headers"` // fixed headers
	// Context map[string]string // extra data for worker side hooks
}

// JobLog TODO: we'll save job log
type JobLog struct {
}
