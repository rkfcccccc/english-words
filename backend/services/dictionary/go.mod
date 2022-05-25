module github.com/rkfcccccc/english_words/services/dictionary

go 1.18

require github.com/stretchr/testify v1.7.1

require (
	github.com/andybalholm/cascadia v1.3.1 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/go-stack/stack v1.8.1 // indirect
	github.com/golang/snappy v0.0.4 // indirect
	github.com/google/go-cmp v0.5.6 // indirect
	github.com/klauspost/compress v1.15.4 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/xdg-go/pbkdf2 v1.0.0 // indirect
	github.com/xdg-go/scram v1.1.1 // indirect
	github.com/xdg-go/stringprep v1.0.3 // indirect
	github.com/youmark/pkcs8 v0.0.0-20201027041543-1326539a0a0a // indirect
	golang.org/x/crypto v0.0.0-20220518034528-6f7dac969898 // indirect
	golang.org/x/net v0.0.0-20220520000938-2e3eb7b945c2 // indirect
	golang.org/x/sync v0.0.0-20220513210516-0976fa681c29 // indirect
	golang.org/x/text v0.3.7 // indirect
	golang.org/x/xerrors v0.0.0-20200804184101-5ec99f83aff1 // indirect
	gopkg.in/check.v1 v1.0.0-20190902080502-41f04d3bba15 // indirect
	gopkg.in/yaml.v3 v3.0.0-20200313102051-9f266ea9e77c // indirect
)

replace github.com/rkfcccccc/english_words/proto => ../../proto

require github.com/rkfcccccc/english_words/shared_pkg/dsync v0.0.0

replace github.com/rkfcccccc/english_words/shared_pkg/dsync => ../../shared_pkg/dsync

require github.com/rkfcccccc/english_words/shared_pkg/cache v0.0.0

replace github.com/rkfcccccc/english_words/shared_pkg/cache => ../../shared_pkg/cache

require github.com/rkfcccccc/english_words/shared_pkg/mongodb v0.0.0

replace github.com/rkfcccccc/english_words/shared_pkg/mongodb => ../../shared_pkg/mongodb

require (
	github.com/PuerkitoBio/goquery v1.8.0
	github.com/joho/godotenv v1.4.0
	go.mongodb.org/mongo-driver v1.9.1
)

replace github.com/rkfcccccc/english_words/shared_pkg/redis => ../../shared_pkg/redis
