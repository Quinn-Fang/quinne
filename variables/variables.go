package variables

type VTypeEnum int

const (
	vTypeInt     VTypeEnum = 0
	vTypeString            = 1
	vTypeFloat64           = 2
	vTypeBool              = 3
	vTypeMap               = 4
)

type Variable struct {
	vName  string
	vType  VTypeEnum
	vValue interface{}
}

func NewVariable(variableName string, variableType VTypeEnum, variableValue interface{}) *Variable {
	newVariable := Variable{
		vName:  variableName,
		vType:  variableType,
		vValue: variableValue,
	}

	return &newVariable
}

func (this *Variable) GetVariableName() string {
	return this.vName
}

func (this *Variable) GetVariableType() VTypeEnum {
	return this.vType
}

func (this *Variable) GetVariableValue() interface{} {
	return this.vValue
}
