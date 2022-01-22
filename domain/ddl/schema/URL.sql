CREATE TABLE URL (
	"url_id"                   INTEGER NOT NULL PRIMARY KEY,		
	"long_url"                 TEXT NOT NULL UNIQUE,
	"short_url"                TEXT NOT NULL,
    "expiration_dt"            INTEGER
);