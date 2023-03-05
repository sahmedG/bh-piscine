package piscine

func AlphaCount(s string) int {
    count := 0
    for i := 0; i <= len(s); i++ {
        if !(byte_str[start] >= 'a' && byte_str[start] <= 'z') {
            continue
        } else {
            count++
        }
    }
    return count
}
