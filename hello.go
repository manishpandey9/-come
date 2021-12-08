package main

import "fmt"

const englishHelloPrefix = "Hello, "

/*Constants should improve performance of your application as it saves you creating the "Hello, " string instance every time Hello is called.
To be clear, the performance boost is incredibly negligible for this example! But it's worth thinking about creating constants to capture the
meaning of values and sometimes to aid performance.
CHECK MORE ABOUT THE variables location in memory and how go works
*/
func Hello(name string) string {
	return englishHelloPrefix + name
}
func main() {
	fmt.Println(Hello("Manish"))
}
