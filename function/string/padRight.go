package string

import (
	"github.com/project-flogo/core/data"
	"github.com/project-flogo/core/data/coerce"
	"github.com/project-flogo/core/data/expression/function"
)

type fnPadRight struct {
}

func init() {
	function.Register(&fnPadRight{})
}

func (s *fnPadRight) Name() string {
	return "padRight"
}

func (s *fnPadRight) Sig() (paramTypes []data.Type, isVariadic bool) {
	return []data.Type{data.TypeString, data.TypeInt, data.TypeString}, false
}

func (s *fnPadRight) Eval(in ...interface{}) (interface{}, error) {

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
		input += padCharacter
		if len(input) >= length {
			return input[0:length], nil
		}
	}

}
