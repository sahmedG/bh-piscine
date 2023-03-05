package piscine

func AlphaCount(s string) int {
    count := 0
    if len (s) < 0 {
        return count
    }
    for i := 0; i <= len(s); i++ {
        if !(s[i] >= 'a' && s[i] <= 'z') {
            continue
        } else {
            count++
        }
    }
    return count
}
