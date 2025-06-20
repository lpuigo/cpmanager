package persist

import (
	"fmt"
	"github.com/google/uuid"
	log "github.com/lpuig/cpmanager/log"
	"io"
	"os"
	"path/filepath"
	"sync"
	"time"
)

type Recorder interface {
	GetId() string
	SetId(id string)
	Dirty()
	Persist(path string) error
	Remove(path string) error
	GetFileName() string
	GetIdFromFileName(string) (string, error)
	Marshall(writer io.Writer) error
}

const (
	DefaultPersistDelay = 1 * time.Second
	ParallelPersister   = 10
)

type Persister struct {
	name      string
	directory string
	delay     time.Duration
	records   map[string]Recorder
	nextId    func(Recorder) string
	log       *log.Logger

	mut         sync.RWMutex
	persistDone *sync.Cond
	dirtyIds    []string

	persistTimer *time.Timer
}

// NewPersister creates a new persister with given name and storing its record in given dir directory
func NewPersister(name, dir string, logger *log.Logger) *Persister {
	p := &Persister{
		name:      name,
		directory: dir,
		delay:     DefaultPersistDelay,
		log:       logger,
		nextId: func(Recorder) string {
			return uuid.NewString()
		},
	}
	p.persistDone = sync.NewCond(&p.mut)
	p.Reinit()
	return p
}

func (p *Persister) NbRecords() int {
	return len(p.records)
}

// MutLock call Lock on reciever sync.RWMutex. Should not be used aside of others reciever's method, to prevent deadlock or panic
func (p *Persister) MutLock() {
	p.mut.Lock()
}

// MutUnLock call Unlock on reciever sync.RWMutex. Should not be used aside of others reciever's method, to prevent deadlock or panic
func (p *Persister) MutUnLock() {
	p.mut.Unlock()
}

// Reinit waits persister mechanism to finish (if any) and reset the persister (empty record and id counter reset to 0)
func (p *Persister) Reinit() {
	p.WaitPersistDone()
	p.mut.Lock()
	defer p.mut.Unlock()
	p.records = make(map[string]Recorder)
}

// SetPersistDelay sets the Pesistance Delay of the Persister
//
// if persistDelay is set to 0, dirty records will be synchronously persisted (writen to disk)
func (p *Persister) SetPersistDelay(persistDelay time.Duration) {
	p.delay = persistDelay
}

// SetNextIdFunc change the default nextId method (random uuid)
func (p *Persister) SetNextIdFunc(f func(Recorder) string) {
	p.nextId = f
}

// NoDelay suppresses receiver persist delay : any record marked as dirt will be persisted synchronously
func (p *Persister) NoDelay() {
	p.SetPersistDelay(0)
}

// CheckDirectory checks if Persister directory exists and create deleted dir if missing
func (p *Persister) CheckDirectory() error {
	fi, err := os.Stat(p.directory)
	if err != nil {
		return err
	}
	if !fi.IsDir() {
		return fmt.Errorf("not a proper directory: %s\n", p.directory)
	}
	dpath := filepath.Join(p.directory, "deleted")
	if _, err := os.Stat(dpath); os.IsNotExist(err) {
		return os.Mkdir(dpath, os.ModePerm)
	}
	return nil
}

// LoadDirectory loads all persisted Records. Records Id are set from file names
func (p *Persister) LoadDirectory(recordFactory func(string) (Recorder, error)) error {
	p.WaitPersistDone() // first wait for all ongoing operations to end
	p.mut.Lock()
	defer p.mut.Unlock()

	p.records = make(map[string]Recorder)

	files, err := p.GetFilesList("deleted")
	if err != nil {
		return fmt.Errorf("could not get files from %s persister: %v", p.name, err)
	}

	for _, file := range files {
		ar, err := recordFactory(file)
		if err != nil {
			return fmt.Errorf("could not instantiate %s from '%s': %v", p.name, filepath.Base(file), err)
		}
		bfile := filepath.Base(file)
		id, err := ar.GetIdFromFileName(bfile)
		if err != nil {
			return fmt.Errorf("could not get id from '%s': %v", bfile, err)
		}
		ar.SetId(id)
		err = p.load(ar)
		if err != nil {
			return fmt.Errorf("error while loading %s: %s", file, err.Error())
		}
	}
	return nil
}

// HasId returns true if persister contains a record with given id, false otherwise
func (p *Persister) HasId(id string) bool {
	if _, ok := p.records[id]; ok {
		return true
	}
	return false
}

// GetFilesList returns all the record files contained in persister directory (User class is responsible to Load the record)
func (p *Persister) GetFilesList(skipdir string) (list []string, err error) {
	err = filepath.Walk(p.directory, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			if info.Name() == skipdir {
				return filepath.SkipDir
			}
			return nil
		}
		list = append(list, path)
		return nil
	})
	if err != nil {
		return
	}
	return
}

// GetRecords returns a slice with all Records from receiver Persister
func (p *Persister) GetRecords() []Recorder {
	p.mut.RLock()
	defer p.mut.RUnlock()
	res := make([]Recorder, len(p.records))
	i := 0
	for _, recorder := range p.records {
		res[i] = recorder
		i++
	}
	return res
}

// GetById returns a recorder with given Id (or nil if Id not found)
func (p *Persister) GetById(id string) (Recorder, bool) {
	p.mut.RLock()
	defer p.mut.RUnlock()
	r, found := p.records[id]
	return r, found
}

// Add adds the given Record to the Persister, assigns it a new id, triggers Persit mechanism and returns its (new) id
func (p *Persister) Add(r Recorder) string {
	p.mut.Lock()
	defer p.mut.Unlock()

	id := p.nextId(r)
	r.SetId(id)
	p.records[id] = r
	p.markDirty(r)

	return id
}

// Update the given Record to the Persister and triggers Persit mechanism
func (p *Persister) Update(r Recorder) error {
	rId := r.GetId()
	if !p.HasId(rId) {
		return fmt.Errorf("record with id %s not found", rId)
	}
	p.mut.Lock()
	defer p.mut.Unlock()
	p.records[rId] = r
	p.markDirty(r)

	return nil
}

// Load adds the given Record to the Persister
func (p *Persister) Load(r Recorder) error {
	p.mut.Lock()
	defer p.mut.Unlock()

	return p.load(r)
}

// load checks given record's Id, and if not already known, adds it to the reciever
func (p *Persister) load(r Recorder) error {
	rId := r.GetId()
	if p.HasId(rId) {
		return fmt.Errorf("persister already contains a record with Id %s", rId)
	}
	p.records[rId] = r
	//if p.nextId <= rId {
	//	p.nextId = rId + 1
	//}
	return nil
}

// markDirty marks the given recorder as dirty and triggers the persistence mechanism
func (p *Persister) MarkDirty(r Recorder) {
	p.mut.Lock()
	defer p.mut.Unlock()
	p.markDirty(r)
}

func (p *Persister) markDirty(r Recorder) {
	if _, ok := p.records[r.GetId()]; !ok {
		return
	}
	r.Dirty()
	p.dirtyIds = append(p.dirtyIds, r.GetId())
	p.triggerPersist()
}

// Remove removes the given recorder from the persister (pertaining persisted file is deleted)
func (p *Persister) Remove(r Recorder) error {
	id := r.GetId()
	if _, ok := p.records[id]; !ok {
		return fmt.Errorf("persister does not contains given record with Id %s", id)
	}
	p.mut.Lock()
	defer p.mut.Unlock()
	go func(dr Recorder) {
		err := dr.Remove(p.directory)
		if err != nil {
			p.log.Error(fmt.Sprintf("error removing record id %s: %v\n", id, err))
		}
	}(r)
	delete(p.records, id)
	return nil
}

// PersistAll immediatly persist all contained recorder(persistance delay is ignored)
func (p *Persister) PersistAll() {
	p.mut.Lock()
	defer p.mut.Unlock()
	// desactivate persistMechanism if activated
	if p.persistTimer != nil {
		p.persistTimer.Stop()
		p.persistTimer = nil
		p.dirtyIds = []string{}
	}

	token := make(chan struct{}, ParallelPersister)
	defer close(token)
	for _, r := range p.records {
		token <- struct{}{}
		go func(pr Recorder) {
			err := r.Persist(p.directory)
			if err != nil {
				p.log.Error(fmt.Sprintf("error persisting record id %s: %v\n", r.GetId(), err))
			}
			_ = <-token
		}(r)
	}
	// wait for all persister completion
	for i := 0; i < ParallelPersister; i++ {
		token <- struct{}{}
	}
}

func (p *Persister) triggerPersist() {
	if p.delay == 0 {
		if p.persistTimer != nil {
			p.persistTimer.Stop()
			p.persistTimer = nil
		}
		p.persist()
		return
	}
	if p.persistTimer != nil {
		return
	}
	p.persistTimer = time.AfterFunc(p.delay, func() {
		p.mut.Lock()
		defer p.mut.Unlock()
		p.persistTimer = nil
		p.persist()
	})
}

func (p *Persister) persist() {
	token := make(chan struct{}, ParallelPersister)
	defer close(token)
	for _, id := range p.dirtyIds {
		r, found := p.records[id]
		if !found { // can happen if record was remove before persistence was triggered
			continue
		}
		token <- struct{}{}
		go func(pr Recorder) {
			err := pr.Persist(p.directory)
			if err != nil {
				p.log.Error(fmt.Sprintf("error persisting record ID %s: %v\n", pr.GetId(), err))
			}
			_ = <-token
		}(r)
	}
	// wait for all persister completion
	for i := 0; i < ParallelPersister; i++ {
		token <- struct{}{}
	}
	p.dirtyIds = []string{}
	p.persistDone.Broadcast()
}

// WaitPersistDone waits for current persisting mechanism to end and return (return instantly if no persist in progress)
func (p *Persister) WaitPersistDone() {
	if p.persistTimer == nil && len(p.dirtyIds) == 0 {
		return
	}
	p.persistDone.L.Lock()
	p.persistDone.Wait()
	p.persistDone.L.Unlock()
}

// GetName returns persister's name
func (p *Persister) GetDirectory() string {
	return p.directory
}

// GetName returns persister's name
func (p *Persister) GetName() string {
	return p.name
}
