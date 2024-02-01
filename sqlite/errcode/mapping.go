package errcode

// ErrorNum is a number error code.
type ErrorNum int

// Name returns a more human friendly rendering of the error code, namely the
// "condition name".
func (ec ErrorNum) Name() string {
	return errorNumNames[ec]
}

var errorNumNames = map[ErrorNum]string{
	SQLITE_ABORT:             "Callback routine requested an abort (SQLITE_ABORT)",
	SQLITE_AUTH:              "Authorization denied (SQLITE_AUTH)",
	SQLITE_BUSY:              "The database file is locked (SQLITE_BUSY)",
	SQLITE_CANTOPEN:          "Unable to open the database file (SQLITE_CANTOPEN)",
	SQLITE_CONSTRAINT:        "Abort due to constraint violation (SQLITE_CONSTRAINT)",
	SQLITE_CORRUPT:           "The database disk image is malformed (SQLITE_CORRUPT)",
	SQLITE_DONE:              "sqlite3_step() has finished executing (SQLITE_DONE)",
	SQLITE_EMPTY:             "Internal use only (SQLITE_EMPTY)",
	SQLITE_ERROR:             "Generic error (SQLITE_ERROR)",
	SQLITE_FORMAT:            "Not used (SQLITE_FORMAT)",
	SQLITE_FULL:              "Insertion failed because database is full (SQLITE_FULL)",
	SQLITE_INTERNAL:          "Internal logic error in SQLite (SQLITE_INTERNAL)",
	SQLITE_INTERRUPT:         "Operation terminated by sqlite3_interrupt()(SQLITE_INTERRUPT)",
	SQLITE_IOERR | (1 << 8):  "(SQLITE_IOERR_READ)",
	SQLITE_IOERR | (10 << 8): "(SQLITE_IOERR_DELETE)",
	SQLITE_IOERR | (11 << 8): "(SQLITE_IOERR_BLOCKED)",
	SQLITE_IOERR | (12 << 8): "(SQLITE_IOERR_NOMEM)",
	SQLITE_IOERR | (13 << 8): "(SQLITE_IOERR_ACCESS)",
	SQLITE_IOERR | (14 << 8): "(SQLITE_IOERR_CHECKRESERVEDLOCK)",
	SQLITE_IOERR | (15 << 8): "(SQLITE_IOERR_LOCK)",
	SQLITE_IOERR | (16 << 8): "(SQLITE_IOERR_CLOSE)",
	SQLITE_IOERR | (17 << 8): "(SQLITE_IOERR_DIR_CLOSE)",
	SQLITE_IOERR | (2 << 8):  "(SQLITE_IOERR_SHORT_READ)",
	SQLITE_IOERR | (3 << 8):  "(SQLITE_IOERR_WRITE)",
	SQLITE_IOERR | (4 << 8):  "(SQLITE_IOERR_FSYNC)",
	SQLITE_IOERR | (5 << 8):  "(SQLITE_IOERR_DIR_FSYNC)",
	SQLITE_IOERR | (6 << 8):  "(SQLITE_IOERR_TRUNCATE)",
	SQLITE_IOERR | (7 << 8):  "(SQLITE_IOERR_FSTAT)",
	SQLITE_IOERR | (8 << 8):  "(SQLITE_IOERR_UNLOCK)",
	SQLITE_IOERR | (9 << 8):  "(SQLITE_IOERR_RDLOCK)",
	SQLITE_IOERR:             "Some kind of disk I/O error occurred (SQLITE_IOERR)",
	SQLITE_LOCKED | (1 << 8): "(SQLITE_LOCKED_SHAREDCACHE)",
	SQLITE_LOCKED:            "A table in the database is locked (SQLITE_LOCKED)",
	SQLITE_MISMATCH:          "Data type mismatch (SQLITE_MISMATCH)",
	SQLITE_MISUSE:            "Library used incorrectly (SQLITE_MISUSE)",
	SQLITE_NOLFS:             "Uses OS features not supported on host (SQLITE_NOLFS)",
	SQLITE_NOMEM:             "A malloc() failed (SQLITE_NOMEM)",
	SQLITE_NOTADB:            "File opened that is not a database file (SQLITE_NOTADB)",
	SQLITE_NOTFOUND:          "Unknown opcode in sqlite3_file_control() (SQLITE_NOTFOUND)",
	SQLITE_NOTICE:            "Notifications from sqlite3_log() (SQLITE_NOTICE)",
	SQLITE_PERM:              "Access permission denied (SQLITE_PERM)",
	SQLITE_PROTOCOL:          "Database lock protocol error (SQLITE_PROTOCOL)",
	SQLITE_RANGE:             "2nd parameter to sqlite3_bind out of range (SQLITE_RANGE)",
	SQLITE_READONLY:          "Attempt to write a readonly database (SQLITE_READONLY)",
	SQLITE_ROW:               "sqlite3_step() has another row ready (SQLITE_ROW)",
	SQLITE_SCHEMA:            "The database schema changed (SQLITE_SCHEMA)",
	SQLITE_TOOBIG:            "String or BLOB exceeds size limit (SQLITE_TOOBIG)",
	SQLITE_WARNING:           "Warnings from sqlite3_log() (SQLITE_WARNING)",
}
