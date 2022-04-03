package err

import "log"

//CheckNilErr ..
func CheckNilErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
