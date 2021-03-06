package main

import (
	"encoding/json"
	"io/ioutil"

	"github.com/iakrevetkho/components-tests/cott/domain"
	"github.com/iakrevetkho/components-tests/cott/internal/helpers"

	cl_usecase "github.com/iakrevetkho/components-tests/cott/container_launcher/usecase"
	dt_usecase "github.com/iakrevetkho/components-tests/cott/database_tester/usecase"
	tester_usecase "github.com/iakrevetkho/components-tests/cott/tester/usecase"

	"github.com/jinzhu/configor"
	"github.com/sirupsen/logrus"
)

var cfg domain.Config

func init() {
	if err := configor.Load(&cfg, "config.yaml"); err != nil {
		logrus.WithError(err).Fatal("Can't parse conf")
	}

	if err := helpers.InitLogger(&cfg); err != nil {
		logrus.WithError(err).Fatal("Couldn't init logger")
	}

	if cfgJson, err := json.Marshal(cfg); err != nil {
		logrus.WithError(err).Fatal("Couldn't serialize config to JSON")
	} else {
		// Use Infof to prevent \" symbols if using WithField
		logrus.Infof("Loaded config: %s", cfgJson)
	}
}

func main() {
	cluc, err := cl_usecase.NewContainerLauncherUsecase()
	if err != nil {
		logrus.WithError(err).Fatal(domain.COULDNT_INIT_CONTAINER_LAUNCHER)
	}

	dtuc := dt_usecase.NewDatabaseTesterUsecase(cluc)

	tuc := tester_usecase.NewTesterUsecase(cluc, dtuc)

	report, err := tuc.RunCases(cfg.TestCases)
	if err != nil {
		logrus.WithError(err).Error("test case error")
	}
	logrus.WithField("report", report).Info("test cases done")

	reportBytes, err := json.Marshal(report)
	if err != nil {
		logrus.WithError(err).Fatal("couldn't serialise report")
	}

	if err := ioutil.WriteFile(cfg.Report.FilePath, reportBytes, 0644); err != nil {
		logrus.WithError(err).Fatal("couldn't write report")
	}
}
