package main

import (
	"fmt"

	"github.com/seancfoley/ipaddress-go/ipaddr"
)

func main() {
	blockStrs := []string{
		"134.70.24.0/21", "134.70.25.0/30", "134.70.27.0/28", "147.154.59.128/26", "147.154.60.0/30", "130.35.203.118/32", "147.154.27.228/32",
	}
	blocks := make([]*ipaddr.IPAddress, 0, len(blockStrs))
	for _, str := range blockStrs {
		blocks = append(blocks,
			ipaddr.NewIPAddressString(str).GetAddress().ToPrefixBlock())
	}
	trie := ipaddr.AddressTrie{}
	for _, block := range blocks {
		trie.Add(block.ToAddressBase())
	}
	fmt.Printf("trie is %v\n", trie)
	for _, block := range blocks {
		intersecting(trie, block)
	}
}

func intersecting(trie ipaddr.AddressTrie, cidr *ipaddr.IPAddress) {
	intersecting := make([]*ipaddr.IPAddress, 0, trie.Size())

	addr := cidr.ToAddressBase() // convert IPAddress to Address
	containingBlocks := trie.ElementsContaining(addr)
	containedBlocks := trie.ElementsContainedBy(addr)

	for block := containingBlocks.ShortestPrefixMatch(); block != nil; block = block.Next() {
		next := block.GetKey().ToIP()
		if !next.Equal(cidr) {
			intersecting = append(intersecting, next)
		}
	}
	iter := containedBlocks.Iterator()
	for block := iter.Next(); block != nil; block = iter.Next() {
		next := block.ToIP()
		if !next.Equal(cidr) {
			intersecting = append(intersecting, next)
		}
	}
	fmt.Printf("CIDR %s intersects with %v\n", cidr, intersecting)
}
