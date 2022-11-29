package typeconverter

import (
	"github.com/jinzhu/copier"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func WrapperBoolValueToBool() copier.TypeConverter {
	return copier.TypeConverter{
		SrcType: &wrapperspb.BoolValue{},
		DstType: false,
		Fn: func(src interface{}) (interface{}, error) {
			if bv, ok := src.(*wrapperspb.BoolValue); ok {
				return bv.Value, nil
			}

			return nil, typeNotMatchError
		},
	}
}

func BoolToWrapperBoolValue() copier.TypeConverter {
	return copier.TypeConverter{
		SrcType: false,
		DstType: &wrapperspb.BoolValue{},
		Fn: func(src interface{}) (interface{}, error) {
			if b, ok := src.(bool); ok {
				return wrapperspb.Bool(b), nil
			}
			return nil, typeNotMatchError
		},
	}
}
