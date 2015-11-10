package logger

import (
  "io"
  "os"
  "fmt"
  "sync"
  "log"
  "log/syslog"
)

type SyslogdLogger struct {
  // Указатель на систему логгирования
  logger *syslog.Writer
  // Системный mutex
  mutex sync.Mutex
}


func NewSyslogdLogger(tag, proto, addr string) LoggerInterface {
  logger, err := syslog.Dial(proto, addr, syslog.LOG_ERR | syslog.LOG_INFO, tag)
  if err != nil {
    panic("Can't init SyslogD connection")
  }
  log.SetOutput(logger)
  log.SetPrefix(tag)
  return &SyslogdLogger{ logger: logger }
}

// Функция устанавливает новый интерфейс для записи логов (out).
func (logger *SyslogdLogger) SetOutput(out io.Writer) {
}

func (logger *SyslogdLogger) SetPrefix(prefix string) {
}

func (logger *SyslogdLogger) SetEngine(engine LoggerInterface) {
}

func (logger *SyslogdLogger) Print(args ...interface{}) {
  logger.logger.Info(fmt.Sprint(args...))
}

func (logger *SyslogdLogger) Printf(format string, args ...interface{}) {
  logger.logger.Info(fmt.Sprintf(format, args...))
}

func (logger *SyslogdLogger) Println(args ...interface{}) {
  logger.logger.Info(fmt.Sprint(args...))
}

func (logger *SyslogdLogger) Error(args ...interface{}) {
  logger.logger.Err(fmt.Sprint(args...))
}

func (logger *SyslogdLogger) Errorf(format string, args ...interface{}) {
  logger.logger.Err(fmt.Sprintf(format, args...))
}

func (logger *SyslogdLogger) Errorln(args ...interface{}) {
  logger.logger.Err(fmt.Sprint(args...))
}

func (logger *SyslogdLogger) Fatal(args ...interface{}) {
  logger.logger.Crit(fmt.Sprint(args...))
  os.Exit(1)
}

func (logger *SyslogdLogger) Fatalf(format string, args ...interface{}) {
  logger.logger.Crit(fmt.Sprintf(format, args...))
  os.Exit(1)
}

func (logger *SyslogdLogger) Fatalln(args ...interface{}) {
  logger.logger.Crit(fmt.Sprint(args...))
  os.Exit(1)
}

func (logger *SyslogdLogger) Panic(args ...interface{}) {
  logger.logger.Emerg(fmt.Sprint(args...))
  panic(fmt.Sprint(args...))
}

func (logger *SyslogdLogger) Panicf(format string, args ...interface{}) {
  logger.logger.Emerg(fmt.Sprintf(format, args...))
  panic(fmt.Sprintf(format, args...))
}

func (logger *SyslogdLogger) Panicln(args ...interface{}) {
  logger.logger.Emerg(fmt.Sprint(args...))
  panic(fmt.Sprint(args...))
}
