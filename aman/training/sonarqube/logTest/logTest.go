package logTest

import (
	"io/ioutil"
	"strings"
	"testing"
)

func Test(t *testing.T) bool {
	Log.Println("A line with key word in it")
	body, err := ioutil.ReadFile("logTest/test.log")
	var result bool
	if err != nil {
		t.Errorf("Couldn't log")
		result = false
	} else {
		result = strings.Contains(string(body), "key")
	}
	return result
}
