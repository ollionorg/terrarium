package db

import (
	"github.com/google/uuid"
)

type TFResourceAttribute struct {
	Model

	ResourceTypeID uuid.UUID `gorm:"uniqueIndex:resource_attribute_unique"`
	ProviderID     uuid.UUID `gorm:"uniqueIndex:resource_attribute_unique"`
	AttributePath  string    `gorm:"uniqueIndex:resource_attribute_unique"`
	DataType       string
	Description    string
	Optional       bool
	Computed       bool

	ResourceType TFResourceType `gorm:"foreignKey:ResourceTypeID"`
	Provider     TFProvider     `gorm:"foreignKey:ProviderID"`
}

// insert a row in DB or in case of conflict in unique fields, update the existing record and set existing record ID in the given object
func (db *gDB) CreateTFResourceAttribute(e *TFResourceAttribute) (uuid.UUID, error) {
	return createOrUpdate(db.g(), e, []string{"resource_type_id", "provider_id", "attribute_path"})
}
