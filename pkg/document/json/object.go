package json

import (
	"github.com/hackerwins/rottie/pkg/document/json/datatype"
	"github.com/hackerwins/rottie/pkg/document/time"
)

type Object struct {
	members   *datatype.RHT
	createdAt *time.Ticket
}

func NewObject(members *datatype.RHT, createdAt *time.Ticket) *Object {
	return &Object{
		members:   members,
		createdAt: createdAt,
	}
}

func (o *Object) Set(k string, v datatype.Element) {
	o.members.Set(k, v)
}

func (o *Object) Members() map[string]datatype.Element {
	return o.members.Members()
}

func (o *Object) Marshal() string {
	return o.members.Marshal()
}

func (o *Object) CreatedAt() *time.Ticket {
	return o.createdAt
}
