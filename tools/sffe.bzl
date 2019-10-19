# author: adamyi

def sffe_files(files):
    native.filegroup(
        name = "sffe",
        srcs = files,
        visibility = ["//infra/sffe:__pkg__"],
    )
