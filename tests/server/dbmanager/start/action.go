package start

import "go-skv/server/dbmanager"

func doStart(mgr dbmanager.Manager) error {
	return mgr.Start()
}
