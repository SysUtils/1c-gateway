package odata

import "sync"

type Splitter struct {
	watcher   *Watcher
	m         sync.RWMutex
	childrens []*Watcher
}

func NewSplitter(watcher *Watcher) *Splitter {
	return &Splitter{watcher: watcher}
}

func (s *Splitter) NewWatcher() *Watcher {
	newWatcher := s.watcher.InternalCopy()
	s.childrens = append(s.childrens, newWatcher)
	return newWatcher
}

func (s *Splitter) emitDelete(e interface{}) {
	s.m.RLock()
	defer s.m.RUnlock()
	for _, w := range s.childrens {
		w.Deleted <- e
	}
}

func (s *Splitter) emitUpdate(e interface{}) {
	s.m.RLock()
	defer s.m.RUnlock()
	for _, w := range s.childrens {
		w.Updated <- e
	}
}

func (s *Splitter) emitCreate(e interface{}) {
	s.m.RLock()
	defer s.m.RUnlock()
	for _, w := range s.childrens {
		w.Created <- e
	}
}

func (s *Splitter) Start() {
	go s.loop()
}

func (s *Splitter) loop() {
	for {
		select {
		case e := <-s.watcher.Created:
			s.emitCreate(e)
		case e := <-s.watcher.Updated:
			s.emitUpdate(e)
		case e := <-s.watcher.Deleted:
			s.emitDelete(e)
		}
	}
}
