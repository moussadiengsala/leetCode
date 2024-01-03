package hashtable

func LengthOfLongestSubstring(s string) int {
	substr := make(map[rune]int)
	maxLength := 0
	start := 0

	for i, value := range s {
		if lastIndex, found := substr[value]; found && lastIndex >= start {
			start = lastIndex + 1
		}

		

		substr[value] = i
		currentLength := i - start + 1

		if currentLength > maxLength {
			maxLength = currentLength
		}
	}

	return maxLength
}

/*


// Here start is an index
if lastIndex, found := substr[value]; found && lastIndex >= start {
	start = lastIndex + 1
}
the goal of that is move start at the index of the character which follow the character that we already encountered in the string.

Suppose you have the string "abcb".

Iteration 1 (a):

substr is empty, so you add 'a' to it.
start remains 0.
Iteration 2 (b):

'b' is not in substr, so you add it to the substring.
start remains 0.
Iteration 3 (c):

'c' is not in substr, so you add it to the substring.
start remains 0.
Iteration 4 (b):

'b' is already in substr.
The condition lastIndex >= start checks if the last occurrence of 'b' is after or at the current start index.
In this case, the last occurrence of 'b' is at index 1, and start is currently 0.
To avoid repeating characters, you update start to be the next index after the last occurrence of 'b', so start = lastIndex + 1, which is 2.
The current substring becomes "cb".



Consider the string s := "abcabcbb"

Now, let's walk through the logic with the expression currentLength := i - start + 1:

** First Iteration (i = 0, start = 0):
currentLength := 0 - 0 + 1 => currentLength := 1
At the first character 'a', the current substring is "a" with a length of 1.
currentLength > maxLength => maxLength = 1

** Second Iteration (i = 1, start = 0):
currentLength := 1 - 0 + 1 => currentLength := 2
At the second character 'b', the current substring is "ab" with a length of 2.
currentLength > maxLength => maxLength = 2

** Third Iteration (i = 2, start = 0):
currentLength := 2 - 0 + 1 => currentLength := 3
At the third character 'c', the current substring is "abc" with a length of 3.
currentLength > maxLength => maxLength = 3

** Fourth Iteration (i = 3, start = 1):
Now, a repeating character 'a' is encountered. The start index is updated to the next index of 'a' (i.e., start := 1).
currentLength := 3 - 1 + 1 => currentLength := 3
The current substring is now "bca" with a length of 3.
currentLength > maxLength : false => maxLength = 3


** Fifth Iteration (i = 4, start = 2):
Now, a repeating character 'b' is encountered. The start index is updated to the next index of 'b' (i.e., start := 2).
currentLength := 4 - 2 + 1 => currentLength := 3
The current substring is now "cab" with a length of 3.
currentLength > maxLength : false => maxLength = 3


** Sixth Iteration (i = 5, start = 3):
Now, a repeating character 'c' is encountered. The start index is updated to the next index of 'c' (i.e., start := 5).
currentLength := 5 - 3 + 1 => currentLength := 3
The current substring is now "abc" with a length of 3.
currentLength > maxLength : false => maxLength = 3

** Seventh Iteration (i = 6, start = 4):
Now, a repeating character 'b' is encountered again. The start index is updated to the next index of 'b' (i.e., start := 6).
currentLength := 6 - 4 + 1 => currentLength := 3
The current substring is now "bcb" with a length of 3.
currentLength > maxLength : false => maxLength = 3

** Eighth Iteration (i = 7, start = 7):
Now, a repeating character 'b' is encountered again. The start index is updated to the next index of 'b' (i.e., start := 7).
currentLength := 7 - 7 + 1 => currentLength := 1
The current substring is now "b" with a length of 1.
currentLength > maxLength : false => maxLength = 3
*/
