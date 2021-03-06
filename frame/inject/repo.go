package inject

import "fmt"

//IStartRepo 定义IStartRepo接口
type IStartRepo interface {
	Speak(message string) string
}

//StartRepo 注入数据库
type StartRepo struct {
	//Source IDb `inject:""`
}

//Speak 实现Speak方法
func (s *StartRepo) Speak(message string) string {
	//使用注入的IDb访问数据库
	//s.Source.DB().Where("name = ?", "jinzhu").First(&user)
	return fmt.Sprintf("[Repository] speak: %s", message)
}
