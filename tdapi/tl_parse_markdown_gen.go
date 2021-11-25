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
	"github.com/gotd/td/tdjson"
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
	_ = tdjson.Encoder{}
)

// ParseMarkdownRequest represents TL type `parseMarkdown#2d153aef`.
type ParseMarkdownRequest struct {
	// The text to parse. For example, "__italic__ ~~strikethrough~~ **bold** `code`
	// ```pre``` __[italic__ text_url](telegram.org) __italic**bold italic__bold**"
	Text FormattedText
}

// ParseMarkdownRequestTypeID is TL type id of ParseMarkdownRequest.
const ParseMarkdownRequestTypeID = 0x2d153aef

// Ensuring interfaces in compile-time for ParseMarkdownRequest.
var (
	_ bin.Encoder     = &ParseMarkdownRequest{}
	_ bin.Decoder     = &ParseMarkdownRequest{}
	_ bin.BareEncoder = &ParseMarkdownRequest{}
	_ bin.BareDecoder = &ParseMarkdownRequest{}
)

func (p *ParseMarkdownRequest) Zero() bool {
	if p == nil {
		return true
	}
	if !(p.Text.Zero()) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (p *ParseMarkdownRequest) String() string {
	if p == nil {
		return "ParseMarkdownRequest(nil)"
	}
	type Alias ParseMarkdownRequest
	return fmt.Sprintf("ParseMarkdownRequest%+v", Alias(*p))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ParseMarkdownRequest) TypeID() uint32 {
	return ParseMarkdownRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*ParseMarkdownRequest) TypeName() string {
	return "parseMarkdown"
}

// TypeInfo returns info about TL type.
func (p *ParseMarkdownRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "parseMarkdown",
		ID:   ParseMarkdownRequestTypeID,
	}
	if p == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Text",
			SchemaName: "text",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (p *ParseMarkdownRequest) Encode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode parseMarkdown#2d153aef as nil")
	}
	b.PutID(ParseMarkdownRequestTypeID)
	return p.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (p *ParseMarkdownRequest) EncodeBare(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode parseMarkdown#2d153aef as nil")
	}
	if err := p.Text.Encode(b); err != nil {
		return fmt.Errorf("unable to encode parseMarkdown#2d153aef: field text: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (p *ParseMarkdownRequest) Decode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode parseMarkdown#2d153aef to nil")
	}
	if err := b.ConsumeID(ParseMarkdownRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode parseMarkdown#2d153aef: %w", err)
	}
	return p.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (p *ParseMarkdownRequest) DecodeBare(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode parseMarkdown#2d153aef to nil")
	}
	{
		if err := p.Text.Decode(b); err != nil {
			return fmt.Errorf("unable to decode parseMarkdown#2d153aef: field text: %w", err)
		}
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (p *ParseMarkdownRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if p == nil {
		return fmt.Errorf("can't encode parseMarkdown#2d153aef as nil")
	}
	b.ObjStart()
	b.PutID("parseMarkdown")
	b.FieldStart("text")
	if err := p.Text.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode parseMarkdown#2d153aef: field text: %w", err)
	}
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (p *ParseMarkdownRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if p == nil {
		return fmt.Errorf("can't decode parseMarkdown#2d153aef to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("parseMarkdown"); err != nil {
				return fmt.Errorf("unable to decode parseMarkdown#2d153aef: %w", err)
			}
		case "text":
			if err := p.Text.DecodeTDLibJSON(b); err != nil {
				return fmt.Errorf("unable to decode parseMarkdown#2d153aef: field text: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetText returns value of Text field.
func (p *ParseMarkdownRequest) GetText() (value FormattedText) {
	return p.Text
}

// ParseMarkdown invokes method parseMarkdown#2d153aef returning error if any.
func (c *Client) ParseMarkdown(ctx context.Context, text FormattedText) (*FormattedText, error) {
	var result FormattedText

	request := &ParseMarkdownRequest{
		Text: text,
	}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}