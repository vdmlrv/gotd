// Code generated by gotdgen, DO NOT EDIT.

package tdapi

import (
	"context"
	"errors"
	"fmt"
	"sort"
	"strings"

	"go.uber.org/multierr"

	"github.com/gotd/td/bin"
	"github.com/gotd/td/jsontd"
	"github.com/gotd/td/tdp"
	"github.com/gotd/td/tgerr"
)

// No-op definition for keeping imports.
var (
	_ = bin.Buffer{}
	_ = context.Background()
	_ = fmt.Stringer(nil)
	_ = strings.Builder{}
	_ = errors.Is
	_ = multierr.AppendInto
	_ = sort.Ints
	_ = tdp.Format
	_ = tgerr.Error{}
	_ = jsontd.Encoder{}
)

// StatisticalGraphData represents TL type `statisticalGraphData#89732e2c`.
type StatisticalGraphData struct {
	// Graph data in JSON format
	JSONData string
	// If non-empty, a token which can be used to receive a zoomed in graph
	ZoomToken string
}

// StatisticalGraphDataTypeID is TL type id of StatisticalGraphData.
const StatisticalGraphDataTypeID = 0x89732e2c

// construct implements constructor of StatisticalGraphClass.
func (s StatisticalGraphData) construct() StatisticalGraphClass { return &s }

// Ensuring interfaces in compile-time for StatisticalGraphData.
var (
	_ bin.Encoder     = &StatisticalGraphData{}
	_ bin.Decoder     = &StatisticalGraphData{}
	_ bin.BareEncoder = &StatisticalGraphData{}
	_ bin.BareDecoder = &StatisticalGraphData{}

	_ StatisticalGraphClass = &StatisticalGraphData{}
)

func (s *StatisticalGraphData) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.JSONData == "") {
		return false
	}
	if !(s.ZoomToken == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *StatisticalGraphData) String() string {
	if s == nil {
		return "StatisticalGraphData(nil)"
	}
	type Alias StatisticalGraphData
	return fmt.Sprintf("StatisticalGraphData%+v", Alias(*s))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*StatisticalGraphData) TypeID() uint32 {
	return StatisticalGraphDataTypeID
}

// TypeName returns name of type in TL schema.
func (*StatisticalGraphData) TypeName() string {
	return "statisticalGraphData"
}

// TypeInfo returns info about TL type.
func (s *StatisticalGraphData) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "statisticalGraphData",
		ID:   StatisticalGraphDataTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "JSONData",
			SchemaName: "json_data",
		},
		{
			Name:       "ZoomToken",
			SchemaName: "zoom_token",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (s *StatisticalGraphData) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode statisticalGraphData#89732e2c as nil")
	}
	b.PutID(StatisticalGraphDataTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *StatisticalGraphData) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode statisticalGraphData#89732e2c as nil")
	}
	b.PutString(s.JSONData)
	b.PutString(s.ZoomToken)
	return nil
}

// Decode implements bin.Decoder.
func (s *StatisticalGraphData) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode statisticalGraphData#89732e2c to nil")
	}
	if err := b.ConsumeID(StatisticalGraphDataTypeID); err != nil {
		return fmt.Errorf("unable to decode statisticalGraphData#89732e2c: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *StatisticalGraphData) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode statisticalGraphData#89732e2c to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode statisticalGraphData#89732e2c: field json_data: %w", err)
		}
		s.JSONData = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode statisticalGraphData#89732e2c: field zoom_token: %w", err)
		}
		s.ZoomToken = value
	}
	return nil
}

// EncodeTDLibJSON implements jsontd.TDLibEncoder.
func (s *StatisticalGraphData) EncodeTDLibJSON(b jsontd.Encoder) error {
	if s == nil {
		return fmt.Errorf("can't encode statisticalGraphData#89732e2c as nil")
	}
	b.ObjStart()
	b.PutID("statisticalGraphData")
	b.FieldStart("json_data")
	b.PutString(s.JSONData)
	b.FieldStart("zoom_token")
	b.PutString(s.ZoomToken)
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements jsontd.TDLibDecoder.
func (s *StatisticalGraphData) DecodeTDLibJSON(b jsontd.Decoder) error {
	if s == nil {
		return fmt.Errorf("can't decode statisticalGraphData#89732e2c to nil")
	}

	return b.Obj(func(b jsontd.Decoder, key []byte) error {
		switch string(key) {
		case jsontd.TypeField:
			if err := b.ConsumeID("statisticalGraphData"); err != nil {
				return fmt.Errorf("unable to decode statisticalGraphData#89732e2c: %w", err)
			}
		case "json_data":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode statisticalGraphData#89732e2c: field json_data: %w", err)
			}
			s.JSONData = value
		case "zoom_token":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode statisticalGraphData#89732e2c: field zoom_token: %w", err)
			}
			s.ZoomToken = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetJSONData returns value of JSONData field.
func (s *StatisticalGraphData) GetJSONData() (value string) {
	return s.JSONData
}

// GetZoomToken returns value of ZoomToken field.
func (s *StatisticalGraphData) GetZoomToken() (value string) {
	return s.ZoomToken
}

// StatisticalGraphAsync represents TL type `statisticalGraphAsync#19fb2b9f`.
type StatisticalGraphAsync struct {
	// The token to use for data loading
	Token string
}

// StatisticalGraphAsyncTypeID is TL type id of StatisticalGraphAsync.
const StatisticalGraphAsyncTypeID = 0x19fb2b9f

// construct implements constructor of StatisticalGraphClass.
func (s StatisticalGraphAsync) construct() StatisticalGraphClass { return &s }

// Ensuring interfaces in compile-time for StatisticalGraphAsync.
var (
	_ bin.Encoder     = &StatisticalGraphAsync{}
	_ bin.Decoder     = &StatisticalGraphAsync{}
	_ bin.BareEncoder = &StatisticalGraphAsync{}
	_ bin.BareDecoder = &StatisticalGraphAsync{}

	_ StatisticalGraphClass = &StatisticalGraphAsync{}
)

func (s *StatisticalGraphAsync) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.Token == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *StatisticalGraphAsync) String() string {
	if s == nil {
		return "StatisticalGraphAsync(nil)"
	}
	type Alias StatisticalGraphAsync
	return fmt.Sprintf("StatisticalGraphAsync%+v", Alias(*s))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*StatisticalGraphAsync) TypeID() uint32 {
	return StatisticalGraphAsyncTypeID
}

// TypeName returns name of type in TL schema.
func (*StatisticalGraphAsync) TypeName() string {
	return "statisticalGraphAsync"
}

// TypeInfo returns info about TL type.
func (s *StatisticalGraphAsync) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "statisticalGraphAsync",
		ID:   StatisticalGraphAsyncTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Token",
			SchemaName: "token",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (s *StatisticalGraphAsync) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode statisticalGraphAsync#19fb2b9f as nil")
	}
	b.PutID(StatisticalGraphAsyncTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *StatisticalGraphAsync) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode statisticalGraphAsync#19fb2b9f as nil")
	}
	b.PutString(s.Token)
	return nil
}

// Decode implements bin.Decoder.
func (s *StatisticalGraphAsync) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode statisticalGraphAsync#19fb2b9f to nil")
	}
	if err := b.ConsumeID(StatisticalGraphAsyncTypeID); err != nil {
		return fmt.Errorf("unable to decode statisticalGraphAsync#19fb2b9f: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *StatisticalGraphAsync) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode statisticalGraphAsync#19fb2b9f to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode statisticalGraphAsync#19fb2b9f: field token: %w", err)
		}
		s.Token = value
	}
	return nil
}

// EncodeTDLibJSON implements jsontd.TDLibEncoder.
func (s *StatisticalGraphAsync) EncodeTDLibJSON(b jsontd.Encoder) error {
	if s == nil {
		return fmt.Errorf("can't encode statisticalGraphAsync#19fb2b9f as nil")
	}
	b.ObjStart()
	b.PutID("statisticalGraphAsync")
	b.FieldStart("token")
	b.PutString(s.Token)
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements jsontd.TDLibDecoder.
func (s *StatisticalGraphAsync) DecodeTDLibJSON(b jsontd.Decoder) error {
	if s == nil {
		return fmt.Errorf("can't decode statisticalGraphAsync#19fb2b9f to nil")
	}

	return b.Obj(func(b jsontd.Decoder, key []byte) error {
		switch string(key) {
		case jsontd.TypeField:
			if err := b.ConsumeID("statisticalGraphAsync"); err != nil {
				return fmt.Errorf("unable to decode statisticalGraphAsync#19fb2b9f: %w", err)
			}
		case "token":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode statisticalGraphAsync#19fb2b9f: field token: %w", err)
			}
			s.Token = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetToken returns value of Token field.
func (s *StatisticalGraphAsync) GetToken() (value string) {
	return s.Token
}

// StatisticalGraphError represents TL type `statisticalGraphError#c3fda052`.
type StatisticalGraphError struct {
	// The error message
	ErrorMessage string
}

// StatisticalGraphErrorTypeID is TL type id of StatisticalGraphError.
const StatisticalGraphErrorTypeID = 0xc3fda052

// construct implements constructor of StatisticalGraphClass.
func (s StatisticalGraphError) construct() StatisticalGraphClass { return &s }

// Ensuring interfaces in compile-time for StatisticalGraphError.
var (
	_ bin.Encoder     = &StatisticalGraphError{}
	_ bin.Decoder     = &StatisticalGraphError{}
	_ bin.BareEncoder = &StatisticalGraphError{}
	_ bin.BareDecoder = &StatisticalGraphError{}

	_ StatisticalGraphClass = &StatisticalGraphError{}
)

func (s *StatisticalGraphError) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.ErrorMessage == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *StatisticalGraphError) String() string {
	if s == nil {
		return "StatisticalGraphError(nil)"
	}
	type Alias StatisticalGraphError
	return fmt.Sprintf("StatisticalGraphError%+v", Alias(*s))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*StatisticalGraphError) TypeID() uint32 {
	return StatisticalGraphErrorTypeID
}

// TypeName returns name of type in TL schema.
func (*StatisticalGraphError) TypeName() string {
	return "statisticalGraphError"
}

// TypeInfo returns info about TL type.
func (s *StatisticalGraphError) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "statisticalGraphError",
		ID:   StatisticalGraphErrorTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ErrorMessage",
			SchemaName: "error_message",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (s *StatisticalGraphError) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode statisticalGraphError#c3fda052 as nil")
	}
	b.PutID(StatisticalGraphErrorTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *StatisticalGraphError) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode statisticalGraphError#c3fda052 as nil")
	}
	b.PutString(s.ErrorMessage)
	return nil
}

// Decode implements bin.Decoder.
func (s *StatisticalGraphError) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode statisticalGraphError#c3fda052 to nil")
	}
	if err := b.ConsumeID(StatisticalGraphErrorTypeID); err != nil {
		return fmt.Errorf("unable to decode statisticalGraphError#c3fda052: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *StatisticalGraphError) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode statisticalGraphError#c3fda052 to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode statisticalGraphError#c3fda052: field error_message: %w", err)
		}
		s.ErrorMessage = value
	}
	return nil
}

// EncodeTDLibJSON implements jsontd.TDLibEncoder.
func (s *StatisticalGraphError) EncodeTDLibJSON(b jsontd.Encoder) error {
	if s == nil {
		return fmt.Errorf("can't encode statisticalGraphError#c3fda052 as nil")
	}
	b.ObjStart()
	b.PutID("statisticalGraphError")
	b.FieldStart("error_message")
	b.PutString(s.ErrorMessage)
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements jsontd.TDLibDecoder.
func (s *StatisticalGraphError) DecodeTDLibJSON(b jsontd.Decoder) error {
	if s == nil {
		return fmt.Errorf("can't decode statisticalGraphError#c3fda052 to nil")
	}

	return b.Obj(func(b jsontd.Decoder, key []byte) error {
		switch string(key) {
		case jsontd.TypeField:
			if err := b.ConsumeID("statisticalGraphError"); err != nil {
				return fmt.Errorf("unable to decode statisticalGraphError#c3fda052: %w", err)
			}
		case "error_message":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode statisticalGraphError#c3fda052: field error_message: %w", err)
			}
			s.ErrorMessage = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetErrorMessage returns value of ErrorMessage field.
func (s *StatisticalGraphError) GetErrorMessage() (value string) {
	return s.ErrorMessage
}

// StatisticalGraphClass represents StatisticalGraph generic type.
//
// Example:
//  g, err := tdapi.DecodeStatisticalGraph(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *tdapi.StatisticalGraphData: // statisticalGraphData#89732e2c
//  case *tdapi.StatisticalGraphAsync: // statisticalGraphAsync#19fb2b9f
//  case *tdapi.StatisticalGraphError: // statisticalGraphError#c3fda052
//  default: panic(v)
//  }
type StatisticalGraphClass interface {
	bin.Encoder
	bin.Decoder
	bin.BareEncoder
	bin.BareDecoder
	construct() StatisticalGraphClass

	// TypeID returns type id in TL schema.
	//
	// See https://core.telegram.org/mtproto/TL-tl#remarks.
	TypeID() uint32
	// TypeName returns name of type in TL schema.
	TypeName() string
	// String implements fmt.Stringer.
	String() string
	// Zero returns true if current object has a zero value.
	Zero() bool

	EncodeTDLibJSON(b jsontd.Encoder) error
	DecodeTDLibJSON(b jsontd.Decoder) error
}

// DecodeStatisticalGraph implements binary de-serialization for StatisticalGraphClass.
func DecodeStatisticalGraph(buf *bin.Buffer) (StatisticalGraphClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case StatisticalGraphDataTypeID:
		// Decoding statisticalGraphData#89732e2c.
		v := StatisticalGraphData{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode StatisticalGraphClass: %w", err)
		}
		return &v, nil
	case StatisticalGraphAsyncTypeID:
		// Decoding statisticalGraphAsync#19fb2b9f.
		v := StatisticalGraphAsync{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode StatisticalGraphClass: %w", err)
		}
		return &v, nil
	case StatisticalGraphErrorTypeID:
		// Decoding statisticalGraphError#c3fda052.
		v := StatisticalGraphError{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode StatisticalGraphClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode StatisticalGraphClass: %w", bin.NewUnexpectedID(id))
	}
}

// DecodeTDLibJSONStatisticalGraph implements binary de-serialization for StatisticalGraphClass.
func DecodeTDLibJSONStatisticalGraph(buf jsontd.Decoder) (StatisticalGraphClass, error) {
	id, err := buf.FindTypeID()
	if err != nil {
		return nil, err
	}
	switch id {
	case "statisticalGraphData":
		// Decoding statisticalGraphData#89732e2c.
		v := StatisticalGraphData{}
		if err := v.DecodeTDLibJSON(buf); err != nil {
			return nil, fmt.Errorf("unable to decode StatisticalGraphClass: %w", err)
		}
		return &v, nil
	case "statisticalGraphAsync":
		// Decoding statisticalGraphAsync#19fb2b9f.
		v := StatisticalGraphAsync{}
		if err := v.DecodeTDLibJSON(buf); err != nil {
			return nil, fmt.Errorf("unable to decode StatisticalGraphClass: %w", err)
		}
		return &v, nil
	case "statisticalGraphError":
		// Decoding statisticalGraphError#c3fda052.
		v := StatisticalGraphError{}
		if err := v.DecodeTDLibJSON(buf); err != nil {
			return nil, fmt.Errorf("unable to decode StatisticalGraphClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode StatisticalGraphClass: %w", jsontd.NewUnexpectedID(id))
	}
}

// StatisticalGraph boxes the StatisticalGraphClass providing a helper.
type StatisticalGraphBox struct {
	StatisticalGraph StatisticalGraphClass
}

// Decode implements bin.Decoder for StatisticalGraphBox.
func (b *StatisticalGraphBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode StatisticalGraphBox to nil")
	}
	v, err := DecodeStatisticalGraph(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.StatisticalGraph = v
	return nil
}

// Encode implements bin.Encode for StatisticalGraphBox.
func (b *StatisticalGraphBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.StatisticalGraph == nil {
		return fmt.Errorf("unable to encode StatisticalGraphClass as nil")
	}
	return b.StatisticalGraph.Encode(buf)
}

// DecodeTDLibJSON implements bin.Decoder for StatisticalGraphBox.
func (b *StatisticalGraphBox) DecodeTDLibJSON(buf jsontd.Decoder) error {
	if b == nil {
		return fmt.Errorf("unable to decode StatisticalGraphBox to nil")
	}
	v, err := DecodeTDLibJSONStatisticalGraph(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.StatisticalGraph = v
	return nil
}

// EncodeTDLibJSON implements bin.Encode for StatisticalGraphBox.
func (b *StatisticalGraphBox) EncodeTDLibJSON(buf jsontd.Encoder) error {
	if b == nil || b.StatisticalGraph == nil {
		return fmt.Errorf("unable to encode StatisticalGraphClass as nil")
	}
	return b.StatisticalGraph.EncodeTDLibJSON(buf)
}
