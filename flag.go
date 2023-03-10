package flagvar

type FlagVar interface {
	IsAssigned() bool
	String() string
}
