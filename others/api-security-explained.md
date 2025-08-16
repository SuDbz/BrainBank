# API Security Explained

## Table of Contents
1. [Rate Limiting](#rate-limiting)
    - [Per Endpoint](#per-endpoint)
    - [Per User/IP](#per-userip)
    - [Overall (DDos Mitigation)](#overall-ddos-mitigation)
2. [CORS (Cross-Origin Resource Sharing)](#cors-cross-origin-resource-sharing)
3. [SQL & NoSQL Injection](#sql--nosql-injection)
4. [CSRF (Cross-Site Request Forgery)](#csrf-cross-site-request-forgery)
5. [Authentication & Authorization](#authentication--authorization)
6. [HTTPS/TLS](#httpstls)
7. [Input Validation](#input-validation)
8. [Logging & Monitoring](#logging--monitoring)
9. [Security Headers](#security-headers)
10. [Error Handling](#error-handling)
11. [XSS (Cross-Site Scripting)](#xss-cross-site-scripting)

---


## Rate Limiting
Rate limiting restricts how many requests clients can make to an API in a given time window. It helps prevent abuse, protects resources, and mitigates denial-of-service attacks.

### Per Endpoint

**Python (Flask):**
```python
from flask_limiter import Limiter  # Import rate limiting library for Flask
from flask import Flask  # Import Flask web framework
app = Flask(__name__)  # Create Flask application instance
limiter = Limiter(app, key_func=lambda: "global")  # Create rate limiter with global key (affects all users equally)

@app.route('/login')  # Define HTTP route for /login endpoint
@limiter.limit("5 per minute")  # Apply rate limit: maximum 5 requests per minute
def login():  # Function to handle login requests
    return "Login endpoint"  # Return simple response (in real app, handle authentication)
```

**Go (Gin):**
```go
package main  // Define package name - main package for executable programs
import (  // Import required packages
    "github.com/gin-gonic/gin"  // Gin web framework for building REST APIs
    "github.com/ulule/limiter/v3"  // Rate limiting library
    ginlimiter "github.com/ulule/limiter/v3/drivers/middleware/gin"  // Gin middleware for rate limiter
    memory "github.com/ulule/limiter/v3/drivers/store/memory"  // In-memory store for rate limit data
)

func main() {  // Main function - entry point of the program
    r := gin.Default()  // Create Gin router with default middleware (logger, recovery)
    rate, _ := limiter.NewRateFromFormatted("5-M")  // Create rate limit: 5 requests per minute ("5-M" format)
    store := memory.NewStore()  // Create in-memory store to track rate limit counters
    middleware := ginlimiter.NewMiddleware(limiter.New(store, rate))  // Create rate limiting middleware
    r.Use(middleware)  // Apply rate limiting middleware to all routes
    r.GET("/login", func(c *gin.Context) {  // Define GET route for /login endpoint
        c.String(200, "Login endpoint")  // Return HTTP 200 with plain text response
    })
    r.Run()  // Start HTTP server on default port 8080
}
```

### Per User/IP
Limits requests based on user identity or IP address.

**Python:**
```python
limiter = Limiter(app, key_func=lambda: request.remote_addr)  # Create limiter using client IP address as key for per-IP limiting

@app.route('/api/data')  # Define route for /api/data endpoint
@limiter.limit("100 per hour")  # Apply rate limit: 100 requests per hour per IP address
def get_data():  # Function to handle data requests
    return "Data endpoint"  # Return simple response (in real app, return actual data)
```

**Go:**
```go
// Gin + Limiter automatically uses IP for rate limiting by default
```

### Overall (DDos Mitigation)
Global limits to prevent overwhelming the API server.

**Python:**
```python
limiter = Limiter(app, key_func=lambda: "global")  # Create limiter with global key (same limit for all users)
limiter.limit("1000 per minute")  # Set global rate limit: 1000 requests per minute across all users
```

**Go:**
```go
// Set a global rate limit in Gin using Limiter as above
```

---


## CORS (Cross-Origin Resource Sharing)
CORS controls which domains can access your API from browsers. It prevents unauthorized cross-origin requests.

**Python (Flask):**
```python
from flask_cors import CORS  # Import CORS (Cross-Origin Resource Sharing) extension
app = Flask(__name__)  # Create Flask application instance
CORS(app, origins=["https://yourdomain.com"])  # Enable CORS only for specific domain (yourdomain.com)
# This prevents other websites from making requests to your API from browsers
```
**Go (Gin):**
```go
package main  // Define package name
import (  // Import required packages    
    "github.com/gin-gonic/gin"  // Gin web framework
    "github.com/gin-contrib/cors"  // CORS middleware for Gin
    "time"  // Standard library for time operations
)

func main() {  // Main function
    r := gin.Default()  // Create Gin router with default middleware
    r.Use(cors.New(cors.Config{  // Apply CORS middleware with custom configuration
        AllowOrigins:     []string{"https://yourdomain.com"},  // Only allow requests from this domain
        AllowMethods:     []string{"GET", "POST"},  // Only allow these HTTP methods
        AllowHeaders:     []string{"Origin"},  // Allow these headers in requests
        ExposeHeaders:    []string{"Content-Length"},  // Expose these headers to client
        AllowCredentials: true,  // Allow cookies and credentials in CORS requests
        MaxAge: 12 * time.Hour,  // Cache preflight response for 12 hours
    }))
    r.GET("/", func(c *gin.Context) {  // Define GET route for root path
        c.String(200, "Hello World")  // Return HTTP 200 with plain text
    })
    r.Run()  // Start HTTP server
}
```

---


## SQL & NoSQL Injection
Injection attacks exploit insecure query construction to execute malicious code or access unauthorized data.
- **SQL Injection:** Manipulates SQL queries via user input.
- **NoSQL Injection:** Manipulates NoSQL queries (e.g., MongoDB) via user input.

**Prevention:**
- Always use parameterized queries or ORM methods.

**Python (SQL):**
```python
import sqlite3  # Import SQLite database library
conn = sqlite3.connect('db.sqlite3')  # Create connection to SQLite database file
cursor = conn.cursor()  # Create cursor object to execute SQL commands
username = request.args.get('username')  # Get username parameter from HTTP request
# BAD: Vulnerable to injection
# cursor.execute(f"SELECT * FROM users WHERE username = '{username}'")  # DON'T DO THIS - allows SQL injection
# GOOD: Safe
cursor.execute("SELECT * FROM users WHERE username = ?", (username,))  # Use parameterized query - safe from injection
# The ? placeholder is replaced by the database driver, preventing malicious SQL
```

**Go (SQL):**
```go
import (  // Import required packages
    "database/sql"  // Standard SQL database interface
    _ "github.com/mattn/go-sqlite3"  // SQLite driver (blank import to register driver)
)

func getUser(db *sql.DB, username string) {  // Function to get user from database
    // BAD: Vulnerable to injection
    // db.Query("SELECT * FROM users WHERE username = '" + username + "'")  // DON'T DO THIS
    // GOOD: Safe
    db.Query("SELECT * FROM users WHERE username = ?", username)  // Use placeholder ? for safe parameterized query
    // The database driver safely escapes the username parameter
}
```

**Python (NoSQL):**
```python
from pymongo import MongoClient  # Import MongoDB client library
client = MongoClient()  # Create MongoDB client (connects to localhost:27017 by default)
db = client.mydatabase  # Get reference to 'mydatabase' database
username = request.args.get('username')  # Get username from HTTP request parameters
# BAD: Vulnerable to injection
# db.users.find({"$where": f"this.username == '{username}'"})  # DON'T DO THIS - allows JavaScript injection
# GOOD: Safe
user = db.users.find_one({"username": username})  # Use direct field matching - safe from injection
# MongoDB automatically escapes the username value when using direct field queries
```

**Go (NoSQL - MongoDB):**
```go
import (  // Import required packages
    "go.mongodb.org/mongo-driver/mongo"  // Official MongoDB driver for Go
    "go.mongodb.org/mongo-driver/bson"  // BSON (Binary JSON) encoding/decoding
    "context"  // Standard library for handling request contexts and timeouts
)

func getUser(collection *mongo.Collection, username string) {  // Function to get user from MongoDB
    // BAD: Vulnerable to injection
    // collection.Find(context.TODO(), bson.M{"$where": "this.username == '" + username + "'"})  // DON'T DO THIS
    // GOOD: Safe
    collection.FindOne(context.TODO(), bson.M{"username": username})  // Use BSON document for safe querying
    // context.TODO() provides empty context, bson.M creates a map for the query filter
    // The driver safely handles the username value, preventing injection
}
```

---


## CSRF (Cross-Site Request Forgery)
CSRF tricks users into submitting unwanted actions to an API where they're authenticated. Prevent by using anti-CSRF tokens and checking request origins.

**Python (Flask-WTF):**
```python
from flask_wtf.csrf import CSRFProtect  # Import CSRF protection extension for Flask
app = Flask(__name__)  # Create Flask application instance
csrf = CSRFProtect(app)  # Enable CSRF protection for all forms in the application
# This automatically adds CSRF tokens to forms and validates them on submission
```

**Go (Gin):**
```go
package main  // Define package name
import (  // Import required packages
    "github.com/gin-gonic/gin"  // Gin web framework
    "github.com/utrack/gin-csrf"  // CSRF protection middleware for Gin
)

func main() {  // Main function
    r := gin.Default()  // Create Gin router with default middleware
    r.Use(csrf.Middleware(csrf.Options{  // Apply CSRF middleware with configuration
        Secret: "secret123",  // Secret key used to sign CSRF tokens (change in production!)
    }))
    r.GET("/form", func(c *gin.Context) {  // Define GET route to display form
        token := csrf.GetToken(c)  // Generate CSRF token for this request
        c.String(200, "CSRF token: %s", token)  // Return the token (in real app, embed in form)
    })
    r.Run()  // Start HTTP server
}
```

---


## Authentication & Authorization
Authentication verifies who the user is, and authorization determines what they can do.

**Python (Flask-JWT-Extended):**
```python
from flask_jwt_extended import JWTManager, jwt_required, create_access_token  # Import JWT library components
from flask import Flask, request, jsonify  # Import Flask framework and utilities

app = Flask(__name__)  # Create Flask application instance
app.config['JWT_SECRET_KEY'] = 'super-secret'  # Set secret key for signing JWT tokens (CHANGE IN PRODUCTION!)
jwt = JWTManager(app)  # Initialize JWT manager for the Flask app

@app.route('/login', methods=['POST'])  # Define POST route for user login
def login():  # Function to handle login requests
    username = request.json.get('username')  # Extract username from JSON request body
    password = request.json.get('password')  # Extract password from JSON request body
    # Authenticate user (check username/password)
    if username == 'admin' and password == 'password':  # Simple auth check (replace with real auth)
        access_token = create_access_token(identity=username)  # Create JWT token with username as identity
        return jsonify(access_token=access_token)  # Return token as JSON response
    return jsonify(message='Invalid credentials'), 401  # Return error if auth fails

@app.route('/protected', methods=['GET'])  # Define GET route for protected resource
@jwt_required()  # Decorator that requires valid JWT token to access this route
def protected():  # Function to handle protected requests
    return jsonify(message='This is a protected route')  # Return success message if token is valid
```

**Go (Gin + JWT):**
```go
package main  // Define package name

import (  // Import required packages
    "net/http"  // Standard HTTP constants and utilities
    "time"  // Standard library for time operations
    
    "github.com/gin-gonic/gin"  // Gin web framework
    "github.com/golang-jwt/jwt/v4"  // JWT library for Go
)

var jwtKey = []byte("my_secret_key")  // Secret key for signing JWT tokens (store securely in production!)

type Claims struct {  // Define custom claims structure for JWT payload
    Username string `json:"username"`  // Username field to store in JWT
    jwt.RegisteredClaims  // Embed standard JWT claims (exp, iat, etc.)
}

func generateToken(username string) (string, error) {  // Function to create JWT token
    expirationTime := time.Now().Add(5 * time.Minute)  // Set token expiration to 5 minutes from now
    claims := &Claims{  // Create claims object with user data
        Username: username,  // Set username in claims
        RegisteredClaims: jwt.RegisteredClaims{  // Set standard claims
            ExpiresAt: jwt.NewNumericDate(expirationTime),  // Set expiration time
        },
    }
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)  // Create new JWT with HMAC SHA256 signing
    return token.SignedString(jwtKey)  // Sign token with secret key and return string representation
}

func validateToken(tokenStr string) (*Claims, error) {  // Function to validate and parse JWT token
    claims := &Claims{}  // Create empty claims object to populate
    token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
        return jwtKey, nil  // Return the secret key for token verification
    })
    if err != nil || !token.Valid {  // Check if parsing failed or token is invalid
        return nil, err  // Return error if token is invalid
    }
    return claims, nil  // Return parsed claims if token is valid
}

func authMiddleware() gin.HandlerFunc {  // Create middleware function for authentication
    return func(c *gin.Context) {  // Return Gin handler function
        tokenString := c.GetHeader("Authorization")  // Get Authorization header from request
        if tokenString == "" {  // Check if token is missing
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing token"})  // Return 401 error
            c.Abort()  // Stop processing this request
            return
        }
        
        claims, err := validateToken(tokenString)  // Validate the provided token
        if err != nil {  // Check if token validation failed
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})  // Return 401 error
            c.Abort()  // Stop processing this request
            return
        }
        
        c.Set("username", claims.Username)  // Store username in request context for later use
        c.Next()  // Continue to next middleware/handler
    }
}

func main() {  // Main function - entry point
    r := gin.Default()  // Create Gin router with default middleware
    
    r.POST("/login", func(c *gin.Context) {  // Define POST route for login
        var loginData struct {  // Define anonymous struct for login request body
            Username string `json:"username"`  // Username field
            Password string `json:"password"`  // Password field
        }
        
        if err := c.ShouldBindJSON(&loginData); err != nil {  // Parse JSON request body into struct
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})  // Return 400 if JSON is invalid
            return
        }
        
        // Replace with real authentication
        if loginData.Username == "admin" && loginData.Password == "password" {  // Simple auth check
            token, err := generateToken(loginData.Username)  // Generate JWT token for user
            if err != nil {  // Check if token generation failed
                c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})  // Return 500 error
                return
            }
            c.JSON(http.StatusOK, gin.H{"access_token": token})  // Return token in response
        } else {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})  // Return 401 for bad credentials
        }
    })
    
    r.GET("/protected", authMiddleware(), func(c *gin.Context) {  // Define protected route with auth middleware
        username := c.GetString("username")  // Get username from context (set by middleware)
        c.JSON(http.StatusOK, gin.H{"message": "This is a protected route", "user": username})  // Return success with user info
    })
    
    r.Run()  // Start HTTP server on default port 8080
}
```

---


## HTTPS/TLS
Always use HTTPS to encrypt data in transit. Obtain a TLS certificate and configure your server to redirect HTTP to HTTPS.

**Python (Flask):**
```python
if __name__ == "__main__":  # Check if script is run directly (not imported)
    app.run(ssl_context=("path/to/cert.pem", "path/to/key.pem"))  # Start Flask app with SSL/TLS encryption
    # cert.pem contains the SSL certificate, key.pem contains the private key
    # This enables HTTPS on the server to encrypt data in transit
```

**Go (Gin):**
```go
func main() {  // Main function
    r := gin.Default()  // Create Gin router with default middleware
    r.RunTLS(":443", "path/to/cert.pem", "path/to/key.pem")  // Start HTTPS server on port 443
    // First parameter is port, second is certificate file, third is private key file
    // This enables TLS encryption for all communications with the server
}
```

---


## Input Validation
Validate and sanitize all user inputs to prevent malformed data from causing issues.

**Python:**
```python
from wtforms import Form, StringField, validators  # Import form validation library

class LoginForm(Form):  # Define form class for validation
    username = StringField('Username', [validators.Length(min=4, max=25)])  # Username field with length validation
    password = StringField('Password', [validators.Length(min=6, max=35)])  # Password field with length validation
    # validators.Length ensures input is within specified character limits
    # This prevents overly short passwords and excessively long inputs that could cause issues
```

**Go:**
```go
import "github.com/go-playground/validator/v10"  // Import validation library for Go

type User struct {  // Define struct for user data
    Username string `json:"username" validate:"required,min=4,max=25"`  // Username with validation tags
    Password string `json:"password" validate:"required,min=6,max=35"`  // Password with validation tags
    // Validation tags: required=field must be present, min/max=character length limits
}

func validateInput(user User) error {  // Function to validate user input
    validate := validator.New()  // Create new validator instance
    return validate.Struct(user)  // Validate the user struct against the defined tags
    // Returns nil if valid, error with details if validation fails
}
```

---


## Logging & Monitoring
Implement logging and monitoring to detect and respond to security incidents.

**Python:**
```python
import logging  # Import Python's built-in logging library
logging.basicConfig(level=logging.INFO)  # Configure logging to show INFO level and above messages

@app.route('/api/data')  # Define route for API endpoint
def get_data():  # Function to handle data requests
    app.logger.info('Data requested')  # Log information about the request
    # This creates an audit trail of who accessed what data and when
    # Useful for security monitoring and debugging
```

**Go:**
```go
import "log"  // Import Go's standard logging package

func main() {  // Main function
    log.Println("Server started")  // Log server startup message
    r.GET("/api/data", func(c *gin.Context) {  // Define GET route for data endpoint
        log.Println("Data requested")  // Log each data request
        // These logs help track server activity and can be used for security auditing
        // In production, consider using structured logging libraries like logrus or zap
    })
}
```

---


## Security Headers
Set security-related HTTP headers to protect against common attacks.

**Python (Flask-Talisman):**
```python
from flask_talisman import Talisman  # Import Talisman security headers library
Talisman(app)  # Apply default security headers to all responses
# Talisman automatically adds headers like Content-Security-Policy, X-Frame-Options,
# Strict-Transport-Security, etc. to protect against XSS, clickjacking, and other attacks
```

**Go (Gin):**
```go
import "github.com/gin-contrib/securityheaders"  // Import security headers middleware

func main() {  // Main function
    r := gin.Default()  // Create Gin router
    r.Use(securityheaders.Default())  // Apply default security headers middleware
    // This adds headers like X-Frame-Options, X-Content-Type-Options, etc.
    // to protect against common web vulnerabilities
}
```

---


## Error Handling
Handle errors gracefully to avoid exposing stack traces or sensitive information.

**Python:**
```python
@app.errorhandler(404)  # Decorator to handle 404 Not Found errors
def not_found(error):  # Function to handle 404 errors
    return {"message": "Not found"}, 404  # Return generic error message, don't expose internal details
    # This prevents leaking information about internal file structure or routes
```

**Go:**
```go
r.NoRoute(func(c *gin.Context) {  // Handle requests to non-existent routes
    c.JSON(http.StatusNotFound, gin.H{"message": "Not found"})  // Return 404 with generic message
    // This prevents exposing internal application structure through error messages
})
```

---


## XSS (Cross-Site Scripting)
XSS is a vulnerability that allows attackers to inject malicious scripts into web pages viewed by other users. These scripts can steal cookies, session tokens, or perform actions on behalf of the user.

### How XSS Works
- An attacker finds a way to inject JavaScript (or other scripts) into a web page.
- When another user loads the page, the script runs in their browser, potentially stealing data or hijacking sessions.

### Types of XSS
- **Stored XSS:** Malicious script is saved on the server (e.g., in a database) and served to users.
- **Reflected XSS:** Malicious script is reflected off the server (e.g., via a query parameter) and executed in the browser.
- **DOM-based XSS:** The vulnerability is in client-side code that modifies the DOM based on user input.

### Prevention
- Always escape and sanitize user input before rendering it in HTML.
- Use security headers like Content-Security-Policy.
- Validate and encode output.

**Python (Flask):**
```python
from markupsafe import escape  # Import HTML escaping function from MarkupSafe library
@app.route('/greet')  # Define route for greeting endpoint
def greet():  # Function to handle greeting requests
    name = request.args.get('name', '')  # Get 'name' parameter from URL query string, default to empty string
    return f"Hello, {escape(name)}!"  # Escape HTML special characters in the name to prevent XSS
    # escape() converts characters like <, >, &, " to their HTML entities (&lt;, &gt;, etc.)
    # This prevents malicious scripts from being executed in the browser
```

**Go (Gin):**
```go
import "html"  // Import Go's standard HTML package for escaping
r.GET("/greet", func(c *gin.Context) {  // Define GET route for greeting endpoint
    name := c.Query("name")  // Get 'name' parameter from URL query string
    safeName := html.EscapeString(name)  // Escape HTML special characters in the name
    c.String(200, "Hello, %s!", safeName)  // Return escaped name in response
    // html.EscapeString() converts <, >, &, ', " to HTML entities
    // This prevents XSS attacks by ensuring user input is treated as text, not executable code
})
```
