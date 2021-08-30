CREATE TABLE applications(
   id UUID NOT NULL PRIMARY KEY,
   first_name varchar(40) NOT NULL,
   last_name varchar(40) NOT NULL,
   status varchar(40) DEFAULT "created",
   created_at TIMESTAMPZ NOT NULL,
   updated_at TIMESTAMPZ NOT NULL
);

CREATE TABLE jobs(
   id UUID NOT NULL PRIMARY KEY,
   application_id UUID NOT NULL FOREIGN KEY,
   status varchar(20) DEFAULT "created",
   created_at TIMESTAMPZ NOT NULL,
   updated_at TIMESTAMPZ NOT NULL
);