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

// DatabaseStatistics represents TL type `databaseStatistics#bd027350`.
type DatabaseStatistics struct {
	// Database statistics in an unspecified human-readable format
	Statistics string
}

// DatabaseStatisticsTypeID is TL type id of DatabaseStatistics.
const DatabaseStatisticsTypeID = 0xbd027350

// Ensuring interfaces in compile-time for DatabaseStatistics.
var (
	_ bin.Encoder     = &DatabaseStatistics{}
	_ bin.Decoder     = &DatabaseStatistics{}
	_ bin.BareEncoder = &DatabaseStatistics{}
	_ bin.BareDecoder = &DatabaseStatistics{}
)

func (d *DatabaseStatistics) Zero() bool {
	if d == nil {
		return true
	}
	if !(d.Statistics == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (d *DatabaseStatistics) String() string {
	if d == nil {
		return "DatabaseStatistics(nil)"
	}
	type Alias DatabaseStatistics
	return fmt.Sprintf("DatabaseStatistics%+v", Alias(*d))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*DatabaseStatistics) TypeID() uint32 {
	return DatabaseStatisticsTypeID
}

// TypeName returns name of type in TL schema.
func (*DatabaseStatistics) TypeName() string {
	return "databaseStatistics"
}

// TypeInfo returns info about TL type.
func (d *DatabaseStatistics) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "databaseStatistics",
		ID:   DatabaseStatisticsTypeID,
	}
	if d == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Statistics",
			SchemaName: "statistics",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (d *DatabaseStatistics) Encode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't encode databaseStatistics#bd027350 as nil")
	}
	b.PutID(DatabaseStatisticsTypeID)
	return d.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (d *DatabaseStatistics) EncodeBare(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't encode databaseStatistics#bd027350 as nil")
	}
	b.PutString(d.Statistics)
	return nil
}

// Decode implements bin.Decoder.
func (d *DatabaseStatistics) Decode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't decode databaseStatistics#bd027350 to nil")
	}
	if err := b.ConsumeID(DatabaseStatisticsTypeID); err != nil {
		return fmt.Errorf("unable to decode databaseStatistics#bd027350: %w", err)
	}
	return d.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (d *DatabaseStatistics) DecodeBare(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't decode databaseStatistics#bd027350 to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode databaseStatistics#bd027350: field statistics: %w", err)
		}
		d.Statistics = value
	}
	return nil
}

// EncodeTDLibJSON implements jsontd.TDLibEncoder.
func (d *DatabaseStatistics) EncodeTDLibJSON(b jsontd.Encoder) error {
	if d == nil {
		return fmt.Errorf("can't encode databaseStatistics#bd027350 as nil")
	}
	b.ObjStart()
	b.PutID("databaseStatistics")
	b.FieldStart("statistics")
	b.PutString(d.Statistics)
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements jsontd.TDLibDecoder.
func (d *DatabaseStatistics) DecodeTDLibJSON(b jsontd.Decoder) error {
	if d == nil {
		return fmt.Errorf("can't decode databaseStatistics#bd027350 to nil")
	}

	return b.Obj(func(b jsontd.Decoder, key []byte) error {
		switch string(key) {
		case jsontd.TypeField:
			if err := b.ConsumeID("databaseStatistics"); err != nil {
				return fmt.Errorf("unable to decode databaseStatistics#bd027350: %w", err)
			}
		case "statistics":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode databaseStatistics#bd027350: field statistics: %w", err)
			}
			d.Statistics = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetStatistics returns value of Statistics field.
func (d *DatabaseStatistics) GetStatistics() (value string) {
	return d.Statistics
}
