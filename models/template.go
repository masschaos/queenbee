package models

import "github.com/masschaos/queenbee/types"

type Template struct {
	types.Template
}

func (t *Template) TableName() string {
	return "template"
}

func (t *Template) Create() error {
	t.Version = 0
	return db.Create(t).Error
}

func (t *Template) Read() error {
	return db.Where("id = ?", t.ID).First(t).Error
}

func (t *Template) Update(data map[string]interface{}) error {
	t.Version += 1
	err := db.Model(t).Update(data).Error
	if err != nil {
		return err
	}

	return nil
}

func (t *Template) Delete() error {
	return db.Where("id = ?", t.ID).Delete(t).Error
}

func (t *Template) List(pageSize, page int) ([]Template, error) {
	templates := []Template{}
	err := db.Find(&templates).Offset((page - 1) * pageSize).Limit(pageSize).Error
	if err != nil {
		return nil, err
	}

	return templates, nil
}
