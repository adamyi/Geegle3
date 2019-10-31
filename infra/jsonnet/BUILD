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
    name = "all-docker-compose",
    src = "docker-compose.jsonnet",
    outs = ["all-docker-compose.json"],
    ext_strs = {
        "cluster": "all",
    },
    deps = [
        ":utils",
        "//chals:challenges_jsonnet",
        "//infra:infra_jsonnet",
    ],
)

jsonnet_to_json(
    name = "cluster-master-docker-compose",
    src = "docker-compose.jsonnet",
    outs = ["cluster-master-docker-compose.json"],
    ext_strs = {
        "cluster": "master",
    },
    deps = [
        ":utils",
        "//chals:challenges_jsonnet",
        "//infra:infra_jsonnet",
    ],
)

jsonnet_to_json(
    name = "cluster-team-docker-compose",
    src = "docker-compose.jsonnet",
    outs = ["cluster-team-docker-compose.json"],
    ext_strs = {
        "cluster": "team",
    },
    deps = [
        ":utils",
        "//chals:challenges_jsonnet",
        "//infra:infra_jsonnet",
    ],
)
