module github.com/rkfcccccc/english_words/services/movie

go 1.18

require (
	github.com/georgysavva/scany v1.0.0
	github.com/jackc/pgx/v4 v4.16.1
)

require (
	github.com/jackc/chunkreader/v2 v2.0.1 // indirect
	github.com/jackc/pgconn v1.12.1 // indirect
	github.com/jackc/pgio v1.0.0 // indirect
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgproto3/v2 v2.3.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20200714003250-2b9c44734f2b // indirect
	github.com/jackc/pgtype v1.11.0 // indirect
	github.com/jackc/puddle v1.2.1 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	golang.org/x/crypto v0.0.0-20220622213112-05595931fe9d // indirect
	golang.org/x/text v0.3.7 // indirect
)

require (
	github.com/golang/protobuf v1.5.2 // indirect
	golang.org/x/net v0.0.0-20220706163947-c90051bbdb60 // indirect
	golang.org/x/sys v0.0.0-20220520151302-bc2c85ada10a // indirect
	google.golang.org/genproto v0.0.0-20200526211855-cb27e3aa2013 // indirect
	google.golang.org/grpc v1.48.0
	google.golang.org/protobuf v1.27.1 // indirect
)

replace github.com/rkfcccccc/english_words/proto => ../../proto

require github.com/rkfcccccc/english_words/proto v0.0.0

replace github.com/rkfcccccc/english_words/shared_pkg/postgres => ../../shared_pkg/postgres

require github.com/rkfcccccc/english_words/shared_pkg/postgres v0.0.0

replace github.com/rkfcccccc/english_words/shared_pkg/services => ../../shared_pkg/services

require github.com/rkfcccccc/english_words/shared_pkg/services v0.0.0

replace github.com/rkfcccccc/english_words/shared_pkg/redis => ../../shared_pkg/redis

require github.com/rkfcccccc/english_words/shared_pkg/redis v0.0.0

replace github.com/rkfcccccc/english_words/shared_pkg/dsync => ../../shared_pkg/dsync

require github.com/rkfcccccc/english_words/shared_pkg/dsync v0.0.0

require (
	github.com/cespare/xxhash/v2 v2.1.2 // indirect
	github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect
	github.com/go-redis/redis/v8 v8.11.5 // indirect
	github.com/go-redsync/redsync/v4 v4.5.0 // indirect
	github.com/hashicorp/errwrap v1.0.0 // indirect
	github.com/hashicorp/go-multierror v1.1.0 // indirect
	github.com/klauspost/compress v1.15.7 // indirect
	github.com/pierrec/lz4/v4 v4.1.15 // indirect
	github.com/segmentio/kafka-go v0.4.34 // indirect
)
