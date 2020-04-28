### 使用golang+sql-driver写一个go orm
1. 使用sql-lite
安装 sudo apt install sqlite3
2. 链接数据库,如果不存在则创建
> sqlite3 golang.db
3. 创建表和字段
> sqlite3 create table user(name text, age integer);
> sqlite3 .head on //打开列名的开关
4. 插入数据
> sqlite3 insert into user(name,age) values ("lili",18),("lihua",19);
5. 查询数据
> sqlite3 select * from user where age > 18;
统计个数: select count(*) from user;