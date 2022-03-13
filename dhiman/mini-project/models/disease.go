package models

type Disease struct {
	modelImpl
	Name      string   `json:"name"`      // Name of this disease
	Medicines []string `json:"medicines"` // Medicines that can cure this disease
}

func NewDisease(name string, medicines []string) *Disease {
	d := &Disease{
		Name:      name,
		Medicines: medicines,
	}
	d.SetId(name)
	return d
}

func (d *Disease) GetId() string {
	return d.Name
}
