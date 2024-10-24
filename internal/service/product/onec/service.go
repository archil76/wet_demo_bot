package onec

import (
	"errors"

	model "github.com/wet_demo_bot/internal/model/product/onec"
)

type OnecService interface {
	Describe(onecID uint64) (*model.Onec, error)
	List(cursor uint64, limit uint64) ([]model.Onec, error)
	Create(model.Onec) (uint64, error)
	Update(onecID uint64, onec model.Onec) error
	Remove(onecID uint64) (bool, error)
}

type DummyOnecService struct{}

func NewDummyOnecService() *DummyOnecService {
	return &DummyOnecService{}
}

func (s *DummyOnecService) Describe(idx uint64) (*model.Onec, error) {
	return &model.AllEntities[idx], nil
}

func (s *DummyOnecService) List(cursor uint64, limit uint64) ([]model.Onec, error) {

	entitiesLen := uint64(len(model.AllEntities) - 1)
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

func (s *DummyOnecService) Create(onec model.Onec) (uint64, error) {
	return 0, errors.New("Command Create isn't implemented")
}

func (s *DummyOnecService) Update(idx uint64, onec model.Onec) error {
	return errors.New("Command Update isn't implemented")
}

func (s *DummyOnecService) Remove(idx uint64) (bool, error) {
	model.AllEntities = append(model.AllEntities[:idx], model.AllEntities[idx+1:]...)
	return true, nil
}
