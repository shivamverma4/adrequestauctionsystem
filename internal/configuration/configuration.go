package configuration

import (
	"path/filepath"
	"runtime"

	"github.com/op/go-logging"
)

type DeploymentType string

const DEBUG DeploymentType = "debug"
const STAGING DeploymentType = "staging"
const PRODUCTION DeploymentType = "production"

var BASEPATH = getBasePath()

type Config struct {
	Namespace          string
	Deployment         DeploymentType
	LogLevel           logging.Level
	LogFilePath        string
	RequestLogFilePath string
	Port               int
}

func getBasePath() string {
	_, b, _, _ := runtime.Caller(0)
	basePath := filepath.Dir(b)
	return filepath.Join(basePath, "../..")
}
