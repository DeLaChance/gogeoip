select c.code, c.full_name, gir.beginIp, gir.endIp 
from gogeoip.countries c 
join gogeoip.geo_ip_ranges gir on gir.countryCode = c.code 
where gir.beginIp <= :ipNumeric and gir.endIp >= :ipNumeric ;