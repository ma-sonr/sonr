package state

import (
	"fmt"
	"time"

	"github.com/briandowns/spinner"
	"github.com/fatih/color"
	"github.com/sonr-io/sonr/third_party/types/common"
)

type Logger struct {
	common.MotorCallback
	spinner *spinner.Spinner
	logIdx  int
}

func GetCallback() *Logger {
	s := spinner.New(spinner.CharSets[4], 100*time.Millisecond)
	return &Logger{
		spinner: s,
		logIdx:  1,
	}
}

func LogErr(msg string) {
	color.Red("ERROR - %s", msg)
}

func LogWarn(msg string) {
	color.Yellow("Warning. %s", msg)
}

func LogSuccess(msg string) {
	color.Green("Success! %s", msg)
}

func LogInfo(msg string) {
	color.Blue("FYI: %s", msg)
}

func (cb *Logger) StartSpinner() {
	fmt.Println()
	cb.spinner.Start()
}

func (cb *Logger) StopSpinner(msg string) {
	cb.spinner.Stop()
	cb.logIdx = 1
	LogSuccess(msg)
}

func (cb *Logger) OnDiscover(data []byte) {
	// LogInfo("OnDiscover")
}

func (cb *Logger) OnWalletCreated(ok bool) {
	// LogInfo("OnWalletCreated")
}

func (cb *Logger) OnLog(msg string) {
	cb.logIdx++
	cb.spinner.Suffix = fmt.Sprintf(" %d. %s", cb.logIdx, msg)
}
