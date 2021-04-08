package godatamodel

// NOTE: this file must contain only data types. no code, no functions

// type DataModelOptions struct {
// 	// CheckPermission func(permission_type   subject *DataModelSubject)
// }

type DataModelSubjectDocumentItem struct {
	Name    string
	Value   interface{}
	Null    bool
	Default bool
}

// this is used to insert or get documents
type DataModelSubjectDocument struct {
	Items []*DataModelSubjectDocumentItem
}

// this is root structure of tree like structure of this package functionality
type DataModel struct {
	Subjects []*DataModelSubject
}

// type DataModelPermissionLists struct {
// 	WhoCanCreate []string
// 	WhoCanEdit   []string
// 	WhoCanDelete []string
// 	WhoCanList   []string
// 	WhoCanSearch []string
// }

// describes table and it's fields
type DataModelSubject struct {
	Name         string
	DisplayTitle string
	Description  string
	ForExport    bool // clients can request this Model represintation
	Fields       []*DataModelSubjectField
	// DataModelPermissionLists
}

type DataModelSubjectFieldType string

const (
	DataModelSubjectFieldTypeBool        DataModelSubjectFieldType = "bool"
	DataModelSubjectFieldTypeInt                                   = "int"
	DataModelSubjectFieldTypeUInt                                  = "uint"
	DataModelSubjectFieldTypeFloat                                 = "float"
	DataModelSubjectFieldTypeDataDateUTC                           = "dateutc"
	DataModelSubjectFieldTypeString                                = "string"
	DataModelSubjectFieldTypeBytes                                 = "bytes"
	DataModelSubjectFieldTypeUUID                                  = "uuid"
)

// type DataModelSubjectFieldMode uint

// const (
// 	DataModelSubjectFieldModeNormal DataModelSubjectFieldMode = iota
// 	DataModelSubjectFieldModeRelation
// )

type DataModelSubjectField struct {
	// Mode                  DataModelSubjectFieldMode
	Name                  string
	Type                  DataModelSubjectFieldType
	CanBeNull             bool
	DefaultCanBeUsed      bool
	DefaultValue          interface{}
	DisplayTitle          string
	Description           string
	DisplayMustBeCensored bool
	ForStorage            bool // store/load to/from db
	ForUserView           bool // user can view
	ForUserSet            bool // user can change
	ForTransport          bool // for network transfers // TODO: experimental
	// RelativeTable         string

	// OtherModelValues      bool
	// OtherModelName        string
	// OtherModelIdColumn    string
	// OtherModelValueColumn string
	// DataModelPermissionLists
}
