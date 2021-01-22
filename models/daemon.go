package models

import (
	"time"
)

// DaemonJob ...
type DaemonJob struct {
	ID              int64     `json:"id" gorm:"id"`
	Name            string    `json:"name" gorm:"index;not null"`
	GroupID         uint      `json:"groupID" grom:"index"`
	CommandName     string    `json:"commandName" grom:"commandName"`
	Command         []string  `json:"command" gorm:"type:varchar(1000)"`
	ErrorMailNotify bool      `json:"errorMailNotify"`
	ErrorAPINotify  bool      `json:"errorAPINotify"`
	Status          int       `json:"status"`
	MailTo          []string  `json:"mailTo" gorm:"type:varchar(1000)"`
	APITo           []string  `json:"APITo" gorm:"type:varchar(1000)"`
	FailRestart     bool      `json:"failRestart"`
	RetryNum        int       `json:"retryNum"`
	StartAt         time.Time `json:"startAt"`
	WorkUser        string    `json:"workUser"`
	WorkIP          []string  `json:"workIp" gorm:"type:varchar(1000)"`
	WorkEnv         []string  `json:"workEnv" gorm:"type:varchar(1000)"`
	WorkDir         string    `json:"workDir"`
	CreatedUserID   uint      `json:"createdUserId"`
	CreatedUsername string    `json:"createdUsername"`
	UpdatedUserID   uint      `json:"updatedUserID"`
	UpdatedUsername string    `json:"updatedUsername"`
}
