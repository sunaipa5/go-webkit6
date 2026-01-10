package webkit

type NetworkProxyMode uint

const (
	WEBKIT_NETWORK_PROXY_MODE_DEFAULT  NetworkProxyMode = 0
	WEBKIT_NETWORK_PROXY_MODE_NO_PROXY NetworkProxyMode = 1
	WEBKIT_NETWORK_PROXY_MODE_CUSTOM   NetworkProxyMode = 2
)

type ProxyScheme string

const (
	Socks ProxyScheme = "socks://"
	Http  ProxyScheme = "http://"
)
