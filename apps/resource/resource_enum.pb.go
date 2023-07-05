package resource

import (
	"bytes"
	"fmt"
	"strings"
)

// ParseVENDORFromString Parse VENDOR from string
func ParseVENDORFromString(str string) (VENDOR, error) {
	key := strings.Trim(string(str), `"`)
	v, ok := VENDOR_value[strings.ToUpper(key)]
	if !ok {
		return 0, fmt.Errorf("unknown VENDOR: %s", str)
	}

	return VENDOR(v), nil
}

// Equal type compare
func (t VENDOR) Equal(target VENDOR) bool {
	return t == target
}

// IsIn todo
func (t VENDOR) IsIn(targets ...VENDOR) bool {
	for _, target := range targets {
		if t.Equal(target) {
			return true
		}
	}

	return false
}

// MarshalJSON todo
func (t VENDOR) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(strings.ToUpper(t.String()))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

// UnmarshalJSON todo
func (t *VENDOR) UnmarshalJSON(b []byte) error {
	ins, err := ParseVENDORFromString(string(b))
	if err != nil {
		return err
	}
	*t = ins
	return nil
}

// ParseTYPEFromString Parse TYPE from string
func ParseTYPEFromString(str string) (TYPE, error) {
	key := strings.Trim(string(str), `"`)
	v, ok := TYPE_value[strings.ToUpper(key)]
	if !ok {
		return 0, fmt.Errorf("unknown TYPE: %s", str)
	}

	return TYPE(v), nil
}

// Equal type compare
func (t TYPE) Equal(target TYPE) bool {
	return t == target
}

// IsIn todo
func (t TYPE) IsIn(targets ...TYPE) bool {
	for _, target := range targets {
		if t.Equal(target) {
			return true
		}
	}

	return false
}

// MarshalJSON todo
func (t TYPE) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(strings.ToUpper(t.String()))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

// UnmarshalJSON todo
func (t *TYPE) UnmarshalJSON(b []byte) error {
	ins, err := ParseTYPEFromString(string(b))
	if err != nil {
		return err
	}
	*t = ins
	return nil
}

// ParsePAY_MODEFromString Parse PAY_MODE from string
func ParsePAY_MODEFromString(str string) (PAY_MODE, error) {
	key := strings.Trim(string(str), `"`)
	v, ok := PAY_MODE_value[strings.ToUpper(key)]
	if !ok {
		return 0, fmt.Errorf("unknown PAY_MODE: %s", str)
	}

	return PAY_MODE(v), nil
}

// Equal type compare
func (t PAY_MODE) Equal(target PAY_MODE) bool {
	return t == target
}

// IsIn todo
func (t PAY_MODE) IsIn(targets ...PAY_MODE) bool {
	for _, target := range targets {
		if t.Equal(target) {
			return true
		}
	}

	return false
}

// MarshalJSON todo
func (t PAY_MODE) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(strings.ToUpper(t.String()))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

// UnmarshalJSON todo
func (t *PAY_MODE) UnmarshalJSON(b []byte) error {
	ins, err := ParsePAY_MODEFromString(string(b))
	if err != nil {
		return err
	}
	*t = ins
	return nil
}

// ParseUSAGE_MODEFromString Parse USAGE_MODE from string
func ParseUSAGE_MODEFromString(str string) (USAGE_MODE, error) {
	key := strings.Trim(string(str), `"`)
	v, ok := USAGE_MODE_value[strings.ToUpper(key)]
	if !ok {
		return 0, fmt.Errorf("unknown USAGE_MODE: %s", str)
	}

	return USAGE_MODE(v), nil
}

// Equal type compare
func (t USAGE_MODE) Equal(target USAGE_MODE) bool {
	return t == target
}

// IsIn todo
func (t USAGE_MODE) IsIn(targets ...USAGE_MODE) bool {
	for _, target := range targets {
		if t.Equal(target) {
			return true
		}
	}

	return false
}

// MarshalJSON todo
func (t USAGE_MODE) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(strings.ToUpper(t.String()))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

// UnmarshalJSON todo
func (t *USAGE_MODE) UnmarshalJSON(b []byte) error {
	ins, err := ParseUSAGE_MODEFromString(string(b))
	if err != nil {
		return err
	}
	*t = ins
	return nil
}

// ParseTAG_PURPOSEFromString Parse TAG_PURPOSE from string
func ParseTAG_PURPOSEFromString(str string) (TAG_PURPOSE, error) {
	key := strings.Trim(string(str), `"`)
	v, ok := TAG_PURPOSE_value[strings.ToUpper(key)]
	if !ok {
		return 0, fmt.Errorf("unknown TAG_PURPOSE: %s", str)
	}

	return TAG_PURPOSE(v), nil
}

// Equal type compare
func (t TAG_PURPOSE) Equal(target TAG_PURPOSE) bool {
	return t == target
}

// IsIn todo
func (t TAG_PURPOSE) IsIn(targets ...TAG_PURPOSE) bool {
	for _, target := range targets {
		if t.Equal(target) {
			return true
		}
	}

	return false
}

// MarshalJSON todo
func (t TAG_PURPOSE) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(strings.ToUpper(t.String()))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

// UnmarshalJSON todo
func (t *TAG_PURPOSE) UnmarshalJSON(b []byte) error {
	ins, err := ParseTAG_PURPOSEFromString(string(b))
	if err != nil {
		return err
	}
	*t = ins
	return nil
}

// ParseUpdateActionFromString Parse UpdateAction from string
func ParseUpdateActionFromString(str string) (UpdateAction, error) {
	key := strings.Trim(string(str), `"`)
	v, ok := UpdateAction_value[strings.ToUpper(key)]
	if !ok {
		return 0, fmt.Errorf("unknown UpdateAction: %s", str)
	}

	return UpdateAction(v), nil
}

// Equal type compare
func (t UpdateAction) Equal(target UpdateAction) bool {
	return t == target
}

// IsIn todo
func (t UpdateAction) IsIn(targets ...UpdateAction) bool {
	for _, target := range targets {
		if t.Equal(target) {
			return true
		}
	}

	return false
}

// MarshalJSON todo
func (t UpdateAction) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(strings.ToUpper(t.String()))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

// UnmarshalJSON todo
func (t *UpdateAction) UnmarshalJSON(b []byte) error {
	ins, err := ParseUpdateActionFromString(string(b))
	if err != nil {
		return err
	}
	*t = ins
	return nil
}
