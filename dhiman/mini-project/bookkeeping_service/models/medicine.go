package models

// Medicine Data Model
type Medicine struct {
	modelImpl
	Name     string   `json:"name"`     // Name of this medicine
	Diseases []string `json:"diseases"` // Diseases that this medicine can cure
}

// Generate data for a new Medicine.
func NewMedicine(name string, diseases []string) *Medicine {
	d := &Medicine{
		Name:      name,
		Diseases: diseases,
	}
	d.SetId(name)
	return d
}

// Get ID of a Medicine, equivalent to its name.
func (d *Medicine) GetId() string {
	return d.Name
}
