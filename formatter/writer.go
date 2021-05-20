package formatter

import (
	"io"
	"strings"
)

type Formatter struct {
	writer io.Writer
}

func New(w io.Writer) *Formatter {
	return &Formatter{w}
}

// Print formats the words for output by writing new lines
// after each output. It also comma separates all slices of
// more than one word and may return an error
func (f *Formatter) Print(words []string) (err error) {
	commaSeparatedWords := strings.Join(words[:], ",")
	_, err = f.writer.Write([]byte(commaSeparatedWords))
	_, err = f.writer.Write([]byte("\n"))
	_, err = f.writer.Write([]byte("\n"))
	if err != nil {
		return err
	}
	return nil
}
