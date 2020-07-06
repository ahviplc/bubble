package conf

// MYSQL数据库连接配置
const Mysql_Dialect = "mysql"
const Mysql_User = "root"
const Mysql_Password = "root"
const Mysql_Database = "bubble"
const Mysql_Ip = "192.168.0.10" //const Mysql_Ip = "127.0.0.1"
const Mysql_Port = "3306"
const Mysql_Link_Info = Mysql_User + ":" + Mysql_Password + "@(" + Mysql_Ip + ":" + Mysql_Port + ")/" + Mysql_Database + "?charset=utf8&parseTime=True&loc=Local"
