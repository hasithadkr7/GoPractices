package main

import "fmt"

var sampleMap map[string]int

func main() {
	sampleMap = map[string]int{
		"James": 50,
		"Ali":   39,
	}

	currencyMap := map[string]string{
		"AUD": "Australia Dollar",
		"GBP": "Great Britain Pound",
		"LKR": "Sri Lankan Ruppee",
		"CHF": "Switzerland Franc",
	}

	currencyMap["USD"] = "USA Dollar"
	fmt.Println("Currency with USD added: ", currencyMap)

	delete(currencyMap, "CHF")
	fmt.Println("Currency with CHF deleted: ", currencyMap)

	currencyMap["AUD"] = "New Zeland Dollar"
	fmt.Println("Currency with UD value replaced with NZD: ", currencyMap)

	for key, value := range currencyMap {
		fmt.Println("%v might be equal to: %v\n", key, value)
	}

}
