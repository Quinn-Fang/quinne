package consts

type OCType int

type MCType int

type ICType int

// indicators of the Line scope context
const (
	OCTypeUnSet    OCType = 0
	OCTypeIf              = 1
	OCTypeElseIf          = 2
	OCTypeElse            = 3
	OCTypeFuncDecl        = 4
)

const (
	MCTypeUnset MCType = 0
	MCTypeExpr         = 1
)

const (
	ICTypeUnset          ICType = 0
	ICTypeFuncName              = 1
	ICTypeFuncArgs              = 2
	ICTypeLambdaParams          = 3
	ICTypeLambdaExpr            = 4
	ICTypeLambdaRet             = 5
	ICTypeLambdaIfExpr          = 6
	ICTypeLambdaIfClause        = 7
)

//type LogicContextType int
//
//const (
//	LogicContextTypeIf     LogicContextType = 1
//	LogicContextTypeElseIf                  = 2
//	LogicContextTypeElse                    = 3
//)
