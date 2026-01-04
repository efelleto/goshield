package main

import (
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/fatih/color"
	"github.com/hpcloud/tail"
)

func main() {
	blacklist := make(map[string]int)

	warnStyle := color.New(color.FgHiYellow, color.Bold).SprintFunc()
	blockStyle := color.New(color.FgHiRed, color.Bold).SprintFunc()

	blueIP := color.RGB(0, 150, 255).SprintFunc()

	timeGray := color.RGB(90, 90, 90).SprintFunc()
	descGray := color.RGB(150, 150, 150).SprintFunc()
	pathGray := color.RGB(110, 110, 110).SprintFunc()

	header := color.New(color.FgHiCyan, color.Bold).SprintFunc()

	t, _ := tail.TailFile("./logs/access.log", tail.Config{Follow: true, ReOpen: true, Poll: true})
	re := regexp.MustCompile(`(\d+\.\d+\.\d+\.\d+).*?"GET\s+(.*?)\s+HTTP`)

	fmt.Printf("\n  %s\n", header("goshield - github.com/efelleto/goshield"))
	fmt.Printf("  %s\n\n", blueIP("engine: monitoring active traffic..."))

	for line := range t.Lines {
		match := re.FindStringSubmatch(line.Text)
		if len(match) > 2 {
			ip := match[1]
			path := match[2]

			if strings.Contains(path, "admin") || strings.Contains(path, ".env") || strings.Contains(path, "config") {
				blacklist[ip]++
				score := blacklist[ip]
				timestamp := time.Now().Format("15:04:05")

				if score >= 5 {
					// formato: [blocked] 15:04:05 | critical threshold: ip
					fmt.Printf("  %s %s %s %s\n",
						blockStyle("[blocked]"),
						timeGray(timestamp),
						descGray("| critical threshold:"),
						blueIP(ip))
				} else {
					// formato: [warning] 15:04:05 | suspicious: ip -> path
					fmt.Printf("  %s %s %s %s %s %s\n",
						warnStyle("[warning]"),
						timeGray(timestamp),
						descGray("| suspicious access:"),
						blueIP(ip),
						pathGray("»"),
						pathGray(path))
				}
			}
		}
	}
}
