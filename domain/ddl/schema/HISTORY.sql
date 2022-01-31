CREATE TABLE HISTORY (
	"history_id"                   INTEGER NOT NULL PRIMARY KEY,		
	"url_id"                       INTEGER NOT NULL,
    "access_dt"                    TEXT NOT NULL,

	FOREIGN KEY("url_id") REFERENCES URL("url_id")
);