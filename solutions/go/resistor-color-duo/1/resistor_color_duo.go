package resistorcolorduo

// Value should retu7rn the resistance value of a resistor with a given colors.
func Value(colors []string) int {
	r:= map[string]int{
        "black": 0,
"brown": 1,
"red": 2,
"orange": 3,
"yellow": 4,
"green": 5,
"blue": 6,
"violet": 7,
"grey": 8,
"white": 9,
    }
    if len(colors) >=2{
    return r[colors[0]] * 10 + r[colors[1]]
        }
    return 0
}
