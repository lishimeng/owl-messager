//go:generate go-bindata -pkg static -o bindata.go --prefix ../ui/dist ../ui/dist/...
package static

// go-bindata -pkg static -o bindata.go --prefix ../ui/dist ../ui/dist/...
/**
-pkg 指定生成文件的包名
-o 指定生成文件名
--prefix 替换文件路径
 */