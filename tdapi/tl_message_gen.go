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

// Message represents TL type `message#a9c04bbc`.
type Message struct {
	// Message identifier; unique for the chat to which the message belongs
	ID int64
	// Identifier of the sender of the message
	SenderID MessageSenderClass
	// Chat identifier
	ChatID int64
	// The sending state of the message; may be null if the message isn't being sent and
	// didn't fail to be sent
	SendingState MessageSendingStateClass
	// The scheduling state of the message; may be null if the message isn't scheduled
	SchedulingState MessageSchedulingStateClass
	// True, if the message is outgoing
	IsOutgoing bool
	// True, if the message is pinned
	IsPinned bool
	// True, if the message was sent because of a scheduled action by the message sender, for
	// example, as away, or greeting service message
	IsFromOffline bool
	// True, if content of the message can be saved locally or copied using
	// inputMessageForwarded or forwardMessages with copy options
	CanBeSaved bool
	// True, if media timestamp entities refers to a media in this message as opposed to a
	// media in the replied message
	HasTimestampedMedia bool
	// True, if the message is a channel post. All messages to channels are channel posts,
	// all other messages are not channel posts
	IsChannelPost bool
	// True, if the message is a forum topic message
	IsTopicMessage bool
	// True, if the message contains an unread mention for the current user
	ContainsUnreadMention bool
	// Point in time (Unix timestamp) when the message was sent
	Date int32
	// Point in time (Unix timestamp) when the message was last edited
	EditDate int32
	// Information about the initial message sender; may be null if none or unknown
	ForwardInfo MessageForwardInfo
	// Information about the initial message for messages created with importMessages; may be
	// null if the message isn't imported
	ImportInfo MessageImportInfo
	// Information about interactions with the message; may be null if none
	InteractionInfo MessageInteractionInfo
	// Information about unread reactions added to the message
	UnreadReactions []UnreadReaction
	// Information about fact-check added to the message; may be null if none
	FactCheck FactCheck
	// Information about the message or the story this message is replying to; may be null if
	// none
	ReplyTo MessageReplyToClass
	// If non-zero, the identifier of the message thread the message belongs to; unique
	// within the chat to which the message belongs
	MessageThreadID int64
	// Identifier of the Saved Messages topic for the message; 0 for messages not from Saved
	// Messages
	SavedMessagesTopicID int64
	// The message's self-destruct type; may be null if none
	SelfDestructType MessageSelfDestructTypeClass
	// Time left before the message self-destruct timer expires, in seconds; 0 if
	// self-destruction isn't scheduled yet
	SelfDestructIn float64
	// Time left before the message will be automatically deleted by message_auto_delete_time
	// setting of the chat, in seconds; 0 if never
	AutoDeleteIn float64
	// If non-zero, the user identifier of the inline bot through which this message was sent
	ViaBotUserID int64
	// If non-zero, the user identifier of the business bot that sent this message
	SenderBusinessBotUserID int64
	// Number of times the sender of the message boosted the supergroup at the time the
	// message was sent; 0 if none or unknown. For messages sent by the current user,
	// supergroupFullInfo.my_boost_count must be used instead
	SenderBoostCount int32
	// For channel posts and anonymous group messages, optional author signature
	AuthorSignature string
	// Unique identifier of an album this message belongs to; 0 if none. Only audios,
	// documents, photos and videos can be grouped together in albums
	MediaAlbumID int64
	// Unique identifier of the effect added to the message; 0 if none
	EffectID int64
	// True, if media content of the message must be hidden with 18+ spoiler
	HasSensitiveContent bool
	// If non-empty, contains a human-readable description of the reason why access to this
	// message must be restricted
	RestrictionReason string
	// Content of the message
	Content MessageContentClass
	// Reply markup for the message; may be null if none
	ReplyMarkup ReplyMarkupClass
}

// MessageTypeID is TL type id of Message.
const MessageTypeID = 0xa9c04bbc

// Ensuring interfaces in compile-time for Message.
var (
	_ bin.Encoder     = &Message{}
	_ bin.Decoder     = &Message{}
	_ bin.BareEncoder = &Message{}
	_ bin.BareDecoder = &Message{}
)

func (m *Message) Zero() bool {
	if m == nil {
		return true
	}
	if !(m.ID == 0) {
		return false
	}
	if !(m.SenderID == nil) {
		return false
	}
	if !(m.ChatID == 0) {
		return false
	}
	if !(m.SendingState == nil) {
		return false
	}
	if !(m.SchedulingState == nil) {
		return false
	}
	if !(m.IsOutgoing == false) {
		return false
	}
	if !(m.IsPinned == false) {
		return false
	}
	if !(m.IsFromOffline == false) {
		return false
	}
	if !(m.CanBeSaved == false) {
		return false
	}
	if !(m.HasTimestampedMedia == false) {
		return false
	}
	if !(m.IsChannelPost == false) {
		return false
	}
	if !(m.IsTopicMessage == false) {
		return false
	}
	if !(m.ContainsUnreadMention == false) {
		return false
	}
	if !(m.Date == 0) {
		return false
	}
	if !(m.EditDate == 0) {
		return false
	}
	if !(m.ForwardInfo.Zero()) {
		return false
	}
	if !(m.ImportInfo.Zero()) {
		return false
	}
	if !(m.InteractionInfo.Zero()) {
		return false
	}
	if !(m.UnreadReactions == nil) {
		return false
	}
	if !(m.FactCheck.Zero()) {
		return false
	}
	if !(m.ReplyTo == nil) {
		return false
	}
	if !(m.MessageThreadID == 0) {
		return false
	}
	if !(m.SavedMessagesTopicID == 0) {
		return false
	}
	if !(m.SelfDestructType == nil) {
		return false
	}
	if !(m.SelfDestructIn == 0) {
		return false
	}
	if !(m.AutoDeleteIn == 0) {
		return false
	}
	if !(m.ViaBotUserID == 0) {
		return false
	}
	if !(m.SenderBusinessBotUserID == 0) {
		return false
	}
	if !(m.SenderBoostCount == 0) {
		return false
	}
	if !(m.AuthorSignature == "") {
		return false
	}
	if !(m.MediaAlbumID == 0) {
		return false
	}
	if !(m.EffectID == 0) {
		return false
	}
	if !(m.HasSensitiveContent == false) {
		return false
	}
	if !(m.RestrictionReason == "") {
		return false
	}
	if !(m.Content == nil) {
		return false
	}
	if !(m.ReplyMarkup == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (m *Message) String() string {
	if m == nil {
		return "Message(nil)"
	}
	type Alias Message
	return fmt.Sprintf("Message%+v", Alias(*m))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*Message) TypeID() uint32 {
	return MessageTypeID
}

// TypeName returns name of type in TL schema.
func (*Message) TypeName() string {
	return "message"
}

// TypeInfo returns info about TL type.
func (m *Message) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "message",
		ID:   MessageTypeID,
	}
	if m == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ID",
			SchemaName: "id",
		},
		{
			Name:       "SenderID",
			SchemaName: "sender_id",
		},
		{
			Name:       "ChatID",
			SchemaName: "chat_id",
		},
		{
			Name:       "SendingState",
			SchemaName: "sending_state",
		},
		{
			Name:       "SchedulingState",
			SchemaName: "scheduling_state",
		},
		{
			Name:       "IsOutgoing",
			SchemaName: "is_outgoing",
		},
		{
			Name:       "IsPinned",
			SchemaName: "is_pinned",
		},
		{
			Name:       "IsFromOffline",
			SchemaName: "is_from_offline",
		},
		{
			Name:       "CanBeSaved",
			SchemaName: "can_be_saved",
		},
		{
			Name:       "HasTimestampedMedia",
			SchemaName: "has_timestamped_media",
		},
		{
			Name:       "IsChannelPost",
			SchemaName: "is_channel_post",
		},
		{
			Name:       "IsTopicMessage",
			SchemaName: "is_topic_message",
		},
		{
			Name:       "ContainsUnreadMention",
			SchemaName: "contains_unread_mention",
		},
		{
			Name:       "Date",
			SchemaName: "date",
		},
		{
			Name:       "EditDate",
			SchemaName: "edit_date",
		},
		{
			Name:       "ForwardInfo",
			SchemaName: "forward_info",
		},
		{
			Name:       "ImportInfo",
			SchemaName: "import_info",
		},
		{
			Name:       "InteractionInfo",
			SchemaName: "interaction_info",
		},
		{
			Name:       "UnreadReactions",
			SchemaName: "unread_reactions",
		},
		{
			Name:       "FactCheck",
			SchemaName: "fact_check",
		},
		{
			Name:       "ReplyTo",
			SchemaName: "reply_to",
		},
		{
			Name:       "MessageThreadID",
			SchemaName: "message_thread_id",
		},
		{
			Name:       "SavedMessagesTopicID",
			SchemaName: "saved_messages_topic_id",
		},
		{
			Name:       "SelfDestructType",
			SchemaName: "self_destruct_type",
		},
		{
			Name:       "SelfDestructIn",
			SchemaName: "self_destruct_in",
		},
		{
			Name:       "AutoDeleteIn",
			SchemaName: "auto_delete_in",
		},
		{
			Name:       "ViaBotUserID",
			SchemaName: "via_bot_user_id",
		},
		{
			Name:       "SenderBusinessBotUserID",
			SchemaName: "sender_business_bot_user_id",
		},
		{
			Name:       "SenderBoostCount",
			SchemaName: "sender_boost_count",
		},
		{
			Name:       "AuthorSignature",
			SchemaName: "author_signature",
		},
		{
			Name:       "MediaAlbumID",
			SchemaName: "media_album_id",
		},
		{
			Name:       "EffectID",
			SchemaName: "effect_id",
		},
		{
			Name:       "HasSensitiveContent",
			SchemaName: "has_sensitive_content",
		},
		{
			Name:       "RestrictionReason",
			SchemaName: "restriction_reason",
		},
		{
			Name:       "Content",
			SchemaName: "content",
		},
		{
			Name:       "ReplyMarkup",
			SchemaName: "reply_markup",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (m *Message) Encode(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't encode message#a9c04bbc as nil")
	}
	b.PutID(MessageTypeID)
	return m.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (m *Message) EncodeBare(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't encode message#a9c04bbc as nil")
	}
	b.PutInt53(m.ID)
	if m.SenderID == nil {
		return fmt.Errorf("unable to encode message#a9c04bbc: field sender_id is nil")
	}
	if err := m.SenderID.Encode(b); err != nil {
		return fmt.Errorf("unable to encode message#a9c04bbc: field sender_id: %w", err)
	}
	b.PutInt53(m.ChatID)
	if m.SendingState == nil {
		return fmt.Errorf("unable to encode message#a9c04bbc: field sending_state is nil")
	}
	if err := m.SendingState.Encode(b); err != nil {
		return fmt.Errorf("unable to encode message#a9c04bbc: field sending_state: %w", err)
	}
	if m.SchedulingState == nil {
		return fmt.Errorf("unable to encode message#a9c04bbc: field scheduling_state is nil")
	}
	if err := m.SchedulingState.Encode(b); err != nil {
		return fmt.Errorf("unable to encode message#a9c04bbc: field scheduling_state: %w", err)
	}
	b.PutBool(m.IsOutgoing)
	b.PutBool(m.IsPinned)
	b.PutBool(m.IsFromOffline)
	b.PutBool(m.CanBeSaved)
	b.PutBool(m.HasTimestampedMedia)
	b.PutBool(m.IsChannelPost)
	b.PutBool(m.IsTopicMessage)
	b.PutBool(m.ContainsUnreadMention)
	b.PutInt32(m.Date)
	b.PutInt32(m.EditDate)
	if err := m.ForwardInfo.Encode(b); err != nil {
		return fmt.Errorf("unable to encode message#a9c04bbc: field forward_info: %w", err)
	}
	if err := m.ImportInfo.Encode(b); err != nil {
		return fmt.Errorf("unable to encode message#a9c04bbc: field import_info: %w", err)
	}
	if err := m.InteractionInfo.Encode(b); err != nil {
		return fmt.Errorf("unable to encode message#a9c04bbc: field interaction_info: %w", err)
	}
	b.PutInt(len(m.UnreadReactions))
	for idx, v := range m.UnreadReactions {
		if err := v.EncodeBare(b); err != nil {
			return fmt.Errorf("unable to encode bare message#a9c04bbc: field unread_reactions element with index %d: %w", idx, err)
		}
	}
	if err := m.FactCheck.Encode(b); err != nil {
		return fmt.Errorf("unable to encode message#a9c04bbc: field fact_check: %w", err)
	}
	if m.ReplyTo == nil {
		return fmt.Errorf("unable to encode message#a9c04bbc: field reply_to is nil")
	}
	if err := m.ReplyTo.Encode(b); err != nil {
		return fmt.Errorf("unable to encode message#a9c04bbc: field reply_to: %w", err)
	}
	b.PutInt53(m.MessageThreadID)
	b.PutInt53(m.SavedMessagesTopicID)
	if m.SelfDestructType == nil {
		return fmt.Errorf("unable to encode message#a9c04bbc: field self_destruct_type is nil")
	}
	if err := m.SelfDestructType.Encode(b); err != nil {
		return fmt.Errorf("unable to encode message#a9c04bbc: field self_destruct_type: %w", err)
	}
	b.PutDouble(m.SelfDestructIn)
	b.PutDouble(m.AutoDeleteIn)
	b.PutInt53(m.ViaBotUserID)
	b.PutInt53(m.SenderBusinessBotUserID)
	b.PutInt32(m.SenderBoostCount)
	b.PutString(m.AuthorSignature)
	b.PutLong(m.MediaAlbumID)
	b.PutLong(m.EffectID)
	b.PutBool(m.HasSensitiveContent)
	b.PutString(m.RestrictionReason)
	if m.Content == nil {
		return fmt.Errorf("unable to encode message#a9c04bbc: field content is nil")
	}
	if err := m.Content.Encode(b); err != nil {
		return fmt.Errorf("unable to encode message#a9c04bbc: field content: %w", err)
	}
	if m.ReplyMarkup == nil {
		return fmt.Errorf("unable to encode message#a9c04bbc: field reply_markup is nil")
	}
	if err := m.ReplyMarkup.Encode(b); err != nil {
		return fmt.Errorf("unable to encode message#a9c04bbc: field reply_markup: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (m *Message) Decode(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't decode message#a9c04bbc to nil")
	}
	if err := b.ConsumeID(MessageTypeID); err != nil {
		return fmt.Errorf("unable to decode message#a9c04bbc: %w", err)
	}
	return m.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (m *Message) DecodeBare(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't decode message#a9c04bbc to nil")
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode message#a9c04bbc: field id: %w", err)
		}
		m.ID = value
	}
	{
		value, err := DecodeMessageSender(b)
		if err != nil {
			return fmt.Errorf("unable to decode message#a9c04bbc: field sender_id: %w", err)
		}
		m.SenderID = value
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode message#a9c04bbc: field chat_id: %w", err)
		}
		m.ChatID = value
	}
	{
		value, err := DecodeMessageSendingState(b)
		if err != nil {
			return fmt.Errorf("unable to decode message#a9c04bbc: field sending_state: %w", err)
		}
		m.SendingState = value
	}
	{
		value, err := DecodeMessageSchedulingState(b)
		if err != nil {
			return fmt.Errorf("unable to decode message#a9c04bbc: field scheduling_state: %w", err)
		}
		m.SchedulingState = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode message#a9c04bbc: field is_outgoing: %w", err)
		}
		m.IsOutgoing = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode message#a9c04bbc: field is_pinned: %w", err)
		}
		m.IsPinned = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode message#a9c04bbc: field is_from_offline: %w", err)
		}
		m.IsFromOffline = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode message#a9c04bbc: field can_be_saved: %w", err)
		}
		m.CanBeSaved = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode message#a9c04bbc: field has_timestamped_media: %w", err)
		}
		m.HasTimestampedMedia = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode message#a9c04bbc: field is_channel_post: %w", err)
		}
		m.IsChannelPost = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode message#a9c04bbc: field is_topic_message: %w", err)
		}
		m.IsTopicMessage = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode message#a9c04bbc: field contains_unread_mention: %w", err)
		}
		m.ContainsUnreadMention = value
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode message#a9c04bbc: field date: %w", err)
		}
		m.Date = value
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode message#a9c04bbc: field edit_date: %w", err)
		}
		m.EditDate = value
	}
	{
		if err := m.ForwardInfo.Decode(b); err != nil {
			return fmt.Errorf("unable to decode message#a9c04bbc: field forward_info: %w", err)
		}
	}
	{
		if err := m.ImportInfo.Decode(b); err != nil {
			return fmt.Errorf("unable to decode message#a9c04bbc: field import_info: %w", err)
		}
	}
	{
		if err := m.InteractionInfo.Decode(b); err != nil {
			return fmt.Errorf("unable to decode message#a9c04bbc: field interaction_info: %w", err)
		}
	}
	{
		headerLen, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode message#a9c04bbc: field unread_reactions: %w", err)
		}

		if headerLen > 0 {
			m.UnreadReactions = make([]UnreadReaction, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			var value UnreadReaction
			if err := value.DecodeBare(b); err != nil {
				return fmt.Errorf("unable to decode bare message#a9c04bbc: field unread_reactions: %w", err)
			}
			m.UnreadReactions = append(m.UnreadReactions, value)
		}
	}
	{
		if err := m.FactCheck.Decode(b); err != nil {
			return fmt.Errorf("unable to decode message#a9c04bbc: field fact_check: %w", err)
		}
	}
	{
		value, err := DecodeMessageReplyTo(b)
		if err != nil {
			return fmt.Errorf("unable to decode message#a9c04bbc: field reply_to: %w", err)
		}
		m.ReplyTo = value
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode message#a9c04bbc: field message_thread_id: %w", err)
		}
		m.MessageThreadID = value
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode message#a9c04bbc: field saved_messages_topic_id: %w", err)
		}
		m.SavedMessagesTopicID = value
	}
	{
		value, err := DecodeMessageSelfDestructType(b)
		if err != nil {
			return fmt.Errorf("unable to decode message#a9c04bbc: field self_destruct_type: %w", err)
		}
		m.SelfDestructType = value
	}
	{
		value, err := b.Double()
		if err != nil {
			return fmt.Errorf("unable to decode message#a9c04bbc: field self_destruct_in: %w", err)
		}
		m.SelfDestructIn = value
	}
	{
		value, err := b.Double()
		if err != nil {
			return fmt.Errorf("unable to decode message#a9c04bbc: field auto_delete_in: %w", err)
		}
		m.AutoDeleteIn = value
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode message#a9c04bbc: field via_bot_user_id: %w", err)
		}
		m.ViaBotUserID = value
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode message#a9c04bbc: field sender_business_bot_user_id: %w", err)
		}
		m.SenderBusinessBotUserID = value
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode message#a9c04bbc: field sender_boost_count: %w", err)
		}
		m.SenderBoostCount = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode message#a9c04bbc: field author_signature: %w", err)
		}
		m.AuthorSignature = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode message#a9c04bbc: field media_album_id: %w", err)
		}
		m.MediaAlbumID = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode message#a9c04bbc: field effect_id: %w", err)
		}
		m.EffectID = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode message#a9c04bbc: field has_sensitive_content: %w", err)
		}
		m.HasSensitiveContent = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode message#a9c04bbc: field restriction_reason: %w", err)
		}
		m.RestrictionReason = value
	}
	{
		value, err := DecodeMessageContent(b)
		if err != nil {
			return fmt.Errorf("unable to decode message#a9c04bbc: field content: %w", err)
		}
		m.Content = value
	}
	{
		value, err := DecodeReplyMarkup(b)
		if err != nil {
			return fmt.Errorf("unable to decode message#a9c04bbc: field reply_markup: %w", err)
		}
		m.ReplyMarkup = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (m *Message) EncodeTDLibJSON(b tdjson.Encoder) error {
	if m == nil {
		return fmt.Errorf("can't encode message#a9c04bbc as nil")
	}
	b.ObjStart()
	b.PutID("message")
	b.Comma()
	b.FieldStart("id")
	b.PutInt53(m.ID)
	b.Comma()
	b.FieldStart("sender_id")
	if m.SenderID == nil {
		return fmt.Errorf("unable to encode message#a9c04bbc: field sender_id is nil")
	}
	if err := m.SenderID.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode message#a9c04bbc: field sender_id: %w", err)
	}
	b.Comma()
	b.FieldStart("chat_id")
	b.PutInt53(m.ChatID)
	b.Comma()
	b.FieldStart("sending_state")
	if m.SendingState == nil {
		return fmt.Errorf("unable to encode message#a9c04bbc: field sending_state is nil")
	}
	if err := m.SendingState.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode message#a9c04bbc: field sending_state: %w", err)
	}
	b.Comma()
	b.FieldStart("scheduling_state")
	if m.SchedulingState == nil {
		return fmt.Errorf("unable to encode message#a9c04bbc: field scheduling_state is nil")
	}
	if err := m.SchedulingState.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode message#a9c04bbc: field scheduling_state: %w", err)
	}
	b.Comma()
	b.FieldStart("is_outgoing")
	b.PutBool(m.IsOutgoing)
	b.Comma()
	b.FieldStart("is_pinned")
	b.PutBool(m.IsPinned)
	b.Comma()
	b.FieldStart("is_from_offline")
	b.PutBool(m.IsFromOffline)
	b.Comma()
	b.FieldStart("can_be_saved")
	b.PutBool(m.CanBeSaved)
	b.Comma()
	b.FieldStart("has_timestamped_media")
	b.PutBool(m.HasTimestampedMedia)
	b.Comma()
	b.FieldStart("is_channel_post")
	b.PutBool(m.IsChannelPost)
	b.Comma()
	b.FieldStart("is_topic_message")
	b.PutBool(m.IsTopicMessage)
	b.Comma()
	b.FieldStart("contains_unread_mention")
	b.PutBool(m.ContainsUnreadMention)
	b.Comma()
	b.FieldStart("date")
	b.PutInt32(m.Date)
	b.Comma()
	b.FieldStart("edit_date")
	b.PutInt32(m.EditDate)
	b.Comma()
	b.FieldStart("forward_info")
	if err := m.ForwardInfo.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode message#a9c04bbc: field forward_info: %w", err)
	}
	b.Comma()
	b.FieldStart("import_info")
	if err := m.ImportInfo.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode message#a9c04bbc: field import_info: %w", err)
	}
	b.Comma()
	b.FieldStart("interaction_info")
	if err := m.InteractionInfo.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode message#a9c04bbc: field interaction_info: %w", err)
	}
	b.Comma()
	b.FieldStart("unread_reactions")
	b.ArrStart()
	for idx, v := range m.UnreadReactions {
		if err := v.EncodeTDLibJSON(b); err != nil {
			return fmt.Errorf("unable to encode message#a9c04bbc: field unread_reactions element with index %d: %w", idx, err)
		}
		b.Comma()
	}
	b.StripComma()
	b.ArrEnd()
	b.Comma()
	b.FieldStart("fact_check")
	if err := m.FactCheck.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode message#a9c04bbc: field fact_check: %w", err)
	}
	b.Comma()
	b.FieldStart("reply_to")
	if m.ReplyTo == nil {
		return fmt.Errorf("unable to encode message#a9c04bbc: field reply_to is nil")
	}
	if err := m.ReplyTo.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode message#a9c04bbc: field reply_to: %w", err)
	}
	b.Comma()
	b.FieldStart("message_thread_id")
	b.PutInt53(m.MessageThreadID)
	b.Comma()
	b.FieldStart("saved_messages_topic_id")
	b.PutInt53(m.SavedMessagesTopicID)
	b.Comma()
	b.FieldStart("self_destruct_type")
	if m.SelfDestructType == nil {
		return fmt.Errorf("unable to encode message#a9c04bbc: field self_destruct_type is nil")
	}
	if err := m.SelfDestructType.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode message#a9c04bbc: field self_destruct_type: %w", err)
	}
	b.Comma()
	b.FieldStart("self_destruct_in")
	b.PutDouble(m.SelfDestructIn)
	b.Comma()
	b.FieldStart("auto_delete_in")
	b.PutDouble(m.AutoDeleteIn)
	b.Comma()
	b.FieldStart("via_bot_user_id")
	b.PutInt53(m.ViaBotUserID)
	b.Comma()
	b.FieldStart("sender_business_bot_user_id")
	b.PutInt53(m.SenderBusinessBotUserID)
	b.Comma()
	b.FieldStart("sender_boost_count")
	b.PutInt32(m.SenderBoostCount)
	b.Comma()
	b.FieldStart("author_signature")
	b.PutString(m.AuthorSignature)
	b.Comma()
	b.FieldStart("media_album_id")
	b.PutLong(m.MediaAlbumID)
	b.Comma()
	b.FieldStart("effect_id")
	b.PutLong(m.EffectID)
	b.Comma()
	b.FieldStart("has_sensitive_content")
	b.PutBool(m.HasSensitiveContent)
	b.Comma()
	b.FieldStart("restriction_reason")
	b.PutString(m.RestrictionReason)
	b.Comma()
	b.FieldStart("content")
	if m.Content == nil {
		return fmt.Errorf("unable to encode message#a9c04bbc: field content is nil")
	}
	if err := m.Content.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode message#a9c04bbc: field content: %w", err)
	}
	b.Comma()
	b.FieldStart("reply_markup")
	if m.ReplyMarkup == nil {
		return fmt.Errorf("unable to encode message#a9c04bbc: field reply_markup is nil")
	}
	if err := m.ReplyMarkup.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode message#a9c04bbc: field reply_markup: %w", err)
	}
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (m *Message) DecodeTDLibJSON(b tdjson.Decoder) error {
	if m == nil {
		return fmt.Errorf("can't decode message#a9c04bbc to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("message"); err != nil {
				return fmt.Errorf("unable to decode message#a9c04bbc: %w", err)
			}
		case "id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode message#a9c04bbc: field id: %w", err)
			}
			m.ID = value
		case "sender_id":
			value, err := DecodeTDLibJSONMessageSender(b)
			if err != nil {
				return fmt.Errorf("unable to decode message#a9c04bbc: field sender_id: %w", err)
			}
			m.SenderID = value
		case "chat_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode message#a9c04bbc: field chat_id: %w", err)
			}
			m.ChatID = value
		case "sending_state":
			value, err := DecodeTDLibJSONMessageSendingState(b)
			if err != nil {
				return fmt.Errorf("unable to decode message#a9c04bbc: field sending_state: %w", err)
			}
			m.SendingState = value
		case "scheduling_state":
			value, err := DecodeTDLibJSONMessageSchedulingState(b)
			if err != nil {
				return fmt.Errorf("unable to decode message#a9c04bbc: field scheduling_state: %w", err)
			}
			m.SchedulingState = value
		case "is_outgoing":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode message#a9c04bbc: field is_outgoing: %w", err)
			}
			m.IsOutgoing = value
		case "is_pinned":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode message#a9c04bbc: field is_pinned: %w", err)
			}
			m.IsPinned = value
		case "is_from_offline":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode message#a9c04bbc: field is_from_offline: %w", err)
			}
			m.IsFromOffline = value
		case "can_be_saved":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode message#a9c04bbc: field can_be_saved: %w", err)
			}
			m.CanBeSaved = value
		case "has_timestamped_media":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode message#a9c04bbc: field has_timestamped_media: %w", err)
			}
			m.HasTimestampedMedia = value
		case "is_channel_post":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode message#a9c04bbc: field is_channel_post: %w", err)
			}
			m.IsChannelPost = value
		case "is_topic_message":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode message#a9c04bbc: field is_topic_message: %w", err)
			}
			m.IsTopicMessage = value
		case "contains_unread_mention":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode message#a9c04bbc: field contains_unread_mention: %w", err)
			}
			m.ContainsUnreadMention = value
		case "date":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode message#a9c04bbc: field date: %w", err)
			}
			m.Date = value
		case "edit_date":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode message#a9c04bbc: field edit_date: %w", err)
			}
			m.EditDate = value
		case "forward_info":
			if err := m.ForwardInfo.DecodeTDLibJSON(b); err != nil {
				return fmt.Errorf("unable to decode message#a9c04bbc: field forward_info: %w", err)
			}
		case "import_info":
			if err := m.ImportInfo.DecodeTDLibJSON(b); err != nil {
				return fmt.Errorf("unable to decode message#a9c04bbc: field import_info: %w", err)
			}
		case "interaction_info":
			if err := m.InteractionInfo.DecodeTDLibJSON(b); err != nil {
				return fmt.Errorf("unable to decode message#a9c04bbc: field interaction_info: %w", err)
			}
		case "unread_reactions":
			if err := b.Arr(func(b tdjson.Decoder) error {
				var value UnreadReaction
				if err := value.DecodeTDLibJSON(b); err != nil {
					return fmt.Errorf("unable to decode message#a9c04bbc: field unread_reactions: %w", err)
				}
				m.UnreadReactions = append(m.UnreadReactions, value)
				return nil
			}); err != nil {
				return fmt.Errorf("unable to decode message#a9c04bbc: field unread_reactions: %w", err)
			}
		case "fact_check":
			if err := m.FactCheck.DecodeTDLibJSON(b); err != nil {
				return fmt.Errorf("unable to decode message#a9c04bbc: field fact_check: %w", err)
			}
		case "reply_to":
			value, err := DecodeTDLibJSONMessageReplyTo(b)
			if err != nil {
				return fmt.Errorf("unable to decode message#a9c04bbc: field reply_to: %w", err)
			}
			m.ReplyTo = value
		case "message_thread_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode message#a9c04bbc: field message_thread_id: %w", err)
			}
			m.MessageThreadID = value
		case "saved_messages_topic_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode message#a9c04bbc: field saved_messages_topic_id: %w", err)
			}
			m.SavedMessagesTopicID = value
		case "self_destruct_type":
			value, err := DecodeTDLibJSONMessageSelfDestructType(b)
			if err != nil {
				return fmt.Errorf("unable to decode message#a9c04bbc: field self_destruct_type: %w", err)
			}
			m.SelfDestructType = value
		case "self_destruct_in":
			value, err := b.Double()
			if err != nil {
				return fmt.Errorf("unable to decode message#a9c04bbc: field self_destruct_in: %w", err)
			}
			m.SelfDestructIn = value
		case "auto_delete_in":
			value, err := b.Double()
			if err != nil {
				return fmt.Errorf("unable to decode message#a9c04bbc: field auto_delete_in: %w", err)
			}
			m.AutoDeleteIn = value
		case "via_bot_user_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode message#a9c04bbc: field via_bot_user_id: %w", err)
			}
			m.ViaBotUserID = value
		case "sender_business_bot_user_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode message#a9c04bbc: field sender_business_bot_user_id: %w", err)
			}
			m.SenderBusinessBotUserID = value
		case "sender_boost_count":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode message#a9c04bbc: field sender_boost_count: %w", err)
			}
			m.SenderBoostCount = value
		case "author_signature":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode message#a9c04bbc: field author_signature: %w", err)
			}
			m.AuthorSignature = value
		case "media_album_id":
			value, err := b.Long()
			if err != nil {
				return fmt.Errorf("unable to decode message#a9c04bbc: field media_album_id: %w", err)
			}
			m.MediaAlbumID = value
		case "effect_id":
			value, err := b.Long()
			if err != nil {
				return fmt.Errorf("unable to decode message#a9c04bbc: field effect_id: %w", err)
			}
			m.EffectID = value
		case "has_sensitive_content":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode message#a9c04bbc: field has_sensitive_content: %w", err)
			}
			m.HasSensitiveContent = value
		case "restriction_reason":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode message#a9c04bbc: field restriction_reason: %w", err)
			}
			m.RestrictionReason = value
		case "content":
			value, err := DecodeTDLibJSONMessageContent(b)
			if err != nil {
				return fmt.Errorf("unable to decode message#a9c04bbc: field content: %w", err)
			}
			m.Content = value
		case "reply_markup":
			value, err := DecodeTDLibJSONReplyMarkup(b)
			if err != nil {
				return fmt.Errorf("unable to decode message#a9c04bbc: field reply_markup: %w", err)
			}
			m.ReplyMarkup = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetID returns value of ID field.
func (m *Message) GetID() (value int64) {
	if m == nil {
		return
	}
	return m.ID
}

// GetSenderID returns value of SenderID field.
func (m *Message) GetSenderID() (value MessageSenderClass) {
	if m == nil {
		return
	}
	return m.SenderID
}

// GetChatID returns value of ChatID field.
func (m *Message) GetChatID() (value int64) {
	if m == nil {
		return
	}
	return m.ChatID
}

// GetSendingState returns value of SendingState field.
func (m *Message) GetSendingState() (value MessageSendingStateClass) {
	if m == nil {
		return
	}
	return m.SendingState
}

// GetSchedulingState returns value of SchedulingState field.
func (m *Message) GetSchedulingState() (value MessageSchedulingStateClass) {
	if m == nil {
		return
	}
	return m.SchedulingState
}

// GetIsOutgoing returns value of IsOutgoing field.
func (m *Message) GetIsOutgoing() (value bool) {
	if m == nil {
		return
	}
	return m.IsOutgoing
}

// GetIsPinned returns value of IsPinned field.
func (m *Message) GetIsPinned() (value bool) {
	if m == nil {
		return
	}
	return m.IsPinned
}

// GetIsFromOffline returns value of IsFromOffline field.
func (m *Message) GetIsFromOffline() (value bool) {
	if m == nil {
		return
	}
	return m.IsFromOffline
}

// GetCanBeSaved returns value of CanBeSaved field.
func (m *Message) GetCanBeSaved() (value bool) {
	if m == nil {
		return
	}
	return m.CanBeSaved
}

// GetHasTimestampedMedia returns value of HasTimestampedMedia field.
func (m *Message) GetHasTimestampedMedia() (value bool) {
	if m == nil {
		return
	}
	return m.HasTimestampedMedia
}

// GetIsChannelPost returns value of IsChannelPost field.
func (m *Message) GetIsChannelPost() (value bool) {
	if m == nil {
		return
	}
	return m.IsChannelPost
}

// GetIsTopicMessage returns value of IsTopicMessage field.
func (m *Message) GetIsTopicMessage() (value bool) {
	if m == nil {
		return
	}
	return m.IsTopicMessage
}

// GetContainsUnreadMention returns value of ContainsUnreadMention field.
func (m *Message) GetContainsUnreadMention() (value bool) {
	if m == nil {
		return
	}
	return m.ContainsUnreadMention
}

// GetDate returns value of Date field.
func (m *Message) GetDate() (value int32) {
	if m == nil {
		return
	}
	return m.Date
}

// GetEditDate returns value of EditDate field.
func (m *Message) GetEditDate() (value int32) {
	if m == nil {
		return
	}
	return m.EditDate
}

// GetForwardInfo returns value of ForwardInfo field.
func (m *Message) GetForwardInfo() (value MessageForwardInfo) {
	if m == nil {
		return
	}
	return m.ForwardInfo
}

// GetImportInfo returns value of ImportInfo field.
func (m *Message) GetImportInfo() (value MessageImportInfo) {
	if m == nil {
		return
	}
	return m.ImportInfo
}

// GetInteractionInfo returns value of InteractionInfo field.
func (m *Message) GetInteractionInfo() (value MessageInteractionInfo) {
	if m == nil {
		return
	}
	return m.InteractionInfo
}

// GetUnreadReactions returns value of UnreadReactions field.
func (m *Message) GetUnreadReactions() (value []UnreadReaction) {
	if m == nil {
		return
	}
	return m.UnreadReactions
}

// GetFactCheck returns value of FactCheck field.
func (m *Message) GetFactCheck() (value FactCheck) {
	if m == nil {
		return
	}
	return m.FactCheck
}

// GetReplyTo returns value of ReplyTo field.
func (m *Message) GetReplyTo() (value MessageReplyToClass) {
	if m == nil {
		return
	}
	return m.ReplyTo
}

// GetMessageThreadID returns value of MessageThreadID field.
func (m *Message) GetMessageThreadID() (value int64) {
	if m == nil {
		return
	}
	return m.MessageThreadID
}

// GetSavedMessagesTopicID returns value of SavedMessagesTopicID field.
func (m *Message) GetSavedMessagesTopicID() (value int64) {
	if m == nil {
		return
	}
	return m.SavedMessagesTopicID
}

// GetSelfDestructType returns value of SelfDestructType field.
func (m *Message) GetSelfDestructType() (value MessageSelfDestructTypeClass) {
	if m == nil {
		return
	}
	return m.SelfDestructType
}

// GetSelfDestructIn returns value of SelfDestructIn field.
func (m *Message) GetSelfDestructIn() (value float64) {
	if m == nil {
		return
	}
	return m.SelfDestructIn
}

// GetAutoDeleteIn returns value of AutoDeleteIn field.
func (m *Message) GetAutoDeleteIn() (value float64) {
	if m == nil {
		return
	}
	return m.AutoDeleteIn
}

// GetViaBotUserID returns value of ViaBotUserID field.
func (m *Message) GetViaBotUserID() (value int64) {
	if m == nil {
		return
	}
	return m.ViaBotUserID
}

// GetSenderBusinessBotUserID returns value of SenderBusinessBotUserID field.
func (m *Message) GetSenderBusinessBotUserID() (value int64) {
	if m == nil {
		return
	}
	return m.SenderBusinessBotUserID
}

// GetSenderBoostCount returns value of SenderBoostCount field.
func (m *Message) GetSenderBoostCount() (value int32) {
	if m == nil {
		return
	}
	return m.SenderBoostCount
}

// GetAuthorSignature returns value of AuthorSignature field.
func (m *Message) GetAuthorSignature() (value string) {
	if m == nil {
		return
	}
	return m.AuthorSignature
}

// GetMediaAlbumID returns value of MediaAlbumID field.
func (m *Message) GetMediaAlbumID() (value int64) {
	if m == nil {
		return
	}
	return m.MediaAlbumID
}

// GetEffectID returns value of EffectID field.
func (m *Message) GetEffectID() (value int64) {
	if m == nil {
		return
	}
	return m.EffectID
}

// GetHasSensitiveContent returns value of HasSensitiveContent field.
func (m *Message) GetHasSensitiveContent() (value bool) {
	if m == nil {
		return
	}
	return m.HasSensitiveContent
}

// GetRestrictionReason returns value of RestrictionReason field.
func (m *Message) GetRestrictionReason() (value string) {
	if m == nil {
		return
	}
	return m.RestrictionReason
}

// GetContent returns value of Content field.
func (m *Message) GetContent() (value MessageContentClass) {
	if m == nil {
		return
	}
	return m.Content
}

// GetReplyMarkup returns value of ReplyMarkup field.
func (m *Message) GetReplyMarkup() (value ReplyMarkupClass) {
	if m == nil {
		return
	}
	return m.ReplyMarkup
}
