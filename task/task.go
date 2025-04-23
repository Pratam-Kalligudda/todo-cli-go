package task

import "time"

// import(
// 	"fmt"
//)

type Task struct {
	ID          int
	Description string
	Status      string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
