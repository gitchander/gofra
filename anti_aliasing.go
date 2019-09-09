package gofra

import (
	"encoding/json"
	"errors"
	"strings"
)

type AntiAliasing int

const (
	AA_NONE AntiAliasing = iota
	AA_4X
	AA_9X
	AA_16X
	AA_25X
)

var key_AntiAliasing = map[AntiAliasing]string{
	AA_NONE: "NONE",
	AA_4X:   "4X",
	AA_9X:   "9X",
	AA_16X:  "16X",
	AA_25X:  "25X",
}

var val_AntiAliasing = map[string]AntiAliasing{
	"NONE": AA_NONE,
	"4X":   AA_4X,
	"9X":   AA_9X,
	"16X":  AA_16X,
	"25X":  AA_25X,
}

func (aa AntiAliasing) MarshalJSON() ([]byte, error) {

	s, ok := key_AntiAliasing[aa]
	if !ok {
		return nil, errors.New("AntiAliasing.MarshalJSON")
	}

	return json.Marshal(s)
}

func (aa *AntiAliasing) UnmarshalJSON(data []byte) error {

	var s string

	err := json.Unmarshal(data, &s)
	if err != nil {
		return err
	}

	s = strings.ToUpper(s)
	v, ok := val_AntiAliasing[s]
	if !ok {
		return errors.New("AntiAliasing.UnmarshalJSON")
	}

	*aa = v

	return nil
}
