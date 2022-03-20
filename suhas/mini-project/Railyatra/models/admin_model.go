package models

// type Admin struct {
//     Id       primitive.ObjectID `json:"id,omitempty"`
//     Name     string             `json:"name,omitempty" validate:"required"`
//     Location string             `json:"location,omitempty" validate:"required"`
//     Title    string             `json:"title,omitempty" validate:"required"`
// }

type Admin struct {
	Name  string `json:"name,omitempty" validate:"required"`
	Email string `json:"email,omitempty" validate:"required"`
}
