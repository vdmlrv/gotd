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

// MessagesTranslateTextRequest represents TL type `messages.translateText#63183030`.
// Translate a given text
//
// See https://core.telegram.org/method/messages.translateText for reference.
type MessagesTranslateTextRequest struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// If the text is a chat message, the peer ID
	//
	// Use SetPeer and GetPeer helpers.
	Peer InputPeerClass
	//
	//
	// Use SetID and GetID helpers.
	ID []int
	// The text to translate
	//
	// Use SetText and GetText helpers.
	Text []TextWithEntities
	// Two-letter ISO 639-1 language code of the language to which the message is translated
	ToLang string
}

// MessagesTranslateTextRequestTypeID is TL type id of MessagesTranslateTextRequest.
const MessagesTranslateTextRequestTypeID = 0x63183030

// Ensuring interfaces in compile-time for MessagesTranslateTextRequest.
var (
	_ bin.Encoder     = &MessagesTranslateTextRequest{}
	_ bin.Decoder     = &MessagesTranslateTextRequest{}
	_ bin.BareEncoder = &MessagesTranslateTextRequest{}
	_ bin.BareDecoder = &MessagesTranslateTextRequest{}
)

func (t *MessagesTranslateTextRequest) Zero() bool {
	if t == nil {
		return true
	}
	if !(t.Flags.Zero()) {
		return false
	}
	if !(t.Peer == nil) {
		return false
	}
	if !(t.ID == nil) {
		return false
	}
	if !(t.Text == nil) {
		return false
	}
	if !(t.ToLang == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (t *MessagesTranslateTextRequest) String() string {
	if t == nil {
		return "MessagesTranslateTextRequest(nil)"
	}
	type Alias MessagesTranslateTextRequest
	return fmt.Sprintf("MessagesTranslateTextRequest%+v", Alias(*t))
}

// FillFrom fills MessagesTranslateTextRequest from given interface.
func (t *MessagesTranslateTextRequest) FillFrom(from interface {
	GetPeer() (value InputPeerClass, ok bool)
	GetID() (value []int, ok bool)
	GetText() (value []TextWithEntities, ok bool)
	GetToLang() (value string)
}) {
	if val, ok := from.GetPeer(); ok {
		t.Peer = val
	}

	if val, ok := from.GetID(); ok {
		t.ID = val
	}

	if val, ok := from.GetText(); ok {
		t.Text = val
	}

	t.ToLang = from.GetToLang()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessagesTranslateTextRequest) TypeID() uint32 {
	return MessagesTranslateTextRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*MessagesTranslateTextRequest) TypeName() string {
	return "messages.translateText"
}

// TypeInfo returns info about TL type.
func (t *MessagesTranslateTextRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messages.translateText",
		ID:   MessagesTranslateTextRequestTypeID,
	}
	if t == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Peer",
			SchemaName: "peer",
			Null:       !t.Flags.Has(0),
		},
		{
			Name:       "ID",
			SchemaName: "id",
			Null:       !t.Flags.Has(0),
		},
		{
			Name:       "Text",
			SchemaName: "text",
			Null:       !t.Flags.Has(1),
		},
		{
			Name:       "ToLang",
			SchemaName: "to_lang",
		},
	}
	return typ
}

// SetFlags sets flags for non-zero fields.
func (t *MessagesTranslateTextRequest) SetFlags() {
	if !(t.Peer == nil) {
		t.Flags.Set(0)
	}
	if !(t.ID == nil) {
		t.Flags.Set(0)
	}
	if !(t.Text == nil) {
		t.Flags.Set(1)
	}
}

// Encode implements bin.Encoder.
func (t *MessagesTranslateTextRequest) Encode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode messages.translateText#63183030 as nil")
	}
	b.PutID(MessagesTranslateTextRequestTypeID)
	return t.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (t *MessagesTranslateTextRequest) EncodeBare(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode messages.translateText#63183030 as nil")
	}
	t.SetFlags()
	if err := t.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.translateText#63183030: field flags: %w", err)
	}
	if t.Flags.Has(0) {
		if t.Peer == nil {
			return fmt.Errorf("unable to encode messages.translateText#63183030: field peer is nil")
		}
		if err := t.Peer.Encode(b); err != nil {
			return fmt.Errorf("unable to encode messages.translateText#63183030: field peer: %w", err)
		}
	}
	if t.Flags.Has(0) {
		b.PutVectorHeader(len(t.ID))
		for _, v := range t.ID {
			b.PutInt(v)
		}
	}
	if t.Flags.Has(1) {
		b.PutVectorHeader(len(t.Text))
		for idx, v := range t.Text {
			if err := v.Encode(b); err != nil {
				return fmt.Errorf("unable to encode messages.translateText#63183030: field text element with index %d: %w", idx, err)
			}
		}
	}
	b.PutString(t.ToLang)
	return nil
}

// Decode implements bin.Decoder.
func (t *MessagesTranslateTextRequest) Decode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode messages.translateText#63183030 to nil")
	}
	if err := b.ConsumeID(MessagesTranslateTextRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.translateText#63183030: %w", err)
	}
	return t.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (t *MessagesTranslateTextRequest) DecodeBare(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode messages.translateText#63183030 to nil")
	}
	{
		if err := t.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode messages.translateText#63183030: field flags: %w", err)
		}
	}
	if t.Flags.Has(0) {
		value, err := DecodeInputPeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.translateText#63183030: field peer: %w", err)
		}
		t.Peer = value
	}
	if t.Flags.Has(0) {
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode messages.translateText#63183030: field id: %w", err)
		}

		if headerLen > 0 {
			t.ID = make([]int, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := b.Int()
			if err != nil {
				return fmt.Errorf("unable to decode messages.translateText#63183030: field id: %w", err)
			}
			t.ID = append(t.ID, value)
		}
	}
	if t.Flags.Has(1) {
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode messages.translateText#63183030: field text: %w", err)
		}

		if headerLen > 0 {
			t.Text = make([]TextWithEntities, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			var value TextWithEntities
			if err := value.Decode(b); err != nil {
				return fmt.Errorf("unable to decode messages.translateText#63183030: field text: %w", err)
			}
			t.Text = append(t.Text, value)
		}
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode messages.translateText#63183030: field to_lang: %w", err)
		}
		t.ToLang = value
	}
	return nil
}

// SetPeer sets value of Peer conditional field.
func (t *MessagesTranslateTextRequest) SetPeer(value InputPeerClass) {
	t.Flags.Set(0)
	t.Peer = value
}

// GetPeer returns value of Peer conditional field and
// boolean which is true if field was set.
func (t *MessagesTranslateTextRequest) GetPeer() (value InputPeerClass, ok bool) {
	if t == nil {
		return
	}
	if !t.Flags.Has(0) {
		return value, false
	}
	return t.Peer, true
}

// SetID sets value of ID conditional field.
func (t *MessagesTranslateTextRequest) SetID(value []int) {
	t.Flags.Set(0)
	t.ID = value
}

// GetID returns value of ID conditional field and
// boolean which is true if field was set.
func (t *MessagesTranslateTextRequest) GetID() (value []int, ok bool) {
	if t == nil {
		return
	}
	if !t.Flags.Has(0) {
		return value, false
	}
	return t.ID, true
}

// SetText sets value of Text conditional field.
func (t *MessagesTranslateTextRequest) SetText(value []TextWithEntities) {
	t.Flags.Set(1)
	t.Text = value
}

// GetText returns value of Text conditional field and
// boolean which is true if field was set.
func (t *MessagesTranslateTextRequest) GetText() (value []TextWithEntities, ok bool) {
	if t == nil {
		return
	}
	if !t.Flags.Has(1) {
		return value, false
	}
	return t.Text, true
}

// GetToLang returns value of ToLang field.
func (t *MessagesTranslateTextRequest) GetToLang() (value string) {
	if t == nil {
		return
	}
	return t.ToLang
}

// MessagesTranslateText invokes method messages.translateText#63183030 returning error if any.
// Translate a given text
//
// Possible errors:
//
//	400 INPUT_TEXT_EMPTY: The specified text is empty.
//	400 MSG_ID_INVALID: Invalid message ID provided.
//	400 PEER_ID_INVALID: The provided peer id is invalid.
//	400 TO_LANG_INVALID: The specified destination language is invalid.
//
// See https://core.telegram.org/method/messages.translateText for reference.
func (c *Client) MessagesTranslateText(ctx context.Context, request *MessagesTranslateTextRequest) (*MessagesTranslateResult, error) {
	var result MessagesTranslateResult

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
