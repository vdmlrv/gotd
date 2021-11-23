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

// MessageLinkInfo represents TL type `messageLinkInfo#c57d442a`.
type MessageLinkInfo struct {
	// True, if the link is a public link for a message in a chat
	IsPublic bool
	// If found, identifier of the chat to which the message belongs, 0 otherwise
	ChatID int64
	// If found, the linked message; may be null
	Message Message
	// Timestamp from which the video/audio/video note/voice note playing should start, in
	// seconds; 0 if not specified. The media can be in the message content or in its web
	// page preview
	MediaTimestamp int32
	// True, if the whole media album to which the message belongs is linked
	ForAlbum bool
	// True, if the message is linked as a channel post comment or from a message thread
	ForComment bool
}

// MessageLinkInfoTypeID is TL type id of MessageLinkInfo.
const MessageLinkInfoTypeID = 0xc57d442a

// Ensuring interfaces in compile-time for MessageLinkInfo.
var (
	_ bin.Encoder     = &MessageLinkInfo{}
	_ bin.Decoder     = &MessageLinkInfo{}
	_ bin.BareEncoder = &MessageLinkInfo{}
	_ bin.BareDecoder = &MessageLinkInfo{}
)

func (m *MessageLinkInfo) Zero() bool {
	if m == nil {
		return true
	}
	if !(m.IsPublic == false) {
		return false
	}
	if !(m.ChatID == 0) {
		return false
	}
	if !(m.Message.Zero()) {
		return false
	}
	if !(m.MediaTimestamp == 0) {
		return false
	}
	if !(m.ForAlbum == false) {
		return false
	}
	if !(m.ForComment == false) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (m *MessageLinkInfo) String() string {
	if m == nil {
		return "MessageLinkInfo(nil)"
	}
	type Alias MessageLinkInfo
	return fmt.Sprintf("MessageLinkInfo%+v", Alias(*m))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessageLinkInfo) TypeID() uint32 {
	return MessageLinkInfoTypeID
}

// TypeName returns name of type in TL schema.
func (*MessageLinkInfo) TypeName() string {
	return "messageLinkInfo"
}

// TypeInfo returns info about TL type.
func (m *MessageLinkInfo) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messageLinkInfo",
		ID:   MessageLinkInfoTypeID,
	}
	if m == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "IsPublic",
			SchemaName: "is_public",
		},
		{
			Name:       "ChatID",
			SchemaName: "chat_id",
		},
		{
			Name:       "Message",
			SchemaName: "message",
		},
		{
			Name:       "MediaTimestamp",
			SchemaName: "media_timestamp",
		},
		{
			Name:       "ForAlbum",
			SchemaName: "for_album",
		},
		{
			Name:       "ForComment",
			SchemaName: "for_comment",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (m *MessageLinkInfo) Encode(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't encode messageLinkInfo#c57d442a as nil")
	}
	b.PutID(MessageLinkInfoTypeID)
	return m.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (m *MessageLinkInfo) EncodeBare(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't encode messageLinkInfo#c57d442a as nil")
	}
	b.PutBool(m.IsPublic)
	b.PutLong(m.ChatID)
	if err := m.Message.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messageLinkInfo#c57d442a: field message: %w", err)
	}
	b.PutInt32(m.MediaTimestamp)
	b.PutBool(m.ForAlbum)
	b.PutBool(m.ForComment)
	return nil
}

// Decode implements bin.Decoder.
func (m *MessageLinkInfo) Decode(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't decode messageLinkInfo#c57d442a to nil")
	}
	if err := b.ConsumeID(MessageLinkInfoTypeID); err != nil {
		return fmt.Errorf("unable to decode messageLinkInfo#c57d442a: %w", err)
	}
	return m.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (m *MessageLinkInfo) DecodeBare(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't decode messageLinkInfo#c57d442a to nil")
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode messageLinkInfo#c57d442a: field is_public: %w", err)
		}
		m.IsPublic = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode messageLinkInfo#c57d442a: field chat_id: %w", err)
		}
		m.ChatID = value
	}
	{
		if err := m.Message.Decode(b); err != nil {
			return fmt.Errorf("unable to decode messageLinkInfo#c57d442a: field message: %w", err)
		}
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode messageLinkInfo#c57d442a: field media_timestamp: %w", err)
		}
		m.MediaTimestamp = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode messageLinkInfo#c57d442a: field for_album: %w", err)
		}
		m.ForAlbum = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode messageLinkInfo#c57d442a: field for_comment: %w", err)
		}
		m.ForComment = value
	}
	return nil
}

// EncodeTDLibJSON implements jsontd.TDLibEncoder.
func (m *MessageLinkInfo) EncodeTDLibJSON(b jsontd.Encoder) error {
	if m == nil {
		return fmt.Errorf("can't encode messageLinkInfo#c57d442a as nil")
	}
	b.ObjStart()
	b.PutID("messageLinkInfo")
	b.FieldStart("is_public")
	b.PutBool(m.IsPublic)
	b.FieldStart("chat_id")
	b.PutLong(m.ChatID)
	b.FieldStart("message")
	if err := m.Message.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode messageLinkInfo#c57d442a: field message: %w", err)
	}
	b.FieldStart("media_timestamp")
	b.PutInt32(m.MediaTimestamp)
	b.FieldStart("for_album")
	b.PutBool(m.ForAlbum)
	b.FieldStart("for_comment")
	b.PutBool(m.ForComment)
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements jsontd.TDLibDecoder.
func (m *MessageLinkInfo) DecodeTDLibJSON(b jsontd.Decoder) error {
	if m == nil {
		return fmt.Errorf("can't decode messageLinkInfo#c57d442a to nil")
	}

	return b.Obj(func(b jsontd.Decoder, key []byte) error {
		switch string(key) {
		case jsontd.TypeField:
			if err := b.ConsumeID("messageLinkInfo"); err != nil {
				return fmt.Errorf("unable to decode messageLinkInfo#c57d442a: %w", err)
			}
		case "is_public":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode messageLinkInfo#c57d442a: field is_public: %w", err)
			}
			m.IsPublic = value
		case "chat_id":
			value, err := b.Long()
			if err != nil {
				return fmt.Errorf("unable to decode messageLinkInfo#c57d442a: field chat_id: %w", err)
			}
			m.ChatID = value
		case "message":
			if err := m.Message.DecodeTDLibJSON(b); err != nil {
				return fmt.Errorf("unable to decode messageLinkInfo#c57d442a: field message: %w", err)
			}
		case "media_timestamp":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode messageLinkInfo#c57d442a: field media_timestamp: %w", err)
			}
			m.MediaTimestamp = value
		case "for_album":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode messageLinkInfo#c57d442a: field for_album: %w", err)
			}
			m.ForAlbum = value
		case "for_comment":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode messageLinkInfo#c57d442a: field for_comment: %w", err)
			}
			m.ForComment = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetIsPublic returns value of IsPublic field.
func (m *MessageLinkInfo) GetIsPublic() (value bool) {
	return m.IsPublic
}

// GetChatID returns value of ChatID field.
func (m *MessageLinkInfo) GetChatID() (value int64) {
	return m.ChatID
}

// GetMessage returns value of Message field.
func (m *MessageLinkInfo) GetMessage() (value Message) {
	return m.Message
}

// GetMediaTimestamp returns value of MediaTimestamp field.
func (m *MessageLinkInfo) GetMediaTimestamp() (value int32) {
	return m.MediaTimestamp
}

// GetForAlbum returns value of ForAlbum field.
func (m *MessageLinkInfo) GetForAlbum() (value bool) {
	return m.ForAlbum
}

// GetForComment returns value of ForComment field.
func (m *MessageLinkInfo) GetForComment() (value bool) {
	return m.ForComment
}
