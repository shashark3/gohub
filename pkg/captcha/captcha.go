package captcha

import (
	"gohub/pkg/app"
	"gohub/pkg/config"
	"gohub/pkg/redis"
	"sync"

	"github.com/mojocn/base64Captcha"
)

type Captcha struct {
	Base64Captcha *base64Captcha.Captcha
}

var once sync.Once

var internalCaptcha *Captcha

func NewCaptcha() *Captcha {
	once.Do(func() {
		//初始化Captcha
		internalCaptcha = &Captcha{}

		//使用全局Redis对象，配置Key前缀
		store := RedisStore{
			RedisClient: redis.Redis,
			keyPrefix:   config.GetString("app.name") + ":captcha:",
		}

		//配置base64captcha信息  driver使用captcha包自带的Driverdigit
		driver := base64Captcha.NewDriverDigit(
			config.GetInt("captcha.height"),      // 宽
			config.GetInt("captcha.width"),       // 高
			config.GetInt("captcha.length"),      // 长度
			config.GetFloat64("captcha.maxskew"), // 数字的最大倾斜角度
			config.GetInt("captcha.dotcount"),    // 图片背景里的混淆点数量
		)

		//实例化base64captcha 并赋值给自定义的internalCaptcha
		internalCaptcha.Base64Captcha = base64Captcha.NewCaptcha(driver, &store)
	})
	return internalCaptcha
}

// GenerateCaptcha 生成图片验证码
func (c *Captcha) GenerateCaptcha() (id, b64s string, err error) {
	return c.Base64Captcha.Generate()
}

// VerifyCaptcha 验证是否正确
func (c *Captcha) VerifyCaptcha(id string, answer string) (match bool) {

	//方便本地和Api自动测试
	// 方便本地和 API 自动测试
	if !app.IsProduction() && id == config.GetString("captcha.testing_key") {
		return true
	}
	// 第三个参数是验证后是否删除，我们选择 false
	// 这样方便用户多次提交，防止表单提交错误需要多次输入图片验证码
	return c.Base64Captcha.Verify(id, answer, false)
}
