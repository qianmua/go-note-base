package test

import (
	"fmt"
	"sync"
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