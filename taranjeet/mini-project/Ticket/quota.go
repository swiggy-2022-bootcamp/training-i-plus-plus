package trs

import "fmt"

type generalQuota struct {
	name string
}

type tatkalQuota struct {
	name string
}

type ladiesQuota struct {
	name string
}

type seniorCitizenQuota struct {
	name string
}

type Quota interface {
	CalculateFare() int
}

func (g generalQuota) CalculateFare() int {
	return 200

}

func (l ladiesQuota) CalculateFare() int {
	return 150

}

func (t tatkalQuota) CalculateFare() int {
	return 350

}

func (s seniorCitizenQuota) CalculateFare() int {
	return 120

}

// Factory Method to create a instance of a passed quota type

func CreateQuotaFactory(typeOfQuota string) (Quota, error) {
	var quota Quota
	switch typeOfQuota {
	case "General Quota":
		quota = generalQuota{name: "General Quota"}
		return quota, nil
	case "Ladies Quota":
		quota = ladiesQuota{name: "Ladies Quota"}
		return quota, nil
	case "Tatkal Quota":
		quota = tatkalQuota{name: "Tatkal Quota"}
		return quota, nil
	case "Senior Citizen Quota":
		quota = seniorCitizenQuota{name: "Senior Citizen Quota"}
		return quota, nil
	}

	return nil, fmt.Errorf("Wrong type of quota passed")

}
