package typeconverter

import (
	"github.com/jinzhu/copier"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func WrapperBytesValueToBytes() copier.TypeConverter {
	return copier.TypeConverter{
		SrcType: &wrapperspb.BytesValue{},
		DstType: []byte{},
		Fn: func(src interface{}) (interface{}, error) {
			if bv, ok := src.(*wrapperspb.BytesValue); ok {
				return bv.Value, nil
			}

			return nil, typeNotMatchError
		},
	}
}

func BytesToWrapperBytesValue() copier.TypeConverter {
	return copier.TypeConverter{
		SrcType: []byte{},
		DstType: &wrapperspb.BytesValue{},
		Fn: func(src interface{}) (interface{}, error) {
			if b, ok := src.([]byte); ok {
				return wrapperspb.Bytes(b), nil
			}
			return nil, typeNotMatchError
		},
	}
}
