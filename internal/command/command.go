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

func (command *Command) UpdateCmd(cmd string) {
	command.Cmd = cmd
}

func (command *Command) UpdateArgs(args []string) {
	command.Args = args
}
