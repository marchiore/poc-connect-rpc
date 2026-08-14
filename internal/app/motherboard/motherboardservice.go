package motherboard

import (
	repo "poc/internal/repository/motherboard"
)

type MotherboardService struct {
	repository *repo.MotherboardRepository
}

func NewMotherboardService(repository *repo.MotherboardRepository) *MotherboardService {
	return &MotherboardService{
		repository: repository,
	}
}

func (s *MotherboardService) GetMotherboardByName(name string) (*repo.Motherboard, error) {
	return s.repository.FindByName(name)
}

func (s *MotherboardService) CreateMotherboard(description string, colorID, manufacturerID int32) (string, error) {
	return s.repository.Insert(description, colorID, manufacturerID)
}
