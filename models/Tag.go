package models

type Tag struct {
	Uid string `json:"uid,omitempty"`

	Xid  string `json:"xid,omitempty"`
	Type string `json:"dgraph.type,omitempty"`

	Name        string `json:"name,omitempty" validate:"alphanumunicode"`
	Description string `json:"description,omitempty" validate:"alphanumunicode"`
	Citation    uint64 `json:"citation,omitempty" validate:"min=0"`
	IsKey       bool   `json:"isKey,omitempty"`
	Status      int    `json:"status,omitempty"`
	Tag         []Tag  `json:"tag,omitempty"`
}

const TagXid = "Tag"
