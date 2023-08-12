## 说明

本项目主要提供一个基于 go-zero 的基本开发模板，然后快速开发，有两种思路

### 1.从 go-zero 从头构建


``` bash
go install github.com/zeromicro/go-zero/tools/goctl@latest
goctl env check --install --verbose --force
goctl api new project
cd project
go mod tidy
go run ./
```
即可获取一个最基础的项目，然后从本项目中获取你需要的相关包

### 2.基于本项目开发

``` bash
git clone https://github.com/xemxx/go-zero-template.git
cd go-zero-template
rm -rf .git
git init
go mod tidy
```
这就完成了项目初始化，再根据需要，重新编写api文件，根据上面的教程步骤实现开发，最好将 go.mod 文件的包名改为自己的项目地址


## 基本本项目的开发步骤

> 本项目基础采用 go-zero 进行开发，其他组件按需引用，主要利用到了api生成功能和db结构生成功能

### 目录结构

- /api
    - 存放实际api定义文件，根据go-zero 的现状，需要把所有api定义写在同一个文件
- /model
    - 存放数据库模型，主要是根据 go-zero 生成数据库表的定义，快速开发
- /internal
    - 存放内部库，按照一般的规则，整个项目是有可能被引用的，但是internal目录下的包是不能被其他项目引用的
    - 内部的目录结构由 go-zero 自动生成，暂时不变更，没有出现太大的逻辑问题
    - 特别说明
        - /handler 这个目录主要是包装入参和错误码，自动生成，不需要关注
        - /logic 这个目录主要是实现代码逻辑，是主要编写代码的地方，在go里面一般是数据库逻辑和业务逻辑不会有太明显的分层，如果出现逻辑复杂通常在同一个文件里写多个函数，来实现控制函数复杂度
        - /types 这个目录比较特殊，按理说每个接口的定义会区分文件，这里生成的就不管了
        - /config 这个目录的配置文件是有默认值
- /db
    - 存放我们的数据库表定义
- /pkg
    - 通常存放公共组件，可以复用的
    - /mysql
        - mysql链接池，目前非常简陋
- /etc
    - 这个表存放了配置文件的示例

### 基本开发步骤

这个项目基础主要用于api项目，所以主要如下

1、定义api

根据语法 https://go-zero.dev/docs/tasks/dsl/api 在/api 文件夹下的文件中编写你的api接口

然后通过命令

> goctl api go --api ./api/foo.api  --dir .

生成自动化的代码，通常执行路径我们的项目根目录，其他自行替换路径

2、生成需要的db结构

根据实际情况，如果已经编写在db，最好是直连db生成，命令如下：

> goctl model mysql datasource --url "root:123456@tcp(127.0.0.1:13306)/markman"   -t user -t note -d internal/model

说明：

1. url 是可以替换成实际的dsn链接
2. -t 参数表示你需要的数据表，可以多个
3. -d 代表了生成文件的输出目录

go里面通常会平铺代码，然后用不同的文件名称来代表模块，所以这个model生成的代码我们通常放置于一起

其他除非你手写sql文件，则可以利用文件生成

> goctl model mysql ddl --src db/user.sql --dir internal/model

3、保存数据库sql

方便的话，将数据库的建表 sql 作为版本文件存放于 db 目录中，这样方便沟通是否修改

按理说存在一个db迁移工具，方便做版本切换，但是实际生产中的db修改都是很少的，通常会单独编写一个升级脚本，所以暂时不做太过于集成化的东西。

4、修改启动配置

如果我们需要新增配置，比如指定什么什么文件，端口什么的，可以修改 config 目录下的配置结构体，然后修改 etc 目录下的示例文件

通常开发配置可以在本地新建一个配置，比如在新增一个配置文件是 etc/foo.yaml.dev

然后启动项目就可以是

> go run ./foo.go -f etc/foo.yaml.dev

这样就可以使用我们本地的配置

5、错误返回

参考 https://go-zero.dev/docs/tutorials/http/server/response/ext

引入一个新的errors包，然后在types目录中添加我们的Errors定义文件，每次直接引入即可

6、逻辑编写

通常在生成api和db之后就能编写逻辑了，db操作参考 https://go-zero.dev/docs/tutorials/mysql/curd

整个流程可以参考 internal/logic/g1/loginlogic.go 的 Login 方法以及 internal/handler/g1/loginhandler.go


