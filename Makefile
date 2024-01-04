dbhost="mysql://fanchann:kalijaga@127.0.0.1:3306/wikipedia_api"



dbmate.up:
	dbmate --url $(dbhost) up
dbmate.down:
	dbmate --url $(dbhost) down