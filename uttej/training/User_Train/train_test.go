package train

import "testing"

func TestDetails(t *testing.T) {

	t.Run("Train Details", func(t *testing.T) {
		train1 := Train{name: "GoodTrain", seats: 500}
		got := Info(train1)

		want := "GoodTrain has a capacity of 500 seats"

		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}

	})

	t.Run("User Details", func(t *testing.T) {

		user := User{userName: "Uttej", age: 23, gender: "Male"}
		got := Info(user)

		want := "Hey Uttej, you are 23 years old and Male"

		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})

}
