package consts

type OCType int

type MCType int

type ICType int

// indicators of the Line scope context
const (
	OCTypeDefault  OCType = 0
	OCTypeIf              = 1
	OCTypeElseIf          = 2
	OCTypeElse            = 3
	OCTypeFuncDecl        = 4
)

const (
	MCTypeDefault MCType = 0
	MCTypeExpr           = 1
)

const (
	ICTypeDefault  ICType = 0
	ICTypeFuncName        = 1
	ICTypeFuncArgs        = 2
)

//type LogicContextType int
//
//const (
//	LogicContextTypeIf     LogicContextType = 1
//	LogicContextTypeElseIf                  = 2
//	LogicContextTypeElse                    = 3
//)
