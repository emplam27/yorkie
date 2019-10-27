package operation

import (
	"fmt"

	"github.com/hackerwins/rottie/pkg/document/json/datatype"

	"github.com/hackerwins/rottie/pkg/document/json"
	"github.com/hackerwins/rottie/pkg/document/time"
	"github.com/hackerwins/rottie/pkg/log"
)

type Set struct {
	key             string
	value           datatype.Element
	parentCreatedAt *time.Ticket
	executedAt      *time.Ticket
}

func NewSet(
	key string,
	value datatype.Element,
	parentCreatedAt *time.Ticket,
	executedAt *time.Ticket,
) *Set {
	return &Set{
		key:             key,
		value:           value,
		parentCreatedAt: parentCreatedAt,
		executedAt:      executedAt,
	}
}

func (o *Set) Execute(root *json.Root) error {
	parent := root.FindByCreatedAt(o.parentCreatedAt)

	obj, ok := parent.(*json.Object)
	if !ok {
		err := fmt.Errorf("fail to execute, only Object can execute Set")
		log.Logger.Error(err)
		return err
	}

	obj.Set(o.key, o.value)
	root.RegisterElement(o.value)
	return nil
}

func (o *Set) ExecutedAt() *time.Ticket {
	return o.executedAt
}
