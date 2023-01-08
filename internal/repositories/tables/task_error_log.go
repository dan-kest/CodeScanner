package tables

import "time"

type TaskErrorLog struct {
	ID        int       `gorm:"type:int8;primaryKey;autoIncrement" json:"id"`
	Body      string    `gorm:"type:text" json:"body"`
	Message   string    `gorm:"type:varchar(255)" json:"message"`
	CreatedAt time.Time `gorm:"type:timestamp(3)" json:"created_at"`
}

func (m *TaskErrorLog) TableName() string {
	return "task_error_log"
}
