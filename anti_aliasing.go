package gofra

import (
	"encoding/json"
	"fmt"
)

// Anti-aliasing

// https://en.wikipedia.org/wiki/Spatial_anti-aliasing
// https://en.wikipedia.org/wiki/Multisample_anti-aliasing
// https://en.wikipedia.org/wiki/Supersampling

/*

+---+
|   |
+---+

4x - 2*2
+---+---+
|   |   |
+---+---+
|   |   |
+---+---+

9x - 3*3
+---+---+---+
|   |   |   |
+---+---+---+
|   |   |   |
+---+---+---+
|   |   |   |
+---+---+---+

16x - 4*4
+---+---+---+---+
|   |   |   |   |
+---+---+---+---+
|   |   |   |   |
+---+---+---+---+
|   |   |   |   |
+---+---+---+---+
|   |   |   |   |
+---+---+---+---+

25x - 5*5
+---+---+---+---+---+
|   |   |   |   |   |
+---+---+---+---+---+
|   |   |   |   |   |
+---+---+---+---+---+
|   |   |   |   |   |
+---+---+---+---+---+
|   |   |   |   |   |
+---+---+---+---+---+
|   |   |   |   |   |
+---+---+---+---+---+

*/

type AntiAliasing int

const (
	AA_NONE AntiAliasing = iota
	AA_4X                // 4 samples per pixel
	AA_9X                // 9 samples per pixel
	AA_16X               // 16 samples per pixel
	AA_25X               // 25 samples per pixel
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

func (aa AntiAliasing) String() string {
	name, ok := aaNames[aa]
	if ok {
		return name
	}
	return fmt.Sprintf("AntiAliasing(%d)", aa)
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

func ParseAntiAliasing(s string) (AntiAliasing, error) {
	value, ok := aaValues[s]
	if !ok {
		return 0, fmt.Errorf("gofra.ParseAntiAliasing: undefined name %q", s)
	}
	return value, nil
}
