package arraysslices

import "fmt"

func Arraysslices() {
	fmt.Println("Arrays e Slices:")
	var testarray = [2]string{"test", "test"}
	var testslice = []string{"test"}
	// O array só pode ter dois valores
	fmt.Printf("Antes\nArray: %v\tSlice:%v\n", testarray, testslice)
	testslice = append(testslice, "test")
	fmt.Printf("Depois\nArray: %v\tSlice:%v\n", testarray, testslice)
}