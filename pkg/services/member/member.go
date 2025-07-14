package member

import (
	"context"
	"fmt"
)

type MemberSvcs interface {
	AddMember(ctx context.Context, name string) (string, error)
}

type memberSvcs struct{}

func NewMember() MemberSvcs {
	return &memberSvcs{}
}

func (m *memberSvcs) AddMember(ctx context.Context, name string) (string, error) {

	return fmt.Sprintf("member %s add success", name), nil
}
