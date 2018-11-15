# MariaDB

查看当前数据库连接
```
MariaDB [gin_project]> show processlist;
+-------+------+------------------+-------------+---------+------+-------+------------------+----------+
| Id    | User | Host             | db          | Command | Time | State | Info             | Progress |
+-------+------+------------------+-------------+---------+------+-------+------------------+----------+
| 57174 | root | 172.17.0.6:37864 | gin_project | Query   |    0 | init  | show processlist |    0.000 |
+-------+------+------------------+-------------+---------+------+-------+------------------+----------+
1 row in set (0.00 sec)
```

查看数据库最大连接数
```
MariaDB [gin_project]> show variables like 'max_connections';
+-----------------+-------+
| Variable_name   | Value |
+-----------------+-------+
| max_connections | 100   |
+-----------------+-------+
1 row in set (0.00 sec)
```

查看数据库连接等待超时时间(global为准)
```
MariaDB [gin_project]> show global variables like 'wait_timeout';
+---------------+-------+
| Variable_name | Value |
+---------------+-------+
| wait_timeout  | 600   |
+---------------+-------+
1 row in set (0.00 sec)
MariaDB [gin_project]> show variables like 'wait_timeout';
+---------------+-------+
| Variable_name | Value |
+---------------+-------+
| wait_timeout  | 28800 |
+---------------+-------+
1 row in set (0.00 sec)
```
