package greetings

import(
	"regexp"
	"testing"
)

func TestHelloAmin(t *testing.T){
	name:= "Amin"
	safeName:=  regexp.MustCompile(`\b` + name + `\b`)
	msg, err:= Hello("Amin")
	if !safeName.MatchString(msg) || err != nil{
		t.Errorf(`Hello("Amin") = %q , %v , want match for %#q, nil`, msg, err, safeName)
	}
}

func TestHelloEmpty(t *testing.T){
	msg, err:= Hello("")
	if msg != "" || err == nil{
		t.Errorf(`Hello("") = %q, %v , want "", error`, msg, err)
	}
}