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

查询数据库连接数状态信息
```
MariaDB [gin_project]> show status like '%conn%';
+-----------------------------------------------+-------+
| Variable_name                                 | Value |
+-----------------------------------------------+-------+
| Aborted_connects                              | 0     |   <- 被中断的连接, 可能由于密码错误或者权限错误, 或者其他原因断开连接
| Connection_errors_accept                      | 0     |
| Connection_errors_internal                    | 0     |
| Connection_errors_max_connections             | 10204 |
| Connection_errors_peer_address                | 0     |
| Connection_errors_select                      | 0     |
| Connection_errors_tcpwrap                     | 0     |
| Connections                                   | 57418 |   <- 服务器启动至今, 曾经完成过的客户端连接次数
| Max_used_connections                          | 101   |   <- 服务器启动至今, 曾经出现过的最大并发连接
| Performance_schema_session_connect_attrs_lost | 0     |
| Slave_connections                             | 0     |
| Slaves_connected                              | 0     |
| Ssl_client_connects                           | 0     |   <- 通过 ssl 方式的客户端连接
| Ssl_connect_renegotiates                      | 0     |
| Ssl_finished_connects                         | 0     |
| Threads_connected                             | 1     |   <- 当前正在与 mariadb 交互的客户端连接数
| wsrep_connected                               | OFF   |
+-----------------------------------------------+-------+
17 rows in set (0.03 sec)
```

查询数据库连接数配置信息
```
MariaDB [flask_project]> show variables like '%conn%';
+-----------------------------------------------+-----------------+
| Variable_name                                 | Value           |
+-----------------------------------------------+-----------------+
| character_set_connection                      | utf8            |
| collation_connection                          | utf8_general_ci |
| connect_timeout                               | 5               |
| default_master_connection                     |                 |
| extra_max_connections                         | 1               |
| init_connect                                  |                 |
| max_connect_errors                            | 100             |
| max_connections                               | 100             |
| max_user_connections                          | 0               |
| performance_schema_session_connect_attrs_size | -1              |
+-----------------------------------------------+-----------------+
10 rows in set (0.00 sec)
```
