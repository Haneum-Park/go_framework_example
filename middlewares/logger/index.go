package logger

// import (
// 	"os"
// 	"sort"

// 	"github.com/sirupsen/logrus"
// )

// var Log = logrus.New()

// var fieldSeq = map[string]int{
// 	"TIME":    0,
// 	"LEVEL":   1,
// 	"METHOD":  2,
// 	"HOST":    3,
// 	"URI":     4,
// 	"status":  5,
// 	"MESSAGE": 6,
// }

// func Logger() *logrus.Logger {
// 	Log.SetFormatter(&logrus.TextFormatter{
// 		TimestampFormat:  "2006-01-02 15:04:05",
// 		FullTimestamp:    true,
// 		QuoteEmptyFields: true,
// 		FieldMap: logrus.FieldMap{
// 			logrus.FieldKeyTime:  "TIME",
// 			logrus.FieldKeyLevel: "LEVEL",
// 			logrus.FieldKeyMsg:   "MESSAGE",
// 		},
// 		SortingFunc: func(fields []string) {
// 			sort.Slice(fields, func(i, j int) bool {
// 				if fields[i] == "MESSAGE" {
// 					return false
// 				}
// 				if iIdx, oki := fieldSeq[fields[i]]; oki {
// 					if jIdx, okj := fieldSeq[fields[j]]; okj {
// 						return iIdx < jIdx
// 					}
// 					return true
// 				}
// 				return false
// 			})
// 		},
// 		PadLevelText: true,
// 	})
// 	Log.SetOutput(os.Stderr)
// 	Log.SetLevel(logrus.DebugLevel)

// 	return Log
// }
