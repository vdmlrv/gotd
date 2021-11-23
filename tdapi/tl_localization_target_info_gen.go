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

// LocalizationTargetInfo represents TL type `localizationTargetInfo#2ca3903b`.
type LocalizationTargetInfo struct {
	// List of available language packs for this application
	LanguagePacks []LanguagePackInfo
}

// LocalizationTargetInfoTypeID is TL type id of LocalizationTargetInfo.
const LocalizationTargetInfoTypeID = 0x2ca3903b

// Ensuring interfaces in compile-time for LocalizationTargetInfo.
var (
	_ bin.Encoder     = &LocalizationTargetInfo{}
	_ bin.Decoder     = &LocalizationTargetInfo{}
	_ bin.BareEncoder = &LocalizationTargetInfo{}
	_ bin.BareDecoder = &LocalizationTargetInfo{}
)

func (l *LocalizationTargetInfo) Zero() bool {
	if l == nil {
		return true
	}
	if !(l.LanguagePacks == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (l *LocalizationTargetInfo) String() string {
	if l == nil {
		return "LocalizationTargetInfo(nil)"
	}
	type Alias LocalizationTargetInfo
	return fmt.Sprintf("LocalizationTargetInfo%+v", Alias(*l))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*LocalizationTargetInfo) TypeID() uint32 {
	return LocalizationTargetInfoTypeID
}

// TypeName returns name of type in TL schema.
func (*LocalizationTargetInfo) TypeName() string {
	return "localizationTargetInfo"
}

// TypeInfo returns info about TL type.
func (l *LocalizationTargetInfo) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "localizationTargetInfo",
		ID:   LocalizationTargetInfoTypeID,
	}
	if l == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "LanguagePacks",
			SchemaName: "language_packs",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (l *LocalizationTargetInfo) Encode(b *bin.Buffer) error {
	if l == nil {
		return fmt.Errorf("can't encode localizationTargetInfo#2ca3903b as nil")
	}
	b.PutID(LocalizationTargetInfoTypeID)
	return l.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (l *LocalizationTargetInfo) EncodeBare(b *bin.Buffer) error {
	if l == nil {
		return fmt.Errorf("can't encode localizationTargetInfo#2ca3903b as nil")
	}
	b.PutInt(len(l.LanguagePacks))
	for idx, v := range l.LanguagePacks {
		if err := v.EncodeBare(b); err != nil {
			return fmt.Errorf("unable to encode bare localizationTargetInfo#2ca3903b: field language_packs element with index %d: %w", idx, err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (l *LocalizationTargetInfo) Decode(b *bin.Buffer) error {
	if l == nil {
		return fmt.Errorf("can't decode localizationTargetInfo#2ca3903b to nil")
	}
	if err := b.ConsumeID(LocalizationTargetInfoTypeID); err != nil {
		return fmt.Errorf("unable to decode localizationTargetInfo#2ca3903b: %w", err)
	}
	return l.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (l *LocalizationTargetInfo) DecodeBare(b *bin.Buffer) error {
	if l == nil {
		return fmt.Errorf("can't decode localizationTargetInfo#2ca3903b to nil")
	}
	{
		headerLen, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode localizationTargetInfo#2ca3903b: field language_packs: %w", err)
		}

		if headerLen > 0 {
			l.LanguagePacks = make([]LanguagePackInfo, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			var value LanguagePackInfo
			if err := value.DecodeBare(b); err != nil {
				return fmt.Errorf("unable to decode bare localizationTargetInfo#2ca3903b: field language_packs: %w", err)
			}
			l.LanguagePacks = append(l.LanguagePacks, value)
		}
	}
	return nil
}

// EncodeTDLibJSON implements jsontd.TDLibEncoder.
func (l *LocalizationTargetInfo) EncodeTDLibJSON(b jsontd.Encoder) error {
	if l == nil {
		return fmt.Errorf("can't encode localizationTargetInfo#2ca3903b as nil")
	}
	b.ObjStart()
	b.PutID("localizationTargetInfo")
	b.FieldStart("language_packs")
	b.ArrStart()
	for idx, v := range l.LanguagePacks {
		if err := v.EncodeTDLibJSON(b); err != nil {
			return fmt.Errorf("unable to encode localizationTargetInfo#2ca3903b: field language_packs element with index %d: %w", idx, err)
		}
	}
	b.ArrEnd()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements jsontd.TDLibDecoder.
func (l *LocalizationTargetInfo) DecodeTDLibJSON(b jsontd.Decoder) error {
	if l == nil {
		return fmt.Errorf("can't decode localizationTargetInfo#2ca3903b to nil")
	}

	return b.Obj(func(b jsontd.Decoder, key []byte) error {
		switch string(key) {
		case jsontd.TypeField:
			if err := b.ConsumeID("localizationTargetInfo"); err != nil {
				return fmt.Errorf("unable to decode localizationTargetInfo#2ca3903b: %w", err)
			}
		case "language_packs":
			if err := b.Arr(func(b jsontd.Decoder) error {
				var value LanguagePackInfo
				if err := value.DecodeTDLibJSON(b); err != nil {
					return fmt.Errorf("unable to decode localizationTargetInfo#2ca3903b: field language_packs: %w", err)
				}
				l.LanguagePacks = append(l.LanguagePacks, value)
				return nil
			}); err != nil {
				return fmt.Errorf("unable to decode localizationTargetInfo#2ca3903b: field language_packs: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetLanguagePacks returns value of LanguagePacks field.
func (l *LocalizationTargetInfo) GetLanguagePacks() (value []LanguagePackInfo) {
	return l.LanguagePacks
}
