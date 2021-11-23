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

// RequestPasswordRecoveryRequest represents TL type `requestPasswordRecovery#ff2dc552`.
type RequestPasswordRecoveryRequest struct {
}

// RequestPasswordRecoveryRequestTypeID is TL type id of RequestPasswordRecoveryRequest.
const RequestPasswordRecoveryRequestTypeID = 0xff2dc552

// Ensuring interfaces in compile-time for RequestPasswordRecoveryRequest.
var (
	_ bin.Encoder     = &RequestPasswordRecoveryRequest{}
	_ bin.Decoder     = &RequestPasswordRecoveryRequest{}
	_ bin.BareEncoder = &RequestPasswordRecoveryRequest{}
	_ bin.BareDecoder = &RequestPasswordRecoveryRequest{}
)

func (r *RequestPasswordRecoveryRequest) Zero() bool {
	if r == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (r *RequestPasswordRecoveryRequest) String() string {
	if r == nil {
		return "RequestPasswordRecoveryRequest(nil)"
	}
	type Alias RequestPasswordRecoveryRequest
	return fmt.Sprintf("RequestPasswordRecoveryRequest%+v", Alias(*r))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*RequestPasswordRecoveryRequest) TypeID() uint32 {
	return RequestPasswordRecoveryRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*RequestPasswordRecoveryRequest) TypeName() string {
	return "requestPasswordRecovery"
}

// TypeInfo returns info about TL type.
func (r *RequestPasswordRecoveryRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "requestPasswordRecovery",
		ID:   RequestPasswordRecoveryRequestTypeID,
	}
	if r == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (r *RequestPasswordRecoveryRequest) Encode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode requestPasswordRecovery#ff2dc552 as nil")
	}
	b.PutID(RequestPasswordRecoveryRequestTypeID)
	return r.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (r *RequestPasswordRecoveryRequest) EncodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode requestPasswordRecovery#ff2dc552 as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (r *RequestPasswordRecoveryRequest) Decode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode requestPasswordRecovery#ff2dc552 to nil")
	}
	if err := b.ConsumeID(RequestPasswordRecoveryRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode requestPasswordRecovery#ff2dc552: %w", err)
	}
	return r.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (r *RequestPasswordRecoveryRequest) DecodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode requestPasswordRecovery#ff2dc552 to nil")
	}
	return nil
}

// EncodeTDLibJSON implements jsontd.TDLibEncoder.
func (r *RequestPasswordRecoveryRequest) EncodeTDLibJSON(b jsontd.Encoder) error {
	if r == nil {
		return fmt.Errorf("can't encode requestPasswordRecovery#ff2dc552 as nil")
	}
	b.ObjStart()
	b.PutID("requestPasswordRecovery")
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements jsontd.TDLibDecoder.
func (r *RequestPasswordRecoveryRequest) DecodeTDLibJSON(b jsontd.Decoder) error {
	if r == nil {
		return fmt.Errorf("can't decode requestPasswordRecovery#ff2dc552 to nil")
	}

	return b.Obj(func(b jsontd.Decoder, key []byte) error {
		switch string(key) {
		case jsontd.TypeField:
			if err := b.ConsumeID("requestPasswordRecovery"); err != nil {
				return fmt.Errorf("unable to decode requestPasswordRecovery#ff2dc552: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// RequestPasswordRecovery invokes method requestPasswordRecovery#ff2dc552 returning error if any.
func (c *Client) RequestPasswordRecovery(ctx context.Context) (*EmailAddressAuthenticationCodeInfo, error) {
	var result EmailAddressAuthenticationCodeInfo

	request := &RequestPasswordRecoveryRequest{}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
