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

// AddContactRequest represents TL type `addContact#6f707140`.
type AddContactRequest struct {
	// The contact to add or edit; phone number can be empty and needs to be specified only
	// if known, vCard is ignored
	Contact Contact
	// True, if the new contact needs to be allowed to see current user's phone number. A
	// corresponding rule to userPrivacySettingShowPhoneNumber will be added if needed. Use
	// the field UserFullInfo.need_phone_number_privacy_exception to check whether the
	// current user needs to be asked to share their phone number
	SharePhoneNumber bool
}

// AddContactRequestTypeID is TL type id of AddContactRequest.
const AddContactRequestTypeID = 0x6f707140

// Ensuring interfaces in compile-time for AddContactRequest.
var (
	_ bin.Encoder     = &AddContactRequest{}
	_ bin.Decoder     = &AddContactRequest{}
	_ bin.BareEncoder = &AddContactRequest{}
	_ bin.BareDecoder = &AddContactRequest{}
)

func (a *AddContactRequest) Zero() bool {
	if a == nil {
		return true
	}
	if !(a.Contact.Zero()) {
		return false
	}
	if !(a.SharePhoneNumber == false) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (a *AddContactRequest) String() string {
	if a == nil {
		return "AddContactRequest(nil)"
	}
	type Alias AddContactRequest
	return fmt.Sprintf("AddContactRequest%+v", Alias(*a))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*AddContactRequest) TypeID() uint32 {
	return AddContactRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*AddContactRequest) TypeName() string {
	return "addContact"
}

// TypeInfo returns info about TL type.
func (a *AddContactRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "addContact",
		ID:   AddContactRequestTypeID,
	}
	if a == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Contact",
			SchemaName: "contact",
		},
		{
			Name:       "SharePhoneNumber",
			SchemaName: "share_phone_number",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (a *AddContactRequest) Encode(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't encode addContact#6f707140 as nil")
	}
	b.PutID(AddContactRequestTypeID)
	return a.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (a *AddContactRequest) EncodeBare(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't encode addContact#6f707140 as nil")
	}
	if err := a.Contact.Encode(b); err != nil {
		return fmt.Errorf("unable to encode addContact#6f707140: field contact: %w", err)
	}
	b.PutBool(a.SharePhoneNumber)
	return nil
}

// Decode implements bin.Decoder.
func (a *AddContactRequest) Decode(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't decode addContact#6f707140 to nil")
	}
	if err := b.ConsumeID(AddContactRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode addContact#6f707140: %w", err)
	}
	return a.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (a *AddContactRequest) DecodeBare(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't decode addContact#6f707140 to nil")
	}
	{
		if err := a.Contact.Decode(b); err != nil {
			return fmt.Errorf("unable to decode addContact#6f707140: field contact: %w", err)
		}
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode addContact#6f707140: field share_phone_number: %w", err)
		}
		a.SharePhoneNumber = value
	}
	return nil
}

// EncodeTDLibJSON implements jsontd.TDLibEncoder.
func (a *AddContactRequest) EncodeTDLibJSON(b jsontd.Encoder) error {
	if a == nil {
		return fmt.Errorf("can't encode addContact#6f707140 as nil")
	}
	b.ObjStart()
	b.PutID("addContact")
	b.FieldStart("contact")
	if err := a.Contact.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode addContact#6f707140: field contact: %w", err)
	}
	b.FieldStart("share_phone_number")
	b.PutBool(a.SharePhoneNumber)
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements jsontd.TDLibDecoder.
func (a *AddContactRequest) DecodeTDLibJSON(b jsontd.Decoder) error {
	if a == nil {
		return fmt.Errorf("can't decode addContact#6f707140 to nil")
	}

	return b.Obj(func(b jsontd.Decoder, key []byte) error {
		switch string(key) {
		case jsontd.TypeField:
			if err := b.ConsumeID("addContact"); err != nil {
				return fmt.Errorf("unable to decode addContact#6f707140: %w", err)
			}
		case "contact":
			if err := a.Contact.DecodeTDLibJSON(b); err != nil {
				return fmt.Errorf("unable to decode addContact#6f707140: field contact: %w", err)
			}
		case "share_phone_number":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode addContact#6f707140: field share_phone_number: %w", err)
			}
			a.SharePhoneNumber = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetContact returns value of Contact field.
func (a *AddContactRequest) GetContact() (value Contact) {
	return a.Contact
}

// GetSharePhoneNumber returns value of SharePhoneNumber field.
func (a *AddContactRequest) GetSharePhoneNumber() (value bool) {
	return a.SharePhoneNumber
}

// AddContact invokes method addContact#6f707140 returning error if any.
func (c *Client) AddContact(ctx context.Context, request *AddContactRequest) error {
	var ok Ok

	if err := c.rpc.Invoke(ctx, request, &ok); err != nil {
		return err
	}
	return nil
}
