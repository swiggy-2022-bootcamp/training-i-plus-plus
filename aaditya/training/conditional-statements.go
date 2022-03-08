package purchase

//Concepts - Conditional if-else statements

// NeedsLicense determines whether a license is needed to drive a type of vehicle. Only "car" and "truck" require a license.
func NeedsLicense(kind string) bool {
	if kind=="car" || kind=="truck" {
        return true
    }else{
    	return false
    }
}

// ChooseVehicle recommends a vehicle for selection. It always recommends the vehicle that comes first in dictionary order.
func ChooseVehicle(option1, option2 string) string {
    var car string
	if option1 < option2 {
       car = option1
    }else{
        car = option2
    }
	return car + " is clearly the better choice."
}

// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.
func CalculateResellPrice(originalPrice, age float64) float64 {
	var sellingPrice float64 = 0.0
    if age < 3 {
        sellingPrice = 0.80 * originalPrice
    }else if age >= 10 {
        sellingPrice = 0.50 * originalPrice
    }else {
        sellingPrice = 0.70 * originalPrice
    }
	return sellingPrice
}