package models

import (
	"time"

	"gorm.io/gorm"
)

type Todo struct {
	ID        uint64
	UserID    uint64
	Title     string
	Detail    string
	Complete  bool
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

func TodosGetAll(user *User) *[]Todo {
	var todos []Todo
	DB.Where("deleted_at is NULL and user_id = ?", user.ID).Order("updated_at desc").Find(&todos)
	return &todos
}

func TodoCreate(user *User, title string, detail string) *Todo {
	entry := Todo{Title: title, Detail: detail, UserID: user.ID, Complete: false}
	DB.Create(&entry)
	return &entry
}

func TodoFind(user *User, id uint64) *Todo {
	var todo Todo
	DB.Where("id = ? and user_id = ?", id, user.ID).First(&todo)
	return &todo
}

func (todo *Todo) MarkComplete(state bool) {
	todo.Complete = state
	DB.Save(todo)
}
