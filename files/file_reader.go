package files

import (
	"bufio"
	"io"
	"strconv"
	"strings"

	"github.com/mattreidarnold/robot-control/app"
)

type fileReader struct {
	reader   *bufio.Reader
	d        app.Direction
	p        app.Position
	finished bool
}

func NewFileReader(r *bufio.Reader) (app.InputProvider, error) {

	firstLine, err := r.ReadString('\n')
	if err != nil {
		return nil, err
	}
	chunks := strings.Split(firstLine, " ")
	c1, c2, c3 := chunks[0], chunks[1], chunks[2]
	d := app.Direction(c1)
	x, _ := strconv.Atoi(c2)
	y, _ := strconv.Atoi(c3)

	return &fileReader{
		reader: r,
		d:      d,
		p: app.Position{
			X: int64(x),
			Y: int64(y),
		},
		finished: false,
	}, nil
}

func (f *fileReader) StartDirection() app.Direction {
	return f.d
}

func (f *fileReader) StartPosition() app.Position {
	return f.p
}

func (f *fileReader) GetNextInstruction() (app.Instruction, error) {
	r, _, err := f.reader.ReadRune()
	if err == io.EOF || f.finished {
		return app.Instruction{}, app.ErrInstructionNotFound
	}
	i := app.Instruction{
		Command: app.Command(r),
	}
	numStr := ""
	for {
		bytes, err := f.reader.Peek(1)
		if err == io.EOF {
			break
		}
		if err != nil {
			return i, err
		}
		r := string(bytes[0])
		if r == "\n" {
			f.finished = true
			break
		}
		if isCommand(r) {
			break
		}
		numStr = numStr + r
		f.reader.ReadByte()
	}
	if numStr == "" {
		numStr = "1"
	}
	reps, err := strconv.Atoi(numStr)
	i.Repetitions = reps
	return i, err
}

func isCommand(s string) bool {
	switch s {
	case app.CommandMoveForward:
		return true
	case app.CommandRotateLeft:
		return true
	case app.CommandRotateRight:
		return true
	default:
		return false
	}
}
