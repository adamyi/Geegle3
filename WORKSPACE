workspace(
    name = "geegle3",
    managed_directories = {"@npm": ["node_modules"]},
)

load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive", "http_file")

http_archive(
    name = "io_bazel_rules_docker",
    sha256 = "e513c0ac6534810eb7a14bf025a0f159726753f97f74ab7863c650d26e01d677",
    strip_prefix = "rules_docker-0.9.0",
    urls = ["https://github.com/bazelbuild/rules_docker/archive/v0.9.0.tar.gz"],
)

load(
    "@io_bazel_rules_docker//repositories:repositories.bzl",
    container_repositories = "repositories",
)

container_repositories()

http_archive(
    name = "io_bazel_rules_go",
    sha256 = "8df59f11fb697743cbb3f26cfb8750395f30471e9eabde0d174c3aebc7a1cd39",
    urls = [
        "https://storage.googleapis.com/bazel-mirror/github.com/bazelbuild/rules_go/releases/download/0.19.1/rules_go-0.19.1.tar.gz",
        "https://github.com/bazelbuild/rules_go/releases/download/0.19.1/rules_go-0.19.1.tar.gz",
    ],
)

load("@io_bazel_rules_go//go:deps.bzl", "go_register_toolchains", "go_rules_dependencies")

go_rules_dependencies()

go_register_toolchains()

http_archive(
    name = "bazel_gazelle",
    sha256 = "3c681998538231a2d24d0c07ed5a7658cb72bfb5fd4bf9911157c0e9ac6a2687",
    urls = ["https://github.com/bazelbuild/bazel-gazelle/releases/download/0.17.0/bazel-gazelle-0.17.0.tar.gz"],
)

load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies", "go_repository")

gazelle_dependencies()

http_archive(
    name = "com_google_protobuf",
    sha256 = "3040a5b946d9df7aa89c0bf6981330bf92b7844fd90e71b61da0c721e421a421",
    strip_prefix = "protobuf-3.9.1",
    urls = ["https://github.com/protocolbuffers/protobuf/releases/download/v3.9.1/protobuf-all-3.9.1.tar.gz"],
)

load("@com_google_protobuf//:protobuf_deps.bzl", "protobuf_deps")

protobuf_deps()

http_archive(
    name = "com_github_bazelbuild_buildtools",
    strip_prefix = "buildtools-master",
    url = "https://github.com/bazelbuild/buildtools/archive/master.zip",
)

load("@bazel_tools//tools/build_defs/repo:git.bzl", "git_repository")

git_repository(
    name = "io_bazel_rules_python",
    commit = "fdbb17a4118a1728d19e638a5291b4c4266ea5b8",
    remote = "https://github.com/bazelbuild/rules_python.git",
)

load("@io_bazel_rules_docker//container:new_pull.bzl", "new_container_pull")
load("@io_bazel_rules_docker//container:pull.bzl", "container_pull")
load("@io_bazel_rules_docker//java:image.bzl", _java_image_repos = "repositories")
load("@io_bazel_rules_docker//go:image.bzl", _go_image_repos = "repositories")
load("@io_bazel_rules_docker//python:image.bzl", _py_image_repos = "repositories")

_java_image_repos()

_go_image_repos()

_py_image_repos()

http_archive(
    name = "base_images_docker",
    sha256 = "ce6043d38aa7fad421910311aecec865beb060eb56d8c3eb5af62b2805e9379c",
    strip_prefix = "base-images-docker-7657d04ad9e30b9b8d981b96ae57634cd45ba18a",
    urls = ["https://github.com/GoogleCloudPlatform/base-images-docker/archive/7657d04ad9e30b9b8d981b96ae57634cd45ba18a.tar.gz"],
)

new_container_pull(
    name = "tomcat9",
    registry = "index.docker.io",
    repository = "tomcat",
    tag = "9.0.21-jdk8",
)

container_pull(
    name = "tomcat-jython",
    registry = "index.docker.io",
    repository = "adamyi/tomcat-jython",
    tag = "latest",
)

container_pull(
    name = "python-with-latex",
    registry = "index.docker.io",
    repository = "adamyi/python-with-latex",
    tag = "latest",
)

container_pull(
    name = "nginx-php-fpm-with-imagick",
    registry = "index.docker.io",
    repository = "adamyi/nginx-php-fpm-with-imagick",
    tag = "latest",
)

container_pull(
    name = "nginx-php-fpm",
    registry = "index.docker.io",
    repository = "richarvey/nginx-php-fpm",
    tag = "1.7.4",
)

new_container_pull(
    name = "alpine_linux_amd64",
    registry = "index.docker.io",
    repository = "library/alpine",
    tag = "3.8",
)

new_container_pull(
    name = "ubuntu1804",
    registry = "gcr.io",
    repository = "cloud-marketplace/google/ubuntu1804",
    tag = "latest",
)

new_container_pull(
    name = "python2-base",
    registry = "index.docker.io",
    repository = "python",
    tag = "2.7",
)

new_container_pull(
    name = "python3-base",
    registry = "index.docker.io",
    repository = "python",
    tag = "3.7",
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

RULES_NODEJS_VERSION = "0.32.2"

RULES_NODEJS_SHA256 = "6d4edbf28ff6720aedf5f97f9b9a7679401bf7fca9d14a0fff80f644a99992b4"

http_archive(
    name = "build_bazel_rules_nodejs",
    sha256 = RULES_NODEJS_SHA256,
    url = "https://github.com/bazelbuild/rules_nodejs/releases/download/%s/rules_nodejs-%s.tar.gz" % (RULES_NODEJS_VERSION, RULES_NODEJS_VERSION),
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

load("@npm_bazel_karma//:package.bzl", "rules_karma_dependencies")

rules_karma_dependencies()

load("@io_bazel_rules_webtesting//web:repositories.bzl", "web_test_repositories")

web_test_repositories()

load("@npm_bazel_karma//:browser_repositories.bzl", "browser_repositories")

browser_repositories()

load("@npm_bazel_typescript//:index.bzl", "ts_setup_workspace")

ts_setup_workspace()

load("@io_bazel_rules_sass//sass:sass_repositories.bzl", "sass_repositories")

sass_repositories()

load("@io_bazel_rules_python//python:pip.bzl", "pip_import")

pip_import(
    name = "gae_pip",
    requirements = "//infra/gae:src/requirements.txt",
)

pip_import(
    name = "kix_pip",
    requirements = "//advanced/web_kix:app/requirements.txt",
)

pip_import(
    name = "pasteweb_pip",
    requirements = "//advanced/web_pasteweb:app/requirements.txt",
)

load(
    "@gae_pip//:requirements.bzl",
    _gae_install = "pip_install",
)

_gae_install()

load(
    "@kix_pip//:requirements.bzl",
    _kix_install = "pip_install",
)

_kix_install()

load(
    "@pasteweb_pip//:requirements.bzl",
    _pasteweb_install = "pip_install",
)

_pasteweb_install()

http_archive(
    name = "websocketd",
    build_file = "//third_party:BUILD.websocketd",
    sha256 = "fdadea9886e942c1d766aaa4303c3b8fe746caa66e7d012f726bbdb71e2cef3a",
    url = "https://github.com/joewalnes/websocketd/releases/download/v0.3.0/websocketd-0.3.0-linux_amd64.zip",
)
