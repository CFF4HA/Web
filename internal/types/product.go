package types

type Product struct {
	OriginLink  string  `json:"origin_link" gorm:"type:text;not null;"`
	Description *string `json:"description,omitempty" gorm:"type:text;"`
}
