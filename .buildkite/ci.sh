git-crypt unlock /g3.key
bazel run //:gazelle -- --mode=diff || (echo \"ERROR: Bazel files out-of-date, please run \\`bazel run :gazelle\\`\" >&2; exit 1)
bazel run //:buildifier_check || (echo \"ERROR: Bazel files not formatted, please run \\`bazel run :buildifier\\`\" >&2; exit 1)
bazel build //infra/jsonnet:all-docker-compose || (echo \"ERROR: Bazel build docker-compose failed\" >&2; exit 1)
bazel build //infra/jsonnet:cluster-master-docker-compose || (echo \"ERROR: Bazel build docker-compose failed\" >&2; exit 1)
bazel build //infra/jsonnet:cluster-team-docker-compose || (echo \"ERROR: Bazel build docker-compose failed\" >&2; exit 1)
bazel build //:all_containers || (echo \"ERROR: Bazel build all_containers failed\" >&2; exit 1)
