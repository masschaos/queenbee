package types

import (
	"encoding/json"
	"time"

	"github.com/masschaos/x/xtype"
)

// WorkerJob defines a job which can run on workers.
type WorkerJob struct {
	ID        string `json:"id" gorm:"type:varchar(20);primary_key"` // use xid for privacy protection
	Version   int    `json:"version"`                                // start from 0, +1 after each update
	Steps     Steps  `json:"steps" gorm:"type:json"`                 // job steps
	TimeLimit int    `json:"timeLimit"`                              // tell worker the max run seconds
}

// Job struct is the complete definition of job.
// This is a db table: jobs
type Job struct {
	WorkerJob
	Name         string    `json:"name" gorm:"type:varchar(50)"`  // user defined name
	Scope        xtype.Set `json:"scope" gorm:"type:json"`        // required scope, union of the step scopes
	RetryLimit   int       `json:"retryLimit"`                    // tell queen the max retry time for this job run
	NetworkLimit *Network  `json:"networkLimit" gorm:"type:json"` // network limit in run will overwrite it in job define
	QueenSteps   Steps     `json:"queenSteps" gorm:"type:json"`   // job steps runs on queen
	CreatedAt    time.Time `json:"createdAt"`                     // for gorm
	UpdatedAt    time.Time `json:"updatedAt"`                     // for gorm
}

// Steps is a step array, defined for custom db driver
type Steps []Step

// Step will be process one by one in a job run.
type Step struct {
	ID         string   `json:"id" gorm:"type:varchar(20)"`       // user defined id
	Name       string   `json:"name" gorm:"type:varchar(50)"`     // user defined name
	If         string   `json:"if"`                               // You can use the if conditional to prevent a step from running unless a condition is met.
	RetryLimit int      `json:"retryLimit"`                       // tell worker the max retry times for this step
	TimeLimit  int      `json:"timeLimit"`                        // tell worker the max run seconds
	Type       StepType `json:"stepType" gorm:"type:varchar(20)"` // determines content struct
	Content    json.RawMessage
}

// StepType determines the definition of the step
type StepType string

// All step types
const (
	StepTypeDocker StepType = "docker"
)
