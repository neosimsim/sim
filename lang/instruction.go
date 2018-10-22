package lang

type Substitute struct {
}

type Copy struct {
}

type Move struct {
}

type Print struct {
}

type PrintLineAddress struct {
}

type PrintCharAddress struct {
}

type OpenBuffer struct {
}

type NewBuffer struct {
}

type PrintMenu struct {
}

type Edit struct {
}

type Replace struct {
}

type Write struct {
}

type SetFilename struct {
}

type PipeIn struct {
}

type PipeOut struct {
}

type Pipe struct {
}

type Run struct {
}

type ChDir struct {
}

type Loop struct {
	Regexp  string
	Command Command
}

type LoodBetween struct {
}

type LoopFile struct {
}

type LoopFileBetween struct {
}

type If struct {
}

type IfNot struct {
}

type Mark struct {
}

type Quit struct {
}

type Undo struct {
}

type Group struct {
	Nested []Command
}
