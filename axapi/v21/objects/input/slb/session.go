package slb

type SessionFilter struct {
	FwdSrcAddr          string `param:"fwd_src_addr" long:"fwd-src-addr" description:"IPv4 address in dotted decimal form. Default: 0.0.0.0"`
	FwdSrcMask          string `param:"fwd_src_mask" long:"fwd-src-mask" description:"IPv4 net mask in dotted decimal form. Default: 0.0.0.0."`
	FwdDstAddr          string `param:"fwd_dst_addr" long:"fwd-dst-addr" description:"IPv4 address in dotted decimal form. Default: 0.0.0.0"`
	FwdDstMask          string `param:"fwd_dst_mask" long:"fwd-dst-mask" description:"IPv4 net mask in dotted decimal form. Default: 0.0.0.0."`
	FwdSrcIpv6Addr      string `param:"fwd_src_ipv6_addr" long:"fwd-src-ipv6_addr" description:"IPv6 address. Default: ::."`
	FwdSrcIpv6PrefixLen string `param:"fwd_src_ipv6_prefix_len" long:"fwd-src-ipv6_prefix_len" description:"Scope: 0 - 128, Default: 0."`
	FwdDstIpv6Addr      string `param:"fwd_dst_ipv6_addr" long:"fwd-dst-ipv6_addr" description:"IPv6 address. Default: ::."`
	FwdDstIpv6PrefixLen string `param:"fwd_dst_ipv6_prefix_len" long:"fwd-dst-ipv6_prefix_len" description:"Scope: 0 - 128, Default: 0."`
	FwdSrcPort          string `param:"fwd_src_port" long:"fwd-src-port" description:"Scope: 0 - 9223372036854775807, Default: 0."`
	FwdDstPort          string `param:"fwd_dst_port" long:"fwd-dst-port" description:"Scope: 0 - 9223372036854775807, Default: 0."`
	FilterName          string `param:"filter_name" long:"filter-name" description:"Length: 0 - 63. Default: ."`
}
