CREATE TABLE IF NOT EXISTS buckets (
    id bigserial PRIMARY KEY,
    name text NOT NULL,
    balance bigint NOT NULL, -- stored in cents
    created_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    _version integer NOT NULL DEFAULT 1
);
