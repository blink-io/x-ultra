module github.com/blink-io/x

go 1.22.4

replace github.com/gobwas/glob => github.com/blink-io/glob v0.0.0-20231227024915-2e8bc4bf1fee

require (
	github.com/Masterminds/squirrel v1.5.4
	github.com/ProtonMail/gopenpgp/v2 v2.7.5
	github.com/alexedwards/argon2id v1.0.0
	github.com/andybalholm/brotli v1.1.0
	github.com/cenkalti/backoff/v4 v4.3.0
	github.com/cespare/xxhash/v2 v2.3.0
	github.com/creasty/defaults v1.7.0
	github.com/disgoorg/snowflake/v2 v2.0.3
	github.com/frankban/quicktest v1.14.6 //test
	github.com/fxamacker/cbor/v2 v2.7.0
	github.com/getsentry/sentry-go v0.28.1
	github.com/go-chi/chi/v5 v5.1.0
	github.com/go-kratos/kratos/v2 v2.8.0
	github.com/go-playground/validator/v10 v10.22.0
	github.com/go-resty/resty/v2 v2.13.1
	github.com/go-sql-driver/mysql v1.8.1
	github.com/go-task/slim-sprig/v3 v3.0.0
	github.com/goccy/go-json v0.10.3
	github.com/goccy/go-reflect v1.2.0
	github.com/gofrs/uuid/v5 v5.2.0
	github.com/gogo/protobuf v1.3.2
	github.com/google/uuid v1.6.0
	github.com/gopacket/gopacket v1.2.0
	github.com/grpc-ecosystem/go-grpc-middleware/v2 v2.1.0
	github.com/hashicorp/golang-lru/v2 v2.0.7
	github.com/jackc/pgx-zap v0.0.0-20221202020421-94b1cb2f889f
	github.com/jackc/pgx/v5 v5.6.0
	github.com/jaevor/go-nanoid v1.4.0
	github.com/jarcoal/httpmock v1.3.1
	github.com/jellydator/ttlcache/v3 v3.2.0
	github.com/joho/godotenv v1.5.1
	github.com/karlseguin/ccache/v3 v3.0.5
	github.com/karrick/godirwalk v1.17.0
	github.com/klauspost/compress v1.17.9
	github.com/lithammer/shortuuid/v4 v4.0.0
	github.com/lmittmann/tint v1.0.5
	github.com/matthewhartstonge/argon2 v1.0.0
	github.com/miekg/dns v1.1.61
	github.com/mitchellh/mapstructure v1.5.0
	github.com/natefinch/lumberjack/v3 v3.0.0-alpha
	github.com/nats-io/nats.go v1.36.0
	github.com/nicksnyder/go-i18n/v2 v2.4.0
	github.com/npillmayer/nestext v0.1.3
	github.com/oklog/ulid/v2 v2.1.0
	github.com/onsi/ginkgo/v2 v2.19.1
	github.com/onsi/gomega v1.34.0
	github.com/outcaste-io/ristretto v0.2.3
	github.com/pelletier/go-toml/v2 v2.2.2
	github.com/pierrec/lz4/v4 v4.1.21
	github.com/quic-go/quic-go v0.45.1
	github.com/redis/go-redis/extra/rediscmd/v9 v9.5.3
	github.com/redis/go-redis/v9 v9.6.1
	github.com/redis/rueidis v1.0.43
	github.com/rs/xid v1.5.0
	github.com/sanity-io/litter v1.5.5
	github.com/segmentio/ksuid v1.0.4
	github.com/stretchr/testify v1.9.0
	github.com/teris-io/shortid v0.0.0-20220617161101-71ec9f2aa569
	github.com/tink-crypto/tink-go/v2 v2.2.0
	github.com/twmb/murmur3 v1.1.8
	github.com/ulikunitz/xz v0.5.12
	github.com/unrolled/render v1.6.1
	github.com/uptrace/bun v1.2.1
	github.com/uptrace/bun/dialect/mysqldialect v1.2.1
	github.com/uptrace/bun/dialect/pgdialect v1.2.1
	github.com/uptrace/bun/dialect/sqlitedialect v1.2.1
	github.com/uptrace/opentelemetry-go-extra/otelsql v0.3.1
	github.com/vmihailenco/go-tinylfu v0.2.2
	github.com/vmihailenco/msgpack/v5 v5.4.1
	github.com/zeebo/xxh3 v1.0.2
	gitlab.com/greyxor/slogor v1.2.10
	go.etcd.io/bbolt v1.3.10
	go.etcd.io/etcd/client/v3 v3.5.15
	go.temporal.io/api v1.36.0
	go.temporal.io/sdk v1.28.1
	go.uber.org/multierr v1.11.0
	go.uber.org/zap v1.27.0
	golang.org/x/crypto v0.25.0
	golang.org/x/text v0.16.0
	google.golang.org/genproto/googleapis/api v0.0.0-20240725223205-93522f1f2a9f
	google.golang.org/grpc v1.65.0
	google.golang.org/protobuf v1.34.2
	gopkg.in/yaml.v3 v3.0.1
	modernc.org/sqlite v1.31.1
)

require (
	connectrpc.com/grpchealth v1.3.0
	connectrpc.com/grpcreflect v1.2.0
	github.com/DATA-DOG/go-sqlmock v1.5.2
	github.com/IBM/sarama v1.43.2
	github.com/VictoriaMetrics/fastcache v1.12.2
	github.com/XSAM/otelsql v0.32.0
	github.com/aarondl/opt v0.0.0-20240623220848-083f18ab9536
	github.com/ahmetb/go-linq/v3 v3.2.0
	github.com/alicebob/miniredis/v2 v2.33.0
	github.com/allegro/bigcache/v3 v3.1.0
	github.com/ammario/tlru v0.4.0
	github.com/apache/thrift v0.20.0
	github.com/avast/retry-go/v4 v4.6.0
	github.com/beevik/guid v1.0.0
	github.com/bits-and-blooms/bloom/v3 v3.7.0
	github.com/blockloop/scan/v2 v2.5.0
	github.com/bmatcuk/doublestar/v4 v4.6.1
	github.com/brianvoe/gofakeit/v6 v6.28.0
	github.com/bsm/redislock v0.9.4
	github.com/bwmarrin/snowflake v0.3.0
	github.com/caarlos0/env/v10 v10.0.0
	github.com/carlmjohnson/requests v0.24.1
	github.com/cohesivestack/valgo v0.4.1
	github.com/dchest/siphash v1.2.3
	github.com/deckarep/golang-set/v2 v2.6.0
	github.com/didip/tollbooth/v7 v7.0.2
	github.com/doug-martin/goqu/v9 v9.19.0
	github.com/eapache/go-resiliency v1.7.0
	github.com/eapache/queue/v2 v2.0.0-20230407133247-75960ed334e4
	github.com/emicklei/proto v1.13.2
	github.com/exaring/otelpgx v0.6.2
	github.com/expr-lang/expr v1.16.9
	github.com/gabriel-vasile/mimetype v1.4.5
	github.com/go-co-op/gocron/v2 v2.11.0
	github.com/go-crypt/crypt v0.2.25
	github.com/go-faster/city v1.0.1
	github.com/go-faster/xor v1.0.0
	github.com/go-kit/log v0.2.1
	github.com/go-kratos/aegis v0.2.0
	github.com/go-slog/otelslog v0.1.0
	github.com/gobwas/glob v0.2.3
	github.com/gocraft/dbr/v2 v2.7.6
	github.com/golang-cz/devslog v0.0.8
	github.com/gorilla/mux v1.8.1
	github.com/gorilla/websocket v1.5.3
	github.com/h2non/filetype v1.1.3
	github.com/huandu/go-sqlbuilder v1.28.0
	github.com/imroc/req/v3 v3.43.7
	github.com/jackc/chunkreader/v2 v2.0.1
	github.com/jackc/puddle/v2 v2.2.1
	github.com/jhump/protoreflect v1.16.0
	github.com/jonboulle/clockwork v0.4.0
	github.com/k0kubun/pp/v3 v3.2.0
	github.com/kardianos/osext v0.0.0-20190222173326-2bc1f35cddc0
	github.com/leonelquinteros/gotext v1.6.1
	github.com/lesismal/nbio v1.5.9
	github.com/lib/pq v1.10.9
	github.com/life4/genesis v1.10.3
	github.com/matryer/is v1.4.1
	github.com/mattn/go-runewidth v0.0.16
	github.com/mattn/go-sqlite3 v1.14.22
	github.com/matttproud/golang_protobuf_extensions/v2 v2.0.0
	github.com/mholt/acmez v1.2.0
	github.com/microcosm-cc/bluemonday v1.0.27
	github.com/mileusna/useragent v1.3.4
	github.com/montanaflynn/stats v0.7.1
	github.com/orcaman/concurrent-map/v2 v2.0.1
	github.com/panjf2000/ants/v2 v2.10.0
	github.com/pashagolub/pgxmock/v3 v3.4.0
	github.com/philhofer/fwd v1.1.2
	github.com/proullon/ramsql v0.1.4
	github.com/qustavo/dotsql v1.2.0
	github.com/redis/rueidis/rueidishook v1.0.43
	github.com/remychantenay/slog-otel v1.3.2
	github.com/russross/blackfriday v1.6.0
	github.com/samber/do/v2 v2.0.0-beta.7
	github.com/samber/slog-common v0.17.1
	github.com/samber/slog-multi v1.2.0
	github.com/samber/slog-nats v0.4.0
	github.com/samber/slog-sentry/v2 v2.7.0
	github.com/samber/slog-syslog v1.0.0
	github.com/samber/slog-webhook v1.0.0
	github.com/samber/slog-zap/v2 v2.6.0
	github.com/segmentio/encoding v0.4.0
	github.com/segmentio/kafka-go v0.4.47
	github.com/shopspring/decimal v1.4.0
	github.com/smartystreets/goconvey v1.8.1
	github.com/sourcegraph/conc v0.3.0
	github.com/stephenafamo/bob v0.28.1
	github.com/stephenafamo/scan v0.5.0
	github.com/thanhpk/randstr v1.0.6
	github.com/uptrace/opentelemetry-go-extra/otelutil v0.3.1
	github.com/uptrace/opentelemetry-go-extra/otelzap v0.3.1
	github.com/vmihailenco/tagparser/v2 v2.0.0
	github.com/wk8/go-ordered-map/v2 v2.1.8
	github.com/xo/dburl v0.23.2
	github.com/yuin/goldmark v1.7.4
	github.com/zitadel/passwap v0.6.0
	go.opentelemetry.io/otel v1.28.0
	go.opentelemetry.io/otel/metric v1.28.0
	go.opentelemetry.io/otel/trace v1.28.0
	golang.org/x/net v0.27.0
	golang.org/x/sys v0.22.0
)

require (
	connectrpc.com/connect v1.16.2 // indirect
	filippo.io/edwards25519 v1.1.0 // indirect
	github.com/ProtonMail/go-crypto v1.0.0 // indirect
	github.com/ProtonMail/go-mime v0.0.0-20230322103455-7d82a3887f2f // indirect
	github.com/aarondl/json v0.0.0-20221020222930-8b0db17ef1bf // indirect
	github.com/alicebob/gopher-json v0.0.0-20230218143504-906a9b012302 // indirect
	github.com/armon/go-radix v1.0.0 // indirect
	github.com/aymerick/douceur v0.2.0 // indirect
	github.com/bahlo/generic-list-go v0.2.0 // indirect
	github.com/bits-and-blooms/bitset v1.13.0 // indirect
	github.com/bufbuild/protocompile v0.14.0 // indirect
	github.com/buger/jsonparser v1.1.1 // indirect
	github.com/cloudflare/circl v1.3.9 // indirect
	github.com/coreos/go-semver v0.3.1 // indirect
	github.com/coreos/go-systemd/v22 v22.5.0 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect
	github.com/dustin/go-humanize v1.0.1 // indirect
	github.com/eapache/go-xerial-snappy v0.0.0-20230731223053-c322873962e3 // indirect
	github.com/eapache/queue v1.1.0 // indirect
	github.com/facebookgo/clock v0.0.0-20150410010913-600d898af40a // indirect
	github.com/fsnotify/fsnotify v1.7.0 // indirect
	github.com/go-crypt/x v0.2.18 // indirect
	github.com/go-logfmt/logfmt v0.6.0 // indirect
	github.com/go-logr/logr v1.4.2 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/go-pkgz/expirable-cache/v3 v3.0.0 // indirect
	github.com/go-playground/form/v4 v4.2.1 // indirect
	github.com/go-playground/locales v0.14.1 // indirect
	github.com/go-playground/universal-translator v0.18.1 // indirect
	github.com/golang-sql/civil v0.0.0-20220223132316-b832511892a9 // indirect
	github.com/golang/mock v1.6.0 // indirect
	github.com/golang/protobuf v1.5.4 // indirect
	github.com/golang/snappy v0.0.4 // indirect
	github.com/google/go-cmp v0.6.0 // indirect
	github.com/google/pprof v0.0.0-20240727154555-813a5fbdbec8 // indirect
	github.com/gopherjs/gopherjs v1.17.2 // indirect
	github.com/gorilla/css v1.0.1 // indirect
	github.com/grpc-ecosystem/go-grpc-middleware v1.4.0 // indirect
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.21.0 // indirect
	github.com/hashicorp/errwrap v1.1.0 // indirect
	github.com/hashicorp/go-multierror v1.1.1 // indirect
	github.com/hashicorp/go-uuid v1.0.3 // indirect
	github.com/huandu/xstrings v1.5.0 // indirect
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20240606120523-5a60cdf6a761 // indirect
	github.com/jcmturner/aescts/v2 v2.0.0 // indirect
	github.com/jcmturner/dnsutils/v2 v2.0.0 // indirect
	github.com/jcmturner/gofork v1.7.6 // indirect
	github.com/jcmturner/gokrb5/v8 v8.4.4 // indirect
	github.com/jcmturner/rpc/v2 v2.0.3 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jmoiron/sqlx v1.3.5 // indirect
	github.com/jtolds/gls v4.20.0+incompatible // indirect
	github.com/klauspost/cpuid/v2 v2.2.8 // indirect
	github.com/kr/pretty v0.3.1 // indirect
	github.com/kr/text v0.2.0 // indirect
	github.com/lann/builder v0.0.0-20180802200727-47ae307949d0 // indirect
	github.com/lann/ps v0.0.0-20150810152359-62de8c46ede0 // indirect
	github.com/leodido/go-urn v1.4.0 // indirect
	github.com/mailru/easyjson v0.7.7 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/nats-io/nkeys v0.4.7 // indirect
	github.com/nats-io/nuid v1.0.1 // indirect
	github.com/ncruces/go-strftime v0.1.9 // indirect
	github.com/nexus-rpc/sdk-go v0.0.9 // indirect
	github.com/pborman/uuid v1.2.1 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/qdm12/reprint v0.0.0-20200326205758-722754a53494 // indirect
	github.com/quic-go/qpack v0.4.0 // indirect
	github.com/rcrowley/go-metrics v0.0.0-20201227073835-cf1acfcdf475 // indirect
	github.com/refraction-networking/utls v1.6.7 // indirect
	github.com/remyoudompheng/bigfft v0.0.0-20230129092748-24d4a6f8daec // indirect
	github.com/rivo/uniseg v0.4.7 // indirect
	github.com/robfig/cron v1.2.0 // indirect
	github.com/robfig/cron/v3 v3.0.1 // indirect
	github.com/rogpeppe/go-internal v1.12.0 // indirect
	github.com/samber/go-type-to-string v1.7.0 // indirect
	github.com/samber/lo v1.46.0 // indirect
	github.com/segmentio/asm v1.2.0 // indirect
	github.com/smarty/assertions v1.16.0 // indirect
	github.com/stretchr/objx v0.5.2 // indirect
	github.com/tmthrgd/go-hex v0.0.0-20190904060850-447a3041c3bc // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/valyala/fasttemplate v1.2.2 // indirect
	github.com/x448/float16 v0.8.4 // indirect
	github.com/yuin/gopher-lua v1.1.1 // indirect
	go.etcd.io/etcd/api/v3 v3.5.15 // indirect
	go.etcd.io/etcd/client/pkg/v3 v3.5.15 // indirect
	go.opentelemetry.io/otel/log v0.4.0 // indirect
	go.opentelemetry.io/otel/sdk v1.28.0 // indirect
	go.uber.org/atomic v1.11.0 // indirect
	go.uber.org/mock v0.4.0 // indirect
	golang.org/x/exp v0.0.0-20240719175910-8a7402abbf56 // indirect
	golang.org/x/mod v0.19.0 // indirect
	golang.org/x/sync v0.7.0 // indirect
	golang.org/x/term v0.22.0 // indirect
	golang.org/x/time v0.5.0 // indirect
	golang.org/x/tools v0.23.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240725223205-93522f1f2a9f // indirect
	modernc.org/gc/v3 v3.0.0-20240722195230-4a140ff9c08e // indirect
	modernc.org/libc v1.55.5 // indirect
	modernc.org/mathutil v1.6.0 // indirect
	modernc.org/memory v1.8.0 // indirect
	modernc.org/strutil v1.2.0 // indirect
	modernc.org/token v1.1.0 // indirect
)
