// Code generated for package migr by go-bindata DO NOT EDIT. (@generated)
// sources:
// db/migrations/0001_setup.down.sql
// db/migrations/0001_setup.up.sql
// db/migrations/0002_bot.down.sql
// db/migrations/0002_bot.up.sql
package migr

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"net/http"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

// Name return file name
func (fi bindataFileInfo) Name() string {
	return fi.name
}

// Size return file size
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}

// Mode return file mode
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}

// Mode return file modify time
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}

// IsDir return file whether a directory
func (fi bindataFileInfo) IsDir() bool {
	return fi.mode&os.ModeDir != 0
}

// Sys return file is sys mode
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}


type assetFile struct {
	*bytes.Reader
	name            string
	childInfos      []os.FileInfo
	childInfoOffset int
}

type assetOperator struct{}

// Open implement http.FileSystem interface
func (f *assetOperator) Open(name string) (http.File, error) {
	var err error
	if len(name) > 0 && name[0] == '/' {
		name = name[1:]
	}
	content, err := Asset(name)
	if err == nil {
		return &assetFile{name: name, Reader: bytes.NewReader(content)}, nil
	}
	children, err := AssetDir(name)
	if err == nil {
		childInfos := make([]os.FileInfo, 0, len(children))
		for _, child := range children {
			childPath := filepath.Join(name, child)
			info, errInfo := AssetInfo(filepath.Join(name, child))
			if errInfo == nil {
				childInfos = append(childInfos, info)
			} else {
				childInfos = append(childInfos, newDirFileInfo(childPath))
			}
		}
		return &assetFile{name: name, childInfos: childInfos}, nil
	} else {
		// If the error is not found, return an error that will
		// result in a 404 error. Otherwise the server returns
		// a 500 error for files not found.
		if strings.Contains(err.Error(), "not found") {
			return nil, os.ErrNotExist
		}
		return nil, err
	}
}

// Close no need do anything
func (f *assetFile) Close() error {
	return nil
}

// Readdir read dir's children file info
func (f *assetFile) Readdir(count int) ([]os.FileInfo, error) {
	if len(f.childInfos) == 0 {
		return nil, os.ErrNotExist
	}
	if count <= 0 {
		return f.childInfos, nil
	}
	if f.childInfoOffset+count > len(f.childInfos) {
		count = len(f.childInfos) - f.childInfoOffset
	}
	offset := f.childInfoOffset
	f.childInfoOffset += count
	return f.childInfos[offset : offset+count], nil
}

// Stat read file info from asset item
func (f *assetFile) Stat() (os.FileInfo, error) {
	if len(f.childInfos) != 0 {
		return newDirFileInfo(f.name), nil
	}
	return AssetInfo(f.name)
}

// newDirFileInfo return default dir file info
func newDirFileInfo(name string) os.FileInfo {
	return &bindataFileInfo{
		name:    name,
		size:    0,
		mode:    os.FileMode(2147484068), // equal os.FileMode(0644)|os.ModeDir
		modTime: time.Time{}}
}

// AssetFile return a http.FileSystem instance that data backend by asset
func AssetFile() http.FileSystem {
	return &assetOperator{}
}

var __0001_setupDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\x09\xf2\x0f\x50\x08\x71\x74\xf2\x71\x55\x48\x2e\x2d\x2a\x4a\xcd\x4b\xce\x4c\x2d\xb6\xe6\x42\x12\x4e\xad\x48\xce\x48\xcc\x4b\x4f\x2d\x42\x15\x2e\x2d\xc6\x26\x12\x9f\x94\x98\x93\x98\x97\x8c\x66\x44\x41\x62\x65\x6e\x6a\x5e\x49\x7c\x72\x7e\x5e\x5a\x66\x51\x6e\x62\x49\x66\x7e\x1e\xaa\x8a\xfc\xa2\x14\x90\x71\x80\x00\x00\x00\xff\xff\x1d\x54\x4a\x76\x8e\x00\x00\x00")

func _0001_setupDownSqlBytes() ([]byte, error) {
	return bindataRead(
		__0001_setupDownSql,
		"0001_setup.down.sql",
	)
}

func _0001_setupDownSql() (*asset, error) {
	bytes, err := _0001_setupDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "0001_setup.down.sql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __0001_setupUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x55\x41\x6f\xe2\x3c\x10\xbd\xf3\x2b\xe6\x08\x12\xff\xa0\xa7\x40\x4d\x15\x7d\x21\xe1\x0b\x46\x5a\x4e\x96\x6b\x0f\xc5\x2a\xb1\xa9\x9d\xb4\xcb\xbf\x5f\x05\x12\x6a\x8b\x04\xd2\x3d\xef\x35\xf3\x3c\x6f\xf2\xde\x3c\x7b\x9e\x93\x88\x12\xa0\xd1\x2c\x21\x20\x2a\x6b\x51\x0b\x85\x0e\xc6\x23\x00\x25\x61\x16\xbf\xac\x49\x1e\x47\x09\xac\xf2\x78\x19\xe5\x5b\xf8\x8f\x6c\xa7\x23\x00\x61\x24\x02\x25\xbf\x28\x6c\xd2\xf8\xff\x0d\x81\x34\xa3\x90\x6e\x92\xa4\x2e\x4a\x74\xc2\xaa\x63\xa9\x8c\xee\xc4\x8c\x26\x4f\xa3\x80\x18\x7f\x8b\x3d\xd7\x6f\x68\x1f\x13\x2b\x5d\x28\x0d\xcf\xd9\xa6\x3e\xb7\xca\xc9\x3c\x5e\xc7\x59\x7a\x9f\xde\x2f\x5a\xfc\xa8\x94\x45\x76\xe4\xa7\x02\x75\xc9\x3e\xd1\xaa\x9d\x12\xfc\x8c\x9e\x65\x59\x42\xa2\xb0\x9b\xd2\xac\xd1\xe5\xe4\x4d\xe5\x23\xe6\x59\xba\xa6\x79\x14\xa7\x14\x76\xef\x67\x30\x53\x1a\x16\x59\x4e\xe2\x97\xb4\x9e\x7b\xec\xf5\x98\x40\x4e\x16\x24\x27\xe9\x9c\xac\x3d\xc1\xc7\x4a\x4e\xea\x5e\xa6\x2a\xff\x82\xce\x54\x65\xc0\xe7\x77\x79\x40\x58\x69\xf5\x51\xa1\x3f\xe1\x34\x18\x62\x72\x63\x56\xe5\x86\xf8\x84\x05\x57\x87\xde\x0d\xb9\xa8\x8e\xb2\x53\xf1\x23\x77\xee\x6b\xcf\xdd\xfe\xd6\x3d\x2e\x6b\xf7\xbb\x0e\x99\x23\x5a\x5e\x1a\xdb\x59\x2c\xcd\x3b\x76\xec\xc2\x6b\xe5\x4e\x37\xf8\x9b\xff\x7d\xe5\x07\xae\xc5\x80\x4c\xd4\xc2\xb0\x00\xd0\x6f\x5b\x8b\xf5\x5d\x6b\xbe\x05\x86\x9d\xc5\x6e\xbd\x6a\x3d\x19\x4a\xe2\xe3\x7d\x22\xef\xfb\xa0\xed\x80\x76\xb4\xa9\x3f\xc2\x19\xd0\x88\x73\x3f\x90\x5c\x4a\x8b\xce\x0d\xbb\x0b\x04\xb7\x92\x09\xa3\x77\xca\x16\xe7\x54\xfe\x13\xbe\x4b\xf8\x7b\x9a\x3e\x4c\x98\x2a\xf8\x1b\xc2\x6c\x4b\x49\x74\xa3\xbf\xb1\x72\x48\xbe\x7f\xaa\xf9\xcf\x04\x6f\xd3\x3c\x94\xe1\x9a\xfe\xe0\x16\xfc\x6e\xd2\xcf\x74\x7d\x7a\x86\x52\x05\x07\x7c\x3a\xbf\x10\xf0\x7d\xbf\x6e\x2d\x29\x2f\x4c\xa5\x4b\xf6\xe8\x21\x6b\x60\xf5\xed\x7e\x17\x67\x51\xa0\xfa\x44\x16\xec\x84\x0f\x10\x16\x79\x89\x92\xf1\x12\x68\xbc\x24\x6b\x1a\x2d\x57\x57\x00\x3c\x93\x45\xb4\x49\x28\x68\xf3\x35\xbe\xac\x7b\x1d\xe9\xc3\xa1\x67\x79\x76\x4a\x2b\xb7\xef\x29\x36\xc1\x65\xde\x86\x5d\xae\xf4\xcb\x6b\xdb\x94\x3b\x0e\x8f\x26\x4f\x7f\x02\x00\x00\xff\xff\xba\x12\x9f\x60\x8c\x08\x00\x00")

func _0001_setupUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__0001_setupUpSql,
		"0001_setup.up.sql",
	)
}

func _0001_setupUpSql() (*asset, error) {
	bytes, err := _0001_setupUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "0001_setup.up.sql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __0002_botDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\x09\xf2\x0f\x50\x08\x71\x74\xf2\x71\x55\x48\xca\x2f\x89\xcf\x4d\x2d\x2e\x4e\x4c\x4f\x2d\xb6\x06\x04\x00\x00\xff\xff\x10\xe1\xe1\xce\x18\x00\x00\x00")

func _0002_botDownSqlBytes() ([]byte, error) {
	return bindataRead(
		__0002_botDownSql,
		"0002_bot.down.sql",
	)
}

func _0002_botDownSql() (*asset, error) {
	bytes, err := _0002_botDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "0002_bot.down.sql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __0002_botUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\xcf\xc1\xca\xc2\x30\x10\x04\xe0\x7b\x9f\x62\x8e\x2d\xfc\x6f\xf0\x9f\xd2\xb2\x2d\xc1\xb0\x81\x74\x05\x3d\x15\x6d\xa2\x96\x22\x85\x46\xdf\x5f\x8c\x55\xec\xcd\xeb\xec\xcc\x07\x5b\x39\x52\x42\x10\x55\x1a\xc2\x71\xba\x75\xd7\x10\xe3\xe1\x1c\x22\xf2\x0c\x00\xee\x31\xcc\xdd\xe0\x51\xea\xa6\x25\xa7\x95\xf9\x4b\x71\x65\xb9\x15\xa7\x34\x0b\x4e\x63\xf7\x2c\xa1\xb6\x8e\x74\xc3\xd8\xd0\x3e\x5f\x56\x05\x1c\xd5\xe4\x88\x2b\x6a\x93\x14\xf3\xc1\x17\x2f\x61\x9a\xfd\x0f\x72\x6a\xad\xe8\xf7\x6e\x65\xa7\xf0\x0b\x5f\x9e\x80\xd0\x4e\xc0\x56\xc0\x5b\xb3\xf8\xfd\x25\xf4\x63\xf0\x28\xad\x35\xa4\xf8\x73\xcd\x8a\xff\x47\x00\x00\x00\xff\xff\x99\x12\xfd\xd3\x0c\x01\x00\x00")

func _0002_botUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__0002_botUpSql,
		"0002_bot.up.sql",
	)
}

func _0002_botUpSql() (*asset, error) {
	bytes, err := _0002_botUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "0002_bot.up.sql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"0001_setup.down.sql": _0001_setupDownSql,
	"0001_setup.up.sql":   _0001_setupUpSql,
	"0002_bot.down.sql":   _0002_botDownSql,
	"0002_bot.up.sql":     _0002_botUpSql,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}

var _bintree = &bintree{nil, map[string]*bintree{
	"0001_setup.down.sql": &bintree{_0001_setupDownSql, map[string]*bintree{}},
	"0001_setup.up.sql":   &bintree{_0001_setupUpSql, map[string]*bintree{}},
	"0002_bot.down.sql":   &bintree{_0002_botDownSql, map[string]*bintree{}},
	"0002_bot.up.sql":     &bintree{_0002_botUpSql, map[string]*bintree{}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}
