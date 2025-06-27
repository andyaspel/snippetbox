package main

import (
	"net/http"
	"sync"
	"time"

	"log"

	"github.com/fatih/color"
	"github.com/natefinch/lumberjack"
	"golang.org/x/time/rate"

	"github.com/andyaspel/snippetbox/pkg/models"
)

func SecureHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Enforce HTTPS
		w.Header().Set("Strict-Transport-Security", "max-age=63072000; includeSubDomains; preload")
		// Prevent XSS, clickjacking, and other attacks
		w.Header().Set("X-Frame-Options", "DENY")
		w.Header().Set("X-Content-Type-Options", "nosniff")
		w.Header().Set("X-XSS-Protection", "0") // Modern browsers ignore this, CSP is preferred
		// Content Security Policy: adjust as needed for your app
		w.Header().Set("Content-Security-Policy", "frame-ancestors 'none';")
		// Referrer Policy
		w.Header().Set("Referrer-Policy", "strict-origin-when-cross-origin")
		// Permissions Policy (formerly Feature Policy)
		w.Header().Set("Permissions-Policy", "geolocation=(), camera=(), microphone=()")
		// Cross-Origin Resource Policy
		w.Header().Set("Cross-Origin-Resource-Policy", "same-origin")
		// Cache Control
		w.Header().Set("Cache-Control", "no-store, no-cache, must-revalidate, proxy-revalidate")
		w.Header().Set("Pragma", "no-cache")
		w.Header().Set("Expires", "0")

		next.ServeHTTP(w, r)
	})
}

var (
	rateLimiters   = make(map[string]*rate.Limiter)
	rateLimitersMu sync.Mutex
)

func getRateLimiter(ip string) *rate.Limiter {
	rateLimitersMu.Lock()
	defer rateLimitersMu.Unlock()
	limiter, exists := rateLimiters[ip]
	if !exists {
		// 5 requests per second, burst of 10
		limiter = rate.NewLimiter(5, 10)
		rateLimiters[ip] = limiter
	}
	return limiter
}

// Periodic cleanup to prevent memory leaks
func cleanupRateLimiters() {
	ticker := time.NewTicker(10 * time.Minute)
	for range ticker.C {
		rateLimitersMu.Lock()
		for ip, limiter := range rateLimiters {
			if limiter.Allow() {
				delete(rateLimiters, ip)
			}
		}
		rateLimitersMu.Unlock()
	}
}

func RateLimit(next http.Handler) http.Handler {
	// Start cleanup goroutine once
	var once sync.Once
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		once.Do(func() { go cleanupRateLimiters() })
		ip := r.RemoteAddr
		limiter := getRateLimiter(ip)
		if !limiter.Allow() {
			w.Header().Set("Retry-After", "1")
			http.Error(w, http.StatusText(http.StatusTooManyRequests), http.StatusTooManyRequests)
			return
		}
		next.ServeHTTP(w, r)
	})
}

var (
	logger      *log.Logger
	consoleInfo = color.New(color.FgGreen).SprintFunc()
	consoleWarn = color.New(color.FgYellow).SprintFunc()
	consoleErr  = color.New(color.FgRed).SprintFunc()
)

func init() {
	logger = log.New(&lumberjack.Logger{
		Filename:   "logs/access-" + time.Now().Format("2006-01-02") + ".log",
		MaxSize:    10, // megabytes
		MaxBackups: 7,
		MaxAge:     30, //days
		Compress:   true,
	}, "", log.LstdFlags|log.Lshortfile)
}

type responseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (rw *responseWriter) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}

func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		rw := &responseWriter{w, http.StatusOK}
		next.ServeHTTP(rw, r)
		duration := time.Since(start)

		logEntry := logEntry{
			Time:      time.Now().Format(time.RFC3339),
			Method:    r.Method,
			URL:       r.URL.String(),
			Status:    rw.statusCode,
			Duration:  duration.String(),
			Remote:    r.RemoteAddr,
			UserAgent: r.UserAgent(),
		}
		logger.Printf("%+v", logEntry)
		// db2, err := connectToLogs()
		// if err != nil {
		// 	logger.Fatal(err)
		// }
		// var Log models.Log
		// err = db2.AutoMigrate(&Log)
		// if err != nil {
		// 	log.Fatal(err)
		// }
		logToDB(logEntry)

		// Terminal output with color
		var statusColor func(a ...interface{}) string
		switch {
		case rw.statusCode >= 500:
			statusColor = consoleErr
		case rw.statusCode >= 400:
			statusColor = consoleWarn
		default:
			statusColor = consoleInfo
		}
		logLine := time.Now().Format("15:04:05") + " " + r.Method + " " + r.URL.String() + " [" + statusColor(rw.statusCode) + "] " + duration.String() + " " + r.RemoteAddr
		log.Println(logLine)
	})
}

func logToDB(entry logEntry) {
	db, err := connectToLogs()
	if err != nil {
		return
	}
	logModel := models.Log{
		Time:      entry.Time,
		Method:    entry.Method,
		URL:       entry.URL,
		Status:    entry.Status,
		Duration:  entry.Duration,
		Remote:    entry.Remote,
		UserAgent: entry.UserAgent,
	}
	db.Create(&logModel)
}
