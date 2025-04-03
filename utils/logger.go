package utils

import (
	"fmt"
	"log"
	"os"
	"time"
)

// ANSI-farger (brukes kun for terminal)
const (
	colorReset  = "\033[0m"
	colorRed    = "\033[31m"
	colorGreen  = "\033[32m"
	colorYellow = "\033[33m"
)

var logFile *os.File

// InitLogger setter opp logging til fil og terminal (farger)
func InitLogger(path string) {
	// Sørg for at logs/-mappen finnes
	if err := os.MkdirAll("logs", os.ModePerm); err != nil {
		log.Fatalf("kan ikke opprette logs/-mappe: %v", err)
	}

	// Åpne loggfil
	var err error
	logFile, err = os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("kan ikke åpne loggfil: %v", err)
	}

	// Sett standard logger til fil
	log.SetOutput(logFile)
}

// CloseLogger lukkes i main via defer
func CloseLogger() {
	if logFile != nil {
		logFile.Close()
	}
}

// interne formatter
func logMessage(level string, color string, msg string) {
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	formatted := fmt.Sprintf("[%s] [%s] %s\n", timestamp, level, msg)

	// Skriv til loggfil (via standardlogg)
	log.Print(formatted)

	// Skriv farget versjon til terminal
	fmt.Print(color + formatted + colorReset)
}

// Offentlige funksjoner
func Info(msg string)  { logMessage("INFO", colorGreen, msg) }
func Warn(msg string)  { logMessage("WARN", colorYellow, msg) }
func Error(msg string) { logMessage("ERROR", colorRed, msg) }
