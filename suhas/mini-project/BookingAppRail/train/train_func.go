package train

func GetTrainArray () []Train {
	return Train_array
}

func AddTrain (train Train) bool{
	Train_array = append(Train_array,train)
	return true
}