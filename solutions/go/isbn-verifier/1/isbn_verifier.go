package isbn
import(
    "unicode"
    "strings"
)
func IsValidISBN(isbn string) bool {
	isbn= strings.ReplaceAll(isbn, "-","")
    var digit int
    sum:= 0
    step:= 10
    if len(isbn) != 10{
        return false
    }else{
        for i:=0; i<10; i++{
            if i == 9 && (isbn[i] == 'X' || isbn[i] == 'x'){
                digit =10
            }else{
            	if !unicode.IsDigit(rune(isbn[i])){
                	return false
            	}
            	digit= int(isbn[i] - '0')
			}
            sum+= digit * step
            step--
		}
	}
    return sum % 11 == 0
}
