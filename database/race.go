package database

import (
	"kingdom/model"
)

// CreateRace creates new Race object
func (d *GormDatabase) CreateRace(race *model.Race) error { return d.DB.Create(&race).Error }

// GetRaceByID returns Race by ID
func (d *GormDatabase) GetRaceByID(id uint) (*model.Race, error) {
	race := new(model.Race)
	err := d.DB.Find(race, id).Error
	if race.ID == id {
		return race, err
	}
	return nil, err
}

// GetRaces returns all Races
func (d *GormDatabase) GetRaces() ([]*model.Race, error) {
	var races []*model.Race
	err := d.DB.Find(&races).Error
	return races, err
}

// DeleteRaceByID deletes Race by ID
func (d *GormDatabase) DeleteRaceByID(id uint) error {
	return d.DB.Where("id = ?", id).Delete(&model.Race{}).Error
}

// UpdateRace updates Race by ID
func (d *GormDatabase) UpdateRace(race *model.Race) error {
	return d.DB.Save(race).Error
}
