package entity

import "strings"

type Command struct {
	CommandStr string
}

func NewCommand(s ...string) *Command {
	var builder strings.Builder
	for i := 0; i < len(s); i++ {
		builder.WriteString(s[i] + " ")
	}
	r := strings.TrimSuffix(builder.String(), " ")
	return &Command{
		CommandStr: r,
	}
}

func (c *Command) ToString() string {
	return c.CommandStr
}

type CommandBuilder struct {
	Commands []*Command
}

func (b *CommandBuilder) ToString() string {
	var builder strings.Builder
	for i := 0; i < len(b.Commands); i++ {
		builder.WriteString(b.Commands[i].ToString() + ";")
	}
	return builder.String()
}

func (b *CommandBuilder) Next(c *Command) *CommandBuilder {
	if b.Commands == nil {
		b.Commands = make([]*Command, 0)
	}
	b.Commands = append(b.Commands, c)
	return b
}

func (b *CommandBuilder) Add(another *CommandBuilder) *CommandBuilder {
	if b.Commands == nil {
		b.Commands = make([]*Command, 0)
	}
	if another == nil {
		return b
	}
	if another.Commands == nil {
		return b
	}
	b.Commands = append(b.Commands, another.Commands...)
	return b
}
