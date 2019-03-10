package util

import (
	"encoding/json"
	"fmt"
)

type Ticket struct {
	Num string

	Reds  []int
	Blues []int
}

func GetTicket(n string, reds []int, blues []int) Ticket {
	return Ticket{n, reds, blues}
}

func Marshal(i interface{}) []byte {
	data, _ := json.Marshal(i)
	fmt.Println(string(data))
	return data
}
