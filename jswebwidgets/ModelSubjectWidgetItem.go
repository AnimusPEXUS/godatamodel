package jswebwidgets

import (
	"syscall/js"

	"github.com/AnimusPEXUS/godatamodel"
	"github.com/AnimusPEXUS/gojstools/elementtreeconstructor"
	gojstoolsutils "github.com/AnimusPEXUS/gojstools/utils"
)

type ModelSubjectWidgetItemOptions struct {
	Etc          *elementtreeconstructor.ElementTreeConstructor
	Parent       *ModelSubjectWidget
	Item         *godatamodel.DataModelSubjectField
	InitialValue interface{}
	InitialMode  ModelSubjectWidgetMode
}

type ModelSubjectWidgetItem struct {
	options *ModelSubjectWidgetItemOptions

	EditingElement *elementtreeconstructor.ElementMutator

	Element *elementtreeconstructor.ElementMutator
}

func NewModelSubjectWidgetItem(
	options *ModelSubjectWidgetItemOptions,
) (*ModelSubjectWidgetItem, error) {

	self := &ModelSubjectWidgetItem{options: options}

	etc := self.options.Etc

	editing_child := etc.CreateElement("input").SetAttribute("type", "text")
	self.EditingElement = editing_child

	self.Element = etc.CreateElement("div").AppendChildren(
		etc.CreateElement("label").AppendChildren(etc.CreateTextNode(self.options.Item.Name)),
		editing_child,
	)

	if self.options.InitialValue != nil {
		editing_child.SetJsValue(
			"value",
			gojstoolsutils.JSValueLiteralToPointer(
				js.ValueOf(self.options.InitialValue),
			),
		)
	}
	return self, nil
}

func (self *ModelSubjectWidgetItem) GetValue() interface{} {
	ret := self.EditingElement.GetJsValue("value").String()
	return ret
}
