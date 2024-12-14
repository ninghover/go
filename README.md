src目录下就是全部的源码
每个文件夹就是一个工程

consul 简易上手指南：https://www.cnblogs.com/edisonfish/p/17216756.html


初始化一个git远程仓库

```bash
1.在github.com创建一个库,本例中就是go
2.在本地执行: 
    git init
    git add .
    git commit -m "first commit"
    git remote add origin git@github.com:ninghover/go.git       远程库和本地库进行绑定
    git push origin master
```

本地单分支上的多个commit怎么合并为一个分支
```
首先，如果本地还有未暂存，未提交的，先暂存，先提交
git log 查看提交记录
git rebase -i HEAD~n // n为要合并的次数，比如:4 
然后进入一个交互页面，除了第一个pick，其他的都改成s，
然后保存，重新编辑提交信息
```