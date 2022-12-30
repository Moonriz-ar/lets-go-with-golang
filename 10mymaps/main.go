package main

import "fmt"

func main() {
	fmt.Println("maps")

	languages := make(map[string]string)

	languages["JS"] = "Javascript"
	languages["RB"] = "Ruby"
	languages["PY"] = "python"

	fmt.Println("Map of all languages: ", languages)
	fmt.Println("JS is short for: ", languages["JS"])

	delete(languages, "RB")
	fmt.Println("Map of all languages: ", languages)
	
	for key, value := range languages {
		fmt.Printf("For key %v, value is %v\n", key, value)
	}
}