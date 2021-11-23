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

// GetLanguagePackInfoRequest represents TL type `getLanguagePackInfo#7bd8daa8`.
type GetLanguagePackInfoRequest struct {
	// Language pack identifier
	LanguagePackID string
}

// GetLanguagePackInfoRequestTypeID is TL type id of GetLanguagePackInfoRequest.
const GetLanguagePackInfoRequestTypeID = 0x7bd8daa8

// Ensuring interfaces in compile-time for GetLanguagePackInfoRequest.
var (
	_ bin.Encoder     = &GetLanguagePackInfoRequest{}
	_ bin.Decoder     = &GetLanguagePackInfoRequest{}
	_ bin.BareEncoder = &GetLanguagePackInfoRequest{}
	_ bin.BareDecoder = &GetLanguagePackInfoRequest{}
)

func (g *GetLanguagePackInfoRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.LanguagePackID == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *GetLanguagePackInfoRequest) String() string {
	if g == nil {
		return "GetLanguagePackInfoRequest(nil)"
	}
	type Alias GetLanguagePackInfoRequest
	return fmt.Sprintf("GetLanguagePackInfoRequest%+v", Alias(*g))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*GetLanguagePackInfoRequest) TypeID() uint32 {
	return GetLanguagePackInfoRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*GetLanguagePackInfoRequest) TypeName() string {
	return "getLanguagePackInfo"
}

// TypeInfo returns info about TL type.
func (g *GetLanguagePackInfoRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "getLanguagePackInfo",
		ID:   GetLanguagePackInfoRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "LanguagePackID",
			SchemaName: "language_pack_id",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (g *GetLanguagePackInfoRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getLanguagePackInfo#7bd8daa8 as nil")
	}
	b.PutID(GetLanguagePackInfoRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *GetLanguagePackInfoRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getLanguagePackInfo#7bd8daa8 as nil")
	}
	b.PutString(g.LanguagePackID)
	return nil
}

// Decode implements bin.Decoder.
func (g *GetLanguagePackInfoRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getLanguagePackInfo#7bd8daa8 to nil")
	}
	if err := b.ConsumeID(GetLanguagePackInfoRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode getLanguagePackInfo#7bd8daa8: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *GetLanguagePackInfoRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getLanguagePackInfo#7bd8daa8 to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode getLanguagePackInfo#7bd8daa8: field language_pack_id: %w", err)
		}
		g.LanguagePackID = value
	}
	return nil
}

// EncodeTDLibJSON implements jsontd.TDLibEncoder.
func (g *GetLanguagePackInfoRequest) EncodeTDLibJSON(b jsontd.Encoder) error {
	if g == nil {
		return fmt.Errorf("can't encode getLanguagePackInfo#7bd8daa8 as nil")
	}
	b.ObjStart()
	b.PutID("getLanguagePackInfo")
	b.FieldStart("language_pack_id")
	b.PutString(g.LanguagePackID)
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements jsontd.TDLibDecoder.
func (g *GetLanguagePackInfoRequest) DecodeTDLibJSON(b jsontd.Decoder) error {
	if g == nil {
		return fmt.Errorf("can't decode getLanguagePackInfo#7bd8daa8 to nil")
	}

	return b.Obj(func(b jsontd.Decoder, key []byte) error {
		switch string(key) {
		case jsontd.TypeField:
			if err := b.ConsumeID("getLanguagePackInfo"); err != nil {
				return fmt.Errorf("unable to decode getLanguagePackInfo#7bd8daa8: %w", err)
			}
		case "language_pack_id":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode getLanguagePackInfo#7bd8daa8: field language_pack_id: %w", err)
			}
			g.LanguagePackID = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetLanguagePackID returns value of LanguagePackID field.
func (g *GetLanguagePackInfoRequest) GetLanguagePackID() (value string) {
	return g.LanguagePackID
}

// GetLanguagePackInfo invokes method getLanguagePackInfo#7bd8daa8 returning error if any.
func (c *Client) GetLanguagePackInfo(ctx context.Context, languagepackid string) (*LanguagePackInfo, error) {
	var result LanguagePackInfo

	request := &GetLanguagePackInfoRequest{
		LanguagePackID: languagepackid,
	}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
