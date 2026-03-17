package client

import (
	"testing"
)

func TestAllParameterTypeConstants(t *testing.T) {
	tests := []struct {
		name     string
		constant interface{}
		expected string
	}{
		// ViewDirectorToursParamsTimePeriod
		{"FUTURE", FUTURE, "FUTURE"},
		{"PAST", PAST, "PAST"},

		// ViewPlayerActiveResultsParamsRankingSystem
		{"MAIN", ViewPlayerActiveResultsParamsRankingSystemMAIN, "MAIN"},
		{"WOMEN", ViewPlayerActiveResultsParamsRankingSystemWOMEN, "WOMEN"},
		{"YOUTH", ViewPlayerActiveResultsParamsRankingSystemYOUTH, "YOUTH"},

		// ViewPlayerActiveResultsParamsType
		{"ACTIVE", ViewPlayerActiveResultsParamsTypeACTIVE, "ACTIVE"},
		{"INACTIVE", ViewPlayerActiveResultsParamsTypeINACTIVE, "INACTIVE"},
		{"NONACTIVE", ViewPlayerActiveResultsParamsTypeNONACTIVE, "NONACTIVE"},

		// RankingProParamsRankingSystem
		{"RankingPro OPEN", RankingProParamsRankingSystemOPEN, "OPEN"},
		{"RankingPro WOMEN", RankingProParamsRankingSystemWOMEN, "WOMEN"},

		// RankingWomenOpenParamsTournamentType
		{"RankingWomenOpen OPEN", RankingWomenOpenParamsTournamentTypeOPEN, "OPEN"},
		{"RankingWomenOpen WOMEN", RankingWomenOpenParamsTournamentTypeWOMEN, "WOMEN"},

		// SeriesRegionsParamsSeriesCode
		{"ACS", ACS, "ACS"},
		{"NACS", NACS, "NACS"},
		{"WNACSO", WNACSO, "WNACSO"},
		{"WNACSW", WNACSW, "WNACSW"},

		// StatsCountryPlayerCountParamsRankType
		{"StatsCountryPlayerCount OPEN", StatsCountryPlayerCountParamsRankTypeOPEN, "OPEN"},
		{"StatsCountryPlayerCount WOMEN", StatsCountryPlayerCountParamsRankTypeWOMEN, "WOMEN"},

		// StatsEventPeriodParamsRankType
		{"StatsEventPeriod OPEN", StatsEventPeriodParamsRankTypeOPEN, "OPEN"},
		{"StatsEventPeriod WOMEN", StatsEventPeriodParamsRankTypeWOMEN, "WOMEN"},

		// StatsEventsByYearParamsRankType
		{"StatsEventsByYear OPEN", StatsEventsByYearParamsRankTypeOPEN, "OPEN"},
		{"StatsEventsByYear WOMEN", StatsEventsByYearParamsRankTypeWOMEN, "WOMEN"},

		// StatsLargestTournamentsParamsRankType
		{"StatsLargestTournaments OPEN", StatsLargestTournamentsParamsRankTypeOPEN, "OPEN"},
		{"StatsLargestTournaments WOMEN", StatsLargestTournamentsParamsRankTypeWOMEN, "WOMEN"},

		// StatsLucrativeToursParamsRankType
		{"StatsLucrativeTours OPEN", StatsLucrativeToursParamsRankTypeOPEN, "OPEN"},
		{"StatsLucrativeTours WOMEN", StatsLucrativeToursParamsRankTypeWOMEN, "WOMEN"},

		// StatsOverallParamsSystemCode
		{"StatsOverall OPEN", StatsOverallParamsSystemCodeOPEN, "OPEN"},
		{"StatsOverall WOMEN", StatsOverallParamsSystemCodeWOMEN, "WOMEN"},

		// StatsPointsPeriodParamsRankType
		{"StatsPointsPeriod OPEN", StatsPointsPeriodParamsRankTypeOPEN, "OPEN"},
		{"StatsPointsPeriod WOMEN", StatsPointsPeriodParamsRankTypeWOMEN, "WOMEN"},

		// StatsStatePlayerCountParamsRankType
		{"StatsStatePlayerCount OPEN", StatsStatePlayerCountParamsRankTypeOPEN, "OPEN"},
		{"StatsStatePlayerCount WOMEN", StatsStatePlayerCountParamsRankTypeWOMEN, "WOMEN"},

		// StatsStateTourCountParamsRankType
		{"StatsStateTourCount OPEN", StatsStateTourCountParamsRankTypeOPEN, "OPEN"},
		{"StatsStateTourCount WOMEN", StatsStateTourCountParamsRankTypeWOMEN, "WOMEN"},

		// ViewLeagueInfoParamsTimePeriod
		{"ACTIVE", ViewLeagueInfoParamsTimePeriodACTIVE, "ACTIVE"},
		{"HISTORY", ViewLeagueInfoParamsTimePeriodHISTORY, "HISTORY"},
		{"UPCOMING", ViewLeagueInfoParamsTimePeriodUPCOMING, "UPCOMING"},

		// TourSearchParamsRankType
		{"TourSearch MAIN", MAIN, "MAIN"},
		{"TourSearch WOMEN", WOMEN, "WOMEN"},

		// TourSearchParamsEventType
		{"Keague", Keague, "Keague"},
		{"Tournament", Tournament, "Tournament"},

		// TourSearchParamsPreRegistration
		{"TourSearchParamsPreRegistration N", TourSearchParamsPreRegistrationN, "N"},
		{"TourSearchParamsPreRegistration Y", TourSearchParamsPreRegistrationY, "Y"},

		// TourSearchParamsOnlyWithResults
		{"TourSearchParamsOnlyWithResults N", TourSearchParamsOnlyWithResultsN, "N"},
		{"TourSearchParamsOnlyWithResults Y", TourSearchParamsOnlyWithResultsY, "Y"},

		// TourSearchParamsDistanceUnit
		{"K for Kilometers", KForKilometers, "k for Kilometers"},
		{"M for Miles", MForMiles, "m for Miles"},

		// TourSearchParamsPointFilter
		{"N (default)", NDefault, "N (default)"},
		{"Y", Y, "Y"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			constantStr := ""
			switch v := tt.constant.(type) {
			case ViewDirectorToursParamsTimePeriod:
				constantStr = string(v)
			case ViewPlayerActiveResultsParamsRankingSystem:
				constantStr = string(v)
			case ViewPlayerActiveResultsParamsType:
				constantStr = string(v)
			case RankingProParamsRankingSystem:
				constantStr = string(v)
			case RankingWomenOpenParamsTournamentType:
				constantStr = string(v)
			case SeriesRegionsParamsSeriesCode:
				constantStr = string(v)
			case StatsCountryPlayerCountParamsRankType:
				constantStr = string(v)
			case StatsEventPeriodParamsRankType:
				constantStr = string(v)
			case StatsEventsByYearParamsRankType:
				constantStr = string(v)
			case StatsLargestTournamentsParamsRankType:
				constantStr = string(v)
			case StatsLucrativeToursParamsRankType:
				constantStr = string(v)
			case StatsOverallParamsSystemCode:
				constantStr = string(v)
			case StatsPointsPeriodParamsRankType:
				constantStr = string(v)
			case StatsStatePlayerCountParamsRankType:
				constantStr = string(v)
			case StatsStateTourCountParamsRankType:
				constantStr = string(v)
			case ViewLeagueInfoParamsTimePeriod:
				constantStr = string(v)
			case TourSearchParamsRankType:
				constantStr = string(v)
			case TourSearchParamsEventType:
				constantStr = string(v)
			case TourSearchParamsPreRegistration:
				constantStr = string(v)
			case TourSearchParamsOnlyWithResults:
				constantStr = string(v)
			case TourSearchParamsDistanceUnit:
				constantStr = string(v)
			case TourSearchParamsPointFilter:
				constantStr = string(v)
			default:
				t.Fatalf("unhandled constant type: %T", v)
			}

			if constantStr != tt.expected {
				t.Errorf("expected %q, got %q", tt.expected, constantStr)
			}
		})
	}
}

func TestParameterStructInitialization(t *testing.T) {
	t.Run("SearchDirectorsParams", func(t *testing.T) {
		// Test with all fields
		name := "Test Director"
		count := 100
		params := SearchDirectorsParams{
			Name:  &name,
			Count: &count,
		}

		if params.Name == nil || *params.Name != "Test Director" {
			t.Error("Name field not properly set")
		}
		if params.Count == nil || *params.Count != 100 {
			t.Error("Count field not properly set")
		}

		// Test with nil fields
		emptyParams := SearchDirectorsParams{}
		if emptyParams.Name != nil {
			t.Error("Name should be nil by default")
		}
		if emptyParams.Count != nil {
			t.Error("Count should be nil by default")
		}
	})

	t.Run("ViewPlayerMultiParams", func(t *testing.T) {
		params := ViewPlayerMultiParams{
			Players: "1,2,3,4,5",
		}

		if params.Players != "1,2,3,4,5" {
			t.Errorf("expected Players '1,2,3,4,5', got %q", params.Players)
		}
	})

	t.Run("SearchPlayersParams", func(t *testing.T) {
		name := "Test Player"
		country := "US"
		stateprov := "CA"
		tournament := "Test Tournament"
		tourpos := float32(5.5)

		params := SearchPlayersParams{
			Name:       &name,
			Country:    &country,
			Stateprov:  &stateprov,
			Tournament: &tournament,
			Tourpos:    &tourpos,
		}

		if params.Name == nil || *params.Name != "Test Player" {
			t.Error("Name field not properly set")
		}
		if params.Country == nil || *params.Country != "US" {
			t.Error("Country field not properly set")
		}
		if params.Stateprov == nil || *params.Stateprov != "CA" {
			t.Error("Stateprov field not properly set")
		}
		if params.Tournament == nil || *params.Tournament != "Test Tournament" {
			t.Error("Tournament field not properly set")
		}
		if params.Tourpos == nil || *params.Tourpos != 5.5 {
			t.Error("Tourpos field not properly set")
		}
	})

	t.Run("RankingCountryParams", func(t *testing.T) {
		startPos := float32(1.0)
		count := float32(100.0)

		params := RankingCountryParams{
			StartPos: &startPos,
			Count:    &count,
			Country:  "Canada",
		}

		if params.StartPos == nil || *params.StartPos != 1.0 {
			t.Error("StartPos field not properly set")
		}
		if params.Count == nil || *params.Count != 100.0 {
			t.Error("Count field not properly set")
		}
		if params.Country != "Canada" {
			t.Error("Country field not properly set")
		}
	})

	t.Run("RankingCustomIDParams", func(t *testing.T) {
		startPos := float32(10.0)
		count := float32(25.0)

		params := RankingCustomIDParams{
			StartPos: &startPos,
			Count:    &count,
		}

		if params.StartPos == nil || *params.StartPos != 10.0 {
			t.Error("StartPos field not properly set")
		}
		if params.Count == nil || *params.Count != 25.0 {
			t.Error("Count field not properly set")
		}
	})

	t.Run("SeriesRegionOverallStandingsParams", func(t *testing.T) {
		year := float32(2023)

		params := SeriesRegionOverallStandingsParams{
			Year: &year,
		}

		if params.Year == nil || *params.Year != 2023 {
			t.Error("Year field not properly set")
		}

		// Test with nil year
		emptyParams := SeriesRegionOverallStandingsParams{}
		if emptyParams.Year != nil {
			t.Error("Year should be nil by default")
		}
	})
}

func TestFloatParameterHandling(t *testing.T) {
	// Test that float32 parameters work correctly
	testValues := []float32{0.0, 1.0, 5.5, 100.5, -1.0}

	for _, val := range testValues {
		t.Run("float32 parameter", func(t *testing.T) {
			params := RankingCountryParams{
				StartPos: &val,
				Country:  "Test",
			}

			if params.StartPos == nil || *params.StartPos != val {
				t.Errorf("expected StartPos %f, got %f", val, *params.StartPos)
			}
		})
	}
}

func TestApiKeyScopes(t *testing.T) {
	if API_KEYScopes != "API_KEY.Scopes" {
		t.Errorf("expected API_KEYScopes to be 'API_KEY.Scopes', got %q", API_KEYScopes)
	}
}

func TestTypeStringConversions(t *testing.T) {
	// Test that custom types can be converted to strings
	tests := []struct {
		name     string
		value    interface{}
		expected string
	}{
		{"ViewDirectorToursParamsTimePeriod", FUTURE, "FUTURE"},
		{"ViewPlayerActiveResultsParamsRankingSystem", ViewPlayerActiveResultsParamsRankingSystemMAIN, "MAIN"},
		{"SeriesRegionsParamsSeriesCode", ACS, "ACS"},
		{"TourSearchParamsRankType", MAIN, "MAIN"},
		{"TourSearchParamsEventType", Tournament, "Tournament"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			str := ""
			switch v := tt.value.(type) {
			case ViewDirectorToursParamsTimePeriod:
				str = string(v)
			case ViewPlayerActiveResultsParamsRankingSystem:
				str = string(v)
			case SeriesRegionsParamsSeriesCode:
				str = string(v)
			case TourSearchParamsRankType:
				str = string(v)
			case TourSearchParamsEventType:
				str = string(v)
			default:
				t.Fatalf("unhandled type: %T", v)
			}

			if str != tt.expected {
				t.Errorf("expected %q, got %q", tt.expected, str)
			}
		})
	}
}

// Benchmark parameter initialization
func BenchmarkParameterInit(b *testing.B) {
	name := "Test Name"
	count := 50

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = SearchDirectorsParams{
			Name:  &name,
			Count: &count,
		}
	}
}

func BenchmarkComplexParameterInit(b *testing.B) {
	name := "Test Player"
	country := "US"
	stateprov := "CA"
	tournament := "World Championship"
	tourpos := float32(1.0)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = SearchPlayersParams{
			Name:       &name,
			Country:    &country,
			Stateprov:  &stateprov,
			Tournament: &tournament,
			Tourpos:    &tourpos,
		}
	}
}
