# Tips and Notes


This page offers a curated collection of video resources that provide essential tips and best practices for writing 
clean, maintainable, and readable code in software development. Discover expert insights and practical guidance to 
enhance your coding skills and create high-quality software

### Code readability 
 - [The 3 Laws of Writing Readable Code](https://www.youtube.com/watch?v=-AzSRHiV9Cc)
 - [how NASA writes space-proof code](https://www.youtube.com/watch?v=GWYhtksrmhE)
 - [Naming Things in Code](https://www.youtube.com/watch?v=-J3wNP6u5YU)
 - [My 10 “Clean” Code Principles (Start These Now)](https://www.youtube.com/watch?v=wSDyiEjhp8k)


### System-design 
  - [20 System Design Concepts Explained in 10 Minutes](https://www.youtube.com/watch?v=i53Gi_K3o7I)
  - [System Design Concepts Course and Interview Prep](https://www.youtube.com/watch?v=F2FmTdLtb_4)
  - [10 Most Common System Design Interview Mistakes](https://www.youtube.com/watch?v=15sgUqScHgs)
  - [7 Must-know Strategies to Scale Your Database](https://www.youtube.com/watch?v=_1IKwnbscQU)
  - [Basic System Design for Uber or Lyft | System Design Interview Prep](https://www.youtube.com/watch?v=R_agd5qZ26Y)
  - [Design Twitter - System Design Interview](https://www.youtube.com/watch?v=o5n85GRKuzk)
  - [awesome-system-design-resources](https://github.com/ashishps1/awesome-system-design-resources/tree/main?tab=readme-ov-file)
  - [Books](https://github.com/samayun/devbooks/blob/master/Designing%20Data-Intensive%20Applications%20The%20Big%20Ideas%20Behind%20Reliable%2C%20Scalable%2C%20and%20Maintainable%20Systems%20(%20PDFDrive%20).pdf)
  - [system-design-primer](https://github.com/donnemartin/system-design-primer/blob/master/README.md)
  - [Talks and other scale related notes](https://github.com/binhnguyennus/awesome-scalability?tab=readme-ov-file#talk)
  - [awesome-low-level-design](https://github.com/ashishps1/awesome-low-level-design?tab=readme-ov-file)


### API  
  - [Top 7 Ways to 10x Your API Performance](https://www.youtube.com/watch?v=zvWKqUiovAM)
  - [Top 12 Tips For API Security](https://www.youtube.com/watch?v=6WZ6S-qmtqY)
  - [Good APIs Vs Bad APIs: 7 Tips for API Design](https://www.youtube.com/watch?v=_gQaygjm_hg)
  - [Rest API - Best Practices - Design](https://www.youtube.com/watch?v=1Wl-rtew1_E)


### VS code 
   - [25 VS Code Productivity Tips and Speed Hacks](https://www.youtube.com/watch?v=ifTF3ags0XI)
   - [GitHub Copilot Top Features Explained](https://www.youtube.com/watch?v=KjyMQzoJo8Y)

###  Go String Formaters
- %s: Prints the value as a string.
- %d: Prints the value as a decimal integer.
- %f: Prints the value as a floating-point number.
- %b: Prints the value as a binary number.
- %x: Prints the value as a hexadecimal number.
- %o: Prints the value as an octal number.
- %q: Prints the value as a quoted Go string.
- %p: Prints the value as a pointer.
- %+v: Prints the value with the full type information.
 ```go
var x int = 42
var s string = "Hello"
var f float64 = 3.14

fmt.Printf("%s\n", s)    // Output: Hello
fmt.Printf("%d\n", x)    // Output: 42
fmt.Printf("%f\n", f)    // Output: 3.140000
fmt.Printf("%b\n", x)    // Output: 101010
fmt.Printf("%x\n", x)    // Output: 2a
fmt.Printf("%o\n", x)    // Output: 52
fmt.Printf("%q\n", s)    // Output: "Hello"
fmt.Printf("%p\n", &x)   // Output: 0x1040a000 (address of x)
fmt.Printf("%+v\n", x)   // Output: int(42)
 ```

