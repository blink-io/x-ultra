-- name: get-db-version
select version();

-- name: get-db-tsz
select now();

-- name: get-db-detail
select
    current_database(),
    current_schema(),
    inet_server_addr(),
    inet_server_port(),
    pg_backend_pid(),
    pg_conf_load_time();