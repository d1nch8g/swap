package main

import (
	"fmt"

	"ion.lc/d1nhc8g/bitchange/bestchange"
)

func main() {
	client := bestchange.NewClient("2dcd269e14d6bf5060e9df0fc7ab16a7")

	rez, err := client.Rates(21, 209, 10, false, true) // SBP to TON
	if err != nil {
		panic(err)
	}
	fmt.Println(" ===== SBT-TON =====")
	bestchange.PrintTable(rez)

	rez, err = client.Rates(209, 21, 10, true, true) // TON to SBP
	if err != nil {
		panic(err)
	}
	fmt.Println(" ===== TON-SBP =====")
	bestchange.PrintTable(rez)
}
