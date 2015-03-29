package main

import "fmt"

type blockAxis int

const (
	AxisY blockAxis = iota
	AxisZ
	AxisX
	AxisNone
)

func (b blockAxis) String() string {
	switch b {
	case AxisNone:
		return "none"
	case AxisX:
		return "x"
	case AxisY:
		return "y"
	case AxisZ:
		return "z"
	}
	return fmt.Sprintf("blockAxis(%d)", b)
}

type treeVariant int

const (
	treeOak treeVariant = iota
	treeSpruce
	treeBirch
	treeJungle
)

func (t treeVariant) String() string {
	switch t {
	case treeOak:
		return "oak"
	case treeSpruce:
		return "spruce"
	case treeBirch:
		return "birch"
	case treeJungle:
		return "jungle"
	}
	return fmt.Sprintf("treeVariant(%d)", t)
}

type blockLog struct {
	baseBlock
	Variant treeVariant `state:"variant,0-3"`
	Axis    blockAxis   `state:"axis,0-3"`
}

func initLog(name string) *BlockSet {
	l := &blockLog{}
	l.init(name)
	set := alloc(l)
	return set
}

func (l *blockLog) String() string {
	return l.Parent.stringify(l)
}

func (l *blockLog) clone() Block {
	return &blockLog{
		baseBlock: *(l.baseBlock.clone().(*baseBlock)),
		Variant:   l.Variant,
		Axis:      l.Axis,
	}
}

func (l *blockLog) ModelName() string {
	return l.Variant.String() + "_" + l.name
}

func (l *blockLog) ModelVariant() string {
	return fmt.Sprintf("axis=%s", l.Axis)
}

func (l *blockLog) toData() int {
	data := int(l.Variant)
	data |= int(l.Axis) << 2
	return data
}
