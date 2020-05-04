package golang_orm

import (
	"database/sql"
	"golang_orm/dialect"
	"golang_orm/log"
	"golang_orm/session"
)

type Engine struct {
	db      *sql.DB
	dialect dialect.Dialect
}

func NewEngine(driver, source string) (e *Engine, err error) {
	db, err := sql.Open(driver, source)
	if err != nil {
		log.Error(err)
		return
	}
	if err = db.Ping(); err != nil {
		log.Error(err)
		return

	}
	dial, ok := dialect.GetDialect(driver)
	if !ok {
		log.Error("dialect %s Not found", driver)
		return
	}
	e = &Engine{db: db, dialect: dial}
	log.Info("Connect database success!")
	return
}
func (engine *Engine) Close() {
	if err := engine.db.Close(); err != nil {
		log.Error("Fail to close database")
	}
	log.Info("Close database success")
	return
}

func (engine *Engine) NewSession() *session.Session {
	return session.New(engine.db, engine.dialect)
}
