package interfaces

type Models interface {
	Get(interface{}) (interface{}, error)
}
