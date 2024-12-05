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

// AffiliateProgramInfo represents TL type `affiliateProgramInfo#96fce8b5`.
type AffiliateProgramInfo struct {
	// Parameters of the affiliate program
	Parameters AffiliateProgramParameters
	// Point in time (Unix timestamp) when the affiliate program will be closed; 0 if the
	// affiliate program isn't scheduled to be closed.
	EndDate int32
	// The amount of daily revenue per user in Telegram Stars of the bot that created the
	// affiliate program
	DailyRevenuePerUserAmount StarAmount
}

// AffiliateProgramInfoTypeID is TL type id of AffiliateProgramInfo.
const AffiliateProgramInfoTypeID = 0x96fce8b5

// Ensuring interfaces in compile-time for AffiliateProgramInfo.
var (
	_ bin.Encoder     = &AffiliateProgramInfo{}
	_ bin.Decoder     = &AffiliateProgramInfo{}
	_ bin.BareEncoder = &AffiliateProgramInfo{}
	_ bin.BareDecoder = &AffiliateProgramInfo{}
)

func (a *AffiliateProgramInfo) Zero() bool {
	if a == nil {
		return true
	}
	if !(a.Parameters.Zero()) {
		return false
	}
	if !(a.EndDate == 0) {
		return false
	}
	if !(a.DailyRevenuePerUserAmount.Zero()) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (a *AffiliateProgramInfo) String() string {
	if a == nil {
		return "AffiliateProgramInfo(nil)"
	}
	type Alias AffiliateProgramInfo
	return fmt.Sprintf("AffiliateProgramInfo%+v", Alias(*a))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*AffiliateProgramInfo) TypeID() uint32 {
	return AffiliateProgramInfoTypeID
}

// TypeName returns name of type in TL schema.
func (*AffiliateProgramInfo) TypeName() string {
	return "affiliateProgramInfo"
}

// TypeInfo returns info about TL type.
func (a *AffiliateProgramInfo) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "affiliateProgramInfo",
		ID:   AffiliateProgramInfoTypeID,
	}
	if a == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Parameters",
			SchemaName: "parameters",
		},
		{
			Name:       "EndDate",
			SchemaName: "end_date",
		},
		{
			Name:       "DailyRevenuePerUserAmount",
			SchemaName: "daily_revenue_per_user_amount",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (a *AffiliateProgramInfo) Encode(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't encode affiliateProgramInfo#96fce8b5 as nil")
	}
	b.PutID(AffiliateProgramInfoTypeID)
	return a.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (a *AffiliateProgramInfo) EncodeBare(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't encode affiliateProgramInfo#96fce8b5 as nil")
	}
	if err := a.Parameters.Encode(b); err != nil {
		return fmt.Errorf("unable to encode affiliateProgramInfo#96fce8b5: field parameters: %w", err)
	}
	b.PutInt32(a.EndDate)
	if err := a.DailyRevenuePerUserAmount.Encode(b); err != nil {
		return fmt.Errorf("unable to encode affiliateProgramInfo#96fce8b5: field daily_revenue_per_user_amount: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (a *AffiliateProgramInfo) Decode(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't decode affiliateProgramInfo#96fce8b5 to nil")
	}
	if err := b.ConsumeID(AffiliateProgramInfoTypeID); err != nil {
		return fmt.Errorf("unable to decode affiliateProgramInfo#96fce8b5: %w", err)
	}
	return a.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (a *AffiliateProgramInfo) DecodeBare(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't decode affiliateProgramInfo#96fce8b5 to nil")
	}
	{
		if err := a.Parameters.Decode(b); err != nil {
			return fmt.Errorf("unable to decode affiliateProgramInfo#96fce8b5: field parameters: %w", err)
		}
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode affiliateProgramInfo#96fce8b5: field end_date: %w", err)
		}
		a.EndDate = value
	}
	{
		if err := a.DailyRevenuePerUserAmount.Decode(b); err != nil {
			return fmt.Errorf("unable to decode affiliateProgramInfo#96fce8b5: field daily_revenue_per_user_amount: %w", err)
		}
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (a *AffiliateProgramInfo) EncodeTDLibJSON(b tdjson.Encoder) error {
	if a == nil {
		return fmt.Errorf("can't encode affiliateProgramInfo#96fce8b5 as nil")
	}
	b.ObjStart()
	b.PutID("affiliateProgramInfo")
	b.Comma()
	b.FieldStart("parameters")
	if err := a.Parameters.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode affiliateProgramInfo#96fce8b5: field parameters: %w", err)
	}
	b.Comma()
	b.FieldStart("end_date")
	b.PutInt32(a.EndDate)
	b.Comma()
	b.FieldStart("daily_revenue_per_user_amount")
	if err := a.DailyRevenuePerUserAmount.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode affiliateProgramInfo#96fce8b5: field daily_revenue_per_user_amount: %w", err)
	}
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (a *AffiliateProgramInfo) DecodeTDLibJSON(b tdjson.Decoder) error {
	if a == nil {
		return fmt.Errorf("can't decode affiliateProgramInfo#96fce8b5 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("affiliateProgramInfo"); err != nil {
				return fmt.Errorf("unable to decode affiliateProgramInfo#96fce8b5: %w", err)
			}
		case "parameters":
			if err := a.Parameters.DecodeTDLibJSON(b); err != nil {
				return fmt.Errorf("unable to decode affiliateProgramInfo#96fce8b5: field parameters: %w", err)
			}
		case "end_date":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode affiliateProgramInfo#96fce8b5: field end_date: %w", err)
			}
			a.EndDate = value
		case "daily_revenue_per_user_amount":
			if err := a.DailyRevenuePerUserAmount.DecodeTDLibJSON(b); err != nil {
				return fmt.Errorf("unable to decode affiliateProgramInfo#96fce8b5: field daily_revenue_per_user_amount: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetParameters returns value of Parameters field.
func (a *AffiliateProgramInfo) GetParameters() (value AffiliateProgramParameters) {
	if a == nil {
		return
	}
	return a.Parameters
}

// GetEndDate returns value of EndDate field.
func (a *AffiliateProgramInfo) GetEndDate() (value int32) {
	if a == nil {
		return
	}
	return a.EndDate
}

// GetDailyRevenuePerUserAmount returns value of DailyRevenuePerUserAmount field.
func (a *AffiliateProgramInfo) GetDailyRevenuePerUserAmount() (value StarAmount) {
	if a == nil {
		return
	}
	return a.DailyRevenuePerUserAmount
}
