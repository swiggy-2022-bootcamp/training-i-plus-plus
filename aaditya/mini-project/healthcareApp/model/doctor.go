package model

type Doctor struct{
	Id					string		`json:"id"`
	User							`json:"user"`
	Category 			string		`json:"category"`
	Yoe 	 			float64		`json:"yoe"`
	MedicalLicenseLink	string		`json:"medicalLicenseLink"`
}