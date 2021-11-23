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

// PhoneNumberInfo represents TL type `phoneNumberInfo#2163aee1`.
type PhoneNumberInfo struct {
	// Information about the country to which the phone number belongs; may be null
	Country CountryInfo
	// The part of the phone number denoting country calling code or its part
	CountryCallingCode string
	// The phone number without country calling code formatted accordingly to local rules
	FormattedPhoneNumber string
}

// PhoneNumberInfoTypeID is TL type id of PhoneNumberInfo.
const PhoneNumberInfoTypeID = 0x2163aee1

// Ensuring interfaces in compile-time for PhoneNumberInfo.
var (
	_ bin.Encoder     = &PhoneNumberInfo{}
	_ bin.Decoder     = &PhoneNumberInfo{}
	_ bin.BareEncoder = &PhoneNumberInfo{}
	_ bin.BareDecoder = &PhoneNumberInfo{}
)

func (p *PhoneNumberInfo) Zero() bool {
	if p == nil {
		return true
	}
	if !(p.Country.Zero()) {
		return false
	}
	if !(p.CountryCallingCode == "") {
		return false
	}
	if !(p.FormattedPhoneNumber == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (p *PhoneNumberInfo) String() string {
	if p == nil {
		return "PhoneNumberInfo(nil)"
	}
	type Alias PhoneNumberInfo
	return fmt.Sprintf("PhoneNumberInfo%+v", Alias(*p))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*PhoneNumberInfo) TypeID() uint32 {
	return PhoneNumberInfoTypeID
}

// TypeName returns name of type in TL schema.
func (*PhoneNumberInfo) TypeName() string {
	return "phoneNumberInfo"
}

// TypeInfo returns info about TL type.
func (p *PhoneNumberInfo) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "phoneNumberInfo",
		ID:   PhoneNumberInfoTypeID,
	}
	if p == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Country",
			SchemaName: "country",
		},
		{
			Name:       "CountryCallingCode",
			SchemaName: "country_calling_code",
		},
		{
			Name:       "FormattedPhoneNumber",
			SchemaName: "formatted_phone_number",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (p *PhoneNumberInfo) Encode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode phoneNumberInfo#2163aee1 as nil")
	}
	b.PutID(PhoneNumberInfoTypeID)
	return p.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (p *PhoneNumberInfo) EncodeBare(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode phoneNumberInfo#2163aee1 as nil")
	}
	if err := p.Country.Encode(b); err != nil {
		return fmt.Errorf("unable to encode phoneNumberInfo#2163aee1: field country: %w", err)
	}
	b.PutString(p.CountryCallingCode)
	b.PutString(p.FormattedPhoneNumber)
	return nil
}

// Decode implements bin.Decoder.
func (p *PhoneNumberInfo) Decode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode phoneNumberInfo#2163aee1 to nil")
	}
	if err := b.ConsumeID(PhoneNumberInfoTypeID); err != nil {
		return fmt.Errorf("unable to decode phoneNumberInfo#2163aee1: %w", err)
	}
	return p.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (p *PhoneNumberInfo) DecodeBare(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode phoneNumberInfo#2163aee1 to nil")
	}
	{
		if err := p.Country.Decode(b); err != nil {
			return fmt.Errorf("unable to decode phoneNumberInfo#2163aee1: field country: %w", err)
		}
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode phoneNumberInfo#2163aee1: field country_calling_code: %w", err)
		}
		p.CountryCallingCode = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode phoneNumberInfo#2163aee1: field formatted_phone_number: %w", err)
		}
		p.FormattedPhoneNumber = value
	}
	return nil
}

// EncodeTDLibJSON implements jsontd.TDLibEncoder.
func (p *PhoneNumberInfo) EncodeTDLibJSON(b jsontd.Encoder) error {
	if p == nil {
		return fmt.Errorf("can't encode phoneNumberInfo#2163aee1 as nil")
	}
	b.ObjStart()
	b.PutID("phoneNumberInfo")
	b.FieldStart("country")
	if err := p.Country.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode phoneNumberInfo#2163aee1: field country: %w", err)
	}
	b.FieldStart("country_calling_code")
	b.PutString(p.CountryCallingCode)
	b.FieldStart("formatted_phone_number")
	b.PutString(p.FormattedPhoneNumber)
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements jsontd.TDLibDecoder.
func (p *PhoneNumberInfo) DecodeTDLibJSON(b jsontd.Decoder) error {
	if p == nil {
		return fmt.Errorf("can't decode phoneNumberInfo#2163aee1 to nil")
	}

	return b.Obj(func(b jsontd.Decoder, key []byte) error {
		switch string(key) {
		case jsontd.TypeField:
			if err := b.ConsumeID("phoneNumberInfo"); err != nil {
				return fmt.Errorf("unable to decode phoneNumberInfo#2163aee1: %w", err)
			}
		case "country":
			if err := p.Country.DecodeTDLibJSON(b); err != nil {
				return fmt.Errorf("unable to decode phoneNumberInfo#2163aee1: field country: %w", err)
			}
		case "country_calling_code":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode phoneNumberInfo#2163aee1: field country_calling_code: %w", err)
			}
			p.CountryCallingCode = value
		case "formatted_phone_number":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode phoneNumberInfo#2163aee1: field formatted_phone_number: %w", err)
			}
			p.FormattedPhoneNumber = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetCountry returns value of Country field.
func (p *PhoneNumberInfo) GetCountry() (value CountryInfo) {
	return p.Country
}

// GetCountryCallingCode returns value of CountryCallingCode field.
func (p *PhoneNumberInfo) GetCountryCallingCode() (value string) {
	return p.CountryCallingCode
}

// GetFormattedPhoneNumber returns value of FormattedPhoneNumber field.
func (p *PhoneNumberInfo) GetFormattedPhoneNumber() (value string) {
	return p.FormattedPhoneNumber
}
