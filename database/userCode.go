package database

import "kingdom/model"

// CreateUserCode creates UserCode object
func (d *GormDatabase) CreateUserCode(code *model.UserCode) error {
	return d.DB.Create(code).Error
}

// GetUSerCodeByEmail returns UserCode object by email
func (d *GormDatabase) GetUSerCodeByEmail(email string) (*model.UserCode, error) {
	var code model.UserCode
	err := d.DB.Where("email = ?", email).First(&code).Error
	return &code, err
}
