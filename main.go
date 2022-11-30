package main

import (
	"encoding/json"
	"log"
)

func main() {
	// b, err := readJson("example1.json")
	// if err != nil {
	// 	log.SetFlags(log.Lshortfile | log.LstdFlags)
	// 	log.Println("err:", err)
	// }
	// getExample1(b)

	b, err := readJson("example2.json")
	if err != nil {
		log.SetFlags(log.Lshortfile | log.LstdFlags)
		log.Println("err:", err)
	}
	getExample2(b)

}

// sportName => soccer
func getSportMarketLabelMap(b []byte, sportName string) map[string][]string {
	m := map[string]map[string][]string{}
	json.Unmarshal(b, &m)
	value, found := m[sportName]
	if found {
		return value
	}
	return map[string][]string{}
}
