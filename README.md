# go_mruby_test
Test go_mruby package
IMPORTANT: When you run go get command you may get error. To fix that, you need to copy file mruby/libmruby.a to your project root directory. Example bash command to run on your project root directory:
```sh
export GOPATH=$(go env GOPATH)
mkdir mruby
cp $GOPATH/pkg/mod/webimizer.dev/go_mruby@v1.0.0/mruby/libmruby.a ./mruby/libmruby.a
```
