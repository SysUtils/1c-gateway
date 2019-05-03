package odata

type IEntity interface {
	APIEntityType() string
}

type IPrimaryKey interface {
	APIEntityType() string
	Serialize() string
}

type IFunction interface {
	Name() string
	Parameters() string
}
