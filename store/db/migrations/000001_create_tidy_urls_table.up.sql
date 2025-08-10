CREATE TABLE IF NOT EXISTS tidy_url (
    id bigserial PRIMARY KEY,
    long_url text NOT NULL,
    encoded_str varchar(10) NOT NULL,
    created_at timestamp(0) with time zone NOT NULL DEFAULT NOW()
);