package piscine

const (
	Space = ' '
	Minus = '-'
    Plus = '+'
	Start = '0'
	End = '9'
    MaxInt = int(^uint32(0) >> 1)
)

func myAtoi(s string) int {
	var res int
	var neg bool
    var intStarted bool
	runes := []rune(s)
	for i:=0; i<len(runes); i++ {
		r := runes[i]
		if !intStarted && r == Space {
			continue
		}
        if !intStarted && r == Plus {
            intStarted = true
            continue
        }
		if !intStarted && r == Minus {
			intStarted = true
            neg = true
			continue
		}
		if r < Start || r > End  {
			break
		}

        intStarted = true
		res = res * 10 + int(r - Start)

        if res > MaxInt {
            if neg {
                return -MaxInt-1
            }
            return MaxInt
        }
	}

	if neg {
		return -res
	}
	return res
		
}