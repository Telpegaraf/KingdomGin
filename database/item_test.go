package database

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"kingdom/model"
)

func (s *DatabaseSuite) TestItem() {
	item, err := s.db.GetItemByID(1)
	require.NoError(s.T(), err)
	assert.Nil(s.T(), item)

	armor := &model.Armor{
		ArmorClass: 2,
	}
	testArmor := &model.Item{
		Name:        "Test Armor",
		Description: "Test Description",
		Bulk:        0.003,
		Level:       5,
		Price:       "Price",
		OwnerType:   "armors",
	}
	err = s.db.CreateArmor(armor, testArmor)
	require.NoError(s.T(), err)
	assert.Equal(s.T(), testArmor.Name, "Test Armor")

	weapon := &model.Weapon{}
	testWeapon := &model.Item{
		Name:        "Test Weapon",
		Description: "Test Description",
		Bulk:        1.00,
		Level:       3,
		Price:       "Price",
		OwnerType:   "weapons",
	}
	err = s.db.CreateWeapon(weapon, testWeapon)
	require.NoError(s.T(), err)
	assert.Equal(s.T(), testWeapon.Name, "Test Weapon")

	gear := &model.Gear{}
	testGear := &model.Item{
		Name:        "Test Gear",
		Description: "Test Description",
		Bulk:        0.1,
		Level:       3,
		Price:       "Price",
		OwnerType:   "gears",
	}
	err = s.db.CreateGear(gear, testGear)
	require.NoError(s.T(), err)
	assert.Equal(s.T(), testGear.Name, "Test Gear")

	items, err := s.db.GetItems()
	require.NoError(s.T(), err)
	require.Len(s.T(), items, 3)

	armors, err := s.db.GetArmors()
	require.NoError(s.T(), err)
	require.Len(s.T(), armors, 1)

	weapons, err := s.db.GetWeapons()
	require.NoError(s.T(), err)
	require.Len(s.T(), weapons, 1)

	gears, err := s.db.GetGears()
	require.NoError(s.T(), err)
	require.Len(s.T(), gears, 1)

	armor.ArmorClass = 5
	testArmor.Bulk = 2
	err = s.db.UpdateArmor(armor, testArmor)
	require.NoError(s.T(), err)
	assert.Equal(s.T(), testArmor.Bulk, float64(2))

	weapon.DiceQuantity = 2
	testWeapon.Name = "test Weapon 2"
	err = s.db.UpdateWeapon(weapon, testWeapon)
	assert.Equal(s.T(), weapon.DiceQuantity, uint8(2))

	testGear.Name = "test Gear 2"
	err = s.db.UpdateGear(gear, testGear)
	assert.Equal(s.T(), testGear.Name, "test Gear 2")

	err = s.db.DeleteItem(1, "armors", 1)
	require.NoError(s.T(), err)
	err = s.db.DeleteItem(2, "weapons", 1)
	require.NoError(s.T(), err)
	err = s.db.DeleteItem(3, "gears", 1)
	require.NoError(s.T(), err)

	items, err = s.db.GetItems()
	require.NoError(s.T(), err)
	assert.Empty(s.T(), items)

}
