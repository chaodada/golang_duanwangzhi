package conf

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/viper" // 配置文件模块
)

// 定义Conf结构体
type Conf struct {
	Redis Redis // Redis结构体类型
}

// 定义Redis 结构体
type Redis struct {
	Addr string // redis 连接地址
	Pwd  string // redis 密码
	Db   int    // redis 选择数据库
}

var C *Conf // 定义一个变量 类型为Conf的指针

// conf 模块init方法
func init() {
	log.Println("conf包 conf文件 init方法")
	// 获取运行模式
	env := getMode()
	log.Print("运行模式为：", env)

	// 获取项目目录
	realPath, _ := filepath.Abs("./")
	log.Print("项目目录为：", realPath)

	// 配置文件名称
	viper.SetConfigName("app")
	log.Print("配置文件名称：", "app")

	// 如果配置文件没有扩展名 则需要指定扩展名
	viper.SetConfigType("toml")
	log.Print("配置文件类型：", "toml")

	// 在以下位置寻找配置文件
	viper.AddConfigPath(realPath + "/conf/" + env)
	log.Print("配置文件路径：", realPath+"/conf/"+env)

	// 查找并读取配置文件
	err := viper.ReadInConfig()

	// 如果读取配置文件错误错误
	if err != nil {
		panic(fmt.Errorf("读取配置文件错误: %s \n", err))
	}

	C = &Conf{} // 初始化Conf 并将指针给C

	err = viper.Unmarshal(C) //相当于把配置转成指定格式 给Conf 里边指定的元素    ***   比如将读取的配置文件中的rides配置 传递给Conf 的Redis 属性
	if err != nil {
		panic(err)
	}
}

// 检查环境是否开发模式
func getMode() string {
	env := os.Getenv("RUN_MODE") // Getenv检索并返回名为key的环境变量的值。如果不存在该环境变量会返回空字符串。
	if env == "" {               // 环境变量不存在
		env = "dev" // 赋值为dev
	}
	return env
}
