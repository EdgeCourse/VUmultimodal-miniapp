# VUmultimodal-miniapp

In this application, we write Go code, create Go packages and modules, write Go unit tests with regular expressions and bench tests, and compile and install the Go application. You can start with a directory that you name github.com

1. Create a greetings directory for your Go module source code.

mkdir greetings

cd greetings

2. Start your module using the go mod init command.

Run the go mod init command, giving it your module path -- here, use github.com/greetings. If you publish a module, this must be a path from which your module can be downloaded by Go tools. That would be your code's repository.

$ go mod init github.com/greetings

Output:
go: creating new go.mod: module github.com/greetings

The go mod init command creates a go.mod file to track your code's dependencies. So far, the file includes only the name of your module and the Go version your code supports. But as you add dependencies, the go.mod file will list the versions your code depends on. This keeps builds reproducible and gives you direct control over which module versions to use.

3. In your text editor, create a file for your code and call it greetings.go. Put code into your greetings.go file, for example, this, and save the file.

package greetings

import "fmt"

// Hello returns a greeting for the named person.
func Hello(name string) string {
    // Return a greeting that embeds the name in a message.
    message := fmt.Sprintf("Hi, %v. Welcome!", name)
    return message
}

This is the first code for your module. It returns a greeting to any caller that asks for one. You'll write code that calls this function in the next step.

In this code, you:

-Declare a greetings package to collect related functions.
-Implement a Hello function to return the greeting. The function takes a name parameter whose type is string. The function also returns a string. In Go, a function whose name starts with a capital letter can be called by a function not in the same package. This is known in Go as an exported name. 
-Declare a message variable to hold your greeting. In Go, the := operator is a shortcut for declaring and initializing a variable in one line (Go uses the value on the right to determine the variable's type). 
-Use the fmt package's Sprintf function to create a greeting message. The first argument is a format string, and Sprintf substitutes the name parameter's value for the %v format verb. Inserting the value of the name parameter completes the greeting text.
-Return the formatted greeting text to the caller.
-Call your code from another module. In the previous section, you created a greetings module. In this section, you'll write code to make calls to the Hello function in the module you just wrote. You'll write code you can execute as an application, and which calls code in the greetings module.

4. Create a hello directory for your Go module source code. This is where you'll write your caller.
After you create this directory, you should have both a hello and a greetings directory at the same level in the hierarchy, like so:

<home>/
 |-- greetings/
 |-- hello/

For example, if your command prompt is in the greetings directory, you could use the following commands:

cd ..
mkdir hello
cd hello

5. Enable dependency tracking for the code you're about to write: run the go mod init command, giving it the name of the module your code will be in. Let's use github.com/hello for the module path.

$ go mod init github.com/hello

Output:
go: creating new go.mod: module github.com/hello

Note: we will run go mod tidy later

6. In your text editor, in the hello directory, create a file in which to write your code and call it hello.go. Write code to call the Hello function, then print the function's return value, for example:

package main

import (
    "fmt"

    "github.com/greetings"
)

func main() {
    // Get a greeting message and print it.
    message := greetings.Hello("Trent")
    fmt.Println(message)
}

In this code, we:

-Declare a main package. 
-Import two packages: github.com/greetings and the fmt package. This gives your code access to functions in those packages. Importing github.com/greetings (the package contained in the module you created earlier) gives you access to the Hello function. You also import fmt, with functions for handling input and output text (such as printing text to the console).
-Get a greeting by calling the greetings package’s Hello function.
-Edit the github.com/hello module to use your local github.com/greetings module.

For production use, you’d publish the github.com/greetings module from its repository (with a module path that reflected its published location), where Go tools could find it to download it. For now, because you haven't published the module yet, you need to adapt the github.com/hello module so it can find the github.com/greetings code on your local file system. To do that, use the go mod edit command to edit the github.com/hello module to redirect Go tools from its module path (where the module isn't) to the local directory (where it is).

7. From the command prompt in the hello directory, run the following command:
$ go mod edit -replace github.com/greetings=../greetings

The command specifies that github.com/greetings should be replaced with ../greetings for the purpose of locating the dependency. After you run the command, the go.mod file in the hello directory should include a replace directive:

module github.com/hello

go 1.18

replace github.com/greetings => ../greetings

8. From the command prompt in the hello directory, run the go mod tidy command to synchronize the github.com/hello module's dependencies, adding those required by the code, but not yet tracked in the module.

$ go mod tidy

Output:
go: found github.com/greetings in github.com/greetings v0.0.0-00010101000000-000000000000
After the command completes, the github.com/hello module's go.mod file should look like this:

module github.com/hello

go 1.18

replace github.com/greetings => ../greetings

require github.com/greetings v0.0.0-00010101000000-000000000000

The command found the local code in the greetings directory, then added a require directive to specify that github.com/hello requires github.com/greetings. You created this dependency when you imported the greetings package in hello.go.

The number following the module path is a pseudo-version number -- a generated number used in place of a semantic version number (which the module doesn't have yet).

To reference a published module, a go.mod file would typically omit the replace directive and use a require directive with a tagged version number at the end.

require github.com/greetings v1.1.0


9. At the command prompt in the hello directory, run your code to confirm that it works. What would happen if the function were not capitalized?
$ go run .

Output:
Hi, Trent. Welcome!

You now have two functioning modules.


10: Return a random greeting and add error checking: change the code so that instead of returning a single greeting every time, it returns one of several predefined greeting messages. Use a Go slice. A slice is like an array, except that its size changes dynamically as you add and remove items. The slice is one of Go's most useful types. Add a small slice to contain three greeting messages, then have your code return one of the messages randomly.

In greetings/greetings.go, change your code to:

package greetings

import (
    "errors"
    "fmt"
    "math/rand"
    "time"
)

// Hello returns a greeting for the named person.

func Hello(name string) (string, error) {
    // If no name was given, return an error with a message.
    if name == "" {
        return name, errors.New("empty name")
    }
    // Create a message using a random format.
    message := fmt.Sprintf(randomFormat(), name)
    return message, nil
}

// init sets initial values for variables used in the function.
func init() {
    rand.Seed(time.Now().UnixNano())
}

// randomFormat returns one of a set of greeting messages. The returned
// message is selected at random.
func randomFormat() string {
    // A slice of message formats.
    formats := []string{
        "Hi, %v. Welcome!",
        "Great to see you, %v!",
        "Hey, %v! Welcome!",
    }

    // Return a randomly selected message format by specifying
    // a random index for the slice of formats.
    return formats[rand.Intn(len(formats))]
}

In this code, you:

-Add a randomFormat function that returns a randomly selected format for a greeting message. Note that randomFormat starts with a lowercase letter, making it accessible only to code in its own package (in other words, it's not exported). In randomFormat, declare a formats slice with three message formats. When declaring a slice, you omit its size in the brackets, like this: []string. This tells Go that the size of the array underlying the slice can be dynamically changed.
Use the math/rand package to generate a random number for selecting an item from the slice.
-Add an init function to seed the rand package with the current time. Go executes init functions automatically at program startup, after global variables have been initialized. For more about init functions, see Effective Go.
-In Hello, call the randomFormat function to get a format for the message you'll return, then use the format and name value together to create the message.
-Return the message (or an error) as you did before.

11. In hello/hello.go, change your code so it looks like the following: 

package main

import (
    "fmt"
    "log"

    "github.com/greetings"
)

func main() {
    // Set properties of the predefined Logger, including
    // the log entry prefix and a flag to disable printing
    // the time, source file, and line number.
    log.SetPrefix("greetings: ")
    log.SetFlags(0)

    // Request a greeting message.
    message, err := greetings.Hello("Trent")
    // If an error was returned, print it to the console and
    // exit the program.
    if err != nil {
        log.Fatal(err)
    }

    // If no error was returned, print the returned message
    // to the console.
    fmt.Println(message)
}

12. At the command line, in the hello directory, run hello.go to confirm that the code works. Run it multiple times, noticing that the greeting changes.
$ go run .

Sample output:
Great to see you, Trent!

$ go run .

Sample output: 
Hi, Trent. Welcome!

$ go run .

Sample output:
Hail, Trent! Welcome!

13. Return greetings for multiple people: you'll handle a multiple-value input, then pair values in that input with a multiple-value output. To do this, you'll need to pass a set of names to a function that can return a greeting for each of them. Changing the Hello function's parameter from a single name to a set of names would change the function's signature. If you had already published the github.com/greetings module and users had already written code calling Hello, that change would break their programs.

In this situation, a better choice is to write a new function with a different name. The new function will take multiple parameters. That preserves the old function for backward compatibility.

14. In greetings/greetings.go, change your code so it looks like the following.

package greetings

import (
    "errors"
    "fmt"
    "math/rand"
    "time"
)

// Hello returns a greeting for the named person.
func Hello(name string) (string, error) {
    // If no name was given, return an error with a message.
    if name == "" {
        return name, errors.New("empty name")
    }
    // Create a message using a random format.
    message := fmt.Sprintf(randomFormat(), name)
    return message, nil
}

// Hellos returns a map that associates each of the named people
// with a greeting message.
func Hellos(names []string) (map[string]string, error) {
    // A map to associate names with messages.
    messages := make(map[string]string)
    // Loop through the received slice of names, calling
    // the Hello function to get a message for each name.
    for _, name := range names {
        message, err := Hello(name)
        if err != nil {
            return nil, err
        }
        // In the map, associate the retrieved message with
        // the name.
        messages[name] = message
    }
    return messages, nil
}

// Init sets initial values for variables used in the function.
func init() {
    rand.Seed(time.Now().UnixNano())
}

// randomFormat returns one of a set of greeting messages. The returned
// message is selected at random.
func randomFormat() string {
    // A slice of message formats.
    formats := []string{
        "Hi, %v. Welcome!",
        "Great to see you, %v!",
        "Hail, %v! Welcome!",
    }

    // Return one of the message formats selected at random.
    return formats[rand.Intn(len(formats))]
}

In this code, you:

-Add a Hellos function whose parameter is a slice of names rather than a single name. Also, you change one of its return types from a string to a map so you can return names mapped to greeting messages.
-Have the new Hellos function call the existing Hello function. This helps reduce duplication while also leaving both functions in place.
-Create a messages map to associate each of the received names (as a key) with a generated message (as a value). In Go, you initialize a map with the following syntax: make(map[key-type]value-type). -Have the Hellos function return this map to the caller. For more about maps, see Go maps in action on the Go blog.
-Loop through the names your function received, checking that each has a non-empty value, then associate a message with each. In this for loop, range returns two values: the index of the current item in the loop and a copy of the item's value. You don't need the index, so you use the Go blank identifier (an underscore) to ignore it. For more, see The blank identifier in Effective Go.

15. In your hello/hello.go calling code, pass a slice of names, then print the contents of the names/messages map you get back. In hello.go, change your code so it looks like the following.

package main

import (
    "fmt"
    "log"

    "github.com/greetings"
)

func main() {
    // Set properties of the predefined Logger, including
    // the log entry prefix and a flag to disable printing
    // the time, source file, and line number.
    log.SetPrefix("greetings: ")
    log.SetFlags(0)

    // A slice of names.
    names := []string{"Trent", "Sadio", "Bobby"}

    // Request greeting messages for the names.
    messages, err := greetings.Hellos(names)
    if err != nil {
        log.Fatal(err)
    }
    // If no error was returned, print the returned map of
    // messages to the console.
    fmt.Println(messages)
}

With these changes, you:

-Create a names variable as a slice type holding three names.
-Pass the names variable as the argument to the Hellos function.

16. At the command line, change to the directory that contains hello/hello.go, then use go run to confirm that the code works. The output should be a string representation of the map associating names with messages, something like the following:

$ go run .

Output:
map[Bobby:Hail, Bobby! Welcome! Trent:Hi, Trent. Welcome! Sadio:Hail, Sadio! Welcome!]


17. Add a test. Testing your code during development can expose bugs that find their way in as you make changes. Let's add a test for the Hello function. Using Go's built-in support for unit testing with Go's testing package and the go test command, you can quickly write and execute tests.

In the greetings directory, create a file called greetings_test.go. Ending a file's name with _test.go tells the go test command that this file contains test functions.

In greetings_test.go, put in the following code and save the file.

package greetings

import (
    "testing"
    "regexp"
)

// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value.
func TestHelloName(t *testing.T) {
    name := "Trent"
    want := regexp.MustCompile(`\b`+name+`\b`)
    msg, err := Hello("Trent")
    if !want.MatchString(msg) || err != nil {
        t.Fatalf(`Hello("Trent") = %q, %v, want match for %#q, nil`, msg, err, want)
    }
}

// TestHelloEmpty calls greetings.Hello with an empty string,
// checking for an error.
func TestHelloEmpty(t *testing.T) {
    msg, err := Hello("")
    if msg != "" || err == nil {
        t.Fatalf(`Hello("") = %q, %v, want "", error`, msg, err)
    }
}

In this code, you:

-Implement test functions in the same package as the code you're testing.
-Create two test functions to test the greetings.Hello function. Test function names have the form TestName, where Name says something about the specific test. Also, test functions take a pointer to the testing package's testing.T type as a parameter. You use this parameter's methods for reporting and logging from your test.
-Implement two tests: TestHelloName calls the Hello function, passing a name value with which the function should be able to return a valid response message. If the call returns an error or an unexpected response message (one that doesn't include the name you passed in), you use the t parameter's Fatalf method to print a message to the console and end execution.
TestHelloEmpty calls the Hello function with an empty string. This test is designed to confirm that your error handling works. If the call returns a non-empty string or no error, you use the t parameter's Fatalf method to print a message to the console and end execution.

At the command line in the greetings directory, run the go test command to execute the test.
The go test command executes test functions (whose names begin with Test) in test files (whose names end with _test.go). You can add the -v flag to get verbose output that lists all of the tests and their results.

The tests should pass.

$ go test
PASS
ok      github.com/greetings   0.364s

$ go test -v
=== RUN   TestHelloName
--- PASS: TestHelloName (0.00s)
=== RUN   TestHelloEmpty
--- PASS: TestHelloEmpty (0.00s)
PASS
ok      github.com/greetings   0.372s

18. Break the greetings.Hello function to view a failing test: the TestHelloName test function checks the return value for the name you specified as a Hello function parameter. To view a failing test result, change the greetings.Hello function so that it no longer includes the name.

In greetings/greetings.go, paste the following code in place of the Hello function. Note that the highlighted lines change the value that the function returns, as if the name argument had been accidentally removed.

// Hello returns a greeting for the named person.
func Hello(name string) (string, error) {
    // If no name was given, return an error with a message.
    if name == "" {
        return name, errors.New("empty name")
    }
    // Create a message using a random format.
    // message := fmt.Sprintf(randomFormat(), name)
    message := fmt.Sprint(randomFormat())
    return message, nil
}

19. At the command line in the greetings directory, run go test to execute the test. This time, run go test without the -v flag. The output will include results for only the tests that failed, which can be useful when you have a lot of tests. The TestHelloName test should fail -- TestHelloEmpty still passes.

$ go test
--- FAIL: TestHelloName (0.00s)
    greetings_test.go:15: Hello("Trent") = "Hail, %v! Welcome!", <nil>, want match for `\bTrent\b`, nil
FAIL
exit status 1
FAIL    github.com/greetings   0.182s

20.  Compile and install the application: while the go run command is a useful shortcut for compiling and running a program when you're making frequent changes, it doesn't generate a binary executable.

The go build command compiles the packages, along with their dependencies, but it doesn't install the results.

The go install command compiles and installs the packages.

From the command line in the hello directory, run the go build command to compile the code into an executable.

$ go build

From the command line in the hello directory, run the new hello executable to confirm that the code works. Note that your result might differ depending on whether you changed your greetings.go code after testing it.

On Linux or Mac:
$ ./hello
map[Bobby:Great to see you, Bobby! Trent:Hail, Trent! Welcome! Sadio:Hail, Sadio! Welcome!]

On Windows:
$ hello.exe
map[Bobby:Great to see you, Bobby! Trent:Hail, Trent! Welcome! Sadio:Hail, Sadio! Welcome!]
You've compiled the application into an executable so you can run it. But to run it currently, your prompt needs either to be in the executable's directory, or to specify the executable's path.

21. Install the executable so you can run it without specifying its path. Discover the Go install path, where the go command will install the current package. You can discover the install path by running the go list command, as in the following example:

$ go list -f '{{.Target}}'

For example, the command's output might say /home/gopher/bin/hello, meaning that binaries are installed to /home/gopher/bin. You'll need this install directory in the next step.

22. Add the Go install directory to your system's shell path so you can run your program's executable without specifying where the executable is.

On Linux or Mac, run the following command:
$ export PATH=$PATH:/path/to/your/install/directory

On Windows, run the following command:
$ set PATH=%PATH%;C:\path\to\your\install\directory

As an alternative, if you already have a directory like $HOME/bin in your shell path and you'd like to install your Go programs there, you can change the install target by setting the GOBIN variable using the go env command:

$ go env -w GOBIN=/path/to/your/bin

or

$ go env -w GOBIN=C:\path\to\your\bin

Once you've updated the shell path, run the go install command to compile and install the package.
$ go install

Run your application by simply typing its name. To make this interesting, open a new command prompt and run the hello executable name in some other directory.

$ hello
map[Bobby:Hail, Bobby! Welcome! Trent:Great to see you, Trent! Sadio:Hail, Sadio! Welcome!]
