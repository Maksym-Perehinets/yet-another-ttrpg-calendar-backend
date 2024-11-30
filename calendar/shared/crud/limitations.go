package crud

// CreateModel is an interface for models that can be created
type CreateModel interface {
	GetID() uint
}

// UpdateModel is an interface for models that can be updated
type UpdateModel interface {
	GetID() uint
}

// DeleteModel is an interface for models that can be deleted
type DeleteModel interface {
	GetID() uint
}

// ReadModel is an interface for models that can be read
type ReadModel interface {
	GetID() uint
}
