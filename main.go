package main

import "fmt"

// An array is a sequence of elements that has a specific length.
func arrays() {
	// Here we have an array of 10 integers. If we print out this variable, what
	// do you think we'll see?
	var intArray [10]int
	fmt.Printf("integer array: %v\n", intArray)

	// As you'd expect we can get and set values using the array.
	intArray[4] = 42
	fmt.Printf("0th integer: %v\n", intArray[0])
	fmt.Printf("4th integer: %v\n", intArray[4])

	// We can also define literal arrays like so. Here is a 3 element array of
	// strings.
	stringArray := [3]string{"foo", "bar", "baz"}
	fmt.Printf("string array: %v\n", stringArray)

	// We can also let the compiler count the elements for us.
	anotherStringArray := [...]string{"qux", "quux", "quuz"}
	fmt.Printf("another string array: %v\n", anotherStringArray)
}

// In Go slices are built on top of arrays. It is best described as a
// descriptor of an array segment. A slice is made up of a pointer to an
// array, the length of segment and the capacity of the segment.
func slices() {
	// A slice can be declared in a very similar way to an array, we just don't
	// bother to set the length.
	stringSlice := []string{"foo", "bar", "baz"}

	// What do you think the contents and length of the slice will be?
	fmt.Printf("string slice: %v\n", stringSlice)
	fmt.Printf("string slice length: %d\n", len(stringSlice))
}

func makingSlices() {
	// We can define an empty slice like so.
	var intSlice []int

	// What do you think the contents, length and capacity of the slice will be?
	fmt.Printf("int slice: %v\n", intSlice)
	fmt.Printf("int slice length: %d\n", len(intSlice))
	fmt.Printf("int slice capacity: %d\n", cap(intSlice))

	// If we uncomment this line we'll get an error. Why do you think that is?
	// intSlice[0] = 2

	// We use the make function here to create a slice of integers. The slice
	// points to an array 3 elements long. Why might it be preferable to create
	// a slice with a pre-defined length.
	intSlice = make([]int, 3)

	// What do you think the contents and length of the slice will be?
	fmt.Printf("int slice: %v\n", intSlice)
	fmt.Printf("int slice length: %d\n", len(intSlice))
	fmt.Printf("int slice capacity: %d\n", cap(intSlice))

	// Just like arrays we can get and set elements in the slice.
	intSlice[0] = 42
	fmt.Printf("0th integer: %v\n", intSlice[0])
	fmt.Printf("1st integer: %v\n", intSlice[1])

	// If we uncomment this line we'll get an error. Why do you think that is?
	// intSlice[3] = 2
}

func modifyingSlices() {
	// Here we've defined a slice of strings.
	stringSlice := []string{"foo", "bar", "baz", "qux", "quux", "quuz"}

	// We can create a new slice from our original slice like so. The syntax
	// in this case is slice[low:high].
	newSlice := stringSlice[2:5]

	// What will be the output of the new slice?
	fmt.Printf("string slice: %v\n", newSlice)

	// We can "re-slice" our original slice as many times as we like. It does
	// not affect the underlying array.
	anotherSlice := stringSlice[1:4]
	fmt.Printf("another slice: %v\n", anotherSlice)
	fmt.Printf("original slice: %v\n", stringSlice)

	// We can omit the low or high to just get the start or end of the slice.
	startOfSlice := stringSlice[2:]
	endOfSlice := stringSlice[:5]
	fmt.Printf("start of slice: %v\n", startOfSlice)
	fmt.Printf("end of slice: %v\n", endOfSlice)
}

func growingSlices() {
	// Here we have a slice of 3 integers.
	intSlice := []int{1, 2, 3}

	// Here we add more integers to the slice. Watch what happens to the length
	// and capacity as we add elements.
	i := 3
	for i < 9 {
		intSlice = append(intSlice, i)

		fmt.Printf("int slice length: %d, capacity: %d\n", len(intSlice), cap(intSlice))

		i++
	}
}

func main() {
	arrays()

	fmt.Println()

	slices()

	fmt.Println()

	makingSlices()

	fmt.Println()

	modifyingSlices()

	fmt.Println()

	growingSlices()
}
