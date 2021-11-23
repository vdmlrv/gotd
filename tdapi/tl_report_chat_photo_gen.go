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

// ReportChatPhotoRequest represents TL type `reportChatPhoto#2bc9e924`.
type ReportChatPhotoRequest struct {
	// Chat identifier
	ChatID int64
	// Identifier of the photo to report. Only full photos from chatPhoto can be reported
	FileID int32
	// The reason for reporting the chat photo
	Reason ChatReportReasonClass
	// Additional report details; 0-1024 characters
	Text string
}

// ReportChatPhotoRequestTypeID is TL type id of ReportChatPhotoRequest.
const ReportChatPhotoRequestTypeID = 0x2bc9e924

// Ensuring interfaces in compile-time for ReportChatPhotoRequest.
var (
	_ bin.Encoder     = &ReportChatPhotoRequest{}
	_ bin.Decoder     = &ReportChatPhotoRequest{}
	_ bin.BareEncoder = &ReportChatPhotoRequest{}
	_ bin.BareDecoder = &ReportChatPhotoRequest{}
)

func (r *ReportChatPhotoRequest) Zero() bool {
	if r == nil {
		return true
	}
	if !(r.ChatID == 0) {
		return false
	}
	if !(r.FileID == 0) {
		return false
	}
	if !(r.Reason == nil) {
		return false
	}
	if !(r.Text == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (r *ReportChatPhotoRequest) String() string {
	if r == nil {
		return "ReportChatPhotoRequest(nil)"
	}
	type Alias ReportChatPhotoRequest
	return fmt.Sprintf("ReportChatPhotoRequest%+v", Alias(*r))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ReportChatPhotoRequest) TypeID() uint32 {
	return ReportChatPhotoRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*ReportChatPhotoRequest) TypeName() string {
	return "reportChatPhoto"
}

// TypeInfo returns info about TL type.
func (r *ReportChatPhotoRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "reportChatPhoto",
		ID:   ReportChatPhotoRequestTypeID,
	}
	if r == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ChatID",
			SchemaName: "chat_id",
		},
		{
			Name:       "FileID",
			SchemaName: "file_id",
		},
		{
			Name:       "Reason",
			SchemaName: "reason",
		},
		{
			Name:       "Text",
			SchemaName: "text",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (r *ReportChatPhotoRequest) Encode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode reportChatPhoto#2bc9e924 as nil")
	}
	b.PutID(ReportChatPhotoRequestTypeID)
	return r.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (r *ReportChatPhotoRequest) EncodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode reportChatPhoto#2bc9e924 as nil")
	}
	b.PutLong(r.ChatID)
	b.PutInt32(r.FileID)
	if r.Reason == nil {
		return fmt.Errorf("unable to encode reportChatPhoto#2bc9e924: field reason is nil")
	}
	if err := r.Reason.Encode(b); err != nil {
		return fmt.Errorf("unable to encode reportChatPhoto#2bc9e924: field reason: %w", err)
	}
	b.PutString(r.Text)
	return nil
}

// Decode implements bin.Decoder.
func (r *ReportChatPhotoRequest) Decode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode reportChatPhoto#2bc9e924 to nil")
	}
	if err := b.ConsumeID(ReportChatPhotoRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode reportChatPhoto#2bc9e924: %w", err)
	}
	return r.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (r *ReportChatPhotoRequest) DecodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode reportChatPhoto#2bc9e924 to nil")
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode reportChatPhoto#2bc9e924: field chat_id: %w", err)
		}
		r.ChatID = value
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode reportChatPhoto#2bc9e924: field file_id: %w", err)
		}
		r.FileID = value
	}
	{
		value, err := DecodeChatReportReason(b)
		if err != nil {
			return fmt.Errorf("unable to decode reportChatPhoto#2bc9e924: field reason: %w", err)
		}
		r.Reason = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode reportChatPhoto#2bc9e924: field text: %w", err)
		}
		r.Text = value
	}
	return nil
}

// EncodeTDLibJSON implements jsontd.TDLibEncoder.
func (r *ReportChatPhotoRequest) EncodeTDLibJSON(b jsontd.Encoder) error {
	if r == nil {
		return fmt.Errorf("can't encode reportChatPhoto#2bc9e924 as nil")
	}
	b.ObjStart()
	b.PutID("reportChatPhoto")
	b.FieldStart("chat_id")
	b.PutLong(r.ChatID)
	b.FieldStart("file_id")
	b.PutInt32(r.FileID)
	b.FieldStart("reason")
	if r.Reason == nil {
		return fmt.Errorf("unable to encode reportChatPhoto#2bc9e924: field reason is nil")
	}
	if err := r.Reason.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode reportChatPhoto#2bc9e924: field reason: %w", err)
	}
	b.FieldStart("text")
	b.PutString(r.Text)
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements jsontd.TDLibDecoder.
func (r *ReportChatPhotoRequest) DecodeTDLibJSON(b jsontd.Decoder) error {
	if r == nil {
		return fmt.Errorf("can't decode reportChatPhoto#2bc9e924 to nil")
	}

	return b.Obj(func(b jsontd.Decoder, key []byte) error {
		switch string(key) {
		case jsontd.TypeField:
			if err := b.ConsumeID("reportChatPhoto"); err != nil {
				return fmt.Errorf("unable to decode reportChatPhoto#2bc9e924: %w", err)
			}
		case "chat_id":
			value, err := b.Long()
			if err != nil {
				return fmt.Errorf("unable to decode reportChatPhoto#2bc9e924: field chat_id: %w", err)
			}
			r.ChatID = value
		case "file_id":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode reportChatPhoto#2bc9e924: field file_id: %w", err)
			}
			r.FileID = value
		case "reason":
			value, err := DecodeTDLibJSONChatReportReason(b)
			if err != nil {
				return fmt.Errorf("unable to decode reportChatPhoto#2bc9e924: field reason: %w", err)
			}
			r.Reason = value
		case "text":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode reportChatPhoto#2bc9e924: field text: %w", err)
			}
			r.Text = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetChatID returns value of ChatID field.
func (r *ReportChatPhotoRequest) GetChatID() (value int64) {
	return r.ChatID
}

// GetFileID returns value of FileID field.
func (r *ReportChatPhotoRequest) GetFileID() (value int32) {
	return r.FileID
}

// GetReason returns value of Reason field.
func (r *ReportChatPhotoRequest) GetReason() (value ChatReportReasonClass) {
	return r.Reason
}

// GetText returns value of Text field.
func (r *ReportChatPhotoRequest) GetText() (value string) {
	return r.Text
}

// ReportChatPhoto invokes method reportChatPhoto#2bc9e924 returning error if any.
func (c *Client) ReportChatPhoto(ctx context.Context, request *ReportChatPhotoRequest) error {
	var ok Ok

	if err := c.rpc.Invoke(ctx, request, &ok); err != nil {
		return err
	}
	return nil
}
