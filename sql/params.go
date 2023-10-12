package sql

// ConnParamsPostgres doc ref: https://www.postgresql.org/docs/current/config-setting.html#CONFIG-SETTING-NAMES-VALUES
var ConnParamsPostgres = []string{
	"application_name",
	"fallback_application_name",
	"client_encoding",
	"sslmode",
	"sslkey",
	"sslcert",
	"sslrootcert",
	"sslinline",
	"service",
}

// ConnParamsMySQL doc ref: https://dev.mysql.com/doc/refman/8.0/en/connecting-using-uri-or-key-value-pairs.html
var ConnParamsMySQL = []string{
	"ssl-mode",
	"collation",
}

// ConnParamsSQLite doc ref: https://github.com/mattn/go-sqlite3
var ConnParamsSQLite = []string{
	"_auth_user",                // string
	"_auth_pass",                // string
	"_auth_crypt",               // SHA1 | SSHA1 | SHA256 | SSHA256 | SHA384 | SSHA384 | SHA512 | SSHA512
	"_auth_salt",                // string
	"_auto_vacuum",              // 0 | none, 1 | full, 2 | incremental
	"_vacuum",                   // same with _auto_vacuum
	"_busy_timeout",             // int
	"_timeout",                  // same with _busy_timeout
	"_case_sensitive_like",      // boolean
	"_cslike",                   // boolean
	"_defer_foreign_keys",       // boolean
	"_defer_fk",                 // boolean
	"_foreign_keys",             // boolean
	"_fk",                       // boolean
	"_ignore_check_constraints", // boolean
	"immutable",                 // boolean
	"_journal_mode",             // DELETE | TRUNCATE | PERSIST | MEMORY | WAL | OFF
	"_journal",                  // same with _journal
	"_locking_mode",             // NORMAL | EXCLUSIVE
	"_locking",                  // same with _locking_mode
	"mode",                      // ro | rw | rc | rwc | memory
	"_mutex",                    // no | full
	"_query_only",               // boolean
	"_recursive_triggers",       // boolean
	"_rt",                       // same with _recursive_triggers
	"_secure_delete",            // boolean | FAST
	"cache",                     // options: shared|private
	"_synchronous",              // 0 | OFF, 1 | NORMAL, 2 | FULL, 3 | EXTRA
	"_sync",                     // same with _synchronous
	"_loc",                      // options: auto
	"_txlock",                   // options: immediate | deferred | exclusive
	"_writable_schema",          // boolean
	"_cache_size",               // int
}
