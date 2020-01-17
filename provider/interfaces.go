package provider

type ExternalCall interface {
	Call() (interface{},error)
}