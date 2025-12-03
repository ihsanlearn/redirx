package logger

import (
	"fmt"
	"github.com/fatih/color"
)

var (
	Red    = color.New(color.FgRed).SprintFunc()
	Green  = color.New(color.FgGreen).SprintFunc()
	Yellow = color.New(color.FgYellow).SprintFunc()
	Blue   = color.New(color.FgBlue).SprintFunc()
	Magenta = color.New(color.FgMagenta).SprintFunc()
	Cyan   = color.New(color.FgCyan).SprintFunc()
	White  = color.New(color.FgWhite).SprintFunc()
	Black  = color.New(color.FgBlack).SprintFunc()
	Bold   = color.New(color.Bold).SprintFunc()
)

func Info(format string, args ...interface{}) {
	prefix := Blue("[INFO]")
	msg := fmt.Sprintf(format, args...)
	fmt.Printf("%s %s\n", prefix, msg)
}

func Vulnerable(format string, args ...interface{}) {
	prefix := Green("[VULNERABLE]")
	msg := fmt.Sprintf(format, args...)
	fmt.Printf("%s %s\n", prefix, Bold(msg)) 
}

func NotVulnerable(format string, args ...interface{}) {
	prefix := Yellow("[NOT VULNERABLE]")
	msg := fmt.Sprintf(format, args...)
	fmt.Printf("%s %s\n", prefix, msg)
}

func Error(format string, args ...interface{}) {
	prefix := Red("[ERROR]")
	msg := fmt.Sprintf(format, args...)
	fmt.Printf("%s %s\n", prefix, msg)
}

func Debug(format string, args ...interface{}) {
	prefix := Magenta("[DEBUG]")
	msg := fmt.Sprintf(format, args...)
	fmt.Printf("%s %s\n", prefix, msg)
}

func PrintBanner() {
	banner := `
    ____        _ _     _  __
   / __ \___  _| (_)___| |/ /
  / /_/ / _ \/ _ | / __/   /  
 / _, _/  __/ /_ |/ / /   |   
/_/ |_|\___/\____/_/ /_/|_|   
                                
    RedirX - Open Redirect Scanner
	`
	color.Cyan(banner)
}
