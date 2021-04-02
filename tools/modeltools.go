package tools

import (
	"github.com/AnimusPEXUS/godatamodel"
)

func GenerateModelForExport(in *godatamodel.DataModel) *godatamodel.DataModel {
	ret := *in

	ret.Subjects = make([]*godatamodel.DataModelSubject, 0)
	for _, i := range in.Subjects {
		if i.ForExport {
			ret.Subjects = append(ret.Subjects, i)
		}
	}

	return &ret
}
