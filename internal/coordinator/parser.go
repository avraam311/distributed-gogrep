package coordinator

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/spf13/pflag"
)

type Parser struct {
	FlagA    int
	FlagB    int
	FlagC    int
	FlagCc   bool
	FlagI    bool
	FlagV    bool
	FlagF    bool
	FlagN    bool
	Template string
	Strings  []string
}

func NewParser() (*Parser, error) {
	flagA := pflag.IntP("A", "A", 0, "print strings after")
	flagB := pflag.IntP("B", "B", 0, "print string before")
	flagC := pflag.IntP("C", "C", 0, "print string after and before")
	flagCc := pflag.BoolP("c", "c", false, "print number of matches")
	flagI := pflag.BoolP("i", "i", false, "ignore registre")
	flagV := pflag.BoolP("v", "v", false, "print strings that don't match")
	flagF := pflag.BoolP("F", "F", false, "consider template as string")
	flagN := pflag.BoolP("n", "n", false, "print number of string before each string")
	pflag.Parse()
	fileName := pflag.Arg(1)
	template := pflag.Arg(0)

	strs, err := getRows(fileName)
	if err != nil {
		return nil, err
	}
	flags := Parser{
		FlagA:    *flagA,
		FlagB:    *flagB,
		FlagC:    *flagC,
		FlagCc:   *flagCc,
		FlagI:    *flagI,
		FlagV:    *flagV,
		FlagF:    *flagF,
		FlagN:    *flagN,
		Template: template,
		Strings:  strs,
	}

	return &flags, nil
}

func getRows(fileName string) ([]string, error) {
	var data []byte
	var err error

	if fileName == "" {
		data, err = io.ReadAll(os.Stdin)
		if err != nil {
			return nil, fmt.Errorf("error reading from stdin")
		}
	} else {
		data, err = os.ReadFile(fileName)
		if err != nil {
			log.Println(os.Getwd())
			return nil, fmt.Errorf("error reading from file")
		}
	}

	rows := strings.Split(strings.ReplaceAll(string(data), "\r\n", "\n"), "\n")
	return rows, nil
}
