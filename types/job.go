package types

// Job will dispatch to worker
type Job struct {
	ID               string         `json:"id" gorm:"id"`                             // use xid for privacy protection
	TemplateID       string         `json:"template_id" gorm:"template_id"`           // template reference
	TemplateVersion  int            `json:"template_version" gorm:"template_version"` // for worker side cache
	URLVars          stringSlice    `json:"url_variables" gorm:"type:json"`
	BodyVars         stringSlice    `json:"body_variables" gorm:"type:json"`
	ExtraQueryParams mapStrStr      `json:"extra_query_params" gorm:"type:json"`
	ExtraHeaders     mapStrSliceStr `json:"extra_headers" gorm:"type:json"` // fixed headers
	// Context map[string]string // extra data for worker side hooks
}

// JobLog TODO: we'll save job log
type JobLog struct {
}

type JobResult struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
}
