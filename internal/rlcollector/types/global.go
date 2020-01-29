// River Life
// Copyright (C) 2020  Denny Chambers

// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.

// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.

// You should have received a copy of the GNU General Public License
// along with this progrsm.  If not, see <http://www.gnu.org/licenses/>.
package types

import (
	appctx "riverlife/internal/rlcollector/appcontext"
)

var StateURL string = "https://water.weather.gov/ahps2/rss/obs/%s.rss"
var SiteURL string = "https://water.weather.gov/ahps2/hydrograph_to_xml.php?gage=%s&output=xml"

var Ctx *appctx.AppContext = nil
