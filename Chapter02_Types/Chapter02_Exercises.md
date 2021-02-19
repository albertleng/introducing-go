1. How are integers stored on a computer?  
*My answer*: Depending on its types, it's stored in memory locations consisting of 1 byte or more.  
**Book's answer**: Integers are stored on a computer by using a base-2, binary number system.

2. We know that (in base 10) the largest one-digit number is 9 and the largest two-digit number is 99. Given that in binary the largest two-digit number is 11 (3), the largest three-digit number is 111 (7) and the largest four-digit number is 1111 (15), what is the largest eight-digit number? (Hint: 10^1-1 = 9 and 10^2-1 = 99)   
*My answer*: 15+2^4+2^5+2^6+2^7=255
**Book's answer**: 


3. Although overpowered for the task, you can use Go as a calculator. Write a program that computes 32,132 x 42,452 and prints it to the terminal (use the * operator for multiplication).  
*My answer*: Refer to [chapter02_exercise_03.go](chapter02_exercise_03.go) 
**Book's answer**:  Same as mine.

4. What is a string? How do you find its length?  
*My answer*: A sequence of characters. using `len(string)`  
**Book's answer**: 

5. What is the value of the expression `(true && false) || (false && true) || !(false && false)`?  
*My answer*: true

