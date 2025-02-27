package request

type Developer struct {
	Name         string `json:"name"`
	Capacity     int    `json:"capacity"`     // Capacity in terms of difficulty per hour
	WeeklyHours  int    `json:"weeklyHours"`  // Total hours available per week
	CurrentHours int    `json:"currentHours"` // Remaining hours this week
}

type Developers struct {
	Developers []Developer `json:"developers"`
}
