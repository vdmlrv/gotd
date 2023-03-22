// Code generated by gotdgen, DO NOT EDIT.

package tg

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

// MessagesForwardMessagesRequest represents TL type `messages.forwardMessages#c661bbc4`.
// Forwards messages by their IDs.
//
// See https://core.telegram.org/method/messages.forwardMessages for reference.
type MessagesForwardMessagesRequest struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// Whether to send messages silently (no notification will be triggered on the
	// destination clients)
	Silent bool
	// Whether to send the message in background
	Background bool
	// When forwarding games, whether to include your score in the game
	WithMyScore bool
	// Whether to forward messages without quoting the original author
	DropAuthor bool
	// Whether to strip captions from media
	DropMediaCaptions bool
	// Only for bots, disallows further re-forwarding and saving of the messages, even if the
	// destination chat doesn't have content protection¹ enabled
	//
	// Links:
	//  1) https://telegram.org/blog/protected-content-delete-by-date-and-more
	Noforwards bool
	// Source of messages
	FromPeer InputPeerClass
	// IDs of messages
	ID []int
	// Random ID to prevent resending of messages
	RandomID []int64
	// Destination peer
	ToPeer InputPeerClass
	//
	//
	// Use SetTopMsgID and GetTopMsgID helpers.
	TopMsgID int
	// Scheduled message date for scheduled messages
	//
	// Use SetScheduleDate and GetScheduleDate helpers.
	ScheduleDate int
	// Forward the messages as the specified peer
	//
	// Use SetSendAs and GetSendAs helpers.
	SendAs InputPeerClass
}

// MessagesForwardMessagesRequestTypeID is TL type id of MessagesForwardMessagesRequest.
const MessagesForwardMessagesRequestTypeID = 0xc661bbc4

// Ensuring interfaces in compile-time for MessagesForwardMessagesRequest.
var (
	_ bin.Encoder     = &MessagesForwardMessagesRequest{}
	_ bin.Decoder     = &MessagesForwardMessagesRequest{}
	_ bin.BareEncoder = &MessagesForwardMessagesRequest{}
	_ bin.BareDecoder = &MessagesForwardMessagesRequest{}
)

func (f *MessagesForwardMessagesRequest) Zero() bool {
	if f == nil {
		return true
	}
	if !(f.Flags.Zero()) {
		return false
	}
	if !(f.Silent == false) {
		return false
	}
	if !(f.Background == false) {
		return false
	}
	if !(f.WithMyScore == false) {
		return false
	}
	if !(f.DropAuthor == false) {
		return false
	}
	if !(f.DropMediaCaptions == false) {
		return false
	}
	if !(f.Noforwards == false) {
		return false
	}
	if !(f.FromPeer == nil) {
		return false
	}
	if !(f.ID == nil) {
		return false
	}
	if !(f.RandomID == nil) {
		return false
	}
	if !(f.ToPeer == nil) {
		return false
	}
	if !(f.TopMsgID == 0) {
		return false
	}
	if !(f.ScheduleDate == 0) {
		return false
	}
	if !(f.SendAs == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (f *MessagesForwardMessagesRequest) String() string {
	if f == nil {
		return "MessagesForwardMessagesRequest(nil)"
	}
	type Alias MessagesForwardMessagesRequest
	return fmt.Sprintf("MessagesForwardMessagesRequest%+v", Alias(*f))
}

// FillFrom fills MessagesForwardMessagesRequest from given interface.
func (f *MessagesForwardMessagesRequest) FillFrom(from interface {
	GetSilent() (value bool)
	GetBackground() (value bool)
	GetWithMyScore() (value bool)
	GetDropAuthor() (value bool)
	GetDropMediaCaptions() (value bool)
	GetNoforwards() (value bool)
	GetFromPeer() (value InputPeerClass)
	GetID() (value []int)
	GetRandomID() (value []int64)
	GetToPeer() (value InputPeerClass)
	GetTopMsgID() (value int, ok bool)
	GetScheduleDate() (value int, ok bool)
	GetSendAs() (value InputPeerClass, ok bool)
}) {
	f.Silent = from.GetSilent()
	f.Background = from.GetBackground()
	f.WithMyScore = from.GetWithMyScore()
	f.DropAuthor = from.GetDropAuthor()
	f.DropMediaCaptions = from.GetDropMediaCaptions()
	f.Noforwards = from.GetNoforwards()
	f.FromPeer = from.GetFromPeer()
	f.ID = from.GetID()
	f.RandomID = from.GetRandomID()
	f.ToPeer = from.GetToPeer()
	if val, ok := from.GetTopMsgID(); ok {
		f.TopMsgID = val
	}

	if val, ok := from.GetScheduleDate(); ok {
		f.ScheduleDate = val
	}

	if val, ok := from.GetSendAs(); ok {
		f.SendAs = val
	}

}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessagesForwardMessagesRequest) TypeID() uint32 {
	return MessagesForwardMessagesRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*MessagesForwardMessagesRequest) TypeName() string {
	return "messages.forwardMessages"
}

// TypeInfo returns info about TL type.
func (f *MessagesForwardMessagesRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messages.forwardMessages",
		ID:   MessagesForwardMessagesRequestTypeID,
	}
	if f == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Silent",
			SchemaName: "silent",
			Null:       !f.Flags.Has(5),
		},
		{
			Name:       "Background",
			SchemaName: "background",
			Null:       !f.Flags.Has(6),
		},
		{
			Name:       "WithMyScore",
			SchemaName: "with_my_score",
			Null:       !f.Flags.Has(8),
		},
		{
			Name:       "DropAuthor",
			SchemaName: "drop_author",
			Null:       !f.Flags.Has(11),
		},
		{
			Name:       "DropMediaCaptions",
			SchemaName: "drop_media_captions",
			Null:       !f.Flags.Has(12),
		},
		{
			Name:       "Noforwards",
			SchemaName: "noforwards",
			Null:       !f.Flags.Has(14),
		},
		{
			Name:       "FromPeer",
			SchemaName: "from_peer",
		},
		{
			Name:       "ID",
			SchemaName: "id",
		},
		{
			Name:       "RandomID",
			SchemaName: "random_id",
		},
		{
			Name:       "ToPeer",
			SchemaName: "to_peer",
		},
		{
			Name:       "TopMsgID",
			SchemaName: "top_msg_id",
			Null:       !f.Flags.Has(9),
		},
		{
			Name:       "ScheduleDate",
			SchemaName: "schedule_date",
			Null:       !f.Flags.Has(10),
		},
		{
			Name:       "SendAs",
			SchemaName: "send_as",
			Null:       !f.Flags.Has(13),
		},
	}
	return typ
}

// SetFlags sets flags for non-zero fields.
func (f *MessagesForwardMessagesRequest) SetFlags() {
	if !(f.Silent == false) {
		f.Flags.Set(5)
	}
	if !(f.Background == false) {
		f.Flags.Set(6)
	}
	if !(f.WithMyScore == false) {
		f.Flags.Set(8)
	}
	if !(f.DropAuthor == false) {
		f.Flags.Set(11)
	}
	if !(f.DropMediaCaptions == false) {
		f.Flags.Set(12)
	}
	if !(f.Noforwards == false) {
		f.Flags.Set(14)
	}
	if !(f.TopMsgID == 0) {
		f.Flags.Set(9)
	}
	if !(f.ScheduleDate == 0) {
		f.Flags.Set(10)
	}
	if !(f.SendAs == nil) {
		f.Flags.Set(13)
	}
}

// Encode implements bin.Encoder.
func (f *MessagesForwardMessagesRequest) Encode(b *bin.Buffer) error {
	if f == nil {
		return fmt.Errorf("can't encode messages.forwardMessages#c661bbc4 as nil")
	}
	b.PutID(MessagesForwardMessagesRequestTypeID)
	return f.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (f *MessagesForwardMessagesRequest) EncodeBare(b *bin.Buffer) error {
	if f == nil {
		return fmt.Errorf("can't encode messages.forwardMessages#c661bbc4 as nil")
	}
	f.SetFlags()
	if err := f.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.forwardMessages#c661bbc4: field flags: %w", err)
	}
	if f.FromPeer == nil {
		return fmt.Errorf("unable to encode messages.forwardMessages#c661bbc4: field from_peer is nil")
	}
	if err := f.FromPeer.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.forwardMessages#c661bbc4: field from_peer: %w", err)
	}
	b.PutVectorHeader(len(f.ID))
	for _, v := range f.ID {
		b.PutInt(v)
	}
	b.PutVectorHeader(len(f.RandomID))
	for _, v := range f.RandomID {
		b.PutLong(v)
	}
	if f.ToPeer == nil {
		return fmt.Errorf("unable to encode messages.forwardMessages#c661bbc4: field to_peer is nil")
	}
	if err := f.ToPeer.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.forwardMessages#c661bbc4: field to_peer: %w", err)
	}
	if f.Flags.Has(9) {
		b.PutInt(f.TopMsgID)
	}
	if f.Flags.Has(10) {
		b.PutInt(f.ScheduleDate)
	}
	if f.Flags.Has(13) {
		if f.SendAs == nil {
			return fmt.Errorf("unable to encode messages.forwardMessages#c661bbc4: field send_as is nil")
		}
		if err := f.SendAs.Encode(b); err != nil {
			return fmt.Errorf("unable to encode messages.forwardMessages#c661bbc4: field send_as: %w", err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (f *MessagesForwardMessagesRequest) Decode(b *bin.Buffer) error {
	if f == nil {
		return fmt.Errorf("can't decode messages.forwardMessages#c661bbc4 to nil")
	}
	if err := b.ConsumeID(MessagesForwardMessagesRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.forwardMessages#c661bbc4: %w", err)
	}
	return f.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (f *MessagesForwardMessagesRequest) DecodeBare(b *bin.Buffer) error {
	if f == nil {
		return fmt.Errorf("can't decode messages.forwardMessages#c661bbc4 to nil")
	}
	{
		if err := f.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode messages.forwardMessages#c661bbc4: field flags: %w", err)
		}
	}
	f.Silent = f.Flags.Has(5)
	f.Background = f.Flags.Has(6)
	f.WithMyScore = f.Flags.Has(8)
	f.DropAuthor = f.Flags.Has(11)
	f.DropMediaCaptions = f.Flags.Has(12)
	f.Noforwards = f.Flags.Has(14)
	{
		value, err := DecodeInputPeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.forwardMessages#c661bbc4: field from_peer: %w", err)
		}
		f.FromPeer = value
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode messages.forwardMessages#c661bbc4: field id: %w", err)
		}

		if headerLen > 0 {
			f.ID = make([]int, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := b.Int()
			if err != nil {
				return fmt.Errorf("unable to decode messages.forwardMessages#c661bbc4: field id: %w", err)
			}
			f.ID = append(f.ID, value)
		}
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode messages.forwardMessages#c661bbc4: field random_id: %w", err)
		}

		if headerLen > 0 {
			f.RandomID = make([]int64, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := b.Long()
			if err != nil {
				return fmt.Errorf("unable to decode messages.forwardMessages#c661bbc4: field random_id: %w", err)
			}
			f.RandomID = append(f.RandomID, value)
		}
	}
	{
		value, err := DecodeInputPeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.forwardMessages#c661bbc4: field to_peer: %w", err)
		}
		f.ToPeer = value
	}
	if f.Flags.Has(9) {
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.forwardMessages#c661bbc4: field top_msg_id: %w", err)
		}
		f.TopMsgID = value
	}
	if f.Flags.Has(10) {
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.forwardMessages#c661bbc4: field schedule_date: %w", err)
		}
		f.ScheduleDate = value
	}
	if f.Flags.Has(13) {
		value, err := DecodeInputPeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.forwardMessages#c661bbc4: field send_as: %w", err)
		}
		f.SendAs = value
	}
	return nil
}

// SetSilent sets value of Silent conditional field.
func (f *MessagesForwardMessagesRequest) SetSilent(value bool) {
	if value {
		f.Flags.Set(5)
		f.Silent = true
	} else {
		f.Flags.Unset(5)
		f.Silent = false
	}
}

// GetSilent returns value of Silent conditional field.
func (f *MessagesForwardMessagesRequest) GetSilent() (value bool) {
	if f == nil {
		return
	}
	return f.Flags.Has(5)
}

// SetBackground sets value of Background conditional field.
func (f *MessagesForwardMessagesRequest) SetBackground(value bool) {
	if value {
		f.Flags.Set(6)
		f.Background = true
	} else {
		f.Flags.Unset(6)
		f.Background = false
	}
}

// GetBackground returns value of Background conditional field.
func (f *MessagesForwardMessagesRequest) GetBackground() (value bool) {
	if f == nil {
		return
	}
	return f.Flags.Has(6)
}

// SetWithMyScore sets value of WithMyScore conditional field.
func (f *MessagesForwardMessagesRequest) SetWithMyScore(value bool) {
	if value {
		f.Flags.Set(8)
		f.WithMyScore = true
	} else {
		f.Flags.Unset(8)
		f.WithMyScore = false
	}
}

// GetWithMyScore returns value of WithMyScore conditional field.
func (f *MessagesForwardMessagesRequest) GetWithMyScore() (value bool) {
	if f == nil {
		return
	}
	return f.Flags.Has(8)
}

// SetDropAuthor sets value of DropAuthor conditional field.
func (f *MessagesForwardMessagesRequest) SetDropAuthor(value bool) {
	if value {
		f.Flags.Set(11)
		f.DropAuthor = true
	} else {
		f.Flags.Unset(11)
		f.DropAuthor = false
	}
}

// GetDropAuthor returns value of DropAuthor conditional field.
func (f *MessagesForwardMessagesRequest) GetDropAuthor() (value bool) {
	if f == nil {
		return
	}
	return f.Flags.Has(11)
}

// SetDropMediaCaptions sets value of DropMediaCaptions conditional field.
func (f *MessagesForwardMessagesRequest) SetDropMediaCaptions(value bool) {
	if value {
		f.Flags.Set(12)
		f.DropMediaCaptions = true
	} else {
		f.Flags.Unset(12)
		f.DropMediaCaptions = false
	}
}

// GetDropMediaCaptions returns value of DropMediaCaptions conditional field.
func (f *MessagesForwardMessagesRequest) GetDropMediaCaptions() (value bool) {
	if f == nil {
		return
	}
	return f.Flags.Has(12)
}

// SetNoforwards sets value of Noforwards conditional field.
func (f *MessagesForwardMessagesRequest) SetNoforwards(value bool) {
	if value {
		f.Flags.Set(14)
		f.Noforwards = true
	} else {
		f.Flags.Unset(14)
		f.Noforwards = false
	}
}

// GetNoforwards returns value of Noforwards conditional field.
func (f *MessagesForwardMessagesRequest) GetNoforwards() (value bool) {
	if f == nil {
		return
	}
	return f.Flags.Has(14)
}

// GetFromPeer returns value of FromPeer field.
func (f *MessagesForwardMessagesRequest) GetFromPeer() (value InputPeerClass) {
	if f == nil {
		return
	}
	return f.FromPeer
}

// GetID returns value of ID field.
func (f *MessagesForwardMessagesRequest) GetID() (value []int) {
	if f == nil {
		return
	}
	return f.ID
}

// GetRandomID returns value of RandomID field.
func (f *MessagesForwardMessagesRequest) GetRandomID() (value []int64) {
	if f == nil {
		return
	}
	return f.RandomID
}

// GetToPeer returns value of ToPeer field.
func (f *MessagesForwardMessagesRequest) GetToPeer() (value InputPeerClass) {
	if f == nil {
		return
	}
	return f.ToPeer
}

// SetTopMsgID sets value of TopMsgID conditional field.
func (f *MessagesForwardMessagesRequest) SetTopMsgID(value int) {
	f.Flags.Set(9)
	f.TopMsgID = value
}

// GetTopMsgID returns value of TopMsgID conditional field and
// boolean which is true if field was set.
func (f *MessagesForwardMessagesRequest) GetTopMsgID() (value int, ok bool) {
	if f == nil {
		return
	}
	if !f.Flags.Has(9) {
		return value, false
	}
	return f.TopMsgID, true
}

// SetScheduleDate sets value of ScheduleDate conditional field.
func (f *MessagesForwardMessagesRequest) SetScheduleDate(value int) {
	f.Flags.Set(10)
	f.ScheduleDate = value
}

// GetScheduleDate returns value of ScheduleDate conditional field and
// boolean which is true if field was set.
func (f *MessagesForwardMessagesRequest) GetScheduleDate() (value int, ok bool) {
	if f == nil {
		return
	}
	if !f.Flags.Has(10) {
		return value, false
	}
	return f.ScheduleDate, true
}

// SetSendAs sets value of SendAs conditional field.
func (f *MessagesForwardMessagesRequest) SetSendAs(value InputPeerClass) {
	f.Flags.Set(13)
	f.SendAs = value
}

// GetSendAs returns value of SendAs conditional field and
// boolean which is true if field was set.
func (f *MessagesForwardMessagesRequest) GetSendAs() (value InputPeerClass, ok bool) {
	if f == nil {
		return
	}
	if !f.Flags.Has(13) {
		return value, false
	}
	return f.SendAs, true
}

// MessagesForwardMessages invokes method messages.forwardMessages#c661bbc4 returning error if any.
// Forwards messages by their IDs.
//
// Possible errors:
//
//	400 BROADCAST_PUBLIC_VOTERS_FORBIDDEN: You can't forward polls with public voters.
//	400 CHANNEL_INVALID: The provided channel is invalid.
//	406 CHANNEL_PRIVATE: You haven't joined this channel/supergroup.
//	403 CHAT_ADMIN_REQUIRED: You must be an admin in this chat to do this.
//	406 CHAT_FORWARDS_RESTRICTED: You can't forward messages from a protected chat.
//	400 CHAT_ID_INVALID: The provided chat id is invalid.
//	400 CHAT_RESTRICTED: You can't send messages in this chat, you were restricted.
//	403 CHAT_SEND_AUDIOS_FORBIDDEN: You can't send audio messages in this chat.
//	403 CHAT_SEND_GAME_FORBIDDEN: You can't send a game to this chat.
//	403 CHAT_SEND_GIFS_FORBIDDEN: You can't send gifs in this chat.
//	403 CHAT_SEND_MEDIA_FORBIDDEN: You can't send media in this chat.
//	403 CHAT_SEND_PHOTOS_FORBIDDEN: You can't send photos in this chat.
//	403 CHAT_SEND_PLAIN_FORBIDDEN: You can't send non-media (text) messages in this chat.
//	403 CHAT_SEND_POLL_FORBIDDEN: You can't send polls in this chat.
//	403 CHAT_SEND_STICKERS_FORBIDDEN: You can't send stickers in this chat.
//	403 CHAT_SEND_VIDEOS_FORBIDDEN: You can't send videos in this chat.
//	403 CHAT_WRITE_FORBIDDEN: You can't write in this chat.
//	400 GROUPED_MEDIA_INVALID: Invalid grouped media.
//	400 INPUT_USER_DEACTIVATED: The specified user was deleted.
//	400 MEDIA_EMPTY: The provided media object is invalid.
//	400 MESSAGE_IDS_EMPTY: No message ids were provided.
//	400 MESSAGE_ID_INVALID: The provided message id is invalid.
//	400 MSG_ID_INVALID: Invalid message ID provided.
//	400 PEER_ID_INVALID: The provided peer id is invalid.
//	400 QUIZ_ANSWER_MISSING: You can forward a quiz while hiding the original author only after choosing an option in the quiz.
//	500 RANDOM_ID_DUPLICATE: You provided a random ID that was already used.
//	400 RANDOM_ID_INVALID: A provided random ID is invalid.
//	400 SCHEDULE_DATE_TOO_LATE: You can't schedule a message this far in the future.
//	400 SCHEDULE_TOO_MUCH: There are too many scheduled messages.
//	400 SEND_AS_PEER_INVALID: You can't send messages as the specified peer.
//	420 SLOWMODE_WAIT_%d: Slowmode is enabled in this chat: wait %d seconds before sending another message to this chat.
//	400 SLOWMODE_MULTI_MSGS_DISABLED: Slowmode is enabled, you cannot forward multiple messages to this group.
//	406 TOPIC_CLOSED: This topic was closed, you can't send messages to it anymore.
//	406 TOPIC_DELETED: The specified topic was deleted.
//	400 USER_BANNED_IN_CHANNEL: You're banned from sending messages in supergroups/channels.
//	403 USER_IS_BLOCKED: You were blocked by this user.
//	400 USER_IS_BOT: Bots can't send messages to other bots.
//	400 YOU_BLOCKED_USER: You blocked this user.
//
// See https://core.telegram.org/method/messages.forwardMessages for reference.
// Can be used by bots.
func (c *Client) MessagesForwardMessages(ctx context.Context, request *MessagesForwardMessagesRequest) (UpdatesClass, error) {
	var result UpdatesBox

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Updates, nil
}
