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

// AnswerShippingQueryRequest represents TL type `answerShippingQuery#7a3c2432`.
type AnswerShippingQueryRequest struct {
	// Identifier of the shipping query
	ShippingQueryID int64
	// Available shipping options
	ShippingOptions []ShippingOption
	// An error message, empty on success
	ErrorMessage string
}

// AnswerShippingQueryRequestTypeID is TL type id of AnswerShippingQueryRequest.
const AnswerShippingQueryRequestTypeID = 0x7a3c2432

// Ensuring interfaces in compile-time for AnswerShippingQueryRequest.
var (
	_ bin.Encoder     = &AnswerShippingQueryRequest{}
	_ bin.Decoder     = &AnswerShippingQueryRequest{}
	_ bin.BareEncoder = &AnswerShippingQueryRequest{}
	_ bin.BareDecoder = &AnswerShippingQueryRequest{}
)

func (a *AnswerShippingQueryRequest) Zero() bool {
	if a == nil {
		return true
	}
	if !(a.ShippingQueryID == 0) {
		return false
	}
	if !(a.ShippingOptions == nil) {
		return false
	}
	if !(a.ErrorMessage == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (a *AnswerShippingQueryRequest) String() string {
	if a == nil {
		return "AnswerShippingQueryRequest(nil)"
	}
	type Alias AnswerShippingQueryRequest
	return fmt.Sprintf("AnswerShippingQueryRequest%+v", Alias(*a))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*AnswerShippingQueryRequest) TypeID() uint32 {
	return AnswerShippingQueryRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*AnswerShippingQueryRequest) TypeName() string {
	return "answerShippingQuery"
}

// TypeInfo returns info about TL type.
func (a *AnswerShippingQueryRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "answerShippingQuery",
		ID:   AnswerShippingQueryRequestTypeID,
	}
	if a == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ShippingQueryID",
			SchemaName: "shipping_query_id",
		},
		{
			Name:       "ShippingOptions",
			SchemaName: "shipping_options",
		},
		{
			Name:       "ErrorMessage",
			SchemaName: "error_message",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (a *AnswerShippingQueryRequest) Encode(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't encode answerShippingQuery#7a3c2432 as nil")
	}
	b.PutID(AnswerShippingQueryRequestTypeID)
	return a.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (a *AnswerShippingQueryRequest) EncodeBare(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't encode answerShippingQuery#7a3c2432 as nil")
	}
	b.PutLong(a.ShippingQueryID)
	b.PutInt(len(a.ShippingOptions))
	for idx, v := range a.ShippingOptions {
		if err := v.EncodeBare(b); err != nil {
			return fmt.Errorf("unable to encode bare answerShippingQuery#7a3c2432: field shipping_options element with index %d: %w", idx, err)
		}
	}
	b.PutString(a.ErrorMessage)
	return nil
}

// Decode implements bin.Decoder.
func (a *AnswerShippingQueryRequest) Decode(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't decode answerShippingQuery#7a3c2432 to nil")
	}
	if err := b.ConsumeID(AnswerShippingQueryRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode answerShippingQuery#7a3c2432: %w", err)
	}
	return a.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (a *AnswerShippingQueryRequest) DecodeBare(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't decode answerShippingQuery#7a3c2432 to nil")
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode answerShippingQuery#7a3c2432: field shipping_query_id: %w", err)
		}
		a.ShippingQueryID = value
	}
	{
		headerLen, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode answerShippingQuery#7a3c2432: field shipping_options: %w", err)
		}

		if headerLen > 0 {
			a.ShippingOptions = make([]ShippingOption, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			var value ShippingOption
			if err := value.DecodeBare(b); err != nil {
				return fmt.Errorf("unable to decode bare answerShippingQuery#7a3c2432: field shipping_options: %w", err)
			}
			a.ShippingOptions = append(a.ShippingOptions, value)
		}
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode answerShippingQuery#7a3c2432: field error_message: %w", err)
		}
		a.ErrorMessage = value
	}
	return nil
}

// EncodeTDLibJSON implements jsontd.TDLibEncoder.
func (a *AnswerShippingQueryRequest) EncodeTDLibJSON(b jsontd.Encoder) error {
	if a == nil {
		return fmt.Errorf("can't encode answerShippingQuery#7a3c2432 as nil")
	}
	b.ObjStart()
	b.PutID("answerShippingQuery")
	b.FieldStart("shipping_query_id")
	b.PutLong(a.ShippingQueryID)
	b.FieldStart("shipping_options")
	b.ArrStart()
	for idx, v := range a.ShippingOptions {
		if err := v.EncodeTDLibJSON(b); err != nil {
			return fmt.Errorf("unable to encode answerShippingQuery#7a3c2432: field shipping_options element with index %d: %w", idx, err)
		}
	}
	b.ArrEnd()
	b.FieldStart("error_message")
	b.PutString(a.ErrorMessage)
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements jsontd.TDLibDecoder.
func (a *AnswerShippingQueryRequest) DecodeTDLibJSON(b jsontd.Decoder) error {
	if a == nil {
		return fmt.Errorf("can't decode answerShippingQuery#7a3c2432 to nil")
	}

	return b.Obj(func(b jsontd.Decoder, key []byte) error {
		switch string(key) {
		case jsontd.TypeField:
			if err := b.ConsumeID("answerShippingQuery"); err != nil {
				return fmt.Errorf("unable to decode answerShippingQuery#7a3c2432: %w", err)
			}
		case "shipping_query_id":
			value, err := b.Long()
			if err != nil {
				return fmt.Errorf("unable to decode answerShippingQuery#7a3c2432: field shipping_query_id: %w", err)
			}
			a.ShippingQueryID = value
		case "shipping_options":
			if err := b.Arr(func(b jsontd.Decoder) error {
				var value ShippingOption
				if err := value.DecodeTDLibJSON(b); err != nil {
					return fmt.Errorf("unable to decode answerShippingQuery#7a3c2432: field shipping_options: %w", err)
				}
				a.ShippingOptions = append(a.ShippingOptions, value)
				return nil
			}); err != nil {
				return fmt.Errorf("unable to decode answerShippingQuery#7a3c2432: field shipping_options: %w", err)
			}
		case "error_message":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode answerShippingQuery#7a3c2432: field error_message: %w", err)
			}
			a.ErrorMessage = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetShippingQueryID returns value of ShippingQueryID field.
func (a *AnswerShippingQueryRequest) GetShippingQueryID() (value int64) {
	return a.ShippingQueryID
}

// GetShippingOptions returns value of ShippingOptions field.
func (a *AnswerShippingQueryRequest) GetShippingOptions() (value []ShippingOption) {
	return a.ShippingOptions
}

// GetErrorMessage returns value of ErrorMessage field.
func (a *AnswerShippingQueryRequest) GetErrorMessage() (value string) {
	return a.ErrorMessage
}

// AnswerShippingQuery invokes method answerShippingQuery#7a3c2432 returning error if any.
func (c *Client) AnswerShippingQuery(ctx context.Context, request *AnswerShippingQueryRequest) error {
	var ok Ok

	if err := c.rpc.Invoke(ctx, request, &ok); err != nil {
		return err
	}
	return nil
}
