package board

import (
	"context"
	"fmt"

	"github.com/samber/do"
)

type GroupSvcs interface {
	Test(ctx context.Context)
}

type groupsvcs struct {
}

func NewGroupSvcs(i *do.Injector) (GroupSvcs, error) {
	return &groupsvcs{}, nil
}

func (g *groupsvcs) Test(ctx context.Context) {
	fmt.Println("test from group service")
}
