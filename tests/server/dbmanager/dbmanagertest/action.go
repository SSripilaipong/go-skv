package dbmanagertest

import "go-skv/server/dbmanager"

func DoStart(mgr dbmanager.Manager) error {
	return mgr.Start()
}

func DoStop(mgr dbmanager.Manager) error {
	return mgr.Stop()
}
