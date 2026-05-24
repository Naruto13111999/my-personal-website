package data

import "time"

type ActivityDay struct {
	Date  string
	Label string
	Level int
	Count int
}

type ActivityWeek struct {
	Days []ActivityDay
}

type ActivityGraph struct {
	Total       int
	Weeks       []ActivityWeek
	MonthLabels []MonthLabel
	Label       string
	Note        string
}

type MonthLabel struct {
	Name  string
	Index int
}

func GetActivityGraph() ActivityGraph {
	g := buildActivityGraph("Career-wide shipping activity")
	g.Note = "Features shipped, code reviews & production support"
	return g
}

func buildActivityGraph(label string) ActivityGraph {
	end := time.Now().UTC()
	gridStart := end.AddDate(-1, 1, 0)

	for gridStart.Weekday() != time.Sunday {
		gridStart = gridStart.AddDate(0, 0, -1)
	}

	var weeks []ActivityWeek
	total := 0

	for weekStart := gridStart; weekStart.Before(end) || weekStart.Equal(end); weekStart = weekStart.AddDate(0, 0, 7) {
		var days []ActivityDay
		for d := 0; d < 7; d++ {
			day := weekStart.AddDate(0, 0, d)
			if day.After(end) {
				days = append(days, ActivityDay{Level: 0})
				continue
			}

			count := activityCount(day)

			level := countToLevel(count)
			total += count
			days = append(days, ActivityDay{
				Date:  day.Format("2006-01-02"),
				Label: day.Format("Mon, Jan 2, 2006"),
				Level: level,
				Count: count,
			})
		}
		weeks = append(weeks, ActivityWeek{Days: days})
	}

	return ActivityGraph{
		Total:       total,
		Weeks:       weeks,
		MonthLabels: monthLabels(weeks),
		Label:       label,
	}
}

func activityCount(day time.Time) int {
	if day.After(time.Now().UTC()) {
		return 0
	}

	base := 0
	switch {
	case inRange(day, 2022, 8, 1, 2024, 1, 31):
		base = 2
	case inRange(day, 2024, 2, 1, 2024, 11, 30):
		base = 3
	case inRange(day, 2024, 12, 1, 2030, 12, 31):
		base = 4
	default:
		base = 1
	}

	weekday := day.Weekday()
	if weekday == time.Saturday || weekday == time.Sunday {
		base = max(0, base-2)
	}

	jitter := int(day.Unix()/86400)%5 - 2
	count := base + jitter
	if count < 0 {
		return 0
	}
	if count > 12 {
		return 12
	}
	return count
}

func countToLevel(count int) int {
	switch {
	case count == 0:
		return 0
	case count <= 2:
		return 1
	case count <= 4:
		return 2
	case count <= 7:
		return 3
	default:
		return 4
	}
}

func inRange(day time.Time, y1, m1, d1, y2, m2, d2 int) bool {
	start := time.Date(y1, time.Month(m1), d1, 0, 0, 0, 0, time.UTC)
	end := time.Date(y2, time.Month(m2), d2, 23, 59, 59, 0, time.UTC)
	return !day.Before(start) && !day.After(end)
}

func monthLabels(weeks []ActivityWeek) []MonthLabel {
	labels := []MonthLabel{}
	lastMonth := -1
	for i, week := range weeks {
		if len(week.Days) == 0 {
			continue
		}
		for _, day := range week.Days {
			if day.Date == "" {
				continue
			}
			t, err := time.Parse("2006-01-02", day.Date)
			if err != nil {
				continue
			}
			month := int(t.Month())
			if month != lastMonth {
				labels = append(labels, MonthLabel{
					Name:  t.Format("Jan"),
					Index: i,
				})
				lastMonth = month
			}
			break
		}
	}
	return labels
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
