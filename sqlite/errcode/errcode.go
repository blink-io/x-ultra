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
	ErrorCodeString = map[int]string{}
)
