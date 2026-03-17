package client

import "testing"

func TestViewDirectorToursParamsTimePeriodValid(t *testing.T) {
	if !FUTURE.Valid() {
		t.Error("FUTURE.Valid() = false, want true")
	}
	if !PAST.Valid() {
		t.Error("PAST.Valid() = false, want true")
	}
	if ViewDirectorToursParamsTimePeriod("INVALID").Valid() {
		t.Error(`"INVALID".Valid() = true, want false`)
	}
}

func TestViewPlayerActiveResultsParamsRankingSystemValid(t *testing.T) {
	valid := []ViewPlayerActiveResultsParamsRankingSystem{
		ViewPlayerActiveResultsParamsRankingSystemMAIN,
		ViewPlayerActiveResultsParamsRankingSystemWOMEN,
		ViewPlayerActiveResultsParamsRankingSystemYOUTH,
	}
	for _, v := range valid {
		if !v.Valid() {
			t.Errorf("%q.Valid() = false, want true", v)
		}
	}
	if ViewPlayerActiveResultsParamsRankingSystem("INVALID").Valid() {
		t.Error(`"INVALID".Valid() = true, want false`)
	}
}

func TestViewPlayerActiveResultsParamsTypeValid(t *testing.T) {
	valid := []ViewPlayerActiveResultsParamsType{
		ViewPlayerActiveResultsParamsTypeACTIVE,
		ViewPlayerActiveResultsParamsTypeINACTIVE,
		ViewPlayerActiveResultsParamsTypeNONACTIVE,
	}
	for _, v := range valid {
		if !v.Valid() {
			t.Errorf("%q.Valid() = false, want true", v)
		}
	}
	if ViewPlayerActiveResultsParamsType("INVALID").Valid() {
		t.Error(`"INVALID".Valid() = true, want false`)
	}
}

func TestRankingProParamsRankingSystemValid(t *testing.T) {
	if !RankingProParamsRankingSystemOPEN.Valid() {
		t.Error("RankingProParamsRankingSystemOPEN.Valid() = false, want true")
	}
	if !RankingProParamsRankingSystemWOMEN.Valid() {
		t.Error("RankingProParamsRankingSystemWOMEN.Valid() = false, want true")
	}
	if RankingProParamsRankingSystem("INVALID").Valid() {
		t.Error(`"INVALID".Valid() = true, want false`)
	}
}

func TestRankingWomenOpenParamsTournamentTypeValid(t *testing.T) {
	if !RankingWomenOpenParamsTournamentTypeOPEN.Valid() {
		t.Error("RankingWomenOpenParamsTournamentTypeOPEN.Valid() = false, want true")
	}
	if !RankingWomenOpenParamsTournamentTypeWOMEN.Valid() {
		t.Error("RankingWomenOpenParamsTournamentTypeWOMEN.Valid() = false, want true")
	}
	if RankingWomenOpenParamsTournamentType("INVALID").Valid() {
		t.Error(`"INVALID".Valid() = true, want false`)
	}
}

func TestSeriesRegionsParamsSeriesCodeValid(t *testing.T) {
	valid := []SeriesRegionsParamsSeriesCode{ACS, NACS, WNACSO, WNACSW}
	for _, v := range valid {
		if !v.Valid() {
			t.Errorf("%q.Valid() = false, want true", v)
		}
	}
	if SeriesRegionsParamsSeriesCode("INVALID").Valid() {
		t.Error(`"INVALID".Valid() = true, want false`)
	}
}

func TestStatsCountryPlayerCountParamsRankTypeValid(t *testing.T) {
	if !StatsCountryPlayerCountParamsRankTypeOPEN.Valid() {
		t.Error("OPEN.Valid() = false, want true")
	}
	if !StatsCountryPlayerCountParamsRankTypeWOMEN.Valid() {
		t.Error("WOMEN.Valid() = false, want true")
	}
	if StatsCountryPlayerCountParamsRankType("INVALID").Valid() {
		t.Error(`"INVALID".Valid() = true, want false`)
	}
}

func TestStatsEventPeriodParamsRankTypeValid(t *testing.T) {
	if !StatsEventPeriodParamsRankTypeOPEN.Valid() {
		t.Error("OPEN.Valid() = false, want true")
	}
	if !StatsEventPeriodParamsRankTypeWOMEN.Valid() {
		t.Error("WOMEN.Valid() = false, want true")
	}
	if StatsEventPeriodParamsRankType("INVALID").Valid() {
		t.Error(`"INVALID".Valid() = true, want false`)
	}
}

func TestStatsEventsByYearParamsRankTypeValid(t *testing.T) {
	if !StatsEventsByYearParamsRankTypeOPEN.Valid() {
		t.Error("OPEN.Valid() = false, want true")
	}
	if !StatsEventsByYearParamsRankTypeWOMEN.Valid() {
		t.Error("WOMEN.Valid() = false, want true")
	}
	if StatsEventsByYearParamsRankType("INVALID").Valid() {
		t.Error(`"INVALID".Valid() = true, want false`)
	}
}

func TestStatsLargestTournamentsParamsRankTypeValid(t *testing.T) {
	if !StatsLargestTournamentsParamsRankTypeOPEN.Valid() {
		t.Error("OPEN.Valid() = false, want true")
	}
	if !StatsLargestTournamentsParamsRankTypeWOMEN.Valid() {
		t.Error("WOMEN.Valid() = false, want true")
	}
	if StatsLargestTournamentsParamsRankType("INVALID").Valid() {
		t.Error(`"INVALID".Valid() = true, want false`)
	}
}

func TestStatsLucrativeToursParamsRankTypeValid(t *testing.T) {
	if !StatsLucrativeToursParamsRankTypeOPEN.Valid() {
		t.Error("OPEN.Valid() = false, want true")
	}
	if !StatsLucrativeToursParamsRankTypeWOMEN.Valid() {
		t.Error("WOMEN.Valid() = false, want true")
	}
	if StatsLucrativeToursParamsRankType("INVALID").Valid() {
		t.Error(`"INVALID".Valid() = true, want false`)
	}
}

func TestStatsOverallParamsSystemCodeValid(t *testing.T) {
	if !StatsOverallParamsSystemCodeOPEN.Valid() {
		t.Error("OPEN.Valid() = false, want true")
	}
	if !StatsOverallParamsSystemCodeWOMEN.Valid() {
		t.Error("WOMEN.Valid() = false, want true")
	}
	if StatsOverallParamsSystemCode("INVALID").Valid() {
		t.Error(`"INVALID".Valid() = true, want false`)
	}
}

func TestStatsPointsPeriodParamsRankTypeValid(t *testing.T) {
	if !StatsPointsPeriodParamsRankTypeOPEN.Valid() {
		t.Error("OPEN.Valid() = false, want true")
	}
	if !StatsPointsPeriodParamsRankTypeWOMEN.Valid() {
		t.Error("WOMEN.Valid() = false, want true")
	}
	if StatsPointsPeriodParamsRankType("INVALID").Valid() {
		t.Error(`"INVALID".Valid() = true, want false`)
	}
}

func TestStatsStatePlayerCountParamsRankTypeValid(t *testing.T) {
	if !StatsStatePlayerCountParamsRankTypeOPEN.Valid() {
		t.Error("OPEN.Valid() = false, want true")
	}
	if !StatsStatePlayerCountParamsRankTypeWOMEN.Valid() {
		t.Error("WOMEN.Valid() = false, want true")
	}
	if StatsStatePlayerCountParamsRankType("INVALID").Valid() {
		t.Error(`"INVALID".Valid() = true, want false`)
	}
}

func TestStatsStateTourCountParamsRankTypeValid(t *testing.T) {
	if !StatsStateTourCountParamsRankTypeOPEN.Valid() {
		t.Error("OPEN.Valid() = false, want true")
	}
	if !StatsStateTourCountParamsRankTypeWOMEN.Valid() {
		t.Error("WOMEN.Valid() = false, want true")
	}
	if StatsStateTourCountParamsRankType("INVALID").Valid() {
		t.Error(`"INVALID".Valid() = true, want false`)
	}
}

func TestViewLeagueInfoParamsTimePeriodValid(t *testing.T) {
	valid := []ViewLeagueInfoParamsTimePeriod{
		ViewLeagueInfoParamsTimePeriodACTIVE,
		ViewLeagueInfoParamsTimePeriodHISTORY,
		ViewLeagueInfoParamsTimePeriodUPCOMING,
	}
	for _, v := range valid {
		if !v.Valid() {
			t.Errorf("%q.Valid() = false, want true", v)
		}
	}
	if ViewLeagueInfoParamsTimePeriod("INVALID").Valid() {
		t.Error(`"INVALID".Valid() = true, want false`)
	}
}

func TestTourSearchParamsRankTypeValid(t *testing.T) {
	if !MAIN.Valid() {
		t.Error("MAIN.Valid() = false, want true")
	}
	if !WOMEN.Valid() {
		t.Error("WOMEN.Valid() = false, want true")
	}
	if TourSearchParamsRankType("INVALID").Valid() {
		t.Error(`"INVALID".Valid() = true, want false`)
	}
}

func TestTourSearchParamsEventTypeValid(t *testing.T) {
	if !Keague.Valid() {
		t.Error("Keague.Valid() = false, want true")
	}
	if !Tournament.Valid() {
		t.Error("Tournament.Valid() = false, want true")
	}
	if TourSearchParamsEventType("INVALID").Valid() {
		t.Error(`"INVALID".Valid() = true, want false`)
	}
}

func TestTourSearchParamsPreRegistrationValid(t *testing.T) {
	if !TourSearchParamsPreRegistrationN.Valid() {
		t.Error("N.Valid() = false, want true")
	}
	if !TourSearchParamsPreRegistrationY.Valid() {
		t.Error("Y.Valid() = false, want true")
	}
	if TourSearchParamsPreRegistration("INVALID").Valid() {
		t.Error(`"INVALID".Valid() = true, want false`)
	}
}

func TestTourSearchParamsOnlyWithResultsValid(t *testing.T) {
	if !TourSearchParamsOnlyWithResultsN.Valid() {
		t.Error("N.Valid() = false, want true")
	}
	if !TourSearchParamsOnlyWithResultsY.Valid() {
		t.Error("Y.Valid() = false, want true")
	}
	if TourSearchParamsOnlyWithResults("INVALID").Valid() {
		t.Error(`"INVALID".Valid() = true, want false`)
	}
}

func TestTourSearchParamsDistanceUnitValid(t *testing.T) {
	if !KForKilometers.Valid() {
		t.Error("KForKilometers.Valid() = false, want true")
	}
	if !MForMiles.Valid() {
		t.Error("MForMiles.Valid() = false, want true")
	}
	if TourSearchParamsDistanceUnit("INVALID").Valid() {
		t.Error(`"INVALID".Valid() = true, want false`)
	}
}

func TestTourSearchParamsPointFilterValid(t *testing.T) {
	if !NDefault.Valid() {
		t.Error("NDefault.Valid() = false, want true")
	}
	if !Y.Valid() {
		t.Error("Y.Valid() = false, want true")
	}
	if TourSearchParamsPointFilter("INVALID").Valid() {
		t.Error(`"INVALID".Valid() = true, want false`)
	}
}
