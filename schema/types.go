package schema

import "github.com/k9xR7vA2/recharge-common/schema/types"

type SchemaField = types.SchemaField
type SchemaOption = types.SchemaOption
type RawField = types.RawField
type FieldType = types.FieldType
type BusinessSchema = types.BusinessSchema

const (
	FieldTypeRadio       = types.FieldTypeRadio
	FieldTypeCheckbox    = types.FieldTypeCheckbox
	FieldTypeSelect      = types.FieldTypeSelect
	FieldTypeInput       = types.FieldTypeInput
	FieldTypeInputNumber = types.FieldTypeInputNumber
)

var Ptr = types.Ptr
