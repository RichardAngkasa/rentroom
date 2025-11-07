package service

import (
	"errors"
	"rentroom/internal/models"
	repository "rentroom/internal/repositories/property"

	"gorm.io/gorm"
)

type PropertyService struct {
	repo repository.PropertyRepository
}

func NewPropertyService(repo repository.PropertyRepository) *PropertyService {
	return &PropertyService{repo: repo}
}

func (s *PropertyService) ListPublicProperties(countryID uint) ([]models.PropertyResponse, error) {
	properties, err := s.repo.GetPublishedProperties(countryID)
	if err != nil {
		return nil, err
	}
	return NewPropertiesResponse(properties), err
}

func (s *PropertyService) GetPublishedByID(id uint) (*models.PropertyResponse, error) {
	property, err := s.repo.FindByID(id)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("property not found")
	}
	if err != nil {
		return nil, err
	}
	if property.Status != models.StatusPublished {
		return nil, errors.New("property is not published")
	}
	return NewPropertyResponse(property), err
}
