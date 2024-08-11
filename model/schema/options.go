package schema

type Options struct {
	Salt             string   `json:"crypt_salt,omitempty"`
	LowFields        []string `json:"low_fields,omitempty"`
	FieldsSort       []string `json:"fields_sort,omitempty"`
	CryptLen         int      `json:"crypt_len,omitempty"`
	DisabledMigrator *bool    `json:"disabled_migrator,omitempty"`
	// SkipFieldValidation bool     `json:"skip_field_validation,omitempty"`
	SoftDeletes *bool `json:"soft_deletes,omitempty"`
	Timestamps  *bool `json:"timestamps,omitempty"`
	CryptID     *bool `json:"crypt_id,omitempty"`
}

func (o *Options) SetDisabledMigrator(b bool) {
	o.DisabledMigrator = &b
}

func (o *Options) SetSalt(s string) {
	o.Salt = s
}

func (o *Options) SetSoftDeletes(b bool) {
	o.SoftDeletes = &b
}

func (o *Options) SetTimestamps(b bool) {
	o.Timestamps = &b
}

func (o *Options) SetCryptID(b bool) {
	o.CryptID = &b
}
