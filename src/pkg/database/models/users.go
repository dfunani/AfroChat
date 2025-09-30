package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	// Primary Key
	ID uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`

	// Basic Info
	Email       string  `gorm:"uniqueIndex;not null;size:255" json:"email"`
	Username    string  `gorm:"uniqueIndex;not null;size:50" json:"username"`
	DisplayName string  `gorm:"not null;size:100" json:"display_name"`
	FirstName   string  `gorm:"size:50" json:"first_name"`
	LastName    string  `gorm:"size:50" json:"last_name"`
	AvatarURL   *string `gorm:"type:text" json:"avatar_url"`
	Bio         string  `gorm:"type:text" json:"bio"`

	// Contact & Location
	PhoneNumber *string `gorm:"size:20" json:"phone_number"`
	TimeZone    string  `gorm:"default:UTC;size:50" json:"time_zone"`
	Location    *string `gorm:"size:100" json:"location"`

	// Status & Permissions
	Status      string `gorm:"default:offline;size:20" json:"status"`
	IsVerified  bool   `gorm:"default:false" json:"is_verified"`
	IsActive    bool   `gorm:"default:true" json:"is_active"`
	IsSuspended bool   `gorm:"default:false" json:"is_suspended"`
	IsBanned    bool   `gorm:"default:false" json:"is_banned"`
	IsPremium   bool   `gorm:"default:false" json:"is_premium"`

	// Security
	PasswordHash string `gorm:"not null;size:255" json:"-"`
	Salt         string `gorm:"not null;size:255" json:"-"`

	// Timestamps
	LastSeenAt     *time.Time     `json:"last_seen_at"`
	CreatedAt      time.Time      `json:"created_at"`
	UpdatedAt      time.Time      `json:"updated_at"`
	DeletedAt      gorm.DeletedAt `gorm:"index" json:"-"`
	SuspendedAt    *time.Time     `json:"suspended_at"`
	BannedAt       *time.Time     `json:"banned_at"`
	PremiumAt      *time.Time     `json:"premium_at"`
	LastLoginAt    *time.Time     `json:"last_login_at"`
	LastLogoutAt   *time.Time     `json:"last_logout_at"`
	LastActivityAt *time.Time     `json:"last_activity_at"`
}

func (User) TableName() string {
	return "users"
}
