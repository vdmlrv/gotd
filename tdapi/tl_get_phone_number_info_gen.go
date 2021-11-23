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

// GetPhoneNumberInfoRequest represents TL type `getPhoneNumberInfo#a0229bf9`.
type GetPhoneNumberInfoRequest struct {
	// The phone number prefix
	PhoneNumberPrefix string
}

// GetPhoneNumberInfoRequestTypeID is TL type id of GetPhoneNumberInfoRequest.
const GetPhoneNumberInfoRequestTypeID = 0xa0229bf9

// Ensuring interfaces in compile-time for GetPhoneNumberInfoRequest.
var (
	_ bin.Encoder     = &GetPhoneNumberInfoRequest{}
	_ bin.Decoder     = &GetPhoneNumberInfoRequest{}
	_ bin.BareEncoder = &GetPhoneNumberInfoRequest{}
	_ bin.BareDecoder = &GetPhoneNumberInfoRequest{}
)

func (g *GetPhoneNumberInfoRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.PhoneNumberPrefix == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *GetPhoneNumberInfoRequest) String() string {
	if g == nil {
		return "GetPhoneNumberInfoRequest(nil)"
	}
	type Alias GetPhoneNumberInfoRequest
	return fmt.Sprintf("GetPhoneNumberInfoRequest%+v", Alias(*g))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*GetPhoneNumberInfoRequest) TypeID() uint32 {
	return GetPhoneNumberInfoRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*GetPhoneNumberInfoRequest) TypeName() string {
	return "getPhoneNumberInfo"
}

// TypeInfo returns info about TL type.
func (g *GetPhoneNumberInfoRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "getPhoneNumberInfo",
		ID:   GetPhoneNumberInfoRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "PhoneNumberPrefix",
			SchemaName: "phone_number_prefix",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (g *GetPhoneNumberInfoRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getPhoneNumberInfo#a0229bf9 as nil")
	}
	b.PutID(GetPhoneNumberInfoRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *GetPhoneNumberInfoRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getPhoneNumberInfo#a0229bf9 as nil")
	}
	b.PutString(g.PhoneNumberPrefix)
	return nil
}

// Decode implements bin.Decoder.
func (g *GetPhoneNumberInfoRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getPhoneNumberInfo#a0229bf9 to nil")
	}
	if err := b.ConsumeID(GetPhoneNumberInfoRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode getPhoneNumberInfo#a0229bf9: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *GetPhoneNumberInfoRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getPhoneNumberInfo#a0229bf9 to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode getPhoneNumberInfo#a0229bf9: field phone_number_prefix: %w", err)
		}
		g.PhoneNumberPrefix = value
	}
	return nil
}

// EncodeTDLibJSON implements jsontd.TDLibEncoder.
func (g *GetPhoneNumberInfoRequest) EncodeTDLibJSON(b jsontd.Encoder) error {
	if g == nil {
		return fmt.Errorf("can't encode getPhoneNumberInfo#a0229bf9 as nil")
	}
	b.ObjStart()
	b.PutID("getPhoneNumberInfo")
	b.FieldStart("phone_number_prefix")
	b.PutString(g.PhoneNumberPrefix)
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements jsontd.TDLibDecoder.
func (g *GetPhoneNumberInfoRequest) DecodeTDLibJSON(b jsontd.Decoder) error {
	if g == nil {
		return fmt.Errorf("can't decode getPhoneNumberInfo#a0229bf9 to nil")
	}

	return b.Obj(func(b jsontd.Decoder, key []byte) error {
		switch string(key) {
		case jsontd.TypeField:
			if err := b.ConsumeID("getPhoneNumberInfo"); err != nil {
				return fmt.Errorf("unable to decode getPhoneNumberInfo#a0229bf9: %w", err)
			}
		case "phone_number_prefix":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode getPhoneNumberInfo#a0229bf9: field phone_number_prefix: %w", err)
			}
			g.PhoneNumberPrefix = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetPhoneNumberPrefix returns value of PhoneNumberPrefix field.
func (g *GetPhoneNumberInfoRequest) GetPhoneNumberPrefix() (value string) {
	return g.PhoneNumberPrefix
}

// GetPhoneNumberInfo invokes method getPhoneNumberInfo#a0229bf9 returning error if any.
func (c *Client) GetPhoneNumberInfo(ctx context.Context, phonenumberprefix string) (*PhoneNumberInfo, error) {
	var result PhoneNumberInfo

	request := &GetPhoneNumberInfoRequest{
		PhoneNumberPrefix: phonenumberprefix,
	}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
