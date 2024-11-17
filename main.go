package main

import (
	"errors"
	"fmt"
	"regexp"
	// "sort"
	"strings"
	"time"
)

//q1
func resizeSlice(slice []int, targetLen int)[]int{
	if targetLen < 0{
		return []int{}
	}
	if len(slice) < targetLen{
		for i :=len(slice); i < targetLen; i++{
			slice = append(slice, 0)
		}
	}else{
		slice = slice[:targetLen]
	}
	return slice 

}

//q2
func WordFrequency(words []string) map[string]int {

	frequency := make(map[string]int)

	// Regular expression to match words, ignoring punctuation
	regex := regexp.MustCompile(`[a-zA-Z]+`)
	for _, word := range words {
		word = strings.ToLower(word)
		wordsInString := regex.FindAllString(word, -1)
		for _, w := range wordsInString {
			frequency[w]++
		}
	}
	return frequency
}

//q3
func findEmployeeAge(employee map[string]int, name string)(int, error){
	age, exists :=employee[name]
	if !exists{
		return 0, fmt.Errorf(" %s is not an employee", name)
	}
	return age, nil
	
}


//q4

func removeDuplicates(nums []int) []int {
	dup := make(map[int]bool)
	var result []int
	for _, num := range nums {
		if !dup[num] {
			result = append(result, num)
			dup[num] = true
		}
	}

	// Return the slice without duplicates
	return result
}

// q 5 partition even and odd 
func partitionEvenOdd(nums []int)([]int, []int){
	evenNum := []int{}
	oddNum := []int{}
	for _, num := range nums{
		if num % 2 == 0{
			evenNum = append(evenNum, num)
		}else{
			oddNum = append(oddNum, num)
		}
	}
	return evenNum, oddNum
}

// q6 merge and sort



func mergeAndSort(slice1 []int, slice2 []int)( []int , error){
	if checkDuplicate(slice1) || checkDuplicate(slice2){
		return nil, errors.New("slice contains duplicate elements")
	}

	if !isSorted(slice1) || !isSorted(slice2) {
		return nil, errors.New("slice is not sorted, u should sort them and try again")
	}

	slice3 := append(slice1, slice2...)
	// sort.Ints(slice3)
	return slice3, nil

}
// check for duplicate
func checkDuplicate(slice []int)(bool){
	dup := make(map[int]bool)
	for _, num := range slice{
		if dup[num]{
			return true
		}
		dup[num] = true
	}
	return false
}
// check if the inputs are sorted or not
func isSorted(slice []int) bool {
	for i := 1; i < len(slice); i++ {
		if slice[i] < slice[i-1] {
			return false 
		}
	}
	return true 
}

// q7

type MultiKeyMap map[string]map[string]interface{}
func (m MultiKeyMap) Set(key1, key2 string, value interface{}) {
	if _, exists := m[key1]; !exists {
		m[key1] = make(map[string]interface{})
	}
	m[key1][key2] = value
}

// Get retrieves a value from the MultiKeyMap using two keys
func (m MultiKeyMap) Get(key1, key2 string) (interface{}, error) {
	if nestedMap, exists := m[key1]; exists {
		// Check if the second key exists in the nested map
		if value, exists := nestedMap[key2]; exists {
			return value, nil
		}
	}
	return nil, errors.New("key combination not found")
}

// Delete removes a value from the MultiKeyMap using two keys
func (m MultiKeyMap) Delete(key1, key2 string) error {
	
	if nestedMap, exists := m[key1]; exists {
		// Check if the second key exists in the nested map
		if _, exists := nestedMap[key2]; exists {
			// Delete the second key
			delete(nestedMap, key2)
			// If the nested map is now empty, delete the first key
			if len(nestedMap) == 0 {
				delete(m, key1)
			}
			return nil
		}
	}
	return errors.New("key combination not found")

}

//q8

func WindowSum(nums []int, k int)([]int, error){
	if k < 1 || k > len(nums) {
		return nil, errors.New("invalid window size")
	}

	var result []int
	windowSum := 0

	// Initialize the first window
	for i := 0; i < k; i++ {
		windowSum += nums[i]
	}
	result = append(result, windowSum)

	// Slide the window across the slice
	for i := k; i < len(nums); i++ {
		// Subtract the element that goes out of the window
		windowSum -= nums[i-k]
		// Add the element that comes into the window
		windowSum += nums[i]
		result = append(result, windowSum)
	}

	return result, nil
}

// q9
//reverses a given string
func reverseString(s string) string {
	bytes := []byte(s) 
    for i, j := 0, len(bytes)-1; i < j; i, j = i+1, j-1 {
        bytes[i], bytes[j] = bytes[j], bytes[i] 
    }
    return string(bytes) 
}

// replaceMapKeys reverses each key in the given map in-place
func replaceMapKeys(m map[string]int) {
	for key, value := range m {
		reversedKey := reverseString(key)
		delete(m, key)
		m[reversedKey] = value
	}
}

// q10
type SliceError struct {
	operation string 
	index     int    
	message   string 
}

// error method 
func (e *SliceError) Error() string {
	return fmt.Sprintf("Error during %s at index %d: %s", e.operation, e.index, e.message)
}

func safeRemove(slice []int, index int) ([]int, error) {
	if index < 0 || index >= len(slice) {
		return nil, &SliceError{
			operation: "remove",
			index:     index,
			message:   "index out of bounds",
		}
	}

	return append(slice[:index], slice[index+1:]...), nil
}


// q11
func sumVariadic(nums ...int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum
}


// q12
func IsPalindrome(s string) bool {
	var noPunctuation string
	for _, char := range s {
		if char >= 'a' && char <= 'z' || char >= 'A' && char <= 'Z' {
			noPunctuation += string(char)
		}
	}

	// check the word is a palindrome
	left, right :=len(noPunctuation)-1, 0
	for left < right{
		if noPunctuation[left] != noPunctuation[right]{
			return false
		}
		left++
		right--
	}
	return true
}
// q13
func Swap(num1, num2 *int) {
	*num1, *num2 = *num2, *num1
	
}

// q14
func applyToEach(slice []int, f func(int) int) []int {
	for i, num := range slice {
		slice[i] = f(num)
	}
	return slice
}

//  example functions to apply
func square(x int) int {
	return x * x
}



// q15

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
func GCD(a, b int) int{
	
	if a == 0 && b == 0 {
		panic("GCD of 0 and 0 is undefined")
	}

	
	if b == 0 {
		return abs(a) 
	}
	return GCD(b, a%b)
}

// q 16
func memoizeFib(n int, memo map[int]int) int {
	if val, exists := memo[n]; exists {
		return val
	}
	if n <= 1 {
		memo[n] = n
		return n
	}

	memo[n] = memoizeFib(n-1, memo) + memoizeFib(n-2, memo)
	return memo[n]
}

//q 17

func filterSlice(slice []int, f func(int) bool) []int {
	var filtered []int
	for _, num := range slice {
		if f(num) {
			filtered = append(filtered, num)
		}
	}
	return filtered
}

// example predicate functions
func isEven(n int) bool {
    return n%2 == 0
}


// q 18
func revString(s string) string {
	rune := []rune(s)
	for i, j := 0, len(rune)-1; i < j; i, j = i+1, j-1 {
        rune[i], rune[j] = rune[j], rune[i]
    }

    return string(rune)
}


// q19
func executeWithRecovery(f func()) {
	// Defer a function to recover from any panic that occurs
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
		fmt.Println("program runs after recovered from panic..")
	}()
	f()
}


func retry(fun func()error, retries int) error{
	var err error
	for i := 0; i < retries; i++{
		err = fun()
		if err == nil{
			return nil
		}
		fmt.Printf("Attempt %d failed: %v\n", i+1, err)
	}
	return err
}

func mightFailFunction() error {
	if (time.Now().UnixNano() % 2) == 0 {
		return errors.New("random failure")
	}
	return nil
}

func main(){ 
	// question 1
	fmt.Println(resizeSlice([]int{1, 2, 3}, 5)) 
	fmt.Println(resizeSlice([]int{1, 2, 3, 4, 5}, 3))  

	fmt.Println("-------------------------------------")
	fmt.Println("--------------question 2 ------------")
	fmt.Println("-------------------------------------")
	

	// q 2
	words := []string{"Hello", "world!", "hello", "world.", "go", "go", "Golang!"}
	fmt.Println(WordFrequency(words))

	fmt.Println("-------------------------------------")
	fmt.Println("--------------question 3 ------------")
	fmt.Println("-------------------------------------")

	// q 3
	employee := map[string]int{"sara": 27, "yabi": 28, "joel": 29}
	name := "abel"
	age, err := findEmployeeAge(employee, name)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(age)
	}

	fmt.Println("-------------------------------------")
	fmt.Println("--------------question 4 ------------")
	fmt.Println("-------------------------------------")

	//q4
	numbers := []int{1, 2, 3, 2, 4, 1, 5, 3}
	fmt.Println(removeDuplicates(numbers))

	fmt.Println("-------------------------------------")
	fmt.Println("--------------question 5 ------------")
	fmt.Println("-------------------------------------")
	//q5
	nums :=[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	even, odd := partitionEvenOdd(nums)
	fmt.Println("Even numbers:", even)
	fmt.Println("Odd numbers:", odd)

	fmt.Println("-------------------------------------")
	fmt.Println("--------------question 6 ------------")
	fmt.Println("-------------------------------------")

	// q6
	// trying it with unsorted slice
	slice1 := []int{1, 2, 3,4, 5}
	slice2 := []int{5, 4, 3, 2, 1}
	slice3, err := mergeAndSort(slice1, slice2)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(slice3)
	}

	fmt.Println("-------------------------------------")
	fmt.Println("--------------question 7 ------------")
	fmt.Println("-------------------------------------")
	//q7
	mkm := make(MultiKeyMap)

	// Set values
	mkm.Set("first", "last", 25)
	mkm.Set("foo", "bar", "hello")
	mkm.Set("alpha", "beta", 3.14)

	// Get values
	if value, err := mkm.Get("first", "last"); err == nil {
		fmt.Println("Value for ('first', 'last'):", value)
	} else {
		fmt.Println("Error:", err)
	}

	if value, err := mkm.Get("foo", "bar"); err == nil {
		fmt.Println("Value for ('foo', 'bar'):", value)
	} else {
		fmt.Println("Error:", err)
	}

	// Delete a value
	if err := mkm.Delete("alpha", "beta"); err == nil {
		fmt.Println("Deleted ('alpha', 'beta')")
	} else {
		fmt.Println("Error:", err)
	}

	// Try to get a deleted value
	if value, err := mkm.Get("alpha", "beta"); err == nil {
		fmt.Println("Value for ('alpha', 'beta'):", value)
	} else {
		fmt.Println("Error:", err)
	}

	fmt.Println("-------------------------------------")
	fmt.Println("--------------question 8 ------------")
	fmt.Println("-------------------------------------")

	// q8
	numb := []int{1, 2, 3, 4, 5, 6}
	k := 3

	// Call WindowSum function
	result, err := WindowSum(numb, k)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Window sums:", result)
	}

	// Invalid input
	k = 7
	result, err = WindowSum(numb, k)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Window sums:", result)
	}
	fmt.Println("-------------------------------------")
	fmt.Println("--------------question 9 ------------")
	fmt.Println("-------------------------------------")

	// q9
	m := map[string]int{
		"hello": 10,
		"world": 20,
		"go":    30,
	}

	fmt.Println("Original map:", m)
	replaceMapKeys(m)
	fmt.Println("Modified map:", m)

	fmt.Println("-------------------------------------")
	fmt.Println("--------------question 10------------")
	fmt.Println("-------------------------------------")

	//q10
	//  slice
	x := []int{1, 2, 3, 4, 5}

	// Valid removal
	newSlice, err := safeRemove(x, 2)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Updated slice:", newSlice)
	}

	// Invalid removal
	newSlice, err = safeRemove(x, 10)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("-------------------------------------")
	fmt.Println("--------------question 11------------")
	fmt.Println("-------------------------------------")

	res := sumVariadic(1,2,3,4,5,6,7,8,9,10)
	// sum with individual integers
	fmt.Println("Sum with individual integers:", res)
	// sum with slice of integers
	resSlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	res = sumVariadic(resSlice...)
	fmt.Println("Sum with slice of integers:", res)

	fmt.Println("-------------------------------------")
	fmt.Println("--------------question 12------------")
	fmt.Println("-------------------------------------")

	//q12
	palindrome := "racecar!"
	if IsPalindrome(palindrome) {
		fmt.Println(palindrome, "is a palindrome")
	} else {
		fmt.Println(palindrome, "is not a palindrome")
	}

	fmt.Println("-------------------------------------")
	fmt.Println("--------------question 13------------")
	fmt.Println("-------------------------------------")

	num1,num2 :=4, 7
	Swap(&num1, &num2)
	fmt.Println("the swapped numbers: ", num1, num2)

	fmt.Println("-------------------------------------")
	fmt.Println("--------------question 14------------")
	fmt.Println("-------------------------------------")

	kutir := []int{1, 2, 3, 4, 5}

	squaredSlice := applyToEach(kutir, square)
	fmt.Println("Squared slice:", squaredSlice)



	fmt.Println("-------------------------------------")
	fmt.Println("--------------question 15------------")
	fmt.Println("-------------------------------------")

	fmt.Println("GCD of 48 and 18:", GCD(48, 18))


	fmt.Println("-------------------------------------")
	fmt.Println("--------------question 16------------")
	fmt.Println("-------------------------------------")

	memo := make(map[int]int)

	// Test cases
	fmt.Println("Fib(10):", memoizeFib(10, memo))

	fmt.Println("-------------------------------------")
	fmt.Println("--------------question 17------------")
	fmt.Println("-------------------------------------")

	predicate := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

    // Filter for even numbers.
    evens := filterSlice(predicate, isEven)
    fmt.Println("Even numbers:", evens)

	fmt.Println("-------------------------------------")
	fmt.Println("--------------question 18------------")
	fmt.Println("-------------------------------------")
	
	input := "hey you!"
    rev := revString(input)
    fmt.Println("reversed string:", rev)

	fmt.Println("-------------------------------------")
	fmt.Println("--------------question 19------------")
	fmt.Println("-------------------------------------")


	// function that will causes a panic
	causePanic := func() {
		fmt.Println("am about to panic")
		panic("something went wrong!")
	}

	executeWithRecovery(causePanic)

	fmt.Println("-------------------------------------")
	fmt.Println("--------------question 20------------")
	fmt.Println("-------------------------------------")

	error := retry(mightFailFunction, 3)
	if error != nil {
		fmt.Println("final error:", error)
	} else {
		fmt.Println("function succeeded")
	}

}