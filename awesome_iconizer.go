package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
)

func printBattery(r string, value int) {
	fmt.Printf("%s%d%%", r, value)
}

func handleBattery() {
	var charge int
	_, err := fmt.Scan(&charge)
	if err != nil {
		log.Fatal("Bad input")
	}
	switch {
	case charge < 10:
		printBattery("", charge)
	case charge < 20:
		printBattery("", charge)
	case charge < 30:
		printBattery("", charge)
	case charge < 40:
		printBattery("", charge)
	case charge < 50:
		printBattery("", charge)
	case charge < 60:
		printBattery("", charge)
	case charge < 70:
		printBattery("", charge)
	case charge < 80:
		printBattery("", charge)
	case charge < 90:
		printBattery("", charge)
	default:
		printBattery("", charge)
	}
}

func handleBluetooth() {
	reader := bufio.NewReader(os.Stdin)
	word, _ := reader.ReadString(' ')
	switch string(word) {
	case "no\n":
		fmt.Print("")
	case "yes\n":
		fmt.Print("")
	}
}

func handleCPU() {
	stat, err := os.Open("/proc/stat")
	if err != nil {
		fmt.Print("0")
		return
	}
	defer stat.Close()
	reader := bufio.NewReader(stat)
	var (
		user, nice, system, idle, iowait, irq, softirq, steal, guest, guest_nice int
	)
	fmt.Fscanf(reader, "cpu  %d %d %d %d %d %d %d %d %d", &user, &nice, &system, &idle, &iowait, &irq, &softirq, &steal, &guest, &guest_nice)

	total := float64(user + nice + system + idle + iowait + irq + softirq + steal + guest + guest_nice)
	var prevtotal, previdle float64

	prev, err := os.Open("/tmp/awesome_iconizer_cpu")
	if err == nil {
		fmt.Fscanf(bufio.NewReader(prev), "%f %f", &prevtotal, &previdle)
		prev.Close()
	}
	fmt.Print((1 - (float64(idle)-previdle)/(total-prevtotal)))

	save, err := os.Create("/tmp/awesome_iconizer_cpu")
	defer save.Close()
	save.WriteString(fmt.Sprintf("%f %f", total, float64(idle)))
}

func main() {
	mode := flag.String("mode", "", "")
	flag.Parse()
	switch *mode {
	case "battery":
		handleBattery()
	case "bluetooth":
		handleBluetooth()
	case "cpu":
		handleCPU()
	}
}
