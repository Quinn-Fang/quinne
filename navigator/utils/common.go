package utils

import (
	"github.com/Knetic/govaluate"
)

func EvaluateExpr(expr string, parameters map[string]interface{}) bool {
	expression, err := govaluate.NewEvaluableExpression(expr)

	result, err := expression.Evaluate(parameters)
	if err != nil {
		panic(err)
	}

	if ret, ok := result.(bool); !ok {
		panic("Expr result not bool")
	} else {
		return ret
	}
}
