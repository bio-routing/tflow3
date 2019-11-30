package flow

import (
	"sync"

	bnet "github.com/bio-routing/bio-rd/net"
)

var (
	pkgIPCache = newIPCache()
)

// ipCache per VRF
type ipCache struct {
	cache   map[ipCacheKey]*ipCacheEntry
	cacheMu sync.Mutex
}

type ipCacheKey struct {
	vrf uint64
	ip  *bnet.IP
}

type ipCacheEntry struct {
	ip       *IPAddress
	refCount uint64
}

func newIPCache() *ipCache {
	return &ipCache{
		cache: make(map[ipCacheKey]*ipCacheEntry, 0),
	}
}

// GetIP gets an IP
func GetIP(vrf uint64, ip *bnet.IP) *IPAddress {
	return pkgIPCache.get(vrf, ip)
}

// ReleaseIP releases an ip
func ReleaseIP(vrf uint64, ip *bnet.IP) {
	pkgIPCache.release(vrf, ip)
}

func buildIPCacheKey(vrf uint64, ip *bnet.IP) ipCacheKey {
	return ipCacheKey{
		vrf: vrf,
		ip:  ip,
	}
}

func (pfxc *ipCache) get(vrf uint64, ip *bnet.IP) *IPAddress {
	k := buildIPCacheKey(vrf, ip)

	pfxc.cacheMu.Lock()
	defer pfxc.cacheMu.Unlock()

	if e, found := pfxc.cache[k]; found {
		e.refCount++
		return e.ip
	}

	e := &ipCacheEntry{
		ip: &IPAddress{
			Addr:     ip.ToProto(),
			Metadata: make(map[string]string),
		},
		refCount: 1,
	}

	pfxc.cache[k] = e
	return e.ip
}

func (pfxc *ipCache) release(vrf uint64, ip *bnet.IP) {
	k := buildIPCacheKey(vrf, ip)

	pfxc.cacheMu.Lock()
	defer pfxc.cacheMu.Unlock()

	if e, found := pfxc.cache[k]; found {
		e.refCount--
		if e.refCount == 0 {
			delete(pfxc.cache, k)
		}

		return
	}

	panic("release called for non existing object")
}
