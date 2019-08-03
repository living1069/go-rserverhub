package util

import (
    "fmt"
    "time"
)

type JSONTime time.Time

func (t JSONTime)MarshalJSON() ([]byte, error) {
    stamp := fmt.Sprintf("%d", time.Time(t).Unix())
    return []byte(stamp), nil
}
