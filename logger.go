// Пакет предостваляет функции логгировния
package logger

import (
  "io"
  "os"
  "fmt"
  "log"
  "sync"
)

type LoggerInterface interface {
  SetOutput(io.Writer)
  SetPrefix(string)
  SetEngine(LoggerInterface)

  Print(...interface{})
  Printf(string, ...interface{})
  Println(...interface{})

  Error(...interface{})
  Errorf(string, ...interface{})
  Errorln(...interface{})

  Fatal(...interface{})
  Fatalf(string, ...interface{})
  Fatalln(...interface{})

  Panic(...interface{})
  Panicf(string, ...interface{})
  Panicln(...interface{})
}

// Основной тип системы логгирования.
type Logger struct {
  // Указатель на систему логгирования
  engine LoggerInterface
  // Системный mutex
  mutex sync.Mutex
  //
  forcestdout bool
  stdout *log.Logger
}

// Функция возвращает указатель но систему логгирования.
// В качестве параметров принимает указатель на интерфейс для записи логов (out)
// и строку-префикс (prefix) для идентификации в системе логгирования.
func NewLogger(engine LoggerInterface, forcestdout bool) LoggerInterface {
  if forcestdout {
    return &Logger{engine: engine, forcestdout: forcestdout, stdout: log.New(os.Stderr, log.Prefix(), log.LstdFlags) }
  } else {
    return &Logger{engine: engine, forcestdout: false, stdout: nil}
  }
}

// Функция устанавливает новый интерфейс для записи логов (out).
func (logger *Logger) SetOutput(out io.Writer) {
  logger.engine.SetOutput(out)
}

func (logger *Logger) SetPrefix(prefix string) {
  logger.engine.SetPrefix(prefix)
  if logger.forcestdout {
    logger.stdout.SetPrefix(prefix)
  }
}

func (logger *Logger) SetEngine(engine LoggerInterface) {
  logger.mutex.Lock()
  defer logger.mutex.Unlock()

  if engine != nil {
    logger.engine = engine
    logger.engine.SetEngine(engine)
  }
}

func (logger *Logger) Print(args ...interface{}) {
  logger.engine.Print(args...)
  if logger.forcestdout {
    logger.stdout.Print("[INFO] ", fmt.Sprint(args...))
  }
}

func (logger *Logger) Printf(format string, args ...interface{}) {
  logger.engine.Printf(format, args...)
  if logger.forcestdout {
    logger.stdout.Printf("[INFO] "+format, args...)
  }
}

func (logger *Logger) Println(args ...interface{}) {
  logger.engine.Println(args...)
  if logger.forcestdout {
    logger.stdout.Println("[INFO] ", fmt.Sprint(args...))
  }
}

func (logger *Logger) Error(args ...interface{}) {
  logger.engine.Error(args...)
  if logger.forcestdout {
    logger.stdout.Print("[ERROR] ", fmt.Sprint(args...))
  }
}

func (logger *Logger) Errorf(format string, args ...interface{}) {
  logger.engine.Errorf(format, args...)
  if logger.forcestdout {
    logger.stdout.Printf("[ERROR] "+format, args...)
  }
}

func (logger *Logger) Errorln(args ...interface{}) {
  logger.engine.Errorln(args...)
  if logger.forcestdout {
    logger.stdout.Println("[ERROR] ", fmt.Sprint(args...))
  }
}

func (logger *Logger) Fatal(args ...interface{}) {
  if logger.forcestdout {
    logger.stdout.Print("[FATAL] ", fmt.Sprint(args...))
  }
  logger.engine.Fatal(args...)
}

func (logger *Logger) Fatalf(format string, args ...interface{}) {
  if logger.forcestdout {
    logger.stdout.Printf("[FATAL] "+format, args...)
  }
  logger.engine.Fatalf(format, args...)
}

func (logger *Logger) Fatalln(args ...interface{}) {
  if logger.forcestdout {
    logger.stdout.Println("[FATAL] ", fmt.Sprint(args...))
  }
  logger.engine.Fatalln(args...)
}

func (logger *Logger) Panic(args ...interface{}) {
  if logger.forcestdout {
    logger.stdout.Print("[PANIC] ", fmt.Sprint(args...))
  }
  logger.engine.Panic(args...)
}

func (logger *Logger) Panicf(format string, args ...interface{}) {
  if logger.forcestdout {
    logger.stdout.Printf("[PANIC] "+format, args...)
  }
  logger.engine.Panicf(format, args...)
}

func (logger *Logger) Panicln(args ...interface{}) {
  if logger.forcestdout {
    logger.stdout.Println("[PANIC] ", fmt.Sprint(args...))
  }
  logger.engine.Panicln(args...)
}
