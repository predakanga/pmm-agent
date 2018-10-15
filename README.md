# pmm-agent

[![Build Status](https://travis-ci.com/percona/pmm-agent.svg?branch=master)](https://travis-ci.com/percona/pmm-agent)
[![codecov](https://codecov.io/gh/percona/pmm-agent/branch/master/graph/badge.svg)](https://codecov.io/gh/percona/pmm-agent)
[![Go Report Card](https://goreportcard.com/badge/github.com/percona/pmm-agent)](https://goreportcard.com/report/github.com/percona/pmm-agent)

```terminal
$ make serve
go install -v ./...
go install -v ./vendor/github.com/percona/mysqld_exporter
pmm-agent --config=testdata/.pmm-agent.yml serve
2018/10/15 23:04:29 Listen on: 127.0.0.1:7771

```

```terminal
$ make demo
go install -v ./...
go install -v ./vendor/github.com/percona/mysqld_exporter
cp testdata/.pmm-agent.yml ~/.pmm-agent.yml

// Initial list should be empty.	
pmm-agent list

// Add new program.
pmm-agent add mysql-1 --env DATA_SOURCE_NAME=root@/ -- mysqld_exporter --collect.all

// List now should contain new program.
pmm-agent list
Name     Program          Arg              Env                        Running  PID    Err
mysql-1  mysqld_exporter  [--collect.all]  [DATA_SOURCE_NAME=root@/]  true     69519  

// Stop program.
pmm-agent stop mysql-1

// List now should contain stopped program.
pmm-agent list
Name     Program          Arg              Env                        Running  PID  Err
mysql-1  mysqld_exporter  [--collect.all]  [DATA_SOURCE_NAME=root@/]  false    0    signal: killed

// Start program.
pmm-agent start mysql-1

// List now should contain started program.
pmm-agent list
Name     Program          Arg              Env                        Running  PID    Err
mysql-1  mysqld_exporter  [--collect.all]  [DATA_SOURCE_NAME=root@/]  true     69552  

// Add another new program.
pmm-agent add mysql-2 --env DATA_SOURCE_NAME=root@/ -- mysqld_exporter --collect.all --web.listen-address=:9204

// List now should contain started programs.
pmm-agent list
Name     Program          Arg                                         Env                        Running  PID    Err
mysql-1  mysqld_exporter  [--collect.all]                             [DATA_SOURCE_NAME=root@/]  true     69552  
mysql-2  mysqld_exporter  [--collect.all --web.listen-address=:9204]  [DATA_SOURCE_NAME=root@/]  true     69569  

// Stop programs.
pmm-agent stop mysql-1 mysql-2


// List now should contain stopped programs.
pmm-agent list
Name     Program          Arg                                         Env                        Running  PID  Err
mysql-1  mysqld_exporter  [--collect.all]                             [DATA_SOURCE_NAME=root@/]  false    0    signal: killed
mysql-2  mysqld_exporter  [--collect.all --web.listen-address=:9204]  [DATA_SOURCE_NAME=root@/]  false    0    signal: killed

// Start programs.
pmm-agent start mysql-1 mysql-2

// List now should contain started programs again.
pmm-agent list
Name     Program          Arg                                         Env                        Running  PID    Err
mysql-1  mysqld_exporter  [--collect.all]                             [DATA_SOURCE_NAME=root@/]  true     69602  
mysql-2  mysqld_exporter  [--collect.all --web.listen-address=:9204]  [DATA_SOURCE_NAME=root@/]  true     69606  

// Stop all programs.
pmm-agent stop

// List now should contain stopped programs.
pmm-agent list
Name     Program          Arg                                         Env                        Running  PID  Err
mysql-1  mysqld_exporter  [--collect.all]                             [DATA_SOURCE_NAME=root@/]  false    0    signal: killed
mysql-2  mysqld_exporter  [--collect.all --web.listen-address=:9204]  [DATA_SOURCE_NAME=root@/]  false    0    signal: killed

// Start all programs.
pmm-agent start

// List now should contain started programs again.
pmm-agent list
Name     Program          Arg                                         Env                        Running  PID    Err
mysql-1  mysqld_exporter  [--collect.all]                             [DATA_SOURCE_NAME=root@/]  true     69639  
mysql-2  mysqld_exporter  [--collect.all --web.listen-address=:9204]  [DATA_SOURCE_NAME=root@/]  true     69640  

// Remove program.
pmm-agent remove mysql-1

// List should contain just one program.
pmm-agent list
Name     Program          Arg                                         Env                        Running  PID    Err
mysql-2  mysqld_exporter  [--collect.all --web.listen-address=:9204]  [DATA_SOURCE_NAME=root@/]  true     69640  

// Remove all programs.
pmm-agent remove

// List should be empty again.
pmm-agent list
```