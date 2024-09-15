package time

import (
	"encoding"
	xtime "time"
)

var _ encoding.TextUnmarshaler = (*Duration)(nil)
var _ encoding.TextMarshaler = (*Duration)(nil)

// Duration be used toml unmarshal string time, like 1s, 500ms.
type Duration xtime.Duration

func (d *Duration) MarshalText() (text []byte, err error) {
	dd := xtime.Duration(*d)
	return []byte(dd.String()), nil
}

// UnmarshalText unmarshal text to duration.
func (d *Duration) UnmarshalText(text []byte) error {
	tmp, err := xtime.ParseDuration(string(text))
	if err == nil {
		*d = Duration(tmp)
	}
	return err
}

func (d *Duration) AsDuration() xtime.Duration {
	return xtime.Duration(*d)
}
