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

// GetGroupCallInviteLinkRequest represents TL type `getGroupCallInviteLink#2ae14924`.
type GetGroupCallInviteLinkRequest struct {
	// Group call identifier
	GroupCallID int32
	// Pass true if the invite_link should contain an invite hash, passing which to
	// joinGroupCall would allow the invited user to unmute themselves. Requires groupCall
	// can_be_managed group call flag
	CanSelfUnmute bool
}

// GetGroupCallInviteLinkRequestTypeID is TL type id of GetGroupCallInviteLinkRequest.
const GetGroupCallInviteLinkRequestTypeID = 0x2ae14924

// Ensuring interfaces in compile-time for GetGroupCallInviteLinkRequest.
var (
	_ bin.Encoder     = &GetGroupCallInviteLinkRequest{}
	_ bin.Decoder     = &GetGroupCallInviteLinkRequest{}
	_ bin.BareEncoder = &GetGroupCallInviteLinkRequest{}
	_ bin.BareDecoder = &GetGroupCallInviteLinkRequest{}
)

func (g *GetGroupCallInviteLinkRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.GroupCallID == 0) {
		return false
	}
	if !(g.CanSelfUnmute == false) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *GetGroupCallInviteLinkRequest) String() string {
	if g == nil {
		return "GetGroupCallInviteLinkRequest(nil)"
	}
	type Alias GetGroupCallInviteLinkRequest
	return fmt.Sprintf("GetGroupCallInviteLinkRequest%+v", Alias(*g))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*GetGroupCallInviteLinkRequest) TypeID() uint32 {
	return GetGroupCallInviteLinkRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*GetGroupCallInviteLinkRequest) TypeName() string {
	return "getGroupCallInviteLink"
}

// TypeInfo returns info about TL type.
func (g *GetGroupCallInviteLinkRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "getGroupCallInviteLink",
		ID:   GetGroupCallInviteLinkRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "GroupCallID",
			SchemaName: "group_call_id",
		},
		{
			Name:       "CanSelfUnmute",
			SchemaName: "can_self_unmute",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (g *GetGroupCallInviteLinkRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getGroupCallInviteLink#2ae14924 as nil")
	}
	b.PutID(GetGroupCallInviteLinkRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *GetGroupCallInviteLinkRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getGroupCallInviteLink#2ae14924 as nil")
	}
	b.PutInt32(g.GroupCallID)
	b.PutBool(g.CanSelfUnmute)
	return nil
}

// Decode implements bin.Decoder.
func (g *GetGroupCallInviteLinkRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getGroupCallInviteLink#2ae14924 to nil")
	}
	if err := b.ConsumeID(GetGroupCallInviteLinkRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode getGroupCallInviteLink#2ae14924: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *GetGroupCallInviteLinkRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getGroupCallInviteLink#2ae14924 to nil")
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode getGroupCallInviteLink#2ae14924: field group_call_id: %w", err)
		}
		g.GroupCallID = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode getGroupCallInviteLink#2ae14924: field can_self_unmute: %w", err)
		}
		g.CanSelfUnmute = value
	}
	return nil
}

// EncodeTDLibJSON implements jsontd.TDLibEncoder.
func (g *GetGroupCallInviteLinkRequest) EncodeTDLibJSON(b jsontd.Encoder) error {
	if g == nil {
		return fmt.Errorf("can't encode getGroupCallInviteLink#2ae14924 as nil")
	}
	b.ObjStart()
	b.PutID("getGroupCallInviteLink")
	b.FieldStart("group_call_id")
	b.PutInt32(g.GroupCallID)
	b.FieldStart("can_self_unmute")
	b.PutBool(g.CanSelfUnmute)
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements jsontd.TDLibDecoder.
func (g *GetGroupCallInviteLinkRequest) DecodeTDLibJSON(b jsontd.Decoder) error {
	if g == nil {
		return fmt.Errorf("can't decode getGroupCallInviteLink#2ae14924 to nil")
	}

	return b.Obj(func(b jsontd.Decoder, key []byte) error {
		switch string(key) {
		case jsontd.TypeField:
			if err := b.ConsumeID("getGroupCallInviteLink"); err != nil {
				return fmt.Errorf("unable to decode getGroupCallInviteLink#2ae14924: %w", err)
			}
		case "group_call_id":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode getGroupCallInviteLink#2ae14924: field group_call_id: %w", err)
			}
			g.GroupCallID = value
		case "can_self_unmute":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode getGroupCallInviteLink#2ae14924: field can_self_unmute: %w", err)
			}
			g.CanSelfUnmute = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetGroupCallID returns value of GroupCallID field.
func (g *GetGroupCallInviteLinkRequest) GetGroupCallID() (value int32) {
	return g.GroupCallID
}

// GetCanSelfUnmute returns value of CanSelfUnmute field.
func (g *GetGroupCallInviteLinkRequest) GetCanSelfUnmute() (value bool) {
	return g.CanSelfUnmute
}

// GetGroupCallInviteLink invokes method getGroupCallInviteLink#2ae14924 returning error if any.
func (c *Client) GetGroupCallInviteLink(ctx context.Context, request *GetGroupCallInviteLinkRequest) (*HTTPURL, error) {
	var result HTTPURL

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
