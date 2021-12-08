package main

import "fmt"

const englishHelloPrefix = "Hello, "

/*Constants should improve performance of your application as it saves you creating the "Hello, " string instance every time Hello is called.
To be clear, the performance boost is incredibly negligible for this example! But it's worth thinking about creating constants to capture the
meaning of values and sometimes to aid performance.
CHECK MORE ABOUT THE variables location in memory and how go works
*/

//we can switch statemetns like other languages to avoid lot of if statments
func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}
	if language == "Spanish" {
		return "Hola, " + name
	}
	if language == "French" {
		return "Bonjour, " + name
	}
	return englishHelloPrefix + name
}
func main() {
	fmt.Println(Hello("Manish", ""))
}

/*some new syntex in go
func greetingPrefix(language string) (prefix string) {
	prefix := "abcd"
	return
}
the thing to notice here we have set the return of the function
as named one i.e. prefix so now whatever the value of prefix changes during the course
of function it will get return at the end and you don't need to explicitly write `return prefix`
it will always return `prefix` just with mentioning prefix
*/
