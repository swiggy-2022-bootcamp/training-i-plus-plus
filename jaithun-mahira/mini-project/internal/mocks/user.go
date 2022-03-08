package mocks

import (
	"mini-project/internal/modals"
	"time"

	"github.com/google/uuid"
)

var Users = []modals.User{
    {
				Id: uuid.New(),
				Name: "Jaithun Mahira",
				Password: "dummy",
				Contact: modals.ContactInfo{
					Email: "jaithunmahira@gmail.com",
					Phone: "123456789",
				},
				DateOfBirth: time.Date(1998, time.June, 18, 0,0,0,0,time.UTC),
				IdProof: "352435634",
				Role: modals.Admin,
    },
}