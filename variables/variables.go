package variables

type VTypeEnum int

const (
	VTypeUndefined        VTypeEnum = 0
	VTypeInt                        = 1
	VTypeString                     = 2
	VTypeFloat                      = 3
	VTypeBool                       = 4
	VTypeMap                        = 5
	VTypeFunctionReturned           = 6
)

type Variable struct {
	vName                  string
	vType                  VTypeEnum
	vValue                 interface{}
	vIndex                 int
	vFunctionReturnedIndex int
}

func NewEmptyVariable() *Variable {
	newVariable := Variable{}

	return &newVariable
}

func NewVariable(variableName string, variableType VTypeEnum, variableValue interface{}, index int) *Variable {
	newVariable := Variable{
		vName:  variableName,
		vType:  variableType,
		vValue: variableValue,
		vIndex: index,
	}

	return &newVariable
}

func (this *Variable) GetVariableName() string {
	return this.vName
}

func (this *Variable) SetVariableName(variableName string) {
	this.vName = variableName
}

func (this *Variable) GetVariableType() VTypeEnum {
	return this.vType
}

func (this *Variable) GetVariableValue() interface{} {
	return this.vValue
}

func (this *Variable) GetVariableIndex() int {
	return this.vIndex
}
