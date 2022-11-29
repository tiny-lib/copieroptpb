package typeconverter

import (
	"github.com/jinzhu/copier"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func WrapperInt32ValueToInt32() copier.TypeConverter {
	return copier.TypeConverter{
		SrcType: &wrapperspb.Int32Value{},
		DstType: int32(0),
		Fn: func(src interface{}) (interface{}, error) {
			if i32, ok := src.(*wrapperspb.Int32Value); ok {
				return i32.Value, nil
			}

			return nil, typeNotMatchError
		},
	}
}

func Int32ToWrapperInt32Value() copier.TypeConverter {
	return copier.TypeConverter{
		SrcType: int32(0),
		DstType: &wrapperspb.Int32Value{},
		Fn: func(src interface{}) (interface{}, error) {
			if i32, ok := src.(int32); ok {
				return wrapperspb.Int32(i32), nil
			}
			return nil, typeNotMatchError
		},
	}
}

func WrapperUInt32ValueToUInt32() copier.TypeConverter {
	return copier.TypeConverter{
		SrcType: &wrapperspb.UInt32Value{},
		DstType: uint32(0),
		Fn: func(src interface{}) (interface{}, error) {
			if ui32, ok := src.(*wrapperspb.UInt32Value); ok {
				return ui32.Value, nil
			}

			return nil, typeNotMatchError
		},
	}
}

func Uint32ToWrapperUInt32Value() copier.TypeConverter {
	return copier.TypeConverter{
		SrcType: int32(0),
		DstType: &wrapperspb.UInt32Value{},
		Fn: func(src interface{}) (interface{}, error) {
			if ui32, ok := src.(uint32); ok {
				return wrapperspb.UInt32(ui32), nil
			}
			return nil, typeNotMatchError
		},
	}
}
