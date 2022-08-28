package consts

type OCType int

type MCType int

type ICType int

const (
	OCTypeIf       OCType = 1
	OCTypeElseIf          = 2
	OCTypeElse            = 3
	OCTypeFuncDecl        = 4
)

const (
	MCTypeExpr MCType = 1
)

const (
	ICTypeFuncName ICType = 1
	ICTypeFuncArgs        = 2
)

type LogicContextType int

const (
	LogicContextTypeIf     LogicContextType = 1
	LogicContextTypeElseIf                  = 2
	LogicContextTypeElse                    = 3
)
