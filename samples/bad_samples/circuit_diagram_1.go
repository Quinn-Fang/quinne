package main

import "fmt"

type batteryNode struct {
	name string
	prev interface{}
	next interface{}
}

type switchNode struct {
	name    string
	prev    interface{}
	onNode  interface{}
	offNode interface{}
}

type bulbNode struct {
	name string
	prev interface{}
	next interface{}
}

func createBattery(batteryName string) *batteryNode {
	// codes to create a battery node
	bNode := &batteryNode{
		name: batteryName,
	}
	fmt.Println("Successfully created battery.")
	return bNode
}

func createSwitch(switchName string) *switchNode {
	// codes to create a switch node
	sNode := &switchNode{
		name: switchName,
	}
	fmt.Println("Successfully created switch.")
	return sNode
}

func createBulb(bulbName string) *bulbNode {
	// codes to create a bulb node
	bNode := &bulbNode{
		name: bulbName,
	}
	fmt.Println("Successfully created bulb.")
	return bNode
}

func main() {
	battery := createBattery("Battery 1")
	sSwitch := createSwitch("switch A")
	battery.next = sSwitch
	sSwitch.prev = battery

	bulb_1 := createBulb("Bulb 1")
	bulb_2 := createBulb("Bulb 2")

	bulb_1.prev = sSwitch
	bulb_2.prev = sSwitch

	sSwitch.onNode = bulb_1
	sSwitch.offNode = bulb_2
}
