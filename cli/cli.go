package cli

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/paologaleotti/blaze-cli/util"
)

func CloneRepository(projectName string) error {
	cloneCmd := exec.Command("git", "clone", util.RepoUrl, projectName)
	return cloneCmd.Run()
}

func RemoveIgnoredFiles(projectName string) error {
	for _, file := range util.Ignored {
		err := os.RemoveAll(filepath.Join(projectName, file))
		if err != nil {
			return err
		}
	}

	return nil
}

func ReplaceProjectName(projectName string) error {
	goModPath := filepath.Join(projectName, "go.mod")
	err := util.ReplaceInFile(goModPath, "blaze", projectName)
	if err != nil {
		return fmt.Errorf("error replacing text in go.mod: %w", err)
	}

	directories := []string{"cmd", "internal", "pkg"}
	for _, dir := range directories {
		err = util.ReplaceInFiles(filepath.Join(projectName, dir), "blaze", projectName)
		if err != nil {
			return fmt.Errorf("error replacing text in %s : %w", dir, err)
		}
	}

	return nil
}

func InstallDependencies(dir string) error {
	cmd := exec.Command("go", "mod", "tidy")
	cmd.Dir = dir
	return cmd.Run()
}
