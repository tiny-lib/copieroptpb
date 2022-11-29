package wrapCopierOpt

import (
	"github.com/czyt/copieroptpb/internal"
	"github.com/jinzhu/copier"
)

var wrapperOption = copier.Option{Converters: internal.PbWrapperValueConverters()}

func Option() copier.Option {
	return wrapperOption
}
