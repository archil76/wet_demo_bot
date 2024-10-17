package onec

import (
	model "github.com/wet_demo_bot/internal/model/product/onec"
)

type OnecService interface {
	Describe(onecID uint64) (*model.Onec, error)
	List(cursor uint64, limit uint64) ([]model.Onec, error)
	Create(model.Onec) (uint64, error)
	Update(onecID uint64, onec model.Onec) error
	Remove(onecID uint64) (bool, error)
}

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func (s *Service) List(cursor int, limit int) ([]model.Onec, error) {

	entitiesLen := len(model.AllEntities) - 1
	if cursor > entitiesLen {
		return []model.Onec{}, nil
	}
	if cursor == entitiesLen {
		return model.AllEntities[cursor:], nil
	}

	lastIndex := cursor + limit
	if lastIndex >= entitiesLen {
		return model.AllEntities[cursor:], nil
	}

	return model.AllEntities[cursor:lastIndex], nil
}

func (s *Service) Get(idx int) (*model.Onec, error) {
	return &model.AllEntities[idx], nil
}

func (s *Service) Remove(idx int) (bool, error) {
	model.AllEntities = append(model.AllEntities[:idx], model.AllEntities[idx+1:]...)
	return true, nil
}

type DummyOnecService struct{}

func NewDummyOnecService() *DummyOnecService {
	return &DummyOnecService{}
}

func (s *DummyOnecService) List() []model.Onec {
	return model.AllEntities
}

func (s *DummyOnecService) Get(idx int) (*model.Onec, error) {
	return &model.AllEntities[idx], nil
}

func (s *DummyOnecService) Describe(idx int) (*model.Onec, error) {
	return s.Get(idx)
}

func (s *DummyOnecService) Create(onec model.Onec) (uint64, error) {
	return 0, nil
}

func (s *DummyOnecService) Update(idx int, onec model.Onec) error {
	return nil
}

func (s *DummyOnecService) Remove(idx int) (bool, error) {
	return true, nil
}
