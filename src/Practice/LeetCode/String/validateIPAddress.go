// 468. Validate IP Address
// https://leetcode.com/problems/validate-ip-address/

func validIPAddress(queryIP string) string {
    ans1 := "IPv4"
    ans2 := "IPv6"
    ans3 := "Neither"

    if IsIPv6(queryIP) {
        return ans2
    } else if IsIPv4(queryIP) {
        return ans1
    } else {
        return ans3
    }
}

func IsIPv6(str string) bool {
    ip := strings.Split(str, ":")
    if len(ip) != 8 {
        return false
    }
    for _, e := range ip {
        if len(e) > 4 || len(e) < 1 {
            return false
        }
        for _, j := range e {
            if !unicode.IsDigit(j) && !unicode.IsLetter(j) {
                return false
            }
            if unicode.IsLetter(j) && (unicode.ToLower(j) < rune('a') || unicode.ToLower(j) > rune('f')) {
                return false
            }
        }
    }
    return true
}

func IsIPv4(str string) bool {
    ip := strings.Split(str, ".")
    if len(ip) != 4 {
        return false
    }
    for _, e := range ip {
        if len(e) > 1 && e[0] == '0' {
            return false
        }
        num, err := strconv.Atoi(e) 
        if err != nil || num < 0 || num > 255 {
            return false
        }
    }
    return true
}
