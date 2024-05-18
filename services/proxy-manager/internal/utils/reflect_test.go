package reflectutils

import (
	"testing"
)

type testStruct struct {
	FloatValue     float64       `env:"FLOAT_VALUE"`
	IntValue       int           `env:"INT_VALUE"`
	SubStructValue testSubStruct `env:"SUB_STRUCT_VALUE"`
}

type testSubStruct struct {
	StringValue string `env:"STRING_VALUE"`
}

func TestStructToEnv(t *testing.T) {
}

func Test_isZeroValue(t *testing.T) {
}

func Test_fieldValueToString(t *testing.T) {
}

func Test_readStructMetadata(t *testing.T) {
}
