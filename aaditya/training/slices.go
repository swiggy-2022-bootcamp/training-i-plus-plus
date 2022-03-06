package cards

//Concepts - Slices and multiple return values
          

// GetItem retrieves an item from a slice at given position. The second return value indicates whether
// the given index exists in the slice or not.
func GetItem(slice []int, index int) (int, bool) {
	card:= 0
   if index >= len(slice) || index < 0{
	    return card, false
    }else{
    	return slice[index],true
    }
}
// SetItem writes an item to a slice at given position overwriting an existing value.
// If the index is out of range the value needs to be appended.
func SetItem(slice []int, index, value int) []int {
	if index >= len(slice) || index < 0 {
        slice = append(slice,value)
    }else{
        slice[index] = value
    }
	return slice
}

// PrefilledSlice creates a slice of given length and prefills it with the given value.
func PrefilledSlice(value, length int) []int {
	if length < 0 {
        length = 0
    }
    slice:= make([]int, length)
   for i,_:= range slice {
        slice[i]= value
    }
	return slice
}

// RemoveItem removes an item from a slice by modifying the existing slice.
func RemoveItem(slice []int, index int) []int {
	if index >= len(slice) || index < 0 {
        return slice
    }
	return append(slice[:index], slice[index+1:]...)
}