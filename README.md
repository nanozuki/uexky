# xboard (temporarily)

一个新的讨论版程序

# 功能需求

## 用户侧需求
1. 方便的查看回复的方式：参考 4chanX 插件，借鉴并改进
2. 不使用分区和分类，而是主tag(有且仅一个) + 多个子tag的形式。主tag管理员添加，子tag用户自行添加，事后审核。可以合并，迁移
3. 可以自行订阅tag，然后合并在一起显示，（支持多个订阅组）
4. 外链贴图，支持MD5检索
5. 支持匿名回复，一个贴内匿名ID一致，不同贴内不同（可以有多个ID和匿名ID）
6. 一个tag可以有属于自己的专页/wiki

## 管理侧功能：

1. 删帖，删回复，禁言
2. Ban帐号 Ban IP
3. 修改Post的tag
4. 管理子tag
5. 用户提议合并，删除，迁移，重命名tag
6. 置顶，置底post
7. 在别人的帖子内容上添加评注

## 技术需求

1. 响应式设计
2. 实现 PWA
3. 浏览器通知
4. 尽量减少用户刷新界面