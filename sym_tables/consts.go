package sym_tables

type ContextType int

const (
	ContextTypeDefault      ContextType = 1
	ContextTypeFunctionName             = 2
	ContextTypeFunctionArgs             = 3
	ContextTypeIf                       = 4
	ContextTypeElseIf                   = 5
	ContextTypeElse                     = 6
	ContextTypeBlock                    = 7
	ContextTypeFuncDecl                 = 8
	ContextTypeForLoop                  = 9
)
