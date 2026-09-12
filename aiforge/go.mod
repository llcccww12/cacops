module code.gitea.io/gitea

go 1.24.0

toolchain go1.24.11

require (
	gitea.com/jolheiser/gitea-vet v0.1.0
	gitea.com/lunny/levelqueue v0.3.0
	gitea.com/macaron/binding v0.0.0-20190822013154-a5f53841ed2b
	gitea.com/macaron/cache v0.0.0-20190822004001-a6e7fee4ee76
	gitea.com/macaron/captcha v0.0.0-20190822015246-daa973478bae
	gitea.com/macaron/cors v0.0.0-20190826180238-95aec09ea8b4
	gitea.com/macaron/csrf v0.0.0-20190822024205-3dc5a4474439
	gitea.com/macaron/gzip v0.0.0-20191118041502-506895b47aae
	gitea.com/macaron/i18n v0.0.0-20190822004228-474e714e2223
	gitea.com/macaron/inject v0.0.0-20190805023432-d4c86e31027a
	gitea.com/macaron/macaron v1.4.0
	gitea.com/macaron/session v0.0.0-20191207215012-613cebf0674d
	gitea.com/macaron/toolbox v0.0.0-20190822013122-05ff0fc766b7
	github.com/360EntSecGroup-Skylar/excelize/v2 v2.0.2
	github.com/BurntSushi/toml v0.3.1
	github.com/Jeffail/tunny v0.1.4
	github.com/PuerkitoBio/goquery v1.5.0
	github.com/RichardKnop/machinery v1.6.9
	github.com/alecthomas/chroma v0.10.0
	github.com/alibabacloud-go/darabonba-openapi v0.1.18
	github.com/alibabacloud-go/dysmsapi-20170525/v2 v2.0.9
	github.com/alibabacloud-go/tea v1.1.17
	github.com/alibabacloud-go/tea-utils v1.4.3
	github.com/blevesearch/bleve v1.0.7
	github.com/denisenkom/go-mssqldb v0.0.0-20200428022330-06a60b6afbbc
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/disintegration/imaging v1.6.2
	github.com/dustin/go-humanize v1.0.0
	github.com/editorconfig/editorconfig-core-go/v2 v2.1.1
	github.com/elliotchance/orderedmap v1.4.0
	github.com/emirpasic/gods v1.18.1
	github.com/ethantkoenig/rupture v0.0.0-20180203182544-0a76f03a811a
	github.com/gliderlabs/ssh v0.3.5
	github.com/go-enry/go-enry/v2 v2.3.0
	github.com/go-git/go-billy/v5 v5.5.0
	github.com/go-git/go-git/v5 v5.7.0
	github.com/go-http-utils/headers v0.0.0-20181008091004-fed159eddc2a
	github.com/go-macaron/auth v0.0.0-20161228062157-884c0e6c9b92
	github.com/go-redis/redis v6.15.2+incompatible
	github.com/go-resty/resty/v2 v2.3.0
	github.com/go-sql-driver/mysql v1.4.1
	github.com/go-swagger/go-swagger v0.21.0
	github.com/gobwas/glob v0.2.3
	github.com/gogs/chardet v0.0.0-20191104214054-4b6791f73a28
	github.com/gogs/cron v0.0.0-20171120032916-9f6c956d3e14
	github.com/gomodule/redigo v2.0.0+incompatible
	github.com/google/go-github/v24 v24.0.1
	github.com/google/go-querystring v1.0.0
	github.com/google/uuid v1.6.0
	github.com/gorilla/context v1.1.1
	github.com/gorilla/websocket v1.4.0
	github.com/grafana/otel-profiling-go v0.5.1
	github.com/grafana/pyroscope-go v1.2.7
	github.com/hashicorp/go-retryablehttp v0.6.6
	github.com/huandu/xstrings v1.3.0
	github.com/huaweicloud/huaweicloud-sdk-go-obs v3.25.9+incompatible
	github.com/issue9/identicon v1.0.1
	github.com/jaytaylor/html2text v0.0.0-20160923191438-8fb95d837f7d
	github.com/kballard/go-shellquote v0.0.0-20170619183022-cd60e84ee657
	github.com/keybase/go-crypto v0.0.0-20200123153347-de78d2cb44f4
	github.com/klauspost/compress v1.17.8
	github.com/lafriks/xormstore v1.3.2
	github.com/lib/pq v1.2.0
	github.com/lunny/dingtalk_webhook v0.0.0-20171025031554-e3534c89ef96
	github.com/markbates/goth v1.61.2
	github.com/mattn/go-isatty v0.0.11
	github.com/mattn/go-sqlite3 v1.11.0
	github.com/mcuadros/go-version v0.0.0-20190308113854-92cdf37c5b75
	github.com/mgechev/dots v0.0.0-20190921121421-c36f7dcfbb81
	github.com/mgechev/revive v1.0.2
	github.com/microcosm-cc/bluemonday v1.0.3-0.20191119130333-0a75d7616912
	github.com/minio/minio-go v6.0.14+incompatible
	github.com/minio/minio-go/v6 v6.0.57
	github.com/mitchellh/go-homedir v1.1.0
	github.com/msteinert/pam v0.0.0-20151204160544-02ccfbfaf0cc
	github.com/nfnt/resize v0.0.0-20160724205520-891127d8d1b5
	github.com/niklasfasching/go-org v0.1.9
	github.com/oliamb/cutter v0.2.2
	github.com/olivere/elastic/v7 v7.0.9
	github.com/patrickmn/go-cache v2.1.0+incompatible
	github.com/pkg/errors v0.9.1
	github.com/pquerna/otp v1.2.0
	github.com/prometheus/client_golang v1.1.0
	github.com/quasoft/websspi v1.0.0
	github.com/satori/go.uuid v1.2.0
	github.com/sergi/go-diff v1.1.0
	github.com/shirou/gopsutil/v3 v3.20.10
	github.com/shurcooL/vfsgen v0.0.0-20181202132449-6a9ea43bcacd
	github.com/sirupsen/logrus v1.7.0
	github.com/stretchr/testify v1.11.1
	github.com/tstranex/u2f v1.0.0
	github.com/unknwon/cae v1.0.0
	github.com/unknwon/com v1.0.1
	github.com/unknwon/i18n v0.0.0-20190805065654-5c6446a380b6
	github.com/unknwon/paginater v0.0.0-20151104151617-7748a72e0141
	github.com/urfave/cli v1.22.1
	github.com/xanzy/go-gitlab v0.31.0
	github.com/yohcop/openid-go v1.0.0
	github.com/yuin/goldmark v1.7.13
	github.com/yuin/goldmark-highlighting v0.0.0-20220208100518-594be1970594
	github.com/yuin/goldmark-meta v1.1.0
	go.opentelemetry.io/otel v1.39.0
	go.opentelemetry.io/otel/exporters/otlp/otlplog/otlploghttp v0.15.0
	go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetrichttp v1.39.0
	go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp v1.39.0
	go.opentelemetry.io/otel/exporters/stdout/stdoutmetric v1.39.0
	go.opentelemetry.io/otel/log v0.15.0
	go.opentelemetry.io/otel/metric v1.39.0
	go.opentelemetry.io/otel/sdk v1.39.0
	go.opentelemetry.io/otel/sdk/log v0.15.0
	go.opentelemetry.io/otel/sdk/metric v1.39.0
	go.opentelemetry.io/otel/trace v1.39.0
	golang.org/x/crypto v0.46.0
	golang.org/x/net v0.48.0
	golang.org/x/oauth2 v0.34.0
	golang.org/x/sys v0.39.0
	golang.org/x/text v0.32.0
	golang.org/x/time v0.14.0
	golang.org/x/tools v0.39.0
	gopkg.in/gomail.v2 v2.0.0-20160411212932-81ebce5c23df
	gopkg.in/ini.v1 v1.56.0
	gopkg.in/ldap.v3 v3.0.2
	gopkg.in/testfixtures.v2 v2.5.0
	gopkg.in/yaml.v2 v2.4.0
	mvdan.cc/xurls/v2 v2.1.0
	strk.kbt.io/projects/go/libravatar v0.0.0-20191008002943-06d1c002b251
	xorm.io/builder v0.3.7
	xorm.io/xorm v1.0.1
)

require (
	cloud.google.com/go v0.121.6 // indirect
	cloud.google.com/go/auth v0.18.0 // indirect
	cloud.google.com/go/auth/oauth2adapt v0.2.8 // indirect
	cloud.google.com/go/compute/metadata v0.9.0 // indirect
	cloud.google.com/go/iam v1.5.3 // indirect
	cloud.google.com/go/pubsub v1.50.1 // indirect
	cloud.google.com/go/pubsub/v2 v2.3.0 // indirect
	github.com/Microsoft/go-winio v0.5.2 // indirect
	github.com/ProtonMail/go-crypto v0.0.0-20230518184743-7afd39499903 // indirect
	github.com/PuerkitoBio/purell v1.1.1 // indirect
	github.com/PuerkitoBio/urlesc v0.0.0-20170810143723-de5bf2ad4578 // indirect
	github.com/RichardKnop/logging v0.0.0-20181101035820-b1d5d44c82d6 // indirect
	github.com/RichardKnop/redsync v1.2.0 // indirect
	github.com/RoaringBitmap/roaring v0.4.23 // indirect
	github.com/StackExchange/wmi v0.0.0-20190523213315-cbe66965904d // indirect
	github.com/acomagu/bufpipe v1.0.4 // indirect
	github.com/alibabacloud-go/alibabacloud-gateway-spi v0.0.4 // indirect
	github.com/alibabacloud-go/debug v0.0.0-20190504072949-9472017b5c68 // indirect
	github.com/alibabacloud-go/endpoint-util v1.1.0 // indirect
	github.com/alibabacloud-go/openapi-util v0.0.11 // indirect
	github.com/alibabacloud-go/tea-xml v1.1.2 // indirect
	github.com/aliyun/credentials-go v1.1.2 // indirect
	github.com/andybalholm/cascadia v1.0.0 // indirect
	github.com/anmitsu/go-shlex v0.0.0-20200514113438-38f4b401e2be // indirect
	github.com/asaskevich/govalidator v0.0.0-20190424111038-f61b66f89f4a // indirect
	github.com/aws/aws-sdk-go v1.25.25 // indirect
	github.com/aymerick/douceur v0.2.0 // indirect
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/bgentry/speakeasy v0.1.0 // indirect
	github.com/blevesearch/go-porterstemmer v1.0.3 // indirect
	github.com/blevesearch/mmap-go v1.0.2 // indirect
	github.com/blevesearch/segment v0.9.0 // indirect
	github.com/blevesearch/snowballstem v0.9.0 // indirect
	github.com/blevesearch/zap/v11 v11.0.7 // indirect
	github.com/blevesearch/zap/v12 v12.0.7 // indirect
	github.com/boombuler/barcode v1.0.1-0.20190219062509-6c824513bacc // indirect
	github.com/bradfitz/gomemcache v0.0.0-20190329173943-551aad21a668 // indirect
	github.com/cenkalti/backoff/v5 v5.0.3 // indirect
	github.com/cespare/xxhash/v2 v2.3.0 // indirect
	github.com/chris-ramon/douceur v0.2.0 // indirect
	github.com/clbanning/mxj/v2 v2.5.5 // indirect
	github.com/cloudflare/circl v1.3.3 // indirect
	github.com/couchbase/gomemcached v0.0.0-20191004160342-7b5da2ec40b2 // indirect
	github.com/couchbase/goutils v0.0.0-20191018232750-b49639060d85 // indirect
	github.com/couchbase/vellum v1.0.1 // indirect
	github.com/couchbaselabs/go-couchbase v0.0.0-20190708161019-23e7ca2ce2b7 // indirect
	github.com/cpuguy83/go-md2man/v2 v2.0.0-20190314233015-f79a8a8ca69d // indirect
	github.com/cyphar/filepath-securejoin v0.2.4 // indirect
	github.com/cznic/b v0.0.0-20181122101859-a26611c4d92d // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/dlclark/regexp2 v1.4.0 // indirect
	github.com/fatih/color v1.9.0 // indirect
	github.com/fatih/structtag v1.2.0 // indirect
	github.com/felixge/httpsnoop v1.0.4 // indirect
	github.com/fsnotify/fsnotify v1.4.7 // indirect
	github.com/glycerine/go-unsnap-stream v0.0.0-20190901134440-81cf024a9e0a // indirect
	github.com/go-enry/go-oniguruma v1.2.0 // indirect
	github.com/go-git/gcfg v1.5.1-0.20230307220236-3a3c6141e376 // indirect
	github.com/go-ini/ini v1.56.0 // indirect
	github.com/go-logr/logr v1.4.3 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/go-macaron/inject v0.0.0-20160627170012-d8a0b8677191 // indirect
	github.com/go-ole/go-ole v1.2.4 // indirect
	github.com/go-openapi/analysis v0.19.5 // indirect
	github.com/go-openapi/errors v0.19.2 // indirect
	github.com/go-openapi/inflect v0.19.0 // indirect
	github.com/go-openapi/jsonpointer v0.19.3 // indirect
	github.com/go-openapi/jsonreference v0.19.3 // indirect
	github.com/go-openapi/loads v0.19.3 // indirect
	github.com/go-openapi/runtime v0.19.5 // indirect
	github.com/go-openapi/spec v0.19.3 // indirect
	github.com/go-openapi/strfmt v0.19.3 // indirect
	github.com/go-openapi/swag v0.19.5 // indirect
	github.com/go-openapi/validate v0.19.3 // indirect
	github.com/go-stack/stack v1.8.0 // indirect
	github.com/golang-sql/civil v0.0.0-20190719163853-cb61b32ac6fe // indirect
	github.com/golang/groupcache v0.0.0-20210331224755-41bb18bfe9da // indirect
	github.com/golang/protobuf v1.5.4 // indirect
	github.com/golang/snappy v0.0.4 // indirect
	github.com/google/s2a-go v0.1.9 // indirect
	github.com/googleapis/enterprise-certificate-proxy v0.3.8 // indirect
	github.com/googleapis/gax-go/v2 v2.16.0 // indirect
	github.com/gorilla/css v1.0.0 // indirect
	github.com/gorilla/handlers v1.4.2 // indirect
	github.com/gorilla/mux v1.6.2 // indirect
	github.com/gorilla/securecookie v1.1.1 // indirect
	github.com/gorilla/sessions v1.2.0 // indirect
	github.com/grafana/pyroscope-go/godeltaprof v0.1.9 // indirect
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.27.3 // indirect
	github.com/hashicorp/go-cleanhttp v0.5.1 // indirect
	github.com/hashicorp/hcl v1.0.0 // indirect
	github.com/imdario/mergo v0.3.15 // indirect
	github.com/issue9/assert v1.3.2 // indirect
	github.com/jbenet/go-context v0.0.0-20150711004518-d14ea06fba99 // indirect
	github.com/jessevdk/go-flags v1.5.0 // indirect
	github.com/jmespath/go-jmespath v0.0.0-20180206201540-c2b33e8439af // indirect
	github.com/jmhodges/levigo v1.0.0 // indirect
	github.com/joho/godotenv v1.3.0 // indirect
	github.com/json-iterator/go v1.1.10 // indirect
	github.com/kelseyhightower/envconfig v1.3.0 // indirect
	github.com/kevinburke/ssh_config v1.2.0 // indirect
	github.com/klauspost/cpuid v1.2.3 // indirect
	github.com/kr/pretty v0.3.1 // indirect
	github.com/kr/text v0.2.0 // indirect
	github.com/lunny/log v0.0.0-20160921050905-7887c61bf0de // indirect
	github.com/lunny/nodb v0.0.0-20160621015157-fc1ef06ad4af // indirect
	github.com/magiconair/properties v1.8.1 // indirect
	github.com/mailru/easyjson v0.7.0 // indirect
	github.com/mattn/go-colorable v0.1.4 // indirect
	github.com/mattn/go-oci8 v0.0.0-20190320171441-14ba190cf52d // indirect
	github.com/mattn/go-runewidth v0.0.7 // indirect
	github.com/matttproud/golang_protobuf_extensions v1.0.1 // indirect
	github.com/minio/md5-simd v1.1.0 // indirect
	github.com/minio/sha256-simd v0.1.1 // indirect
	github.com/mitchellh/mapstructure v1.1.2 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.1 // indirect
	github.com/mohae/deepcopy v0.0.0-20170929034955-c48cc78d4826 // indirect
	github.com/mrjones/oauth v0.0.0-20180629183705-f4e24b6d100c // indirect
	github.com/mschoch/smat v0.2.0 // indirect
	github.com/olekukonko/tablewriter v0.0.4 // indirect
	github.com/opentracing/opentracing-go v1.1.0 // indirect
	github.com/pelletier/go-toml v1.4.0 // indirect
	github.com/philhofer/fwd v1.0.0 // indirect
	github.com/pjbgf/sha1cd v0.3.0 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/prometheus/client_model v0.0.0-20190812154241-14fe0d1b01d4 // indirect
	github.com/prometheus/common v0.6.0 // indirect
	github.com/prometheus/procfs v0.0.4 // indirect
	github.com/rogpeppe/go-internal v1.14.1 // indirect
	github.com/russross/blackfriday/v2 v2.0.1 // indirect
	github.com/shurcooL/httpfs v0.0.0-20190527155220-6a4d4a70508b // indirect
	github.com/shurcooL/sanitized_anchor_name v1.0.0 // indirect
	github.com/siddontang/go-snappy v0.0.0-20140704025258-d8f7bb82a96d // indirect
	github.com/skeema/knownhosts v1.1.1 // indirect
	github.com/spf13/afero v1.2.2 // indirect
	github.com/spf13/cast v1.3.0 // indirect
	github.com/spf13/jwalterweatherman v1.1.0 // indirect
	github.com/spf13/pflag v1.0.3 // indirect
	github.com/spf13/viper v1.4.0 // indirect
	github.com/steveyen/gtreap v0.1.0 // indirect
	github.com/streadway/amqp v0.0.0-20190214183023-884228600bc9 // indirect
	github.com/syndtr/goleveldb v1.0.0 // indirect
	github.com/tecbot/gorocksdb v0.0.0-20181010114359-8752a9433481 // indirect
	github.com/tinylib/msgp v1.1.2 // indirect
	github.com/tjfoc/gmsm v1.3.2 // indirect
	github.com/toqueteos/trie v1.0.0 // indirect
	github.com/toqueteos/webbrowser v1.2.0 // indirect
	github.com/willf/bitset v1.1.10 // indirect
	github.com/xanzy/ssh-agent v0.3.3 // indirect
	github.com/xdg/scram v0.0.0-20180814205039-7eeb5667e42c // indirect
	github.com/xdg/stringprep v1.0.0 // indirect
	go.etcd.io/bbolt v1.3.4 // indirect
	go.mongodb.org/mongo-driver v1.1.1 // indirect
	go.opencensus.io v0.24.0 // indirect
	go.opentelemetry.io/auto/sdk v1.2.1 // indirect
	go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc v0.64.0 // indirect
	go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp v0.64.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlptrace v1.39.0 // indirect
	go.opentelemetry.io/proto/otlp v1.9.0 // indirect
	golang.org/x/image v0.0.0-20191009234506-e7c1f5e7dbb8 // indirect
	golang.org/x/mod v0.30.0 // indirect
	golang.org/x/sync v0.19.0 // indirect
	golang.org/x/tools/godoc v0.1.0-deprecated // indirect
	google.golang.org/api v0.259.0 // indirect
	google.golang.org/appengine v1.6.8 // indirect
	google.golang.org/genproto v0.0.0-20251222181119-0a764e51fe1b // indirect
	google.golang.org/genproto/googleapis/api v0.0.0-20251222181119-0a764e51fe1b // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20251222181119-0a764e51fe1b // indirect
	google.golang.org/grpc v1.78.0 // indirect
	google.golang.org/protobuf v1.36.11 // indirect
	gopkg.in/alexcesaro/quotedprintable.v3 v3.0.0-20150716171945-2caba252f4dc // indirect
	gopkg.in/asn1-ber.v1 v1.0.0-20150924051756-4e86f4367175 // indirect
	gopkg.in/macaron.v1 v1.3.9 // indirect
	gopkg.in/toqueteos/substring.v1 v1.0.2 // indirect
	gopkg.in/warnings.v0 v0.1.2 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
