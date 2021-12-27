package db

import (
	"time"

	"gorm.io/gorm"

	"github.com/go-sdk/lib/token"
)

type Meta struct {
	ID        string         `gorm:"primaryKey;size:36"`
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	CreatedBy string         `gorm:"size:36;index"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime"`
	UpdatedBy string         `gorm:"size:36;index"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (meta *Meta) BeforeCreate(tx *gorm.DB) error {
	if meta.ID == "" {
		meta.ID = idGenerator(tx.Statement.Context)
		tx.Statement.SetColumn("id", meta.ID)
	}
	t, err := token.FromContext(tx.Statement.Context)
	if err == nil {
		meta.CreatedBy = t.GetUserId()
		tx.Statement.SetColumn("created_by", meta.CreatedBy)
	}
	return nil
}

func (meta *Meta) BeforeUpdate(tx *gorm.DB) error {
	t, err := token.FromContext(tx.Statement.Context)
	if err == nil {
		meta.UpdatedBy = t.GetUserId()
		tx.Statement.SetColumn("updated_by", meta.UpdatedBy)
	}
	return nil
}

func (meta *Meta) BeforeSave(tx *gorm.DB) error {
	t, err := token.FromContext(tx.Statement.Context)
	if err == nil {
		meta.UpdatedBy = t.GetUserId()
		tx.Statement.SetColumn("updated_by", meta.UpdatedBy)
	}
	return nil
}

func (meta *Meta) BeforeDelete(tx *gorm.DB) error {
	t, err := token.FromContext(tx.Statement.Context)
	if err == nil {
		meta.UpdatedBy = t.GetUserId()
		tx.Statement.SetColumn("updated_by", meta.UpdatedBy)
	}
	return nil
}

type MetaD struct {
	ID        string         `gorm:"primaryKey;size:36"`
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	CreatedBy string         `gorm:"size:36;index"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime"`
	UpdatedBy string         `gorm:"size:36;index"`
	DeletedAt gorm.DeletedAt `gorm:"index;uniqueIndex:uk_deleted_at"`
}

func (meta *MetaD) BeforeCreate(tx *gorm.DB) error {
	if meta.ID == "" {
		meta.ID = idGenerator(tx.Statement.Context)
		tx.Statement.SetColumn("id", meta.ID)
	}
	t, err := token.FromContext(tx.Statement.Context)
	if err == nil {
		meta.CreatedBy = t.GetUserId()
		tx.Statement.SetColumn("created_by", meta.CreatedBy)
	}
	return nil
}

func (meta *MetaD) BeforeUpdate(tx *gorm.DB) error {
	t, err := token.FromContext(tx.Statement.Context)
	if err == nil {
		meta.UpdatedBy = t.GetUserId()
		tx.Statement.SetColumn("updated_by", meta.UpdatedBy)
	}
	return nil
}

func (meta *MetaD) BeforeSave(tx *gorm.DB) error {
	t, err := token.FromContext(tx.Statement.Context)
	if err == nil {
		meta.UpdatedBy = t.GetUserId()
		tx.Statement.SetColumn("updated_by", meta.UpdatedBy)
	}
	return nil
}

func (meta *MetaD) BeforeDelete(tx *gorm.DB) error {
	t, err := token.FromContext(tx.Statement.Context)
	if err == nil {
		meta.UpdatedBy = t.GetUserId()
		tx.Statement.SetColumn("updated_by", meta.UpdatedBy)
	}
	return nil
}
