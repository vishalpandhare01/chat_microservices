package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Chats struct {
	ID        string    `gorm:"type:varchar(36);primeryKey;not:null"`
	ChatName  string    `gorm:"type:varchar(100);"`
	IsGroup   bool      `gorm:"type:boolean;default:false"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
}

func (C *Chats) BeforeCreate(tx *gorm.DB) {
	C.ID = uuid.New().String()
}

type ChatParticipants struct {
	ID        string    `gorm:"type:varchar(36);primeryKey;not:null"`
	UserID    string    `gorm:"type:varchar(36);not:null"`
	ChatID    string    `gorm:"type:varchar(36);not:null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	Chats     *Chats    `gorm:"foreignKey:ChatID;constraint:OnDelete:CASCADE;"`
}

func (C *ChatParticipants) BeforeCreate(tx *gorm.DB) {
	C.ID = uuid.New().String()
}

type Messages struct {
	ID          string    `gorm:"type:varchar(36);primeryKey;not:null"`
	SenderID    string    `gorm:"type:varchar(36);not:null"` //user id
	ChatID      string    `gorm:"type:varchar(36);not:null"`
	MediaUrl    string    `gorm:"type:text"`
	Message     string    `gorm:"type:text"`
	MessageType string    `gorm:"enum('text' , 'image' ,'video' ,'document')"`
	IsRead      bool      `gorm:"type:boolean;default:false"`
	CreatedAt   time.Time `gorm:"autoCreateTime"`
	Chats       *Chats    `gorm:"foreignKey:ChatID;constraint:OnDelete:CASCADE;"`
}

func (C *Messages) BeforeCreate(tx *gorm.DB) {
	C.ID = uuid.New().String()
}

type Groups struct {
	ID        string    `gorm:"type:varchar(36);primeryKey;not:null"`
	SchoolID  string    `gorm:"type:varchar(36);not:null"`
	CreatorID string    `gorm:"type:varchar(36);not:null"`
	ClassID   string    `gorm:"type:varchar(36);"`
	GroupName string    `gorm:"type:varchar(100);"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	Chats     *Chats    `gorm:"foreignKey:ChatID;constraint:OnDelete:CASCADE;"`
}

func (C *Groups) BeforeCreate(tx *gorm.DB) {
	C.ID = uuid.New().String()
}

type GroupMember struct {
	ID        string    `gorm:"type:varchar(36);primeryKey;not:null"`
	GroupID   string    `gorm:"type:varchar(36);not:null"`
	ChatID    string    `gorm:"type:varchar(36);not:null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	Chats     *Chats    `gorm:"foreignKey:ChatID;constraint:OnDelete:CASCADE;"`
}

func (C *GroupMember) BeforeCreate(tx *gorm.DB) {
	C.ID = uuid.New().String()
}
