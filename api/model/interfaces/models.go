package interfaces

type Models interface {
	Find(interface{}) (interface{}, error)
	Create(interface{}) (bool, error)
	Remove(interface{}) (bool, error)
}
