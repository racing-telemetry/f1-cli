package printer

import (
	"fmt"
	"github.com/racing-telemetry/f1-dump/internal/text/emoji"
)

func Print(emoji emoji.Emoji, s string, a ...interface{}) {
	fmt.Printf("\r%s %s\n", string(emoji), fmt.Sprintf(s, a...))
}

func Error(err error) error {
	return fmt.Errorf("\r%s Error: %s", string(emoji.Boom), err.Error())
}

func PrintError(format string, a ...interface{}) {
	fmt.Println(Error(fmt.Errorf(format, a...)))
}
