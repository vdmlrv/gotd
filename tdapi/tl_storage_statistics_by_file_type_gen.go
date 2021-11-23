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

// StorageStatisticsByFileType represents TL type `storageStatisticsByFileType#2a8ef8a8`.
type StorageStatisticsByFileType struct {
	// File type
	FileType FileTypeClass
	// Total size of the files
	Size int64
	// Total number of files
	Count int32
}

// StorageStatisticsByFileTypeTypeID is TL type id of StorageStatisticsByFileType.
const StorageStatisticsByFileTypeTypeID = 0x2a8ef8a8

// Ensuring interfaces in compile-time for StorageStatisticsByFileType.
var (
	_ bin.Encoder     = &StorageStatisticsByFileType{}
	_ bin.Decoder     = &StorageStatisticsByFileType{}
	_ bin.BareEncoder = &StorageStatisticsByFileType{}
	_ bin.BareDecoder = &StorageStatisticsByFileType{}
)

func (s *StorageStatisticsByFileType) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.FileType == nil) {
		return false
	}
	if !(s.Size == 0) {
		return false
	}
	if !(s.Count == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *StorageStatisticsByFileType) String() string {
	if s == nil {
		return "StorageStatisticsByFileType(nil)"
	}
	type Alias StorageStatisticsByFileType
	return fmt.Sprintf("StorageStatisticsByFileType%+v", Alias(*s))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*StorageStatisticsByFileType) TypeID() uint32 {
	return StorageStatisticsByFileTypeTypeID
}

// TypeName returns name of type in TL schema.
func (*StorageStatisticsByFileType) TypeName() string {
	return "storageStatisticsByFileType"
}

// TypeInfo returns info about TL type.
func (s *StorageStatisticsByFileType) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "storageStatisticsByFileType",
		ID:   StorageStatisticsByFileTypeTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "FileType",
			SchemaName: "file_type",
		},
		{
			Name:       "Size",
			SchemaName: "size",
		},
		{
			Name:       "Count",
			SchemaName: "count",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (s *StorageStatisticsByFileType) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode storageStatisticsByFileType#2a8ef8a8 as nil")
	}
	b.PutID(StorageStatisticsByFileTypeTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *StorageStatisticsByFileType) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode storageStatisticsByFileType#2a8ef8a8 as nil")
	}
	if s.FileType == nil {
		return fmt.Errorf("unable to encode storageStatisticsByFileType#2a8ef8a8: field file_type is nil")
	}
	if err := s.FileType.Encode(b); err != nil {
		return fmt.Errorf("unable to encode storageStatisticsByFileType#2a8ef8a8: field file_type: %w", err)
	}
	b.PutLong(s.Size)
	b.PutInt32(s.Count)
	return nil
}

// Decode implements bin.Decoder.
func (s *StorageStatisticsByFileType) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode storageStatisticsByFileType#2a8ef8a8 to nil")
	}
	if err := b.ConsumeID(StorageStatisticsByFileTypeTypeID); err != nil {
		return fmt.Errorf("unable to decode storageStatisticsByFileType#2a8ef8a8: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *StorageStatisticsByFileType) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode storageStatisticsByFileType#2a8ef8a8 to nil")
	}
	{
		value, err := DecodeFileType(b)
		if err != nil {
			return fmt.Errorf("unable to decode storageStatisticsByFileType#2a8ef8a8: field file_type: %w", err)
		}
		s.FileType = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode storageStatisticsByFileType#2a8ef8a8: field size: %w", err)
		}
		s.Size = value
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode storageStatisticsByFileType#2a8ef8a8: field count: %w", err)
		}
		s.Count = value
	}
	return nil
}

// EncodeTDLibJSON implements jsontd.TDLibEncoder.
func (s *StorageStatisticsByFileType) EncodeTDLibJSON(b jsontd.Encoder) error {
	if s == nil {
		return fmt.Errorf("can't encode storageStatisticsByFileType#2a8ef8a8 as nil")
	}
	b.ObjStart()
	b.PutID("storageStatisticsByFileType")
	b.FieldStart("file_type")
	if s.FileType == nil {
		return fmt.Errorf("unable to encode storageStatisticsByFileType#2a8ef8a8: field file_type is nil")
	}
	if err := s.FileType.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode storageStatisticsByFileType#2a8ef8a8: field file_type: %w", err)
	}
	b.FieldStart("size")
	b.PutLong(s.Size)
	b.FieldStart("count")
	b.PutInt32(s.Count)
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements jsontd.TDLibDecoder.
func (s *StorageStatisticsByFileType) DecodeTDLibJSON(b jsontd.Decoder) error {
	if s == nil {
		return fmt.Errorf("can't decode storageStatisticsByFileType#2a8ef8a8 to nil")
	}

	return b.Obj(func(b jsontd.Decoder, key []byte) error {
		switch string(key) {
		case jsontd.TypeField:
			if err := b.ConsumeID("storageStatisticsByFileType"); err != nil {
				return fmt.Errorf("unable to decode storageStatisticsByFileType#2a8ef8a8: %w", err)
			}
		case "file_type":
			value, err := DecodeTDLibJSONFileType(b)
			if err != nil {
				return fmt.Errorf("unable to decode storageStatisticsByFileType#2a8ef8a8: field file_type: %w", err)
			}
			s.FileType = value
		case "size":
			value, err := b.Long()
			if err != nil {
				return fmt.Errorf("unable to decode storageStatisticsByFileType#2a8ef8a8: field size: %w", err)
			}
			s.Size = value
		case "count":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode storageStatisticsByFileType#2a8ef8a8: field count: %w", err)
			}
			s.Count = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetFileType returns value of FileType field.
func (s *StorageStatisticsByFileType) GetFileType() (value FileTypeClass) {
	return s.FileType
}

// GetSize returns value of Size field.
func (s *StorageStatisticsByFileType) GetSize() (value int64) {
	return s.Size
}

// GetCount returns value of Count field.
func (s *StorageStatisticsByFileType) GetCount() (value int32) {
	return s.Count
}
