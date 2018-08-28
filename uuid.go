package uuid

import (
	"strings"

	exuuid "github.com/gofrs/uuid"
)

type UUID interface {
	String() string
	NoHyphen() string
}

// private entity struct
type uuid struct {
	exuuid.UUID
}

func (u *uuid) NoHyphen() string {
	return strings.Replace(u.String(), "-", "", -1)
}

func NewUUID() (UUID, error) {
	if uuidv4, err := exuuid.NewV4(); err != nil {
		return nil, err
	} else {
		return &uuid{
			UUID: uuidv4,
		}, nil
	}
}

func FactoryUUIDFromString(value string) (UUID, error) {
	if genUuid, err := exuuid.FromString(value); err != nil {
		return nil, err
	} else {
		return &uuid{
			UUID: genUuid,
		}, nil
	}
}
