package schema

import (
	"github.com/zlsgo/app_module/model/hook"
)

type Options struct {
	DisabledMigrator *bool `json:"disabled_migrator,omitempty"`
	SoftDeletes      *bool `json:"soft_deletes,omitempty"`
	SoftDeleteIsTime *bool `json:"soft_delete_is_time,omitempty"`
	Timestamps       *bool `json:"timestamps,omitempty"`
	CryptID          *bool `json:"crypt_id,omitempty"`
	Hook             func(event hook.Event, data ...any) error
	Salt             string   `json:"crypt_salt,omitempty"`
	LowFields        []string `json:"low_fields,omitempty"`
	FieldsSort       []string `json:"fields_sort,omitempty"`
	CryptLen         int      `json:"crypt_len,omitempty"`
}

func (o *Options) SetDisabledMigrator(b bool) *Options {
	o.DisabledMigrator = &b
	return o
}

func (o *Options) SetSalt(s string) *Options {
	o.Salt = s
	return o
}

func (o *Options) SetSoftDeletes(b bool) *Options {
	o.SoftDeletes = &b
	return o
}

func (o *Options) SetSoftDeleteIsTime(b bool) *Options {
	o.SoftDeleteIsTime = &b
	return o
}

func (o *Options) SetTimestamps(b bool) *Options {
	o.Timestamps = &b
	return o
}

func (o *Options) SetCryptID(b bool) *Options {
	o.CryptID = &b
	return o
}

func (o *Options) SetCryptLen(i int) *Options {
	o.CryptLen = i
	return o
}

func (o *Options) SetHook(h func(event hook.Event, data ...any) error) *Options {
	o.Hook = h
	return o
}
