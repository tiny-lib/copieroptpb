package typeconverter

import (
	"github.com/jinzhu/copier"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func WrapperInt64ValueToInt64() copier.TypeConverter {
	return copier.TypeConverter{
		SrcType: &wrapperspb.Int64Value{},
		DstType: int64(0),
		Fn: func(src interface{}) (interface{}, error) {
			if i64, ok := src.(*wrapperspb.Int64Value); ok {
				return i64.Value, nil
			}

			return nil, typeNotMatchError
		},
	}
}

func Int64ToWrapperInt64Value() copier.TypeConverter {
	return copier.TypeConverter{
		SrcType: int64(0),
		DstType: &wrapperspb.Int64Value{},
		Fn: func(src interface{}) (interface{}, error) {
			if i64, ok := src.(int64); ok {
				return wrapperspb.Int64(i64), nil
			}
			return nil, typeNotMatchError
		},
	}
}

func WrapperUInt64ValueToUInt64() copier.TypeConverter {
	return copier.TypeConverter{
		SrcType: &wrapperspb.UInt64Value{},
		DstType: uint64(0),
		Fn: func(src interface{}) (interface{}, error) {
			if ui64, ok := src.(*wrapperspb.UInt64Value); ok {
				return ui64.Value, nil
			}

			return nil, typeNotMatchError
		},
	}
}

func Uint64ToWrapperUInt64Value() copier.TypeConverter {
	return copier.TypeConverter{
		SrcType: int64(0),
		DstType: &wrapperspb.UInt64Value{},
		Fn: func(src interface{}) (interface{}, error) {
			if ui64, ok := src.(uint64); ok {
				return wrapperspb.UInt64(ui64), nil
			}
			return nil, typeNotMatchError
		},
	}
}
