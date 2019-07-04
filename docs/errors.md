# errors

https://github.com/pkg/errors

```
go get -u github.com/pkg/errors
```

最佳实践参考 [https://studygolang.com/articles/11753](https://studygolang.com/articles/11753)


错误处理的正确姿势
- 失败的原因只有一个时，不使用error，使用bool替代
- 没有失败时，不使用error
- error应放在返回值类型列表的最后，bool作为返回值类型时也一样
- 错误值统一定义，而不是跟着感觉走
- 错误逐层传递时，层层都加日志
- 错误处理使用defer
- 当尝试几次可以避免失败时，不要立即返回错误
- 当上层函数不关心错误时，建议不返回error
- 当发生错误时，不忽略有用的返回值

异常处理的正确姿势
- 在程序开发阶段，坚持速错
- 在程序部署后，应恢复异常避免程序终止
- 对于不应该出现的分支，使用异常处理
- 针对入参不应该有问题的函数，使用panic设计


panic异常处理机制不会自动将错误信息传递给error，需要显示传递


```
func deferDemo() error {
    err := createResource1()
    if err != nil {
        return ERR_CREATE_RESOURCE1_FAILED
    }

    err = createResource2()
    if err != nil {
        destroyResource1()
        return ERR_CREATE_RESOURCE2_FAILED
    }

    err = createResource3()
    if err != nil {
        destroyResource2()
        destroyResource1()
        return ERR_CREATE_RESOURCE3_FAILED
    }

    err = createResource4()
    if err != nil {
        destroyResource3()
        destroyResource2()
        destroyResource1()
        return ERR_CREATE_RESOURCE4_FAILED
    }
    return nil
}
```
重构为：
```
func deferDemo() error {
    err := createResource1()
    if err != nil {
        return ERR_CREATE_RESOURCE1_FAILED
    }
    defer func() {
        if err != nil {
            destroyResource1()
        }
    }()

    err = createResource2()
    if err != nil {
        return ERR_CREATE_RESOURCE2_FAILED
    }
    defer func() {
        if err != nil {
            destroyResource2()
        }
    }()

    err = createResource3()
    if err != nil {
        return ERR_CREATE_RESOURCE3_FAILED
    }
    defer func() {
        if err != nil {
            destroyResource3()
        }
    }()

    err = createResource4()
    if err != nil {
        return ERR_CREATE_RESOURCE4_FAILED
    }
    return nil
}
```
