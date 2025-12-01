package greetings

import(
	"regexp"
	"testing"
)

func TestHelloFunc(t *testing.T){
	name:= "Amin"
	safeName:=  regexp.MustCompile(`\b` + name + `\b`)
	msg, err:= Hello("Amin")
	if !safeName.MatchString(msg) || err != nil{
		t.Errorf(`Hello("Amin") = %q , %v , want match for %#q, nil`, msg, err, safeName)
	}
}