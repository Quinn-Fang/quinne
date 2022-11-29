
# Quinne 
Quinne is a framework to create a new Domain-Specific programming language(DSL)

***Goal***:  To make it easier to create a new domain-specific language for your own purposes, with this new language, you can get these advantages:
1. ***Efficient communication***:  Developers can communicate with other teams in this new,  less complicated, easier-to-understand language, so participators from different teams can reach an agreement shortly.
2. ***Declarative Programming***: Most Business rule engines can be created by using logic branches and functions(actions) in a declarative manner, which is a better way to express solutions to difficult problems and domain experts can easily read and verify a set of rules.
3. ***Knowledge Centralization***: You can build a repository of knowledge, which is executable and strongly tied to the backend code and data, so it will be the most up-to-date version of business logic and can be served as documentation.
4.  ***Logic and Code Separation***: Code and data reside in the actual backend realization and business logic resides in the frontend DSL, meaning domain experts can focus on domain problems and developers can focus on code design and optimization.
5. ***Customized Grammar***: With Golang's basic grammar combined with python style try-catch, and Lamba grammar, developers can write shorter programs with enhanced functionalities by using functional programming.

**More Readings**:
- [Looking at programming language from a different aspect](https://medium.com/@quinnkunfang_5420/looking-at-programming-language-from-a-different-aspect-9e4544047c1e)


## Getting started

### Getting Quinne
```bash
go get -u github.com/Quinn-Fang/quinne
```
## Examples
### 1. Directional programming
#### Building Cloud Service Components(K8S)
```go
package main

func main() {
	Platform := "aws"
	Authenticate()
	Ret := CreateEKSCluster("EKSCluster-1", 3)
	additionalNodeCount := lambda currentNode, expectedNode int: expectedNode-currentNode if expectedNode > currentNode else 0
	currentNodeCount := getCurrentNodeCount()
	expectedNodeCount := 6
	additionalNode := additionalNodeCount(currentNodeCount, expectedNodeCount)
	expandEKSCluster(additionalNode)
}
```
- [Full example here  ](https://github.com/Quinn-Fang/quinne/blob/master/Documents/infra_example_1.md)
#### Driving route planner
```go
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

```

#### Login Design
```go
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

```

#### Strategy design
```go
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

```

### 2. Graphs:
#### Electronic circuit
```go
package samples_001

func main() {
	create_battery()
	create_switch()
	if SWITCH_ON {
		create_bulb("bulb-1")
	} else {
		create_bulb("bulb-2")
	}
}
```

#### Trees(Database index storage)
```go
package graph

func main() {
	searchForTreeNodeWithValue(36)
	if treeNodeExists {
		if rearrangeRequired {
			rearrangeTreeStructure()
			insertNewValue()
		} else {
			insertNewValue()
		}
	} else {
		createTreeNode()
	}
}

```

## Docs:
- [Introduction: ](https://github.com/Quinn-Fang/quinne/blob/master/Documents/info.md)
- [Examples: ](https://github.com/Quinn-Fang/quinne/blob/master/Documents/info.md)


