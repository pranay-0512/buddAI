package models

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Defines the IntegratedApplication struct (integrated_applications table in the database)
type IntegratedApplication struct {
	ID         string    `gorm:"type:string;primary_key" json:"id"`
	Name       string    `gorm:"type:string;not null" json:"name"`
	AppLogoURL string    `gorm:"type:string;column:app_logo_url" json:"app_logo_url"`
	Status     string    `gorm:"type:string" json:"status"`
	CreatedAt  time.Time `gorm:"not null" json:"created_at"`
	UpdatedAt  time.Time `gorm:"not null" json:"updated_at"`

	ConnectedApps []UserConnectedApp `gorm:"foreignKey:IntegratedAppID" json:"connected_apps,omitempty"`
}

func (ia *IntegratedApplication) BeforeCreate(tx *gorm.DB) error {
	ia.ID = uuid.New().String()
	ia.CreatedAt = time.Now()
	ia.UpdatedAt = time.Now()
	return nil
}

func (ia *IntegratedApplication) BeforeSave(tx *gorm.DB) error {
	ia.UpdatedAt = time.Now()
	return nil
}

// Defines the AppAccount struct (app_accounts table in the database)
// each connected app can have multiple accounts
type AppAccount struct {
	ID                 string          `gorm:"type:string;primary_key" json:"id"`
	UserConnectedAppID string          `gorm:"type:string;not null" json:"user_connected_app_id"`
	AccountIdentifier  string          `gorm:"type:string;not null" json:"account_identifier"`
	AccountName        string          `gorm:"type:string" json:"account_name"`
	Credentials        json.RawMessage `gorm:"type:json" json:"credentials"`
	IsActive           bool            `gorm:"not null;default:true" json:"is_active"`
	CreatedAt          time.Time       `gorm:"not null" json:"created_at"`
	UpdatedAt          time.Time       `gorm:"not null" json:"updated_at"`

	UserConnectedApp UserConnectedApp `gorm:"foreignKey:UserConnectedAppID" json:"user_connected_app,omitempty"`
}

func (aa *AppAccount) BeforeCreate(tx *gorm.DB) error {
	aa.ID = uuid.New().String()
	aa.CreatedAt = time.Now()
	aa.UpdatedAt = time.Now()
	return nil
}

func (aa *AppAccount) BeforeSave(tx *gorm.DB) error {
	aa.UpdatedAt = time.Now()
	return nil
}
