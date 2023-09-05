package logger

import (
	"fmt"
	"log"
	"strings"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	//log.SetPrefix(xxx)
}

func Fatal(f interface{}, v ...interface{}) {
	s := "[Fatal] " + formatLog(f, v...)
	log.Output(2, s)
}

func Emer(f interface{}, v ...interface{}) {
	s := "[Emer] " + formatLog(f, v...)
	log.Output(2, s)
}

func Alert(f interface{}, v ...interface{}) {
	s := "[Alert] " + formatLog(f, v...)
	log.Output(2, s)
}

func Crit(f interface{}, v ...interface{}) {
	s := "[Crit] " + formatLog(f, v...)
	log.Output(2, s)
}

func Error(f interface{}, v ...interface{}) {
	s := "[Error] " + formatLog(f, v...)
	log.Output(2, s)
}

func ErrorPrint(err error) {
	s := "[Error] " + err.Error()
	log.Output(2, s)
}

func Warn(f interface{}, v ...interface{}) {
	s := "[Warn] " + formatLog(f, v...)
	log.Output(2, s)
}

func Info(f interface{}, v ...interface{}) {
	s := "[Info] " + formatLog(f, v...)
	log.Output(2, s)
}

func Println(f interface{}, v ...interface{}) {
	s := "[Info] " + formatLog(f, v...)
	log.Output(2, s)
}

func Debug(f interface{}, v ...interface{}) {
	s := "[Debug] " + formatLog(f, v...)
	log.Output(2, s)
}

func Trace(f interface{}, v ...interface{}) {
	s := "[Trace] " + formatLog(f, v...)
	log.Output(2, s)
}

func formatLog(f interface{}, v ...interface{}) string {
	var msg string
	switch f.(type) {
	case string:
		msg = f.(string)
		if len(v) == 0 {
			return msg
		}
		if strings.Contains(msg, "%") && !strings.Contains(msg, "%%") {
			//format string
		} else {
			//do not contain format char
			msg += strings.Repeat(" %v", len(v))
		}
	default:
		msg = fmt.Sprint(f)
		if len(v) == 0 {
			return msg
		}
		msg += strings.Repeat(" %v", len(v))
	}
	return fmt.Sprintf(msg, v...)
}
