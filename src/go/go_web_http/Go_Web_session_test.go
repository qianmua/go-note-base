package test

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"github.com/astaxie/beego/session"
	"net/http"
	"net/url"
	"sync"
	"testing"
	"time"
)

/**
 * @author
 * @date 2020/11/4 14:09
 * create by Goland
 * @version 1.0
 */ 
 
// session and cookie


// session manage
type Manager struct {
	cookieName string // private cookie name
	lock sync.Mutex // protects session
	provider Provider // base api
	maxLifeTime int64
}

//Provider
type Provider interface {
	SessionInit(sid string) (Session, error) // session Init
	SessionRead(sid string) (Session, error) // read session
	SessionDestroy(sid string) error	// 销毁
	SessionGC(maxLifeTime int64) //清理
}
// Session
type Session interface {
	Set(key, value interface{}) error // set session value
	Get(key interface{}) interface{}  // get session value
	Delete(key interface{}) error     // delete session value
	SessionID() string                // back current sessionID
}

// session map
var provides = make(map[string]Provider)

// Register
func Register(name string , provider Provider){
	if provider == nil {
		// not nil
		panic("session: Register provider is nil")
	}
	if _, dup := provides[name]; dup {
		// not nil
		panic("session: Register called twice for provider " + name)
	}
	provides[name] = provider
}
// gen session id
// 全局唯一的（GUID）
func (manager *Manager) sessionId()string {
	b := make([]byte , 32)
	if _ ,err := rand.Read(b) ; err != nil{
		return ""
	}
	return base64.URLEncoding.EncodeToString(b)
}

// 检测是否已经有某个Session与当前来访用户发生了关联，如果没有则创建之
func (manager *Manager) SessionStart(w http.ResponseWriter , r *http.Request) (session Session){
	manager.lock.Lock()
	defer manager.lock.Unlock()
	cookie, err := r.Cookie(manager.cookieName)
	if err != nil || cookie.Value == "" {
		sessionId := manager.sessionId()
		session, _ = manager.provider.SessionInit(sessionId)
		cookie := http.Cookie{
			Name:     manager.cookieName,
			Value:    url.QueryEscape(sessionId),
			Path:     "/",
			Expires:  time.Time{},
			MaxAge:   int(manager.maxLifeTime),
			HttpOnly: true,
		}
		http.SetCookie(w , &cookie)
	}else {
		sid, _ := url.QueryUnescape(cookie.Value)
		session, _ = manager.provider.SessionRead(sid)
	}
	return
}

/**
创建 session
 */
func NewManager(provideName, cookieName string, maxLifeTime int64) (*Manager, error) {
	provider, ok := provides[provideName]
	if !ok {
		return nil, fmt.Errorf("session: unknown provide %q (forgotten import?)", provideName)
	}
	return &Manager{provider: provider, cookieName: cookieName, maxLifeTime: maxLifeTime}, nil
}


// global management
var globalSessions *Manager
//然后在init函数中初始化
func init() {
	globalSessions, _ = NewManager("memory", "gosessionid", 3600)
}

// test login
func TestLogin(t *testing.T) {

}

func login2(w http.ResponseWriter , r *http.Request){
	SessionStart
}