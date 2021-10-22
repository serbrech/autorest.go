//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package complexgroup

import (
	"encoding/json"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"reflect"
	"time"
)

// ArrayGetEmptyOptions contains the optional parameters for the Array.GetEmpty method.
type ArrayGetEmptyOptions struct {
	// placeholder for future optional parameters
}

// ArrayGetNotProvidedOptions contains the optional parameters for the Array.GetNotProvided method.
type ArrayGetNotProvidedOptions struct {
	// placeholder for future optional parameters
}

// ArrayGetValidOptions contains the optional parameters for the Array.GetValid method.
type ArrayGetValidOptions struct {
	// placeholder for future optional parameters
}

// ArrayPutEmptyOptions contains the optional parameters for the Array.PutEmpty method.
type ArrayPutEmptyOptions struct {
	// placeholder for future optional parameters
}

// ArrayPutValidOptions contains the optional parameters for the Array.PutValid method.
type ArrayPutValidOptions struct {
	// placeholder for future optional parameters
}

type ArrayWrapper struct {
	Array []*string `json:"array,omitempty"`
}

// MarshalJSON implements the json.Marshaller interface for type ArrayWrapper.
func (a ArrayWrapper) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "array", a.Array)
	return json.Marshal(objectMap)
}

type Basic struct {
	Color *CMYKColors `json:"color,omitempty"`

	// Basic Id
	ID *int32 `json:"id,omitempty"`

	// Name property with a very long description that does not fit on a single line and a line break.
	Name *string `json:"name,omitempty"`
}

// BasicGetEmptyOptions contains the optional parameters for the Basic.GetEmpty method.
type BasicGetEmptyOptions struct {
	// placeholder for future optional parameters
}

// BasicGetInvalidOptions contains the optional parameters for the Basic.GetInvalid method.
type BasicGetInvalidOptions struct {
	// placeholder for future optional parameters
}

// BasicGetNotProvidedOptions contains the optional parameters for the Basic.GetNotProvided method.
type BasicGetNotProvidedOptions struct {
	// placeholder for future optional parameters
}

// BasicGetNullOptions contains the optional parameters for the Basic.GetNull method.
type BasicGetNullOptions struct {
	// placeholder for future optional parameters
}

// BasicGetValidOptions contains the optional parameters for the Basic.GetValid method.
type BasicGetValidOptions struct {
	// placeholder for future optional parameters
}

// BasicPutValidOptions contains the optional parameters for the Basic.PutValid method.
type BasicPutValidOptions struct {
	// placeholder for future optional parameters
}

type BooleanWrapper struct {
	FieldFalse *bool `json:"field_false,omitempty"`
	FieldTrue  *bool `json:"field_true,omitempty"`
}

type ByteWrapper struct {
	Field []byte `json:"field,omitempty"`
}

// MarshalJSON implements the json.Marshaller interface for type ByteWrapper.
func (b ByteWrapper) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populateByteArray(objectMap, "field", b.Field, runtime.Base64StdFormat)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ByteWrapper.
func (b *ByteWrapper) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "field":
			err = runtime.DecodeByteArray(string(val), &b.Field, runtime.Base64StdFormat)
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

type Cat struct {
	Pet
	Color *string `json:"color,omitempty"`
	Hates []*Dog  `json:"hates,omitempty"`
}

// MarshalJSON implements the json.Marshaller interface for type Cat.
func (c Cat) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	c.marshalInternal(objectMap)
	return json.Marshal(objectMap)
}

func (c Cat) marshalInternal(objectMap map[string]interface{}) {
	c.Pet.marshalInternal(objectMap)
	populate(objectMap, "color", c.Color)
	populate(objectMap, "hates", c.Hates)
}

type Cookiecuttershark struct {
	Shark
}

// MarshalJSON implements the json.Marshaller interface for type Cookiecuttershark.
func (c Cookiecuttershark) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	c.Shark.marshalInternal(objectMap, "cookiecuttershark")
	return json.Marshal(objectMap)
}

type DateWrapper struct {
	Field *time.Time `json:"field,omitempty"`
	Leap  *time.Time `json:"leap,omitempty"`
}

// MarshalJSON implements the json.Marshaller interface for type DateWrapper.
func (d DateWrapper) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populateDateType(objectMap, "field", d.Field)
	populateDateType(objectMap, "leap", d.Leap)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type DateWrapper.
func (d *DateWrapper) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "field":
			err = unpopulateDateType(val, &d.Field)
			delete(rawMsg, key)
		case "leap":
			err = unpopulateDateType(val, &d.Leap)
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

type DatetimeWrapper struct {
	Field *time.Time `json:"field,omitempty"`
	Now   *time.Time `json:"now,omitempty"`
}

// MarshalJSON implements the json.Marshaller interface for type DatetimeWrapper.
func (d DatetimeWrapper) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populateTimeRFC3339(objectMap, "field", d.Field)
	populateTimeRFC3339(objectMap, "now", d.Now)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type DatetimeWrapper.
func (d *DatetimeWrapper) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "field":
			err = unpopulateTimeRFC3339(val, &d.Field)
			delete(rawMsg, key)
		case "now":
			err = unpopulateTimeRFC3339(val, &d.Now)
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

type Datetimerfc1123Wrapper struct {
	Field *time.Time `json:"field,omitempty"`
	Now   *time.Time `json:"now,omitempty"`
}

// MarshalJSON implements the json.Marshaller interface for type Datetimerfc1123Wrapper.
func (d Datetimerfc1123Wrapper) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populateTimeRFC1123(objectMap, "field", d.Field)
	populateTimeRFC1123(objectMap, "now", d.Now)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type Datetimerfc1123Wrapper.
func (d *Datetimerfc1123Wrapper) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "field":
			err = unpopulateTimeRFC1123(val, &d.Field)
			delete(rawMsg, key)
		case "now":
			err = unpopulateTimeRFC1123(val, &d.Now)
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// DictionaryGetEmptyOptions contains the optional parameters for the Dictionary.GetEmpty method.
type DictionaryGetEmptyOptions struct {
	// placeholder for future optional parameters
}

// DictionaryGetNotProvidedOptions contains the optional parameters for the Dictionary.GetNotProvided method.
type DictionaryGetNotProvidedOptions struct {
	// placeholder for future optional parameters
}

// DictionaryGetNullOptions contains the optional parameters for the Dictionary.GetNull method.
type DictionaryGetNullOptions struct {
	// placeholder for future optional parameters
}

// DictionaryGetValidOptions contains the optional parameters for the Dictionary.GetValid method.
type DictionaryGetValidOptions struct {
	// placeholder for future optional parameters
}

// DictionaryPutEmptyOptions contains the optional parameters for the Dictionary.PutEmpty method.
type DictionaryPutEmptyOptions struct {
	// placeholder for future optional parameters
}

// DictionaryPutValidOptions contains the optional parameters for the Dictionary.PutValid method.
type DictionaryPutValidOptions struct {
	// placeholder for future optional parameters
}

type DictionaryWrapper struct {
	// Dictionary of
	DefaultProgram map[string]*string `json:"defaultProgram,omitempty"`
}

// MarshalJSON implements the json.Marshaller interface for type DictionaryWrapper.
func (d DictionaryWrapper) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "defaultProgram", d.DefaultProgram)
	return json.Marshal(objectMap)
}

type Dog struct {
	Pet
	Food *string `json:"food,omitempty"`
}

// MarshalJSON implements the json.Marshaller interface for type Dog.
func (d Dog) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	d.Pet.marshalInternal(objectMap)
	populate(objectMap, "food", d.Food)
	return json.Marshal(objectMap)
}

// DotFishClassification provides polymorphic access to related types.
// Call the interface's GetDotFish() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *DotFish, *DotSalmon
type DotFishClassification interface {
	// GetDotFish returns the DotFish content of the underlying type.
	GetDotFish() *DotFish
}

type DotFish struct {
	// REQUIRED
	FishType *string `json:"fish.type,omitempty"`
	Species  *string `json:"species,omitempty"`
}

// GetDotFish implements the DotFishClassification interface for type DotFish.
func (d *DotFish) GetDotFish() *DotFish { return d }

// UnmarshalJSON implements the json.Unmarshaller interface for type DotFish.
func (d *DotFish) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	return d.unmarshalInternal(rawMsg)
}

func (d DotFish) marshalInternal(objectMap map[string]interface{}, discValue string) {
	d.FishType = &discValue
	objectMap["fish.type"] = d.FishType
	populate(objectMap, "species", d.Species)
}

func (d *DotFish) unmarshalInternal(rawMsg map[string]json.RawMessage) error {
	for key, val := range rawMsg {
		var err error
		switch key {
		case "fish.type":
			err = unpopulate(val, &d.FishType)
			delete(rawMsg, key)
		case "species":
			err = unpopulate(val, &d.Species)
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

type DotFishMarket struct {
	Fishes       []DotFishClassification `json:"fishes,omitempty"`
	Salmons      []*DotSalmon            `json:"salmons,omitempty"`
	SampleFish   DotFishClassification   `json:"sampleFish,omitempty"`
	SampleSalmon *DotSalmon              `json:"sampleSalmon,omitempty"`
}

// MarshalJSON implements the json.Marshaller interface for type DotFishMarket.
func (d DotFishMarket) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "fishes", d.Fishes)
	populate(objectMap, "salmons", d.Salmons)
	populate(objectMap, "sampleFish", d.SampleFish)
	populate(objectMap, "sampleSalmon", d.SampleSalmon)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type DotFishMarket.
func (d *DotFishMarket) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "fishes":
			d.Fishes, err = unmarshalDotFishClassificationArray(val)
			delete(rawMsg, key)
		case "salmons":
			err = unpopulate(val, &d.Salmons)
			delete(rawMsg, key)
		case "sampleFish":
			d.SampleFish, err = unmarshalDotFishClassification(val)
			delete(rawMsg, key)
		case "sampleSalmon":
			err = unpopulate(val, &d.SampleSalmon)
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

type DotSalmon struct {
	DotFish
	Iswild   *bool   `json:"iswild,omitempty"`
	Location *string `json:"location,omitempty"`
}

// MarshalJSON implements the json.Marshaller interface for type DotSalmon.
func (d DotSalmon) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	d.DotFish.marshalInternal(objectMap, "DotSalmon")
	populate(objectMap, "iswild", d.Iswild)
	populate(objectMap, "location", d.Location)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type DotSalmon.
func (d *DotSalmon) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "iswild":
			err = unpopulate(val, &d.Iswild)
			delete(rawMsg, key)
		case "location":
			err = unpopulate(val, &d.Location)
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	if err := d.DotFish.unmarshalInternal(rawMsg); err != nil {
		return err
	}
	return nil
}

type DoubleWrapper struct {
	Field1                                                                          *float64 `json:"field1,omitempty"`
	Field56ZerosAfterTheDotAndNegativeZeroBeforeDotAndThisIsALongFieldNameOnPurpose *float64 `json:"field_56_zeros_after_the_dot_and_negative_zero_before_dot_and_this_is_a_long_field_name_on_purpose,omitempty"`
}

type DurationWrapper struct {
	Field *string `json:"field,omitempty"`
}

// Implements the error and azcore.HTTPResponse interfaces.
type Error struct {
	raw     string
	Message *string `json:"message,omitempty"`
	Status  *int32  `json:"status,omitempty"`
}

// Error implements the error interface for type Error.
// The contents of the error text are not contractual and subject to change.
func (e Error) Error() string {
	return e.raw
}

// FishClassification provides polymorphic access to related types.
// Call the interface's GetFish() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *Cookiecuttershark, *Fish, *Goblinshark, *Salmon, *Sawshark, *Shark, *SmartSalmon
type FishClassification interface {
	// GetFish returns the Fish content of the underlying type.
	GetFish() *Fish
}

type Fish struct {
	// REQUIRED
	Fishtype *string `json:"fishtype,omitempty"`

	// REQUIRED
	Length   *float32             `json:"length,omitempty"`
	Siblings []FishClassification `json:"siblings,omitempty"`
	Species  *string              `json:"species,omitempty"`
}

// GetFish implements the FishClassification interface for type Fish.
func (f *Fish) GetFish() *Fish { return f }

// UnmarshalJSON implements the json.Unmarshaller interface for type Fish.
func (f *Fish) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	return f.unmarshalInternal(rawMsg)
}

func (f Fish) marshalInternal(objectMap map[string]interface{}, discValue string) {
	f.Fishtype = &discValue
	objectMap["fishtype"] = f.Fishtype
	populate(objectMap, "length", f.Length)
	populate(objectMap, "siblings", f.Siblings)
	populate(objectMap, "species", f.Species)
}

func (f *Fish) unmarshalInternal(rawMsg map[string]json.RawMessage) error {
	for key, val := range rawMsg {
		var err error
		switch key {
		case "fishtype":
			err = unpopulate(val, &f.Fishtype)
			delete(rawMsg, key)
		case "length":
			err = unpopulate(val, &f.Length)
			delete(rawMsg, key)
		case "siblings":
			f.Siblings, err = unmarshalFishClassificationArray(val)
			delete(rawMsg, key)
		case "species":
			err = unpopulate(val, &f.Species)
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// FlattencomplexGetValidOptions contains the optional parameters for the Flattencomplex.GetValid method.
type FlattencomplexGetValidOptions struct {
	// placeholder for future optional parameters
}

type FloatWrapper struct {
	Field1 *float32 `json:"field1,omitempty"`
	Field2 *float32 `json:"field2,omitempty"`
}

type Goblinshark struct {
	Shark
	// Colors possible
	Color   *GoblinSharkColor `json:"color,omitempty"`
	Jawsize *int32            `json:"jawsize,omitempty"`
}

// MarshalJSON implements the json.Marshaller interface for type Goblinshark.
func (g Goblinshark) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	g.Shark.marshalInternal(objectMap, "goblin")
	populate(objectMap, "color", g.Color)
	populate(objectMap, "jawsize", g.Jawsize)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type Goblinshark.
func (g *Goblinshark) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "color":
			err = unpopulate(val, &g.Color)
			delete(rawMsg, key)
		case "jawsize":
			err = unpopulate(val, &g.Jawsize)
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	if err := g.Shark.unmarshalInternal(rawMsg); err != nil {
		return err
	}
	return nil
}

// InheritanceGetValidOptions contains the optional parameters for the Inheritance.GetValid method.
type InheritanceGetValidOptions struct {
	// placeholder for future optional parameters
}

// InheritancePutValidOptions contains the optional parameters for the Inheritance.PutValid method.
type InheritancePutValidOptions struct {
	// placeholder for future optional parameters
}

type IntWrapper struct {
	Field1 *int32 `json:"field1,omitempty"`
	Field2 *int32 `json:"field2,omitempty"`
}

type LongWrapper struct {
	Field1 *int64 `json:"field1,omitempty"`
	Field2 *int64 `json:"field2,omitempty"`
}

type MyBaseHelperType struct {
	PropBH1 *string `json:"propBH1,omitempty"`
}

// MyBaseTypeClassification provides polymorphic access to related types.
// Call the interface's GetMyBaseType() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *MyBaseType, *MyDerivedType
type MyBaseTypeClassification interface {
	// GetMyBaseType returns the MyBaseType content of the underlying type.
	GetMyBaseType() *MyBaseType
}

type MyBaseType struct {
	// REQUIRED
	Kind   *MyKind           `json:"kind,omitempty"`
	Helper *MyBaseHelperType `json:"helper,omitempty"`
	PropB1 *string           `json:"propB1,omitempty"`
}

// GetMyBaseType implements the MyBaseTypeClassification interface for type MyBaseType.
func (m *MyBaseType) GetMyBaseType() *MyBaseType { return m }

// UnmarshalJSON implements the json.Unmarshaller interface for type MyBaseType.
func (m *MyBaseType) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	return m.unmarshalInternal(rawMsg)
}

func (m MyBaseType) marshalInternal(objectMap map[string]interface{}, discValue MyKind) {
	populate(objectMap, "helper", m.Helper)
	m.Kind = &discValue
	objectMap["kind"] = m.Kind
	populate(objectMap, "propB1", m.PropB1)
}

func (m *MyBaseType) unmarshalInternal(rawMsg map[string]json.RawMessage) error {
	for key, val := range rawMsg {
		var err error
		switch key {
		case "helper":
			err = unpopulate(val, &m.Helper)
			delete(rawMsg, key)
		case "kind":
			err = unpopulate(val, &m.Kind)
			delete(rawMsg, key)
		case "propB1":
			err = unpopulate(val, &m.PropB1)
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

type MyDerivedType struct {
	MyBaseType
	PropD1 *string `json:"propD1,omitempty"`
}

// MarshalJSON implements the json.Marshaller interface for type MyDerivedType.
func (m MyDerivedType) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	m.MyBaseType.marshalInternal(objectMap, MyKindKind1)
	populate(objectMap, "propD1", m.PropD1)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type MyDerivedType.
func (m *MyDerivedType) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "propD1":
			err = unpopulate(val, &m.PropD1)
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	if err := m.MyBaseType.unmarshalInternal(rawMsg); err != nil {
		return err
	}
	return nil
}

type Pet struct {
	ID   *int32  `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

// MarshalJSON implements the json.Marshaller interface for type Pet.
func (p Pet) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	p.marshalInternal(objectMap)
	return json.Marshal(objectMap)
}

func (p Pet) marshalInternal(objectMap map[string]interface{}) {
	populate(objectMap, "id", p.ID)
	populate(objectMap, "name", p.Name)
}

// PolymorphicrecursiveGetValidOptions contains the optional parameters for the Polymorphicrecursive.GetValid method.
type PolymorphicrecursiveGetValidOptions struct {
	// placeholder for future optional parameters
}

// PolymorphicrecursivePutValidOptions contains the optional parameters for the Polymorphicrecursive.PutValid method.
type PolymorphicrecursivePutValidOptions struct {
	// placeholder for future optional parameters
}

// PolymorphismGetComplicatedOptions contains the optional parameters for the Polymorphism.GetComplicated method.
type PolymorphismGetComplicatedOptions struct {
	// placeholder for future optional parameters
}

// PolymorphismGetComposedWithDiscriminatorOptions contains the optional parameters for the Polymorphism.GetComposedWithDiscriminator method.
type PolymorphismGetComposedWithDiscriminatorOptions struct {
	// placeholder for future optional parameters
}

// PolymorphismGetComposedWithoutDiscriminatorOptions contains the optional parameters for the Polymorphism.GetComposedWithoutDiscriminator method.
type PolymorphismGetComposedWithoutDiscriminatorOptions struct {
	// placeholder for future optional parameters
}

// PolymorphismGetDotSyntaxOptions contains the optional parameters for the Polymorphism.GetDotSyntax method.
type PolymorphismGetDotSyntaxOptions struct {
	// placeholder for future optional parameters
}

// PolymorphismGetValidOptions contains the optional parameters for the Polymorphism.GetValid method.
type PolymorphismGetValidOptions struct {
	// placeholder for future optional parameters
}

// PolymorphismPutComplicatedOptions contains the optional parameters for the Polymorphism.PutComplicated method.
type PolymorphismPutComplicatedOptions struct {
	// placeholder for future optional parameters
}

// PolymorphismPutMissingDiscriminatorOptions contains the optional parameters for the Polymorphism.PutMissingDiscriminator method.
type PolymorphismPutMissingDiscriminatorOptions struct {
	// placeholder for future optional parameters
}

// PolymorphismPutValidMissingRequiredOptions contains the optional parameters for the Polymorphism.PutValidMissingRequired method.
type PolymorphismPutValidMissingRequiredOptions struct {
	// placeholder for future optional parameters
}

// PolymorphismPutValidOptions contains the optional parameters for the Polymorphism.PutValid method.
type PolymorphismPutValidOptions struct {
	// placeholder for future optional parameters
}

// PrimitiveGetBoolOptions contains the optional parameters for the Primitive.GetBool method.
type PrimitiveGetBoolOptions struct {
	// placeholder for future optional parameters
}

// PrimitiveGetByteOptions contains the optional parameters for the Primitive.GetByte method.
type PrimitiveGetByteOptions struct {
	// placeholder for future optional parameters
}

// PrimitiveGetDateOptions contains the optional parameters for the Primitive.GetDate method.
type PrimitiveGetDateOptions struct {
	// placeholder for future optional parameters
}

// PrimitiveGetDateTimeOptions contains the optional parameters for the Primitive.GetDateTime method.
type PrimitiveGetDateTimeOptions struct {
	// placeholder for future optional parameters
}

// PrimitiveGetDateTimeRFC1123Options contains the optional parameters for the Primitive.GetDateTimeRFC1123 method.
type PrimitiveGetDateTimeRFC1123Options struct {
	// placeholder for future optional parameters
}

// PrimitiveGetDoubleOptions contains the optional parameters for the Primitive.GetDouble method.
type PrimitiveGetDoubleOptions struct {
	// placeholder for future optional parameters
}

// PrimitiveGetDurationOptions contains the optional parameters for the Primitive.GetDuration method.
type PrimitiveGetDurationOptions struct {
	// placeholder for future optional parameters
}

// PrimitiveGetFloatOptions contains the optional parameters for the Primitive.GetFloat method.
type PrimitiveGetFloatOptions struct {
	// placeholder for future optional parameters
}

// PrimitiveGetIntOptions contains the optional parameters for the Primitive.GetInt method.
type PrimitiveGetIntOptions struct {
	// placeholder for future optional parameters
}

// PrimitiveGetLongOptions contains the optional parameters for the Primitive.GetLong method.
type PrimitiveGetLongOptions struct {
	// placeholder for future optional parameters
}

// PrimitiveGetStringOptions contains the optional parameters for the Primitive.GetString method.
type PrimitiveGetStringOptions struct {
	// placeholder for future optional parameters
}

// PrimitivePutBoolOptions contains the optional parameters for the Primitive.PutBool method.
type PrimitivePutBoolOptions struct {
	// placeholder for future optional parameters
}

// PrimitivePutByteOptions contains the optional parameters for the Primitive.PutByte method.
type PrimitivePutByteOptions struct {
	// placeholder for future optional parameters
}

// PrimitivePutDateOptions contains the optional parameters for the Primitive.PutDate method.
type PrimitivePutDateOptions struct {
	// placeholder for future optional parameters
}

// PrimitivePutDateTimeOptions contains the optional parameters for the Primitive.PutDateTime method.
type PrimitivePutDateTimeOptions struct {
	// placeholder for future optional parameters
}

// PrimitivePutDateTimeRFC1123Options contains the optional parameters for the Primitive.PutDateTimeRFC1123 method.
type PrimitivePutDateTimeRFC1123Options struct {
	// placeholder for future optional parameters
}

// PrimitivePutDoubleOptions contains the optional parameters for the Primitive.PutDouble method.
type PrimitivePutDoubleOptions struct {
	// placeholder for future optional parameters
}

// PrimitivePutDurationOptions contains the optional parameters for the Primitive.PutDuration method.
type PrimitivePutDurationOptions struct {
	// placeholder for future optional parameters
}

// PrimitivePutFloatOptions contains the optional parameters for the Primitive.PutFloat method.
type PrimitivePutFloatOptions struct {
	// placeholder for future optional parameters
}

// PrimitivePutIntOptions contains the optional parameters for the Primitive.PutInt method.
type PrimitivePutIntOptions struct {
	// placeholder for future optional parameters
}

// PrimitivePutLongOptions contains the optional parameters for the Primitive.PutLong method.
type PrimitivePutLongOptions struct {
	// placeholder for future optional parameters
}

// PrimitivePutStringOptions contains the optional parameters for the Primitive.PutString method.
type PrimitivePutStringOptions struct {
	// placeholder for future optional parameters
}

type ReadonlyObj struct {
	Size *int32 `json:"size,omitempty"`

	// READ-ONLY
	ID *string `json:"id,omitempty" azure:"ro"`
}

// ReadonlypropertyGetValidOptions contains the optional parameters for the Readonlyproperty.GetValid method.
type ReadonlypropertyGetValidOptions struct {
	// placeholder for future optional parameters
}

// ReadonlypropertyPutValidOptions contains the optional parameters for the Readonlyproperty.PutValid method.
type ReadonlypropertyPutValidOptions struct {
	// placeholder for future optional parameters
}

// SalmonClassification provides polymorphic access to related types.
// Call the interface's GetSalmon() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *Salmon, *SmartSalmon
type SalmonClassification interface {
	FishClassification
	// GetSalmon returns the Salmon content of the underlying type.
	GetSalmon() *Salmon
}

type Salmon struct {
	Fish
	Iswild   *bool   `json:"iswild,omitempty"`
	Location *string `json:"location,omitempty"`
}

// GetSalmon implements the SalmonClassification interface for type Salmon.
func (s *Salmon) GetSalmon() *Salmon { return s }

// MarshalJSON implements the json.Marshaller interface for type Salmon.
func (s Salmon) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	s.marshalInternal(objectMap, "salmon")
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type Salmon.
func (s *Salmon) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	return s.unmarshalInternal(rawMsg)
}

func (s Salmon) marshalInternal(objectMap map[string]interface{}, discValue string) {
	s.Fish.marshalInternal(objectMap, discValue)
	populate(objectMap, "iswild", s.Iswild)
	populate(objectMap, "location", s.Location)
}

func (s *Salmon) unmarshalInternal(rawMsg map[string]json.RawMessage) error {
	for key, val := range rawMsg {
		var err error
		switch key {
		case "iswild":
			err = unpopulate(val, &s.Iswild)
			delete(rawMsg, key)
		case "location":
			err = unpopulate(val, &s.Location)
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	if err := s.Fish.unmarshalInternal(rawMsg); err != nil {
		return err
	}
	return nil
}

type Sawshark struct {
	Shark
	Picture []byte `json:"picture,omitempty"`
}

// MarshalJSON implements the json.Marshaller interface for type Sawshark.
func (s Sawshark) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	s.Shark.marshalInternal(objectMap, "sawshark")
	populateByteArray(objectMap, "picture", s.Picture, runtime.Base64StdFormat)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type Sawshark.
func (s *Sawshark) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "picture":
			err = runtime.DecodeByteArray(string(val), &s.Picture, runtime.Base64StdFormat)
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	if err := s.Shark.unmarshalInternal(rawMsg); err != nil {
		return err
	}
	return nil
}

// SharkClassification provides polymorphic access to related types.
// Call the interface's GetShark() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *Cookiecuttershark, *Goblinshark, *Sawshark, *Shark
type SharkClassification interface {
	FishClassification
	// GetShark returns the Shark content of the underlying type.
	GetShark() *Shark
}

type Shark struct {
	Fish
	// REQUIRED
	Birthday *time.Time `json:"birthday,omitempty"`
	Age      *int32     `json:"age,omitempty"`
}

// GetShark implements the SharkClassification interface for type Shark.
func (s *Shark) GetShark() *Shark { return s }

// MarshalJSON implements the json.Marshaller interface for type Shark.
func (s Shark) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	s.marshalInternal(objectMap, "shark")
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type Shark.
func (s *Shark) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	return s.unmarshalInternal(rawMsg)
}

func (s Shark) marshalInternal(objectMap map[string]interface{}, discValue string) {
	s.Fish.marshalInternal(objectMap, discValue)
	populate(objectMap, "age", s.Age)
	populateTimeRFC3339(objectMap, "birthday", s.Birthday)
}

func (s *Shark) unmarshalInternal(rawMsg map[string]json.RawMessage) error {
	for key, val := range rawMsg {
		var err error
		switch key {
		case "age":
			err = unpopulate(val, &s.Age)
			delete(rawMsg, key)
		case "birthday":
			err = unpopulateTimeRFC3339(val, &s.Birthday)
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	if err := s.Fish.unmarshalInternal(rawMsg); err != nil {
		return err
	}
	return nil
}

type Siamese struct {
	Cat
	Breed *string `json:"breed,omitempty"`
}

// MarshalJSON implements the json.Marshaller interface for type Siamese.
func (s Siamese) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	s.Cat.marshalInternal(objectMap)
	populate(objectMap, "breed", s.Breed)
	return json.Marshal(objectMap)
}

type SmartSalmon struct {
	Salmon
	// OPTIONAL; Contains additional key/value pairs not defined in the schema.
	AdditionalProperties map[string]interface{}
	CollegeDegree        *string `json:"college_degree,omitempty"`
}

// MarshalJSON implements the json.Marshaller interface for type SmartSalmon.
func (s SmartSalmon) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	s.Salmon.marshalInternal(objectMap, "smart_salmon")
	populate(objectMap, "college_degree", s.CollegeDegree)
	if s.AdditionalProperties != nil {
		for key, val := range s.AdditionalProperties {
			objectMap[key] = val
		}
	}
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type SmartSalmon.
func (s *SmartSalmon) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "college_degree":
			err = unpopulate(val, &s.CollegeDegree)
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	if err := s.Salmon.unmarshalInternal(rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		if s.AdditionalProperties == nil {
			s.AdditionalProperties = map[string]interface{}{}
		}
		if val != nil {
			var aux interface{}
			err = json.Unmarshal(val, &aux)
			s.AdditionalProperties[key] = aux
		}
		delete(rawMsg, key)
		if err != nil {
			return err
		}
	}
	return nil
}

type StringWrapper struct {
	Empty *string `json:"empty,omitempty"`
	Field *string `json:"field,omitempty"`
	Null  *string `json:"null,omitempty"`
}

func populate(m map[string]interface{}, k string, v interface{}) {
	if v == nil {
		return
	} else if azcore.IsNullValue(v) {
		m[k] = nil
	} else if !reflect.ValueOf(v).IsNil() {
		m[k] = v
	}
}

func populateByteArray(m map[string]interface{}, k string, b []byte, f runtime.Base64Encoding) {
	if azcore.IsNullValue(b) {
		m[k] = nil
	} else if len(b) == 0 {
		return
	} else {
		m[k] = runtime.EncodeByteArray(b, f)
	}
}

func unpopulate(data json.RawMessage, v interface{}) error {
	if data == nil {
		return nil
	}
	return json.Unmarshal(data, v)
}
