package typeconverter

import (
	"github.com/jinzhu/copier"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func WrapperDoubleValueToFloat64() copier.TypeConverter {
	return copier.TypeConverter{
		SrcType: &wrapperspb.DoubleValue{},
		DstType: float64(0),
		Fn: func(src interface{}) (interface{}, error) {
			if dv, ok := src.(*wrapperspb.DoubleValue); ok {
				return dv.Value, nil
			}

			return nil, typeNotMatchError
		},
	}
}

func Float64ToWrapperDoubleValue() copier.TypeConverter {
	return copier.TypeConverter{
		SrcType: float64(0),
		DstType: &wrapperspb.DoubleValue{},
		Fn: func(src interface{}) (interface{}, error) {
			if f64, ok := src.(float64); ok {
				return wrapperspb.Double(f64), nil
			}
			return nil, typeNotMatchError
		},
	}
}
