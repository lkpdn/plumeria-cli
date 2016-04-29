package v21

import (
	"github.com/lkpdn/plumeria-cli/axapi/v21/commands"
	"github.com/lkpdn/plumeria-cli/axapi/v21/commands/axapi"
	"github.com/lkpdn/plumeria-cli/axapi/v21/commands/ha"
	"github.com/lkpdn/plumeria-cli/axapi/v21/commands/nat"
	"github.com/lkpdn/plumeria-cli/axapi/v21/commands/network"
	"github.com/lkpdn/plumeria-cli/axapi/v21/commands/network/acl"
	"github.com/lkpdn/plumeria-cli/axapi/v21/commands/network/arp"
	"github.com/lkpdn/plumeria-cli/axapi/v21/commands/network/dns"
	"github.com/lkpdn/plumeria-cli/axapi/v21/commands/network/interfase"
	"github.com/lkpdn/plumeria-cli/axapi/v21/commands/network/lacp"
	"github.com/lkpdn/plumeria-cli/axapi/v21/commands/network/route"
	"github.com/lkpdn/plumeria-cli/axapi/v21/commands/network/vlan"
	"github.com/lkpdn/plumeria-cli/axapi/v21/commands/slb"
	"github.com/lkpdn/plumeria-cli/axapi/v21/commands/slb/global"
	"github.com/lkpdn/plumeria-cli/axapi/v21/commands/slb/hm"
	"github.com/lkpdn/plumeria-cli/axapi/v21/commands/slb/session"
	"github.com/lkpdn/plumeria-cli/axapi/v21/commands/slb/session/persistent"
	"github.com/lkpdn/plumeria-cli/axapi/v21/commands/slb/ssl"
	"github.com/lkpdn/plumeria-cli/axapi/v21/commands/slb/template"
	"github.com/lkpdn/plumeria-cli/axapi/v21/commands/system"
	"github.com/lkpdn/plumeria-cli/axapi/v21/commands/system/device_info/cpu"
	"github.com/lkpdn/plumeria-cli/axapi/v21/commands/system/log"
	"github.com/lkpdn/plumeria-cli/axapi/v21/commands/vrrp_a"

	"github.com/mitchellh/cli"
)

var Commands map[string]cli.CommandFactory

func init() {
	// Commands
	Commands = map[string]cli.CommandFactory{
		"authenticate": func() (cli.Command, error) {
			return &commands.Authenticate{}, nil
		},
		"axapi.method": func() (cli.Command, error) {
			return &axapi.Method{}, nil
		},
		"axapi.module": func() (cli.Command, error) {
			return &axapi.Module{}, nil
		},
		"ha.global": func() (cli.Command, error) {
			return &ha.Global{}, nil
		},
		"ha.inline_mode": func() (cli.Command, error) {
			return &ha.InlineMode{}, nil
		},
		"ha.interface": func() (cli.Command, error) {
			return &ha.Interface{}, nil
		},
		"ha.l2_inline_status": func() (cli.Command, error) {
			return &ha.L2InlineStatus{}, nil
		},
		"ha.status": func() (cli.Command, error) {
			return &ha.Status{}, nil
		},
		"nat.acl_bind": func() (cli.Command, error) {
			return &nat.AclBind{}, nil
		},
		"nat.global": func() (cli.Command, error) {
			return &nat.Global{}, nil
		},
		"nat.interface": func() (cli.Command, error) {
			return &nat.Interface{}, nil
		},
		"nat.pool": func() (cli.Command, error) {
			return &nat.Pool{}, nil
		},
		"nat.pool_group": func() (cli.Command, error) {
			return &nat.PoolGroup{}, nil
		},
		"nat.static_translation": func() (cli.Command, error) {
			return &nat.StaticTranslation{}, nil
		},
		"nat.range": func() (cli.Command, error) {
			return &nat.Range{}, nil
		},
		"network.acl.ext": func() (cli.Command, error) {
			return &acl.Ext{}, nil
		},
		"network.acl.ipv6": func() (cli.Command, error) {
			return &acl.Ipv6{}, nil
		},
		"network.acl.std": func() (cli.Command, error) {
			return &acl.Std{}, nil
		},
		"network.arp": func() (cli.Command, error) {
			return &network.Arp{}, nil
		},
		"network.arp.global": func() (cli.Command, error) {
			return &arp.Global{}, nil
		},
		"network.dns.server": func() (cli.Command, error) {
			return &dns.Server{}, nil
		},
		"network.interface": func() (cli.Command, error) {
			return &network.Interface{}, nil
		},
		"network.interface.global": func() (cli.Command, error) {
			return &interfase.Global{}, nil
		},
		"network.lacp.global": func() (cli.Command, error) {
			return &lacp.Global{}, nil
		},
		"network.route.ipv4static": func() (cli.Command, error) {
			return &route.Ipv4static{}, nil
		},
		"network.route.ipv6static": func() (cli.Command, error) {
			return &route.Ipv6static{}, nil
		},
		"network.route.mgmt_interface": func() (cli.Command, error) {
			return &route.MgmtInterface{}, nil
		},
		"network.transparent_interface": func() (cli.Command, error) {
			return &network.TransparentInterface{}, nil
		},
		"network.trunk": func() (cli.Command, error) {
			return &network.Trunk{}, nil
		},
		"network.ve": func() (cli.Command, error) {
			return &network.Ve{}, nil
		},
		"network.vlan": func() (cli.Command, error) {
			return &network.Vlan{}, nil
		},
		"network.vlan.global": func() (cli.Command, error) {
			return &vlan.Global{}, nil
		},
		"network.vlan.mac": func() (cli.Command, error) {
			return &vlan.Mac{}, nil
		},
		"slb.class_list": func() (cli.Command, error) {
			return &slb.ClassList{}, nil
		},
		"slb.glid": func() (cli.Command, error) {
			return &slb.Glid{}, nil
		},
		"slb.global.ddos_protection": func() (cli.Command, error) {
			return &global.DdosProtection{}, nil
		},
		"slb.global.log_rate_limiting": func() (cli.Command, error) {
			return &global.LogRateLimiting{}, nil
		},
		"slb.global.settings": func() (cli.Command, error) {
			return &global.Settings{}, nil
		},
		"slb.hm": func() (cli.Command, error) {
			return &slb.Hm{}, nil
		},
		"slb.hm.external": func() (cli.Command, error) {
			return &hm.External{}, nil
		},
		"slb.hm.global": func() (cli.Command, error) {
			return &hm.Global{}, nil
		},
		"slb.pbslb": func() (cli.Command, error) {
			return &slb.Pbslb{}, nil
		},
		"slb.server": func() (cli.Command, error) {
			return &slb.Server{}, nil
		},
		"slb.service_group": func() (cli.Command, error) {
			return &slb.ServiceGroup{}, nil
		},
		"slb.session.all": func() (cli.Command, error) {
			return &session.All{}, nil
		},
		"slb.session.brief": func() (cli.Command, error) {
			return &session.Brief{}, nil
		},
		"slb.session.ipv4": func() (cli.Command, error) {
			return &session.Ipv4{}, nil
		},
		"slb.session.ipv4v6": func() (cli.Command, error) {
			return &session.Ipv4v6{}, nil
		},
		"slb.session.ipv6": func() (cli.Command, error) {
			return &session.Ipv6{}, nil
		},
		"slb.session.persistent.dstip": func() (cli.Command, error) {
			return &persistent.Dstip{}, nil
		},
		"slb.session.persistent.srcip": func() (cli.Command, error) {
			return &persistent.Srcip{}, nil
		},
		"slb.session.persistent.ssl": func() (cli.Command, error) {
			return &persistent.Ssl{}, nil
		},
		"slb.session.persistent.uie": func() (cli.Command, error) {
			return &persistent.Uie{}, nil
		},
		"slb.session.radius": func() (cli.Command, error) {
			return &session.Radius{}, nil
		},
		"slb.session.sip": func() (cli.Command, error) {
			return &session.Sip{}, nil
		},
		"slb.ssl": func() (cli.Command, error) {
			return &slb.Ssl{}, nil
		},
		"slb.ssl.cert": func() (cli.Command, error) {
			return &ssl.Cert{}, nil
		},
		"slb.ssl.crl": func() (cli.Command, error) {
			return &ssl.Crl{}, nil
		},
		"slb.ssl.expiration_email": func() (cli.Command, error) {
			return &ssl.ExpirationEmail{}, nil
		},
		"slb.template.cache": func() (cli.Command, error) {
			return &template.Cache{}, nil
		},
		"slb.template.cookie_persistence": func() (cli.Command, error) {
			return &template.CookiePersistence{}, nil
		},
		"slb.template.diameter": func() (cli.Command, error) {
			return &template.Diameter{}, nil
		},
		"slb.template.dns": func() (cli.Command, error) {
			return &template.Dns{}, nil
		},
		"slb.template.dst_ip_persistence": func() (cli.Command, error) {
			return &template.DstIpPersistence{}, nil
		},
		"slb.template.http": func() (cli.Command, error) {
			return &template.Http{}, nil
		},
		"slb.template.logging": func() (cli.Command, error) {
			return &template.Logging{}, nil
		},
		"slb.template.pbslb": func() (cli.Command, error) {
			return &template.Pbslb{}, nil
		},
		"slb.template.rport": func() (cli.Command, error) {
			return &template.Rport{}, nil
		},
		"slb.template.rtsp": func() (cli.Command, error) {
			return &template.Rtsp{}, nil
		},
		"slb.template.server": func() (cli.Command, error) {
			return &template.Server{}, nil
		},
		"slb.template.server_ssl": func() (cli.Command, error) {
			return &template.ServerSsl{}, nil
		},
		"slb.template.sip": func() (cli.Command, error) {
			return &template.Sip{}, nil
		},
		"slb.template.smtp": func() (cli.Command, error) {
			return &template.Smtp{}, nil
		},
		"slb.template.src_ip_persistence": func() (cli.Command, error) {
			return &template.SrcIpPersistence{}, nil
		},
		"slb.template.ssl_sid_persistence": func() (cli.Command, error) {
			return &template.SslSidPersistence{}, nil
		},
		"slb.template.tcp": func() (cli.Command, error) {
			return &template.Tcp{}, nil
		},
		"slb.template.tcp_proxy": func() (cli.Command, error) {
			return &template.TcpProxy{}, nil
		},
		"slb.template.udp": func() (cli.Command, error) {
			return &template.Udp{}, nil
		},
		"slb.template.vip": func() (cli.Command, error) {
			return &template.Vip{}, nil
		},
		"slb.template.vport": func() (cli.Command, error) {
			return &template.Vport{}, nil
		},
		"slb.virtual_server": func() (cli.Command, error) {
			return &slb.VirtualServer{}, nil
		},
		"slb.virtual_service": func() (cli.Command, error) {
			return &slb.VirtualService{}, nil
		},
		"system.config_status": func() (cli.Command, error) {
			return &system.ConfigStatus{}, nil
		},
		"system.device_info.cpu.historical_usage": func() (cli.Command, error) {
			return &cpu.HistoricalUsage{}, nil
		},
		"system.device_info.cpu.current_usage": func() (cli.Command, error) {
			return &cpu.CurrentUsage{}, nil
		},
		"system.information": func() (cli.Command, error) {
			return &system.Information{}, nil
		},
		"system.log.server": func() (cli.Command, error) {
			return &log.Server{}, nil
		},
		"system.performance": func() (cli.Command, error) {
			return &system.Performance{}, nil
		},
		"system.resource": func() (cli.Command, error) {
			return &system.Resource{}, nil
		},
		"system.show_tech": func() (cli.Command, error) {
			return &system.ShowTech{}, nil
		},
		"system.time": func() (cli.Command, error) {
			return &system.Time{}, nil
		},
		"vrrp_a": func() (cli.Command, error) {
			return &vrrp_a.VrrpA{}, nil
		},
		"vrrp_a.interface": func() (cli.Command, error) {
			return &vrrp_a.Interface{}, nil
		},
		"vrrp_a.template": func() (cli.Command, error) {
			return &vrrp_a.Template{}, nil
		},
	}
}
