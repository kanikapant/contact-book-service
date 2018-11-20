CREATE TABLE IF NOT EXISTS profiling_service_db.endpoint_category_result (
    "firstname" text,
    "lastname" text,
    "contact_no" map<text,bigint>,
    "emailid" text,
    "birthday" date,
    "address" text,
    "updated_at" bigint,
     PRIMARY KEY (firstname, emailid)
  );
