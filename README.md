# gogeoip applicaiton

A simple backend for queries mapping an IPv4 address to a location (e.g. country) via a REST JSON api.

## TODO:

- SQL to file
- Solve remaining TODO's backend
- domain objects to separate dir

- Instructions on how to build

# Rest API:

GET:
[http://localhost:8080/api/geoip/8.8.8.8/country]http://localhost:8080/api/geoip/8.8.8.8/country).

Output for example:

```
{
  "query": {
    "ipAddress": "8.8.8.8",
    "ipAddressNumeric": 134744072
  },
  "country": {
    "code": "US",
    "fullName": "United States of America"
  }
}
```
