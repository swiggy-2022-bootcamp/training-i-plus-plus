package models

// Disease Data Model
type Disease struct {
	modelImpl
	Name      string   `json:"name"`      // Name of this disease
	Medicines []string `json:"medicines"` // Medicines that can cure this disease
	Symptoms  []string `json:"symptoms"`  // Symptoms of this disease
}

// Generate data for a new Disease.
func NewDisease(name string, medicines []string, symptoms []string) Disease {
	d := Disease{
		Name:      name,
		Medicines: medicines,
		Symptoms: symptoms,
	}
	d.SetId(name)
	return d
}

// Get ID of a Disease, equivalent to its name.
func (d *Disease) GetId() string {
	return d.Name
}
