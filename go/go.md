# Golang 



## Table of contents
 - [About](#about)
 - [Key features](#features)
 - [Use Cases](#use-case)
 - [Go Language Notes](#note)
 - [Common notes](#others)


## About <a name = "about"></a>
Golang, often referred to as Go, is a statically typed, compiled programming language designed by Google. It's known for its simplicity, efficiency, and concurrency features, making it a popular choice for various applications.

## Key Features <a name = "features"></a>
 - Simplicity: Golang's syntax is clean and easy to learn, reducing the complexity of writing and maintaining code.
 - Efficiency: It's a compiled language, resulting in fast execution times and efficient memory usage.
 - Concurrency: Golang provides built-in support for concurrency through goroutines and channels, making it well-suited for handling multiple tasks simultaneously.
 - Static Typing: Ensures type safety and helps catch errors early in the development process.
 - Garbage Collection: Automatically manages memory allocation and deallocation, reducing the burden on developers.
 - Cross-Platform Compatibility: Golang can be compiled for various operating systems, making it highly portable.
 - Standard Library: Offers a rich set of built-in packages for common tasks like networking, I/O, and data structures.

## Use Cases <a name = "use-case"></a>
 - Web Development: Building web applications and services.
 - System Programming: Creating operating systems, network tools, and command-line utilities.
 - Cloud Computing: Developing cloud-native applications and microservices.
 - Data Science: Analyzing and processing large datasets.
 - DevOps: Automating tasks and managing infrastructure.

## Go Project Structure and basics 
 - [What can you build in Golang?!](https://www.youtube.com/watch?v=4fjNO9CuqVs)
 - [Golang standard project layout (not official )](https://github.com/golang-standards/project-layout)
 - [This Is The BEST Way To Structure Your GO Projects](https://www.youtube.com/watch?v=dxPakeBsgl4)
    - [readme](https://github.com/Melkeydev/go-blueprint?tab=readme-ov-file)
    - [blueprint](https://go-blueprint.dev/)
 - Mod
   
   Mod files in Go are configuration files that specify the dependencies and version requirements of a Go project.
   They were introduced in Go 1.11 to simplify dependency management and make it more reliable.

        
   ###### 1. go mod init
        Initializes a new Go module and creates a go.mod file in the current directory.

        Example: go mod init example.com/yourproject
   
        Functionality:
            1. Sets the module path in the go.mod file.
            2. Adds any direct dependencies to the require section of the go.mod file.

   
    ###### 2. go mod tidy
         Ensures that the go.mod file and the vendor directory are consistent with the project's dependencies.

         Example: go mod tidy

         Functionality:
            1. Adds missing dependencies to the go.mod file.
            2. Removes unnecessary dependencies from the go.mod file.
            3. Updates dependency versions to the latest compatible ones.
            4. Updates the vendor directory to match the dependencies in the go.mod file.

   ###### 3. go mod vendor
          Copies all dependencies into a vendor directory within the projec

          Example: go mod vendor
   
          Functionality:
             1. Creates a vendor directory at the project's root.
             2. Copies all dependencies listed in the go.mod file into the vendor directory.
             3. This can be useful for offline builds or specific dependency management strategies.


    - [go module dependency](https://www.youtube.com/watch?v=5VKZzVNKodk)
    - [A long video on MOD in golang](https://www.youtube.com/watch?v=O8uUGEobo-Q)
- [Golang Rune - Fully Understanding Runes in Go](https://www.youtube.com/watch?v=7isCXLWPTqI&list=PLve39GJ2D71wKL33k5eZ6Frot74mhiCxz)
- [Go Workspace & Runtime Explained in 5 Minutes](https://www.youtube.com/watch?v=k8LClK96NZ4&list=PLve39GJ2D71wKL33k5eZ6Frot74mhiCxz&index=5)
- [Go Environment Variables Explained in 5 Minutes](https://www.youtube.com/watch?v=Ut-NLq6d694&list=PLve39GJ2D71wKL33k5eZ6Frot74mhiCxz&index=6)
  

## Go Language Notes <a name = "notes"></a>
 - [Intergaces](features/Interfaces.md)
 - [Goroutines](features/Goroutines.md)
 - [Channels and Concurrency](features/Channels%20and%20Concurrency.md)

## Others <a name = "others"></a>
 - [Closures](https://www.youtube.com/watch?v=jHd0FczIjAE&list=PL7g1jYj15RUMMCMDYPyZHN3CaWbt3Rl5y&index=1)
 - [Abstraction](https://www.youtube.com/watch?v=CRY4_-p5FgM&list=PL7g1jYj15RUMMCMDYPyZHN3CaWbt3Rl5y&index=5)
 - Context
    - [Advanced Golang: Channels, Context and Interfaces Explained](https://www.youtube.com/watch?v=VkGQFFl66X4)
    - [How To Use The Context Package In Golang?](https://www.youtube.com/watch?v=kaZOXRqFPCw)
 - Pointers
    - [Go Pointers: When & How To Use Them Efficiently](https://www.youtube.com/watch?v=3WsEDZRif6U)
 - Slice
    - [Golang Slice Tricks Every Beginner Should Know](https://www.youtube.com/watch?v=AL_C9nF_0ss)
  
 - [Important Tips On How To Write Idiomatic Code In Golang](https://www.youtube.com/watch?v=9cJHCoSxbn8)
 - [How To Build And Structure A Microservice In Golang?!](https://www.youtube.com/watch?v=sqj4UzN4OpU)
 - [Don't Make this Golang Beginner Mistake!](https://www.youtube.com/watch?v=M9h6KGFRRwE)
 - Design Pattern
      - [Golang Pattern](https://www.youtube.com/watch?v=pNv0dXBqKsM&list=PLZ6pHHPq5DrnH7XK5NXAXbCFeHfQzIauk)
      - [Go Design Patterns](https://www.youtube.com/watch?v=F365lY5ECGY&list=PLJbE2Yu2zumAKLbWO3E2vKXDlQ8LT_R28)








