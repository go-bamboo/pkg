package casbinx

import (
	"fmt"
	"strings"

	"github.com/emberfarkas/pkg/log"
	"go.uber.org/zap/zapcore"
)

type Logger struct {
	log.ZapLogger
	enabled bool
}

func NewLogger(core zapcore.Core, enabled bool) *Logger {
	logger := log.NewLogger(core, 1)
	l := &Logger{
		ZapLogger: *logger,
		enabled:   enabled,
	}
	return l
}

// EnableLog controls whether print the message.
func (l *Logger) EnableLog(enabled bool) {
	l.enabled = enabled
}

// IsEnabled returns if logger is enabled.
func (l *Logger) IsEnabled() bool {
	return l.enabled
}

// LogModel log info related to model.
func (l *Logger) LogModel(model [][]string) {
	if !l.enabled {
		return
	}
	var str strings.Builder
	str.WriteString("Model: ")
	for _, v := range model {
		str.WriteString(fmt.Sprintf("%v\n", v))
	}
	l.ZapLogger.Infof("%v", str.String())
}

// LogEnforce log info related to enforce.
func (l *Logger) LogEnforce(matcher string, request []interface{}, result bool, explains [][]string) {
	if !l.enabled {
		return
	}

	var reqStr strings.Builder
	reqStr.WriteString("Request: ")
	for i, rval := range request {
		if i != len(request)-1 {
			reqStr.WriteString(fmt.Sprintf("%v, ", rval))
		} else {
			reqStr.WriteString(fmt.Sprintf("%v", rval))
		}
	}
	reqStr.WriteString(fmt.Sprintf(" ---> %t\n", result))

	reqStr.WriteString("Hit Policy: ")
	for i, pval := range explains {
		if i != len(explains)-1 {
			reqStr.WriteString(fmt.Sprintf("%v, ", pval))
		} else {
			reqStr.WriteString(fmt.Sprintf("%v \n", pval))
		}
	}

	l.ZapLogger.Infof("%v", reqStr.String())
}

// LogRole log info related to role.
func (l *Logger) LogRole(roles []string) {
	if !l.enabled {
		return
	}

	l.ZapLogger.Infof("Roles: %v", roles)
}

// LogPolicy log info related to policy.
func (l *Logger) LogPolicy(policy map[string][][]string) {
	if !l.enabled {
		return
	}

	var str strings.Builder
	str.WriteString("Policy: ")
	for k, v := range policy {
		str.WriteString(fmt.Sprintf("%s : %v\n", k, v))
	}

	l.ZapLogger.Infof("%v", str.String())
}
