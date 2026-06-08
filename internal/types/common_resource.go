package types

import "github.com/google/uuid"

type CommonResource struct {
	Model
	PrimaryName  string    `json:"primary_name" gorm:"type:text;not null;"`
	ResourceType string    `json:"resource_type" gorm:"type:text;not null;check:resource_type IN ('product', 'ingredient');"`
	ResourceId   uuid.UUID `json:"resource_id" gorm:"type:uuid;not null;unique;"`

	Labels []Label `json:"labels" gorm:"many2many:label_resources;"`
}
