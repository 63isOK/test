# test

test some idea

- json测试
  - 对固定格式的json的解析和生成(使用encoding/json)
  - 对动态格式或未知格式的json的解析(分别是用json包和gabs包)
- pkg/errors测试
  - 测试结果：对error添加了一个string信息和一个栈信息
  - 和标准库errors相比，多了agiel栈信息，通过Format()来支持fmt.Printf
- 获取本机ip,就像测试例子一样，只能获取本机ip，没有公网ip
