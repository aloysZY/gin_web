package auth

import (
	"time"

	"gin_web/global"

	"github.com/golang-jwt/jwt/v4"
)

// Claims 这是自定义 JWT 的字段信息
type Claims struct {
	// 不要放机密信息，这不是放数据库的s
	UserId               uint64                   `json:"auth_id,string"` // 生成 token 的时候传入这个字段，认证后解析，行下文记录使用这个字段
	jwt.RegisteredClaims `json:"standard_claims"` // 内嵌标准的声明
}

// GetJWTSecret 获取配置文件秘钥
func GetJWTSecret() []byte {
	return []byte(global.AppSetting.JWT.Secret)
}

// CreateToken 获取token，最好还是密码验证后进行 token
func CreateToken(userId uint64) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(global.AppSetting.JWT.Expire)
	// 创建一个我们自己的声明
	claims := Claims{
		// util.EncodeMD5(authId), 额外做了一层加密，因为获取 token，这段信息是可以直接被解密的，不需要知道盐，账号密码存进去就失败了
		// UserId: util.EncodeMD5(userId), // 这里暂时不加密了，用户的id，在认证后进行的操作记录直接用这个 ID
		// AppKey:    util.EncodeMD5(appKey),
		// AppSecret: util.EncodeMD5(appSecret),
		UserId: userId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expireTime),
			Issuer:    global.AppSetting.JWT.Issuer,
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims) // SigningMethodHS256 加密的算法，还有 384 和 512
	token, err := tokenClaims.SignedString(GetJWTSecret())           // 传入秘钥，获取签名字符串
	return token, err
}

// ParseToken 解析 token
func ParseToken(token string) (*Claims, error) {
	// 解析token
	// 如果是自定义Claim结构体则需要使用 ParseWithClaims 方法
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return GetJWTSecret(), nil
	})
	// 如果 token 过去，这里就返回错误了
	if err != nil {
		return nil, err
	}
	if tokenClaims != nil {
		// 对token对象中的Claim进行类型断言
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid { // tokenClaims.Valid 校验 token
			return claims, nil
		}
	}
	return nil, err
}
