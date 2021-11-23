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

// SendInlineQueryResultMessageRequest represents TL type `sendInlineQueryResultMessage#c774e89c`.
type SendInlineQueryResultMessageRequest struct {
	// Target chat
	ChatID int64
	// If not 0, a message thread identifier in which the message will be sent
	MessageThreadID int64
	// Identifier of a message to reply to or 0
	ReplyToMessageID int64
	// Options to be used to send the message
	Options MessageSendOptions
	// Identifier of the inline query
	QueryID int64
	// Identifier of the inline result
	ResultID string
	// If true, there will be no mention of a bot, via which the message is sent. Can be used
	// only for bots GetOption("animation_search_bot_username"),
	// GetOption("photo_search_bot_username") and GetOption("venue_search_bot_username")
	HideViaBot bool
}

// SendInlineQueryResultMessageRequestTypeID is TL type id of SendInlineQueryResultMessageRequest.
const SendInlineQueryResultMessageRequestTypeID = 0xc774e89c

// Ensuring interfaces in compile-time for SendInlineQueryResultMessageRequest.
var (
	_ bin.Encoder     = &SendInlineQueryResultMessageRequest{}
	_ bin.Decoder     = &SendInlineQueryResultMessageRequest{}
	_ bin.BareEncoder = &SendInlineQueryResultMessageRequest{}
	_ bin.BareDecoder = &SendInlineQueryResultMessageRequest{}
)

func (s *SendInlineQueryResultMessageRequest) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.ChatID == 0) {
		return false
	}
	if !(s.MessageThreadID == 0) {
		return false
	}
	if !(s.ReplyToMessageID == 0) {
		return false
	}
	if !(s.Options.Zero()) {
		return false
	}
	if !(s.QueryID == 0) {
		return false
	}
	if !(s.ResultID == "") {
		return false
	}
	if !(s.HideViaBot == false) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *SendInlineQueryResultMessageRequest) String() string {
	if s == nil {
		return "SendInlineQueryResultMessageRequest(nil)"
	}
	type Alias SendInlineQueryResultMessageRequest
	return fmt.Sprintf("SendInlineQueryResultMessageRequest%+v", Alias(*s))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*SendInlineQueryResultMessageRequest) TypeID() uint32 {
	return SendInlineQueryResultMessageRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*SendInlineQueryResultMessageRequest) TypeName() string {
	return "sendInlineQueryResultMessage"
}

// TypeInfo returns info about TL type.
func (s *SendInlineQueryResultMessageRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "sendInlineQueryResultMessage",
		ID:   SendInlineQueryResultMessageRequestTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ChatID",
			SchemaName: "chat_id",
		},
		{
			Name:       "MessageThreadID",
			SchemaName: "message_thread_id",
		},
		{
			Name:       "ReplyToMessageID",
			SchemaName: "reply_to_message_id",
		},
		{
			Name:       "Options",
			SchemaName: "options",
		},
		{
			Name:       "QueryID",
			SchemaName: "query_id",
		},
		{
			Name:       "ResultID",
			SchemaName: "result_id",
		},
		{
			Name:       "HideViaBot",
			SchemaName: "hide_via_bot",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (s *SendInlineQueryResultMessageRequest) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode sendInlineQueryResultMessage#c774e89c as nil")
	}
	b.PutID(SendInlineQueryResultMessageRequestTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *SendInlineQueryResultMessageRequest) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode sendInlineQueryResultMessage#c774e89c as nil")
	}
	b.PutLong(s.ChatID)
	b.PutLong(s.MessageThreadID)
	b.PutLong(s.ReplyToMessageID)
	if err := s.Options.Encode(b); err != nil {
		return fmt.Errorf("unable to encode sendInlineQueryResultMessage#c774e89c: field options: %w", err)
	}
	b.PutLong(s.QueryID)
	b.PutString(s.ResultID)
	b.PutBool(s.HideViaBot)
	return nil
}

// Decode implements bin.Decoder.
func (s *SendInlineQueryResultMessageRequest) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode sendInlineQueryResultMessage#c774e89c to nil")
	}
	if err := b.ConsumeID(SendInlineQueryResultMessageRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode sendInlineQueryResultMessage#c774e89c: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *SendInlineQueryResultMessageRequest) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode sendInlineQueryResultMessage#c774e89c to nil")
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode sendInlineQueryResultMessage#c774e89c: field chat_id: %w", err)
		}
		s.ChatID = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode sendInlineQueryResultMessage#c774e89c: field message_thread_id: %w", err)
		}
		s.MessageThreadID = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode sendInlineQueryResultMessage#c774e89c: field reply_to_message_id: %w", err)
		}
		s.ReplyToMessageID = value
	}
	{
		if err := s.Options.Decode(b); err != nil {
			return fmt.Errorf("unable to decode sendInlineQueryResultMessage#c774e89c: field options: %w", err)
		}
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode sendInlineQueryResultMessage#c774e89c: field query_id: %w", err)
		}
		s.QueryID = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode sendInlineQueryResultMessage#c774e89c: field result_id: %w", err)
		}
		s.ResultID = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode sendInlineQueryResultMessage#c774e89c: field hide_via_bot: %w", err)
		}
		s.HideViaBot = value
	}
	return nil
}

// EncodeTDLibJSON implements jsontd.TDLibEncoder.
func (s *SendInlineQueryResultMessageRequest) EncodeTDLibJSON(b jsontd.Encoder) error {
	if s == nil {
		return fmt.Errorf("can't encode sendInlineQueryResultMessage#c774e89c as nil")
	}
	b.ObjStart()
	b.PutID("sendInlineQueryResultMessage")
	b.FieldStart("chat_id")
	b.PutLong(s.ChatID)
	b.FieldStart("message_thread_id")
	b.PutLong(s.MessageThreadID)
	b.FieldStart("reply_to_message_id")
	b.PutLong(s.ReplyToMessageID)
	b.FieldStart("options")
	if err := s.Options.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode sendInlineQueryResultMessage#c774e89c: field options: %w", err)
	}
	b.FieldStart("query_id")
	b.PutLong(s.QueryID)
	b.FieldStart("result_id")
	b.PutString(s.ResultID)
	b.FieldStart("hide_via_bot")
	b.PutBool(s.HideViaBot)
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements jsontd.TDLibDecoder.
func (s *SendInlineQueryResultMessageRequest) DecodeTDLibJSON(b jsontd.Decoder) error {
	if s == nil {
		return fmt.Errorf("can't decode sendInlineQueryResultMessage#c774e89c to nil")
	}

	return b.Obj(func(b jsontd.Decoder, key []byte) error {
		switch string(key) {
		case jsontd.TypeField:
			if err := b.ConsumeID("sendInlineQueryResultMessage"); err != nil {
				return fmt.Errorf("unable to decode sendInlineQueryResultMessage#c774e89c: %w", err)
			}
		case "chat_id":
			value, err := b.Long()
			if err != nil {
				return fmt.Errorf("unable to decode sendInlineQueryResultMessage#c774e89c: field chat_id: %w", err)
			}
			s.ChatID = value
		case "message_thread_id":
			value, err := b.Long()
			if err != nil {
				return fmt.Errorf("unable to decode sendInlineQueryResultMessage#c774e89c: field message_thread_id: %w", err)
			}
			s.MessageThreadID = value
		case "reply_to_message_id":
			value, err := b.Long()
			if err != nil {
				return fmt.Errorf("unable to decode sendInlineQueryResultMessage#c774e89c: field reply_to_message_id: %w", err)
			}
			s.ReplyToMessageID = value
		case "options":
			if err := s.Options.DecodeTDLibJSON(b); err != nil {
				return fmt.Errorf("unable to decode sendInlineQueryResultMessage#c774e89c: field options: %w", err)
			}
		case "query_id":
			value, err := b.Long()
			if err != nil {
				return fmt.Errorf("unable to decode sendInlineQueryResultMessage#c774e89c: field query_id: %w", err)
			}
			s.QueryID = value
		case "result_id":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode sendInlineQueryResultMessage#c774e89c: field result_id: %w", err)
			}
			s.ResultID = value
		case "hide_via_bot":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode sendInlineQueryResultMessage#c774e89c: field hide_via_bot: %w", err)
			}
			s.HideViaBot = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetChatID returns value of ChatID field.
func (s *SendInlineQueryResultMessageRequest) GetChatID() (value int64) {
	return s.ChatID
}

// GetMessageThreadID returns value of MessageThreadID field.
func (s *SendInlineQueryResultMessageRequest) GetMessageThreadID() (value int64) {
	return s.MessageThreadID
}

// GetReplyToMessageID returns value of ReplyToMessageID field.
func (s *SendInlineQueryResultMessageRequest) GetReplyToMessageID() (value int64) {
	return s.ReplyToMessageID
}

// GetOptions returns value of Options field.
func (s *SendInlineQueryResultMessageRequest) GetOptions() (value MessageSendOptions) {
	return s.Options
}

// GetQueryID returns value of QueryID field.
func (s *SendInlineQueryResultMessageRequest) GetQueryID() (value int64) {
	return s.QueryID
}

// GetResultID returns value of ResultID field.
func (s *SendInlineQueryResultMessageRequest) GetResultID() (value string) {
	return s.ResultID
}

// GetHideViaBot returns value of HideViaBot field.
func (s *SendInlineQueryResultMessageRequest) GetHideViaBot() (value bool) {
	return s.HideViaBot
}

// SendInlineQueryResultMessage invokes method sendInlineQueryResultMessage#c774e89c returning error if any.
func (c *Client) SendInlineQueryResultMessage(ctx context.Context, request *SendInlineQueryResultMessageRequest) (*Message, error) {
	var result Message

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
