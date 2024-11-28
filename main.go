package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sync"

	"github.com/jcelliott/lumber"
)

const Version = "1.0.0"

type (
	Logger interface {
		Fatal(string, ...interface{})
		Error(string, ...interface{})
		Warn(string, ...interface{})
		Info(string, ...interface{})
		Debug(string, ...interface{})
		Trace(string, ...interface{})
	}
	Driver struct {
		mutex   sync.Mutex
		mutexes map[string]*sync.Mutex
		dir     string
		log     Logger
	}
)

type Options struct {
	Logger
}

func New(dir string, options *Options) (*Driver, error) {
	dir = filepath.Clean(dir)
	opts := Options{}
	if options != nil {
		opts = *options
	}
	if opts.Logger == nil {
		opts.Logger = lumber.NewConsoleLogger((lumber.INFO))
	}
	driver := Driver{
		dir:     dir,
		mutexes: make(map[string]*sync.Mutex),
		log:     opts.Logger,
	}
	if _, err := os.Stat(dir); err == nil {
		opts.Logger.Debug("Using '%s' (database alredy exists)\n", dir)
		return &driver, nil
	}
	opts.Logger.Debug("Creating the database at '%s' ...\n", dir)
	return &driver, os.MkdirAll(dir, 0755)
}

func (d *Driver) Write(collection, resource string, v interface{}) error {
	if collection == "" {
		return fmt.Errorf("missing collection - no place to store record")
	}
	if resource == "" {
		return fmt.Errorf("missing resource - unable to store record")
	}
	mutex := d.getOrCreateMutex(collection)
	mutex.Lock()
	defer mutex.Unlock()
	dir := filepath.Join(d.dir, collection)
	fnlPath := filepath.Join(dir, resource+".json")
	tmpPath := fnlPath + ".tmp"
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}
	b, err := json.MarshalIndent(v, "", "\t")
	if err != nil {
		return err
	}
	b = append(b, byte('\n'))
	if err := os.WriteFile(tmpPath, b, 0644); err != nil {
		return err
	}
	return os.Rename(tmpPath, fnlPath)
}

func (d *Driver) Read(collection, resource string, v interface{}) error {
	if collection == "" {
		return fmt.Errorf("missing collection - no place to read record")
	}
	if resource == "" {
		return fmt.Errorf("missing resource - unable to read record")
	}
	record := filepath.Join(d.dir, collection, resource)
	if _, err := stat(record); err != nil {
		return err
	}
	b, err := os.ReadFile(record + ".json")
	if err != nil {
		return err
	}
	return json.Unmarshal(b, &v)
}

func (d *Driver) ReadAll(collection string) ([]string, error) {
	if collection == "" {
		return nil, fmt.Errorf("missing collection - unable to read")
	}
	dir := filepath.Join(d.dir, collection)
	if _, err := os.Stat(dir); err != nil {
		return nil, err
	}
	files, _ := os.ReadDir(dir)
	var records []string
	for _, file := range files {
		b, err := os.ReadFile(filepath.Join(dir, file.Name()))
		if err != nil {
			return nil, err
		}
		records = append(records, string(b))
	}
	return records, nil
}

func (d *Driver) Delete(collection, resource string) error {
	path := filepath.Join(collection, resource)
	mutex := d.getOrCreateMutex(collection)
	mutex.Lock()
	defer mutex.Unlock()
	dir := filepath.Join(d.dir, path)
	switch fi, err := stat(dir); {
	case fi == nil, err == nil:
		return fmt.Errorf("unable to find file or directory named %v", path)
	case fi.Mode().IsDir():
		return os.RemoveAll(dir)
	case fi.Mode().IsRegular():
		return os.RemoveAll(dir + ".json")
	}
	return nil
}

func (d *Driver) getOrCreateMutex(collection string) *sync.Mutex {
	d.mutex.Lock()
	defer d.mutex.Unlock()
	m, ok := d.mutexes[collection]
	if !ok {
		m = &sync.Mutex{}
		d.mutexes[collection] = m
	}
	return m
}

func stat(path string) (fi os.FileInfo, err error) {
	if fi, err = os.Stat(path); os.IsNotExist(err) {
		fi, err = os.Stat(path + ".json")
	}
	return
}

type Address struct {
	City    string
	State   string
	Country string
	Pincode json.Number
}

type User struct {
	Name    string
	Age     json.Number
	Contact string
	Company string
	Address Address
}

func main() {
	dir := "./"

	db, err := New(dir, nil)
	if err != nil {
		fmt.Println("Error: ", err)
	}

	employees := []User{
		{"John", "23", "123456789", "Software Factory", Address{"CABA", "Buenos Aires", "Argentina", "1403"}},
		{"Daniel", "20", "987654321", "Tech Town", Address{"Cordoba", "Cordoba", "Argentina", "3025"}},
		{"Sabrina", "33", "123459876", "SAP", Address{"Rosario", "Santa Fe", "Argentina", "7786"}},
		{"Marta", "30", "543216789", "PoliTech", Address{"Rio Gallegos", "Santa Cruz", "Argentina", "9971"}},
		{"Santiago", "28", "775533995", "Arquitechture", Address{"Gualeguaychu", "Entre Rios", "Argentina", "4562"}},
		{"Michel", "25", "774441258", "Tecno Point", Address{"CABA", "Buenos Aires", "Argentina", "1238"}},
	}

	for _, employee := range employees {
		db.Write("users", employee.Name, User{
			Name:    employee.Name,
			Age:     employee.Age,
			Contact: employee.Contact,
			Company: employee.Company,
			Address: employee.Address,
		})
	}

	records, err := db.ReadAll("users")
	if err != nil {
		fmt.Println("Error: ", err)
	}
	fmt.Println("Records: ", records)

	users := []User{}
	for _, user := range records {
		employeeFound := User{}
		if err := json.Unmarshal([]byte(user), &employeeFound); err != nil {
			fmt.Println("Error: ", err)
		}
		users = append(users, employeeFound)
	}
	fmt.Println((users))

	if err := db.Delete("users", ""); err != nil {
		fmt.Println("Error: ", err)
	}
}
