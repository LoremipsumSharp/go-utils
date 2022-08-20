package json

import "encoding/json"

var (
	Marshal   = json.Marshal
	Unmarshal = json.Unmarshal
)

func UnmarshalTo[T any](bs []byte) (m T, err error) {
	return m, Unmarshal(bs, &m)
}

func MustMarshal(v any) []byte {
	bs, err := Marshal(v)
	if err != nil {
		panic(err)
	}
	return bs
}

func MustMarshalString(v any) string {
	return string(MustMarshal(v))
}

func MustUnmarshal(bs []byte, v any) any {
	err := Unmarshal(bs, v)
	if err != nil {
		panic(err)
	}
	return v
}

func MustUnmarshalTo[T any](bs []byte) (m T) {
	err := Unmarshal(bs, &m)
	if err != nil {
		panic(err)
	}
	return m
}