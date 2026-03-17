# ifpapinballgo

A Go client library for the IFPA (International Flippper Pinball Association) [API](https://api.ifpapinball.com/docs/).

The Go client is generated from the *modified* [spec](ifpapinball-official-api.modified.yaml), not the official one.

The [official IFPA OpenAPI spec 2.1](https://api.ifpapinball.com/docs/) is incomplete (missing fields) and contains structural errors that are incompatible with code generation for statically typed languages. The overlay ([ifpapinball-overlay.yaml](ifpapinball-overlay.yaml)) patches these issues before generation.

The official spec (v2.1) has several issues. The most notable ones are:
- Structural errors: several fields that return JSON arrays are declared as `type: object` instead of `type: array`, which breaks Go code generation (e.g., tournaments, pvp, results, rankings, country_directors, rank_history).
- Missing type annotations: some numeric fields are returned by the API as strings (or inconsistently as both strings and numbers).
- Missing or misspelled fields.

I reached out to the IFPA to see if some of these issues can be fixed. Once they are resolved, I will update the client library. Full CHANGELOG available [here](./CHANGELOG.SPEC.md).

## Backwards compatibility guarantee

Whatever is committed to main should work, and I will try to keep changes backwards compatible. However, at this stage the Go client is still under development and even if it's mostly working there is still some work to do. For example, some dates are still returned as strings, which I plan to fix. Hopefully some other bugs found during development will also be fixed by the IFPA team, which means the client will change accordingly.
