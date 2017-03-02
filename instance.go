package hot

import (
	"path/filepath"
	"fmt"
	"github.com/Unknwon/com"
	"os"
)

type Docs struct {
	LocalRoot string
	GitTarget string
}

func (d *Docs) ReloadDocs() error {
	localRoot := d.LocalRoot
	absRoot, err := filepath.Abs(localRoot)
	if err != nil {
		return fmt.Errorf("filepath.Abs: %v", err)
	}
	// Clone new or pull to update.
	if com.IsDir(absRoot) {
		stdout, stderr, err := com.ExecCmdDir(absRoot, "git", "pull")
		if err != nil {
			return fmt.Errorf("Fail to update docs from remote source(%s): %v - %s", d.GitTarget, err, stderr)
		}
		fmt.Println(stdout)
	} else {
		os.MkdirAll(filepath.Dir(absRoot), os.ModePerm)
		stdout, stderr, err := com.ExecCmd("git", "clone", d.GitTarget, absRoot)
		if err != nil {
			return fmt.Errorf("Fail to clone docs from remote source(%s): %v - %s", d.GitTarget, err, stderr)
		}
		fmt.Println(stdout)
	}

	return nil
}
