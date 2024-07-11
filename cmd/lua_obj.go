package main

type LuaObjType byte

const (
	ObjNil LuaObjType = iota
	ObjBool
	ObjNumber
	ObjString
)

// Base representation of a lua object, all value are a type of this object
type LuaObj struct {
	Kind  LuaObjType
	Value any
}

func NewNilValue() *LuaObj {
	return &LuaObj{Kind: ObjNil, Value: nil}
}

func NewBoolValue(b bool) *LuaObj {
	return &LuaObj{Kind: ObjBool, Value: b}
}
