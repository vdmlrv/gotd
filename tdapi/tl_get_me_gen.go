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

// GetMeRequest represents TL type `getMe#f495b27f`.
type GetMeRequest struct {
}

// GetMeRequestTypeID is TL type id of GetMeRequest.
const GetMeRequestTypeID = 0xf495b27f

// Ensuring interfaces in compile-time for GetMeRequest.
var (
	_ bin.Encoder     = &GetMeRequest{}
	_ bin.Decoder     = &GetMeRequest{}
	_ bin.BareEncoder = &GetMeRequest{}
	_ bin.BareDecoder = &GetMeRequest{}
)

func (g *GetMeRequest) Zero() bool {
	if g == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (g *GetMeRequest) String() string {
	if g == nil {
		return "GetMeRequest(nil)"
	}
	type Alias GetMeRequest
	return fmt.Sprintf("GetMeRequest%+v", Alias(*g))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*GetMeRequest) TypeID() uint32 {
	return GetMeRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*GetMeRequest) TypeName() string {
	return "getMe"
}

// TypeInfo returns info about TL type.
func (g *GetMeRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "getMe",
		ID:   GetMeRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (g *GetMeRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getMe#f495b27f as nil")
	}
	b.PutID(GetMeRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *GetMeRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getMe#f495b27f as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (g *GetMeRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getMe#f495b27f to nil")
	}
	if err := b.ConsumeID(GetMeRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode getMe#f495b27f: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *GetMeRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getMe#f495b27f to nil")
	}
	return nil
}

// EncodeTDLibJSON implements jsontd.TDLibEncoder.
func (g *GetMeRequest) EncodeTDLibJSON(b jsontd.Encoder) error {
	if g == nil {
		return fmt.Errorf("can't encode getMe#f495b27f as nil")
	}
	b.ObjStart()
	b.PutID("getMe")
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements jsontd.TDLibDecoder.
func (g *GetMeRequest) DecodeTDLibJSON(b jsontd.Decoder) error {
	if g == nil {
		return fmt.Errorf("can't decode getMe#f495b27f to nil")
	}

	return b.Obj(func(b jsontd.Decoder, key []byte) error {
		switch string(key) {
		case jsontd.TypeField:
			if err := b.ConsumeID("getMe"); err != nil {
				return fmt.Errorf("unable to decode getMe#f495b27f: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetMe invokes method getMe#f495b27f returning error if any.
func (c *Client) GetMe(ctx context.Context) (*User, error) {
	var result User

	request := &GetMeRequest{}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
