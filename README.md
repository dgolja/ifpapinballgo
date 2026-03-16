# ifpapinballgo

A Go client library for the IFPA (International Flippper Pinball Association) [API](https://api.ifpapinball.com/docs/).

The Go client is generated from the *modified* [spec](ifpapinball-official-api.modified.yaml), not the official one.

The [official IFPA OpenAPI spec 2.1](https://api.ifpapinball.com/docs/) is incomplete (missing fields) and contains structural errors that are incompatible with code generation for statically typed languages. The overlay ([ifpapinball-overlay.yaml](ifpapinball-overlay.yaml)) patches these issues before generation.

The official spec (v2.1) has several issues. The most notable ones are:
- Structural errors: several fields that return JSON arrays are declared as `type: object` instead of `type: array`, which breaks Go code generation (e.g., tournaments, pvp, results, rankings, country_directors, rank_history).
- Missing type annotations: some numeric fields are returned by the API as strings (or inconsistently as both strings and numbers).
- Missing or misspelled fields.

I reached out to the IFPA to see if some of these issues can be fixed. Once they are resolved, I will update the client library.
