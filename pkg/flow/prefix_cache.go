package flow

import (
	"sync"

	bnet "github.com/bio-routing/bio-rd/net"
)

var (
	pkgPrefixCache = newPrefixCache()
)

// prefixCache per agent and VRF
type prefixCache struct {
	cache   map[prefixCacheKey]*prefixCacheEntry
	cacheMu sync.Mutex
}

type prefixCacheKey struct {
	agent  *bnet.IP
	vrf    uint64
	prefix *bnet.Prefix
}

type prefixCacheEntry struct {
	prefix   *Prefix
	refCount uint64
}

func newPrefixCache() *prefixCache {
	return &prefixCache{
		cache: make(map[prefixCacheKey]*prefixCacheEntry, 0),
	}
}

// GetPrefix gets a prefix
func GetPrefix(agent *bnet.IP, vrf uint64, p *bnet.Prefix) *Prefix {
	return pkgPrefixCache.get(agent, vrf, p)
}

// ReleasePrefix releases a prefix
func ReleasePrefix(agent *bnet.IP, vrf uint64, p *bnet.Prefix) {
	pkgPrefixCache.release(agent, vrf, p)
}

func buildPrefixCacheKey(agent *bnet.IP, vrf uint64, p *bnet.Prefix) prefixCacheKey {
	return prefixCacheKey{
		agent:  agent.Dedup(),
		vrf:    vrf,
		prefix: p.Dedup(),
	}
}

func (pfxc *prefixCache) get(agent *bnet.IP, vrf uint64, p *bnet.Prefix) *Prefix {
	k := buildPrefixCacheKey(agent, vrf, p)

	pfxc.cacheMu.Lock()
	defer pfxc.cacheMu.Unlock()

	if e, found := pfxc.cache[k]; found {
		e.refCount++
		return e.prefix
	}

	e := &prefixCacheEntry{
		prefix: &Prefix{
			Pfx:      p.ToProto(),
			Metadata: make(map[string]string),
		},
		refCount: 1,
	}

	pfxc.cache[k] = e
	return e.prefix
}

func (pfxc *prefixCache) release(agent *bnet.IP, vrf uint64, p *bnet.Prefix) {
	k := buildPrefixCacheKey(agent, vrf, p)

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
