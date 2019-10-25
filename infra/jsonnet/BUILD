load("@io_bazel_rules_jsonnet//jsonnet:jsonnet.bzl", "jsonnet_library", "jsonnet_to_json")

exports_files(["cli-static-sffe.jsonnet"])

jsonnet_library(
    name = "utils",
    srcs = [
        "utils.libsonnet",
    ],
    visibility = ["//visibility:public"],
)

jsonnet_to_json(
    name = "docker-compose",
    src = "docker-compose.jsonnet",
    outs = ["docker-compose.json"],
    deps = [
        ":utils",
        "//chals:challenges_jsonnet",
        "//infra:infra_jsonnet",
    ],
)
