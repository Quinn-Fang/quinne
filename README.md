
# Quinne
**Goal**:  Daily used programming languages like Python, Java, C++ are general-purpose languages meant to serve most scenarios, with a price being too complicated for specific use cases within specific context. This project aims to provide an easier way for you to create your own programming language, thus your own way of describing something.

Quinne is a golang grammared programming language framework, making it possible to create your own simple, customized, context-specific programming language  with a focus on being:
* *Light*: Minimal dependencies (Antlr4 golang target mostly)
* *Simpler*: Deal with essential parts of a program only (how if-else expression should be judged, how functions should be run, etc...) without worrying about most compiler details (variables, symbol tables, scopes ...) that need to be done but not closely related to what you want to do.
* *Flexible*: Easy to define your own grammars when comes to specific usecases,  or under specific contexts.
* *Expandable:* Possible to add new grammars, say lambda, try-catch to golang grammar and even execute other languages like c by simply just write it in the middle of your go program.

## Getting started

### Getting Quinne
```bash
go get -u github.com/Quinn-Fang/quinne
```
### Create a golang program
 Create a golang program that resides anywhere in your project. Say you want to write a program to create
  a circuit diagram.
 ```go
 package samples_001
 
func main() {
	create_battery()
	create_switch()
	if SWITCH_ON {
		create_bulb("bulb-1")
	} else {
		create_bulb("buib-2")
	}
}
 ```

### Create Quinne handler
To manipulate this small program, you create a quinne handler, iterate through events that you need to handle
 in order to make the program meaningful.
 ```go
 func pgTest_2() {
	eventHandler := quinne.NewEventHandler("samples/sample_001.go")
	event, err := eventHandler.GetNextEvent()
	for err == nil {
		fmt.Printf("%+v\n", event)
		if event.GetEventType() == uspace.EventTypeFunction {
			// deal with how function show be executed here
			// and provide the return value
			fFunction := event.GetFunction(event.GetEventContext())
			if fFunction.GetFunctionName() == "create_battery" {

				fFunction.SetReturnValue(true)
				fmt.Printf("%+v %+v\n", fFunction, fFunction.GetReturnValue())
			} else if fFunction.GetFunctionName() == "create_switch" {
				fFunction.SetReturnValue("success")
				fmt.Printf("%+v %+v\n", fFunction, fFunction.GetReturnValue())
			} else if fFunction.GetFunctionName() == "create_bulb" {
				// Get Function arguments, maybe put some check on them ...
				params := fFunction.GetParams()
				for _, v := range params {
					fmt.Printf("%+v ", v.GetVariableValue())
				}

				// Set the return value for this particular function
				fFunction.SetReturnValue("success")
				fmt.Printf("%+v %+v\n\n", fFunction, fFunction.GetReturnValue())
			}
		} else if event.GetEventType() == uspace.EventTypeIfElseExpr {
			// Get the if expression or else-if expression and variables within that you
			// should provide value or assigned automatically if has defined previously

			ifElseExpr, ifElseExprVarNames := event.GetExpr(event.GetEventContext())
			if strings.Contains(ifElseExpr, "SWITCH_ON") {
				fmt.Printf("%+v %+v\n", ifElseExpr, ifElseExprVarNames)
				varMap := make(map[string]interface{}, 8)
				varMap["SWITCH_ON"] = false
				event.SetExpr(varMap)

				// filled variables automatically if already defined and can be accessed
				// by scope rules

				//event.FillExpr()
			}

		}

		event, err = eventHandler.GetNextEvent()

		fmt.Println()
	}
}

 func main() {
	pgTest_2()
}

  ```
