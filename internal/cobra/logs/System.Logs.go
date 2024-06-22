package logs

import (
        "io/ioutil"
        "os"
        "log/syslog"
        logrus_syslog "github.com/sirupsen/logrus/hooks/syslog"

        "github.com/sirupsen/logrus"
        "github.com/sirupsen/logrus/hooks/writer"
)


var Log *logrus.Logger

func init() {
        Log = InitLogrus()
}

func InitLogrus() *logrus.Logger {

        log := logrus.New()
        log.SetOutput(ioutil.Discard) // Send all logs to nowhere by default

        hook, err := logrus_syslog.NewSyslogHook("", "", syslog.LOG_INFO, "")
        if err == nil {
                log.Hooks.Add(hook)
        }

        log.AddHook(&writer.Hook{ // Send logs with level higher than warning to stderr
                Writer: os.Stderr,
                LogLevels: []logrus.Level{
                        logrus.PanicLevel,
                        logrus.FatalLevel,
                        logrus.ErrorLevel,
                        logrus.WarnLevel,
                },
        })
        log.AddHook(&writer.Hook{ // Send info and debug logs to stdout
                Writer: os.Stdout,
                LogLevels: []logrus.Level{
                        logrus.InfoLevel,
                        logrus.TraceLevel,
                        logrus.DebugLevel,
                },
        })

        log.SetReportCaller(true)

        log.SetFormatter(&logrus.TextFormatter{
                ForceColors:   true,
                DisableColors: false,
                FullTimestamp: true,
        })

        log.SetLevel(logrus.TraceLevel)

        return log
}

