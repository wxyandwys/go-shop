package config

import (
    "bytes"
    "github.com/dchest/captcha"
    "github.com/gin-contrib/sessions"
    "github.com/gin-contrib/sessions/cookie"
    "github.com/gin-gonic/gin"
    "net/http"
    "time"
    "github.com/garyburd/redigo/redis"
)

// 中间件，处理session
func Session(keyPairs string) gin.HandlerFunc {
    store := SessionConfig()
    //store := cookie.NewStore([]byte("secret11111"))
    return sessions.Sessions(keyPairs, store)
}
func SessionConfig() sessions.Store {
    sessionMaxAge := 3600
    sessionSecret := "topgoer"
    var store sessions.Store
    store = cookie.NewStore([]byte(sessionSecret))
    store.Options(sessions.Options{
        MaxAge: sessionMaxAge, //seconds
        Path:   "/",
    })
    return store
}

func Captcha(c *gin.Context, length ...int) {
    l := captcha.DefaultLen
    w, h := 107, 36
    if len(length) == 1 {
        l = length[0]
    }
    if len(length) == 2 {
        w = length[1]
    }
    if len(length) == 3 {
        h = length[2]
    }
    captchaId := captcha.NewLen(l)
    //session := sessions.Default(c)
    //session.Set("captcha", captchaId)
    // redis测试
    conn := RedisDefaultPool.Get()
    conn.Do("setex", "captcha", 200, captchaId)

   // _ = session.Save()
    _ = Serve(c.Writer, c.Request, captchaId, ".png", "zh", false, w, h)
}
func CaptchaVerify(c *gin.Context, code string) bool {
    //session := sessions.Default(c)
    conn := RedisDefaultPool.Get()
    captchaId ,_ := redis.Bytes(conn.Do("get", "captcha"))
    s := string(captchaId)
    if  s != "" {
        conn.Do("del", "captcha")
        /*
        session.Delete("captcha")
        _ = session.Save()
        */
        
        if captcha.VerifyString(s, code) {
            return true
        } else {
            return false
        }
    } else {
        return false
    }
    /*
    if captchaId := session.Get("captcha"); captchaId != nil {
        session.Delete("captcha")
        _ = session.Save()
        if captcha.VerifyString(captchaId.(string), code) {
            return true
        } else {
            return false
        }
    } else {
        return false
    }
    */
}
func Serve(w http.ResponseWriter, r *http.Request, id, ext, lang string, download bool, width, height int) error {
    w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
    w.Header().Set("Pragma", "no-cache")
    w.Header().Set("Expires", "0")

    var content bytes.Buffer
    switch ext {
    case ".png":
        w.Header().Set("Content-Type", "image/png")
        _ = captcha.WriteImage(&content, id, width, height)
    case ".wav":
        w.Header().Set("Content-Type", "audio/x-wav")
        _ = captcha.WriteAudio(&content, id, lang)
    default:
        return captcha.ErrNotFound
    }

    if download {
        w.Header().Set("Content-Type", "application/octet-stream")
    }
    http.ServeContent(w, r, id+ext, time.Time{}, bytes.NewReader(content.Bytes()))
    return nil
}

/*import (
    "encoding/json"
    captcha "github.com/mojocn/base64Captcha"
    "net/http"
    
)

var store = captcha.DefaultMemStore

func NewDriver() *captcha.DriverString {
    driver := new(captcha.DriverString)
    driver.Height = 44
    driver.Width = 120
    driver.NoiseCount = 5
    driver.ShowLineOptions = captcha.OptionShowSineLine | captcha.OptionShowSlimeLine | captcha.OptionShowHollowLine
    driver.Length = 6
    driver.Source = "1234567890qwertyuipkjhgfdsazxcvbnm"
    driver.Fonts = []string{"wqy-microhei.ttc"}
    return driver
}

// 生成图形验证码
func GenerateCaptchaHandler(w http.ResponseWriter, r *http.Request) {
    var driver = NewDriver().ConvertFonts()
    c := captcha.NewCaptcha(driver, store)
    _, content, answer := c.Driver.GenerateIdQuestionAnswer()
    id := "captcha:yufei"
    item, _ := c.Driver.DrawCaptcha(content)
    c.Store.Set(id, answer)
    item.WriteTo(w)
}

// 验证
func CaptchaVerifyHandle(w http.ResponseWriter, r *http.Request) {

    id := "captcha:yufei"
    code := r.FormValue("code")
    body := map[string]interface{}{"code": 1000, "msg": "failed"}
    if store.Verify(id, code, true) {
        body = map[string]interface{}{"code": 1001, "msg": "ok"}
    }
    w.Header().Set("Content-Type", "application/json; charset=utf-8")
    json.NewEncoder(w).Encode(body)
}
*/