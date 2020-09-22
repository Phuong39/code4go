# db2struct [![Build Status](https://travis-ci.org/Shelnutt2/db2struct.svg?branch=master)](https://travis-ci.org/Shelnutt2/db2struct) [![Coverage Status](https://coveralls.io/repos/github/Shelnutt2/db2struct/badge.svg?branch=1-add-coveralls-support)](https://coveralls.io/github/Shelnutt2/db2struct?branch=1-add-coveralls-support) [![GoDoc](https://godoc.org/github.com/Shelnutt2/db2struct?status.svg)](https://godoc.org/github.com/Shelnutt2/db2struct)

对db2struct的二次封装：
1. 支持全表生成，db2struct只支持单表；
2. 自动由表名推断实体类名；
3. 实体字段顺序与数据库字段顺序保持一致；
