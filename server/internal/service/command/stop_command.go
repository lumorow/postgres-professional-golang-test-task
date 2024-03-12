package command

import (
	"fmt"
	"os/exec"
	"strconv"
)

func (s *Service) StopCommandById(id string) error {
	i, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}
	cmd, err := s.ExecCmdCache.Get(i)
	if err != nil {
		return err
	}
	c := cmd.(*exec.Cmd)
	if err = c.Process.Kill(); err != nil {
		return fmt.Errorf("failed to kill process: %e", err)
	}
	return nil
}
