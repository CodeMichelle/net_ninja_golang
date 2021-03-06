package main

import "fmt"

// var score = 99.5

// func sayGreeting(n string) {
// 	fmt.Printf("Good morning %v \n", n)
// }

// func sayFarewell(n string) {
// 	fmt.Printf("Goodbye %v \n", n)
// }

// func cycleNames(n []string, f func(string)) {
// 	for _, v := range n {
// 		f(v)
// 	}
// }

// func circleArea(r float64) float64 {
// 	return math.Pi * r * r
// }

// func getInitials(n string) (string, string) {
// 	s := strings.ToUpper(n)
// 	names := strings.Split(s, " ")

// 	var initials []string
// 	for _, v := range names {
// 		initials = append(initials, v[:1])
// 	}

// 	if len(initials) > 1 {
// 		return initials[0], initials[1]
// 	}

// 	return initials[0], "_"
// }

func updateName(x *string) {
	*x = "Taylor"
}

// func updateMenu(y map[string]float64) {
// 	y["coffee"] = 2.99
// }

func main() {
	// POINTERS
	name := "Michelle"
	// updateName(name)
	fmt.Println("Memory address of name is: ", &name)
	m := &name
	fmt.Println("Memory address: ", m)
	fmt.Println(name)
	//An asterik with a memory address variable accesses the value of the pointer
	fmt.Println("Value at memory address:", *m)
	updateName(m)
	fmt.Println("Value after updateName:", name)

	// PASS BY VALUE
	// Go makes copies of values when passed into functions because it is a pass-by-value language
	// name := "Michelle"
	// name = updateName(name)
	// fmt.Println(name)

	// menu := map[string]float64{
	// 	"pie":       5.95,
	// 	"ice cream": 3.99,
	// }

	// updateMenu(menu)
	// fmt.Println(menu)
	// fmt.Println(menu["ice cream"])

	// MAPS
	// menu := map[string]float64{
	// 	"soup":    4.99,
	// 	"pie":     7.99,
	// 	"salad":   6.99,
	// 	"pudding": 3.55,
	// }

	// fmt.Println(menu)
	// fmt.Println(menu["soup"])

	// LOOPING MAPS
	// for k, v := range menu {
	// 	fmt.Println(k, "-", v)
	// }

	// INTS AS KEY TYPE
	// phonebook := map[int]string{
	// 	12345567: "Michelle",
	// 	23455789: "Derrick",
	// }

	// fmt.Println(phonebook)
	// fmt.Println(phonebook[23455789])

	// phonebook[23455789] = "Michelle"
	// fmt.Println(phonebook)

	// phonebook[12345567] = "Derrick"
	// fmt.Println(phonebook)

	// PACKAGE SCOPE
	// sayHello("Michelle")

	// for _, v := range points {
	// 	fmt.Println(v)
	// }

	// showScore()

	// MULTIPLE RETURN VALUES
	// firstNameOne, secondNameOne := getInitials("Michelle Taylor")
	// fmt.Println(firstNameOne, secondNameOne)

	// firstNameTwo, secondNameTwo := getInitials("Ilana VanWyk")
	// fmt.Println(firstNameTwo, secondNameTwo)

	// FUNCTIONS
	// sayGreeting("Michelle")
	// sayFarewell("Michelle")
	// cycleNames([]string{"Michelle", "Derrick", "Ilana"}, sayGreeting)
	// cycleNames([]string{"Michelle", "Derrick", "Ilana"}, sayFarewell)
	// a1 := circleArea(10.5)
	// a2 := circleArea(15)
	// fmt.Println(a1, a2)
	// fmt.Printf("Circle 1 is %0.3f and circle 2 is %0.3f", a1, a2)

	// BOOLEANS AND CONDITIONALS
	// age := 45
	// fmt.Println("One", age <= 50)
	// fmt.Println("Two", age >= 50)
	// fmt.Println("Three", age == 45)
	// fmt.Println("Four", age != 50)

	// if age < 30 {
	// 	fmt.Println("Age is less than 30.")
	// } else if age < 40 {
	// 	fmt.Println("Age is less than 40.")
	// } else {
	// 	fmt.Println("Age is not less than 45")
	// }

	// names := []string{"Michelle", "Derrick", "Ilana", "Cynthia", "Layla"}

	// // The continue keyword in a for loop breaks out of the current loop iteration and returns into the main/parent loop
	// // The break keyword breaks out of a loop and doesn't continue cycling through the slice
	// for index, value := range names {
	// 	if index == 1 {
	// 		fmt.Println("Continuing at position", index)
	// 		continue
	// 	}
	// 	if index > 2 {
	// 		fmt.Println("Breaking at position %v", index)
	// 		break
	// 	}
	// 	fmt.Printf("The value at position %v is %v", index, value)
	// }

	// LOOPS
	// x := 0
	// for x < 5 {
	// 	fmt.Println("The value of x is:", x)
	// 	x++
	// }
	// for i := 0; i < 5; i++ {
	// 	fmt.Println("Value of i is:", i)
	// }

	// names := []string{"Ilana", "Michelle"}

	// // for i := 0; i < len(names); i++ {
	// // 	fmt.Println("Value of i is:", i)
	// // }

	// // How to cycle through a slice
	// // for index, value := range names {
	// // 	fmt.Printf("The value at index %v is %v\n", index, value)
	// // }

	// for _, value := range names {
	// 	fmt.Printf("The value is %v\n", value)
	// 	value = "new string"
	// }

	// fmt.Println(names)

	// THE STANDARD LIBRARY
	// greeting := "Hello friends"
	// // fmt.Println(strings.Contains(greeting, "Hello"))
	// // fmt.Println(strings.ReplaceAll(greeting, "Hello", "What's up"))

	// fmt.Println(strings.Index(greeting, "ll"))
	// // Strings.Split returns a slice
	// fmt.Println(strings.Split(greeting, " "))

	// fmt.Println("Original string value = ", greeting)

	// ages := []int{27, 2, 30, 4, 11, 56}
	// sort.Ints(ages)
	// fmt.Println(ages)

	// index := sort.SearchInts(ages, 30)
	// fmt.Println(index)

	// names := []string{"Michelle", "Derrick", "Ilana"}
	// sort.Strings(names)
	// fmt.Println(names)
	// fmt.Println(sort.SearchStrings(names, "Derrick"))
	// String
	// var nameOne string = "Michelle"
	// var nameTwo = "Derrick"
	// var nameThree string

	// fmt.Println(nameOne, nameTwo, nameThree)

	// nameThree = "Ilana"
	// nameFour := "Layla"
	// fmt.Println(nameOne, nameTwo, nameThree, nameFour)

	// INTEGERS
	// var ageOne int = 27
	// var ageTwo = 30
	// ageThree := 2

	// fmt.Println(ageOne, ageTwo, ageThree)

	// BITS AND MEMORY
	// var numOne int8 = 25
	// var numTwo int8 = -128
	// var numThree uint8 = 25

	// fmt.Println(numOne, numTwo, numThree)

	// var scoreOne float32 = 25.98
	// var scoreTwo float64 = 219387120847019201.2
	// scoreThree := 1.5

	// fmt.Println(scoreOne, scoreTwo, scoreThree)

	//Print
	// age := "27"
	// name := "Michelle"
	// fmt.Print("Hello, ")
	// fmt.Print("world!\n")
	// fmt.Print("New line\n")

	// //Println
	// fmt.Println("Hello ninjas!")
	// fmt.Println("Goodbye ninjas!")
	// fmt.Println("My age is,", age, "and my name is", name)

	// //Printf (formatted strings)
	// fmt.Printf("My age is %v and my name is %v \n", age, name)
	// fmt.Printf("My age is %q and my name is %q \n", age, name)
	// fmt.Printf("My age is of type %T \n", age)
	// fmt.Printf("You scored  %f points \n", 225.55)
	// fmt.Printf("You scored %0.1f points! \n", 225.55)

	// Sprintf (save formatted strings)
	// var str = fmt.Sprintf("You scored %0.1f points! \n", 225.55)
	// fmt.Println(str)

	// Arrays and Slices
	// var ages [3]int = [3]int{20, 25, 27}
	// var ages = [3]int{20, 25, 27}
	// names := [4]string{"Yoshi", "Mario", "Peach", "Bowser"}
	// names[1] = "Luigi"

	// fmt.Println(ages, len(ages))
	// fmt.Println(names, len(names))

	// // Slices (use arrays under the hood)
	// var scores = []int{100, 60, 40}
	// scores[2] = 25
	// scores = append(scores, 85)

	// fmt.Println(scores, len(scores))

	// SLICE RANGES
	// rangeOne := names[1:3]
	// rangeTwo := names[2:]
	// rangeThree := names[:3]

	// fmt.Println(rangeOne)
	// fmt.Println(rangeTwo)
	// fmt.Println(rangeThree)

	// rangeOne = append(rangeOne, "Koopa")
	// fmt.Println(rangeOne)
}
