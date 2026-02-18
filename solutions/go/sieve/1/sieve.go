package sieve

func Sieve(limit int) []int {
	if limit < 2{
        return nil
    }
    marked:= make(map[int]bool, limit)
    prime:= make([]int, 0)
    for i:=2; i<= limit; i++{
        if !marked[i]{
            prime= append(prime, i)
            for j:= i*i; j<= limit; j+=i{
                marked[j]= true
            }
        }
    }
    return prime
}
