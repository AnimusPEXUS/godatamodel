package jswebwidgets

import (
	"github.com/AnimusPEXUS/godatamodel"
	"github.com/AnimusPEXUS/gojstools/elementtreeconstructor"
)

type ModelSubjectWidgetMode uint

const (
	ModelSubjectWidgetModeView ModelSubjectWidgetMode = iota
	ModelSubjectWidgetModeEdit
)

type ModelSubjectWidgetOptions struct {
	Etc *elementtreeconstructor.ElementTreeConstructor

	ModelSubject  *godatamodel.DataModelSubject
	InitialValues *godatamodel.DataModelSubjectDocument
	InitialMode   ModelSubjectWidgetMode

	CanEditFieldsSeparately bool

	OnSavePressed        func(*ModelSubjectWidget, map[string]interface{})
	OnCancelPressed      func(*ModelSubjectWidget)
	OnSaveFieldPressed   func(form *ModelSubjectWidget, name string, value interface{})
	OnCancelFieldPressed func(form *ModelSubjectWidget, name string)
}

type ModelSubjectWidget struct {
	options *ModelSubjectWidgetOptions

	Fields []*ModelSubjectWidgetItem

	Values *godatamodel.DataModelSubjectDocument

	Element *elementtreeconstructor.ElementMutator
}

func NewModelSubjectWidget(
	options *ModelSubjectWidgetOptions,
) (
	*ModelSubjectWidget,
	error,
) {
	self := &ModelSubjectWidget{options: options}
	return self, nil
}

func (self *ModelSubjectWidget) GetValues() (map[string]interface{}, error) {
	ret := make(map[string]interface{})
	for _, i := range self.Fields {
		v := i.GetValue()
		ret[i.options.Item.Name] = v
	}
	return ret, nil
}

func (self *ModelSubjectWidget) Render() (
	*elementtreeconstructor.ElementMutator,
	error,
) {

	self.Element = self.options.Etc.CreateElement("div")

	self.Fields = make([]*ModelSubjectWidgetItem, 0)

	for _, i := range self.options.ModelSubject.Fields {

		o := &ModelSubjectWidgetItemOptions{}

		o.Item = i

		{
			var field_init_item *godatamodel.DataModelSubjectDocumentItem

			if self.options.InitialValues != nil {
				for _, j := range self.options.InitialValues.Items {
					if j.Name == i.Name {
						field_init_item = j
						break
					}
				}
			}

			if field_init_item != nil {
				o.InitialValue = field_init_item.Value
			}
		}

		form_item, err := NewModelSubjectWidgetItem(o)
		if err != nil {
			return nil, err
		}
		self.Fields = append(self.Fields, form_item)
		self.Element.AppendChildren(form_item.Element)
	}

	return self.Element, nil
}
