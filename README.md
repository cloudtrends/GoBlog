#Fxh.Go

A fast and simple blog engine with GoInk framework in Golang.

Current version is **0.1.6-beta** on 2014.01.31

### Overview

Fxh.Go is a dynamic blog engine written in Golang. It's fast and very simple configs. Fxh.Go persists data into pieces of json files and support compress them as backup zip for next upgrade or installation.

Fxh.Go supports markdown contents as articles or pages, ajax comments and dynamic administration.

Fxh.Go contains two kinds of content as article and page. They can be customized as you want.

Documentation is writing.

### Requirement

* >= Golang 1.2

### Setup

Fxh.Go is written in Golang with support for Windows, Linux and Mac OSX.

The stable release can use `go get` to install:

    go get github.com/fuxiaohei/GoBlog

remember set `$GOPATH/bin` to global environment variables.

The newest sources is in branch develop, you can download and build in manual (not-recommended).

### Setup

If installed, `GoBlog` binary file is built in `$GOPATH/bin`.

make a new dir to install Fxh.Go:

    cd empty_dir
    Goblog

then it will unzip static files in `empty_dir` , initialize original data and start server at `localhost:9000`

#### Administration

visit `localhost:9000/login/` to enter administrator with username `admin` and password `admin`. You'd better change them after installed successfully.

### Suggestion and Contribution

create issues or pull requests here.

### Products

* [抛弃世俗之浮躁，留我钻研之刻苦](http://wuwen.org)
* [FuXiaoHei.Net](http://fuxiaohei.net)

### Thanks

gladly thank for [@Unknwon](https://github.com/Unknwon) on testing and [zip library](https://github.com/Unknwon/cae) support.

### License

The MIT License

