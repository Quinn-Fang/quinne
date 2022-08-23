package tests

//func Test_4(t *testing.T) {
//	eventHandler := quinne.NewEventHandler("data/sample_if_else.go")
//	event, err := eventHandler.GetNextEvent()
//	for err == nil {
//		fmt.Printf("%+v\n", event)
//		//if event.GetEventType() == uspace.EventTypeFunction {
//		//	// deal with how function show be executed here
//		//	// and provide the return value
//		//	fFunction := event.GetFunction(event.GetEventContext())
//		//	if fFunction.GetFunctionName() == "testFunction_1" {
//
//		//		fFunction.SetReturnValue("success_1")
//		//		// variables should be available
//
//		//		fReturn := fFunction.GetReturnValue().(string)
//		//		expected := "success_1"
//		//		if fReturn != expected {
//		//			errMsg := fmt.Sprintf(funcReturnNotEqual, fReturn, expected)
//		//			t.Error(errMsg)
//		//		}
//		//		var1, _ := event.GetSymTable().GetVariableByName("name")
//		//		var2, _ := event.GetSymTable().GetVariableByName("age")
//		//		vars := MakeVarSlice(var1, var2)
//		//		expects := MakeSlice("naveen", 5)
//
//		//		CompareVars(t, vars, expects)
//		//		eventLog := fmt.Sprintf("%T\n", event.GetEventContext())
//		//		t.Log(eventLog)
//		//	}
//		//} else if event.GetEventType() == uspace.EventTypeIfElseExpr {
//		//	ifElseExpr, ifElseExprVarNames := event.GetExpr(event.GetEventContext())
//		//	if strings.Contains(ifElseExpr, "age") {
//		//		fmt.Printf("%+v %+v\n", ifElseExpr, ifElseExprVarNames)
//		//		varMap := make(map[string]interface{}, 8)
//		//		varMap["age"] = 7
//		//		event.SetExpr(varMap)
//
//		//		//event.FillExpr()
//		//	} else if strings.Contains(ifElseExpr, "566") {
//		//		expected := "566>test1"
//		//		if ifElseExpr != expected {
//		//			t.Error(formatError(ifElseExprDoesNotMatch, ifElseExpr, expected))
//		//		}
//		//	} else if strings.Contains(ifElseExpr, "naveen") {
//		//		expected := "name==naveen"
//		//		if ifElseExpr != expected {
//		//			t.Error(formatError(ifElseExprDoesNotMatch, ifElseExpr, expected))
//		//		}
//
//		//	}
//
//		//}
//
//		event, err = eventHandler.GetNextEvent()
//
//		fmt.Println()
//	}
//}
