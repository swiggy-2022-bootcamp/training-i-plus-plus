package models

type Disease struct {
	modelImpl
	Name      string   `json:"name"`      // Name of this disease
	Medicines []string `json:"medicines"` // Medicines that can cure this disease
	Symptoms  []string `json:"symptoms"`  // Symptoms of this disease
}

func NewDisease(name string, medicines []string, symptoms []string) *Disease {
	d := &Disease{
		Name:      name,
		Medicines: medicines,
		Symptoms: symptoms,
	}
	d.SetId(name)
	return d
}

func (d *Disease) GetId() string {
	return d.Name
}
