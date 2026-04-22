package db

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"time"
)

type SeverityLevel string

const (
	SeverityInfo     SeverityLevel = "info"
	SeverityWarning  SeverityLevel = "warning"
	SeverityCritical SeverityLevel = "critical"
)

func (s SeverityLevel) IsValid() bool {
	switch s {
	case SeverityInfo, SeverityWarning, SeverityCritical:
		return true
	}
	return false
}

type KnowledgeBase struct {
	gorm.Model
	Title           string        `gorm:"not null"`
	Content         string        `gorm:"type:text"`
	Recommendations string        `gorm:"type:text"`
	Severity        SeverityLevel `gorm:"type:varchar(20);not null;default:info"`
}

func (k *KnowledgeBase) BeforeCreate(tx *gorm.DB) error {
	if !k.Severity.IsValid() {
		return fmt.Errorf("invalid severity level: %s", k.Severity)
	}
	return nil
}

type User struct {
	gorm.Model
	Age                  int        `gorm:"not null"`
	LastMenstruationDate *time.Time `gorm:"type:date"`
	PregnancyWeeks       *int
	DisclaimerAccepted   bool `gorm:"not null;default:false"`
	DisclaimerAcceptedAt *time.Time
}

type Session struct {
	gorm.Model
	SessionID string `gorm:"size:64;uniqueIndex;not null"`
	UserID    *uint
	User      *User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

func NewSQLiteDB() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("./data/app.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	if err := db.AutoMigrate(&KnowledgeBase{}, &User{}, &Session{}); err != nil {
		return nil, err
	}

	return db, nil
}
