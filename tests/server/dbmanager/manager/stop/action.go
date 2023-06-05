package stop

import "go-skv/server/dbmanager"

func doStop(mgr dbmanager.Manager) error {
	return mgr.Stop()
}
