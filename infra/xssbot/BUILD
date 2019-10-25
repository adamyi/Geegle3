load("@io_bazel_rules_docker//nodejs:image.bzl", "nodejs_image")
load("@io_bazel_rules_docker//container:container.bzl", "container_image")

container_image(
    name = "image",
    base = ":nodejs_image",
    tars = [
        "@chromium",
    ],
    visibility = ["//visibility:public"],
)

nodejs_image(
    name = "nodejs_image",
    data = [":server.js"],
    entry_point = "server.js",
    node_modules = "@npm//:node_modules",
)
