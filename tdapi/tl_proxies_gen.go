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

// Proxies represents TL type `proxies#5ee27a86`.
type Proxies struct {
	// List of proxy servers
	Proxies []Proxy
}

// ProxiesTypeID is TL type id of Proxies.
const ProxiesTypeID = 0x5ee27a86

// Ensuring interfaces in compile-time for Proxies.
var (
	_ bin.Encoder     = &Proxies{}
	_ bin.Decoder     = &Proxies{}
	_ bin.BareEncoder = &Proxies{}
	_ bin.BareDecoder = &Proxies{}
)

func (p *Proxies) Zero() bool {
	if p == nil {
		return true
	}
	if !(p.Proxies == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (p *Proxies) String() string {
	if p == nil {
		return "Proxies(nil)"
	}
	type Alias Proxies
	return fmt.Sprintf("Proxies%+v", Alias(*p))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*Proxies) TypeID() uint32 {
	return ProxiesTypeID
}

// TypeName returns name of type in TL schema.
func (*Proxies) TypeName() string {
	return "proxies"
}

// TypeInfo returns info about TL type.
func (p *Proxies) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "proxies",
		ID:   ProxiesTypeID,
	}
	if p == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Proxies",
			SchemaName: "proxies",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (p *Proxies) Encode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode proxies#5ee27a86 as nil")
	}
	b.PutID(ProxiesTypeID)
	return p.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (p *Proxies) EncodeBare(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode proxies#5ee27a86 as nil")
	}
	b.PutInt(len(p.Proxies))
	for idx, v := range p.Proxies {
		if err := v.EncodeBare(b); err != nil {
			return fmt.Errorf("unable to encode bare proxies#5ee27a86: field proxies element with index %d: %w", idx, err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (p *Proxies) Decode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode proxies#5ee27a86 to nil")
	}
	if err := b.ConsumeID(ProxiesTypeID); err != nil {
		return fmt.Errorf("unable to decode proxies#5ee27a86: %w", err)
	}
	return p.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (p *Proxies) DecodeBare(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode proxies#5ee27a86 to nil")
	}
	{
		headerLen, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode proxies#5ee27a86: field proxies: %w", err)
		}

		if headerLen > 0 {
			p.Proxies = make([]Proxy, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			var value Proxy
			if err := value.DecodeBare(b); err != nil {
				return fmt.Errorf("unable to decode bare proxies#5ee27a86: field proxies: %w", err)
			}
			p.Proxies = append(p.Proxies, value)
		}
	}
	return nil
}

// EncodeTDLibJSON implements jsontd.TDLibEncoder.
func (p *Proxies) EncodeTDLibJSON(b jsontd.Encoder) error {
	if p == nil {
		return fmt.Errorf("can't encode proxies#5ee27a86 as nil")
	}
	b.ObjStart()
	b.PutID("proxies")
	b.FieldStart("proxies")
	b.ArrStart()
	for idx, v := range p.Proxies {
		if err := v.EncodeTDLibJSON(b); err != nil {
			return fmt.Errorf("unable to encode proxies#5ee27a86: field proxies element with index %d: %w", idx, err)
		}
	}
	b.ArrEnd()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements jsontd.TDLibDecoder.
func (p *Proxies) DecodeTDLibJSON(b jsontd.Decoder) error {
	if p == nil {
		return fmt.Errorf("can't decode proxies#5ee27a86 to nil")
	}

	return b.Obj(func(b jsontd.Decoder, key []byte) error {
		switch string(key) {
		case jsontd.TypeField:
			if err := b.ConsumeID("proxies"); err != nil {
				return fmt.Errorf("unable to decode proxies#5ee27a86: %w", err)
			}
		case "proxies":
			if err := b.Arr(func(b jsontd.Decoder) error {
				var value Proxy
				if err := value.DecodeTDLibJSON(b); err != nil {
					return fmt.Errorf("unable to decode proxies#5ee27a86: field proxies: %w", err)
				}
				p.Proxies = append(p.Proxies, value)
				return nil
			}); err != nil {
				return fmt.Errorf("unable to decode proxies#5ee27a86: field proxies: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetProxies returns value of Proxies field.
func (p *Proxies) GetProxies() (value []Proxy) {
	return p.Proxies
}
