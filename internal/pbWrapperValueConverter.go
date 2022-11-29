package internal

import (
	"github.com/czyt/copieroptpb/internal/typeconverter"
	"github.com/jinzhu/copier"
)

func PbWrapperValueConverters() []copier.TypeConverter {
	return []copier.TypeConverter{
		typeconverter.WrapperStringValueToString(),
		typeconverter.StringToWrapperStringValue(),
		typeconverter.WrapperBoolValueToBool(),
		typeconverter.BoolToWrapperBoolValue(),
		typeconverter.WrapperBytesValueToBytes(),
		typeconverter.BytesToWrapperBytesValue(),
		typeconverter.WrapperDoubleValueToFloat64(),
		typeconverter.Float64ToWrapperDoubleValue(),
		typeconverter.WrapperFloatValueToFloat32(),
		typeconverter.Float32ToWrapperBoolValue(),
		typeconverter.WrapperInt32ValueToInt32(),
		typeconverter.Int32ToWrapperInt32Value(),
		typeconverter.WrapperUInt32ValueToUInt32(),
		typeconverter.Uint32ToWrapperUInt32Value(),
		typeconverter.WrapperInt64ValueToInt64(),
		typeconverter.Int64ToWrapperInt64Value(),
		typeconverter.WrapperUInt64ValueToUInt64(),
		typeconverter.Uint64ToWrapperUInt64Value(),
		typeconverter.TimestampPbValueToTime(),
		typeconverter.TimeToTimestampPbValue(),
	}
}
