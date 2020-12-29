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

// Date ...
type Date struct {
	Year  int
	Month time.Month
	Day   int
}

// Dates ...
type Dates []Date

// PriorDate ...
func PriorDate(timestamp time.Time) Date {
	priorTime := timestamp.AddDate(0, 0, -1)
	return Date{
		Year:  priorTime.Year(),
		Month: priorTime.Month(),
		Day:   priorTime.Day(),
	}
}

// NextDate ...
func NextDate(timestamp time.Time) Date {
	nextTime := timestamp.AddDate(0, 0, 1)
	return Date{
		Year:  nextTime.Year(),
		Month: nextTime.Month(),
		Day:   nextTime.Day(),
	}
}

// DatesBetween ...
func DatesBetween(start, end time.Time) Dates {
	var dates Dates

	dates = append(dates, Date{
		Year:  start.Year(),
		Month: start.Month(),
		Day:   start.Day(),
	})

	var nextDay time.Time
	for {
		nextDay = start.AddDate(0, 0, 1)
		if end.After(nextDay) {
			dates = append(dates, Date{
				Year:  nextDay.Year(),
				Month: nextDay.Month(),
				Day:   nextDay.Day(),
			})
		} else {
			// if final time is later in same day than end time: add final day to list of days
			if end.Year() == nextDay.Year() && end.Month() == nextDay.Month() && end.Day() == nextDay.Day() {
				dates = append(dates, Date{
					Year:  end.Year(),
					Month: end.Month(),
					Day:   end.Day(),
				})
			}

			break
		}
	}

	return dates
}
