package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Defines the User struct (users table in the database)
type User struct {
	ID         string    `gorm:"type:string;primary_key" json:"id"`
	UserName   string    `gorm:"type:string;uniqueIndex;not null" json:"user_name"`
	FirstName  string    `gorm:"type:string" json:"first_name"`
	LastName   string    `gorm:"type:string" json:"last_name"`
	Password   string    `gorm:"type:string;not null" json:"-"`
	AuthSource string    `gorm:"type:string" json:"auth_source"`
	CreatedAt  time.Time `gorm:"not null" json:"created_at"`
	UpdatedAt  time.Time `gorm:"not null" json:"updated_at"`

	ConnectedApps []UserConnectedApp `gorm:"foreignKey:UserID" json:"connected_apps,omitempty"`
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	u.ID = uuid.New().String()
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
	return nil
}
func (u *User) BeforeSave(tx *gorm.DB, auth_source string) error {
	u.AuthSource = auth_source
	u.UpdatedAt = time.Now()
	return nil
}

// Defines the UserConnectedApp struct (user_connected_apps table in the database)
// junction table for many to many relationship between User and IntegratedApplication
type UserConnectedApp struct {
	ID              string    `gorm:"type:string;primary_key" json:"id"`
	UserID          string    `gorm:"type:string;not null" json:"user_id"`
	IntegratedAppID string    `gorm:"type:string;not null" json:"integrated_app_id"`
	CreatedAt       time.Time `gorm:"not null" json:"created_at"`
	UpdatedAt       time.Time `gorm:"not null" json:"updated_at"`

	User          User                  `gorm:"foreignKey:UserID" json:"user,omitempty"`
	IntegratedApp IntegratedApplication `gorm:"foreignKey:IntegratedAppID" json:"integrated_app,omitempty"`
	AppAccounts   []AppAccount          `gorm:"foreignKey:UserConnectedAppID" json:"app_accounts,omitempty"`
}

func (uca *UserConnectedApp) BeforeCreate(tx *gorm.DB) error {
	uca.ID = uuid.New().String()
	uca.CreatedAt = time.Now()
	uca.UpdatedAt = time.Now()
	return nil
}

func (uca *UserConnectedApp) BeforeSave(tx *gorm.DB) error {
	uca.UpdatedAt = time.Now()
	return nil
}
