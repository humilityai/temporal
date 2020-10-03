// Copyright 2020 Hummility AI Incorporated, All Rights Reserved.
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
	FirstQuarter  = 1
	SecondQuarter = 2
	ThirdQuarter  = 3
	FourthQuarter = 4
)

// Quarter will return which yearly quarter the timestamp is in {1,2,3,4}.
// It will return 0 if there was an error in determining the quarter.
func Quarter(t time.Time) int {
	switch t.Month() {
	case time.January, time.February, time.March:
		return FirstQuarter
	case time.April, time.May, time.June:
		return SecondQuarter
	case time.July, time.August, time.September:
		return ThirdQuarter
	case time.October, time.November, time.December:
		return FourthQuarter
	}

	return 0
}

// QuarterStart will return the starting time of the quarter for the time.Time object
func QuarterStart(t time.Time) time.Time {
	quarter := Quarter(t)
	var start time.Time
	switch quarter {
	case FirstQuarter:
		start = time.Date(t.Year(), time.January, 1, 0, 0, 0, 0, t.Location())
	case SecondQuarter:
		start = time.Date(t.Year(), time.April, 1, 0, 0, 0, 0, t.Location())
	case ThirdQuarter:
		start = time.Date(t.Year(), time.July, 1, 0, 0, 0, 0, t.Location())
	case FourthQuarter:
		start = time.Date(t.Year(), time.October, 1, 0, 0, 0, 0, t.Location())
	}
	return start
}

// QuarterFinish will return the final time of the quarter for the time.Time object
func QuarterFinish(t time.Time) time.Time {
	return QuarterStart(t).AddDate(0, 3, 0).Add(-time.Nanosecond)
}
