package http

import (
	"encoding/json"
	"errors"
	"time"
)

type DoneDTO struct {
	IsDone bool
}

type TaskDTO struct {
	Title       string
	Description string
}

func (d TaskDTO) ValidationForCreate() error {
	if d.Title == "" {
		return errors.New("title is required")
	}

	if d.Description == "" {
		return errors.New("description is required")
	}

	return nil
}

type ErrorDTO struct {
	Message string
	Time    time.Time
}

func NewErrorDTO(err error) ErrorDTO {
	return ErrorDTO{
		Message: err.Error(),
		Time:    time.Now(),
	}
}

func (e ErrorDTO) ToString() string {
	b, err := json.MarshalIndent(e, "", "    ")

	if err != nil {
		panic(err)
	}

	return string(b)
}
