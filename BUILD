load("@bazel_gazelle//:def.bzl", "gazelle")

# gazelle:prefix github.com/clashthebunny/advent-of-code-golang
gazelle(name = "gazelle")

gazelle(
    name = "gazelle-update-repos",
    args = [
        "-from_file=go.mod",
        "-to_macro=deps.bzl%go_dependencies",
        "-prune",
        "-build_file_proto_mode=disable_global",
    ],
    command = "update-repos",
    visibility = ["//visibility:public"],
)
