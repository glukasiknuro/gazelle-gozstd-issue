load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")
load("@bazel_gazelle//:def.bzl", "gazelle")

# gazelle:prefix module github.com/glukasiknuro/gazelle-gozstd-issue
gazelle(name = "gazelle")

go_library(
    name = "gazelle-gozstd-issue_lib",
    srcs = ["main.go"],
    data = ["hello_world.zst"],
    importpath = "module github.com/glukasiknuro/gazelle-gozstd-issue",
    visibility = ["//visibility:private"],
    deps = ["@com_github_valyala_gozstd//:go_default_library"],
)

go_binary(
    name = "gazelle-gozstd-issue",
    embed = [":gazelle-gozstd-issue_lib"],
    visibility = ["//visibility:public"],
)
