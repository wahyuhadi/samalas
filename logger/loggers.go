package loggers

import (
	"flag"
	"io/ioutil"

	"github.com/sirupsen/logrus"
)

var (
	// -- validate logger , with default value is true
	logger = flag.Bool("logger", true, "Debug mode --logger=false or --logger=true ")
)

func SetLogger(mode, message string) {
	flag.Parse()
	log := logrus.New()
	if !*logger {
		log.Out = ioutil.Discard
	}
	switch mode {

	// - error log
	case "error":
		log.Errorf(message)

	// - warning log
	case "warning":
		log.Warningf(message)

	// - info log
	case "info":
		log.Infof(message)
	}
}
