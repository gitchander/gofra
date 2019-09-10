package gofra

import (
	"encoding/json"
	"fmt"
)

type AntiAliasing int

const (
	AA_NONE AntiAliasing = iota
	AA_4X
	AA_9X
	AA_16X
	AA_25X
)

const (
	name_AA_NONE = "none"
	name_AA_4X   = "4x"
	name_AA_9X   = "9x"
	name_AA_16X  = "16x"
	name_AA_25X  = "25x"
)

var aaNames = map[AntiAliasing]string{
	AA_NONE: name_AA_NONE,
	AA_4X:   name_AA_4X,
	AA_9X:   name_AA_9X,
	AA_16X:  name_AA_16X,
	AA_25X:  name_AA_25X,
}

var aaValues = map[string]AntiAliasing{
	name_AA_NONE: AA_NONE,
	name_AA_4X:   AA_4X,
	name_AA_9X:   AA_9X,
	name_AA_16X:  AA_16X,
	name_AA_25X:  AA_25X,
}

func (aa AntiAliasing) MarshalJSON() ([]byte, error) {
	value := aa
	s, ok := aaNames[value]
	if !ok {
		return nil, fmt.Errorf("gofra.AntiAliasing.MarshalJSON: undefined value %d", value)
	}
	return json.Marshal(s)
}

func (aa *AntiAliasing) UnmarshalJSON(data []byte) error {
	var name string
	err := json.Unmarshal(data, &name)
	if err != nil {
		return err
	}
	value, ok := aaValues[name]
	if !ok {
		return fmt.Errorf("gofra.AntiAliasing.UnmarshalJSON: undefined name %q", name)
	}
	*aa = value
	return nil
}
