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

// EditInlineMessageMediaRequest represents TL type `editInlineMessageMedia#1676781`.
type EditInlineMessageMediaRequest struct {
	// Inline message identifier
	InlineMessageID string
	// The new message reply markup; for bots only
	ReplyMarkup ReplyMarkupClass
	// New content of the message. Must be one of the following types: inputMessageAnimation,
	// inputMessageAudio, inputMessageDocument, inputMessagePhoto or inputMessageVideo
	InputMessageContent InputMessageContentClass
}

// EditInlineMessageMediaRequestTypeID is TL type id of EditInlineMessageMediaRequest.
const EditInlineMessageMediaRequestTypeID = 0x1676781

// Ensuring interfaces in compile-time for EditInlineMessageMediaRequest.
var (
	_ bin.Encoder     = &EditInlineMessageMediaRequest{}
	_ bin.Decoder     = &EditInlineMessageMediaRequest{}
	_ bin.BareEncoder = &EditInlineMessageMediaRequest{}
	_ bin.BareDecoder = &EditInlineMessageMediaRequest{}
)

func (e *EditInlineMessageMediaRequest) Zero() bool {
	if e == nil {
		return true
	}
	if !(e.InlineMessageID == "") {
		return false
	}
	if !(e.ReplyMarkup == nil) {
		return false
	}
	if !(e.InputMessageContent == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (e *EditInlineMessageMediaRequest) String() string {
	if e == nil {
		return "EditInlineMessageMediaRequest(nil)"
	}
	type Alias EditInlineMessageMediaRequest
	return fmt.Sprintf("EditInlineMessageMediaRequest%+v", Alias(*e))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*EditInlineMessageMediaRequest) TypeID() uint32 {
	return EditInlineMessageMediaRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*EditInlineMessageMediaRequest) TypeName() string {
	return "editInlineMessageMedia"
}

// TypeInfo returns info about TL type.
func (e *EditInlineMessageMediaRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "editInlineMessageMedia",
		ID:   EditInlineMessageMediaRequestTypeID,
	}
	if e == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "InlineMessageID",
			SchemaName: "inline_message_id",
		},
		{
			Name:       "ReplyMarkup",
			SchemaName: "reply_markup",
		},
		{
			Name:       "InputMessageContent",
			SchemaName: "input_message_content",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (e *EditInlineMessageMediaRequest) Encode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't encode editInlineMessageMedia#1676781 as nil")
	}
	b.PutID(EditInlineMessageMediaRequestTypeID)
	return e.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (e *EditInlineMessageMediaRequest) EncodeBare(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't encode editInlineMessageMedia#1676781 as nil")
	}
	b.PutString(e.InlineMessageID)
	if e.ReplyMarkup == nil {
		return fmt.Errorf("unable to encode editInlineMessageMedia#1676781: field reply_markup is nil")
	}
	if err := e.ReplyMarkup.Encode(b); err != nil {
		return fmt.Errorf("unable to encode editInlineMessageMedia#1676781: field reply_markup: %w", err)
	}
	if e.InputMessageContent == nil {
		return fmt.Errorf("unable to encode editInlineMessageMedia#1676781: field input_message_content is nil")
	}
	if err := e.InputMessageContent.Encode(b); err != nil {
		return fmt.Errorf("unable to encode editInlineMessageMedia#1676781: field input_message_content: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (e *EditInlineMessageMediaRequest) Decode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't decode editInlineMessageMedia#1676781 to nil")
	}
	if err := b.ConsumeID(EditInlineMessageMediaRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode editInlineMessageMedia#1676781: %w", err)
	}
	return e.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (e *EditInlineMessageMediaRequest) DecodeBare(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't decode editInlineMessageMedia#1676781 to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode editInlineMessageMedia#1676781: field inline_message_id: %w", err)
		}
		e.InlineMessageID = value
	}
	{
		value, err := DecodeReplyMarkup(b)
		if err != nil {
			return fmt.Errorf("unable to decode editInlineMessageMedia#1676781: field reply_markup: %w", err)
		}
		e.ReplyMarkup = value
	}
	{
		value, err := DecodeInputMessageContent(b)
		if err != nil {
			return fmt.Errorf("unable to decode editInlineMessageMedia#1676781: field input_message_content: %w", err)
		}
		e.InputMessageContent = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (e *EditInlineMessageMediaRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if e == nil {
		return fmt.Errorf("can't encode editInlineMessageMedia#1676781 as nil")
	}
	b.ObjStart()
	b.PutID("editInlineMessageMedia")
	b.FieldStart("inline_message_id")
	b.PutString(e.InlineMessageID)
	b.FieldStart("reply_markup")
	if e.ReplyMarkup == nil {
		return fmt.Errorf("unable to encode editInlineMessageMedia#1676781: field reply_markup is nil")
	}
	if err := e.ReplyMarkup.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode editInlineMessageMedia#1676781: field reply_markup: %w", err)
	}
	b.FieldStart("input_message_content")
	if e.InputMessageContent == nil {
		return fmt.Errorf("unable to encode editInlineMessageMedia#1676781: field input_message_content is nil")
	}
	if err := e.InputMessageContent.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode editInlineMessageMedia#1676781: field input_message_content: %w", err)
	}
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (e *EditInlineMessageMediaRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if e == nil {
		return fmt.Errorf("can't decode editInlineMessageMedia#1676781 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("editInlineMessageMedia"); err != nil {
				return fmt.Errorf("unable to decode editInlineMessageMedia#1676781: %w", err)
			}
		case "inline_message_id":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode editInlineMessageMedia#1676781: field inline_message_id: %w", err)
			}
			e.InlineMessageID = value
		case "reply_markup":
			value, err := DecodeTDLibJSONReplyMarkup(b)
			if err != nil {
				return fmt.Errorf("unable to decode editInlineMessageMedia#1676781: field reply_markup: %w", err)
			}
			e.ReplyMarkup = value
		case "input_message_content":
			value, err := DecodeTDLibJSONInputMessageContent(b)
			if err != nil {
				return fmt.Errorf("unable to decode editInlineMessageMedia#1676781: field input_message_content: %w", err)
			}
			e.InputMessageContent = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetInlineMessageID returns value of InlineMessageID field.
func (e *EditInlineMessageMediaRequest) GetInlineMessageID() (value string) {
	return e.InlineMessageID
}

// GetReplyMarkup returns value of ReplyMarkup field.
func (e *EditInlineMessageMediaRequest) GetReplyMarkup() (value ReplyMarkupClass) {
	return e.ReplyMarkup
}

// GetInputMessageContent returns value of InputMessageContent field.
func (e *EditInlineMessageMediaRequest) GetInputMessageContent() (value InputMessageContentClass) {
	return e.InputMessageContent
}

// EditInlineMessageMedia invokes method editInlineMessageMedia#1676781 returning error if any.
func (c *Client) EditInlineMessageMedia(ctx context.Context, request *EditInlineMessageMediaRequest) error {
	var ok Ok

	if err := c.rpc.Invoke(ctx, request, &ok); err != nil {
		return err
	}
	return nil
}