## What

- 与 [Java 版本 ](https://github.com/ItsFunny/cell) 框架类似, 底层设计是一模一样的,所以具体介绍去看 Java版的吧

## How To Use

- 导入依赖

    - > ```
    > import "github.com/itsfunny/go-cell/application"
    > ```

- 添加指定的module(精力有限,所以只有http和swagger)

    - > ```
    > app := application.New(context.Background(),
    >    http.HttpModule,
    >    swagger.SwaggerModule,
    > )
    > ```

- 启动即可

    - > ```
    > app.Run(os.Args)
    > ```

- 完整的示例:
    - https://github.com/ItsFunny/go-cell/tree/dev/demo/demo1