package train

import "time"

type Train struct {
	Train_id       int
	Source         string
	Destination    string
	Departure_time time.Time
	Arrival_time   time.Time
}

var Train_array = []Train {
	{
		Train_id:201,
		Source:"Mysore",
		Destination:"Banglore",
		Departure_time:time.Date(2022, 2, 1, 3, 0, 0, 0, time.UTC),
		Arrival_time:time.Date(2022, 2, 1, 3, 3, 0, 0, time.UTC),
	},
	{	
		Train_id: 202,
		Source: "Banglore",
		Destination: "Chennai",
		Departure_time:time.Date(2022, 2, 1, 3, 0, 0, 0, time.UTC),
		Arrival_time:time.Date(2022, 2, 1, 3, 3, 0, 0, time.UTC),
	},
	{
		Train_id:203,
		Source:"Mysore",
		Destination:"Chennai",
		Departure_time: time.Date(2022, 2, 1, 3, 0, 0, 0, time.UTC),
		Arrival_time:time.Date(2022, 2, 1, 3, 3, 0, 0, time.UTC),
	},
}

