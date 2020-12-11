package console

import (
	"strconv"
	"strings"
)

type Console struct {
	Instructions []*Instruction
	Index        int
	Accumulator  int
}

func (c *Console) DFSDebug() int {
	q := NewFiLoQueue()

	for {
		i := c.GetInstruction()
		q.Push(i)
		break
	}
	return 0
}

func (c *Console) GetInstruction() *Instruction {
	return c.Instructions[c.Index]
}

func (c *Console) LoadInstruction(i *Instruction) {
	c.Instructions = append(c.Instructions, i)
}

func (c *Console) Process(i *Instruction) bool {
	if i.Visited {
		return true
	} else {
		i.Visited = true
	}
	switch i.Operation {
	case "nop":
		c.Index++
	case "acc":
		c.Accumulator += i.Value
		c.Index++
	case "jmp":
		c.Index += i.Value
	}
	return false
}

func (c *Console) Revert(i *Instruction) {
	i.Visited = false
	i.Reverted = true
	switch i.Operation {
	case "nop":
		c.Index--
	case "acc":
		c.Accumulator -= i.Value
		c.Index--
	case "jmp":
		c.Index -= i.Value
	}

}

func (c *Console) Run() int {
	for {
		i := c.GetInstruction()
		b := c.Process(i)
		if b {
			break
		}
	}
	return c.Accumulator
}

func (c *Console) Write(p []byte) (int, error) {
	rb := len(p)
	rawStr := string(p)
	split := strings.Split(rawStr, "\r\n")
	for _, s := range split {
		if s != "" {
			inS := strings.Split(s, " ")
			inN, err := strconv.Atoi(inS[1])
			if err != nil {
				return 0, err
			}
			i := NewInstruction(inS[0], inN)
			c.LoadInstruction(i)
		}
	}
	return rb, nil
}

type FiLoQueue struct {
	Items []*Instruction
}

func (q *FiLoQueue) Push(i *Instruction) {
	q.Items = append([]*Instruction{i}, q.Items...)
}

func (q *FiLoQueue) Pop() *Instruction {
	i := q.Items[0]
	q.Items = q.Items[1:]
	return i
}

type Instruction struct {
	Operation string
	Value     int
	Visited   bool
	Reverted  bool
	Swapped   bool
}

func (i *Instruction) Swap() bool {
	if !i.Swapped {
		switch i.Operation {
		case "jmp":
			i.Operation = "nop"
			i.Swapped = true
			return true
		case "nop":
			i.Operation = "jmp"
			i.Swapped = true
			return true
		}
	}
	return false
}

func (i *Instruction) Revert() {
	if i.Swapped {
		switch i.Operation {
		case "jmp":
			i.Operation = "nop"
			i.Swapped = false
			i.Reverted = true
		case "nop":
			i.Operation = "jmp"
			i.Swapped = false
			i.Reverted = true
		}
	}
}

func NewConsole() *Console {
	c := Console{
		Instructions: make([]*Instruction, 0),
		Index:        0,
		Accumulator:  0,
	}
	return &c
}

func NewFiLoQueue() *FiLoQueue {
	q := FiLoQueue{Items: make([]*Instruction, 0)}
	return &q
}

func NewInstruction(op string, val int) *Instruction {
	i := Instruction{
		Operation: op,
		Value:     val,
		Visited:   false,
		Reverted:  false,
	}
	return &i
}
