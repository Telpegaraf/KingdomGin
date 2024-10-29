package database

import "kingdom/model"

// CreateUserCode creates UserCode object
func (d *GormDatabase) CreateUserCode(code *model.UserCode) error {
	return d.DB.Create(code).Error
}

// GetUserCodeByEmail returns UserCode object by email
func (d *GormDatabase) GetUserCodeByEmail(email string) (*model.UserCode, error) {
	var code model.UserCode
	err := d.DB.Where("email = ?", email).Last(&code).Error
	return &code, err
}
