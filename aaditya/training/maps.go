package Gross

//Concepts - Maps data structure in golang

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	units := map[string]int{}
    units["quarter_of_a_dozen"] = 3
    units["half_of_a_dozen"] = 6
    units["dozen"] = 12
    units["small_gross"]=120
    units["gross"] = 144
    units["great_gross"]=1728
    return units
}


// NewBill creates a new bill.
func NewBill() map[string]int {
	bill := make(map[string]int)
    return bill
}


// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	_,isPresent := units[unit]
    if !isPresent {
        return false
    }
	quantity,isFound := bill[item]
    if(isFound){
        bill[item] = quantity + units[unit]
    }else{
  		bill[item] = units[unit]
    }
	return true
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	quantity,isItemPresent := bill[item]
    _,isUnitPresent := units[unit]
    if !isItemPresent || !isUnitPresent {
        return false
    }
	quantity = quantity - units[unit]
    if quantity < 0 {
        return false
    }else if quantity == 0{
  	delete(bill, item)
        return true
    }else{
    	bill[item] = quantity
		return true
    }
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {

	quantity,isPresent := bill[item]
    return quantity,isPresent
}