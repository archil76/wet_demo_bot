package task

import (
	"errors"
)

type TaskService struct {
	data *taskModel
}

func NewTaskService() *TaskService {

	taskEntities.Init()

	return &TaskService{taskEntities}
}

func (s *TaskService) List(cursor, limit uint64) (result taskModel) {

	lenSlice := uint64(len(*s.data))

	if cursor > lenSlice {
		cursor = 0
	}

	data := *s.data

	if cursor+limit > lenSlice {
		result = data[cursor:]
	} else {
		result = data[cursor : cursor+limit]
	}

	return result
}

func (s *TaskService) Create(Task_id uint64, name, desc string) error {

	if _, err := s.Find(Task_id); err == nil {
		return errors.New("product id found on another product")
	}

	data := *s.data

	data = append(data, task{Id: Task_id, Championat_id: 1, Difficulty: 5, Title: name, Description: desc})

	s.data = &data

	return nil
}

func (s *TaskService) Update(Task_id uint64, Title, Desc string) error {

	id, err := s.Find(Task_id)
	if err != nil {
		return err
	}

	data := *s.data
	data[id-1].Title = Title
	data[id-1].Description = Desc

	return nil
}

func (s *TaskService) Remove(Task_id uint64) error {

	id, err := s.Find(Task_id)
	if err != nil {
		return err
	}

	data := *s.data
	data = append(data[:id], data[id+1:]...)

	s.data = &data

	return nil
}

func (s *TaskService) Count() int {
	return len(*s.data)
}

func (s *TaskService) Get(Task_id uint64) task {

	id, err := s.Find(Task_id)
	if err != nil {
		return task{}
	}

	result := *s.data

	return result[id]

}

func (s *TaskService) Find(Task_id uint64) (int, error) {

	data := *s.data

	for i, v := range data {
		if v.Id == Task_id {
			return i, nil
		}
	}

	return 0, errors.New("id not found")
}
