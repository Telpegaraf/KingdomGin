package database

import "kingdom/model"

func (d *GormDatabase) CreateFeat(feat *model.Feat) error { return d.DB.Create(&feat).Error }
