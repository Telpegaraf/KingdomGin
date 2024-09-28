package database

import (
	"gorm.io/gorm"
	"kingdom/model"
)

func (d *GormDatabase) GetUserByUsername(name string) (*model.User, error) {
	user := new(model.User)
	err := d.DB.Where("username = ?", name).Find(user).Error
	if err == gorm.ErrRecordNotFound {
		err = nil
	}
	if user.Username == name {
		return user, err
	}
	return nil, err
}

// GetUserByID returns the user by the given id or nil.
func (d *GormDatabase) GetUserByID(id uint) (*model.User, error) {
	user := new(model.User)
	err := d.DB.Find(user, id).Error
	if err == gorm.ErrRecordNotFound {
		err = nil
	}
	if user.ID == id {
		return user, err
	}
	return nil, err
}

// CountUser returns the user count which satisfies the given condition.
func (d *GormDatabase) CountUser(condition ...interface{}) (int, error) {
	c := int64(-1)
	handle := d.DB.Model(new(model.User))
	if len(condition) == 1 {
		handle = handle.Where(condition[0])
	} else if len(condition) > 1 {
		handle = handle.Where(condition[0], condition[1:]...)
	}
	err := handle.Count(&c).Error
	return int(c), err
}

// GetUsers returns all users.
func (d *GormDatabase) GetUsers() ([]*model.User, error) {
	var users []*model.User
	err := d.DB.Find(&users).Error
	return users, err
}

// DeleteUserByID deletes a user by its id.
func (d *GormDatabase) DeleteUserByID(id uint) error {
	return d.DB.Where("id = ?", id).Delete(&model.User{}).Error
}

// UpdateUser updates a user.
func (d *GormDatabase) UpdateUser(user *model.User) error {
	return d.DB.Save(user).Error
}

// CreateUser creates a user.
func (d *GormDatabase) CreateUser(user *model.User) error {
	return d.DB.Create(user).Error
}

// GetUserByToken returns the user for the given token or nil.
func (d *GormDatabase) GetUserByToken(token string) (*model.User, error) {
	user := new(model.User)
	err := d.DB.Where("token = ?", token).Find(user).Error
	if err == gorm.ErrRecordNotFound {
		err = nil
	}
	//if user.Token == token {
	//	return user, err
	//}
	return nil, err
}
