package handler

import (
	"fmt"
	"net/http"
	"time"

	"github.com/hardik302001/hardik-portfolio-backend/internal/response"
)

var serverStartTime = time.Now()

var taglines = []string{
	"I don't write bugs anymore. I write 'undocumented features' and delegate the cleanup.",
	"Senior enough to mass-produce the architecture. Wise enough to blame the infra.",
	"My code reviews have a body count. Interns call me 'The Closer'.",
	"I don't debug. I just stare at the code until it confesses.",
	"They gave me a pager. I gave it separation anxiety.",
	"Promoted for knowing which fires to fight and which to just monitor with a dashboard.",
	"My standup update is 'same as yesterday' but somehow the deploy is on fire.",
	"I estimate in sprints. I deliver in quarters. Management hasn't noticed.",
	"The junior asked how long it would take. I made up a number. We both knew it was fiction.",
	"Architecting systems by day. Mass-producing Jira tickets to explain them by night.",
	"Five years in, my greatest skill is mass-producing convincing root cause analyses.",
	"They wanted a 10x engineer. I gave them 10 microservices and a Grafana link.",
	"I mass-produce elegant solutions to problems that shouldn't exist but do because of my last elegant solution.",
	"Senior dev privilege: mass-producing opinions in design docs and calling it 'leadership'.",
	"I mass-produce the chai, the code, and the chaos — in that order, every morning.",
}

func GetProfile(w http.ResponseWriter, r *http.Request) {
	now := time.Now()
	ist, _ := time.LoadLocation("Asia/Kolkata")
	uptime := now.Sub(serverStartTime)

	idx := int(now.Unix()/60) % len(taglines)

	data := map[string]any{
		"name":       "Hardik Sharma",
		"role":       "SDE II @ Zomato — Hyperpure",
		"location":   "Gurgaon, India",
		"tagline":    taglines[idx],
		"server_time": now.In(ist).Format("02 Jan 2006, 03:04 PM IST"),
		"uptime_seconds": int(uptime.Seconds()),
		"uptime_human":   formatUptime(uptime),
	}

	response.WriteJSON(w, http.StatusOK, data)
}

func formatUptime(d time.Duration) string {
	days := int(d.Hours()) / 24
	hours := int(d.Hours()) % 24
	minutes := int(d.Minutes()) % 60

	if days > 0 {
		return fmt.Sprintf("%dd %dh %dm", days, hours, minutes)
	}
	if hours > 0 {
		return fmt.Sprintf("%dh %dm", hours, minutes)
	}
	return fmt.Sprintf("%dm", minutes)
}
