package app

import "time"

// The wrapper of your app
func yourApp(s server) {

	s.winlog.Info(1, "In app.yourApp")

	for{
		// This is just some sample code to do something
		time.Sleep(1 * time.Second)
		s.winlog.Info(1, "Still running")

		time.Sleep(2 * time.Second)
		s.winlog.Info(1, "And running")

		time.Sleep(3 * time.Second)
		s.winlog.Info(1, "But the service will keep running")

		// Notice that if this exits, the service continues to run
		// You can launch web servers, etc.

		time.Sleep(1 * time.Second)
	}
}
