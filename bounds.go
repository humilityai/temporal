// Copyright 2020 Humility AI Incorporated, All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package temporal

import "time"

// HourStart will return the starting time of the day for the given time.Time object
func HourStart(t time.Time) time.Time {
	y, m, d := t.Date()
	return time.Date(y, m, d, t.Hour(), 0, 0, 0, t.Location())
}

// HourFinish will return the final time of the day for the given time.Time object
func HourFinish(t time.Time) time.Time {
	y, m, d := t.Date()
	return time.Date(y, m, d, t.Hour(), 59, 59, int(time.Second-time.Nanosecond), t.Location())
}

// DayStart will return the starting time of the day for the given time.Time object
func DayStart(t time.Time) time.Time {
	y, m, d := t.Date()
	return time.Date(y, m, d, 0, 0, 0, 0, t.Location())
}

// DayFinish will return the final time of the day for the given time.Time object
func DayFinish(t time.Time) time.Time {
	y, m, d := t.Date()
	return time.Date(y, m, d, 23, 59, 59, int(time.Second-time.Nanosecond), t.Location())
}

// WeekStart will return the starting time of the week for the given time.Time object
func WeekStart(t time.Time) time.Time {
	weekday := int(DayStart(t).Weekday())
	return t.AddDate(0, 0, -weekday)
}

// WeekFinish will return the final time of the week for the given time.Time object
func WeekFinish(t time.Time) time.Time {
	return WeekStart(t).AddDate(0, 0, 7).Add(-time.Nanosecond)
}

// MonthStart will return the starting time of the month for the given time.Time object
func MonthStart(t time.Time) time.Time {
	y, m, _ := t.Date()
	return time.Date(y, m, 1, 0, 0, 0, 0, t.Location())
}

// MonthFinish will return the final time of the month for the given time.Time object
func MonthFinish(t time.Time) time.Time {
	return MonthStart(t).AddDate(0, 1, 0).Add(-time.Nanosecond)
}

// YearStart will return the starting time of the year for the given time.Time object
func YearStart(t time.Time) time.Time {
	return time.Date(t.Year(), time.January, 1, 0, 0, 0, 0, t.Location())
}

// YearFinish will return the final time of the year for the given time.Time object
func YearFinish(t time.Time) time.Time {
	return YearStart(t).AddDate(1, 0, 0).Add(-time.Nanosecond)
}
