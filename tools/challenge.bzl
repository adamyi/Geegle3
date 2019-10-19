# author: adamyi

load("@io_bazel_rules_jsonnet//jsonnet:jsonnet.bzl", "jsonnet_library")

def ctf_challenge():
    jsonnet_library(
        name = "challenge",
        srcs = [
            "challenge.libsonnet",
        ],
        visibility = ["//chals:__pkg__", "//infra:__pkg__"],
    )
