package main

import (
	"errors"
	"fmt"
	"os/exec"
)

var (
	ErrGitNotInstalled        = errors.New("The git tool is not installed.")
	ErrInvalidTag             = errors.New("The given tag is not valid.")
	ErrNoGitDirectory         = errors.New("The given directory is not a git directory. Be sure to run from root.")
	ErrNoInfoPlist            = errors.New("No info.plist file found. Please use directory of workflow.")
	ErrPlistBuddyNotAvailable = errors.New("Couldn't find PlistBuddy in /usr/libexec/PlistBuddy, please make sure the program is there and executable!")
)

type releaseOptions struct {
	Version string `short:"v" long:"version" description:"The version string for the release." required:"true"`

	SetGitTag       bool `short:"g" long:"git" description:"Add a git version tag."`
	SetInfoPlistTag bool `short:"i" long:"infoplist" description:"Add the version in the info.plist"`
}

var release releaseOptions

func (r *releaseOptions) Execute(args []string) error {

	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = "."
	}

	if r.SetInfoPlistTag {
		infoPath, ok := infoPlistAvailable(path)
		if !ok {
			return ErrNoInfoPlist
		}
		if err := setInfoPlistVersion(infoPath, r.Version); err != nil {
			fmt.Println(err)
		}
		commitInfoPlist(infoPath, r.Version)
	}

	if r.SetGitTag {
		if !gitInstalled() {
			return ErrGitNotInstalled
		}
		if !gitInWorkingTree(path) {
			return ErrNoGitDirectory
		}
		if o, err := createTag(); err != nil {
			fmt.Printf(o)
		}
	}

	return nil
}

func createTag() (string, error) {
	if release.Version == "" {
		return "", ErrInvalidTag
	}

	c := exec.Command("git", "tag", release.Version)

	o, err := c.CombinedOutput()
	return string(o), err
}

func gitInstalled() bool {
	_, err := exec.LookPath("git")
	return err == nil
}

func gitInWorkingTree(path string) bool {
	gitPath := fmt.Sprintf("%s/.git", path)
	o, err := exec.Command("ls", gitPath).Output()
	return err == nil && len(string(o)) > 0
}

func commitInfoPlist(path, version string) {
	exec.Command("git", "add", path).Run()
	exec.Command("git", "commit", "-m", fmt.Sprintf("\"Release %s\"", version)).Run()
}

func infoPlistAvailable(path string) (string, bool) {
	infoPath := fmt.Sprintf("%s/info.plist", path)
	o, err := exec.Command("ls", infoPath).Output()
	return infoPath, err == nil && len(string(o)) > 0
}

func plistBuddyAvailable() bool {
	_, err := exec.LookPath("/usr/libexec/PlistBuddy")
	return err == nil
}

func setInfoPlistVersion(path, version string) error {
	if !plistBuddyAvailable() {
		return ErrPlistBuddyNotAvailable
	}

	exec.Command("bash", "-c", fmt.Sprintf("/usr/libexec/PlistBuddy -c \"Delete :version\" %s", path)).CombinedOutput()
	o, err := exec.Command("bash", "-c", fmt.Sprintf("/usr/libexec/PlistBuddy -c \"Add :version string %s\" %s", version, path)).CombinedOutput()
	if err != nil {
		return errors.New(string(o))
	}
	return err
}
