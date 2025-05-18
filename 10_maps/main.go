package main

import (
	"fmt"
	"maps"
)

func main() {
	// creating map
	// m := make(map[string]string)

	// setting an element
	// m["name"] = "golang"
	// m["area"] = "backend"

	// getting an element
	// fmt.Println(m["name"])
	// fmt.Println(m["phone"]) // Note: if the key is not present, then it will return the default value of that type.

	// length
	// fmt.Println(len(m))

	// Delete a value
	// delete(m, "area")

	// Clear a map
	// clear(m)

	// m := map[string]string{"name": "Amal", "lang": "golang"}
	
	// for key, val := range m {
		// 	fmt.Println(key, val)
		// }

		// To check if a particular key exists inside a map
		// data, exists := m["age"]
		
		// fmt.Println(data, exists)
		m1 := map[string]string{"name": "Amal", "lang": "golang"}
		m2 := map[string]string{"name": "Amal", "lang": "golang"}

		fmt.Println(maps.Equal(m1, m2))
		
}