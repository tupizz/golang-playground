package main

import "fmt"

func main() {
	// first way to declare a map
	var m map[string]string
	fmt.Println(m == nil) // true
	//m["Tadeu"] = "Tadeu" // panic: assignment to entry in nil map

	// second way to declare a map
	var myMap = make(map[string]string)
	fmt.Println(myMap == nil) // false

	// third way to declare a map
	var myMap2 = map[string]string{}
	fmt.Println(myMap2 == nil) // false
	myMap2["Tadeu"] = "Tadeu"
	myMap2["Idade"] = "27"
	fmt.Println(myMap2)

	var myMap3 = map[string]string{
		"Nome":      "Tadeu",
		"Idade":     "27",
		"Profissao": "Desenvolvedor",
	}
	fmt.Println(myMap3)

	// map of slices
	var slicesMap = map[string][]int{
		"slice1": []int{1, 2, 3},
		"slice2": {4, 5, 6},
	}
	fmt.Println(slicesMap)

	valor, ok := slicesMap["slice3"]
	if ok {
		fmt.Println(valor)
	} else {
		fmt.Println("Key not found")
	}

	valor, ok = slicesMap["slice2"]
	if ok {
		fmt.Println(valor)
	} else {
		fmt.Println("Key not found")
	}

	// 0 is the default value for int and string in go

	delete(slicesMap, "slice1")
	fmt.Println(slicesMap)

	// important when you have NaN as keys in your map NaN is float64, and NaN is not equal to anything, neither itself
	clear(slicesMap)
	fmt.Println(slicesMap)

	lastMap := map[string]string{
		"Nome":      "Tadeu Humberto",
		"Random":    "1212121ssdadas",
		"Idade":     "30",
		"Profissao": "Empreendedor",
	}

	// Iterating over a map keys and values
	for k, v := range lastMap {
		if k == "Random" {
			delete(lastMap, k)
			continue
		}
		fmt.Println(k, v)
	}
}
