package userprofiles

import "fmt"

type Doctor struct{
	emailId string
	phoneNo string
	category string
	yoe int
}

func (d Doctor) DoctorDetails(){
	fmt.Printf(" name is %s\n",d.emailId)
	fmt.Printf(" password is %s\n",d.phoneNo)
	fmt.Printf(" category is %s\n",d.category)
	fmt.Printf(" age is %d\n",d.yoe)
}