load("@bazel_gazelle//:def.bzl", "gazelle")
load("@io_bazel_rules_docker//container:container.bzl", "container_bundle")
load("@io_bazel_rules_docker//contrib:push-all.bzl", "container_push")
load("@com_github_bazelbuild_buildtools//buildifier:def.bzl", "buildifier")

exports_files(["tsconfig.json"])

# gazelle:exclude dist
# gazelle:prefix geegle.org
gazelle(name = "gazelle")

buildifier(
    name = "buildifier",
)

buildifier(
    name = "buildifier_check",
    mode = "check",
)

container_bundle(
    name = "all_containers",
    images = {
        "gcr.io/geegle/infra/mss:latest": "//infra/mss:image",
        "gcr.io/geegle/infra/sffe:latest": "//infra/sffe:image",
        "gcr.io/geegle/infra/geemail-backend:latest": "//infra/geemail-backend:image",
        "gcr.io/geegle/infra/geemail-frontend:latest": "//infra/geemail-frontend:image",
        "gcr.io/geegle/infra/uberproxy:latest": "//infra/uberproxy:image",
        "gcr.io/geegle/infra/gae:latest": "//infra/gae:image",
        "gcr.io/geegle/highschool/crypto-etcpasswd:latest": "//highschool/crypto-etcpasswd:image",
        "gcr.io/geegle/highschool/misc-quicksort:latest": "//highschool/misc-quicksort:image",
        "gcr.io/geegle/highschool/misc-who-is-attacking-me:latest": "//highschool/misc-who-is-attacking-me:image",
        "gcr.io/geegle/highschool/pwn-intern-project:latest": "//highschool/pwn-intern-project:image",
        "gcr.io/geegle/highschool/web-filesystem:latest": "//highschool/web-filesystem:image",
        "gcr.io/geegle/highschool/web-guest-kiosk/:latest": "//highschool/web-guest-kiosk/app:image",
        "gcr.io/geegle/highschool/web-intern-account-manager:latest": "//highschool/web-intern-account-manager/app:image",
        "gcr.io/geegle/highschool/web-intranet:latest": "//highschool/web-intranet/app:image",
        "gcr.io/geegle/highschool/web-privatefile:latest": "//highschool/web-privatefile/app:image",
    },
)

container_push(
    name = "all_containers_push",
    bundle = ":all_containers",
    format = "Docker",
)
