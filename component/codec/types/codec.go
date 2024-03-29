package types

type StreamPacker interface {
	Pack(data []byte) ([]byte, error)
}

type StreamUnPacker interface {
	UnPack([]byte) ([]byte, error)
}

type Codec interface {
	Marshal(data interface{}) ([]byte, error)
	Unmarshal(data []byte, ret interface{}) error
}

type Unmarshaler interface {
	Unmarshal(data []byte) error
}
type Marshaller interface {
	Marshal() ([]byte, error)
}
