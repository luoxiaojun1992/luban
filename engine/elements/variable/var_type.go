package variable

import "fmt"

type BaseType string

const (
	TypeInt     BaseType = "int"
	TypeFloat   BaseType = "float64"
	TypeString  BaseType = "string"
	TypeBoolean BaseType = "bool"

	TypeRaw   BaseType = "raw"
	TypeArray BaseType = "array"
	TypeMap   BaseType = "map"
)

type ArrayType struct {
	EleType *VarType `json:"ele_type"`
	Size    int      `json:"size"`
}

type MapType struct {
	KeyType   *VarType `json:"key_type"`
	ValueType *VarType `json:"value_type"`
}

type VarType struct {
	BaseType  BaseType   `json:"base_type"`
	Raw       string     `json:"raw"`
	ArrayType *ArrayType `json:"array_type"`
	MapType   *MapType   `json:"map_type"`
}

func NewIntType() *VarType {
	return &VarType{
		BaseType: TypeInt,
	}
}

func NewFloatType() *VarType {
	return &VarType{
		BaseType: TypeFloat,
	}
}

func NewStringType() *VarType {
	return &VarType{
		BaseType: TypeString,
	}
}

func NewBooleanType() *VarType {
	return &VarType{
		BaseType: TypeBoolean,
	}
}

func NewRawType(raw string) *VarType {
	return &VarType{
		BaseType: TypeRaw,
		Raw:      raw,
	}
}

func NewArrayType(eleType *VarType, size int) *VarType {
	return &VarType{
		BaseType: TypeArray,
		ArrayType: &ArrayType{
			EleType: eleType,
			Size:    size,
		},
	}
}

func NewMapType(keyType, valueType *VarType) *VarType {
	return &VarType{
		BaseType: TypeMap,
		MapType: &MapType{
			KeyType:   keyType,
			ValueType: valueType,
		},
	}
}

func (vt *VarType) IsInt() bool {
	return vt.BaseType == TypeInt
}

func (vt *VarType) IsFloat() bool {
	return vt.BaseType == TypeFloat
}

func (vt *VarType) IsString() bool {
	return vt.BaseType == TypeString
}

func (vt *VarType) IsBoolean() bool {
	return vt.BaseType == TypeBoolean
}

func (vt *VarType) IsBasicType() bool {
	return vt.IsInt() || vt.IsFloat() || vt.IsString() || vt.IsBoolean()
}

func (vt *VarType) IsRaw() bool {
	return vt.BaseType == TypeRaw
}

func (vt *VarType) IsArray() bool {
	return vt.BaseType == TypeArray
}

func (vt *VarType) IsMap() bool {
	return vt.BaseType == TypeMap
}

func (vt *VarType) String() string {
	if vt.IsBasicType() {
		return string(vt.BaseType)
	}

	if vt.IsRaw() {
		return vt.Raw
	}

	if vt.IsArray() {
		size := ""
		if vt.ArrayType.Size > 0 {
			size = fmt.Sprintf("%d", vt.ArrayType.Size)
		}
		return fmt.Sprintf("[%s]%s", size, vt.ArrayType.EleType)
	}

	if vt.IsMap() {
		return fmt.Sprintf("map[%s]%s", vt.MapType.KeyType, vt.MapType.ValueType)
	}

	return ""
}

func (vt *VarType) Copy() *VarType {
	return &VarType{
		BaseType:  vt.BaseType,
		Raw:       vt.Raw,
		ArrayType: vt.ArrayType, // todo deep copy
		MapType:   vt.MapType,   // todo deep copy
	}
}
