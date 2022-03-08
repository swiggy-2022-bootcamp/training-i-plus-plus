package tests

import(
	"testing"
	"healthcareApp/service"
)

func TestOpenSlotsForAppointments(t *testing.T){
	service.OpenSlotsForAppointments("Rakesh","01-01-2022 15:00:00",500,false)

	slots, doesExists := service.AppointmentsMappedByDoctorName["Rakesh"]

	if !doesExists {
		t.Errorf("Slot not opened for the given doctor")
	}else{
		lastAddedAppointment := slots[len(slots)-1]

		if lastAddedAppointment.DoctorName != "Rakesh" && lastAddedAppointment.Slot != "01-01-2022 15:00:00" && lastAddedAppointment.Fees != 500 && !lastAddedAppointment.Occupied {
			t.Errorf("Slot not opened for the given doctor")
		}
	}
}

func TestOpenMultipleSlotsForAppointmentsBySameDoctor(t *testing.T){
	service.OpenSlotsForAppointments("Rakesh","01-01-2022 15:00:00",500,false)
	service.OpenSlotsForAppointments("Rakesh","01-01-2022 16:00:00",500,false)
	slots, doesExists := service.AppointmentsMappedByDoctorName["Rakesh"]

	if !doesExists {
		t.Errorf("Slot not opened for the given doctor")
	}else if len(slots) < 2 {
		t.Errorf("Failed to open mulitple slots for same doctor")
	}else {

		lastAddedAppointment := slots[len(slots)-1]

		if lastAddedAppointment.DoctorName != "Rakesh" && lastAddedAppointment.Slot != "01-01-2022 16:00:00" && lastAddedAppointment.Fees != 500 && !lastAddedAppointment.Occupied {
			t.Errorf("Slot not opened for the given doctor")
		}
	}
}