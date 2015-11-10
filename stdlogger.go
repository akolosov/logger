package logger

import (
  "log"
  "io"
  "fmt"
  "sync"
)

type StdLogger struct {
  // Указатель на систему логгирования
  logger *log.Logger
  // Системный mutex
  mutex sync.Mutex
}

func NewStdLogger(out io.Writer, prefix string) LoggerInterface {
  log.SetOutput(out)
  log.SetPrefix(prefix)

  return &StdLogger{logger: log.New(out, prefix, log.LstdFlags)}
}

// Функция устанавливает новый интерфейс для записи логов (out).
func (logger *StdLogger) SetOutput(out io.Writer) {
  logger.mutex.Lock()
  defer logger.mutex.Unlock()

  logger.logger = log.New(out, logger.logger.Prefix(), log.LstdFlags)
  log.SetOutput(out)
}

func (logger *StdLogger) SetPrefix(prefix string) {
  logger.mutex.Lock()
  defer logger.mutex.Unlock()

  logger.logger.SetPrefix(prefix)
  log.SetPrefix(prefix)
}

func (logger *StdLogger) SetEngine(engine LoggerInterface) {
}

func (logger *StdLogger) Print(args ...interface{}) {
  logger.logger.Print("[INFO] ", fmt.Sprint(args...))
}

func (logger *StdLogger) Printf(format string, args ...interface{}) {
  logger.logger.Printf("[INFO] "+format, args...)
}

func (logger *StdLogger) Println(args ...interface{}) {
  logger.logger.Println("[INFO]", fmt.Sprint(args...))
}

func (logger *StdLogger) Error(args ...interface{}) {
  logger.logger.Print("[ERROR] ", fmt.Sprint(args...))
}

func (logger *StdLogger) Errorf(format string, args ...interface{}) {
  logger.logger.Printf("[ERROR] "+format, args...)
}

func (logger *StdLogger) Errorln(args ...interface{}) {
  logger.logger.Println("[ERROR]", fmt.Sprint(args...))
}

func (logger *StdLogger) Fatal(args ...interface{}) {
  logger.logger.Fatal("[FATAL] ", fmt.Sprint(args...))
}

func (logger *StdLogger) Fatalf(format string, args ...interface{}) {
  logger.logger.Fatalf("[FATAL] "+format, args...)
}

func (logger *StdLogger) Fatalln(args ...interface{}) {
  logger.logger.Fatalln("[FATAL]", fmt.Sprint(args...))
}

func (logger *StdLogger) Panic(args ...interface{}) {
  logger.logger.Panic("[PANIC] ", fmt.Sprint(args...))
}

func (logger *StdLogger) Panicf(format string, args ...interface{}) {
  logger.logger.Panicf("[PANIC] "+format, args...)
}

func (logger *StdLogger) Panicln(args ...interface{}) {
  logger.logger.Panicln("[PANIC]", fmt.Sprint(args...))
}
