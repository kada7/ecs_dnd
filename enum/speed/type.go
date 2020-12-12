package speed

import "errors"

type Type string

const (
	Walking Type = "walking"
	Flying  Type = "flying"
	Swing   Type = "swing"
)

func (t Type) MustValid() error {
	for _, typ := range AllType() {
		if t == typ {
			return nil
		}
	}
	return errors.New("speed type is invalid")
}

func (t Type) String() string {
	return string(t)
}

func AllType() []Type {
	return []Type{
		Walking,
		Flying,
		Swing,
	}
}
