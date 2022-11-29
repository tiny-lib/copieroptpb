package typeconverter

import (
	"github.com/jinzhu/copier"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func WrapperStringValueToString() copier.TypeConverter {
	return copier.TypeConverter{
		SrcType: &wrapperspb.StringValue{},
		DstType: copier.String,
		Fn: func(src interface{}) (interface{}, error) {
			if sv, ok := src.(*wrapperspb.StringValue); ok {
				return sv.String(), nil
			}

			return nil, typeNotMatchError
		},
	}
}

func StringToWrapperStringValue() copier.TypeConverter {
	return copier.TypeConverter{
		SrcType: copier.String,
		DstType: &wrapperspb.StringValue{},
		Fn: func(src interface{}) (interface{}, error) {
			if s, ok := src.(string); ok {
				return wrapperspb.String(s), nil
			}
			return nil, typeNotMatchError
		},
	}
}
