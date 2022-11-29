package typeconverter

import (
	"github.com/jinzhu/copier"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

func TimestampPbValueToTime() copier.TypeConverter {
	return copier.TypeConverter{
		SrcType: &timestamppb.Timestamp{},
		DstType: time.Time{},
		Fn: func(src interface{}) (interface{}, error) {
			if ts, ok := src.(*timestamppb.Timestamp); ok {
				if ts.IsValid() {
					return ts.AsTime(), nil
				}
				return time.Time{}, timestampPbIsNotValid
			}

			return nil, typeNotMatchError
		},
	}
}

func TimeToTimestampPbValue() copier.TypeConverter {
	return copier.TypeConverter{
		SrcType: time.Time{},
		DstType: &timestamppb.Timestamp{},
		Fn: func(src interface{}) (interface{}, error) {
			if t, ok := src.(time.Time); ok {
				return timestamppb.New(t), nil
			}
			return nil, typeNotMatchError
		},
	}
}
