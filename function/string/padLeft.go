package string

import (
	"github.com/project-flogo/core/data"
	"github.com/project-flogo/core/data/coerce"
	"github.com/project-flogo/core/data/expression/function"
)

type fnPadLeft struct {
}

func init() {
	function.Register(&fnPadLeft{})
}

func (s *fnPadLeft) Name() string {
	return "padLeft"
}

func (s *fnPadLeft) Sig() (paramTypes []data.Type, isVariadic bool) {
	return []data.Type{data.TypeString, data.TypeInt, data.TypeString}, false
}

func (s *fnPadLeft) Eval(in ...interface{}) (interface{}, error) {

	input, err := coerce.ToString(in[0])
	if err != nil {
		return nil, err
	}
	length, err := coerce.ToInt(in[1])
	if err != nil {
		return nil, err
	}
	padCharacter, err := coerce.ToString(in[2])
	if err != nil {
		return nil, err
	}

	for {
		input = padCharacter + input
		if len(input) >= length {
			return input[0:length], nil
		}
	}

}
