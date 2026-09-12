package utils

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"slices"
)

const opFile = "ops.txt" //can be changed or derived from a settings file

var ops []string

func LoadOps() {
	if _, err := os.Stat(opFile); err != nil {
		_, err = os.Create(opFile)
		Check(err)
	}

	f, err := os.OpenFile(opFile, os.O_CREATE|os.O_RDWR, 0665)
	if err != nil {
		Check(err)
	}

	defer f.Close()

	reader := bufio.NewReader(f)

	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		} else {
			Check(err)
		}

		ops = append(ops, line)
	}
}

func IsOp(name string) bool {
	return slices.Contains(ops, name)
}

func AddOp(name string) bool {
	if IsOp(name) {
		return false
	}

	f, err := os.OpenFile(opFile, os.O_APPEND|os.O_RDWR, 0660)
	if err != nil {
		Check(err)
	}

	defer f.Close()

	fmt.Fprintf(f, "%s\n", name)
	ops = append(ops, name)

	return true
}

func RemoveOp(name string) bool {
	if !IsOp(name) {
		return false
	}

	f, err := os.OpenFile(opFile, os.O_RDWR, 0660)
	if err != nil {
		Check(err)
	}

	defer f.Close()

	removeOp(name)
	for _, op := range ops {
		fmt.Fprintf(f, "%s\n", op)
	}

	return true
}

func removeOp(name string) {
	var index int
	for i, op := range ops {
		if op == name {
			index = i
			break
		}
	}

	ops = append(ops[:index], ops[index+1:]...)
}
