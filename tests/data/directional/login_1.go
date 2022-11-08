package directional

func main() {
	getUserPhoneNum()
	checkUserExistance()
	if exists {
		ssoStatus := checkSSOStatus()
		ssoLogin(ssoStatus)
	} else {
		createNewUser()
		userStatusConfirm()
		userLogin()
	}
}
