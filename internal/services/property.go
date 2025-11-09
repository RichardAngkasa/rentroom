package service

import (
	"errors"
	"rentroom/internal/models"
	"rentroom/utils"

	"gorm.io/gorm"
)

func GetProperty(db *gorm.DB, propertyID int) (models.Property, error) {
	var property models.Property
	err := db.First(&property, propertyID).Error
	if err != nil {
		return property, errors.New("property not found")
	}
	return property, nil
}

func GetPropertyWithImages(db *gorm.DB, propertyID int) (models.PropertyWithImages, error) {
	property, err := GetProperty(db, propertyID)
	if err != nil {
		return models.PropertyWithImages{}, err
	}

	var images []models.Image
	if err := db.Where("property_id = ?", property.ID).Find(&images).Error; err != nil {
		return models.PropertyWithImages{}, err
	}

	imageResponses := make([]models.ImageResponse, len(images))
	for i, img := range images {
		imageResponses[i] = models.ImageResponse{
			ID:         img.ID,
			PropertyID: img.PropertyID,
			Path:       img.Path,
		}
	}

	return models.PropertyWithImages{
		PropertyResponse: utils.ConvertPropertyResponse(property),
		Images:           imageResponses,
	}, nil
}

func GetPropertyIDs(db *gorm.DB, userID uint) ([]uint, error) {
	var propertyIDs []uint
	err := db.Model(&models.UserProperties{}).
		Where("user_id = ?", userID).
		Pluck("property_id", &propertyIDs).Error
	if err != nil {
		return nil, err
	}
	if len(propertyIDs) == 0 {
		return nil, errors.New("no properties found for this user")
	}
	return propertyIDs, nil
}
