package types

// Run will dispatch to worker
type Run struct {
	ID         string            `json:"id"`         // use xid for privacy protection
	JobID      string            `json:"jobID"`      // job reference
	Job        *Job              `json:"job"`        // if worker need
	JobVersion int               `json:"jobVersion"` // for worker side cache
	Params     map[string]string `json:"params" gorm:"type:json"`
	Context    map[string]string // extra data for worker side hooks
}

// RunLog TODO: we'll save job log
type RunLog struct {
}

// RunResult is the body worker reported to queen.
type RunResult struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
}
