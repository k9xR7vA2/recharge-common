package schema

import "github.com/k9xR7vA2/recharge-common/schema/types"

type SchemaField = types.SchemaField
type SchemaOption = types.SchemaOption
type RawField = types.RawField
type FieldType = types.FieldType
type SceneType = types.SceneType
type BusinessSchema = types.BusinessSchema

const (
	FieldTypeRadio       = types.FieldTypeRadio
	FieldTypeCheckbox    = types.FieldTypeCheckbox
	FieldTypeSelect      = types.FieldTypeSelect
	FieldTypeInput       = types.FieldTypeInput
	FieldTypeInputNumber = types.FieldTypeInputNumber
	SceneProduct         = types.SceneProduct
	SceneAccount         = types.SceneAccount
)

var Ptr = types.Ptr
