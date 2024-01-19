package utils

import (
	"fmt"
	"net"
	"strings"

	"github.com/Onyx-Innovators/VRConnectivity-Shield/logger"
	"github.com/miekg/dns"
)

// blockedHosts is a map of hosts that should be blocked
var blockedHosts = map[string]struct{}{
	"api.amplitude.com":                         {},
	"api2.amplitude.com":                        {},
	"api.lab.amplitude.com":                     {},
	"api.eu.amplitude.com":                      {},
	"regionconfig.amplitude.com":                {},
	"regionconfig.eu.amplitude.com":             {},
	"o1125869.ingest.sentry.io":                 {},
	"api3.amplitude.com":                        {},
	"cdn.amplitude.com":                         {},
	"info.amplitude.com":                        {},
	"static.amplitude.com":                      {},
	"api.uca.cloud.unity3d.com":                 {},
	"config.uca.cloud.unity3d.com":              {},
	"perf-events.cloud.unity3d.com":             {},
	"public.cloud.unity3d.com":                  {},
	"cdp.cloud.unity3d.com":                     {},
	"data-optout-service.uca.cloud.unity3d.com": {},
	"ecommerce.iap.unity3d.com":                 {},
}

// HandleDNSRequest handles DNS requests
func HandleDNSRequest(w dns.ResponseWriter, r *dns.Msg) {
	log := logger.GetLogger()

	m := new(dns.Msg)
	m.SetReply(r)

	for _, q := range m.Question {
		host := strings.TrimSuffix(q.Name, ".")

		if _, blocked := blockedHosts[host]; blocked {
			// Block the request by responding with 0.0.0.0
			rr, err := dns.NewRR(fmt.Sprintf("%s A 0.0.0.0", q.Name))
			if err == nil {
				m.Answer = append(m.Answer, rr)
			}
		} else {
			// Allow the request to proceed
			addr, err := net.LookupHost(host)
			if err == nil {
				rr, err := dns.NewRR(fmt.Sprintf("%s A %s", q.Name, addr[0]))
				if err == nil {
					m.Answer = append(m.Answer, rr)
				}
			}
		}
	}

	err := w.WriteMsg(m)
	if err != nil {
		log.Error("Failed to write message: ", err)
	}
}
