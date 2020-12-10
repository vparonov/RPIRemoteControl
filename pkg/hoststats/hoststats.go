package hoststats

import "github.com/go-ping/ping"

//HostStats - ........
type HostStats struct {
	IsAlive         bool
	ConnectionError error
	PingStats       *ping.Statistics
}

//GetHostStats - ..........
func GetHostStats(host string, roundtrips int) (*HostStats, error) {
	pinger, err := ping.NewPinger(host)

	if err != nil {
		return nil, err
	}

	hostStats := &HostStats{}
	pinger.SetPrivileged(true)
	pinger.Count = roundtrips
	err = pinger.Run()

	if err != nil {
		hostStats.IsAlive = false
		hostStats.ConnectionError = err
	} else {
		hostStats.IsAlive = true
		hostStats.PingStats = pinger.Statistics()
	}

	return hostStats, nil
}
