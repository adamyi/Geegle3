load("@bazel_gazelle//:def.bzl", "gazelle")
load("@io_bazel_rules_docker//container:container.bzl", "container_bundle")
load("@io_bazel_rules_docker//contrib:push-all.bzl", "container_push")
load("@com_github_bazelbuild_buildtools//buildifier:def.bzl", "buildifier")

exports_files(["tsconfig.json"])

# gazelle:exclude dist
# gazelle:exclude node_modules
# gazelle:prefix geegle.org
gazelle(name = "gazelle")

buildifier(
    name = "buildifier",
    exclude_patterns = [
        "./dist/*",
        "./node_modules/*",
    ],
)

buildifier(
    name = "buildifier_check",
    exclude_patterns = [
        "./dist/*",
        "./node_modules/*",
    ],
    mode = "check",
)

container_bundle(
    name = "all_containers",
    images = {
        # infra
        "gcr.io/geegle/infra/mss:latest": "//infra/mss:image",
        "gcr.io/geegle/infra/sffe:latest": "//infra/sffe:image",
        "gcr.io/geegle/infra/geemail-backend:latest": "//infra/geemail-backend:image",
        "gcr.io/geegle/infra/mail:latest": "//infra/geemail-frontend:image",
        "gcr.io/geegle/infra/uberproxy:latest": "//infra/uberproxy:image",
        "gcr.io/geegle/infra/dns:latest": "//infra/dns:image",
        "gcr.io/geegle/infra/apps:latest": "//infra/gae:image",

        # web challenges
        "gcr.io/geegle/chals/web/docs:latest": "//chals/web/docs:image",
        "gcr.io/geegle/chals/web/seclearn:latest": "//chals/web/seclearn:image",
        "gcr.io/geegle/chals/web/pasteweb:latest": "//chals/web/pasteweb:image",
        "gcr.io/geegle/chals/web/memegen:latest": "//chals/web/memegen:image",
        "gcr.io/geegle/chals/web/flatearth:latest": "//chals/web/flatearth:image",
        "gcr.io/geegle/chals/web/employees:latest": "//chals/web/employees:image",

        # pwn challenges
        "gcr.io/geegle/chals/pwn/game:latest": "//chals/pwn/game:image",
        "gcr.io/geegle/chals/pwn/geelang:latest": "//chals/pwn/geelang:image",
        "gcr.io/geegle/chals/pwn/shell:latest": "//chals/pwn/shell:image",
        "gcr.io/geegle/chals/pwn/payroll:latest": "//chals/pwn/payroll:image",
        "gcr.io/geegle/chals/re/onboarding:latest": "//chals/re/tellGeegle:image",

        # others challenges
        "gcr.io/geegle/chals/ir/who-is-attacking-me:latest": "//chals/ir/who-is-attacking-me:image",
        "gcr.io/geegle/chals/pwn/intern-project:latest": "//chals/pwn/intern-project:image",
        "gcr.io/geegle/chals/web/guest:latest": "//chals/web/guest:image",
    },
)

container_push(
    name = "all_containers_push",
    bundle = ":all_containers",
    format = "Docker",
)
