package logger

import (
	// "fmt"
	"io"
	"os"

	// "sort"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/sirupsen/logrus"
)

// Logrus : implement Logger
type Logrus struct {
	*logrus.Logger
}

// Logger ...
var Logger = logrus.New()

// GetEchoLogger for e.Logger
func GetEchoLogger() Logrus {
	return Logrus{Logger}
}

// Level returns logger level
func (l Logrus) Level() log.Lvl {
	switch l.Logger.Level {
	case logrus.DebugLevel:
		return log.DEBUG
	case logrus.WarnLevel:
		return log.WARN
	case logrus.ErrorLevel:
		return log.ERROR
	case logrus.InfoLevel:
		return log.INFO
	default:
		l.Panic("Invalid level")
	}

	return log.OFF
}

// SetHeader is a stub to satisfy interface
// It's controlled by Logger
func (l Logrus) SetHeader(_ string) {}

// SetPrefix It's controlled by Logger
func (l Logrus) SetPrefix(s string) {}

// Prefix It's controlled by Logger
func (l Logrus) Prefix() string {
	return ""
}

// SetLevel set level to logger from given log.Lvl
func (l Logrus) SetLevel(lvl log.Lvl) {
	switch lvl {
	case log.DEBUG:
		Logger.SetLevel(logrus.DebugLevel)
	case log.WARN:
		Logger.SetLevel(logrus.WarnLevel)
	case log.ERROR:
		Logger.SetLevel(logrus.ErrorLevel)
	case log.INFO:
		Logger.SetLevel(logrus.InfoLevel)
	default:
		l.Panic("Invalid level")
	}
}

// Output logger output func
func (l Logrus) Output() io.Writer {
	return l.Out
}

// SetOutput change output, default os.Stdout
func (l Logrus) SetOutput(w io.Writer) {
	Logger.SetOutput(w)
}

// Printj print json log
func (l Logrus) Printj(j log.JSON) {
	Logger.WithFields(logrus.Fields(j)).Print()
}

// Debugj debug json log
func (l Logrus) Debugj(j log.JSON) {
	Logger.WithFields(logrus.Fields(j)).Debug()
}

// Infoj info json log
func (l Logrus) Infoj(j log.JSON) {
	Logger.WithFields(logrus.Fields(j)).Info()
}

// Warnj warning json log
func (l Logrus) Warnj(j log.JSON) {
	Logger.WithFields(logrus.Fields(j)).Warn()
}

// Errorj error json log
func (l Logrus) Errorj(j log.JSON) {
	Logger.WithFields(logrus.Fields(j)).Error()
}

// Fatalj fatal json log
func (l Logrus) Fatalj(j log.JSON) {
	Logger.WithFields(logrus.Fields(j)).Fatal()
}

// Panicj panic json log
func (l Logrus) Panicj(j log.JSON) {
	Logger.WithFields(logrus.Fields(j)).Panic()
}

// Print string log
func (l Logrus) Print(i ...interface{}) {
	Logger.Print(i[0].(string))
}

// Debug string log
func (l Logrus) Debug(i ...interface{}) {
	Logger.Debug(i[0].(string))
}

// Info string log
func (l Logrus) Info(i ...interface{}) {
	Logger.Info(i[0].(string))
}

// Warn string log
func (l Logrus) Warn(i ...interface{}) {
	Logger.Warn(i[0].(string))
}

// Error string log
func (l Logrus) Error(i ...interface{}) {
	Logger.Error(i[0].(string))
}

// Fatal string log
func (l Logrus) Fatal(i ...interface{}) {
	Logger.Fatal(i[0].(string))
}

// Panic string log
func (l Logrus) Panic(i ...interface{}) {
	Logger.Panic(i[0].(string))
}

// var fieldSeq = map[string]int{
// 	"TIME":          0,
// 	"LEVEL":         1,
// 	"METHOD":        2,
// 	"HOST":          3,
// 	"URI":           4,
// 	"PATH":          5,
// 	"STATUS":        6,
// 	"REFERER":       7,
// 	"USER_AGENT":    8,
// 	"MESSAGE":       9,
// 	"LATENCY":       10,
// 	"LATENCY_HUMAN": 11,
// 	"BYTES_IN":      12,
// 	"BYTES_OUT":     13,
// }

func logrusMiddlewareHandler(c echo.Context, next echo.HandlerFunc) error {
	req := c.Request()
	res := c.Response()
	start := time.Now()
	if err := next(c); err != nil {
		c.Error(err)
	}
	stop := time.Now()

	p := req.URL.Path

	bytesIn := req.Header.Get(echo.HeaderContentLength)

	commonFields := logrus.Fields{
		"TIME":          time.Now().Format("2006-01-02 15:04:05"),
		"HOST":          req.Host,
		"URI":           req.RequestURI,
		"METHOD":        req.Method,
		"PATH":          p,
		"REFERER":       req.Referer(),
		"USER_AGENT":    req.UserAgent(),
		"STATUS":        res.Status,
		"LATENCY":       strconv.FormatInt(stop.Sub(start).Nanoseconds()/1000, 10),
		"LATENCY_HUMAN": stop.Sub(start).String(),
		"BYTES_IN":      bytesIn,
		"BYTES_OUT":     strconv.FormatInt(res.Size, 10),
	}

	Logger.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
		FieldMap: logrus.FieldMap{
			logrus.FieldKeyTime:  "TIME",
			logrus.FieldKeyLevel: "LEVEL",
			logrus.FieldKeyMsg:   "MESSAGE",
		},
		PrettyPrint: true,
	})

	Logger.SetOutput(os.Stdout)
	Logger.SetLevel(logrus.DebugLevel)

	logFields := Logger.WithFields(commonFields)

	if res.Status >= 400 {
		Logger.SetLevel(logrus.ErrorLevel)
		logFields.Error()
	} else if res.Status == 200 {
		Logger.SetLevel(logrus.InfoLevel)
		logFields.Info()
	} else {
		Logger.SetLevel(logrus.DebugLevel)
		logFields.Debug()
	}

	return nil
}

func HookLogger(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		return logrusMiddlewareHandler(c, next)
	}
}
