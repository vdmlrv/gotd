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

// DeleteSavedOrderInfoRequest represents TL type `deleteSavedOrderInfo#61197474`.
type DeleteSavedOrderInfoRequest struct {
}

// DeleteSavedOrderInfoRequestTypeID is TL type id of DeleteSavedOrderInfoRequest.
const DeleteSavedOrderInfoRequestTypeID = 0x61197474

// Ensuring interfaces in compile-time for DeleteSavedOrderInfoRequest.
var (
	_ bin.Encoder     = &DeleteSavedOrderInfoRequest{}
	_ bin.Decoder     = &DeleteSavedOrderInfoRequest{}
	_ bin.BareEncoder = &DeleteSavedOrderInfoRequest{}
	_ bin.BareDecoder = &DeleteSavedOrderInfoRequest{}
)

func (d *DeleteSavedOrderInfoRequest) Zero() bool {
	if d == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (d *DeleteSavedOrderInfoRequest) String() string {
	if d == nil {
		return "DeleteSavedOrderInfoRequest(nil)"
	}
	type Alias DeleteSavedOrderInfoRequest
	return fmt.Sprintf("DeleteSavedOrderInfoRequest%+v", Alias(*d))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*DeleteSavedOrderInfoRequest) TypeID() uint32 {
	return DeleteSavedOrderInfoRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*DeleteSavedOrderInfoRequest) TypeName() string {
	return "deleteSavedOrderInfo"
}

// TypeInfo returns info about TL type.
func (d *DeleteSavedOrderInfoRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "deleteSavedOrderInfo",
		ID:   DeleteSavedOrderInfoRequestTypeID,
	}
	if d == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (d *DeleteSavedOrderInfoRequest) Encode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't encode deleteSavedOrderInfo#61197474 as nil")
	}
	b.PutID(DeleteSavedOrderInfoRequestTypeID)
	return d.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (d *DeleteSavedOrderInfoRequest) EncodeBare(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't encode deleteSavedOrderInfo#61197474 as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (d *DeleteSavedOrderInfoRequest) Decode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't decode deleteSavedOrderInfo#61197474 to nil")
	}
	if err := b.ConsumeID(DeleteSavedOrderInfoRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode deleteSavedOrderInfo#61197474: %w", err)
	}
	return d.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (d *DeleteSavedOrderInfoRequest) DecodeBare(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't decode deleteSavedOrderInfo#61197474 to nil")
	}
	return nil
}

// EncodeTDLibJSON implements jsontd.TDLibEncoder.
func (d *DeleteSavedOrderInfoRequest) EncodeTDLibJSON(b jsontd.Encoder) error {
	if d == nil {
		return fmt.Errorf("can't encode deleteSavedOrderInfo#61197474 as nil")
	}
	b.ObjStart()
	b.PutID("deleteSavedOrderInfo")
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements jsontd.TDLibDecoder.
func (d *DeleteSavedOrderInfoRequest) DecodeTDLibJSON(b jsontd.Decoder) error {
	if d == nil {
		return fmt.Errorf("can't decode deleteSavedOrderInfo#61197474 to nil")
	}

	return b.Obj(func(b jsontd.Decoder, key []byte) error {
		switch string(key) {
		case jsontd.TypeField:
			if err := b.ConsumeID("deleteSavedOrderInfo"); err != nil {
				return fmt.Errorf("unable to decode deleteSavedOrderInfo#61197474: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// DeleteSavedOrderInfo invokes method deleteSavedOrderInfo#61197474 returning error if any.
func (c *Client) DeleteSavedOrderInfo(ctx context.Context) error {
	var ok Ok

	request := &DeleteSavedOrderInfoRequest{}
	if err := c.rpc.Invoke(ctx, request, &ok); err != nil {
		return err
	}
	return nil
}
