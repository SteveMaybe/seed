package script

import "qlova.org/seed/js"

type AnyFile interface {
	AnyValue
	GetFile() js.Value
}
