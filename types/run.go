package types

import "time"

// WorkerRun is the instance of job which will dispatch to worker
type WorkerRun struct {
	ID         string            `json:"id"`         // use xid for privacy protection
	JobID      string            `json:"jobID"`      // job reference
	Job        *WorkerJob        `json:"job"`        // if worker need
	JobVersion int               `json:"jobVersion"` // for worker side cache
	Params     map[string]string `json:"params" gorm:"type:json"`
	Context    map[string]string `json:"context"` // extra data for worker side hooks
}

// RunResult is the body worker reported to queen.
type RunResult struct {
	Status  RunStatus `json:"status"`
	Message string    `json:"message"`
}

// Run is the job instance in queen side
// This is a db table: runs
type Run struct {
	WorkerRun
	RunResult
	CreatedAt    time.Time  `json:"createdAt"`    // queuing
	DispatchedAt *time.Time `json:"dispatchedAt"` // queuing -> running
	SucceededAt  *time.Time `json:"succeededAt"`  // running -> succeeded
	FailedAt     *time.Time `json:"failedAt"`     // running -> failed
}

// RunStatus is job instance status in queen side
type RunStatus string

const (
	RunStatusQueuing   RunStatus = "queuing"
	RunStatusRunning   RunStatus = "running"
	RunStatusSucceeded RunStatus = "succeeded"
	RunStatusFailed    RunStatus = "failed"
)

// RunLog TODO: we'll save job log
type RunLog struct {
}
