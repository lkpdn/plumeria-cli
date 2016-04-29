# plumeria-cli

Unofficial standalone aXAPI client tool in golang. (v0.1)

* Currently only support aXAPI v2.1.
* Currently only support methods that require read-only priviledge.

### commands

* If you call it with no arguments, you will see aXAPI versions supported. (currently only v2.1)


```
$ plumeria-cli
  
  You have to specify api version. try:
  $ plumeria-cli [API_VERSION] --help
   
  Available API_VERSIONs are:
     - v21
```


* to see all the available commands:


```
$ plumeria-cli v21 --help

  usage: plumeria-cli v21 [--version] [--help] <command> [<args>]

  Available commands are:
      authenticate                               aXAPI methods for starting and closing aXAPI sessions.
      axapi.method                               axapi.method method.
      axapi.module                               axapi.module method.
      ha.global                                  ha.global method.
      ha.inline_mode                             ha.inline_mode method.
      ha.interface                               ha.interface method.
      ha.l2_inline_status                        ha.l2_inline_status method.
      ha.status                                  ha.status method.
      nat.acl_bind                               nat.acl_bind method.
      nat.global                                 nat.global method.
      nat.interface                              nat.interface method.
      nat.pool                                   nat.pool method.
      nat.pool_group                             nat.pool_group method.
      nat.range                                  nat.static_translation method.
      nat.static_translation                     nat.static_translation method.
      network.acl.ext                            network.acl.ext method.
      network.acl.ipv6                           network.acl.ipv6 method.
      network.acl.std                            network.acl.std method.
      network.arp                                network.arp method.
      network.arp.global                         network.arp.global method.
      network.dns.server                         network.dns.server method.
      network.interface                          network.interface method.
      network.interface.global                   network.interface.global method.
      network.lacp.global                        network.lacp.global method.
      network.route.ipv4static                   network.route.ipv4static method.
      network.route.ipv6static                   network.route.ipv6static method.
      network.route.mgmt_interface               network.mgmt_interface method.
      network.transparent_interface              network.transparent_interface method.
      network.trunk                              network.trunk method.
      network.ve                                 network.ve method.
      network.vlan                               network.vlan method.
      network.vlan.global                        network.vlan.global method.
      network.vlan.mac                           network.vlan.mac method.
      slb.class_list                             slb.class_list method.
      slb.glid                                   slb.glid method.
      slb.global.ddos_protection                 slb.global.ddos_protection method.
      slb.global.log_rate_limiting               slb.global.log_rate_limiting method.
      slb.global.settings                        slb.global.settings method.
      slb.hm                                     slb.hm method.
      slb.hm.external                            slb.hm.external method.
      slb.hm.global                              slb.hm.global method.
      slb.pbslb                                  slb.pbslb method.
      slb.server                                 slb.server method.
      slb.service_group                          slb.service_group method.
      slb.session.all                            slb.session.all method.
      slb.session.brief                          slb.session.brief method.
      slb.session.ipv4                           slb.session.ipv4 method.
      slb.session.ipv4v6                         slb.session.ipv4v6 method.
      slb.session.ipv6                           slb.session.ipv6 method.
      slb.session.persistent.dstip               slb.session.persistent.dstip method.
      slb.session.persistent.srcip               slb.session.persistent.srcip method.
      slb.session.persistent.ssl                 slb.session.persistent.ssl method.
      slb.session.persistent.uie                 slb.session.persistent.uie method.
      slb.session.radius                         slb.session.radius method.
      slb.session.sip                            slb.session.sip method.
      slb.ssl                                    slb.ssl method.
      slb.ssl.cert                               slb.ssl.cert method.
      slb.ssl.crl                                slb.ssl.crl method.
      slb.ssl.expiration_email                   slb.ssl.expiration_email method.
      slb.template.cache                         slb.template.cache method.
      slb.template.cookie_persistence            slb.template.cookie_persistence method.
      slb.template.diameter                      slb.template.diameter method.
      slb.template.dns                           slb.template.dns method.
      slb.template.dst_ip_persistence            slb.template.dst_ip_persistence method.
      slb.template.http                          slb.template.http method.
      slb.template.logging                       slb.template.logging method.
      slb.template.pbslb                         slb.template.pbslb method.
      slb.template.rport                         slb.template.rport method.
      slb.template.rtsp                          slb.template.rtsp method.
      slb.template.server                        slb.template.server method.
      slb.template.server_ssl                    slb.template.server_ssl method.
      slb.template.sip                           slb.template.sip method.
      slb.template.smtp                          slb.template.smtp method.
      slb.template.src_ip_persistence            slb.template.src_ip_persistence method.
      slb.template.ssl_sid_persistence           slb.template.ssl_sid_persistence method.
      slb.template.tcp                           slb.template.tcp method.
      slb.template.tcp_proxy                     slb.template.tcp_proxy method.
      slb.template.udp                           slb.template.udp method.
      slb.template.vip                           slb.template.vip method.
      slb.template.vport                         slb.template.vport method.
      slb.virtual_server                         slb.virtual_server method.
      slb.virtual_service                        slb.virtual_service method.
      system.config_status                       system.config_status method.
      system.device_info.cpu.current_usage       system.time method.
      system.device_info.cpu.historical_usage    system.time method.
      system.information                         system.information method.
      system.log.server                          system.log.server method.
      system.performance                         system.performance method.
      system.resource                            system.resource method.
      system.show_tech                           system.show_tech method.
      system.time                                system.time method.
      vrrp_a                                     vrrp_a method.
      vrrp_a.interface                           vrrp_a.interface method.
      vrrp_a.template                            vrrp_a.template method.
```


* to see subcommands:


```
$ plumeria-cli v21 --help slb.template.tcp

  $ ... slb.template.tcp <subcommand>

  To see options for each <subcommand>, try
  $ ... slb.template.tcp <subdommand> --help

  Available Subcommands:
          getAll             : "slb.template.tcp.getAll method."
          search             : "slb.template.tcp.search method."

```


* to see subcommands usage:


```
$ plumeria-cli v21 slb.template.tcp search --help

  Usage:
    plumeria-cli v21 slb.template.tcp search [OPTIONS]

  Application Options:
        --session-id=  session id (optional)
    -d, --debug        debug
        --prettify     prettify JSON output
        --config-json= config JSON file path
        --name=        name of slb.template.tcp

  Help Options:
    -h, --help         Show this help message
```


## Usage example (here with jq)


```
$ plumeria-cli v21 slb.server getAll | jq '.server_list[].name'

  "app1"
  "app2"
  "app3"
```


```
$ plumeria-cli v21 system.device_info.cpu.current_usage get | jq '.Current_cpu_usage[][]."cpu_usage(%)"'

  23
  24
  23
  23
  23
  21
  11
  25
  23
  24
  24
  21
```

## Global options

* ``--version or -v``
  * see plumeria-cli version. any subcommand or arguments will be ignored.
* ``--debug or -d``
  * just to show details of how the actual request would be sent to AX Device. Currently to show http.Request structs' values.
* ``--prettify``
  * prettify JSON output. if you want to process the JSON output, unset the flag and use [jq](http://stedolan.github.io/jq/) or some other stuff is highly recommended.
* ``--config-json``
  * if you built without editing ./config/config.go for any reason, you can specify AX Device info with json file. Please see ./config/config.json.sample
