module github.com/rkfcccccc/english_words/services/user

go 1.18

require (
	github.com/cespare/xxhash/v2 v2.1.2 // indirect
	github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect
	github.com/go-redis/redis/v8 v8.11.5 // indirect
	github.com/go-redsync/redsync/v4 v4.5.0 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/hashicorp/errwrap v1.0.0 // indirect
	github.com/hashicorp/go-multierror v1.1.0 // indirect
	github.com/jackc/chunkreader/v2 v2.0.1 // indirect
	github.com/jackc/pgconn v1.12.1 // indirect
	github.com/jackc/pgio v1.0.0 // indirect
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgproto3/v2 v2.3.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20200714003250-2b9c44734f2b // indirect
	github.com/jackc/pgtype v1.11.0 // indirect
	github.com/jackc/puddle v1.2.1 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	golang.org/x/net v0.0.0-20220425223048-2871e0cb64e4 // indirect
	golang.org/x/sys v0.0.0-20220503163025-988cb79eb6c6 // indirect
	golang.org/x/text v0.3.7 // indirect
	google.golang.org/genproto v0.0.0-20220505152158-f39f71e6c8f3 // indirect
	google.golang.org/protobuf v1.28.0 // indirect
)

require github.com/rkfcccccc/english_words/proto v0.0.0

replace github.com/rkfcccccc/english_words/proto => ../../proto

require github.com/rkfcccccc/english_words/shared_pkg/dsync v0.0.0

replace github.com/rkfcccccc/english_words/shared_pkg/dsync => ../../shared_pkg/dsync

require github.com/rkfcccccc/english_words/shared_pkg/postgres v0.0.0

replace github.com/rkfcccccc/english_words/shared_pkg/postgres => ../../shared_pkg/postgres

require (
	github.com/georgysavva/scany v0.3.0
	github.com/jackc/pgx/v4 v4.16.1
	github.com/rkfcccccc/english_words/shared_pkg/redis v0.0.0
	golang.org/x/crypto v0.0.0-20220518034528-6f7dac969898
	google.golang.org/grpc v1.46.0
)

replace github.com/rkfcccccc/english_words/shared_pkg/redis => ../../shared_pkg/redis
