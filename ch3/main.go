package main

import "fmt"

func main() {
	// Go has arrays...but they are too rigid. This ain't C!
	// declaring an empty array of zeros
	var x [3]int
	// or...to specify an array literal
	var y = [3]int{10, 20, 30}
	var yy = [...]int{1, 2, 3, 4, 5, 6}
	// or a sparse array, where you can set most elements to zero, and some based on their indices
	var z = [12]int{1, 5: 4, 6, 10: 100, 15}
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(yy)
	fmt.Println(z)

	// arrays have the strange limitation where Go considers the size

	//slices
	var s = []int{1, 5: 4, 6, 10: 100, 15}
	fmt.Println(s)

	fmt.Println("here are some maps")
	// empty map[type]type
	var nilMap map[string]int
	fmt.Println(nilMap)

	totalWins := map[string]int{}
	fmt.Println(totalWins)

	teams := map[string][]string{
		"Orcas":   {"Fred", "Ralph", "Bijou"},
		"Lions":   {"Sarah", "Peter", "Simba"},
		"Kittens": {"Sunset", "Cheetah", "Smokey"},
	}
	fmt.Println(teams)
 // let's read and write some maps
 totalWins["Orcas"] = 1 
 totalWins["Lions"] = 2
 fmt.Println(totalWins["Orcas"])
 fmt.Println(totalWins["Kittens"])

 totalWins["Kittens"]++
 fmt.Println(totalWins["Kittens"])
 totalWins["Lions"] = 3
 fmt.Println(totalWins["Lions"])

 m := map[string]int {
   "hello": 5,
   "world": 0,
 }

 v, ok := m["hello"]
 fmt.Println(v, ok)
 v, ok = m["world"]
 fmt.Println(v, ok)
 v, ok = m["goodbye"]
 fmt.Println(v, ok)

 //map as a set
 intSet := map[int]bool{}
 vals := []int{5, 10, 2, 5, 8, 7, 3, 9, 1, 2, 10}
 for _, v := range vals {
   intSet[v] = true
 }

 fmt.Println(len(vals), len(intSet))
 fmt.Println(intSet[5])
  fmt.Println(intSet[500])
  if intSet[100] {
    fmt.Println("100 is in the set")
  }
  vals = append(vals, 20)
  if intSet[20] {
    fmt.Println("100 is in the set")
  }

}
