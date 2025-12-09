1. Check if a fast attack can be made
   Define the CanFastAttack() function that takes a boolean value that indicates if the knight is awake. This function returns true if a fast attack can be made based on the state of the knight. Otherwise, returns false:

var knightIsAwake = true
fmt.Println(CanFastAttack(knightIsAwake))
// Output: false
2. Check if the group can be spied upon
   Define the CanSpy() function that takes three boolean values, indicating if the knight, archer and the prisoner, respectively, are awake. The function returns true if the group can be spied upon, based on the state of the three characters. Otherwise, returns false:

var knightIsAwake = false
var archerIsAwake = true
var prisonerIsAwake = false
fmt.Println(CanSpy(knightIsAwake, archerIsAwake, prisonerIsAwake))
// Output: true
3. Check if the prisoner can be signaled
   Define the CanSignalPrisoner() function that takes two boolean values, indicating if the archer and the prisoner, respectively, are awake. The function returns true if the prisoner can be signaled, based on the state of the two characters. Otherwise, returns false:

var archerIsAwake = false
var prisonerIsAwake = true
fmt.Println(CanSignalPrisoner(archerIsAwake, prisonerIsAwake))
// Output: true
4. Check if the prisoner can be freed
   Define the CanFreePrisoner() function that takes four boolean values. The first three parameters indicate if the knight, archer and the prisoner, respectively, are awake. The last parameter indicates if Annalyn's pet dog is present. The function returns true if the prisoner can be freed based on the state of the three characters and Annalyn's pet dog presence. Otherwise, it returns false:

var knightIsAwake = false
var archerIsAwake = true
var prisonerIsAwake = false
var petDogIsPresent = false
fmt.Println(CanFreePrisoner(knightIsAwake, archerIsAwake, prisonerIsAwake, petDogIsPresent))
// Output: false

