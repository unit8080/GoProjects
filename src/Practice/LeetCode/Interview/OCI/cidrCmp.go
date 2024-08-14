/*
*
Ip List: {134.70.24.0/21, 134.70.25.0/30, 134.70.27.0/28, 147.154.59.128/26, 147.154.60/30, 130.35.203.118/32, 147.154.27.228/32}

Answer =
130.35.203.118/32
134.70.24.0/21
147.154.27.228/32
147.154.59.128/26

Approach 1:

- check the cidr belongs to bigger cidr
- sort the cidrs based on the masklength
- use net package to check if contains (firstIP, lastIP)
ipA,ipnetA,_ := net.ParseCIDR(A)
ipnetA.Contains(ipB)
ip.Mask(ipnet.Mask)
*/
package main

import (
	"fmt"
	"net"
)

// A > B
// 130.70.24.1/32 contains  130.70.24.0/24 - false
// 130.70.24.0/24 contains  130.70.24.1/32 - true
// true : contain
// false : does not contain
func helper(A, B string) bool {
	_, ipnetA, err := net.ParseCIDR(A)
	if err != nil {
		return false
	}
	ipB, _, err1 := net.ParseCIDR(B)
	if err1 != nil {
		return false
	}
	return ipnetA.Contains(ipB)
}

func addNumbers(a uint32, b uint32) uint32 {
	return (a + b)
}

func main() {
	var num1, num2, res uint32
	fmt.Scanf("%v\n%v", &num1, &num2)
	res = addNumbers(num1, num2)
	fmt.Println(res)
	ans := []string{}
	List := []string{"134.70.24.0/21", "134.70.25.0/30", "134.70.27.0/28", "147.154.59.128/26", "147.154.60.0/30", "130.35.203.118/32", "147.154.27.228/32"}
	for i := 0; i < len(List); i++ {
		isPresent := false
		A := List[i]
		for j := 0; j < len(List); j++ {

			B := List[j]
			if A == B {
				continue
			}

			result := helper(A, B)
			// if B is contained then skip "B"
			if !0 {
				ans = append(ans, B)
			} else {
				isPresent = true
			}
		}
		if !isPresent {
			ans = append(ans, A)
		}
	}
	// return ans

	// result := helper("130.70.25.0/24", "130.70.24.1/32")
	// result := helper("130.70.24.1/32", "130.70.24.0/24")
	fmt.Println("Result:", result)
}
