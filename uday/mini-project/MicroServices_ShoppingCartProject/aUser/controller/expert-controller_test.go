package controller
import "testing" 
func equal(actual []string,expected []string)bool{
	if len(actual)!=len(expected){
		return true
	}
	for index,_:=range actual{
		if actual[index]!=expected[index]{
			return true
		}
	}
	return false
}
func TestGetSkills(t *testing.T) {
    actual := GetSkills()
    expected := []string{"painter","plumber","carpenter"}
	if equal(actual,expected){
		t.Errorf("Expected String(%s) is not same as"+
         " actual string (%s)", expected,actual)
	}
}