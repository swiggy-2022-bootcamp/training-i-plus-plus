package main

type generalQuota struct {
	source      string
	destination string
	train       string
	fare        int
}

type tatkalQuota struct {
	source      string
	destination string
	train       string
	fare        int
}

type ladiesQuota struct {
	source      string
	destination string
	train       string
	fare        int
}

type seniorCitizenQuota struct {
	source      string
	destination string
	train       string
	fare        int
}

func (g generalQuota) calculateFare() int {
	return 200

}

func (l ladiesQuota) calculateFare() int {
	return 150

}

func (t tatkalQuota) calculateFare() int {
	return 350

}

func (s seniorCitizenQuota) calculateFare() int {
	return 120

}

func main() {

}
