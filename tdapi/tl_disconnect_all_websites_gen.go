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

// DisconnectAllWebsitesRequest represents TL type `disconnectAllWebsites#bf72f203`.
type DisconnectAllWebsitesRequest struct {
}

// DisconnectAllWebsitesRequestTypeID is TL type id of DisconnectAllWebsitesRequest.
const DisconnectAllWebsitesRequestTypeID = 0xbf72f203

// Ensuring interfaces in compile-time for DisconnectAllWebsitesRequest.
var (
	_ bin.Encoder     = &DisconnectAllWebsitesRequest{}
	_ bin.Decoder     = &DisconnectAllWebsitesRequest{}
	_ bin.BareEncoder = &DisconnectAllWebsitesRequest{}
	_ bin.BareDecoder = &DisconnectAllWebsitesRequest{}
)

func (d *DisconnectAllWebsitesRequest) Zero() bool {
	if d == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (d *DisconnectAllWebsitesRequest) String() string {
	if d == nil {
		return "DisconnectAllWebsitesRequest(nil)"
	}
	type Alias DisconnectAllWebsitesRequest
	return fmt.Sprintf("DisconnectAllWebsitesRequest%+v", Alias(*d))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*DisconnectAllWebsitesRequest) TypeID() uint32 {
	return DisconnectAllWebsitesRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*DisconnectAllWebsitesRequest) TypeName() string {
	return "disconnectAllWebsites"
}

// TypeInfo returns info about TL type.
func (d *DisconnectAllWebsitesRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "disconnectAllWebsites",
		ID:   DisconnectAllWebsitesRequestTypeID,
	}
	if d == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (d *DisconnectAllWebsitesRequest) Encode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't encode disconnectAllWebsites#bf72f203 as nil")
	}
	b.PutID(DisconnectAllWebsitesRequestTypeID)
	return d.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (d *DisconnectAllWebsitesRequest) EncodeBare(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't encode disconnectAllWebsites#bf72f203 as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (d *DisconnectAllWebsitesRequest) Decode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't decode disconnectAllWebsites#bf72f203 to nil")
	}
	if err := b.ConsumeID(DisconnectAllWebsitesRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode disconnectAllWebsites#bf72f203: %w", err)
	}
	return d.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (d *DisconnectAllWebsitesRequest) DecodeBare(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't decode disconnectAllWebsites#bf72f203 to nil")
	}
	return nil
}

// EncodeTDLibJSON implements jsontd.TDLibEncoder.
func (d *DisconnectAllWebsitesRequest) EncodeTDLibJSON(b jsontd.Encoder) error {
	if d == nil {
		return fmt.Errorf("can't encode disconnectAllWebsites#bf72f203 as nil")
	}
	b.ObjStart()
	b.PutID("disconnectAllWebsites")
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements jsontd.TDLibDecoder.
func (d *DisconnectAllWebsitesRequest) DecodeTDLibJSON(b jsontd.Decoder) error {
	if d == nil {
		return fmt.Errorf("can't decode disconnectAllWebsites#bf72f203 to nil")
	}

	return b.Obj(func(b jsontd.Decoder, key []byte) error {
		switch string(key) {
		case jsontd.TypeField:
			if err := b.ConsumeID("disconnectAllWebsites"); err != nil {
				return fmt.Errorf("unable to decode disconnectAllWebsites#bf72f203: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// DisconnectAllWebsites invokes method disconnectAllWebsites#bf72f203 returning error if any.
func (c *Client) DisconnectAllWebsites(ctx context.Context) error {
	var ok Ok

	request := &DisconnectAllWebsitesRequest{}
	if err := c.rpc.Invoke(ctx, request, &ok); err != nil {
		return err
	}
	return nil
}
