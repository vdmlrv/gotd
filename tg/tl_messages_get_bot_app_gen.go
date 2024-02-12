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

// MessagesGetBotAppRequest represents TL type `messages.getBotApp#34fdc5c3`.
// Obtain information about a direct link Mini App¹
//
// Links:
//  1. https://core.telegram.org/api/bots/webapps#direct-link-mini-apps
//
// See https://core.telegram.org/method/messages.getBotApp for reference.
type MessagesGetBotAppRequest struct {
	// Bot app information obtained from a Direct Mini App deep link »¹.
	//
	// Links:
	//  1) https://core.telegram.org/api/links#direct-mini-app-links
	App InputBotAppClass
	// Hash for pagination, for more info click here¹
	//
	// Links:
	//  1) https://core.telegram.org/api/offsets#hash-generation
	Hash int64
}

// MessagesGetBotAppRequestTypeID is TL type id of MessagesGetBotAppRequest.
const MessagesGetBotAppRequestTypeID = 0x34fdc5c3

// Ensuring interfaces in compile-time for MessagesGetBotAppRequest.
var (
	_ bin.Encoder     = &MessagesGetBotAppRequest{}
	_ bin.Decoder     = &MessagesGetBotAppRequest{}
	_ bin.BareEncoder = &MessagesGetBotAppRequest{}
	_ bin.BareDecoder = &MessagesGetBotAppRequest{}
)

func (g *MessagesGetBotAppRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.App == nil) {
		return false
	}
	if !(g.Hash == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *MessagesGetBotAppRequest) String() string {
	if g == nil {
		return "MessagesGetBotAppRequest(nil)"
	}
	type Alias MessagesGetBotAppRequest
	return fmt.Sprintf("MessagesGetBotAppRequest%+v", Alias(*g))
}

// FillFrom fills MessagesGetBotAppRequest from given interface.
func (g *MessagesGetBotAppRequest) FillFrom(from interface {
	GetApp() (value InputBotAppClass)
	GetHash() (value int64)
}) {
	g.App = from.GetApp()
	g.Hash = from.GetHash()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessagesGetBotAppRequest) TypeID() uint32 {
	return MessagesGetBotAppRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*MessagesGetBotAppRequest) TypeName() string {
	return "messages.getBotApp"
}

// TypeInfo returns info about TL type.
func (g *MessagesGetBotAppRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messages.getBotApp",
		ID:   MessagesGetBotAppRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "App",
			SchemaName: "app",
		},
		{
			Name:       "Hash",
			SchemaName: "hash",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (g *MessagesGetBotAppRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode messages.getBotApp#34fdc5c3 as nil")
	}
	b.PutID(MessagesGetBotAppRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *MessagesGetBotAppRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode messages.getBotApp#34fdc5c3 as nil")
	}
	if g.App == nil {
		return fmt.Errorf("unable to encode messages.getBotApp#34fdc5c3: field app is nil")
	}
	if err := g.App.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.getBotApp#34fdc5c3: field app: %w", err)
	}
	b.PutLong(g.Hash)
	return nil
}

// Decode implements bin.Decoder.
func (g *MessagesGetBotAppRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode messages.getBotApp#34fdc5c3 to nil")
	}
	if err := b.ConsumeID(MessagesGetBotAppRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.getBotApp#34fdc5c3: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *MessagesGetBotAppRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode messages.getBotApp#34fdc5c3 to nil")
	}
	{
		value, err := DecodeInputBotApp(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.getBotApp#34fdc5c3: field app: %w", err)
		}
		g.App = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode messages.getBotApp#34fdc5c3: field hash: %w", err)
		}
		g.Hash = value
	}
	return nil
}

// GetApp returns value of App field.
func (g *MessagesGetBotAppRequest) GetApp() (value InputBotAppClass) {
	if g == nil {
		return
	}
	return g.App
}

// GetHash returns value of Hash field.
func (g *MessagesGetBotAppRequest) GetHash() (value int64) {
	if g == nil {
		return
	}
	return g.Hash
}

// MessagesGetBotApp invokes method messages.getBotApp#34fdc5c3 returning error if any.
// Obtain information about a direct link Mini App¹
//
// Links:
//  1. https://core.telegram.org/api/bots/webapps#direct-link-mini-apps
//
// Possible errors:
//
//	400 BOT_APP_INVALID: The specified bot app is invalid.
//
// See https://core.telegram.org/method/messages.getBotApp for reference.
func (c *Client) MessagesGetBotApp(ctx context.Context, request *MessagesGetBotAppRequest) (*MessagesBotApp, error) {
	var result MessagesBotApp

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
