package model

import "time"

type Problem struct {
	ID            int64                  `json:"problem_id"`
	Index         string                 `json:"index" gorm:"NOT NULL"`
	Title         string                 `json:"title" gorm:"NOT NULL;type:varchar(64)"`
	TimeLimit     uint                   `json:"time_limit"`
	MemoryLimit   uint                   `json:"memory_limit"`
	Author        string                 `json:"author" gorm:"NOT NULL";type:varchar(64)`
	CreatedAt     time.Time              `json:"created_at"`
	Source        string                 `json:"source" gorm:"type:varchar(32)"`
	Background    string                 `json:"background" gorm:"type:text"`
	Statement     string                 `json:"statement" gorm:"type:text"`
	Input         string                 `json:"input" gorm:"type:text"`
	Output        string                 `json:"output" gorm:"type:text"`
	ExamplesIn    string                 `json:"examples_in" gorm:"type:text"`
	ExamplesOut   string                 `json:"examples_out" gorm:"type:text"`
	Hint          string                 `json:"hint" gorm:"type:text"`
	IsOpen        bool                   `json:"is_open" gorm:"NOT NULL;default:0"`
	Tags          string                 `json:"tags" gorm:"NOT NULL;default:'[]'"`
	Testdatas     JSON                   `json:"testdatas"`
}
