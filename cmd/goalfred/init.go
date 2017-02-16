package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path"
	"strings"
)

const (
	defaultVersion = "1.0.0"
	defaultKeyword = "keyword"

	Tools        CategoryType = "Tools"
	Internet     CategoryType = "Internet"
	Productivity CategoryType = "Productivity"
)

type InitCommand struct {
	Version string `short:"v" description:"The version of your workflow. default: 1.0.0"`
}

var Init InitCommand

type CategoryType string

type Workflow struct {
	Author      string
	Name        string
	Path        string
	BundleId    string
	Version     string
	Category    CategoryType
	Description string
	Website     string
	Keyword     string
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
	workflow.Author = readString("What is the name of the author? ")
	workflow.BundleId = readString("What is the bundle identifier of the workflow? ")
	read := readString("Please select a category \n(0 - Tools 1 - Internet 2 - Productivity) ")
	switch read {
	case "1":
		workflow.Category = Internet
	case "2":
		workflow.Category = Productivity
	default:
		workflow.Category = Tools
	}
	workflow.Description = readString("How would you shortly describe your workflow? ")
	workflow.Website = readString("If you have any, what is the website of your workflow? ")
	workflow.Keyword = readString("What should be the keyword to trigger the workflow? (keyword)")

	if workflow.Keyword == "" {
		workflow.Keyword = defaultKeyword
	}

	err = createInfoPlist(workflow)
	if err != nil {
		return err
	}

	err = createMainFile()
	if err != nil {
		return err
	}

	compileMainFile()

	fmt.Println("\nFinished initializing. Start by editing the main.go file! 🚀")

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

	file.Write([]byte(injectWorkflowInInfoPlist(workflow)))

	return nil
}

func injectWorkflowInInfoPlist(workflow Workflow) string {
	content := infoPlistTemplate

	content = strings.Replace(content, ":bundle:", workflow.BundleId, 1)
	content = strings.Replace(content, ":category:", string(workflow.Category), 1)
	content = strings.Replace(content, ":author:", workflow.Author, 1)
	content = strings.Replace(content, ":description:", workflow.Description, 1)
	content = strings.Replace(content, ":name:", workflow.Name, 2)
	content = strings.Replace(content, ":website:", workflow.Website, 1)
	content = strings.Replace(content, ":keyword:", workflow.Keyword, 1)

	return content
}

func createMainFile() error {
	file, err := os.Create("main.go")
	if err != nil {
		return err
	}

	file.Write([]byte(mainFileTemplate))

	return nil
}

func compileMainFile() {
	exec.Command("bash", "-c", "go build *.go").Run()
}
