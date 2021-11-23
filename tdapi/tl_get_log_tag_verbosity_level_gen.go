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

// GetLogTagVerbosityLevelRequest represents TL type `getLogTagVerbosityLevel#38af2d83`.
type GetLogTagVerbosityLevelRequest struct {
	// Logging tag to change verbosity level
	Tag string
}

// GetLogTagVerbosityLevelRequestTypeID is TL type id of GetLogTagVerbosityLevelRequest.
const GetLogTagVerbosityLevelRequestTypeID = 0x38af2d83

// Ensuring interfaces in compile-time for GetLogTagVerbosityLevelRequest.
var (
	_ bin.Encoder     = &GetLogTagVerbosityLevelRequest{}
	_ bin.Decoder     = &GetLogTagVerbosityLevelRequest{}
	_ bin.BareEncoder = &GetLogTagVerbosityLevelRequest{}
	_ bin.BareDecoder = &GetLogTagVerbosityLevelRequest{}
)

func (g *GetLogTagVerbosityLevelRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.Tag == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *GetLogTagVerbosityLevelRequest) String() string {
	if g == nil {
		return "GetLogTagVerbosityLevelRequest(nil)"
	}
	type Alias GetLogTagVerbosityLevelRequest
	return fmt.Sprintf("GetLogTagVerbosityLevelRequest%+v", Alias(*g))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*GetLogTagVerbosityLevelRequest) TypeID() uint32 {
	return GetLogTagVerbosityLevelRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*GetLogTagVerbosityLevelRequest) TypeName() string {
	return "getLogTagVerbosityLevel"
}

// TypeInfo returns info about TL type.
func (g *GetLogTagVerbosityLevelRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "getLogTagVerbosityLevel",
		ID:   GetLogTagVerbosityLevelRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Tag",
			SchemaName: "tag",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (g *GetLogTagVerbosityLevelRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getLogTagVerbosityLevel#38af2d83 as nil")
	}
	b.PutID(GetLogTagVerbosityLevelRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *GetLogTagVerbosityLevelRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getLogTagVerbosityLevel#38af2d83 as nil")
	}
	b.PutString(g.Tag)
	return nil
}

// Decode implements bin.Decoder.
func (g *GetLogTagVerbosityLevelRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getLogTagVerbosityLevel#38af2d83 to nil")
	}
	if err := b.ConsumeID(GetLogTagVerbosityLevelRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode getLogTagVerbosityLevel#38af2d83: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *GetLogTagVerbosityLevelRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getLogTagVerbosityLevel#38af2d83 to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode getLogTagVerbosityLevel#38af2d83: field tag: %w", err)
		}
		g.Tag = value
	}
	return nil
}

// EncodeTDLibJSON implements jsontd.TDLibEncoder.
func (g *GetLogTagVerbosityLevelRequest) EncodeTDLibJSON(b jsontd.Encoder) error {
	if g == nil {
		return fmt.Errorf("can't encode getLogTagVerbosityLevel#38af2d83 as nil")
	}
	b.ObjStart()
	b.PutID("getLogTagVerbosityLevel")
	b.FieldStart("tag")
	b.PutString(g.Tag)
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements jsontd.TDLibDecoder.
func (g *GetLogTagVerbosityLevelRequest) DecodeTDLibJSON(b jsontd.Decoder) error {
	if g == nil {
		return fmt.Errorf("can't decode getLogTagVerbosityLevel#38af2d83 to nil")
	}

	return b.Obj(func(b jsontd.Decoder, key []byte) error {
		switch string(key) {
		case jsontd.TypeField:
			if err := b.ConsumeID("getLogTagVerbosityLevel"); err != nil {
				return fmt.Errorf("unable to decode getLogTagVerbosityLevel#38af2d83: %w", err)
			}
		case "tag":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode getLogTagVerbosityLevel#38af2d83: field tag: %w", err)
			}
			g.Tag = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetTag returns value of Tag field.
func (g *GetLogTagVerbosityLevelRequest) GetTag() (value string) {
	return g.Tag
}

// GetLogTagVerbosityLevel invokes method getLogTagVerbosityLevel#38af2d83 returning error if any.
func (c *Client) GetLogTagVerbosityLevel(ctx context.Context, tag string) (*LogVerbosityLevel, error) {
	var result LogVerbosityLevel

	request := &GetLogTagVerbosityLevelRequest{
		Tag: tag,
	}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
