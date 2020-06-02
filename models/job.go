package models

import "github.com/masschaos/queenbee/types"

type Job struct {
	types.Job
}

func (j *Job) TableName() string {
	return "job"
}

func (j *Job) Create() error {
	return db.Create(j).Error
}

func (j *Job) Read() error {
	return db.First(j).Error
}

func (j *Job) Delete() error {
	return db.Where("id = ?", j.ID).Delete(j).Error
}
