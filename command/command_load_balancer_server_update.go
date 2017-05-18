package command

import (
	"fmt"
	"github.com/sacloud/libsacloud/sacloud"
)

func LoadBalancerServerUpdate(ctx Context, params *ServerUpdateLoadBalancerParam) error {

	client := ctx.GetAPIClient()
	api := client.GetLoadBalancerAPI()
	p, e := api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("LoadBalancerServerUpdate is failed: %s", e)
	}

	// validation keys
	if (!ctx.IsSet("vip") || !ctx.IsSet("port")) && !ctx.IsSet("vip-index") {
		return fmt.Errorf("%q or %q and %q: is required", "vip-index", "vip", "port")
	}

	var vip *sacloud.LoadBalancerSetting
	if params.VIPIndex > 0 {
		// use VIP index
		if params.VIPIndex-1 >= len(p.Settings.LoadBalancer) {
			return fmt.Errorf("vip-index(%d) is out of range", params.VIPIndex)
		}
		vip = p.Settings.LoadBalancer[params.VIPIndex-1]
	} else {
		// use VIP and port
		for _, v := range p.Settings.LoadBalancer {
			if v.VirtualIPAddress == params.VIP && v.Port == fmt.Sprintf("%d", params.Port) {
				vip = v
				break
			}
		}
		if vip == nil {
			return fmt.Errorf("VIP(%s:%d) is not found", params.VIP, params.Port)
		}
	}

	// find target server
	var server *sacloud.LoadBalancerServer
	for _, s := range vip.Servers {
		if s.IPAddress == params.Ipaddress {
			server = s
			break
		}
	}
	if server == nil {
		return fmt.Errorf("server(%s) is not found", params.Ipaddress)
	}

	if ctx.IsSet("enabled") {
		server.Enabled = "False"
		if params.Enabled {
			server.Enabled = "True"
		}
	}

	if ctx.IsSet("protocol") {

		// validate when http/https
		if params.Protocol == "http" || params.Protocol == "https" {
			if server.HealthCheck.Protocol != "http" && server.HealthCheck.Protocol != "https" {
				if !ctx.IsSet("path") {
					return fmt.Errorf("%q: is required when protocol is http/https", "path")
				}
				if !ctx.IsSet("response-code") {
					return fmt.Errorf("%q: is required when protocol is http/https", "response-code")
				}
			}
		}

		server.HealthCheck.Protocol = params.Protocol
	}

	switch server.HealthCheck.Protocol {
	case "http", "https":
		if ctx.IsSet("path") {
			server.HealthCheck.Path = params.Path
		}
		if ctx.IsSet("response-code") {
			server.HealthCheck.Status = fmt.Sprintf("%d", params.ResponseCode)
		}
	default:
		server.HealthCheck.Path = ""
		server.HealthCheck.Status = ""
	}

	p, err := client.LoadBalancer.Update(params.Id, p)
	if err != nil {
		return fmt.Errorf("LoadBalancerServerUpdate is failed: %s", err)
	}
	_, err = client.LoadBalancer.Config(params.Id)
	if err != nil {
		return fmt.Errorf("LoadBalancerServerUpdate is failed: %s", err)
	}

	return nil
}
