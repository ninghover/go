src目录下就是全部的源码
每个文件夹就是一个工程

consul 简易上手指南：https://www.cnblogs.com/edisonfish/p/17216756.html


初始化一个git远程仓库

1.在github.com创建一个库,本例中就是go
2.在本地执行: 
    git init
    git add .
    git commit -m "first commit"
    git remote add origin git@github.com:ninghover/go.git       远程库和本地库进行绑定
    git push origin master
