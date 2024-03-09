package command

func (s *Service) Runner() {
	// Запуск скрипта
	// добавляем в кеш значения
	// запускаем 10 скриптов максимум
	// обновляем статус скриптов в кеше и в бд
	// cmd = Command
	// map[int]*exec.CMD = pid:cmd
	// Start
	// [pid] -> cmd -> cmd.Process.Kill()

	// канал семаформ

	// ch := make(chan int, 10)

	// for {
	// 	for _, val := range s.Cache.Get() {
	// 		ch <- val
	// 		go func() {
	// 			//
	// 		}()
	// 	}
	// }
}
