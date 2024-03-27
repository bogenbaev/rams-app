package models

import "time"

type RealProperty struct {
	ID             int       `json:"id,omitempty" db:"id"`
	PropertyTypeID int       `json:"property_type_id,omitempty" db:"property_type_id" validate:"required"`
	PropertyType   string    `json:"property_type,omitempty" db:"property_type" validate:"required"`
	Address        string    `json:"address,omitempty" db:"address" validate:"required"`
	Price          float32   `json:"price,omitempty" db:"price" validate:"required"`
	Rooms          int       `json:"rooms,omitempty" db:"rooms" validate:"required,gte=1"`
	Area           float32   `json:"area,omitempty" db:"area" validate:"required,gte=0"`
	Description    string    `json:"description,omitempty" db:"description"`
	CreatedAt      time.Time `json:"-" db:"created_at"`
}
