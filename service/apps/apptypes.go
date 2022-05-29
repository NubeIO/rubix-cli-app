package apps

type Types int

//go:generate stringer -type=Types
const (
	Protocol Types = iota
)
