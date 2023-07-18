# Task manager

## Migrations

```sql
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Task manager jobs
CREATE TABLE IF NOT EXISTS gue_jobs
(
    job_id      TEXT        NOT NULL PRIMARY KEY,
    priority    SMALLINT    NOT NULL,
    run_at      TIMESTAMPTZ NOT NULL,
    job_type    TEXT        NOT NULL,
    args        BYTEA       NOT NULL,
    error_count INTEGER     NOT NULL DEFAULT 0,
    last_error  TEXT,
    queue       TEXT        NOT NULL,
    created_at  TIMESTAMPTZ NOT NULL,
    updated_at  TIMESTAMPTZ NOT NULL,
    UNIQUE (job_type, queue, args)
);
CREATE INDEX IF NOT EXISTS gue_jobs_selector_idx ON gue_jobs USING btree (queue, run_at, priority);
CREATE INDEX IF NOT EXISTS gue_jobs_selector2_idx ON gue_jobs USING btree (queue, priority);
CREATE INDEX IF NOT EXISTS gue_jobs_queue_idx ON gue_jobs USING btree (queue);
```
