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

// Season holds the name and starting and ending times for the season in a given year.
type Season struct {
	Name  string
	Start time.Time
	End   time.Time
}

// Seasons holds a year and a list of Seasons for that year
type Seasons struct {
	Year int
	List []Season
}

// YearSeasons can be used as a list of seasons for a given year.
// Optimal usage of seasons is to simply make(YearSeasons) and call GetSeason() method on timestamps.
type YearSeasons map[int]*Seasons

// GetSeason will get the season for the given timestamp.
// If the list of seasons for the timestamps given year have not yet been
// generated then they will be generated and utilized to determine the season
// of the given timestamp.
func (y YearSeasons) GetSeason(t time.Time) Season {
	year := t.Year()
	seasons, ok := y[year]
	if !ok {
		y[year] = ListSeasons(t)
		seasons = y[year]
	}

	return seasons.GetSeason(t)
}

// GetSeason ...
// "Winter", "Spring", "Summer", or "Fall"
// All you have to do is call (this function).
// https://www.youtube.com/watch?v=eAR_Ff5A8Rk
func (s Seasons) GetSeason(t time.Time) Season {
	var season Season
	for _, ssn := range s.List {
		if t.After(ssn.Start) && ssn.End.After(t) {
			season = ssn
		}
	}

	return season
}

// ListSeasons will generate a list of seasons for the given year of the timestamp
// It is recommended that you store the list and reuse it for processing multiple timestamps.
func ListSeasons(t time.Time) *Seasons {
	year := t.Year()
	location := t.Location()

	seasons := &Seasons{
		Year: year,
	}

	winter1 := Season{
		Name:  "winter",
		Start: DayStart(time.Date(year, time.January, 1, 0, 0, 0, 0, location)),
		End:   DayFinish(time.Date(year, time.March, 20, 0, 0, 0, 0, location)),
	}
	seasons.List = append(seasons.List, winter1)
	spring := Season{
		Name:  "spring",
		Start: DayStart(time.Date(year, time.March, 21, 0, 0, 0, 0, location)),
		End:   DayFinish(time.Date(year, time.June, 20, 0, 0, 0, 0, location)),
	}
	seasons.List = append(seasons.List, spring)
	summer := Season{
		Name:  "summer",
		Start: DayStart(time.Date(year, time.June, 21, 0, 0, 0, 0, location)),
		End:   DayFinish(time.Date(year, time.September, 22, 0, 0, 0, 0, location)),
	}
	seasons.List = append(seasons.List, summer)
	fall := Season{
		Name:  "fall",
		Start: DayStart(time.Date(year, time.September, 23, 0, 0, 0, 0, location)),
		End:   DayFinish(time.Date(year, time.December, 20, 0, 0, 0, 0, location)),
	}
	seasons.List = append(seasons.List, fall)
	winter2 := Season{
		Name:  "winter",
		Start: DayStart(time.Date(year, time.December, 21, 0, 0, 0, 0, location)),
		End:   DayFinish(time.Date(year, time.December, 31, 0, 0, 0, 0, location)),
	}
	seasons.List = append(seasons.List, winter2)

	return seasons
}
