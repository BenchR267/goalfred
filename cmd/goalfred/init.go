package main

import (
	"bufio"
	"fmt"
	"os"
	"path"
)

const (
	defaultVersion = "1.0.0"
)

type InitCommand struct {
	Version string `short:"v" description:"The version of your workflow. default: 1.0.0"`
}

var Init InitCommand

type Workflow struct {
	Name    string
	Path    string
	Version string
}

func (i *InitCommand) Execute(args []string) error {

	path, err := pathFromArgs(args)
	if err != nil {
		return err
	}

	if i.Version == "" {
		i.Version = defaultVersion
	}

	workflow := Workflow{
		Path:    path,
		Version: i.Version,
	}

	workflow.Name = readString("What should be the name of the Workflow? ")

	err = createInfoPlist(workflow)
	if err != nil {
		return err
	}

	fmt.Println("Initialize:", workflow)

	return nil
}

func pathFromArgs(args []string) (string, error) {
	var path string
	if len(args) > 0 {
		path = args[0]
		stat, err := os.Stat(path)
		if err != nil {
			return "", err
		}
		if !stat.IsDir() {
			return "", fmt.Errorf("The given path is no directory: %s", path)
		}
	} else {
		p, err := os.Getwd()
		if err != nil {
			return "", err
		}
		path = p
	}
	return path, nil
}

func readString(desc string) string {
	fmt.Print(desc)
	s := bufio.NewScanner(os.Stdin)
	s.Scan()
	return s.Text()
}

func createInfoPlist(workflow Workflow) error {
	const fileName = "info.plist"

	filePath := path.Join(workflow.Path, fileName)

	if _, err := os.Stat(filePath); err == nil {
		return fmt.Errorf("File at path %s already exists. Aborting..", filePath)
	}

	os.Chdir(workflow.Path)
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}

	file.Write([]byte(infoPlist))

	return nil
}
