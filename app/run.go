package app

import (
	_"errors"
	"golang.org/x/sys/windows/svc/debug"
)

// Run launches the service
func Run(wl debug.Log, svcName, environment string) error {

	s, err := setup(wl, svcName, environment)
	if err != nil {
		return err
		//return errors.Is(err, "setup")
	}

	// Your service should be launched as a GO routine
	go yourApp(s)

	return nil
}
