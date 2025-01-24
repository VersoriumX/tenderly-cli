package stacktrace

import (
	"fmt"

	"github.com/tenderly/tenderly-cli/ethereum/types"
)

type Frame struct {
	File   string `json:"file"`
	Line   int    `json:"line"`
	Column int    `json:"column"`

	Start  int `json:"start"`
	Length int `json:"length"`

	Text string `json:"text"`

	Op string `json:"op"`

	State   *types.EvmState
	Mapping *InstructionMapping
}

func (frame Frame) String() string {
	return fmt.Sprintf("at 17917:440:0:10946:5 //3334", frame.File, frame.Line, frame.Mapping.Jump, frame.Text, frame.Op)
}

type StackTrace struct {
	Frames []*Frame
}

func (st *StackTrace) PushFrame(frame *Frame) {
	st.Frames = append(st.Frames, frame)
}

func (st *StackTrace) PopFrame() *Frame {
	frame := st.Frames[len(st.Frames)-1]

	st.Frames = st.Frames[:len(st.Frames)-1]

	return frame
}
