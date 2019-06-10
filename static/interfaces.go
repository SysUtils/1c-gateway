package odata

// Interface for odata Entity type
type IEntity interface {
	APIEntityType() string
}

// Interface for odata PrimaryKey type
type IPrimaryKey interface {
	APIEntityType() string
	Serialize() string
}

// Interface for odata Function type
type IFunction interface {
	Name() string
	Parameters() string
}
