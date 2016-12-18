package compiler

import "gitlab.com/alehander42/melt/compiler/types"

type Nil struct {
	Value string

	Info
}

func (n *Nil) TypeCheck(ctx *Context) error {
	n.meltType = types.Empty{}
	return nil
}
