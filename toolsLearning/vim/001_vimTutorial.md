# vim turtorial

> 学习参考自 https://github.com/iggredible/

## Chapter 01

### Installation

### Exec

### Exit

### Saving A File

- 如果是新文件，需要给出文件名

```zsh
: w file.txx
```

### Helping

- 输入`:help {command}`

### Opening Files

### Arguments

- 情境
 - 打开文件后迅速执行一条命令，可以向 vim 传递一个 `+{cmd}` 选项
  1. 将文本中的 pancake 替换为 bagel，在终端中输入以下内容
   - `vim +%s /pancake/bagel/g hello.txt`
  2. 且命令可以叠加执行
   - `vim +%s /pancake/bagel/g +%s/bagel/egg/g +%s/egg/donut/g hello.text`
    - 替换流程： pancake -> bagel -> egg -> donut

### Opening Windows

使用 `o` `0` 选项，vim 打开后分别显示为水平或垂直分割窗口

```zsh
# 水平分割
vim -o2

# 垂直分割
vim -02

# 五个水平分割窗口，并使前两个显示 hello1.text hello2.text
vim -o2 hello1.text hello2.text
```

### Suspending

```zsh
C + z
:stop
:suspend

# 终端输入以下命令，从挂起状态中返回
fg
```

### Startig Vim the smart way

- 可以向 vim 命令传递不同的选项和标志
 1. 命令行命令（`+{cmd}` 或 `-c cmd`）
 2. 可以将 vim 命令和其他终端命令联合起来
  - eg: 将 ls 命令的输出重定向打 Vim 中编辑
   - `ls -l | vim -`


## Chapter 02

### Buffers

- buffers: 内存中的一块空间，可以在其中写入或编辑文本。每当 Vim 打开一个文件，文件数据就与一个 buffer 绑定

运行 `:buffers` 命令可以查看所有的buffers（也可以使用`:ls` 和 `files` 命令）

遍历 buffers 的方法：
 - `:bnext` 切换到下一个buffer
  - `:bprevious` 切换到上一个buffer
 - `:buffer` + filename（Tab 键补全文件名）
 - `:buffer`+`n`，n 是buffer 编号

删除 buffers：
 - `:bdelete`
 - `:bdelete n`
 - `:bdelete <filename>`

> buffers: 理解为多个平面，在某一时刻只有一个平面可见

### Exit

```zsh
# 关闭所有 Buffers
:qall

# 退出不保存
:qall!

# 保存并退出
:wqall
```

### Windows

> Windows 视为 buffer 的一个可编辑窗口(关闭窗口后，buffer仍继续存在)，一个buffer可以有多个编辑窗口（在其中一个窗口编辑时，同一buffer的其他窗口也在同步更新）

```zsh
# 垂直分割当前窗口，并在新窗口打开文件
:vsplit filename

# 水平分割当前窗口，并在新窗口打开文件
:split filename

# 新窗口打开文件
:new filename
```

### Tabs

>  Windows 的集合（布局等信息）

```zsh
:tabnew file.txt
:tabclose
:tabnext
:tabprevious
:tablast
:tabfirst

# 下一个标签页
gt
# 上一个标签页
gT
# 切换到第三个tab
3gt
```
### Using Buffers, Windows, Tabs the smart way

## Chapter 03

### Opening and editing

```zsh
# 打开文件，支持通配符，tab补全操作
:edit filename
```

### Find

### Path





























