package pgtype

import (
	"math"
	"reflect"
	"time"

	"github.com/pkg/errors"
)

const maxUint = ^uint(0)
const maxInt = int(maxUint >> 1)
const minInt = -maxInt - 1

// underlyingNumberType gets the underlying type that can be converted to Int2, Int4, Int8, Float4, or Float8
func underlyingNumberType(val interface{}) (interface{}, bool) {
	refVal := reflect.ValueOf(val)

	switch refVal.Kind() {
	case reflect.Ptr:
		if refVal.IsNil() {
			return nil, false
		}
		convVal := refVal.Elem().Interface()
		return convVal, true
	case reflect.Int:
		convVal := int(refVal.Int())
		return convVal, reflect.TypeOf(convVal) != refVal.Type()
	case reflect.Int8:
		convVal := int8(refVal.Int())
		return convVal, reflect.TypeOf(convVal) != refVal.Type()
	case reflect.Int16:
		convVal := int16(refVal.Int())
		return convVal, reflect.TypeOf(convVal) != refVal.Type()
	case reflect.Int32:
		convVal := int32(refVal.Int())
		return convVal, reflect.TypeOf(convVal) != refVal.Type()
	case reflect.Int64:
		convVal := int64(refVal.Int())
		return convVal, reflect.TypeOf(convVal) != refVal.Type()
	case reflect.Uint:
		convVal := uint(refVal.Uint())
		return convVal, reflect.TypeOf(convVal) != refVal.Type()
	case reflect.Uint8:
		convVal := uint8(refVal.Uint())
		return convVal, reflect.TypeOf(convVal) != refVal.Type()
	case reflect.Uint16:
		convVal := uint16(refVal.Uint())
		return convVal, reflect.TypeOf(convVal) != refVal.Type()
	case reflect.Uint32:
		convVal := uint32(refVal.Uint())
		return convVal, reflect.TypeOf(convVal) != refVal.Type()
	case reflect.Uint64:
		convVal := uint64(refVal.Uint())
		return convVal, reflect.TypeOf(convVal) != refVal.Type()
	case reflect.Float32:
		convVal := float32(refVal.Float())
		return convVal, reflect.TypeOf(convVal) != refVal.Type()
	case reflect.Float64:
		convVal := refVal.Float()
		return convVal, reflect.TypeOf(convVal) != refVal.Type()
	case reflect.String:
		convVal := refVal.String()
		return convVal, reflect.TypeOf(convVal) != refVal.Type()
	}

	return nil, false
}

// underlyingBoolType gets the underlying type that can be converted to Bool
func underlyingBoolType(val interface{}) (interface{}, bool) {
	refVal := reflect.ValueOf(val)

	switch refVal.Kind() {
	case reflect.Ptr:
		if refVal.IsNil() {
			return nil, false
		}
		convVal := refVal.Elem().Interface()
		return convVal, true
	case reflect.Bool:
		convVal := refVal.Bool()
		return convVal, reflect.TypeOf(convVal) != refVal.Type()
	}

	return nil, false
}

// underlyingBytesType gets the underlying type that can be converted to []byte
func underlyingBytesType(val interface{}) (interface{}, bool) {
	refVal := reflect.ValueOf(val)

	switch refVal.Kind() {
	case reflect.Ptr:
		if refVal.IsNil() {
			return nil, false
		}
		convVal := refVal.Elem().Interface()
		return convVal, true
	case reflect.Slice:
		if refVal.Type().Elem().Kind() == reflect.Uint8 {
			convVal := refVal.Bytes()
			return convVal, reflect.TypeOf(convVal) != refVal.Type()
		}
	}

	return nil, false
}

// underlyingStringType gets the underlying type that can be converted to String
func underlyingStringType(val interface{}) (interface{}, bool) {
	refVal := reflect.ValueOf(val)

	switch refVal.Kind() {
	case reflect.Ptr:
		if refVal.IsNil() {
			return nil, false
		}
		convVal := refVal.Elem().Interface()
		return convVal, true
	case reflect.String:
		convVal := refVal.String()
		return convVal, reflect.TypeOf(convVal) != refVal.Type()
	}

	return nil, false
}

// underlyingPtrType dereferences a pointer
func underlyingPtrType(val interface{}) (interface{}, bool) {
	refVal := reflect.ValueOf(val)

	switch refVal.Kind() {
	case reflect.Ptr:
		if refVal.IsNil() {
			return nil, false
		}
		convVal := refVal.Elem().Interface()
		return convVal, true
	}

	return nil, false
}

// underlyingTimeType gets the underlying type that can be converted to time.Time
func underlyingTimeType(val interface{}) (interface{}, bool) {
	refVal := reflect.ValueOf(val)

	switch refVal.Kind() {
	case reflect.Ptr:
		if refVal.IsNil() {
			return time.Time{}, false
		}
		convVal := refVal.Elem().Interface()
		return convVal, true
	}

	timeType := reflect.TypeOf(time.Time{})
	if refVal.Type().ConvertibleTo(timeType) {
		return refVal.Convert(timeType).Interface(), true
	}

	return time.Time{}, false
}

// underlyingSliceType gets the underlying slice type
func underlyingSliceType(val interface{}) (interface{}, bool) {
	refVal := reflect.ValueOf(val)

	switch refVal.Kind() {
	case reflect.Ptr:
		if refVal.IsNil() {
			return nil, false
		}
		convVal := refVal.Elem().Interface()
		return convVal, true
	case reflect.Slice:
		baseSliceType := reflect.SliceOf(refVal.Type().Elem())
		if refVal.Type().ConvertibleTo(baseSliceType) {
			convVal := refVal.Convert(baseSliceType)
			return convVal.Interface(), reflect.TypeOf(convVal.Interface()) != refVal.Type()
		}
	}

	return nil, false
}

func int64AssignTo(srcVal int64, srcStatus Status, dst interface{}) error {
	if srcStatus == Present {
		switch v := dst.(type) {
		case *int:
			if srcVal < int64(minInt) {
				return errors.Errorf("%d is less than minimum value for int", srcVal)
			} else if srcVal > int64(maxInt) {
				return errors.Errorf("%d is greater than maximum value for int", srcVal)
			}
			*v = int(srcVal)
		case *int8:
			if srcVal < math.MinInt8 {
				return errors.Errorf("%d is less than minimum value for int8", srcVal)
			} else if srcVal > math.MaxInt8 {
				return errors.Errorf("%d is greater than maximum value for int8", srcVal)
			}
			*v = int8(srcVal)
		case *int16:
			if srcVal < math.MinInt16 {
				return errors.Errorf("%d is less than minimum value for int16", srcVal)
			} else if srcVal > math.MaxInt16 {
				return errors.Errorf("%d is greater than maximum value for int16", srcVal)
			}
			*v = int16(srcVal)
		case *int32:
			if srcVal < math.MinInt32 {
				return errors.Errorf("%d is less than minimum value for int32", srcVal)
			} else if srcVal > math.MaxInt32 {
				return errors.Errorf("%d is greater than maximum value for int32", srcVal)
			}
			*v = int32(srcVal)
		case *int64:
			if srcVal < math.MinInt64 {
				return errors.Errorf("%d is less than minimum value for int64", srcVal)
			} else if srcVal > math.MaxInt64 {
				return errors.Errorf("%d is greater than maximum value for int64", srcVal)
			}
			*v = int64(srcVal)
		case *uint:
			if srcVal < 0 {
				return errors.Errorf("%d is less than zero for uint", srcVal)
			} else if uint64(srcVal) > uint64(maxUint) {
				return errors.Errorf("%d is greater than maximum value for uint", srcVal)
			}
			*v = uint(srcVal)
		case *uint8:
			if srcVal < 0 {
				return errors.Errorf("%d is less than zero for uint8", srcVal)
			} else if srcVal > math.MaxUint8 {
				return errors.Errorf("%d is greater than maximum value for uint8", srcVal)
			}
			*v = uint8(srcVal)
		case *uint16:
			if srcVal < 0 {
				return errors.Errorf("%d is less than zero for uint32", srcVal)
			} else if srcVal > math.MaxUint16 {
				return errors.Errorf("%d is greater than maximum value for uint16", srcVal)
			}
			*v = uint16(srcVal)
		case *uint32:
			if srcVal < 0 {
				return errors.Errorf("%d is less than zero for uint32", srcVal)
			} else if srcVal > math.MaxUint32 {
				return errors.Errorf("%d is greater than maximum value for uint32", srcVal)
			}
			*v = uint32(srcVal)
		case *uint64:
			if srcVal < 0 {
				return errors.Errorf("%d is less than zero for uint64", srcVal)
			}
			*v = uint64(srcVal)
		default:
			if v := reflect.ValueOf(dst); v.Kind() == reflect.Ptr {
				el := v.Elem()
				switch el.Kind() {
				// if dst is a pointer to pointer, strip the pointer and try again
				case reflect.Ptr:
					if el.IsNil() {
						// allocate destination
						el.Set(reflect.New(el.Type().Elem()))
					}
					return int64AssignTo(srcVal, srcStatus, el.Interface())
				case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
					if el.OverflowInt(int64(srcVal)) {
						return errors.Errorf("cannot put %d into %T", srcVal, dst)
					}
					el.SetInt(int64(srcVal))
					return nil
				case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
					if srcVal < 0 {
						return errors.Errorf("%d is less than zero for %T", srcVal, dst)
					}
					if el.OverflowUint(uint64(srcVal)) {
						return errors.Errorf("cannot put %d into %T", srcVal, dst)
					}
					el.SetUint(uint64(srcVal))
					return nil
				}
			}
			return errors.Errorf("cannot assign %v into %T", srcVal, dst)
		}
		return nil
	}
	if srcStatus == Null {
		return NullAssignTo(dst)
	}

	return errors.Errorf("cannot assign %v %v into %T", srcVal, srcStatus, dst)
}

func float64AssignTo(srcVal float64, srcStatus Status, dst interface{}) error {
	if srcStatus == Present {
		switch v := dst.(type) {
		case *float32:
			*v = float32(srcVal)
		case *float64:
			*v = srcVal
		default:
			if v := reflect.ValueOf(dst); v.Kind() == reflect.Ptr {
				el := v.Elem()
				switch el.Kind() {
				// if dst is a pointer to pointer, strip the pointer and try again
				case reflect.Ptr:
					if el.IsNil() {
						// allocate destination
						el.Set(reflect.New(el.Type().Elem()))
					}
					return float64AssignTo(srcVal, srcStatus, el.Interface())
				case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
					i64 := int64(srcVal)
					if float64(i64) == srcVal {
						return int64AssignTo(i64, srcStatus, dst)
					}
				}
			}
			return errors.Errorf("cannot assign %v into %T", srcVal, dst)
		}
		return nil
	}
	if srcStatus == Null {
		return NullAssignTo(dst)
	}

	return errors.Errorf("cannot assign %v %v into %T", srcVal, srcStatus, dst)
}

func NullAssignTo(dst interface{}) error {
	dstPtr := reflect.ValueOf(dst)

	// AssignTo dst must always be a pointer
	if dstPtr.Kind() != reflect.Ptr {
		return errors.Errorf("cannot assign NULL to %T", dst)
	}

	dstVal := dstPtr.Elem()

	switch dstVal.Kind() {
	case reflect.Ptr, reflect.Slice, reflect.Map, reflect.Interface:
		dstVal.Set(reflect.Zero(dstVal.Type()))
		return nil
	}

	return errors.Errorf("cannot assign NULL to %T", dst)
}

var kindTypes map[reflect.Kind]reflect.Type

// GetAssignToDstType attempts to convert dst to something AssignTo can assign
// to. If dst is a pointer to pointer it allocates a value and returns the
// dereferences pointer. If dst is a named type such as *Foo where Foo is type
// Foo int16, it converts dst to *int16.
//
// GetAssignToDstType returns the converted dst and a bool representing if any
// change was made.
func GetAssignToDstType(dst interface{}) (interface{}, bool) {
	dstPtr := reflect.ValueOf(dst)

	// AssignTo dst must always be a pointer
	if dstPtr.Kind() != reflect.Ptr {
		return nil, false
	}

	dstVal := dstPtr.Elem()

	// if dst is a pointer to pointer, allocate space try again with the dereferenced pointer
	if dstVal.Kind() == reflect.Ptr {
		dstVal.Set(reflect.New(dstVal.Type().Elem()))
		return dstVal.Interface(), true
	}

	// if dst is pointer to a base type that has been renamed
	if baseValType, ok := kindTypes[dstVal.Kind()]; ok {
		nextDst := dstPtr.Convert(reflect.PtrTo(baseValType))
		return nextDst.Interface(), dstPtr.Type() != nextDst.Type()
	}

	if dstVal.Kind() == reflect.Slice {
		if baseElemType, ok := kindTypes[dstVal.Type().Elem().Kind()]; ok {
			baseSliceType := reflect.PtrTo(reflect.SliceOf(baseElemType))
			nextDst := dstPtr.Convert(baseSliceType)
			return nextDst.Interface(), dstPtr.Type() != nextDst.Type()
		}
	}

	return nil, false
}

func init() {
	kindTypes = map[reflect.Kind]reflect.Type{
		reflect.Bool:    reflect.TypeOf(false),
		reflect.Float32: reflect.TypeOf(float32(0)),
		reflect.Float64: reflect.TypeOf(float64(0)),
		reflect.Int:     reflect.TypeOf(int(0)),
		reflect.Int8:    reflect.TypeOf(int8(0)),
		reflect.Int16:   reflect.TypeOf(int16(0)),
		reflect.Int32:   reflect.TypeOf(int32(0)),
		reflect.Int64:   reflect.TypeOf(int64(0)),
		reflect.Uint:    reflect.TypeOf(uint(0)),
		reflect.Uint8:   reflect.TypeOf(uint8(0)),
		reflect.Uint16:  reflect.TypeOf(uint16(0)),
		reflect.Uint32:  reflect.TypeOf(uint32(0)),
		reflect.Uint64:  reflect.TypeOf(uint64(0)),
		reflect.String:  reflect.TypeOf(""),
	}
}
