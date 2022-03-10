package main

import (
	"fmt"
	"log"
)

func print(r string, value int) {
	fmt.Printf("%s%d%%", r, value)
}

func main() {
	var charge int
	_, err := fmt.Scan(&charge)
	if err != nil {
		log.Fatal("Bad input")
	}
	switch {
	case charge < 10:
		print("", charge)
	case charge < 20:
		print("", charge)
	case charge < 30:
		print("", charge)
	case charge < 40:
		print("", charge)
	case charge < 50:
		print("", charge)
	case charge < 60:
		print("", charge)
	case charge < 70:
		print("", charge)
	case charge < 80:
		print("", charge)
	case charge < 90:
		print("", charge)
	default:
		print("", charge)
	}
}
