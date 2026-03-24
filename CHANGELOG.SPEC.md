# CHANGELOG

Current changes between [ifpapinball-official-api.yaml](./ifpapinball-official-api.yaml) and [ifpapinball-official-api.modified.yaml](./ifpapinball-official-api.modified.yaml).

---

## `/director/{id}`

### x-go-type annotations
- `director_id`, `country_id` — added `x-go-type: types.StringInt`
- `stats.tournament_count`, `unique_location_count`, `women_tournament_count`, `league_count`, `total_player_count`, `unique_player_count`, `first_time_player_count`, `repeat_player_count`, `largest_event_count`, `single_format_count`, `multiple_format_count`, `unknown_format_count` — added `x-go-type: types.StringInt`
- `stats.highest_value`, `average_value` — added `x-go-type: types.StringFloat64`
- `stats.formats[].count` — added `x-go-type: types.StringInt`
- `player_id` (new field) — added `x-go-type: types.StringInt`

### Spec changes
- `stats.formats` — **fixed**: changed from `type: object` to `type: array` with proper `items` schema
- **Added** new response fields: `default_tournament_photo` (string), `player_id` (nullable number)
- Removed top-level `type: object` from response schema

---

## `/director/{id}/tournaments/{time_period}`

### x-go-type annotations
- `director_id` — added `x-go-type: types.StringInt`
- `tournaments[].tournament_id`, `player_count` — added `x-go-type: types.StringInt`

### Spec changes
- `tournaments` — **fixed**: changed from `type: object` to `type: array` with proper `items` schema

---

## `/director/country`

### x-go-type annotations
- `country_directors[].player_profile.player_id` — added `x-go-type: types.StringInt`

### Spec changes
- `country_directors` — **fixed**: changed from `type: object` to `type: array` with proper `items` schema
- Removed top-level `type: object` from response schema

---

## `/director/search`

### x-go-type annotations
- `directors[].director_id` — added `x-go-type: types.StringInt`
- `directors[].stats.event_count`, `future_event_count`, `unique_player_count`, `total_player_count` — added `x-go-type: types.StringInt`

### Spec changes
- All query parameters (`name`, `count`) — marked `required: false`
- `directors` — **fixed**: changed from `type: object` to `type: array` with proper `items` schema
- Removed top-level `type: object` from response schema

---

## `/other/countries`

### x-go-type annotations
- `country[].country_id` — added `x-go-type: types.StringInt`

### Spec changes
- Response restructured: flat top-level fields (`country_id`, `country_name`, `country_code`, `active_flag`) removed; replaced with a `country` array field

---

## `/other/stateprovs`

### x-go-type annotations
- `stateprov[].country_id` — added `x-go-type: types.StringInt`

### Spec changes
- Response restructured: flat top-level fields removed; replaced with a `stateprov` array field
- `stateprov[].regions` — properly typed as `type: array` with `items`

---

## `/player` and `/player/{id}`

### x-go-type annotations
- `player[].player_id` — added `x-go-type: types.StringInt`
- `player[].excluded_flag`, `ifpa_registered`, `womens_flag` — added `x-go-type: types.StringBool`
- `player[].age` — added `x-go-type: types.FlexibleInt`
- `player[].virtual_player_flag` (new field) — added `x-go-type: types.StringBool`
- `player[].matchplay_events.id`, `rating`, `rank` — added `x-go-type: types.StringInt`
- All rank/points fields within `open` and `womens` — added `x-go-type: types.StringInt` or `types.StringFloat64`
- `pro_rank` (new field within `open`/`womens`) — added `x-go-type: types.StringInt`
- `player_stats.years_active` — added `x-go-type: types.StringInt`
- `player[].series` — all numeric series fields annotated with `x-go-type`

### Spec changes
- `player` — **fixed**: changed from `type: object` to `type: array` with proper `items` schema
- **Added** new field: `player[].virtual_player_flag` (boolean)
- `player_stats.system` keys **renamed**: `MAIN` → `open`, `WOMEN` → `womens`
- **Added** new fields within `open`/`womens`: `highest_rank_date` (string), `pro_rank` (number)
- **Added** `player_stats.years_active` (number)
- `player[].series` — **fixed**: changed from `type: object` to `type: array` with proper `items` schema
- Removed top-level `type: object` from response schema

---

## `/player/{id}/pvp`

### x-go-type annotations
- `pvp[].player_id`, `win_count`, `loss_count`, `tie_count`, `current_rank` — added `x-go-type: types.StringInt`

### Spec changes
- `pvp` — **fixed**: changed from `type: object` to `type: array` with proper `items` schema
- Removed `content: {}` from `400` and `404` responses
- Removed top-level `type` from response schema

---

## `/player/{id}/pvp/{id2}`

### x-go-type annotations
- `player_1.player_id`, `player_2.player_id` — added `x-go-type: types.StringInt`
- `player_1.excluded_flag`, `player_2.excluded_flag` — added `x-go-type: types.StringBool`
- `pvp[].tournament_id` — added `x-go-type: types.StringInt`
- `pvp[].finish_position.player_1`, `player_2` — added `x-go-type: types.StringInt`

### Spec changes
- `pvp` — **fixed**: changed from `type: object` to `type: array` with proper `items` schema
- `pvp[].event_end_date` → **renamed** to `event_end_dt`
- Removed top-level `type` from response schema

---

## `/player/{id}/rank_history`

### x-go-type annotations
- `active_flag` — added `x-go-type: types.StringBool`
- `rank_history[].rank_position` — added `x-go-type: types.StringInt`
- `rank_history[].wppr_points` — added `x-go-type: types.StringFloat64`
- `rank_history[].tournaments_played_count` — added `x-go-type: types.StringInt`
- `rating_history[].rating` — added `x-go-type: types.StringFloat64`

### Spec changes
- `rank_history` — **fixed**: changed from `type: object` to `type: array` with proper `items` schema
- `rank_history[].rank_postiion` → **renamed/corrected** to `rank_position`
- `rating_history` — **fixed**: changed from `type: object` to `type: array` with proper `items` schema
- Removed top-level `type` from response schema

---

## `/player/{id}/results/{ranking_system}/{type}`

### x-go-type annotations
- `results[].tournament_id`, `position` — added `x-go-type: types.StringInt`
- `results[].original_points`, `current_points` — added `x-go-type: types.StringFloat64`

### Spec changes
- `results` — **fixed**: changed from `type: object` to `type: array` with proper `items` schema
- **Added** new response field: `rank_type` (string)
- Removed top-level `type` from response schema

---

## `/player/search`

### x-go-type annotations
- `results[].player_id`, `wppr_rank` — added `x-go-type: types.StringInt`
- `results[].rating_value` — added `x-go-type: types.StringFloat64`
- `total_results` — added `x-go-type: types.StringInt`

### Spec changes
- All query parameters (`name`, `country`, `stateprov`, `tournament`, `tourpos`) — marked `required: false`
- `results` — **fixed**: changed from `type: object` to `type: array` with proper `items` schema
- `total_results` — removed `example`
- `count` — removed `example`
- Removed top-level `type` from response schema

---

## `/rankings/youth`

### x-go-type annotations
- `total_count` — added `x-go-type: types.StringInt`
- `rankings[].player_id`, `current_wppr_rank`, `last_month_rank`, `rating_deviation`, `event_count`, `best_finish_position`, `best_tournament_id` — added `x-go-type: types.StringInt`
- `rankings[].wppr_points`, `rating_value`, `efficiency_percent` — added `x-go-type: types.StringFloat64`

### Spec changes
- Parameters (`start_pos`, `count`) — marked `required: false`
- `rankings` — **fixed**: changed from `type: object` to `type: array` with proper `items` schema
- `rankings[].best_finish` — **type changed** from `number` to `string` (tournament name, not position number)
- Removed top-level `type` from response schema

---

## `/rankings/wppr`

### x-go-type annotations
- `total_count` — added `x-go-type: types.StringInt`
- `rankings[].player_id`, `current_rank`, `last_month_rank`, `rating_deviation`, `event_count`, `best_finish_position`, `best_tournament_id`, `total_wins_last_3_years`, `top_3_last_3_years`, `top_10_last_3_years` — added `x-go-type: types.StringInt`
- `rankings[].age` — added `x-go-type: types.FlexibleInt`
- `rankings[].wppr_points`, `rating_value`, `efficiency_percent` — added `x-go-type: types.StringFloat64`

### Spec changes
- Parameters (`start_pos`, `count`) — marked `required: false`
- `rankings` — **fixed**: changed from `type: object` to `type: array` with proper `items` schema
- **Added** `sort_order` field to response
- `rankings[].best_finish` — **type changed** from `number` to `string` (tournament name)
- **Added** new fields: `total_wins_last_3_years`, `top_3_last_3_years`, `top_10_last_3_years`
- Removed top-level `type` from response schema

---

## `/rankings/virtual`

### x-go-type annotations
- `total_count` — added `x-go-type: types.StringInt`
- `rankings` — same fields annotated as `/rankings/wppr`

### Spec changes
- Parameters (`start_pos`, `count`) — marked `required: false`
- `rankings` — **fixed**: changed from `type: object` to `type: array` with proper `items` schema
- **Added** `sort_order` field; **added** `total_wins_last_3_years`, `top_3_last_3_years`, `top_10_last_3_years`
- Removed top-level `type` from response schema

---

## `/rankings/women/{tournament_type}`

### x-go-type annotations
- `total_count` — added `x-go-type: types.StringInt`
- `rankings[].player_id`, `current_wppr_rank`, `last_month_wppr_rank`, `rating_deviation`, `event_count`, `best_finish_position`, `best_tournament_id` — added `x-go-type: types.StringInt`
- `rankings[].age` — added `x-go-type: types.FlexibleInt`
- `rankings[].current_rank` — added `x-go-type: types.FlexibleInt`
- `rankings[].wppr_points`, `rating_value`, `efficiency_percent` — added `x-go-type: types.StringFloat64`

### Spec changes
- Parameters (`start_pos`, `count`) — marked `required: false`
- `rankings` — **fixed**: changed from `type: object` to `type: array` with proper `items` schema
- Removed top-level `type` from response schema

---

## `/rankings/pro/{ranking_system}`

### x-go-type annotations
- `rankings[].player_id`, `current_rank` — added `x-go-type: types.StringInt`
- `rankings[].pro_points`, `orginal_wppr_points`, `efficiency_percent`, `adj_efficiency_percent`, `excess_percent`, `wpprtunity`, `wppr_adjustment` — added `x-go-type: types.StringFloat64`

### Spec changes
- `rankings` — **fixed**: changed from `type: object` to `type: array` with proper `items` schema
- Removed top-level `type` from response schema

---

## `/rankings/country`

### x-go-type annotations
- `total_count` — added `x-go-type: types.StringInt`
- `rankings[].player_id`, `current_wppr_rank`, `last_month_rank`, `rating_deviation`, `event_count`, `best_finish_position`, `best_tournament_id` — added `x-go-type: types.StringInt`
- `rankings[].age` — added `x-go-type: types.FlexibleInt`
- `rankings[].wppr_points`, `rating`, `efficiency_percent` — added `x-go-type: types.StringFloat64`

### Spec changes
- Parameters (`start_pos`, `count`) — marked `required: false`
- **Added** new response field: `rank_country_name` (string)
- `rankings` — **fixed**: changed from `type: object` to `type: array` with proper `items` schema
- **Added** `profile_photo_url` field alongside `profile_photo`
- Removed top-level `type` from response schema

---

## `/rankings/country_list`

### x-go-type annotations
- `country[].player_count` — added `x-go-type: types.StringInt`

### Spec changes
- Response restructured: flat top-level fields (`country_name`, `country_code`, `player_count`) removed; replaced with a `country` array field and added `count`
- Removed top-level `type` from response schema

---

## `/rankings/custom/list`

### x-go-type annotations
- `custom_view[].view_id` — added `x-go-type: types.StringInt`

### Spec changes
- `custom_view` — **fixed**: changed from `type: object` to `type: array` with proper `items` schema
- Removed top-level `type` from response schema

---

## `/rankings/custom/{id}`

### x-go-type annotations
- `view_id` — added `x-go-type: types.StringInt`
- `view_results[].player_id`, `wppr_rank`, `event_count`, `position` — added `x-go-type: types.StringInt`
- `view_results[].wppr_points` — added `x-go-type: types.StringFloat64`

### Spec changes
- Parameters (`start_pos`, `count`) — marked `required: false`
- **Added** new response fields:
  - `view_results` — array with `player_id`, `name`, `country_code`, `country_name`, `city`, `stateprov`, `wppr_rank`, `wppr_points`, `event_count`, `position`
  - `tournaments` — array with `tournament_id`, `tournament_name`, `event_name`, `event_date`, `city`
  - `view_filters` — array with `name`, `setting`
- Removed `custom_view` field (**renamed/restructured**)
- Removed top-level `type` from response schema

---

## `/series/list`

### Spec changes
- `series` — **fixed**: changed from `type: object` to `type: array` with proper `items` schema; `years` sub-field properly typed as `type: array` of strings
- Removed top-level `type` from response schema

---

## `/series/{series_code}/regions`

### Spec changes
- Parameter `year` — marked `required: false`; type changed from `number` to `integer`
- Response `year` field — type changed from `number` to `integer`
- `active_regions` — **fixed**: changed from `type: object` to `type: array` with proper `items` schema
- Removed top-level `type` from response schema

---

## `/series/{series_code}/player_card/{player_id}`

### x-go-type annotations
- `player_id` — added `x-go-type: types.StringInt`
- `player_card[].tournament_id`, `region_event_rank` — added `x-go-type: types.StringInt`
- `player_card[].wppr_points` — added `x-go-type: types.StringFloat64`

### Spec changes
- Parameter `year` — marked `required: false`
- `series_code` — **fixed**: type changed from implied object to `string`
- `player_card` — **fixed**: changed from `type: object` to `type: array` with proper `items` schema
- Removed top-level `type` from response schema

---

## `/series/{series_code}/region_reps`

### x-go-type annotations
- `representative[].player_key`, `player_id` — added `x-go-type: types.StringInt`

### Spec changes
- `series_code` — **fixed**: type changed to `string`
- `representative` — **fixed**: changed from `type: object` to `type: array` with proper `items` schema
- Removed top-level `type` from response schema

---

## `/series/{series_code}/overall_standings`

### x-go-type annotations
- `overall_results[].player_count`, `unique_player_count`, `tournament_count` — added `x-go-type: types.StringInt`
- `overall_results[].current_leader.player_id` — added `x-go-type: types.StringInt`

### Spec changes
- Parameter `year` — marked `required: false`
- **Added** new response field: `championship_prize_fund` (number)
- `overall_results` — **fixed**: changed from `type: object` to `type: array` with proper `items` schema
- Removed top-level `type` from response schema

---

## `/series/{series_code}/standings`

### x-go-type annotations
- `prize_fund` — added `x-go-type: types.StringFloat64`
- `standings[].player_id`, `event_count`, `win_count` — added `x-go-type: types.StringInt`
- `standings[].wppr_points` — added `x-go-type: types.StringFloat64`

### Spec changes
- Parameter `year` — marked `required: false`
- `standings` — **fixed**: changed from `type: object` to `type: array` with proper `items` schema
- Removed top-level `type` from response schema

---

## `/series/{series_code}/stats`

### x-go-type annotations
- `monthly_stats[].month`, `tournament_count`, `player_count`, `unique_player_count` — added `x-go-type: types.StringInt`
- `monthly_stats[].prize_fund` — added `x-go-type: types.StringFloat64`
- `yearly_stats.player_count`, `unique_player_count` — added `x-go-type: types.StringInt`

### Spec changes
- Parameter `year` — marked `required: false`
- `monthly_stats` — **fixed**: changed from `type: object` to `type: array` with proper `items` schema
- `payouts` — **fixed**: changed from `type: object` to `type: array` with proper `items` schema
- Removed top-level `type` from response schema

---

## `/series/{series_code}/tournaments`

### x-go-type annotations
- `submitted_tournaments[].tournament_id`, `player_count` — added `x-go-type: types.StringInt`

### Spec changes
- Parameter `year` — marked `required: false`
- `submitted_tournaments` — **fixed**: changed from `type: object` to `type: array` with proper `items` schema; includes `winner` sub-object
- `unsubmitted_tournaments` — **fixed**: changed from `type: object` to `type: array` with proper `items` schema
- `future_tournament` — **fixed**: changed from `type: object` to `type: array` with proper `items` schema
- Removed top-level `type` from response schema

---

## `/stats/country_players`

### x-go-type annotations
- `stats[].player_count` — added `x-go-type: types.StringInt`

### Spec changes
- Parameter `count` — marked `required: false`
- `stats` — **fixed**: changed from `type: object` to `type: array` with proper `items` schema

---

## `/stats/state_players`

### x-go-type annotations
- `stats[].player_count` — added `x-go-type: types.StringInt`

### Spec changes
- Parameter `count` — marked `required: false`
- `stats` — **fixed**: changed from `type: object` to `type: array` with proper `items` schema

---

## `/stats/state_tournaments`

### x-go-type annotations
- `stats[].tournament_count` — added `x-go-type: types.StringInt`
- `stats[].total_points_all`, `total_points_tournament_value` — added `x-go-type: types.StringFloat64`

### Spec changes
- Parameter `count` — marked `required: false`
- `stats` — **fixed**: changed from `type: object` to `type: array` with proper `items` schema

---

## `/stats/events_by_year`

### x-go-type annotations
- `stats[].year`, `country_count`, `player_count`, `tournament_count` — added `x-go-type: types.StringInt`

### Spec changes
- Parameters (`count`, `year`) — marked `required: false`
- `stats` — **fixed**: changed from `type: object` to `type: array` with proper `items` schema

---

## `/stats/players_by_year`

### x-go-type annotations
- `stats[].year`, `current_year_count`, `previous_year_count`, `previous_2_year_count` — added `x-go-type: types.StringInt`

### Spec changes
- `stats` — **fixed**: changed from `type: object` to `type: array` with proper `items` schema

---

## `/stats/largest_tournaments`

### x-go-type annotations
- `stats[].player_count`, `tournament_id` — added `x-go-type: types.StringInt`

### Spec changes
- Parameters (`start_pos`, `count`) — marked `required: false`
- `stats` — **fixed**: changed from `type: object` to `type: array` with proper `items` schema

---

## `/stats/lucrative_tournaments`

### x-go-type annotations
- `stats[].tournament_id` — added `x-go-type: types.StringInt`

### Spec changes
- Parameters (`start_pos`, `count`, `year`) — marked `required: false`
- `stats` — **fixed**: changed from `type: object` to `type: array` with proper `items` schema

---

## `/stats/points_given_period`

### x-go-type annotations
- `stats[].player_id` — added `x-go-type: types.StringInt`
- `stats[].wppr_points` — added `x-go-type: types.StringFloat64`

### Spec changes
- Parameters (`start_date`, `end_date`, `country`, `stateprov`, `count`) — marked `required: false`
- `stats` — **fixed**: changed from `type: object` to `type: array` with proper `items` schema

---

## `/stats/events_attended_period`

### x-go-type annotations
- `stats[].player_id`, `tournament_count` — added `x-go-type: types.StringInt`

### Spec changes
- Parameters (`start_date`, `end_date`, `country`, `stateprov`, `count`) — marked `required: false`
- `start_date` and `end_date` — added example values
- `stats` — **fixed**: changed from `type: object` to `type: array` with proper `items` schema

---

## `/stats/overall`

### Spec changes
- Parameter `year` — marked `required: false`
- `stats.age_gender` — **removed**; replaced with `stats.age` object containing explicit age bracket fields: `age_under_18`, `age_18_to_29`, `age_30_to_39`, `age_40_to_49`, `age_50_to_99`
- `stats.overall_player_count`, `active_player_count`, `tournament_count`, `tournament_count_last_month`, `tournament_count_this_year`, `tournament_player_count` — type changed from `number` to `integer`
- `stats.tournament_player_count_average` example corrected to decimal (`22.9`); age bracket examples corrected to decimal percentages

---

## `/tournament/formats`

### x-go-type annotations
- `qualifying_formats[].format_id` — added `x-go-type: types.StringInt`
- `finals_formats[].format_id` — added `x-go-type: types.StringInt`

### Spec changes
- `qualifying_formats` — **fixed**: changed from `type: object` to `type: array` with proper `items` schema
- `finals_formats` — **fixed**: changed from `type: object` to `type: array` with proper `items` schema

---

## `/tournament/{id}`

### x-go-type annotations
- `tournament_id`, `director_id`, `qualify_hours`, `eligible_player_count`, `player_count`, `player_limit`, `matchplay_id`, `games_to_win` — added `x-go-type: types.StringInt`
- `private_flag`, `qualify_flag`, `prestige_flag`, `unlimited_qualifying_flag` — added `x-go-type: types.StringBool`
- `latitude`, `longitude`, `ratings_strength`, `rankings_strength`, `base_value`, `tournament_percentage_grade`, `tournament_value`, `event_weight` — added `x-go-type: types.StringFloat64`

### Spec changes
- **Added** new fields: `prestige_flag`, `event_weight`, `games_to_win`, `unlimited_qualifying_flag`
- **Removed** field: `unlimited_qualify_flag` (replaced by `unlimited_qualifying_flag`)
- `event_name` example corrected; `details` example updated; `player_limit` example changed to `0`

---

## `/tournament/{id}/related`

### x-go-type annotations
- `tournament[].tournament_id` — added `x-go-type: types.StringInt`
- `tournament[].winner` — added `x-go-type: types.TourRelatedWinner`

### Spec changes
- `tournament` — **fixed**: changed from `type: object` to `type: array` with proper `items` schema
- `tournament[].tournament_id` — type changed to `string` (was `number`)
- **Added** `tournament[].winner` field (polymorphic: `""` or JSON object)
- Removed top-level `type` from response schema

---

## `/tournament/{id}/results`

### x-go-type annotations
- `tournament_id` — added `x-go-type: types.StringInt`
- `results[].age` — added `x-go-type: types.FlexibleInt`
- `results[].player_id` — added `x-go-type: types.StringInt`
- `results[].position`, `wppr_rank`, `player_tournament_count`, `wppr_pro_rank` — added `x-go-type: types.StringInt`
- `results[].points`, `ratings_value`, `efficiency_value`, `post_rank_pos`, `post_rating_value`, `post_efficiency_value`, `post_wppr_pro_rank` — added `x-go-type: types.StringFloat64`
- `results[].excluded_flag` — added `x-go-type: types.StringBool`

### Spec changes
- `results` — **fixed**: changed from `type: object` to `type: array` with significantly expanded `items` schema
- **Added** new result fields: `age`, `profile_photo`, `wppr_rank`, `ratings_value`, `excluded_flag`, `player_tournament_count`, `wppr_pro_rank`, `efficiency_value`, `post_rank_pos`, `post_rating_value`, `post_efficiency_value`, `post_wppr_pro_rank`
- `results[].player_id` — type changed to `string`

---

## `/tournament/leagues/{time_period}`

### x-go-type annotations
- `leagues[].private_flag` — added `x-go-type: types.StringBool`
- `leagues[].director_id` — added `x-go-type: types.StringInt`

### Spec changes
- **Added** new response field: `league_status` (string)
- **Added** `leagues` array field with full tournament/league object schema
- **Removed** fields: `status`, `results`

---

## `/tournament/search`

### x-go-type annotations
- `total_results` — added `x-go-type: types.StringInt`
- `tournaments[].tournament_id`, `player_count`, `director_id` — added `x-go-type: types.StringInt`
- `tournaments[].private_flag`, `winner.excluded_flag` — added `x-go-type: types.StringBool`
- `tournaments[].winner.wppr_points` — added `x-go-type: types.StringFloat64`

### Spec changes
- All 19 query parameters — marked `required: false`
- Parameter `end_date` description updated; `type` enum corrected (typo `Keague`)
- `search_filter.start_date` and `end_date` — updated examples, removed `format: date`
- `total_results` — type changed from implicit to `number`
- `tournaments` — **fixed**: changed from `type: object` to `type: array` with significantly expanded `items` schema
- **Added** new tournament search result fields: `event_type`, `address1`, `address2`, `postal_code`, `event_start_date`, `latitude`, `longitude`, `raw_address`, `preregistration_date`, `qualifying_format`, `finals_format`, `director_id`, `director_name`, `website`, `details`, `profile_photo`, `certified_flag`, `winner` (object with `player_id`, `player_name`, `wppr_points`, `profile_photo`, `excluded_flag`, `country_cd`)

---

## History

### 2026-03-24

- `/series/{series_code}/regions`: `year` parameter and response field type changed from `number` to `integer`
- `/stats/overall`: count fields (`overall_player_count`, `active_player_count`, `tournament_count`, `tournament_count_last_month`, `tournament_count_this_year`, `tournament_player_count`) type changed from `number` to `integer`; corrected `tournament_player_count_average` example to decimal (`22.9`) and age bracket examples to decimal percentages
