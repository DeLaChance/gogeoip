# gogeoip applicaiton

A simple backend for queries mapping an IPv4 address to a location (e.g. country) via a REST JSON api.

## TODO:

- Parameter use in query
- SQL to file
- Solve TODO's backend
- Instructions on how to build

# Rest API:

GET:
[http://localhost:8080/api/geoip/:ipAddress/country](http://localhost:8080/api/geoip/:ipAddress/country).

Output for example:

```
{
  "code": "ES",
  "fullName": "Spain"
}
```
