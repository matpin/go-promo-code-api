package app

import (
	"go-promo-code-api/domain"
	"go-promo-code-api/infrastructure/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CodeService struct {
	repo domain.Repository
}

func NewCodeService(repo domain.Repository) *CodeService {
	return &CodeService{repo: repo}
}

func (s *CodeService) GetAllCodes() ([]models.Code, error) {
	return s.repo.FindAll()
}

func (s *CodeService) InsertCode(code models.Code) error {
	return s.repo.Insert(code)
}

func (s *CodeService) UpdateCode(id primitive.ObjectID, code models.Code) error {
	return s.repo.Update(id, code)
}

func (s *CodeService) DeleteCode(id primitive.ObjectID) error {
	return s.repo.Delete(id)
}
