package logger

import (
  "os"
  "testing"

  . "github.com/onsi/ginkgo"
  . "github.com/onsi/gomega"
)

func TestLoggers(t *testing.T) {
  RegisterFailHandler(Fail)
  RunSpecs(t, "Logger Suite")
}

var logger LoggerInterface

var _ = Describe("Logger", func() {
  It("Should not be nil", func(){
    logger = NewLogger(NewStdLogger(os.Stderr, "TEST: "), true)
    Expect(logger).ShouldNot(BeNil())
  })

  It("Should not be nil", func(){
    logger.SetOutput(os.Stderr)
    Expect(logger).ShouldNot(BeNil())
  })

  It("Should not be nil", func(){
    logger.SetPrefix("BLA-BLA-BLA: ")
    Expect(logger).ShouldNot(BeNil())
  })

  It("Should not be nil", func(){
    logger.SetEngine(NewStdLogger(os.Stderr, "TEST: "))
    Expect(logger).ShouldNot(BeNil())
  })

  Describe("#Outputs", func() {
    It("Should be nil", func(){
      catchPanic := func() {
        catchPanicInside := func() {
          catchPanicInsideInside := func() {
            if e := recover(); e != nil {
            }
          }
          defer catchPanicInsideInside()

          if e := recover(); e != nil {
            logger.Panicln("Testing 1 2 3")
          }
        }

        defer catchPanicInside()

        if e := recover(); e != nil {
          logger.Panicf("%s", "Testing 1 2 3")
        }
      }

      defer catchPanic()

      logger.Print("Testing 1 2 3")
      logger.Printf("%s", "Testing 1 2 3")
      logger.Println("Testing 1 2 3")
      logger.Error("Testing 1 2 3")
      logger.Errorf("%s", "Testing 1 2 3")
      logger.Errorln("Testing 1 2 3")
      logger.Panic("Testing 1 2 3")
      logger.Panicln("Testing 1 2 3")
    })
  })
})
