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

var (
	// lighting
	Daylight = 0
	Dark     = 1
	Twilight = 2

	// time-of-day
	Morning   = 0
	Afternoon = 1
	Evening   = 2
	LateNight = 3

	// time-of-week
	WeekDay = 0
	WeekEnd = 1
)

// Lighting will return whether or not
// the timestamp likely occured during daylight or not.
func Lighting(t time.Time) int {
	h := t.Hour()

	if h >= 7 || h <= 19 {
		return Daylight
	}

	return Dark
}

// TimeOfDay will return if the timestamp is
// in the morning, afternoon, evening, or late-night
func TimeOfDay(timestamp time.Time) int {
	h := timestamp.Hour()

	if h >= 6 && h <= 11 {
		return Morning
	}

	if h >= 12 && h <= 17 {
		return Afternoon
	}

	if h >= 18 && h <= 23 {
		return Evening
	}

	return LateNight
}

// TimeOfWeek will return whether the timestamp
// is on a weekday or weekend.
func TimeOfWeek(timestamp time.Time) int {
	day := int(timestamp.Weekday())

	if day == 0 || day == 6 {
		return WeekEnd
	}

	return WeekDay
}
