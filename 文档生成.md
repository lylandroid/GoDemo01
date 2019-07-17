## 一.代码[文档生成]
0. 查看相关命令: go help doc 
1. cd 到具体目录
2. go doc

        package queue // import "."
        type Queue []int
3. 查看文档：go doc Queue //输出结果如下

        func (q *Queue) IsEmpty() bool
        func (q *Queue) Pop() int
        func (q *Queue) Push(value int)

4. 查看具体api调用：go doc IsEmpty //输出结果如下
    
        func (q *Queue) IsEmpty() bool
5. 查看系统api：go doc fmt.Println

## 二.godoc 使用[godoc help]
1. go web文档 ~~(在浏览器中访问查看)~~
 
        godoc -http :6060
        
        
