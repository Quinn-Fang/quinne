
# Quinne 
Quinne is a framework to create a new Domain-Specific programming language(DSL)

**Goal**:  To make it possible and easier to create a new domain specific language in a short time, with this newly created language in hand, development teams can have these advantages:
1. Developers can communicate with other teams in this new, most likely, easier to understand language. Thus most misunderstandings can be eliminated, decompositions and abstractions can alse be done at the stage.
2.  Protocals(Rules) can be set at this moment, any further changes can be described and applied in this new language, thus easier to know if any substantial code changes are required or exactly who to make these changes.
3. With this new language describing the domain-specific logic, less documentation is required.

**More Readings**:
- [Looking at programming language from a different aspect](https://medium.com/@quinnkunfang_5420/looking-at-programming-language-from-a-different-aspect-9e4544047c1e)


**General Design**: Parse the code stream and traverse ast, store essential data needed to provide a event-driven environment for the user to modify runtime
execution flow.

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
		create_bulb("bulb-2")
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

                // by calling SetExpr with emptyMap would filled variables automatically 
                // if already defined and can be accessed by scope rules
                
                // varMap := make(map[string]interface{}, 8)
                // event.SetExpr()

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

## Docs:
- [Introduction: ](https://github.com/Quinn-Fang/quinne/blob/master/Documents/info.md)
- [Examples: ](https://github.com/Quinn-Fang/quinne/blob/master/Documents/info.md)

