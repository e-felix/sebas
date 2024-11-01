package command

type Command struct {
	Cmd  string
	Args []string
}

func NewCommand(cmd string, args []string) *Command {
	return &Command{
		Cmd:  cmd,
		Args: args,
	}
}
