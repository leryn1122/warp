package git

import (
	"bytes"
	"github.com/pkg/errors"
	"os/exec"
)

func Clone(path, url, branch string) error {
	return runcmd("git", "clone", "-b", branch, "--single-branch", url, path)
}

func Update(path, commit string) error {
	return runcmd("git", "-C", path, "checkout", commit)
}

func runcmd(name string, args ...string) error {
	cmd := exec.Command(name, args...)
	bufErr := &bytes.Buffer{}
	cmd.Stderr = bufErr
	if err := cmd.Run(); err != nil {
		return errors.Wrap(err, bufErr.String())
	}
	return nil
}
