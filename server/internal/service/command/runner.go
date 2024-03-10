package command

import (
	"bytes"
	"context"
	"io"
	"log"
	"os/exec"
	"sync"
)

func (s *Service) Runner() {
	ch := make(chan struct{}, 10)
	wg := sync.WaitGroup{}
	ctx := context.Context(context.Background())
	for {
		scriptIds, err := s.ScriptsCache.GetAllKeys()
		if err != nil {
			log.Println(err)
		}
		wg.Add(len(scriptIds))
		for _, id := range scriptIds {
			ch <- struct{}{}
			go func(id int64) {
				defer func() {
					_ = s.ScriptsCache.Delete(id)
					_ = s.ExecCmdCache.Delete(id)
					wg.Done()
					<-ch
				}()

				val, _ := s.ScriptsCache.Get(id)

				script := val.(string)
				cmd := exec.Command("/bin/sh", "-c", script)
				_ = s.ExecCmdCache.Set(id, cmd)

				var outb bytes.Buffer
				cmd.Stdout = &outb

				if err = cmd.Run(); err != nil {
					log.Println(err)
					return
				}

				for {
					line, err := outb.ReadString('\n')
					if err == io.EOF {
						log.Println(string(line))
						break
					}
					err = s.Repository.CreateCommandOutput(ctx, id, string(line))
					if err != nil {
						break
					}
				}

				log.Println("command executed successfully!")
			}(id)
		}
		wg.Wait()
	}
}
