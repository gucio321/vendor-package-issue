## Problem

## Starting conditions:

module1 simulates [cimgui-go](https://github.com/AllenDang/cimgui-go).
This repo uses external C dependencies which are however statically added to the repo (not submodules).
One of these dependencies use `vendor` as its folder for some other dependencies.

cimgui-go's workflow adds temporary `dummy.go` files with a phantom-modules to froce go toolchain to include all C files in go get.

module2 simulates any module using module1.

## Current situation

what happens is, that when module2 tries to get latest version of module1 (where I managed to reproduce my issue) by running `update.sh`,
the following happens

```console
$ ./update.sh
go: github.com/gucio321/vendor-package-issue/module1 imports
	github.com/gucio321/vendor-package-issue/module1/pkg/vendor: cannot find module providing package github.com/gucio321/vendor-package-issue/module1/pkg/vendor
```
