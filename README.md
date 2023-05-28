# gogeoip applicaiton

A simple backend for queries mapping an IPv4 address to a location (e.g. country) via a REST JSON api.

## TODO:

- Return original in query

- SQL to file
- Solve remaining TODO's backend

- Instructions on how to build

# Rest API:

GET:
[http://localhost:8080/api/geoip/8.8.8.8/country]http://localhost:8080/api/geoip/8.8.8.8/country).

Output for example:

```
{
  "code": "US",
  "fullName": "United States of America"
}
```
