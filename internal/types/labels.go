package types

type Label struct {
	Model

	Name        string           `json:"name" gorm:"type:text;not null;"`
	Description string           `json:"description" gorm:"type:text;not null;"`
	Resources   []CommonResource `json:"resources" gorm:"many2many:label_resources;"`
}
