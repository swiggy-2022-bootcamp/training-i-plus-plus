package thefarm
import (
	"errors"
    "strconv"
)

        
//Concept - Error type in golang

type WeightFodder interface {
	FodderAmount() (float64, error)
}

var ErrScaleMalfunction = errors.New("sensor error")

type SillyNephewError struct {
    numberOfCows int
}

func (e *SillyNephewError) Error() string {
    return "silly nephew, there cannot be " + strconv.Itoa(e.numberOfCows) + " cows";
}

// DivideFood computes the fodder amount per cow for the given cows.
func DivideFood(weightFodder WeightFodder, cows int) (float64, error) {
	fodder,err := weightFodder.FodderAmount()
    //Task1
    if(err != nil){
      if err == ErrScaleMalfunction{
            if fodder > 0  {
                ans := float64(2)*fodder/float64(cows)
                return ans,nil
        	}else{
        		return 0, errors.New("negative fodder")
        	}
        }else{
        	return 0, err
        }
    }
    if fodder < 0{
    	return 0,errors.New("negative fodder")        

    }
    if cows == 0 {
        return 0,errors.New("division by zero")   

    }
    if cows < 0 {
    	return 0, &SillyNephewError{numberOfCows:cows}
    }

    return fodder/float64(cows), nil
}