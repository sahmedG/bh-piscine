package piscine

func IsPrime(num int) bool {
    primcount := 0

    for i := 2; i < num; i++ {
        if num%i == 0 {
            primcount++
            break
        }
    }

    if primcount == 0 && num != 1 {
        return true
    } else {
        return false
    }
}
