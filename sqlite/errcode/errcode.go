package errcode

const (
	SQLITE_ABORT                   = 4
	SQLITE_AUTH                    = 23
	SQLITE_BUSY                    = 5
	SQLITE_IOERR                   = 10
	SQLITE_CANTOPEN                = 14
	SQLITE_CONSTRAINT              = 19
	SQLITE_CORRUPT                 = 11
	SQLITE_DONE                    = 101
	SQLITE_EMPTY                   = 16
	SQLITE_ERROR                   = 1
	SQLITE_FORMAT                  = 24
	SQLITE_FULL                    = 13
	SQLITE_INTERNAL                = 2
	SQLITE_INTERRUPT               = 9
	SQLITE_IOERR_ACCESS            = 3338
	SQLITE_IOERR_AUTH              = 7178
	SQLITE_IOERR_BEGIN_ATOMIC      = 7434
	SQLITE_IOERR_BLOCKED           = 2826
	SQLITE_IOERR_CHECKRESERVEDLOCK = 3594
	SQLITE_IOERR_CLOSE             = 4106
	SQLITE_IOERR_COMMIT_ATOMIC     = 7690
	SQLITE_IOERR_CONVPATH          = 6666
	SQLITE_IOERR_CORRUPTFS         = 8458
	SQLITE_IOERR_DATA              = 8202
	SQLITE_IOERR_DELETE            = 2570
	SQLITE_IOERR_DELETE_NOENT      = 5898
	SQLITE_IOERR_DIR_CLOSE         = 4362
	SQLITE_IOERR_DIR_FSYNC         = 1290
	SQLITE_IOERR_FSTAT             = 1802
	SQLITE_IOERR_FSYNC             = 1034
	SQLITE_IOERR_GETTEMPPATH       = 6410
	SQLITE_IOERR_LOCK              = 3850
	SQLITE_IOERR_MMAP              = 6154
	SQLITE_IOERR_NOMEM             = 3082
	SQLITE_IOERR_RDLOCK            = 2314
	SQLITE_IOERR_READ              = 266
	SQLITE_IOERR_ROLLBACK_ATOMIC   = 7946
	SQLITE_IOERR_SEEK              = 5642
	SQLITE_IOERR_SHMLOCK           = 5130
	SQLITE_IOERR_SHMMAP            = 5386
	SQLITE_IOERR_SHMOPEN           = 4618
	SQLITE_IOERR_SHMSIZE           = 4874
	SQLITE_IOERR_SHORT_READ        = 522
	SQLITE_IOERR_TRUNCATE          = 1546
	SQLITE_IOERR_UNLOCK            = 2058
	SQLITE_IOERR_VNODE             = 6922
	SQLITE_IOERR_WRITE             = 778
	SQLITE_LOCKED                  = 6
	SQLITE_MISMATCH                = 20
	SQLITE_MISUSE                  = 21
	SQLITE_NOLFS                   = 22
	SQLITE_NOMEM                   = 7
	SQLITE_NOTADB                  = 26
	SQLITE_NOTFOUND                = 12
	SQLITE_NOTICE                  = 27
	SQLITE_PERM                    = 3
	SQLITE_PROTOCOL                = 15
	SQLITE_RANGE                   = 25
	SQLITE_READONLY                = 8
	SQLITE_ROW                     = 100
	SQLITE_SCHEMA                  = 17
	SQLITE_TOOBIG                  = 18
	SQLITE_WARNING                 = 28
)

var (
	// ErrorCodeString maps Error.Code() to its string representation.
	// http://www.sqlite.org/c3ref/c_abort_rollback.html
	ErrorCodeString = map[int]string{
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
)
