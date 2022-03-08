package main

import "fmt"


var intSlice []int
var strSlice []string



func main() {
	intSlice = []int{10,20,30,40}
	//fmt.Println("Integer Slice:",intSlice)

	strSlice = []string{"one","two","third"}
	//fmt.Println("String Slice:",strSlice)

	// 
	var arr1 []int
	arr1 = []int{1,1,2}
	arr1 = append(arr1,4)
	//fmt.Println(arr1)

	// make([]T, len, cap)
	var nums = make([]int,4,6)
	//fmt.Println("numbers :",nums,len(nums),cap(nums))
	nums = append(nums,1)
	//fmt.Println("numbers :",nums,len(nums),cap(nums))
	nums = append(nums,2)
	//fmt.Println("numbers :",nums,len(nums),cap(nums))
	nums = append(nums,3)
	//fmt.Println("numbers :",nums,len(nums),cap(nums)) // Doubles capacity from 6 to 12


	var nums1 [2]int
	//fmt.Printf("%v %v\n",len(nums1),cap(nums1))
	nums1[0] = 100
	nums1[1] = 200
	//fmt.Printf("%d %d",len(nums1),cap(nums1))

	//s := []int{10,20,40,60}

	// for key,value := range s {
	// 	fmt.Println(key,value)
	// }

	// for _,value := range s {
	// 	fmt.Println(value)
	// }

	// for i:=0; i<len(s);i++ {
	// 	fmt.Println(s[i])
	// }

	// nameAgeMap := map[string]int{
	// 	"James":30,
	// 	"Ali":40,
	// }

	// for key,value := range nameAgeMap {
	// 	fmt.Printf("%v is %d years old\n",key,value)
	// }

	// namesAndHobby := map[string][]string {
	// 	"steven": []string{"basket ball","table tennis","coding"},
	// 	"Nnamdi": []string{"sleeping","watching news","eating"},
	// }

	// namesAndHobby["Tim"] = []string{"Watching Cartoon","Dreaming","Laughing"}

	// for i, v := range namesAndHobby {
	// 	fmt.Printf("%v likes \n",i)
	// 	for j,val := range v {
	// 		fmt.Printf(" %v %v\n",j,val)
	// 	}
	// }

	// delete(namesAndHobby,"steven")

	// for i, v := range namesAndHobby {
	// 	fmt.Printf("%v likes \n",i)
	// 	for j,val := range v {
	// 		fmt.Printf(" %v %v\n",j,val)
	// 	}
	// }

	// currency := map[string]map[string]int {
	// 	"Pound" : {"GBP":1},
	// 	"Euro"  : {"EUR":2},
	// 	"Dollar": {"USD":3},
	// }

	// for key,value := range currency {
	// 	fmt.Printf("Currency Name: %v\n",key)
	// 	for k,v := range value {
	// 		fmt.Printf("\t Currency code %v \n \t Ranking : %v\n\n",k,v)
	// 	}
	// }

	// type person struct {
	// 	firstName string
	// 	lastName  string
	// 	age 	  int
	// }

	// p1 := person {
	// 	firstName : "Mark",
	// 	lastName : "Kedu",
	// 	age :30,
	// }

	//fmt.Println(p1)

	type animal struct {
			name string	
			characterstics []string
	}

	animal1 := animal{
		name : "Lion",
		characterstics : []string{"Eats human",
				"Wild Animal",
				"King of the Jungle",
		},
	}

	fmt.Println("Animal name:",animal1.name)
	for _,v := range animal1.characterstics {
		fmt.Printf("\t %v \n",v)
	}
}