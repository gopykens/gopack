module github.com/goplus/gop

go 1.16

replace (
	github.com/goplus/gox => ../gox
)

require (
	github.com/goplus/gox v1.7.10
	github.com/qiniu/x v1.11.5
	golang.org/x/mod v0.5.1
	golang.org/x/tools v0.1.7
)
