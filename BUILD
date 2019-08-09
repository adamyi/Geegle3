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
    },
)

container_push(
    name = "all_containers_push",
    bundle = ":all_containers",
    format = "Docker",
)
