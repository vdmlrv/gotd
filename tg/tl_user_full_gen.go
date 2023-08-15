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

// UserFull represents TL type `userFull#4fe1cc86`.
// Extended user info
//
// See https://core.telegram.org/constructor/userFull for reference.
type UserFull struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// Whether you have blocked this user
	Blocked bool
	// Whether this user can make VoIP calls
	PhoneCallsAvailable bool
	// Whether this user's privacy settings allow you to call them
	PhoneCallsPrivate bool
	// Whether you can pin messages in the chat with this user, you can do this only for a
	// chat with yourself
	CanPinMessage bool
	// Whether scheduled messages¹ are available
	//
	// Links:
	//  1) https://core.telegram.org/api/scheduled-messages
	HasScheduled bool
	// Whether the user can receive video calls
	VideoCallsAvailable bool
	// Whether this user doesn't allow sending voice messages in a private chat with them
	VoiceMessagesForbidden bool
	// Whether the real-time chat translation popup¹ should be hidden.
	//
	// Links:
	//  1) https://core.telegram.org/api/translation
	TranslationsDisabled bool
	// StoriesPinnedAvailable field of UserFull.
	StoriesPinnedAvailable bool
	// BlockedMyStoriesFrom field of UserFull.
	BlockedMyStoriesFrom bool
	// User ID
	ID int64
	// Bio of the user
	//
	// Use SetAbout and GetAbout helpers.
	About string
	// Peer settings
	Settings PeerSettings
	// Personal profile photo, to be shown instead of profile_photo.
	//
	// Use SetPersonalPhoto and GetPersonalPhoto helpers.
	PersonalPhoto PhotoClass
	// Profile photo
	//
	// Use SetProfilePhoto and GetProfilePhoto helpers.
	ProfilePhoto PhotoClass
	// Fallback profile photo, displayed if no photo is present in profile_photo or
	// personal_photo, due to privacy settings.
	//
	// Use SetFallbackPhoto and GetFallbackPhoto helpers.
	FallbackPhoto PhotoClass
	// Notification settings
	NotifySettings PeerNotifySettings
	// For bots, info about the bot (bot commands, etc)
	//
	// Use SetBotInfo and GetBotInfo helpers.
	BotInfo BotInfo
	// Message ID of the last pinned message¹
	//
	// Links:
	//  1) https://core.telegram.org/api/pin
	//
	// Use SetPinnedMsgID and GetPinnedMsgID helpers.
	PinnedMsgID int
	// Chats in common with this user
	CommonChatsCount int
	// Peer folder ID, for more info click here¹
	//
	// Links:
	//  1) https://core.telegram.org/api/folders#peer-folders
	//
	// Use SetFolderID and GetFolderID helpers.
	FolderID int
	// Time To Live of all messages in this chat; once a message is this many seconds old, it
	// must be deleted.
	//
	// Use SetTTLPeriod and GetTTLPeriod helpers.
	TTLPeriod int
	// Emoji associated with chat theme
	//
	// Use SetThemeEmoticon and GetThemeEmoticon helpers.
	ThemeEmoticon string
	// Anonymized text to be shown instead of the user's name on forwarded messages
	//
	// Use SetPrivateForwardName and GetPrivateForwardName helpers.
	PrivateForwardName string
	// A suggested set of administrator rights¹ for the bot, to be shown when adding the bot
	// as admin to a group, see here for more info on how to handle them »².
	//
	// Links:
	//  1) https://core.telegram.org/api/rights#suggested-bot-rights
	//  2) https://core.telegram.org/api/rights#suggested-bot-rights
	//
	// Use SetBotGroupAdminRights and GetBotGroupAdminRights helpers.
	BotGroupAdminRights ChatAdminRights
	// A suggested set of administrator rights¹ for the bot, to be shown when adding the bot
	// as admin to a channel, see here for more info on how to handle them »².
	//
	// Links:
	//  1) https://core.telegram.org/api/rights#suggested-bot-rights
	//  2) https://core.telegram.org/api/rights#suggested-bot-rights
	//
	// Use SetBotBroadcastAdminRights and GetBotBroadcastAdminRights helpers.
	BotBroadcastAdminRights ChatAdminRights
	// Telegram Premium subscriptions gift options
	//
	// Use SetPremiumGifts and GetPremiumGifts helpers.
	PremiumGifts []PremiumGiftOption
	// Wallpaper¹ to use in the private chat with the user.
	//
	// Links:
	//  1) https://core.telegram.org/api/wallpapers
	//
	// Use SetWallpaper and GetWallpaper helpers.
	Wallpaper WallPaperClass
	// Stories field of UserFull.
	//
	// Use SetStories and GetStories helpers.
	Stories UserStories
}

// UserFullTypeID is TL type id of UserFull.
const UserFullTypeID = 0x4fe1cc86

// Ensuring interfaces in compile-time for UserFull.
var (
	_ bin.Encoder     = &UserFull{}
	_ bin.Decoder     = &UserFull{}
	_ bin.BareEncoder = &UserFull{}
	_ bin.BareDecoder = &UserFull{}
)

func (u *UserFull) Zero() bool {
	if u == nil {
		return true
	}
	if !(u.Flags.Zero()) {
		return false
	}
	if !(u.Blocked == false) {
		return false
	}
	if !(u.PhoneCallsAvailable == false) {
		return false
	}
	if !(u.PhoneCallsPrivate == false) {
		return false
	}
	if !(u.CanPinMessage == false) {
		return false
	}
	if !(u.HasScheduled == false) {
		return false
	}
	if !(u.VideoCallsAvailable == false) {
		return false
	}
	if !(u.VoiceMessagesForbidden == false) {
		return false
	}
	if !(u.TranslationsDisabled == false) {
		return false
	}
	if !(u.StoriesPinnedAvailable == false) {
		return false
	}
	if !(u.BlockedMyStoriesFrom == false) {
		return false
	}
	if !(u.ID == 0) {
		return false
	}
	if !(u.About == "") {
		return false
	}
	if !(u.Settings.Zero()) {
		return false
	}
	if !(u.PersonalPhoto == nil) {
		return false
	}
	if !(u.ProfilePhoto == nil) {
		return false
	}
	if !(u.FallbackPhoto == nil) {
		return false
	}
	if !(u.NotifySettings.Zero()) {
		return false
	}
	if !(u.BotInfo.Zero()) {
		return false
	}
	if !(u.PinnedMsgID == 0) {
		return false
	}
	if !(u.CommonChatsCount == 0) {
		return false
	}
	if !(u.FolderID == 0) {
		return false
	}
	if !(u.TTLPeriod == 0) {
		return false
	}
	if !(u.ThemeEmoticon == "") {
		return false
	}
	if !(u.PrivateForwardName == "") {
		return false
	}
	if !(u.BotGroupAdminRights.Zero()) {
		return false
	}
	if !(u.BotBroadcastAdminRights.Zero()) {
		return false
	}
	if !(u.PremiumGifts == nil) {
		return false
	}
	if !(u.Wallpaper == nil) {
		return false
	}
	if !(u.Stories.Zero()) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (u *UserFull) String() string {
	if u == nil {
		return "UserFull(nil)"
	}
	type Alias UserFull
	return fmt.Sprintf("UserFull%+v", Alias(*u))
}

// FillFrom fills UserFull from given interface.
func (u *UserFull) FillFrom(from interface {
	GetBlocked() (value bool)
	GetPhoneCallsAvailable() (value bool)
	GetPhoneCallsPrivate() (value bool)
	GetCanPinMessage() (value bool)
	GetHasScheduled() (value bool)
	GetVideoCallsAvailable() (value bool)
	GetVoiceMessagesForbidden() (value bool)
	GetTranslationsDisabled() (value bool)
	GetStoriesPinnedAvailable() (value bool)
	GetBlockedMyStoriesFrom() (value bool)
	GetID() (value int64)
	GetAbout() (value string, ok bool)
	GetSettings() (value PeerSettings)
	GetPersonalPhoto() (value PhotoClass, ok bool)
	GetProfilePhoto() (value PhotoClass, ok bool)
	GetFallbackPhoto() (value PhotoClass, ok bool)
	GetNotifySettings() (value PeerNotifySettings)
	GetBotInfo() (value BotInfo, ok bool)
	GetPinnedMsgID() (value int, ok bool)
	GetCommonChatsCount() (value int)
	GetFolderID() (value int, ok bool)
	GetTTLPeriod() (value int, ok bool)
	GetThemeEmoticon() (value string, ok bool)
	GetPrivateForwardName() (value string, ok bool)
	GetBotGroupAdminRights() (value ChatAdminRights, ok bool)
	GetBotBroadcastAdminRights() (value ChatAdminRights, ok bool)
	GetPremiumGifts() (value []PremiumGiftOption, ok bool)
	GetWallpaper() (value WallPaperClass, ok bool)
	GetStories() (value UserStories, ok bool)
}) {
	u.Blocked = from.GetBlocked()
	u.PhoneCallsAvailable = from.GetPhoneCallsAvailable()
	u.PhoneCallsPrivate = from.GetPhoneCallsPrivate()
	u.CanPinMessage = from.GetCanPinMessage()
	u.HasScheduled = from.GetHasScheduled()
	u.VideoCallsAvailable = from.GetVideoCallsAvailable()
	u.VoiceMessagesForbidden = from.GetVoiceMessagesForbidden()
	u.TranslationsDisabled = from.GetTranslationsDisabled()
	u.StoriesPinnedAvailable = from.GetStoriesPinnedAvailable()
	u.BlockedMyStoriesFrom = from.GetBlockedMyStoriesFrom()
	u.ID = from.GetID()
	if val, ok := from.GetAbout(); ok {
		u.About = val
	}

	u.Settings = from.GetSettings()
	if val, ok := from.GetPersonalPhoto(); ok {
		u.PersonalPhoto = val
	}

	if val, ok := from.GetProfilePhoto(); ok {
		u.ProfilePhoto = val
	}

	if val, ok := from.GetFallbackPhoto(); ok {
		u.FallbackPhoto = val
	}

	u.NotifySettings = from.GetNotifySettings()
	if val, ok := from.GetBotInfo(); ok {
		u.BotInfo = val
	}

	if val, ok := from.GetPinnedMsgID(); ok {
		u.PinnedMsgID = val
	}

	u.CommonChatsCount = from.GetCommonChatsCount()
	if val, ok := from.GetFolderID(); ok {
		u.FolderID = val
	}

	if val, ok := from.GetTTLPeriod(); ok {
		u.TTLPeriod = val
	}

	if val, ok := from.GetThemeEmoticon(); ok {
		u.ThemeEmoticon = val
	}

	if val, ok := from.GetPrivateForwardName(); ok {
		u.PrivateForwardName = val
	}

	if val, ok := from.GetBotGroupAdminRights(); ok {
		u.BotGroupAdminRights = val
	}

	if val, ok := from.GetBotBroadcastAdminRights(); ok {
		u.BotBroadcastAdminRights = val
	}

	if val, ok := from.GetPremiumGifts(); ok {
		u.PremiumGifts = val
	}

	if val, ok := from.GetWallpaper(); ok {
		u.Wallpaper = val
	}

	if val, ok := from.GetStories(); ok {
		u.Stories = val
	}

}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*UserFull) TypeID() uint32 {
	return UserFullTypeID
}

// TypeName returns name of type in TL schema.
func (*UserFull) TypeName() string {
	return "userFull"
}

// TypeInfo returns info about TL type.
func (u *UserFull) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "userFull",
		ID:   UserFullTypeID,
	}
	if u == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Blocked",
			SchemaName: "blocked",
			Null:       !u.Flags.Has(0),
		},
		{
			Name:       "PhoneCallsAvailable",
			SchemaName: "phone_calls_available",
			Null:       !u.Flags.Has(4),
		},
		{
			Name:       "PhoneCallsPrivate",
			SchemaName: "phone_calls_private",
			Null:       !u.Flags.Has(5),
		},
		{
			Name:       "CanPinMessage",
			SchemaName: "can_pin_message",
			Null:       !u.Flags.Has(7),
		},
		{
			Name:       "HasScheduled",
			SchemaName: "has_scheduled",
			Null:       !u.Flags.Has(12),
		},
		{
			Name:       "VideoCallsAvailable",
			SchemaName: "video_calls_available",
			Null:       !u.Flags.Has(13),
		},
		{
			Name:       "VoiceMessagesForbidden",
			SchemaName: "voice_messages_forbidden",
			Null:       !u.Flags.Has(20),
		},
		{
			Name:       "TranslationsDisabled",
			SchemaName: "translations_disabled",
			Null:       !u.Flags.Has(23),
		},
		{
			Name:       "StoriesPinnedAvailable",
			SchemaName: "stories_pinned_available",
			Null:       !u.Flags.Has(26),
		},
		{
			Name:       "BlockedMyStoriesFrom",
			SchemaName: "blocked_my_stories_from",
			Null:       !u.Flags.Has(27),
		},
		{
			Name:       "ID",
			SchemaName: "id",
		},
		{
			Name:       "About",
			SchemaName: "about",
			Null:       !u.Flags.Has(1),
		},
		{
			Name:       "Settings",
			SchemaName: "settings",
		},
		{
			Name:       "PersonalPhoto",
			SchemaName: "personal_photo",
			Null:       !u.Flags.Has(21),
		},
		{
			Name:       "ProfilePhoto",
			SchemaName: "profile_photo",
			Null:       !u.Flags.Has(2),
		},
		{
			Name:       "FallbackPhoto",
			SchemaName: "fallback_photo",
			Null:       !u.Flags.Has(22),
		},
		{
			Name:       "NotifySettings",
			SchemaName: "notify_settings",
		},
		{
			Name:       "BotInfo",
			SchemaName: "bot_info",
			Null:       !u.Flags.Has(3),
		},
		{
			Name:       "PinnedMsgID",
			SchemaName: "pinned_msg_id",
			Null:       !u.Flags.Has(6),
		},
		{
			Name:       "CommonChatsCount",
			SchemaName: "common_chats_count",
		},
		{
			Name:       "FolderID",
			SchemaName: "folder_id",
			Null:       !u.Flags.Has(11),
		},
		{
			Name:       "TTLPeriod",
			SchemaName: "ttl_period",
			Null:       !u.Flags.Has(14),
		},
		{
			Name:       "ThemeEmoticon",
			SchemaName: "theme_emoticon",
			Null:       !u.Flags.Has(15),
		},
		{
			Name:       "PrivateForwardName",
			SchemaName: "private_forward_name",
			Null:       !u.Flags.Has(16),
		},
		{
			Name:       "BotGroupAdminRights",
			SchemaName: "bot_group_admin_rights",
			Null:       !u.Flags.Has(17),
		},
		{
			Name:       "BotBroadcastAdminRights",
			SchemaName: "bot_broadcast_admin_rights",
			Null:       !u.Flags.Has(18),
		},
		{
			Name:       "PremiumGifts",
			SchemaName: "premium_gifts",
			Null:       !u.Flags.Has(19),
		},
		{
			Name:       "Wallpaper",
			SchemaName: "wallpaper",
			Null:       !u.Flags.Has(24),
		},
		{
			Name:       "Stories",
			SchemaName: "stories",
			Null:       !u.Flags.Has(25),
		},
	}
	return typ
}

// SetFlags sets flags for non-zero fields.
func (u *UserFull) SetFlags() {
	if !(u.Blocked == false) {
		u.Flags.Set(0)
	}
	if !(u.PhoneCallsAvailable == false) {
		u.Flags.Set(4)
	}
	if !(u.PhoneCallsPrivate == false) {
		u.Flags.Set(5)
	}
	if !(u.CanPinMessage == false) {
		u.Flags.Set(7)
	}
	if !(u.HasScheduled == false) {
		u.Flags.Set(12)
	}
	if !(u.VideoCallsAvailable == false) {
		u.Flags.Set(13)
	}
	if !(u.VoiceMessagesForbidden == false) {
		u.Flags.Set(20)
	}
	if !(u.TranslationsDisabled == false) {
		u.Flags.Set(23)
	}
	if !(u.StoriesPinnedAvailable == false) {
		u.Flags.Set(26)
	}
	if !(u.BlockedMyStoriesFrom == false) {
		u.Flags.Set(27)
	}
	if !(u.About == "") {
		u.Flags.Set(1)
	}
	if !(u.PersonalPhoto == nil) {
		u.Flags.Set(21)
	}
	if !(u.ProfilePhoto == nil) {
		u.Flags.Set(2)
	}
	if !(u.FallbackPhoto == nil) {
		u.Flags.Set(22)
	}
	if !(u.BotInfo.Zero()) {
		u.Flags.Set(3)
	}
	if !(u.PinnedMsgID == 0) {
		u.Flags.Set(6)
	}
	if !(u.FolderID == 0) {
		u.Flags.Set(11)
	}
	if !(u.TTLPeriod == 0) {
		u.Flags.Set(14)
	}
	if !(u.ThemeEmoticon == "") {
		u.Flags.Set(15)
	}
	if !(u.PrivateForwardName == "") {
		u.Flags.Set(16)
	}
	if !(u.BotGroupAdminRights.Zero()) {
		u.Flags.Set(17)
	}
	if !(u.BotBroadcastAdminRights.Zero()) {
		u.Flags.Set(18)
	}
	if !(u.PremiumGifts == nil) {
		u.Flags.Set(19)
	}
	if !(u.Wallpaper == nil) {
		u.Flags.Set(24)
	}
	if !(u.Stories.Zero()) {
		u.Flags.Set(25)
	}
}

// Encode implements bin.Encoder.
func (u *UserFull) Encode(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't encode userFull#4fe1cc86 as nil")
	}
	b.PutID(UserFullTypeID)
	return u.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (u *UserFull) EncodeBare(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't encode userFull#4fe1cc86 as nil")
	}
	u.SetFlags()
	if err := u.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode userFull#4fe1cc86: field flags: %w", err)
	}
	b.PutLong(u.ID)
	if u.Flags.Has(1) {
		b.PutString(u.About)
	}
	if err := u.Settings.Encode(b); err != nil {
		return fmt.Errorf("unable to encode userFull#4fe1cc86: field settings: %w", err)
	}
	if u.Flags.Has(21) {
		if u.PersonalPhoto == nil {
			return fmt.Errorf("unable to encode userFull#4fe1cc86: field personal_photo is nil")
		}
		if err := u.PersonalPhoto.Encode(b); err != nil {
			return fmt.Errorf("unable to encode userFull#4fe1cc86: field personal_photo: %w", err)
		}
	}
	if u.Flags.Has(2) {
		if u.ProfilePhoto == nil {
			return fmt.Errorf("unable to encode userFull#4fe1cc86: field profile_photo is nil")
		}
		if err := u.ProfilePhoto.Encode(b); err != nil {
			return fmt.Errorf("unable to encode userFull#4fe1cc86: field profile_photo: %w", err)
		}
	}
	if u.Flags.Has(22) {
		if u.FallbackPhoto == nil {
			return fmt.Errorf("unable to encode userFull#4fe1cc86: field fallback_photo is nil")
		}
		if err := u.FallbackPhoto.Encode(b); err != nil {
			return fmt.Errorf("unable to encode userFull#4fe1cc86: field fallback_photo: %w", err)
		}
	}
	if err := u.NotifySettings.Encode(b); err != nil {
		return fmt.Errorf("unable to encode userFull#4fe1cc86: field notify_settings: %w", err)
	}
	if u.Flags.Has(3) {
		if err := u.BotInfo.Encode(b); err != nil {
			return fmt.Errorf("unable to encode userFull#4fe1cc86: field bot_info: %w", err)
		}
	}
	if u.Flags.Has(6) {
		b.PutInt(u.PinnedMsgID)
	}
	b.PutInt(u.CommonChatsCount)
	if u.Flags.Has(11) {
		b.PutInt(u.FolderID)
	}
	if u.Flags.Has(14) {
		b.PutInt(u.TTLPeriod)
	}
	if u.Flags.Has(15) {
		b.PutString(u.ThemeEmoticon)
	}
	if u.Flags.Has(16) {
		b.PutString(u.PrivateForwardName)
	}
	if u.Flags.Has(17) {
		if err := u.BotGroupAdminRights.Encode(b); err != nil {
			return fmt.Errorf("unable to encode userFull#4fe1cc86: field bot_group_admin_rights: %w", err)
		}
	}
	if u.Flags.Has(18) {
		if err := u.BotBroadcastAdminRights.Encode(b); err != nil {
			return fmt.Errorf("unable to encode userFull#4fe1cc86: field bot_broadcast_admin_rights: %w", err)
		}
	}
	if u.Flags.Has(19) {
		b.PutVectorHeader(len(u.PremiumGifts))
		for idx, v := range u.PremiumGifts {
			if err := v.Encode(b); err != nil {
				return fmt.Errorf("unable to encode userFull#4fe1cc86: field premium_gifts element with index %d: %w", idx, err)
			}
		}
	}
	if u.Flags.Has(24) {
		if u.Wallpaper == nil {
			return fmt.Errorf("unable to encode userFull#4fe1cc86: field wallpaper is nil")
		}
		if err := u.Wallpaper.Encode(b); err != nil {
			return fmt.Errorf("unable to encode userFull#4fe1cc86: field wallpaper: %w", err)
		}
	}
	if u.Flags.Has(25) {
		if err := u.Stories.Encode(b); err != nil {
			return fmt.Errorf("unable to encode userFull#4fe1cc86: field stories: %w", err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (u *UserFull) Decode(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't decode userFull#4fe1cc86 to nil")
	}
	if err := b.ConsumeID(UserFullTypeID); err != nil {
		return fmt.Errorf("unable to decode userFull#4fe1cc86: %w", err)
	}
	return u.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (u *UserFull) DecodeBare(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't decode userFull#4fe1cc86 to nil")
	}
	{
		if err := u.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode userFull#4fe1cc86: field flags: %w", err)
		}
	}
	u.Blocked = u.Flags.Has(0)
	u.PhoneCallsAvailable = u.Flags.Has(4)
	u.PhoneCallsPrivate = u.Flags.Has(5)
	u.CanPinMessage = u.Flags.Has(7)
	u.HasScheduled = u.Flags.Has(12)
	u.VideoCallsAvailable = u.Flags.Has(13)
	u.VoiceMessagesForbidden = u.Flags.Has(20)
	u.TranslationsDisabled = u.Flags.Has(23)
	u.StoriesPinnedAvailable = u.Flags.Has(26)
	u.BlockedMyStoriesFrom = u.Flags.Has(27)
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode userFull#4fe1cc86: field id: %w", err)
		}
		u.ID = value
	}
	if u.Flags.Has(1) {
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode userFull#4fe1cc86: field about: %w", err)
		}
		u.About = value
	}
	{
		if err := u.Settings.Decode(b); err != nil {
			return fmt.Errorf("unable to decode userFull#4fe1cc86: field settings: %w", err)
		}
	}
	if u.Flags.Has(21) {
		value, err := DecodePhoto(b)
		if err != nil {
			return fmt.Errorf("unable to decode userFull#4fe1cc86: field personal_photo: %w", err)
		}
		u.PersonalPhoto = value
	}
	if u.Flags.Has(2) {
		value, err := DecodePhoto(b)
		if err != nil {
			return fmt.Errorf("unable to decode userFull#4fe1cc86: field profile_photo: %w", err)
		}
		u.ProfilePhoto = value
	}
	if u.Flags.Has(22) {
		value, err := DecodePhoto(b)
		if err != nil {
			return fmt.Errorf("unable to decode userFull#4fe1cc86: field fallback_photo: %w", err)
		}
		u.FallbackPhoto = value
	}
	{
		if err := u.NotifySettings.Decode(b); err != nil {
			return fmt.Errorf("unable to decode userFull#4fe1cc86: field notify_settings: %w", err)
		}
	}
	if u.Flags.Has(3) {
		if err := u.BotInfo.Decode(b); err != nil {
			return fmt.Errorf("unable to decode userFull#4fe1cc86: field bot_info: %w", err)
		}
	}
	if u.Flags.Has(6) {
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode userFull#4fe1cc86: field pinned_msg_id: %w", err)
		}
		u.PinnedMsgID = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode userFull#4fe1cc86: field common_chats_count: %w", err)
		}
		u.CommonChatsCount = value
	}
	if u.Flags.Has(11) {
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode userFull#4fe1cc86: field folder_id: %w", err)
		}
		u.FolderID = value
	}
	if u.Flags.Has(14) {
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode userFull#4fe1cc86: field ttl_period: %w", err)
		}
		u.TTLPeriod = value
	}
	if u.Flags.Has(15) {
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode userFull#4fe1cc86: field theme_emoticon: %w", err)
		}
		u.ThemeEmoticon = value
	}
	if u.Flags.Has(16) {
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode userFull#4fe1cc86: field private_forward_name: %w", err)
		}
		u.PrivateForwardName = value
	}
	if u.Flags.Has(17) {
		if err := u.BotGroupAdminRights.Decode(b); err != nil {
			return fmt.Errorf("unable to decode userFull#4fe1cc86: field bot_group_admin_rights: %w", err)
		}
	}
	if u.Flags.Has(18) {
		if err := u.BotBroadcastAdminRights.Decode(b); err != nil {
			return fmt.Errorf("unable to decode userFull#4fe1cc86: field bot_broadcast_admin_rights: %w", err)
		}
	}
	if u.Flags.Has(19) {
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode userFull#4fe1cc86: field premium_gifts: %w", err)
		}

		if headerLen > 0 {
			u.PremiumGifts = make([]PremiumGiftOption, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			var value PremiumGiftOption
			if err := value.Decode(b); err != nil {
				return fmt.Errorf("unable to decode userFull#4fe1cc86: field premium_gifts: %w", err)
			}
			u.PremiumGifts = append(u.PremiumGifts, value)
		}
	}
	if u.Flags.Has(24) {
		value, err := DecodeWallPaper(b)
		if err != nil {
			return fmt.Errorf("unable to decode userFull#4fe1cc86: field wallpaper: %w", err)
		}
		u.Wallpaper = value
	}
	if u.Flags.Has(25) {
		if err := u.Stories.Decode(b); err != nil {
			return fmt.Errorf("unable to decode userFull#4fe1cc86: field stories: %w", err)
		}
	}
	return nil
}

// SetBlocked sets value of Blocked conditional field.
func (u *UserFull) SetBlocked(value bool) {
	if value {
		u.Flags.Set(0)
		u.Blocked = true
	} else {
		u.Flags.Unset(0)
		u.Blocked = false
	}
}

// GetBlocked returns value of Blocked conditional field.
func (u *UserFull) GetBlocked() (value bool) {
	if u == nil {
		return
	}
	return u.Flags.Has(0)
}

// SetPhoneCallsAvailable sets value of PhoneCallsAvailable conditional field.
func (u *UserFull) SetPhoneCallsAvailable(value bool) {
	if value {
		u.Flags.Set(4)
		u.PhoneCallsAvailable = true
	} else {
		u.Flags.Unset(4)
		u.PhoneCallsAvailable = false
	}
}

// GetPhoneCallsAvailable returns value of PhoneCallsAvailable conditional field.
func (u *UserFull) GetPhoneCallsAvailable() (value bool) {
	if u == nil {
		return
	}
	return u.Flags.Has(4)
}

// SetPhoneCallsPrivate sets value of PhoneCallsPrivate conditional field.
func (u *UserFull) SetPhoneCallsPrivate(value bool) {
	if value {
		u.Flags.Set(5)
		u.PhoneCallsPrivate = true
	} else {
		u.Flags.Unset(5)
		u.PhoneCallsPrivate = false
	}
}

// GetPhoneCallsPrivate returns value of PhoneCallsPrivate conditional field.
func (u *UserFull) GetPhoneCallsPrivate() (value bool) {
	if u == nil {
		return
	}
	return u.Flags.Has(5)
}

// SetCanPinMessage sets value of CanPinMessage conditional field.
func (u *UserFull) SetCanPinMessage(value bool) {
	if value {
		u.Flags.Set(7)
		u.CanPinMessage = true
	} else {
		u.Flags.Unset(7)
		u.CanPinMessage = false
	}
}

// GetCanPinMessage returns value of CanPinMessage conditional field.
func (u *UserFull) GetCanPinMessage() (value bool) {
	if u == nil {
		return
	}
	return u.Flags.Has(7)
}

// SetHasScheduled sets value of HasScheduled conditional field.
func (u *UserFull) SetHasScheduled(value bool) {
	if value {
		u.Flags.Set(12)
		u.HasScheduled = true
	} else {
		u.Flags.Unset(12)
		u.HasScheduled = false
	}
}

// GetHasScheduled returns value of HasScheduled conditional field.
func (u *UserFull) GetHasScheduled() (value bool) {
	if u == nil {
		return
	}
	return u.Flags.Has(12)
}

// SetVideoCallsAvailable sets value of VideoCallsAvailable conditional field.
func (u *UserFull) SetVideoCallsAvailable(value bool) {
	if value {
		u.Flags.Set(13)
		u.VideoCallsAvailable = true
	} else {
		u.Flags.Unset(13)
		u.VideoCallsAvailable = false
	}
}

// GetVideoCallsAvailable returns value of VideoCallsAvailable conditional field.
func (u *UserFull) GetVideoCallsAvailable() (value bool) {
	if u == nil {
		return
	}
	return u.Flags.Has(13)
}

// SetVoiceMessagesForbidden sets value of VoiceMessagesForbidden conditional field.
func (u *UserFull) SetVoiceMessagesForbidden(value bool) {
	if value {
		u.Flags.Set(20)
		u.VoiceMessagesForbidden = true
	} else {
		u.Flags.Unset(20)
		u.VoiceMessagesForbidden = false
	}
}

// GetVoiceMessagesForbidden returns value of VoiceMessagesForbidden conditional field.
func (u *UserFull) GetVoiceMessagesForbidden() (value bool) {
	if u == nil {
		return
	}
	return u.Flags.Has(20)
}

// SetTranslationsDisabled sets value of TranslationsDisabled conditional field.
func (u *UserFull) SetTranslationsDisabled(value bool) {
	if value {
		u.Flags.Set(23)
		u.TranslationsDisabled = true
	} else {
		u.Flags.Unset(23)
		u.TranslationsDisabled = false
	}
}

// GetTranslationsDisabled returns value of TranslationsDisabled conditional field.
func (u *UserFull) GetTranslationsDisabled() (value bool) {
	if u == nil {
		return
	}
	return u.Flags.Has(23)
}

// SetStoriesPinnedAvailable sets value of StoriesPinnedAvailable conditional field.
func (u *UserFull) SetStoriesPinnedAvailable(value bool) {
	if value {
		u.Flags.Set(26)
		u.StoriesPinnedAvailable = true
	} else {
		u.Flags.Unset(26)
		u.StoriesPinnedAvailable = false
	}
}

// GetStoriesPinnedAvailable returns value of StoriesPinnedAvailable conditional field.
func (u *UserFull) GetStoriesPinnedAvailable() (value bool) {
	if u == nil {
		return
	}
	return u.Flags.Has(26)
}

// SetBlockedMyStoriesFrom sets value of BlockedMyStoriesFrom conditional field.
func (u *UserFull) SetBlockedMyStoriesFrom(value bool) {
	if value {
		u.Flags.Set(27)
		u.BlockedMyStoriesFrom = true
	} else {
		u.Flags.Unset(27)
		u.BlockedMyStoriesFrom = false
	}
}

// GetBlockedMyStoriesFrom returns value of BlockedMyStoriesFrom conditional field.
func (u *UserFull) GetBlockedMyStoriesFrom() (value bool) {
	if u == nil {
		return
	}
	return u.Flags.Has(27)
}

// GetID returns value of ID field.
func (u *UserFull) GetID() (value int64) {
	if u == nil {
		return
	}
	return u.ID
}

// SetAbout sets value of About conditional field.
func (u *UserFull) SetAbout(value string) {
	u.Flags.Set(1)
	u.About = value
}

// GetAbout returns value of About conditional field and
// boolean which is true if field was set.
func (u *UserFull) GetAbout() (value string, ok bool) {
	if u == nil {
		return
	}
	if !u.Flags.Has(1) {
		return value, false
	}
	return u.About, true
}

// GetSettings returns value of Settings field.
func (u *UserFull) GetSettings() (value PeerSettings) {
	if u == nil {
		return
	}
	return u.Settings
}

// SetPersonalPhoto sets value of PersonalPhoto conditional field.
func (u *UserFull) SetPersonalPhoto(value PhotoClass) {
	u.Flags.Set(21)
	u.PersonalPhoto = value
}

// GetPersonalPhoto returns value of PersonalPhoto conditional field and
// boolean which is true if field was set.
func (u *UserFull) GetPersonalPhoto() (value PhotoClass, ok bool) {
	if u == nil {
		return
	}
	if !u.Flags.Has(21) {
		return value, false
	}
	return u.PersonalPhoto, true
}

// SetProfilePhoto sets value of ProfilePhoto conditional field.
func (u *UserFull) SetProfilePhoto(value PhotoClass) {
	u.Flags.Set(2)
	u.ProfilePhoto = value
}

// GetProfilePhoto returns value of ProfilePhoto conditional field and
// boolean which is true if field was set.
func (u *UserFull) GetProfilePhoto() (value PhotoClass, ok bool) {
	if u == nil {
		return
	}
	if !u.Flags.Has(2) {
		return value, false
	}
	return u.ProfilePhoto, true
}

// SetFallbackPhoto sets value of FallbackPhoto conditional field.
func (u *UserFull) SetFallbackPhoto(value PhotoClass) {
	u.Flags.Set(22)
	u.FallbackPhoto = value
}

// GetFallbackPhoto returns value of FallbackPhoto conditional field and
// boolean which is true if field was set.
func (u *UserFull) GetFallbackPhoto() (value PhotoClass, ok bool) {
	if u == nil {
		return
	}
	if !u.Flags.Has(22) {
		return value, false
	}
	return u.FallbackPhoto, true
}

// GetNotifySettings returns value of NotifySettings field.
func (u *UserFull) GetNotifySettings() (value PeerNotifySettings) {
	if u == nil {
		return
	}
	return u.NotifySettings
}

// SetBotInfo sets value of BotInfo conditional field.
func (u *UserFull) SetBotInfo(value BotInfo) {
	u.Flags.Set(3)
	u.BotInfo = value
}

// GetBotInfo returns value of BotInfo conditional field and
// boolean which is true if field was set.
func (u *UserFull) GetBotInfo() (value BotInfo, ok bool) {
	if u == nil {
		return
	}
	if !u.Flags.Has(3) {
		return value, false
	}
	return u.BotInfo, true
}

// SetPinnedMsgID sets value of PinnedMsgID conditional field.
func (u *UserFull) SetPinnedMsgID(value int) {
	u.Flags.Set(6)
	u.PinnedMsgID = value
}

// GetPinnedMsgID returns value of PinnedMsgID conditional field and
// boolean which is true if field was set.
func (u *UserFull) GetPinnedMsgID() (value int, ok bool) {
	if u == nil {
		return
	}
	if !u.Flags.Has(6) {
		return value, false
	}
	return u.PinnedMsgID, true
}

// GetCommonChatsCount returns value of CommonChatsCount field.
func (u *UserFull) GetCommonChatsCount() (value int) {
	if u == nil {
		return
	}
	return u.CommonChatsCount
}

// SetFolderID sets value of FolderID conditional field.
func (u *UserFull) SetFolderID(value int) {
	u.Flags.Set(11)
	u.FolderID = value
}

// GetFolderID returns value of FolderID conditional field and
// boolean which is true if field was set.
func (u *UserFull) GetFolderID() (value int, ok bool) {
	if u == nil {
		return
	}
	if !u.Flags.Has(11) {
		return value, false
	}
	return u.FolderID, true
}

// SetTTLPeriod sets value of TTLPeriod conditional field.
func (u *UserFull) SetTTLPeriod(value int) {
	u.Flags.Set(14)
	u.TTLPeriod = value
}

// GetTTLPeriod returns value of TTLPeriod conditional field and
// boolean which is true if field was set.
func (u *UserFull) GetTTLPeriod() (value int, ok bool) {
	if u == nil {
		return
	}
	if !u.Flags.Has(14) {
		return value, false
	}
	return u.TTLPeriod, true
}

// SetThemeEmoticon sets value of ThemeEmoticon conditional field.
func (u *UserFull) SetThemeEmoticon(value string) {
	u.Flags.Set(15)
	u.ThemeEmoticon = value
}

// GetThemeEmoticon returns value of ThemeEmoticon conditional field and
// boolean which is true if field was set.
func (u *UserFull) GetThemeEmoticon() (value string, ok bool) {
	if u == nil {
		return
	}
	if !u.Flags.Has(15) {
		return value, false
	}
	return u.ThemeEmoticon, true
}

// SetPrivateForwardName sets value of PrivateForwardName conditional field.
func (u *UserFull) SetPrivateForwardName(value string) {
	u.Flags.Set(16)
	u.PrivateForwardName = value
}

// GetPrivateForwardName returns value of PrivateForwardName conditional field and
// boolean which is true if field was set.
func (u *UserFull) GetPrivateForwardName() (value string, ok bool) {
	if u == nil {
		return
	}
	if !u.Flags.Has(16) {
		return value, false
	}
	return u.PrivateForwardName, true
}

// SetBotGroupAdminRights sets value of BotGroupAdminRights conditional field.
func (u *UserFull) SetBotGroupAdminRights(value ChatAdminRights) {
	u.Flags.Set(17)
	u.BotGroupAdminRights = value
}

// GetBotGroupAdminRights returns value of BotGroupAdminRights conditional field and
// boolean which is true if field was set.
func (u *UserFull) GetBotGroupAdminRights() (value ChatAdminRights, ok bool) {
	if u == nil {
		return
	}
	if !u.Flags.Has(17) {
		return value, false
	}
	return u.BotGroupAdminRights, true
}

// SetBotBroadcastAdminRights sets value of BotBroadcastAdminRights conditional field.
func (u *UserFull) SetBotBroadcastAdminRights(value ChatAdminRights) {
	u.Flags.Set(18)
	u.BotBroadcastAdminRights = value
}

// GetBotBroadcastAdminRights returns value of BotBroadcastAdminRights conditional field and
// boolean which is true if field was set.
func (u *UserFull) GetBotBroadcastAdminRights() (value ChatAdminRights, ok bool) {
	if u == nil {
		return
	}
	if !u.Flags.Has(18) {
		return value, false
	}
	return u.BotBroadcastAdminRights, true
}

// SetPremiumGifts sets value of PremiumGifts conditional field.
func (u *UserFull) SetPremiumGifts(value []PremiumGiftOption) {
	u.Flags.Set(19)
	u.PremiumGifts = value
}

// GetPremiumGifts returns value of PremiumGifts conditional field and
// boolean which is true if field was set.
func (u *UserFull) GetPremiumGifts() (value []PremiumGiftOption, ok bool) {
	if u == nil {
		return
	}
	if !u.Flags.Has(19) {
		return value, false
	}
	return u.PremiumGifts, true
}

// SetWallpaper sets value of Wallpaper conditional field.
func (u *UserFull) SetWallpaper(value WallPaperClass) {
	u.Flags.Set(24)
	u.Wallpaper = value
}

// GetWallpaper returns value of Wallpaper conditional field and
// boolean which is true if field was set.
func (u *UserFull) GetWallpaper() (value WallPaperClass, ok bool) {
	if u == nil {
		return
	}
	if !u.Flags.Has(24) {
		return value, false
	}
	return u.Wallpaper, true
}

// SetStories sets value of Stories conditional field.
func (u *UserFull) SetStories(value UserStories) {
	u.Flags.Set(25)
	u.Stories = value
}

// GetStories returns value of Stories conditional field and
// boolean which is true if field was set.
func (u *UserFull) GetStories() (value UserStories, ok bool) {
	if u == nil {
		return
	}
	if !u.Flags.Has(25) {
		return value, false
	}
	return u.Stories, true
}

// GetPersonalPhotoAsNotEmpty returns mapped value of PersonalPhoto conditional field and
// boolean which is true if field was set.
func (u *UserFull) GetPersonalPhotoAsNotEmpty() (*Photo, bool) {
	if value, ok := u.GetPersonalPhoto(); ok {
		return value.AsNotEmpty()
	}
	return nil, false
}

// GetProfilePhotoAsNotEmpty returns mapped value of ProfilePhoto conditional field and
// boolean which is true if field was set.
func (u *UserFull) GetProfilePhotoAsNotEmpty() (*Photo, bool) {
	if value, ok := u.GetProfilePhoto(); ok {
		return value.AsNotEmpty()
	}
	return nil, false
}

// GetFallbackPhotoAsNotEmpty returns mapped value of FallbackPhoto conditional field and
// boolean which is true if field was set.
func (u *UserFull) GetFallbackPhotoAsNotEmpty() (*Photo, bool) {
	if value, ok := u.GetFallbackPhoto(); ok {
		return value.AsNotEmpty()
	}
	return nil, false
}
