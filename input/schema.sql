use gogeoip;

create table if not exists countries (
	code char(3) not null primary key,
	full_name varchar(100) not null
);

create table if not exists geo_ip_ranges (
	id int unsigned not null primary key auto_increment,
	beginIp int unsigned not null,
	endIp int unsigned not null,
	countryCode char(3) not null,
	constraint gipc_fk foreign key (`countryCode`) references countries(`code`)
);
