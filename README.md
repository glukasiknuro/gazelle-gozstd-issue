Relevant bug: https://github.com/bazelbuild/bazel-gazelle/issues/1120

## Problematic output

```shell
$ bazel-4.2.0 run //:gazelle-gozstd-issue

gcc: error: ./libzstd_linux_amd64.a: No such file or directory
compilepkg: error running subcommand /usr/bin/gcc: exit status 1
Target //:gazelle-gozstd-issue failed to build
```

## Fix
Uncomment those lines in WORKSPACE:

```python
    # Uncomment below lines to fix issue:
    # build_file_generation = "off", # keep
    # patch_args = ["-p1"],  # keep
    # patches = ["gozstd.patch"],  # keep
```

Result:
```shell
$ bazel-4.2.0 run //:gazelle-gozstd-issue

Got string from zstd file: hello world
```