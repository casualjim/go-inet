package tree

import (
	"encoding/binary"
	"fmt"
	"math/rand"

	"github.com/gaissmai/go-inet/inet"
)

// #####################################################################
// ### generators for IPs and CIDRs
// #####################################################################

func genV4(n int) []inet.IP {
	out := make([]inet.IP, n)
	for i := 0; i < n; i++ {
		buf := make([]byte, 4)
		binary.BigEndian.PutUint32(buf, rand.Uint32())
		ip, _ := inet.NewIP(buf)
		out[i] = ip
	}
	return out
}

func genV6(n int) []inet.IP {
	out := make([]inet.IP, n)
	for i := 0; i < n; i++ {
		buf := make([]byte, 16)
		binary.BigEndian.PutUint64(buf[:8], rand.Uint64())
		binary.BigEndian.PutUint64(buf[8:], rand.Uint64())
		ip, _ := inet.NewIP(buf)
		out[i] = ip
	}
	return out
}

// #####################################################################

func genBlockV4(n int) []inet.Block {
	set := make(map[inet.Block]bool, n)

	rs := make([]inet.Block, 0, n)
	for _, v := range genV4(n) {
		ones := rand.Intn(32)
		b, _ := inet.NewBlock(fmt.Sprintf("%s/%d", v, ones))
		if val := set[b]; val {
			continue
		}
		set[b] = true
		rs = append(rs, b)
	}
	rand.Shuffle(len(rs), func(i, j int) { rs[i], rs[j] = rs[j], rs[i] })
	return rs
}

func genBlockV6(n int) []inet.Block {
	set := make(map[inet.Block]bool, n)

	rs := make([]inet.Block, 0, n)
	for _, v := range genV6(n) {
		ones := rand.Intn(128)
		b, _ := inet.NewBlock(fmt.Sprintf("%s/%d", v, ones))
		if val := set[b]; val {
			continue
		}
		set[b] = true
		rs = append(rs, b)
	}
	rand.Shuffle(len(rs), func(i, j int) { rs[i], rs[j] = rs[j], rs[i] })
	return rs
}

func genBlockMixed(n int) []inet.Block {
	rs := make([]inet.Block, 0, n)
	rs = append(rs, genBlockV4(n/2)...)
	rs = append(rs, genBlockV6(n/2)...)
	rand.Shuffle(len(rs), func(i, j int) { rs[i], rs[j] = rs[j], rs[i] })
	return rs
}