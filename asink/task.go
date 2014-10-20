package asink

type Execer interface {
	Exec() bool
}

type Task struct {
	Name 	string
	Process Execer
	Require string
	Group   string
	Tag     string
}

