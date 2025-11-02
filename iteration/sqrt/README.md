# implement the square root function
As a simple way to play with functions and loops, implement the square root function using Newton's method.
In this case, Newton's method is to approximate Sqrt(x) by picking a starting point z and then repeating: z - (z*z - x) / (2 * z)
To begin with, just repeat that calculation 10 times and see how close you get to the answer for various values (1, 2, 3, ...).
Next, change the loop condition to stop once the value has stopped changing (or only changes by a very small delta). See if that's more or fewer iterations. How close are you to the math.Sqrt?
Hint: to declare and initialize a floating point value, give it floating point syntax or use a conversion:
	z := float64(1)
	z := 1.0

(Note: If you are interested in the details of the algorithm, the z² − x above is how far away z² is from where it needs to be (x), and the division by 2z is the derivative of z², to scale how much we adjust z by how quickly z² is changing. This general approach is called Newton's method. It works well for many functions but especially well for square root.)
