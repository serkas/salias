package model

import (
	"gopkg.in/mgo.v2"
	"log"
)


type Storage struct  {
	Session *mgo.Session
	Col *mgo.Collection
}

func InitStorage() *Storage {
	s, err := mgo.Dial("mongodb://localhost")

	// Check if connection error, is mongo running?
	if err != nil {
		panic(err)
	}
	//defer s.Close()
	// Optional. Switch the session to a monotonic behavior.
	s.SetMode(mgo.Monotonic, true)
	c := s.DB("test").C("tasks")
	return &Storage{s, c}
}

func (s Storage) SaveTask(task *Task)  {
	err := s.Col.Insert(task)
	if err != nil {
		log.Fatal(err)
	}
}

func (s Storage) Tasks() []Task {
	tasks := make([]Task, 10)

	if err := s.Col.Find(nil).All(&tasks); err != nil {
		log.Fatal(err)
		return tasks
	}
	return tasks
}
