package logger

import (
  "io"
  "os"
  "fmt"
  "sync"

  "github.com/fluent/fluent-logger-golang/fluent"
)

type FluentdLogger struct {
  // Указатель на систему логгирования
  logger *fluent.Fluent
  // Системный mutex
  mutex sync.Mutex
  //
  tag, host, pid string
}

type FluentdMessage struct {
  Hostname string   `msg:"host",json:"host"`
  PID string        `msg:"pid",json:"pid"`
  Facility string   `msg:"facility",json:"facility"`
  Message string    `msg:"message",json:"message"`
}

func NewFluentdLogger(tag, host, pid string) LoggerInterface {
  logger, _ := fluent.New(fluent.Config{})
  return &FluentdLogger{ logger: logger, tag: tag, host: host, pid: pid }
}

// Функция устанавливает новый интерфейс для записи логов (out).
func (logger *FluentdLogger) SetOutput(out io.Writer) {
}

func (logger *FluentdLogger) SetPrefix(prefix string) {
  logger.tag = prefix
}

func (logger *FluentdLogger) SetEngine(engine LoggerInterface) {
}

func (logger *FluentdLogger) Print(args ...interface{}) {
  logger.logger.Post(logger.tag, &FluentdMessage{ Hostname: logger.host, PID: logger.pid, Facility: "INFO", Message: fmt.Sprint(args...) })
}

func (logger *FluentdLogger) Printf(format string, args ...interface{}) {
  logger.logger.Post(logger.tag, &FluentdMessage{ Hostname: logger.host, PID: logger.pid, Facility: "INFO", Message: fmt.Sprintf(format, args...) })
}

func (logger *FluentdLogger) Println(args ...interface{}) {
  logger.logger.Post(logger.tag, &FluentdMessage{ Hostname: logger.host, PID: logger.pid, Facility: "INFO", Message: fmt.Sprint(args...) })
}

func (logger *FluentdLogger) Error(args ...interface{}) {
  logger.logger.Post(logger.tag, &FluentdMessage{ Hostname: logger.host, PID: logger.pid, Facility: "ERROR", Message: fmt.Sprint(args...) })
}

func (logger *FluentdLogger) Errorf(format string, args ...interface{}) {
  logger.logger.Post(logger.tag, &FluentdMessage{ Hostname: logger.host, PID: logger.pid, Facility: "ERROR", Message: fmt.Sprintf(format, args...) })
}

func (logger *FluentdLogger) Errorln(args ...interface{}) {
  logger.logger.Post(logger.tag, &FluentdMessage{ Hostname: logger.host, PID: logger.pid, Facility: "ERROR", Message: fmt.Sprint(args...) })
}

func (logger *FluentdLogger) Fatal(args ...interface{}) {
  logger.logger.Post(logger.tag, &FluentdMessage{ Hostname: logger.host, PID: logger.pid, Facility: "FATAL", Message: fmt.Sprint(args...) })
  os.Exit(1)
}

func (logger *FluentdLogger) Fatalf(format string, args ...interface{}) {
  logger.logger.Post(logger.tag, &FluentdMessage{ Hostname: logger.host, PID: logger.pid, Facility: "FATAL", Message: fmt.Sprintf(format, args...) })
  os.Exit(1)
}

func (logger *FluentdLogger) Fatalln(args ...interface{}) {
  logger.logger.Post(logger.tag, &FluentdMessage{ Hostname: logger.host, PID: logger.pid, Facility: "FATAL", Message: fmt.Sprint(args...) })
  os.Exit(1)
}

func (logger *FluentdLogger) Panic(args ...interface{}) {
  logger.logger.Post(logger.tag, &FluentdMessage{ Hostname: logger.host, PID: logger.pid, Facility: "PANIC", Message: fmt.Sprint(args...) })
  panic(fmt.Sprint(args...))
}

func (logger *FluentdLogger) Panicf(format string, args ...interface{}) {
  logger.logger.Post(logger.tag, &FluentdMessage{ Hostname: logger.host, PID: logger.pid, Facility: "PANIC", Message: fmt.Sprintf(format, args...) })
  panic(fmt.Sprintf(format, args...))
}

func (logger *FluentdLogger) Panicln(args ...interface{}) {
  logger.logger.Post(logger.tag, &FluentdMessage{ Hostname: logger.host, PID: logger.pid, Facility: "PANIC", Message: fmt.Sprint(args...) })
  panic(fmt.Sprint(args...))
}
