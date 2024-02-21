module github.com/blink-io/x

go 1.22.0

replace github.com/gobwas/glob => github.com/blink-io/glob v0.0.0-20231227024915-2e8bc4bf1fee

require (
	github.com/Masterminds/squirrel v1.5.4
	github.com/ProtonMail/gopenpgp/v2 v2.7.5
	github.com/alexedwards/argon2id v1.0.0
	github.com/andybalholm/brotli v1.1.0
	github.com/cenkalti/backoff/v4 v4.2.1
	github.com/cespare/xxhash/v2 v2.2.0
	github.com/creasty/defaults v1.7.0
	github.com/disgoorg/snowflake/v2 v2.0.1
	github.com/frankban/quicktest v1.14.6 //test
	github.com/fxamacker/cbor/v2 v2.6.0
	github.com/getsentry/sentry-go v0.27.0
	github.com/go-chi/chi/v5 v5.0.12
	github.com/go-kratos/kratos/v2 v2.7.2
	github.com/go-playground/validator/v10 v10.18.0
	github.com/go-resty/resty/v2 v2.11.0
	github.com/go-sql-driver/mysql v1.7.1
	github.com/go-task/slim-sprig/v3 v3.0.0
	github.com/goccy/go-json v0.10.2
	github.com/goccy/go-reflect v1.2.0
	github.com/gofrs/uuid/v5 v5.0.0
	github.com/gogo/protobuf v1.3.2
	github.com/google/flatbuffers v23.5.26+incompatible
	github.com/google/uuid v1.6.0
	github.com/gopacket/gopacket v1.2.0
	github.com/grpc-ecosystem/go-grpc-middleware/v2 v2.0.1
	github.com/hashicorp/golang-lru/v2 v2.0.7
	github.com/jackc/pgx-zap v0.0.0-20221202020421-94b1cb2f889f
	github.com/jackc/pgx/v5 v5.5.3
	github.com/jaevor/go-nanoid v1.3.0
	github.com/jarcoal/httpmock v1.3.1
	github.com/jellydator/ttlcache/v3 v3.2.0
	github.com/joho/godotenv v1.5.1
	github.com/karlseguin/ccache/v3 v3.0.5
	github.com/karrick/godirwalk v1.17.0
	github.com/klauspost/compress v1.17.6
	github.com/lithammer/shortuuid/v4 v4.0.0
	github.com/lmittmann/tint v1.0.4
	github.com/matthewhartstonge/argon2 v1.0.0
	github.com/miekg/dns v1.1.58
	github.com/mitchellh/mapstructure v1.5.0
	github.com/natefinch/lumberjack/v3 v3.0.0-alpha
	github.com/nats-io/nats.go v1.33.1
	github.com/nicksnyder/go-i18n/v2 v2.4.0
	github.com/npillmayer/nestext v0.1.3
	github.com/oklog/ulid/v2 v2.1.0
	github.com/onsi/ginkgo/v2 v2.15.0
	github.com/onsi/gomega v1.31.1
	github.com/outcaste-io/ristretto v0.2.3
	github.com/pelletier/go-toml/v2 v2.1.1
	github.com/pierrec/lz4/v4 v4.1.21
	github.com/quic-go/quic-go v0.41.0
	github.com/redis/go-redis/extra/rediscmd/v9 v9.0.5
	github.com/redis/go-redis/v9 v9.5.1
	github.com/redis/rueidis v1.0.28
	github.com/rs/xid v1.5.0
	github.com/sanity-io/litter v1.5.5
	github.com/segmentio/ksuid v1.0.4
	github.com/stretchr/testify v1.8.4
	github.com/teris-io/shortid v0.0.0-20220617161101-71ec9f2aa569
	github.com/tink-crypto/tink-go/v2 v2.1.0
	github.com/twmb/murmur3 v1.1.8
	github.com/ulikunitz/xz v0.5.11
	github.com/unrolled/render v1.6.1
	github.com/uptrace/bun v1.1.17
	github.com/uptrace/bun/dialect/mysqldialect v1.1.17
	github.com/uptrace/bun/dialect/pgdialect v1.1.17
	github.com/uptrace/bun/dialect/sqlitedialect v1.1.17
	github.com/uptrace/opentelemetry-go-extra/otelsql v0.2.3
	github.com/vmihailenco/go-tinylfu v0.2.2
	github.com/vmihailenco/msgpack/v5 v5.4.1
	github.com/zeebo/xxh3 v1.0.2
	gitlab.com/greyxor/slogor v1.2.3
	go.etcd.io/bbolt v1.3.8
	go.etcd.io/etcd/client/v3 v3.5.12
	go.temporal.io/api v1.27.0
	go.temporal.io/sdk v1.26.0-rc.2
	go.uber.org/multierr v1.11.0
	go.uber.org/zap v1.27.0
	golang.org/x/crypto v0.19.0
	golang.org/x/text v0.14.0
	google.golang.org/genproto/googleapis/api v0.0.0-20240221002015-b0ce06bbee7c
	google.golang.org/grpc v1.61.1
	google.golang.org/protobuf v1.32.0
	gopkg.in/yaml.v3 v3.0.1
	modernc.org/sqlite v1.29.1
)

require (
	connectrpc.com/grpchealth v1.3.0
	connectrpc.com/grpcreflect v1.2.0
	github.com/DATA-DOG/go-sqlmock v1.5.2
	github.com/VictoriaMetrics/fastcache v1.12.2
	github.com/XSAM/otelsql v0.28.0
	github.com/ahmetb/go-linq/v3 v3.2.0
	github.com/alicebob/miniredis/v2 v2.31.1
	github.com/allegro/bigcache/v3 v3.1.0
	github.com/ammario/tlru v0.3.0
	github.com/apache/incubator-fury/go/fury v0.0.0-20240221062422-8e14efad1a54
	github.com/apache/thrift v0.19.0
	github.com/apple/pkl-go v0.5.3
	github.com/avast/retry-go/v4 v4.5.1
	github.com/beevik/guid v1.0.0
	github.com/bits-and-blooms/bloom/v3 v3.6.0
	github.com/bmatcuk/doublestar/v4 v4.6.1
	github.com/brianvoe/gofakeit/v6 v6.28.0
	github.com/bwmarrin/snowflake v0.3.0
	github.com/caarlos0/env/v10 v10.0.0
	github.com/carlmjohnson/requests v0.23.5
	github.com/chzyer/readline v1.5.1
	github.com/dchest/siphash v1.2.3
	github.com/deckarep/golang-set/v2 v2.6.0
	github.com/didip/tollbooth/v7 v7.0.1
	github.com/doug-martin/goqu/v9 v9.19.0
	github.com/eapache/go-resiliency v1.6.0
	github.com/eapache/queue/v2 v2.0.0-20230407133247-75960ed334e4
	github.com/emicklei/proto v1.13.2
	github.com/exaring/otelpgx v0.5.4
	github.com/expr-lang/expr v1.16.1
	github.com/gabriel-vasile/mimetype v1.4.3
	github.com/go-co-op/gocron/v2 v2.2.4
	github.com/go-crypt/crypt v0.2.18
	github.com/go-faster/city v1.0.1
	github.com/go-faster/xor v1.0.0
	github.com/go-kit/log v0.2.1
	github.com/go-kratos/aegis v0.2.0
	github.com/go-rel/mysql v0.12.0
	github.com/go-rel/postgres v0.11.0
	github.com/go-rel/rel v0.41.0
	github.com/go-rel/sqlite3 v0.11.0
	github.com/go-slog/otelslog v0.1.0
	github.com/gobwas/glob v0.2.3
	github.com/gocraft/dbr/v2 v2.7.6
	github.com/golang-cz/devslog v0.0.8
	github.com/gorilla/mux v1.8.1
	github.com/gorilla/websocket v1.5.1
	github.com/h2non/filetype v1.1.3
	github.com/huandu/go-sqlbuilder v1.25.0
	github.com/imroc/req/v3 v3.42.3
	github.com/jackc/chunkreader/v2 v2.0.1
	github.com/jackc/puddle/v2 v2.2.1
	github.com/jhump/protoreflect v1.15.6
	github.com/jonboulle/clockwork v0.4.0
	github.com/k0kubun/pp/v3 v3.2.0
	github.com/kardianos/osext v0.0.0-20190222173326-2bc1f35cddc0
	github.com/leonelquinteros/gotext v1.5.2
	github.com/leporo/sqlf v1.4.0
	github.com/lesismal/nbio v1.5.0
	github.com/lib/pq v1.10.9
	github.com/life4/genesis v1.10.2
	github.com/mailgun/raymond/v2 v2.0.48
	github.com/matryer/is v1.4.1
	github.com/mattn/go-runewidth v0.0.15
	github.com/mattn/go-sqlite3 v1.14.22
	github.com/matttproud/golang_protobuf_extensions/v2 v2.0.0
	github.com/mholt/acmez v1.2.0
	github.com/microcosm-cc/bluemonday v1.0.26
	github.com/mileusna/useragent v1.3.4
	github.com/montanaflynn/stats v0.7.1
	github.com/nleof/goyesql v1.0.2
	github.com/nsqio/go-nsq v1.1.0
	github.com/orcaman/concurrent-map/v2 v2.0.1
	github.com/ory/dockertest v3.3.5+incompatible
	github.com/panjf2000/ants/v2 v2.9.0
	github.com/pashagolub/pgxmock/v3 v3.3.0
	github.com/philhofer/fwd v1.1.2
	github.com/proullon/ramsql v0.1.3
	github.com/qustavo/dotsql v1.2.0
	github.com/remychantenay/slog-otel v1.2.3
	github.com/reugn/go-streams v0.10.0
	github.com/rickar/cal/v2 v2.1.15
	github.com/russross/blackfriday v1.6.0
	github.com/samber/do/v2 v2.0.0-beta.5
	github.com/samber/slog-common v0.15.0
	github.com/samber/slog-multi v1.0.2
	github.com/samber/slog-nats v0.1.1
	github.com/samber/slog-sentry/v2 v2.4.0
	github.com/samber/slog-syslog v1.0.0
	github.com/samber/slog-webhook v1.0.0
	github.com/samber/slog-zap/v2 v2.3.0
	github.com/segmentio/encoding v0.4.0
	github.com/shopspring/decimal v1.3.1
	github.com/smartystreets/goconvey v1.8.1
	github.com/sourcegraph/conc v0.3.0
	github.com/stephenafamo/bob v0.25.0
	github.com/stephenafamo/scan v0.5.0
	github.com/thanhpk/randstr v1.0.6
	github.com/uptrace/bun/driver/pgdriver v1.1.17
	github.com/uptrace/opentelemetry-go-extra/otelutil v0.2.3
	github.com/uptrace/opentelemetry-go-extra/otelzap v0.2.3
	github.com/vingarcia/ksql v1.12.0
	github.com/vmihailenco/tagparser/v2 v2.0.0
	github.com/wk8/go-ordered-map/v2 v2.1.8
	github.com/xo/dburl v0.21.1
	github.com/yuin/goldmark v1.7.0
	github.com/zitadel/passwap v0.5.0
	go.opentelemetry.io/otel v1.23.1
	go.opentelemetry.io/otel/metric v1.23.1
	go.opentelemetry.io/otel/trace v1.23.1
	golang.org/x/net v0.21.0
	golang.org/x/sys v0.17.0
)

require (
	connectrpc.com/connect v1.15.0 // indirect
	github.com/Azure/go-ansiterm v0.0.0-20230124172434-306776ec8161 // indirect
	github.com/Microsoft/go-winio v0.6.1 // indirect
	github.com/Nvveen/Gotty v0.0.0-20120604004816-cd527374f1e5 // indirect
	github.com/ProtonMail/go-crypto v1.0.0 // indirect
	github.com/ProtonMail/go-mime v0.0.0-20230322103455-7d82a3887f2f // indirect
	github.com/alicebob/gopher-json v0.0.0-20230218143504-906a9b012302 // indirect
	github.com/armon/go-radix v1.0.0 // indirect
	github.com/aymerick/douceur v0.2.0 // indirect
	github.com/bahlo/generic-list-go v0.2.0 // indirect
	github.com/bits-and-blooms/bitset v1.13.0 // indirect
	github.com/bufbuild/protocompile v0.8.0 // indirect
	github.com/buger/jsonparser v1.1.1 // indirect
	github.com/cenkalti/backoff v2.2.1+incompatible // indirect
	github.com/cloudflare/circl v1.3.7 // indirect
	github.com/containerd/continuity v0.4.3 // indirect
	github.com/coreos/go-semver v0.3.1 // indirect
	github.com/coreos/go-systemd/v22 v22.5.0 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect
	github.com/docker/go-connections v0.5.0 // indirect
	github.com/docker/go-units v0.5.0 // indirect
	github.com/dustin/go-humanize v1.0.1 // indirect
	github.com/facebookgo/clock v0.0.0-20150410010913-600d898af40a // indirect
	github.com/fsnotify/fsnotify v1.7.0 // indirect
	github.com/go-crypt/x v0.2.12 // indirect
	github.com/go-logfmt/logfmt v0.6.0 // indirect
	github.com/go-logr/logr v1.4.1 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/go-pkgz/expirable-cache v1.0.0 // indirect
	github.com/go-playground/form/v4 v4.2.1 // indirect
	github.com/go-playground/locales v0.14.1 // indirect
	github.com/go-playground/universal-translator v0.18.1 // indirect
	github.com/go-rel/sql v0.16.0 // indirect
	github.com/go-task/slim-sprig v0.0.0-20230315185526-52ccab3ef572 // indirect
	github.com/golang-sql/civil v0.0.0-20220223132316-b832511892a9 // indirect
	github.com/golang/mock v1.6.0 // indirect
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/golang/snappy v0.0.4 // indirect
	github.com/google/go-cmp v0.6.0 // indirect
	github.com/google/pprof v0.0.0-20240207164012-fb44976bdcd5 // indirect
	github.com/gopherjs/gopherjs v1.17.2 // indirect
	github.com/gorilla/css v1.0.1 // indirect
	github.com/gotestyourself/gotestyourself v2.2.0+incompatible // indirect
	github.com/grpc-ecosystem/go-grpc-middleware v1.4.0 // indirect
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.19.1 // indirect
	github.com/hashicorp/errwrap v1.1.0 // indirect
	github.com/hashicorp/go-multierror v1.1.1 // indirect
	github.com/huandu/xstrings v1.4.0 // indirect
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20231201235250-de7065d80cb9 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jmoiron/sqlx v1.3.5 // indirect
	github.com/jtolds/gls v4.20.0+incompatible // indirect
	github.com/klauspost/cpuid/v2 v2.2.7 // indirect
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
	github.com/opencontainers/go-digest v1.0.0 // indirect
	github.com/opencontainers/image-spec v1.1.0 // indirect
	github.com/opencontainers/runc v1.1.12 // indirect
	github.com/pborman/uuid v1.2.1 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/qdm12/reprint v0.0.0-20200326205758-722754a53494 // indirect
	github.com/quic-go/qpack v0.4.0 // indirect
	github.com/refraction-networking/utls v1.6.2 // indirect
	github.com/remyoudompheng/bigfft v0.0.0-20230129092748-24d4a6f8daec // indirect
	github.com/rivo/uniseg v0.4.7 // indirect
	github.com/robfig/cron v1.2.0 // indirect
	github.com/robfig/cron/v3 v3.0.1 // indirect
	github.com/rogpeppe/go-internal v1.12.0 // indirect
	github.com/samber/go-type-to-string v1.2.0 // indirect
	github.com/samber/lo v1.39.0 // indirect
	github.com/segmentio/asm v1.2.0 // indirect
	github.com/serenize/snaker v0.0.0-20201027110005-a7ad2135616e // indirect
	github.com/sirupsen/logrus v1.9.3 // indirect
	github.com/smarty/assertions v1.15.1 // indirect
	github.com/stretchr/objx v0.5.1 // indirect
	github.com/tmthrgd/go-hex v0.0.0-20190904060850-447a3041c3bc // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/x448/float16 v0.8.4 // indirect
	github.com/yuin/gopher-lua v1.1.1 // indirect
	go.etcd.io/etcd/api/v3 v3.5.12 // indirect
	go.etcd.io/etcd/client/pkg/v3 v3.5.12 // indirect
	go.opentelemetry.io/otel/sdk v1.23.1 // indirect
	go.uber.org/atomic v1.11.0 // indirect
	go.uber.org/mock v0.4.0 // indirect
	golang.org/x/exp v0.0.0-20240213143201-ec583247a57a // indirect
	golang.org/x/mod v0.15.0 // indirect
	golang.org/x/sync v0.6.0 // indirect
	golang.org/x/term v0.17.0 // indirect
	golang.org/x/time v0.5.0 // indirect
	golang.org/x/tools v0.18.0 // indirect
	google.golang.org/genproto v0.0.0-20240221002015-b0ce06bbee7c // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240221002015-b0ce06bbee7c // indirect
	gotest.tools v2.2.0+incompatible // indirect
	mellium.im/sasl v0.3.1 // indirect
	modernc.org/gc/v3 v3.0.0-20240107210532-573471604cb6 // indirect
	modernc.org/libc v1.41.0 // indirect
	modernc.org/mathutil v1.6.0 // indirect
	modernc.org/memory v1.7.2 // indirect
	modernc.org/strutil v1.2.0 // indirect
	modernc.org/token v1.1.0 // indirect
)
