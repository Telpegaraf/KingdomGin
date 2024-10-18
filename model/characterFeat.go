package model

type CharacterFeat struct {
	ID          uint `gorm:"primary_key;AUTO_INCREMENT"`
	CharacterID uint `gorm:"foreignKey:CharacterID;uniqueIndex:idx_character_feat"`
	FeatID      uint `gorm:"foreignKey:FeatID;uniqueIndex:idx_character_feat"`
}

type CreateCharacterFeat struct {
	CharacterID uint `json:"character_id" query:"character_id" form:"character_id" binding:"required"`
	FeatID      uint `json:"feat_id" query:"feat_id" form:"feat_id" binding:"required"`
}

type UpdateCharacterFeat struct {
	CharacterID uint `json:"character_id" query:"character_id" form:"character_id"`
	FeatID      uint `json:"feat_id" query:"feat_id" form:"feat_id"`
}

type CharacterFeatExternal struct {
	ID          uint `json:"id" query:"id" form:"id"`
	CharacterID uint `json:"character_id" query:"character_id" form:"character_id"`
	FeatID      uint `json:"feat_id" query:"feat_id" form:"feat_id"`
}
