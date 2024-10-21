package database

import (
	"errors"
	"gorm.io/gorm"
	"kingdom/model"
)

// GetActionByID returns Action by ID
func (d *GormDatabase) GetActionByID(id uint) (*model.Action, error) {
	action := new(model.Action)
	err := d.DB.Find(action, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	if action.ID == id {
		return action, err
	}
	return nil, err
}

// CreateAction create new Action
func (d *GormDatabase) CreateAction(action *model.Action) error { return d.DB.Create(action).Error }

// GetActions returns all Actions
func (d *GormDatabase) GetActions() ([]*model.Action, error) {
	var actions []*model.Action
	err := d.DB.Find(&actions).Error
	return actions, err
}

// UpdateAction updates Action
func (d *GormDatabase) UpdateAction(action *model.Action) error { return d.DB.Save(action).Error }

// DeleteAction deletes Action
func (d *GormDatabase) DeleteAction(id uint) error {
	return d.DB.Where("id = ?", id).
		Delete(&model.Action{}).Error
}
