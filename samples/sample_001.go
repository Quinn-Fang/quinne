package main

func main() {
	create_battery()
	create_switch()
	if SWITCH_ON {
		create_bulb("bulb-1")
	} else {
		create_bulb("buib-2")
	}
}
