package models

type TodoList struct {
	Task_id   int    `json:"task_id" gorm:"primaryKey;autoIncrement"`
	Task_name string `json:"task_name"`
	Status    string `json:"status" gorm:"default:active"`
}
