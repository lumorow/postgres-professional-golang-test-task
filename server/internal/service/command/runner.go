package command

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os/exec"
	"sync"
)

const (
	ON                = true
	OFF               = false
	MaxCommandProcess = 10
)

// Runner Check a new scripts for start
// MaxCommandProcess - const for set maximum launch command
func (s *Service) Runner() {
	сommandLimit := make(chan struct{}, MaxCommandProcess)
	defer close(сommandLimit)
	wg := sync.WaitGroup{}
	ctx := context.Background()
	select {
	case <-s.stopSignal:
		return
	default:
		for {
			if l, err := s.ScriptsCache.GetLen(); l > 0 && err == nil {
				scriptIds, _ := s.ScriptsCache.GetAllKeys()
				wg.Add(len(scriptIds))

				consoleMode := OFF
				if len(scriptIds) == 1 {
					consoleMode = ON
				}

				for _, id := range scriptIds {
					сommandLimit <- struct{}{}
					go func(id int64) {
						s.executeCommand(ctx, id, consoleMode, &wg, сommandLimit)
					}(id)
				}
				wg.Wait()
			}
		}
	}
}

// executeCommand stating execution command
func (s *Service) executeCommand(ctx context.Context, id int64, consoleMode bool, wg *sync.WaitGroup, ch chan struct{}) {
	defer func() {
		_ = s.ScriptsCache.Delete(id)
		_ = s.ExecCmdCache.Delete(id)
		wg.Done()
		<-ch
	}()

	scanner, cmd, err := s.commandStart(id)
	if err != nil {
		log.Println(err)
		return
	}

	outputScriptCh := make(chan string, 100)
	writeDoneCh := make(chan struct{})

	defer close(writeDoneCh)

	go s.readCommandOutput(scanner, outputScriptCh)
	go s.writeCommandOutput(ctx, id, consoleMode, outputScriptCh, writeDoneCh)

	if err := scanner.Err(); err != nil {
		log.Printf("error: scanning command_id = %d output: %s", id, err)
	}

	err = cmd.Wait()
	<-writeDoneCh
	if err != nil {
		log.Printf("error: command id = %d %s", id, err)
	} else {
		log.Printf("command_id = %d executed successfully!", id)
	}
}

// commandStart Get id command and starting it in cmd.Start
// scanner reads the command output stream
func (s *Service) commandStart(id int64) (*bufio.Scanner, *exec.Cmd, error) {
	val, _ := s.ScriptsCache.Get(id)

	script := val.(string)

	cmd := exec.Command("/bin/sh", "-c", script)
	_ = s.ExecCmdCache.Set(id, cmd)

	stdout, err := cmd.StdoutPipe()

	if err != nil {
		return nil, nil, fmt.Errorf("error creating stdout pipe: %s", err)
	}

	if err = cmd.Start(); err != nil {
		return nil, nil, fmt.Errorf("error: unsuccessful starting command_id = %d: %s", id, err)
	}

	scanner := bufio.NewScanner(stdout)

	return scanner, cmd, nil
}

// ReadCommandOutput Read output and write to chan
func (s *Service) readCommandOutput(scanner *bufio.Scanner, outputScriptCh chan string) {
	defer close(outputScriptCh)

	for scanner.Scan() {
		outputScriptCh <- scanner.Text()
	}
}

// WriteCommandOutput Write output to DB && console from channel
func (s *Service) writeCommandOutput(ctx context.Context, id int64, consoleMode bool, outputScriptCh chan string, writeDoneCh chan struct{}) {
	defer func() {
		writeDoneCh <- struct{}{}
	}()

	for consoleScriptLine := range outputScriptCh {
		if consoleMode {
			log.Println(consoleScriptLine)
		}
		if err := s.Repository.CreateCommandOutput(ctx, id, consoleScriptLine); err != nil {
			log.Printf("error writing command_id = %d output to database: %s", id, err)
		}
	}
}

func (s *Service) StopRunner() {
	s.stopSignal <- struct{}{}
}
