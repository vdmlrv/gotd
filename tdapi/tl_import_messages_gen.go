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

// ImportMessagesRequest represents TL type `importMessages#7e98592b`.
type ImportMessagesRequest struct {
	// Identifier of a chat to which the messages will be imported. It must be an identifier
	// of a private chat with a mutual contact or an identifier of a supergroup chat with
	// can_change_info administrator right
	ChatID int64
	// File with messages to import. Only inputFileLocal and inputFileGenerated are supported
	// The file must not be previously uploaded
	MessageFile InputFileClass
	// Files used in the imported messages. Only inputFileLocal and inputFileGenerated are
	// supported. The files must not be previously uploaded
	AttachedFiles []InputFileClass
}

// ImportMessagesRequestTypeID is TL type id of ImportMessagesRequest.
const ImportMessagesRequestTypeID = 0x7e98592b

// Ensuring interfaces in compile-time for ImportMessagesRequest.
var (
	_ bin.Encoder     = &ImportMessagesRequest{}
	_ bin.Decoder     = &ImportMessagesRequest{}
	_ bin.BareEncoder = &ImportMessagesRequest{}
	_ bin.BareDecoder = &ImportMessagesRequest{}
)

func (i *ImportMessagesRequest) Zero() bool {
	if i == nil {
		return true
	}
	if !(i.ChatID == 0) {
		return false
	}
	if !(i.MessageFile == nil) {
		return false
	}
	if !(i.AttachedFiles == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (i *ImportMessagesRequest) String() string {
	if i == nil {
		return "ImportMessagesRequest(nil)"
	}
	type Alias ImportMessagesRequest
	return fmt.Sprintf("ImportMessagesRequest%+v", Alias(*i))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ImportMessagesRequest) TypeID() uint32 {
	return ImportMessagesRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*ImportMessagesRequest) TypeName() string {
	return "importMessages"
}

// TypeInfo returns info about TL type.
func (i *ImportMessagesRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "importMessages",
		ID:   ImportMessagesRequestTypeID,
	}
	if i == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ChatID",
			SchemaName: "chat_id",
		},
		{
			Name:       "MessageFile",
			SchemaName: "message_file",
		},
		{
			Name:       "AttachedFiles",
			SchemaName: "attached_files",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (i *ImportMessagesRequest) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode importMessages#7e98592b as nil")
	}
	b.PutID(ImportMessagesRequestTypeID)
	return i.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (i *ImportMessagesRequest) EncodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode importMessages#7e98592b as nil")
	}
	b.PutLong(i.ChatID)
	if i.MessageFile == nil {
		return fmt.Errorf("unable to encode importMessages#7e98592b: field message_file is nil")
	}
	if err := i.MessageFile.Encode(b); err != nil {
		return fmt.Errorf("unable to encode importMessages#7e98592b: field message_file: %w", err)
	}
	b.PutInt(len(i.AttachedFiles))
	for idx, v := range i.AttachedFiles {
		if v == nil {
			return fmt.Errorf("unable to encode importMessages#7e98592b: field attached_files element with index %d is nil", idx)
		}
		if err := v.EncodeBare(b); err != nil {
			return fmt.Errorf("unable to encode bare importMessages#7e98592b: field attached_files element with index %d: %w", idx, err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (i *ImportMessagesRequest) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode importMessages#7e98592b to nil")
	}
	if err := b.ConsumeID(ImportMessagesRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode importMessages#7e98592b: %w", err)
	}
	return i.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (i *ImportMessagesRequest) DecodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode importMessages#7e98592b to nil")
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode importMessages#7e98592b: field chat_id: %w", err)
		}
		i.ChatID = value
	}
	{
		value, err := DecodeInputFile(b)
		if err != nil {
			return fmt.Errorf("unable to decode importMessages#7e98592b: field message_file: %w", err)
		}
		i.MessageFile = value
	}
	{
		headerLen, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode importMessages#7e98592b: field attached_files: %w", err)
		}

		if headerLen > 0 {
			i.AttachedFiles = make([]InputFileClass, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeInputFile(b)
			if err != nil {
				return fmt.Errorf("unable to decode importMessages#7e98592b: field attached_files: %w", err)
			}
			i.AttachedFiles = append(i.AttachedFiles, value)
		}
	}
	return nil
}

// EncodeTDLibJSON implements jsontd.TDLibEncoder.
func (i *ImportMessagesRequest) EncodeTDLibJSON(b jsontd.Encoder) error {
	if i == nil {
		return fmt.Errorf("can't encode importMessages#7e98592b as nil")
	}
	b.ObjStart()
	b.PutID("importMessages")
	b.FieldStart("chat_id")
	b.PutLong(i.ChatID)
	b.FieldStart("message_file")
	if i.MessageFile == nil {
		return fmt.Errorf("unable to encode importMessages#7e98592b: field message_file is nil")
	}
	if err := i.MessageFile.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode importMessages#7e98592b: field message_file: %w", err)
	}
	b.FieldStart("attached_files")
	b.ArrStart()
	for idx, v := range i.AttachedFiles {
		if v == nil {
			return fmt.Errorf("unable to encode importMessages#7e98592b: field attached_files element with index %d is nil", idx)
		}
		if err := v.EncodeTDLibJSON(b); err != nil {
			return fmt.Errorf("unable to encode importMessages#7e98592b: field attached_files element with index %d: %w", idx, err)
		}
	}
	b.ArrEnd()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements jsontd.TDLibDecoder.
func (i *ImportMessagesRequest) DecodeTDLibJSON(b jsontd.Decoder) error {
	if i == nil {
		return fmt.Errorf("can't decode importMessages#7e98592b to nil")
	}

	return b.Obj(func(b jsontd.Decoder, key []byte) error {
		switch string(key) {
		case jsontd.TypeField:
			if err := b.ConsumeID("importMessages"); err != nil {
				return fmt.Errorf("unable to decode importMessages#7e98592b: %w", err)
			}
		case "chat_id":
			value, err := b.Long()
			if err != nil {
				return fmt.Errorf("unable to decode importMessages#7e98592b: field chat_id: %w", err)
			}
			i.ChatID = value
		case "message_file":
			value, err := DecodeTDLibJSONInputFile(b)
			if err != nil {
				return fmt.Errorf("unable to decode importMessages#7e98592b: field message_file: %w", err)
			}
			i.MessageFile = value
		case "attached_files":
			if err := b.Arr(func(b jsontd.Decoder) error {
				value, err := DecodeTDLibJSONInputFile(b)
				if err != nil {
					return fmt.Errorf("unable to decode importMessages#7e98592b: field attached_files: %w", err)
				}
				i.AttachedFiles = append(i.AttachedFiles, value)
				return nil
			}); err != nil {
				return fmt.Errorf("unable to decode importMessages#7e98592b: field attached_files: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetChatID returns value of ChatID field.
func (i *ImportMessagesRequest) GetChatID() (value int64) {
	return i.ChatID
}

// GetMessageFile returns value of MessageFile field.
func (i *ImportMessagesRequest) GetMessageFile() (value InputFileClass) {
	return i.MessageFile
}

// GetAttachedFiles returns value of AttachedFiles field.
func (i *ImportMessagesRequest) GetAttachedFiles() (value []InputFileClass) {
	return i.AttachedFiles
}

// ImportMessages invokes method importMessages#7e98592b returning error if any.
func (c *Client) ImportMessages(ctx context.Context, request *ImportMessagesRequest) error {
	var ok Ok

	if err := c.rpc.Invoke(ctx, request, &ok); err != nil {
		return err
	}
	return nil
}
