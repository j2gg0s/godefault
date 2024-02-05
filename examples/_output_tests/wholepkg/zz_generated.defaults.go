//go:build !ignore_autogenerated
// +build !ignore_autogenerated

package wholepkg

import (
	"fmt"
	"reflect"
)

// RegisterDefaults adds defaulters functions to the given map.
// Public to allow building arbitrary maps.
// All generated defaulters are covering - they call all nested defaulters.
func RegisterDefaults(m map[reflect.Type]func(interface{})) error {
	m[reflect.TypeOf(&Struct_Everything{})] = func(obj interface{}) { SetObjectDefaults_Struct_Everything(obj.(*Struct_Everything)) }
	m[reflect.TypeOf(&Struct_Pointer{})] = func(obj interface{}) { SetObjectDefaults_Struct_Pointer(obj.(*Struct_Pointer)) }
	m[reflect.TypeOf(&Struct_Primitives{})] = func(obj interface{}) { SetObjectDefaults_Struct_Primitives(obj.(*Struct_Primitives)) }
	m[reflect.TypeOf(&Struct_Slices{})] = func(obj interface{}) { SetObjectDefaults_Struct_Slices(obj.(*Struct_Slices)) }
	m[reflect.TypeOf(&Struct_Struct_Primitives{})] = func(obj interface{}) { SetObjectDefaults_Struct_Struct_Primitives(obj.(*Struct_Struct_Primitives)) }
	return nil
}

var m map[reflect.Type]func(interface{})

func init() {
	m = map[reflect.Type]func(interface{}){}
	RegisterDefaults(m)
}

// Default set default value for input object.
func Default(obj interface{}) error {
	fn, ok := m[reflect.TypeOf(obj)]
	if !ok {
		return fmt.Errorf("unknown type: %T", obj)
	}
	fn(obj)
	return nil
}

func SetObjectDefaults_Struct_Everything(in *Struct_Everything) {
	SetObjectDefaults_Struct_Pointer(&in.PointerStructField)
	SetObjectDefaults_Struct_Slices(&in.SlicesStructField)
}

func SetObjectDefaults_Struct_Pointer(in *Struct_Pointer) {
	SetObjectDefaults_Struct_Primitives(&in.PointerStructPrimitivesField)
	if in.PointerPointerStructPrimitivesField != nil {
		SetObjectDefaults_Struct_Primitives(in.PointerPointerStructPrimitivesField)
	}
	SetObjectDefaults_Struct_Struct_Primitives(&in.PointerStructStructPrimitives)
	if in.PointerPointerStructStructPrimitives != nil {
		SetObjectDefaults_Struct_Struct_Primitives(in.PointerPointerStructStructPrimitives)
	}
}

func SetObjectDefaults_Struct_Primitives(in *Struct_Primitives) {
	SetDefaults_Struct_Primitives(in)
}

func SetObjectDefaults_Struct_Slices(in *Struct_Slices) {
	for i := range in.SliceStructPrimitivesField {
		a := &in.SliceStructPrimitivesField[i]
		SetObjectDefaults_Struct_Primitives(a)
	}
	for i := range in.SlicePointerStructPrimitivesField {
		a := in.SlicePointerStructPrimitivesField[i]
		if a != nil {
			SetObjectDefaults_Struct_Primitives(a)
		}
	}
	for i := range in.SliceStructStructPrimitives {
		a := &in.SliceStructStructPrimitives[i]
		SetObjectDefaults_Struct_Struct_Primitives(a)
	}
	for i := range in.SlicePointerStructStructPrimitives {
		a := in.SlicePointerStructStructPrimitives[i]
		if a != nil {
			SetObjectDefaults_Struct_Struct_Primitives(a)
		}
	}
}

func SetObjectDefaults_Struct_Struct_Primitives(in *Struct_Struct_Primitives) {
	SetObjectDefaults_Struct_Primitives(&in.StructField)
}