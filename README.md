## 1、项目介绍
> 该项目只是对于go语言进行网络编程的起手项目，所以着重于服务器的编写
* [gin的学习记录](https://github.com/547173318/redo-golang/tree/main/doc/3-Golang%E4%B9%8B%E7%BD%91%E7%BB%9C%E7%BC%96%E7%A8%8B/1-gin%E5%AD%A6%E4%B9%A0%E8%AE%B0%E5%BD%95)
* 服务器开发使用到的技术栈
    * 使用gin搭建web相应服务器
    * 使用gorm调用mysql服务
* 完成的功能（json交互）
    * 待办事项的添加
    * 待办事项的查询
    * 待办事项的完成情况更改
    * 待办事项的删除
* 学习的感悟
    * go
        * go相比于java，多了指针，处理起来不方便但是灵活，如gorm进行绑定model是，必须依靠与指针
        * 错误处理有些冗余，有许多的err
        * 语言层面的强制规定（如不使用到的变量不可存在，除非使用'_'接受）使得代码更加统一简洁
    * gorm
        * 指针声明后，如果想要附上实体，应该new出空间
            > 如DB.Find(todo),todo是指针的话，提前new
        * gorm的使用，可能自己还不熟悉，感觉开发效率不如手写sql
            > 对于复杂的sql，gorm可能比较有效。不过sql的语法经常不写就会忘记了，有时太依赖工具，也不太好
    * gin
        * gin呢，老套路了，url->controller->logic->model->db
## 2、使用说明
* 服务器（port:9000）
    * 数据库：mysql
    * 启动项目
    ```
    go run main.go
    ```
* vue前端（port:8080）
    * 启动项目
    ```
    npm install
    npm run serve
    ```


