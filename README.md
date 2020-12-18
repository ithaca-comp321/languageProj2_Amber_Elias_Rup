# languageProj2_Amber_Elias_Rup

Golang concurrency vs Sequential


Goal:
To use if Golang concurrency is more reliable then sequential 

Problem: 
Created a file with phone numbers and random strings for it to be read in a concurrent way compared to a sequential way and time it to see which one is cost more in terms of sebastian four matrices (readability, writability, reliability and cost) .

How to: 
In order to run a program concurrency golang uses goroutines and channels.

“go” keyword tells us don’t wait for the function to finish; run it in the background and can continue to the next line which is called a goroutine; which will run concurrently. 

When the main go routine finishes, the program will exit unless there is an infinite loop. To fix this problem, we will use a waitgroup. 

Wg.Wait suspends all execution until all goroutines have finished
wg.Done decrements “counter”

channels are a way for goroutines to communicate with each other and allows us to send and receive a message. Channels have types and you can actually send channels through channels
Channels are how we get all the results into a single value.




Conclusion: 

Readability:  If you don’t know the language then the short hand cuts will be difficult to understand unless you read about golang. For example the syntax like the assignment operator := and the blank identifier _  a beginner programmer might not understand. The variable type declaration is strange and doesn’t read as naturally as the conventional type - name (Go is var - name - type)
Otherwise, the language was designed to be fairly simplistic, such as with implicit type checking (which could be positive or negative for readability depending on POV)



Writability: It wasn’t easy to create a program. Concurrency required more lines of code then Sequential. 


 Although it was difficult initially, writing concurrency in Go seems significantly easier/simpler to understand and setup, at least compared to Java.  Referring to the Quora source in the Cost section below - Go was designed for speed of development and you can see that in our code and the goroutine/channel overhead.
Additionally (mentioned in the readability section), Go has options for implicit type checking (and the assignment operator “:=” ) which can speed up writing and testing code.
Writing simple programs in Go will take more time and lines of code compared to python 

Reliability: 

Time to read file: 465.119µs 
Time to process file with goroutines: 2.208016ms 
Time to process file sequentially: 95.131µs 
Time to read large file: 8.9442ms 
Time to process large file with goroutines: 654.523099ms 
Time to process large file sequentially: 721.212979ms 

LANGUAGE SCALES TO LARGE APPLICATIONS WELL



Cons because when running concurrency read file vs sequential, concurrency would be fast most times with 500-1000 workers, and other times sequential would be fast with the same amount of worker. 

We found that lower than 500 workers would alternate who would be faster in terms of reading two files with  3000 phone numbers and random strings.    

Errors can be difficult to figure out from console error output alone (when it crashes due to an error it’s referred to as “panic”)

Cost: Golang is faster than Java, Python, PHP and many other languages. Go is garbage collected and also has support for pointers.

For large projects where this may be an issue, Go apparently has a faster compiling time than C++ and other popular languages. "Go doesn't rely on a virtual machine to compile its code. It gets compiled directly into a binary file."



Resources: 

http://www.golangbootcamp.com/book/concurrency
https://gobyexample.com/


Examples on how to use Maps, Closures, Timers:  https://gobyexample.com 
https://tour.golang.org/concurrency/1


Concurrent Web crawler:
https://itnext.io/create-your-first-web-scraper-in-go-with-goquery-2dcd45743165
https://gobyexample.com/http-clients

Cost:
This is just random googling about how fast Go is
https://benchmarksgame-team.pages.debian.net/benchmarksgame/fastest/go.html
https://www.quora.com/Is-the-performance-of-Go-comparable-to-C++
https://www.quora.com/How-fast-is-Golang-compared-to-other-programming-languages



