package models

type Medicine struct {
	modelImpl
	Name     string   `json:"name"`     // Name of this medicine
	Diseases []string `json:"diseases"` // Diseases that this medicine can cure
}

func NewMedicine(name string, diseases []string) *Medicine {
	d := &Medicine{
		Name:      name,
		Diseases: diseases,
	}
	d.SetId(name)
	return d
}

func (d *Medicine) GetId() string {
	return d.Name
}
