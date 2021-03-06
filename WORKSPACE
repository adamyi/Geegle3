workspace(
    name = "geegle3",
    managed_directories = {"@npm": ["node_modules"]},
)

load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive", "http_file")

http_archive(
    name = "io_bazel_rules_go",
    sha256 = "842ec0e6b4fbfdd3de6150b61af92901eeb73681fd4d185746644c338f51d4c0",
    urls = [
        "https://storage.googleapis.com/bazel-mirror/github.com/bazelbuild/rules_go/releases/download/v0.20.1/rules_go-v0.20.1.tar.gz",
        "https://github.com/bazelbuild/rules_go/releases/download/v0.20.1/rules_go-v0.20.1.tar.gz",
    ],
)

load("@io_bazel_rules_go//go:deps.bzl", "go_register_toolchains", "go_rules_dependencies")

go_rules_dependencies()

go_register_toolchains()

# memegen requires https://github.com/bazelbuild/rules_docker/pull/1200
# which isn't in release yet
# TODO(adamyi): change to a stable release
http_archive(
    name = "io_bazel_rules_docker",
    sha256 = "5f9b5fc1431f03e5b62541cf691e6b2311dff0e698ab8241a777199351a51ad7",
    strip_prefix = "rules_docker-e878e185bef129391d7847076bd1d377d5c16b41",
    urls = ["https://github.com/bazelbuild/rules_docker/archive/e878e185bef129391d7847076bd1d377d5c16b41.tar.gz"],
    #urls = ["https://github.com/bazelbuild/rules_docker/releases/download/v0.12.0/rules_docker-v0.12.0.tar.gz"],
)

load(
    "@io_bazel_rules_docker//repositories:repositories.bzl",
    container_repositories = "repositories",
)

container_repositories()

http_archive(
    name = "bazel_gazelle",
    sha256 = "41bff2a0b32b02f20c227d234aa25ef3783998e5453f7eade929704dcff7cd4b",
    urls = [
        "https://storage.googleapis.com/bazel-mirror/github.com/bazelbuild/bazel-gazelle/releases/download/v0.19.0/bazel-gazelle-v0.19.0.tar.gz",
        "https://github.com/bazelbuild/bazel-gazelle/releases/download/v0.19.0/bazel-gazelle-v0.19.0.tar.gz",
    ],
)

load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies", "go_repository")

gazelle_dependencies()

http_archive(
    name = "com_google_protobuf",
    sha256 = "e11f901c62f6a35e295b7e9c49123a96ef7a47503afd40ed174860ad4c3f294f",
    strip_prefix = "protobuf-3.10.0",
    urls = ["https://github.com/protocolbuffers/protobuf/releases/download/v3.10.0/protobuf-all-3.10.0.tar.gz"],
)

load("@com_google_protobuf//:protobuf_deps.bzl", "protobuf_deps")

protobuf_deps()

http_archive(
    name = "com_github_bazelbuild_buildtools",
    strip_prefix = "buildtools-master",
    url = "https://github.com/bazelbuild/buildtools/archive/master.zip",
)

load("@bazel_tools//tools/build_defs/repo:git.bzl", "git_repository")

http_archive(
    name = "rules_python",
    sha256 = "e220053c4454664c09628ffbb33f245e65f5fe92eb285fbd0bc3a26f173f99d0",
    strip_prefix = "rules_python-5aa465d5d91f1d9d90cac10624e3d2faf2057bd5",
    urls = ["https://github.com/bazelbuild/rules_python/archive/5aa465d5d91f1d9d90cac10624e3d2faf2057bd5.tar.gz"],
)

RULES_NODEJS_VERSION = "0.39.0"

RULES_NODEJS_SHA256 = "26c39450ce2d825abee5583a43733863098ed29d3cbaebf084ebaca59a21a1c8"

http_archive(
    name = "build_bazel_rules_nodejs",
    sha256 = RULES_NODEJS_SHA256,
    url = "https://github.com/bazelbuild/rules_nodejs/releases/download/%s/rules_nodejs-%s.tar.gz" % (RULES_NODEJS_VERSION, RULES_NODEJS_VERSION),
)

load("@io_bazel_rules_docker//container:pull.bzl", "container_pull")
load("@io_bazel_rules_docker//java:image.bzl", _java_image_repos = "repositories")
load("@io_bazel_rules_docker//go:image.bzl", _go_image_repos = "repositories")
load("@io_bazel_rules_docker//python:image.bzl", _py_image_repos = "repositories")
load("@io_bazel_rules_docker//nodejs:image.bzl", _nodejs_image_repos = "repositories")

_java_image_repos()

_go_image_repos()

_py_image_repos()

_nodejs_image_repos()

http_archive(
    name = "io_bazel_rules_jsonnet",
    sha256 = "68b5bcb0779599065da1056fc8df60d970cffe8e6832caf13819bb4d6e832459",
    strip_prefix = "rules_jsonnet-0.2.0",
    urls = ["https://github.com/bazelbuild/rules_jsonnet/archive/0.2.0.tar.gz"],
)

load("@io_bazel_rules_jsonnet//jsonnet:jsonnet.bzl", "jsonnet_repositories")

jsonnet_repositories()

load("@jsonnet_go//bazel:repositories.bzl", "jsonnet_go_repositories")

jsonnet_go_repositories()

load("@jsonnet_go//bazel:deps.bzl", "jsonnet_go_dependencies")

jsonnet_go_dependencies()

http_archive(
    name = "base_images_docker",
    sha256 = "ce6043d38aa7fad421910311aecec865beb060eb56d8c3eb5af62b2805e9379c",
    strip_prefix = "base-images-docker-7657d04ad9e30b9b8d981b96ae57634cd45ba18a",
    urls = ["https://github.com/GoogleCloudPlatform/base-images-docker/archive/7657d04ad9e30b9b8d981b96ae57634cd45ba18a.tar.gz"],
)

container_pull(
    name = "tomcat9",
    registry = "index.docker.io",
    repository = "tomcat",
    tag = "9.0.21-jdk8",
)

container_pull(
    name = "tomcat-jython",
    digest = "sha256:27526ffde703e09cdf8cbb3cb781c169ac48f2e2ba3a6fbe3238c9fff9b80fc7",
    registry = "index.docker.io",
    repository = "adamyi/tomcat-jython",
)

container_pull(
    name = "python-with-latex",
    digest = "sha256:db92134dd530dd3b666a5b420029886c01faf72a7ea366726c6eeac45ae4ed64",
    registry = "index.docker.io",
    repository = "adamyi/python-with-latex",
)

container_pull(
    name = "nginx-php-fpm-with-imagick",
    digest = "sha256:2ac175a4b6faff45ca12325de3ed3899c8acba2d38955fab5b8b877c8cb7c6d5",
    registry = "index.docker.io",
    repository = "adamyi/nginx-php-fpm-with-imagick",
)

container_pull(
    name = "ubuntu1804-with-32bit-libc",
    digest = "sha256:3225563499e60d3bacd4db8f05920ae5d86635372d1c77024fd73d6db9d04cca",
    registry = "index.docker.io",
    repository = "adamyi/ubuntu1804-with-32bit-libc",
)

container_pull(
    name = "ubuntu1804-with-zbar",
    digest = "sha256:cc47d8fc8309178954287c6419f3f39aa3741b6c540c351bd5d71c3662b9d6ba",
    registry = "index.docker.io",
    repository = "adamyi/ubuntu1804-with-zbar",
)

container_pull(
    name = "chrome-base-without-chrome",
    digest = "sha256:b5c86894a56352eb4f91c462d7cb95b5475b3e3735d4faea9893cfaca668c467",
    registry = "index.docker.io",
    repository = "adamyi/chrome-base-without-chrome",
)

container_pull(
    name = "nginx-php-fpm",
    digest = "sha256:2e9718f4bdca05f577cb8cf046327cb9232e4fd817fe32f470db0a65660a6e46",
    registry = "index.docker.io",
    repository = "richarvey/nginx-php-fpm",
)

container_pull(
    name = "alpine_linux_amd64",
    registry = "index.docker.io",
    repository = "library/alpine",
    tag = "3.8",
)

container_pull(
    name = "ubuntu1804",
    digest = "sha256:3942b604b2f23e9b08aa6f3c51dc19efa2b570ae93ce8aaabf94e02111ddcca9",
    registry = "gcr.io",
    repository = "cloud-marketplace/google/ubuntu1804",
)

container_pull(
    name = "python2-base",
    digest = "sha256:938d21703d929295337f5aafd038a8d93172e11e1746f6e87f02eb53e32bcea0",
    registry = "index.docker.io",
    repository = "python",
)

container_pull(
    name = "python3-base",
    digest = "sha256:d182a775e372d82d92b59ff5debeabcb699964fe163320eb9fc5ebb971c51ec3",
    registry = "index.docker.io",
    repository = "python",
)

go_repository(
    name = "com_github_dgrijalva_jwt_go",
    importpath = "github.com/dgrijalva/jwt-go",
    tag = "v3.2.0",
)

go_repository(
    name = "com_github_gomodule_redigo",
    importpath = "github.com/gomodule/redigo",
    tag = "v1.7.0",
)

go_repository(
    name = "com_github_gorilla_mux",
    importpath = "github.com/gorilla/mux",
    tag = "v1.7.3",
)

go_repository(
    name = "com_github_gorilla_websocket",
    importpath = "github.com/gorilla/websocket",
    tag = "v1.4.0",
)

go_repository(
    name = "com_github_clsung_grcode",
    commit = "e7f3c16fb531",
    importpath = "github.com/clsung/grcode",
)

go_repository(
    name = "com_github_tuotoo_qrcode",
    commit = "ac9c44189bf2",
    importpath = "github.com/tuotoo/qrcode",
)

go_repository(
    name = "com_github_maruel_rs",
    commit = "2c81c4312fe4",
    importpath = "github.com/maruel/rs",
)

go_repository(
    name = "com_github_google_uuid",
    importpath = "github.com/google/uuid",
    tag = "v1.1.1",
)

go_repository(
    name = "com_github_skip2_go_qrcode",
    commit = "dc11ecdae0a9",
    importpath = "github.com/skip2/go-qrcode",
)

go_repository(
    name = "com_github_syndtr_goleveldb",
    commit = "02440ea7a285",
    importpath = "github.com/syndtr/goleveldb",
)

go_repository(
    name = "com_github_mattn_go_sqlite3",
    importpath = "github.com/mattn/go-sqlite3",
    tag = "v1.10.0",
)

go_repository(
    name = "com_github_go_sql_driver_mysql",
    importpath = "github.com/go-sql-driver/mysql",
    tag = "v1.4.1",
)

go_repository(
    name = "com_github_golang_snappy",
    importpath = "github.com/golang/snappy",
    tag = "v0.0.1",
)

go_repository(
    name = "com_sajari_code_word2vec",
    importpath = "code.sajari.com/word2vec",
    tag = "v1.0.0",
)

go_repository(
    name = "com_github_ziutek_blas",
    commit = "da4ca23e90bb",
    importpath = "github.com/ziutek/blas",
)

go_repository(
    name = "com_github_joewalnes_websocketd",
    importpath = "github.com/joewalnes/websocketd",
    tag = "v0.3.1",
)

go_repository(
    name = "com_github_miekg_dns",
    importpath = "github.com/miekg/dns",
    tag = "v1.1.22",
)

go_repository(
    name = "org_golang_x_crypto",
    commit = "9ee001bba392",
    importpath = "golang.org/x/crypto",
)

go_repository(
    name = "org_golang_x_net",
    commit = "aa69164e4478",
    importpath = "golang.org/x/net",
)

go_repository(
    name = "org_golang_x_sync",
    commit = "112230192c58",
    importpath = "golang.org/x/sync",
)

go_repository(
    name = "org_golang_x_sys",
    commit = "2837fb4f24fe",
    importpath = "golang.org/x/sys",
)

go_repository(
    name = "org_golang_x_text",
    importpath = "golang.org/x/text",
    tag = "v0.3.2",
)

go_repository(
    name = "org_golang_x_tools",
    commit = "2ca718005c18",
    importpath = "golang.org/x/tools",
)

go_repository(
    name = "com_github_emersion_go_smtp",
    importpath = "github.com/emersion/go-smtp",
    tag = "v0.11.2",
)

go_repository(
    name = "com_github_emersion_go_sasl",
    commit = "36b50694675c",
    importpath = "github.com/emersion/go-sasl",
)

go_repository(
    name = "com_github_jhillyerd_enmime",
    importpath = "github.com/jhillyerd/enmime",
    tag = "v0.6.0",
)

go_repository(
    name = "com_github_gogs_chardet",
    commit = "2404f7772561",
    importpath = "github.com/gogs/chardet",
)

go_repository(
    name = "com_github_jaytaylor_html2text",
    commit = "57d518f124b0",
    importpath = "github.com/jaytaylor/html2text",
)

go_repository(
    name = "com_github_ssor_bom",
    commit = "6386211fdfcf",
    importpath = "github.com/ssor/bom",
)

go_repository(
    name = "com_github_olekukonko_tablewriter",
    commit = "be2c049b30cc",
    importpath = "github.com/olekukonko/tablewriter",
)

go_repository(
    name = "com_github_mattn_go_runewidth",
    importpath = "github.com/mattn/go-runewidth",
    tag = "v0.0.3",
)

go_repository(
    name = "cc_mvdan_xurls_v2",
    importpath = "github.com/mvdan/xurls",
    tag = "v2.1.0",
)

# Rules for compiling sass
RULES_SASS_VERSION = "86ca977cf2a8ed481859f83a286e164d07335116"

RULES_SASS_SHA256 = "4f05239080175a3f4efa8982d2b7775892d656bb47e8cf56914d5f9441fb5ea6"

http_archive(
    name = "io_bazel_rules_sass",
    sha256 = RULES_SASS_SHA256,
    strip_prefix = "rules_sass-%s" % RULES_SASS_VERSION,
    url = "https://github.com/bazelbuild/rules_sass/archive/%s.zip" % RULES_SASS_VERSION,
)

####################################
# Load and install our dependencies downloaded above.

load("@build_bazel_rules_nodejs//:defs.bzl", "node_repositories", "yarn_install")

# Setup the Node repositories. We need a NodeJS version that is more recent than v10.15.0
# because "selenium-webdriver" which is required for "ng e2e" cannot be installed.
# TODO: remove the custom repositories once "rules_nodejs" supports v10.16.0 by default.
node_repositories(
    node_repositories = {
        "10.16.0-darwin_amd64": ("node-v10.16.0-darwin-x64.tar.gz", "node-v10.16.0-darwin-x64", "6c009df1b724026d84ae9a838c5b382662e30f6c5563a0995532f2bece39fa9c"),
        "10.16.0-linux_amd64": ("node-v10.16.0-linux-x64.tar.xz", "node-v10.16.0-linux-x64", "1827f5b99084740234de0c506f4dd2202a696ed60f76059696747c34339b9d48"),
        "10.16.0-windows_amd64": ("node-v10.16.0-win-x64.zip", "node-v10.16.0-win-x64", "aa22cb357f0fb54ccbc06b19b60e37eefea5d7dd9940912675d3ed988bf9a059"),
    },
    node_version = "10.16.0",
)

yarn_install(
    name = "npm",
    package_json = "//:package.json",
    yarn_lock = "//:yarn.lock",
)

load("@npm//:install_bazel_dependencies.bzl", "install_bazel_dependencies")

install_bazel_dependencies()

load("@npm_bazel_karma//:package.bzl", "npm_bazel_karma_dependencies")

npm_bazel_karma_dependencies()

load("@io_bazel_rules_webtesting//web:repositories.bzl", "web_test_repositories")

web_test_repositories()

load("@npm_bazel_typescript//:index.bzl", "ts_setup_workspace")

ts_setup_workspace()

load("@io_bazel_rules_sass//sass:sass_repositories.bzl", "sass_repositories")

sass_repositories()

load("@rules_python//python:pip.bzl", "pip_import")

pip_import(
    name = "docs_pip",
    requirements = "//chals/web/docs:app/requirements.txt",
)

pip_import(
    name = "pasteweb_pip",
    requirements = "//chals/web/pasteweb:app/requirements.txt",
)

pip_import(
    name = "search_pip",
    requirements = "//chals/web/search:app/requirements.txt",
)

load(
    "@docs_pip//:requirements.bzl",
    _docs_install = "pip_install",
)

_docs_install()

load(
    "@pasteweb_pip//:requirements.bzl",
    _pasteweb_install = "pip_install",
)

_pasteweb_install()

load(
    "@search_pip//:requirements.bzl",
    _search_install = "pip_install",
)

_search_install()

http_archive(
    name = "chromium",
    build_file = "@//third_party:chromium.BUILD",
    sha256 = "10ae4e05d9f01a8b646dd2ccc2ac1135e597c472abe5be71552aae7d8a35e2ac",
    url = "https://www.googleapis.com/download/storage/v1/b/chromium-browser-snapshots/o/Linux_x64%2F650583%2Fchrome-linux.zip?generation=1555131417316559&alt=media",
)
