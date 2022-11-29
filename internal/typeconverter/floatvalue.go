package typeconverter

import (
	"github.com/jinzhu/copier"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func WrapperFloatValueToFloat32() copier.TypeConverter {
	return copier.TypeConverter{
		SrcType: &wrapperspb.FloatValue{},
		DstType: float32(0),
		Fn: func(src interface{}) (interface{}, error) {
			if fv, ok := src.(*wrapperspb.FloatValue); ok {
				return fv.Value, nil
			}

			return nil, typeNotMatchError
		},
	}
}

func Float32ToWrapperBoolValue() copier.TypeConverter {
	return copier.TypeConverter{
		SrcType: float32(0),
		DstType: &wrapperspb.FloatValue{},
		Fn: func(src interface{}) (interface{}, error) {
			if f32, ok := src.(float32); ok {
				return wrapperspb.Float(f32), nil
			}
			return nil, typeNotMatchError
		},
	}
}
