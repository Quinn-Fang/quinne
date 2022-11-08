package directional

func main() {
	activateVehicle("Mercedes_1")
	weather := GetWeatherStatus()
	if weather == "raining" {
		showNotificationOnDevice("IPhone_device_1", "it's raining heavily, bring umbrella")
		driveToDestination("pandas_1")
	} else if weather == "snowing" {
		showNotificationOnDevice("Apple_watch_device_2", "it's snowing, wear more clothes")
		driveToBackDoor()
		openBackDoor()
		driveToCompany()
	}
}
