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

// GetUserPrivacySettingRulesRequest represents TL type `getUserPrivacySettingRules#84301671`.
type GetUserPrivacySettingRulesRequest struct {
	// The privacy setting
	Setting UserPrivacySettingClass
}

// GetUserPrivacySettingRulesRequestTypeID is TL type id of GetUserPrivacySettingRulesRequest.
const GetUserPrivacySettingRulesRequestTypeID = 0x84301671

// Ensuring interfaces in compile-time for GetUserPrivacySettingRulesRequest.
var (
	_ bin.Encoder     = &GetUserPrivacySettingRulesRequest{}
	_ bin.Decoder     = &GetUserPrivacySettingRulesRequest{}
	_ bin.BareEncoder = &GetUserPrivacySettingRulesRequest{}
	_ bin.BareDecoder = &GetUserPrivacySettingRulesRequest{}
)

func (g *GetUserPrivacySettingRulesRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.Setting == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *GetUserPrivacySettingRulesRequest) String() string {
	if g == nil {
		return "GetUserPrivacySettingRulesRequest(nil)"
	}
	type Alias GetUserPrivacySettingRulesRequest
	return fmt.Sprintf("GetUserPrivacySettingRulesRequest%+v", Alias(*g))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*GetUserPrivacySettingRulesRequest) TypeID() uint32 {
	return GetUserPrivacySettingRulesRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*GetUserPrivacySettingRulesRequest) TypeName() string {
	return "getUserPrivacySettingRules"
}

// TypeInfo returns info about TL type.
func (g *GetUserPrivacySettingRulesRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "getUserPrivacySettingRules",
		ID:   GetUserPrivacySettingRulesRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Setting",
			SchemaName: "setting",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (g *GetUserPrivacySettingRulesRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getUserPrivacySettingRules#84301671 as nil")
	}
	b.PutID(GetUserPrivacySettingRulesRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *GetUserPrivacySettingRulesRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getUserPrivacySettingRules#84301671 as nil")
	}
	if g.Setting == nil {
		return fmt.Errorf("unable to encode getUserPrivacySettingRules#84301671: field setting is nil")
	}
	if err := g.Setting.Encode(b); err != nil {
		return fmt.Errorf("unable to encode getUserPrivacySettingRules#84301671: field setting: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (g *GetUserPrivacySettingRulesRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getUserPrivacySettingRules#84301671 to nil")
	}
	if err := b.ConsumeID(GetUserPrivacySettingRulesRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode getUserPrivacySettingRules#84301671: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *GetUserPrivacySettingRulesRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getUserPrivacySettingRules#84301671 to nil")
	}
	{
		value, err := DecodeUserPrivacySetting(b)
		if err != nil {
			return fmt.Errorf("unable to decode getUserPrivacySettingRules#84301671: field setting: %w", err)
		}
		g.Setting = value
	}
	return nil
}

// EncodeTDLibJSON implements jsontd.TDLibEncoder.
func (g *GetUserPrivacySettingRulesRequest) EncodeTDLibJSON(b jsontd.Encoder) error {
	if g == nil {
		return fmt.Errorf("can't encode getUserPrivacySettingRules#84301671 as nil")
	}
	b.ObjStart()
	b.PutID("getUserPrivacySettingRules")
	b.FieldStart("setting")
	if g.Setting == nil {
		return fmt.Errorf("unable to encode getUserPrivacySettingRules#84301671: field setting is nil")
	}
	if err := g.Setting.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode getUserPrivacySettingRules#84301671: field setting: %w", err)
	}
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements jsontd.TDLibDecoder.
func (g *GetUserPrivacySettingRulesRequest) DecodeTDLibJSON(b jsontd.Decoder) error {
	if g == nil {
		return fmt.Errorf("can't decode getUserPrivacySettingRules#84301671 to nil")
	}

	return b.Obj(func(b jsontd.Decoder, key []byte) error {
		switch string(key) {
		case jsontd.TypeField:
			if err := b.ConsumeID("getUserPrivacySettingRules"); err != nil {
				return fmt.Errorf("unable to decode getUserPrivacySettingRules#84301671: %w", err)
			}
		case "setting":
			value, err := DecodeTDLibJSONUserPrivacySetting(b)
			if err != nil {
				return fmt.Errorf("unable to decode getUserPrivacySettingRules#84301671: field setting: %w", err)
			}
			g.Setting = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetSetting returns value of Setting field.
func (g *GetUserPrivacySettingRulesRequest) GetSetting() (value UserPrivacySettingClass) {
	return g.Setting
}

// GetUserPrivacySettingRules invokes method getUserPrivacySettingRules#84301671 returning error if any.
func (c *Client) GetUserPrivacySettingRules(ctx context.Context, setting UserPrivacySettingClass) (*UserPrivacySettingRules, error) {
	var result UserPrivacySettingRules

	request := &GetUserPrivacySettingRulesRequest{
		Setting: setting,
	}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
