package main

import (
	"time"
	"usecase/notification"
)

func main() {
	chNotifications, stop := notification.Notifier()

	{

	}
	notifs := []notification.Notification{
		{Email: notification.Email{Body: "Only Email"}},
		{Sms: notification.SMS{Body: "Only Sms"}},
		{Email: notification.Email{Body: "Email + SMS"}, Sms: notification.SMS{Body: "Email + SMS"}},
	}

	for _, n := range notifs {
		chNotifications <- n
	}

	time.Sleep(2 * time.Second)
	stop <- true
}
