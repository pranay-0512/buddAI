package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/google"
)

func main() {
	r := gin.Default()

	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	// Load HTML templates
	r.LoadHTMLGlob("templates/*")

	// Get configuration from environment variables
	clientID := os.Getenv("GOOGLE_CLIENT_ID")
	clientSecret := os.Getenv("GOOGLE_CLIENT_SECRET")
	clientCallbackURL := os.Getenv("GOOGLE_CLIENT_CALLBACK_URL")
	sessionSecret := os.Getenv("SESSION_SECRET")

	// Validate required environment variables
	if clientID == "" || clientSecret == "" || clientCallbackURL == "" {
		log.Fatal("Missing required Google OAuth configuration. Please set GOOGLE_CLIENT_ID, GOOGLE_CLIENT_SECRET, and GOOGLE_CLIENT_CALLBACK_URL in .env file")
	}

	if sessionSecret == "" {
		log.Fatal("SESSION_SECRET must be set in .env file")
	}

	// Initialize session middleware
	store := cookie.NewStore([]byte(sessionSecret))
	r.Use(sessions.Sessions("gothsession", store))

	// Initialize Gothic
	gothic.Store = store

	// Setup Google OAuth provider
	goth.UseProviders(
		google.New(clientID, clientSecret, clientCallbackURL),
	)

	// Setup routes
	r.GET("/", home)
	r.GET("/auth/:provider", signInWithProvider)
	r.GET("/auth/:provider/callback", callbackHandler)
	r.GET("/success", success)

	// Start server
	if err := r.Run(":5000"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}

func home(c *gin.Context) {
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		c.String(http.StatusInternalServerError, "Error loading template")
		return
	}

	if err := tmpl.Execute(c.Writer, gin.H{}); err != nil {
		c.String(http.StatusInternalServerError, "Error executing template")
		return
	}
}

func signInWithProvider(c *gin.Context) {
	provider := c.Param("provider")
	if provider == "" {
		c.String(http.StatusBadRequest, "Provider is required")
		return
	}

	q := c.Request.URL.Query()
	q.Add("provider", provider)
	c.Request.URL.RawQuery = q.Encode()

	gothic.BeginAuthHandler(c.Writer, c.Request)
}

func callbackHandler(c *gin.Context) {
	provider := c.Param("provider")
	if provider == "" {
		c.String(http.StatusBadRequest, "Provider is required")
		return
	}

	q := c.Request.URL.Query()
	q.Add("provider", provider)
	c.Request.URL.RawQuery = q.Encode()

	user, err := gothic.CompleteUserAuth(c.Writer, c.Request)
	if err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("Error completing auth: %v", err))
		return
	}

	// Store user info in session
	session := sessions.Default(c)
	session.Set("user_id", user.UserID)
	session.Set("email", user.Email)
	session.Set("name", user.Name)
	if err := session.Save(); err != nil {
		c.String(http.StatusInternalServerError, "Error saving session")
		return
	}
	fmt.Print(session)
	c.Redirect(http.StatusTemporaryRedirect, "/success")
}

func success(c *gin.Context) {
	// Get user info from session
	session := sessions.Default(c)
	name := session.Get("name")
	email := session.Get("email")

	// Render HTML template with user data
	c.HTML(http.StatusOK, "success.html", gin.H{
		"Name":  name,
		"Email": email,
	})
}
