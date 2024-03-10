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
	ON  = true
	OFF = false
)

// Runner check a new scripts for start
func (s *Service) Runner() {
	ch := make(chan struct{}, 10)
	wg := sync.WaitGroup{}
	ctx := context.Background()
	for {
		scriptIds, _ := s.ScriptsCache.GetAllKeys()
		wg.Add(len(scriptIds))
		ConsoleMode := OFF
		if len(scriptIds) == 1 {
			ConsoleMode = ON
		}

		for _, id := range scriptIds {
			ch <- struct{}{}
			go func(id int64) {
				fmt.Println(len(scriptIds))
				defer func() {
					_ = s.ScriptsCache.Delete(id)
					_ = s.ExecCmdCache.Delete(id)
					wg.Done()
					<-ch
				}()

				val, _ := s.ScriptsCache.Get(id)

				script := val.(string)

				_ = script

				cmd := exec.Command("/bin/sh", "-c", script)
				_ = s.ExecCmdCache.Set(id, cmd)

				stdout, err := cmd.StdoutPipe()

				if err != nil {
					log.Println("error creating stdout pipe:", err)
					return
				}

				if err := cmd.Start(); err != nil {
					log.Println("error starting command:", err)
					return
				}

				scanner := bufio.NewScanner(stdout)

				outputScriptCh := make(chan string)

				fmt.Println(ConsoleMode)

				go s.ReadCommandOutput(scanner, outputScriptCh)

				s.WriteCommandOutput(ctx, id, outputScriptCh, ConsoleMode)

				if err := scanner.Err(); err != nil {
					log.Println("error scanning command output:", err)
				}

				if err := cmd.Wait(); err != nil {
					log.Println("error waiting for command:", err)
				} else {
					log.Println("command executed successfully!")
				}
			}(id)
		}
		wg.Wait()
	}
}

// ReadCommandOutput Read output and write to chan
func (s *Service) ReadCommandOutput(scanner *bufio.Scanner, outputScriptCh chan string) {
	defer close(outputScriptCh)
	for scanner.Scan() {
		outputScriptCh <- scanner.Text()
	}
}

// WriteCommandOutput Write output to DB && console from channel
func (s *Service) WriteCommandOutput(ctx context.Context, id int64, outputCh chan string, consoleMode bool) {
	for consoleScriptLine := range outputCh {
		if consoleMode {
			log.Println(consoleScriptLine)
		}
		if err := s.Repository.CreateCommandOutput(ctx, id, consoleScriptLine); err != nil {
			log.Println("Error writing command output to database:", err)
		}
	}
}
