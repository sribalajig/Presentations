package mongo

import (
	"fmt"
	mgo "gopkg.in/mgo.v2"
	"log"
)

type SessionFactory struct {
	hostName string
	session  *mgo.Session
}

func NewSessionFactory(hostName string) SessionFactory {
	session, err := mgo.Dial(hostName)

	if err != nil {
		log.Println(fmt.Sprintf("Could not establish connection with mongo : %s", err))
	}

	return SessionFactory{
		hostName: hostName,
		session:  session,
	}
}

func (sessionFactory SessionFactory) Get() *mgo.Session {
	return sessionFactory.session.Copy()
}
