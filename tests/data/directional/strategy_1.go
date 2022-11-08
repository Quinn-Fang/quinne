package directional

func main() {
	getUserInfo()
	if isNewUser() {
		expireDate := "2022-7-31"
		amount = 3
		sendCoupon(amount, expireDate)
		if isDriver() {
			amount = 10
			expireDate := "2023-7-31"
			sendCoupon(amount, expireDate)
		}
	} else {
		if isUnderBadWeather() {
			amount = 10
			expireDate := "2023-7-31"
			nodificationMsg := "Bad weather coming!"
			sendCoupon(amount, expireDate)
			sendNotification(nodificationMsg)
		} else {

		}

		if OpenApp() {
			time := getLocalTime()
			sendMsgByTime(time)
		}
	}
}
