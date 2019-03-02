package apis

import (
	"encoding/json"
	"fmt"
)

func TestMapStringInterface() string {
	str := `{
		"Person": {
			  "Name": "BLABLA",
			  "Age": 23
		},
		"Position": "Full-Stack Developer",
		"Main": "Back-End Developer"
	}`
	var result map[string]interface{}
	json.Unmarshal([]byte(str), &result)
	fmt.Println("result ========> ", result)
	fmt.Println()
	var person map[string]interface{}
	// if p, ok := result["Person"].(map[string]interface{}); ok {
	// 	fmt.Println("p ========> ", p)
	// 	fmt.Println()
	// 	p["Name"] = "Change Name"
	// 	p["Age"] = 24
	// 	fmt.Println("result ========> ", result)
	// 	fmt.Println()
	// }
	if p, ok := result["Person"].(map[string]interface{}); ok {
		person = p
	}
	person["Name"] = "LazyDev"
	person["Age"] = 25
	fmt.Println("person ========> ", person)
	fmt.Println()
	fmt.Println("result ========> ", result)
	fmt.Println()

	return str
}
