package notification

import "fmt"

type Notification struct {
	Email Email
	Sms   SMS
}

// notification can have multiple types. right now we have 2 (email sms)

// notifications received are sent on specific type channels
// if they have that corresponding type of notification
// i.e Fan Out

func Notifier() (chan<- Notification, chan<- bool) {

	chEmail := EmailSender()
	chSms := SmsSender()

	chNotifications := make(chan Notification)
	stop := make(chan bool)

	go notificationFanOut(stop, chNotifications, chEmail, chSms)

	return chNotifications, stop
}

func notificationFanOut(stop <-chan bool, in <-chan Notification, chEmail chan<- Email, chSms chan<- SMS) {

	for {

		select {
		case <-stop:
			fmt.Println("Stopping Notifier")
			close(chEmail)
			close(chSms)
			return

		case notif := <-in:
			// fmt.Println("<--Received ", notif, " -->")
			// forward to email sender if notification has email component
			if (Email{}) != notif.Email {
				chEmail <- notif.Email
			}

			// forward to sms sender if notification has sms component
			if (SMS{}) != notif.Sms {
				chSms <- notif.Sms
			}
		}

	}
}
