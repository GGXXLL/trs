// Code generated by go-bindata. DO NOT EDIT.
// sources:
// NAME-service/gen/client/grpc/client.gotemplate (3.184kB)
// NAME-service/gen/client/http/client.gotemplate (105B)
// NAME-service/gen/endpoints.gotemplate (4.609kB)
// NAME-service/gen/transport_grpc.gotemplate (3.149kB)
// NAME-service/gen/transport_http.gotemplate (106B)
// NAME-service/service.gotemplate (62B)

package template

import (
	"bytes"
	"compress/gzip"
	"crypto/sha256"
	"fmt"
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
		return nil, fmt.Errorf("read %q: %w", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("read %q: %w", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes  []byte
	info   os.FileInfo
	digest [sha256.Size]byte
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _genClientGrpcClientGotemplate = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x56\x4d\x6f\xdb\x38\x10\x3d\x8b\xbf\x62\xd6\x08\x16\x52\xa0\xd0\xf7\x2c\x7c\xa9\xd3\x2d\xba\xd8\xa6\x46\x1a\x74\x0f\x45\x51\x30\xd4\x58\x26\x2c\x93\x2a\x49\x3b\x31\x04\xfd\xf7\xc5\x90\x94\x23\x27\x8e\xdb\x43\x10\x8b\xf3\x38\x1f\xef\x0d\x39\x9c\x4e\x61\x6e\x2a\x84\x1a\x35\x5a\xe1\xb1\x82\x87\x3d\x78\xbb\x75\x8e\xc3\xcd\x67\xb8\xfd\x7c\x0f\xef\x6f\x3e\xde\x73\x36\x9d\xc2\x1d\xda\xad\xd6\x4a\xd7\x11\x00\x8f\xaa\x69\xc0\xec\xd0\x3e\x5a\xe5\x11\xfc\x4a\x39\x58\xaa\x06\x03\xf8\x2b\x5a\xa7\x8c\xbe\x86\xae\xe3\xe9\x77\xdf\x8f\x0c\x70\x23\x3c\x8e\xad\xf4\xdd\xf7\x8c\x20\x0b\x21\xd7\xa2\x46\xa8\x6d\x2b\xa1\xb5\x66\xa7\x2a\x74\x20\xa0\xbe\x5b\xcc\x41\x36\x0a\xb5\x87\xa5\xb1\xe0\x57\x48\x0e\xbe\xa0\xdd\x29\x89\xfc\x56\x6c\xb0\xef\xc1\xa5\x4f\xd6\x8e\xdc\x30\xa6\x36\xad\xb1\x1e\x72\x96\x4d\xa4\xd1\x1e\x9f\xfc\x84\x65\x93\xda\x98\xba\x41\x5e\x9b\x46\xe8\x9a\x1b\x5b\x4f\x09\xfd\xb6\x65\xba\x41\x2f\x2a\xe1\x45\x80\x28\xbf\xda\x3e\x70\x69\x36\xd3\x76\x5d\x4f\xd1\x5a\x63\xdd\x84\x1d\x5b\x6a\x73\xb5\x56\x7e\x4a\x7f\xa8\xab\xd6\x28\x4d\x81\xc9\x97\xb7\x42\xbb\x90\xd4\x1b\xf8\x03\x20\x25\xc5\xb2\xe9\x14\xee\x89\xe6\x54\x32\xcb\x26\x5d\xc7\x3f\x86\xca\x16\xc2\xaf\xe0\xaa\xef\x61\xea\x76\x54\x40\xfb\x00\x64\x5c\xbc\x3b\x36\x4f\x58\x11\x38\xbe\xc5\x47\xb0\xe8\xb7\x56\x3b\x10\x7a\x20\x0d\x1e\x84\x5c\xc7\x26\x38\xa6\x5b\x1a\xad\x51\x7a\x65\x34\x87\x8f\x1e\x94\x23\xf2\xc9\x8f\x45\xd7\x1a\xed\xd4\x83\x6a\x94\xdf\x83\x59\x06\x55\xa4\x68\x1a\xb4\xe0\x0d\x54\x4a\x34\x25\x08\x5d\x41\x23\x3c\x5a\x90\x8d\x71\x58\x46\xd0\xb3\x4f\xb6\xdc\x6a\x49\x39\xe5\xb4\x08\x97\x54\x2f\x9f\x87\xd0\x73\xa3\x75\x09\xa6\x25\x9c\x03\xce\xd3\xf2\xe7\xb0\x50\x40\xde\x3e\xf0\x57\x3d\x40\x5f\x68\x4b\x08\x8a\x14\xd0\xb1\x6c\x27\x2c\x48\x99\xaa\x99\x1b\xbd\x54\x35\x63\x19\x35\xd1\x8f\x12\x96\x70\x3d\x03\x2b\x74\x8d\x87\x38\x1d\xcb\x32\xb4\x96\x0c\xcb\xfc\x4f\x29\x0b\x96\x65\x6a\x49\x0e\xe1\x8f\x19\x68\xd5\x04\x44\x16\x19\xa4\xef\x14\xcc\xf1\xff\xac\x68\x73\xb4\xb6\x84\x89\x14\x5a\x1b\x0f\xa2\x6d\x9b\x7d\xf2\x3c\x21\x47\x3d\xcb\x7a\xc6\x32\x39\x2a\xc4\x51\xa4\x6f\xdf\x8f\xda\xe2\xa8\x52\x0a\x77\xca\xfa\x0e\x97\xc6\x62\x4e\xc9\xa4\xb6\xfe\x2a\x9a\x2d\xba\x7b\xf3\xe1\x6e\x31\xff\x94\xba\x35\x97\x92\xaf\x50\x54\x68\x5d\x51\x94\x14\x3e\xeb\xba\x2b\x78\x54\x7e\x05\x17\x1e\x29\x38\xef\x7b\x96\x8d\x56\xdb\x75\x4d\x64\x92\xe9\xc2\x23\x4f\x67\x32\xf2\x4b\xd1\x08\x19\x39\xbb\x50\x03\x68\x50\xe1\x13\xfa\x95\xa9\x5c\x04\x06\xee\xbb\xee\xde\xfc\x6b\x1e\xd1\xc2\x85\x4a\x22\xbd\x4f\xa7\x01\x86\x63\xc1\x87\x95\xb0\x2b\xf0\x4b\x61\xde\xde\x38\x83\x63\x46\x6e\xf1\x31\x92\x92\xc7\xbd\xc4\x88\x2e\xd3\xef\x49\xd7\x0d\x35\xf5\x3d\xef\xba\x71\xbe\x71\x71\x32\x86\xaa\x97\x8b\xef\xb5\x34\x15\x12\xa9\x23\xeb\x1d\xfe\xdc\xa2\xf3\x03\xe6\x06\x4f\x62\xc2\x09\xc1\x01\x14\x1a\xf6\x83\x09\xe4\x5e\x28\x3e\x98\xef\xf7\xed\x90\x48\xd7\x0f\xd8\xa3\x16\xe1\x9c\xa7\xf5\xe2\x40\x55\x5e\x84\x95\xa4\x08\xea\x2a\xa9\x98\x7e\x0d\x3f\xd8\xd0\xa9\x6e\x27\x0f\x7b\x5d\x47\x80\xb1\x86\x2f\x05\xa4\x0b\x23\xb8\x7b\xc5\xfd\x35\x00\x9c\x13\xb5\x7c\x8e\x9d\xf5\x25\x1d\x10\x16\xef\x76\x22\x07\xa2\x4a\x10\xe9\x62\xe7\x73\x88\x53\xe3\x2c\xb3\x74\x1d\x09\x38\xbe\x2d\x79\xdc\x31\x40\xfe\xa6\xfb\xc5\xaf\x44\xb8\xc9\x76\x68\xbd\x03\x41\x7e\xc3\x1d\x77\xa2\x0e\xb0\x48\x87\xd6\x1b\x10\xb0\x75\x68\xaf\x2a\xb3\x11\x4a\xbf\x01\x8d\x31\x38\x2c\xac\xda\x08\xab\x9a\x3d\xed\x59\x6e\x1b\x50\x1a\x44\xba\x74\xd2\x1d\x77\xb6\x90\xfc\x07\xa4\x43\xcc\xe7\xf1\x7f\x19\x5a\xfc\x2e\x24\xa3\xb4\x47\xbb\x14\x12\xbb\xbe\x80\x7c\xf4\x35\xbe\xe8\x62\xde\xd7\xb3\xe7\x7d\x3c\xbf\xfc\x75\xcb\x15\x87\x0e\x09\x0e\x06\xc5\x0e\xfd\xf3\x42\xb9\x78\x18\x7e\x4b\xb9\x73\xe7\xe6\xa4\x70\x71\x43\x42\xbc\xa5\xdb\xaf\x35\x89\x01\x82\x80\x67\x44\x0e\xa8\xdf\x12\xee\x5c\x1d\xa7\x74\x1b\x32\xf8\x4d\xd5\x7e\x86\x19\x94\xf2\x39\xa1\x58\x30\xbc\x21\xd8\xcf\x57\x72\x31\xbf\x6f\xf1\x68\xda\x81\xf3\x76\x2b\x3d\x05\x4b\x83\x00\xbe\x7d\x77\xde\x2a\x5d\xa7\x93\x39\x9e\x36\x51\x18\xaa\x3b\x7c\x05\x01\x36\xa6\x52\x4b\x85\x2e\xce\xee\xc3\xb3\x80\x26\x69\x88\x76\xb4\x9f\xb6\xe6\x97\xe3\x04\x8a\x58\x2e\x8b\x6c\xce\xfd\xd3\x30\xa7\xbe\xa0\xae\xf2\x35\xee\xc3\x70\x8f\x19\x15\xc7\xce\xba\x43\xad\xc1\xad\x81\x53\x8e\xc3\x40\x36\xc3\x94\x83\x19\x90\x4b\x36\x1e\xd1\x34\xf6\xfa\x14\xff\xdc\xac\x0c\xb9\x0c\xe4\x14\x70\x6a\xea\x8e\xbb\xf3\x45\x76\xd2\x3f\xbd\x6e\x86\x4d\x05\x97\xc3\xcb\x91\x7f\xba\x29\x5e\x22\x42\xf2\x34\x27\x5b\xa1\xc6\xca\x64\xc3\x13\x65\xfd\xfc\x44\x09\xe9\x85\xe9\xa8\x96\xb0\x2b\xc1\x04\x9b\xf4\x4f\x3c\x54\x93\xaf\x0b\x9e\xa7\xdc\xff\x22\x63\x1c\xa4\xd1\xf1\x8c\x1e\x23\xc4\x77\xf8\x2c\x61\x5d\xc2\x2e\x4c\x90\x3e\x3c\x4b\xe2\x23\x27\x42\xc7\xcf\x9c\xcb\x4d\x05\x33\x38\x14\xf0\x8f\x51\x3a\xbf\xdc\x54\xe5\xf3\xd2\x82\xf6\x44\xaf\x9c\xf3\xa2\x18\xdc\x25\x66\xa4\x7f\x8a\xec\xff\x1f\x00\x00\xff\xff\x00\xce\x0e\xa6\x70\x0c\x00\x00")

func genClientGrpcClientGotemplateBytes() ([]byte, error) {
	return bindataRead(
		_genClientGrpcClientGotemplate,
		"gen/client/grpc/client.gotemplate",
	)
}

func genClientGrpcClientGotemplate() (*asset, error) {
	bytes, err := genClientGrpcClientGotemplateBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "gen/client/grpc/client.gotemplate", size: 3184, mode: os.FileMode(0644), modTime: time.Unix(1464111000, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x63, 0x72, 0x1f, 0xe5, 0x3a, 0x45, 0x1, 0x91, 0xd8, 0x5b, 0xa8, 0x47, 0x45, 0x45, 0x98, 0xee, 0x0, 0xf5, 0xc1, 0x3c, 0x43, 0xf0, 0x86, 0x3c, 0xec, 0xbe, 0x2d, 0x84, 0xed, 0x1a, 0x17, 0x6c}}
	return a, nil
}

var _genClientHttpClientGotemplate = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2c\xcb\xb1\x0d\xc2\x30\x10\x05\xd0\x3e\x53\x5c\x1d\x09\xdf\x10\x34\x29\x91\xc8\x02\x56\xf8\x98\x88\xc3\x67\x9d\x3f\x95\xe5\xdd\x69\x18\xe0\x8d\xa1\xab\xdc\x01\x29\x7e\x61\x7c\x7b\xd7\x82\x5a\xfc\x7d\x52\x5f\x64\x63\xe4\xda\x9b\x07\x95\xf8\x34\xcb\x44\x2a\x2e\x4f\x0f\x39\xfc\x01\x59\x75\xce\x65\x8c\x23\x9b\x49\xda\xf6\xfd\xb6\xc1\x1a\x22\x5d\xed\x44\xe5\xfe\x27\x92\xe6\x5c\x7e\x01\x00\x00\xff\xff\x0b\x3c\x4c\x9e\x69\x00\x00\x00")

func genClientHttpClientGotemplateBytes() ([]byte, error) {
	return bindataRead(
		_genClientHttpClientGotemplate,
		"gen/client/http/client.gotemplate",
	)
}

func genClientHttpClientGotemplate() (*asset, error) {
	bytes, err := genClientHttpClientGotemplateBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "gen/client/http/client.gotemplate", size: 105, mode: os.FileMode(0644), modTime: time.Unix(1464111000, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xa1, 0xf0, 0x36, 0xf9, 0x16, 0xea, 0x9d, 0x4e, 0x73, 0x64, 0xc5, 0xad, 0xb3, 0x1b, 0x4, 0xe, 0xd8, 0xc8, 0x1e, 0xf7, 0x7a, 0x39, 0x40, 0x4c, 0xb2, 0x12, 0x83, 0x35, 0xca, 0x82, 0x6f, 0xd0}}
	return a, nil
}

var _genEndpointsGotemplate = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x57\xcf\x6f\xdc\xba\x11\x3e\x8b\x7f\xc5\x64\xe1\xc2\xbb\x81\x2c\xdf\x1d\xec\xa1\x4d\xdc\xd6\x40\xe2\x04\xb1\xdb\x1e\x82\x20\xe0\x4a\xb3\x2b\xc2\x14\xc9\x90\xd4\xfe\xa8\xa0\xff\xbd\x18\x52\xd2\x6a\xb3\xb2\xeb\xf7\x8e\x0f\xef\x60\xd8\x26\x87\xc3\x6f\xbe\x6f\x66\x38\xba\xbe\x86\xf7\xba\x40\xd8\xa0\x42\xcb\x3d\x16\xb0\x3a\x80\xb7\xb5\x73\x19\x7c\xf8\x0c\xf7\x9f\x1f\xe1\xf6\xc3\xdd\x63\xc6\xae\xaf\xe1\x2b\xda\x5a\x29\xa1\x36\xd1\x00\x76\x42\x4a\xd0\x5b\xb4\x3b\x2b\x3c\x82\x2f\x85\x83\xb5\x90\x18\x8c\xff\x8d\xd6\x09\xad\x6e\xa0\x69\xb2\xee\xef\xb6\x1d\x6d\xc0\x07\xee\x71\xbc\x4b\xff\xb7\x2d\x63\x86\xe7\x4f\x7c\x13\x10\x31\xb2\x7f\xec\xdd\x42\xae\x95\xe7\x42\x39\xa8\xd0\x97\xba\x70\xe0\x35\x54\xfc\x09\x41\xa8\x42\x6c\x45\x51\x73\x09\xa8\x0a\xa3\x85\xf2\x0e\xd6\x56\x57\xe0\xd0\x6e\x45\x8e\x2e\x25\x4f\x16\x7f\xd6\xe8\x3c\x70\x55\x80\x45\x67\xb4\x72\x08\xfe\x60\x30\x78\x22\x53\x0a\x42\x3b\x3c\x7a\x49\x81\x3b\xd8\xa1\x94\xf4\x1b\x55\xae\x0b\xb4\x8e\x1c\x90\xbf\x02\xbb\xff\xd7\xda\x76\x07\x83\xb7\x34\x2c\x70\x22\x67\x0d\xba\xb6\xe0\x6a\x63\xb4\x25\x72\xbd\xe5\xca\xd1\xdf\x74\x9d\xe0\x52\xfc\x97\x7b\xa1\x15\x79\x5b\x6b\x5b\x71\xef\x32\xc6\x44\x15\x2c\xe6\x2c\x99\xad\x2b\x3f\x63\xc9\x8c\x22\xc7\xbd\x9f\x31\x96\xcc\x36\xc2\x97\xf5\x2a\xcb\x75\x75\xbd\xd1\x57\x4f\xc2\x5f\xd3\x4f\x8f\x98\x4c\xcc\x0a\x66\x4d\x93\x7d\xf9\xdb\x5d\x70\xf4\x85\xfb\x12\xae\xda\x76\xc6\x16\x81\xd0\xdb\x81\xa2\x5c\x4b\x89\xb9\x77\x3d\x56\x5f\x8e\x42\x07\x5f\x72\x0f\xb9\xae\x0c\x05\xc6\x15\xf0\xa2\xe8\xf9\xcc\xe0\xce\x5f\x3a\x72\x56\x21\x57\x9e\xe8\x5b\x21\xd4\x0e\x0b\xe2\x89\x43\x89\xd2\xa0\x05\xe7\x6d\x9d\xfb\x94\xb6\xbb\xab\xa6\x6f\x12\xca\x6b\xe0\xe4\xce\x09\xb5\x91\x08\x86\x5b\x5e\xa1\x47\x4b\xa9\x44\xeb\x77\x0a\x78\x54\xc8\xa6\x20\xfc\xa5\xa3\xcb\xd6\xb5\x0c\x4c\xaf\x6b\x95\x13\x8b\x1d\x64\x85\x44\xb4\x06\x6d\x42\x46\x83\xa6\xb3\x06\xed\x55\x7f\x21\x39\x5c\x71\x27\x5c\x06\x7f\xd7\x16\x70\xcf\x2b\x23\x31\x85\x83\xae\xa1\x12\x9b\xd2\x83\xe1\x8e\x54\x1e\x51\x45\x00\x87\x8b\xe2\x3d\xc6\xea\xa2\xce\x31\xd0\xc0\x15\x94\xde\x9b\xec\x9f\x5c\x15\x92\x30\xee\x84\x2f\x01\x79\x5e\x76\xc9\x0a\xf3\xfe\xf6\x05\xec\x84\xc5\x02\x6a\x13\x9d\x3a\x83\xb9\x58\x8b\x1c\x0c\xf7\x65\x06\xf3\xbb\x80\x4f\x38\xf2\xbf\xe2\x2b\x79\x00\x0e\x95\x70\x3e\x26\x3a\x14\xe8\xc4\x46\xd1\x51\xa1\xb6\xfa\x09\x03\x95\x0f\x51\x96\xa1\x30\x02\x44\x3c\x15\x3b\x8a\x41\x2e\x7a\x26\xb3\xc5\x98\xdd\x5c\x0a\x54\xfe\x94\xdd\x91\x70\xc7\x1a\x93\x07\xaa\xc4\xe8\x0e\x8b\x97\x64\xa4\x6a\x88\x5c\x09\x62\xb8\xc2\x98\x56\x47\xbc\x42\x79\xb4\x6b\x4e\x09\x35\xad\x04\x39\x1b\x2e\x9b\xae\xf3\xda\xc5\x8e\xd4\x15\xd6\x75\xd0\xe1\x1e\x77\xef\xbb\x78\x72\x5d\xad\x84\x0a\x3c\x55\x1d\xc4\x91\xb0\x69\xd7\x0d\x7c\x6d\x15\x88\x90\xc9\x04\x30\xe7\x52\xa2\x8d\xc9\xdc\x81\xcd\x58\x08\xe7\x8c\xd0\x86\x35\x8d\xe5\x6a\x83\x70\x21\xe0\x66\x09\x59\x6f\xff\x29\x8a\xd1\xb6\x2c\x69\x9a\x0b\x91\xdd\xf3\x0a\xdb\xb6\x3f\x0f\x00\x43\x10\x59\xbf\xc8\x9a\xe6\x8a\x56\xdb\x96\xb5\x8c\x51\xba\xc1\x3d\xee\x86\x2b\xe7\x5d\xf9\x81\x59\x65\x4d\x33\x5c\x14\x1d\x3f\x04\x4d\x17\x23\x80\x0d\x63\xc9\x28\x05\xa0\xd0\x15\x17\x2a\x63\xc9\x96\x5b\xea\x2e\x2f\xe3\xa6\x96\xc1\x92\xa4\x69\x1e\xf5\x47\xbd\x43\x0b\xe7\x21\x2c\xe1\x13\x7f\xc2\x89\xe0\x7a\xa0\x0b\xba\x24\xc6\x93\x2c\x18\x4b\x8e\xaa\xdd\x2c\x8f\x40\x9b\x57\x43\x39\xbb\xe8\x86\x68\x7c\x01\x62\x3a\x02\xd0\x32\x96\x74\x42\x0f\x38\x88\xe6\x71\x95\xbc\x42\xcb\xa0\xca\x7c\x94\x08\x0b\x18\x01\x9b\xe7\x7e\x0f\x5d\xbb\xce\xde\xc7\xdf\x29\x15\xdd\xdb\xa0\xd9\x3f\x34\x99\x11\xce\xaf\xf1\x31\x7a\x3c\x98\x5e\xc0\x05\xcc\xcf\x8d\xe2\x2b\x35\xb2\x4a\x01\xad\xd5\x76\x01\x0d\x4b\x92\xfe\x15\x0b\x8b\x04\x18\xb3\x29\x35\x72\xbf\x27\x0c\x0b\x96\x24\x62\x1d\x4c\xdf\x2c\x41\x09\x19\x7c\xf4\x9c\x28\x21\x83\x1b\x96\x24\x2d\x1b\x56\xfb\x1b\xb2\xd7\x60\x5b\xa4\xe4\x85\x98\xee\x49\x27\x72\x29\x47\x4e\x18\x0e\xbd\xf1\xc2\x63\x60\x38\x96\xc7\x98\xf4\x0b\x8f\x53\xbc\x47\xe2\x9f\x4d\xb8\x58\x13\xe3\xb3\xa7\x65\x71\x56\x6b\x27\xc1\x93\xef\x69\xe9\xfa\xa1\x61\x68\x55\x0d\x09\x35\x8c\x0f\xa3\xe5\x28\xc2\x48\x1d\xf2\xfe\x93\x22\xea\x7c\x4c\x71\x78\x96\x04\xe1\xdc\x76\x10\xd4\x65\xbf\x24\x57\x40\x14\xad\x26\xb4\x9c\x52\x33\xea\x39\xec\x6c\x3b\x91\xe2\x72\x7b\x2c\x90\xb1\x66\xff\xb1\xdc\xfc\x55\xca\xdb\x7d\x8e\xc6\xc3\xce\x72\xe3\xe2\x6b\x36\xb0\xb7\x16\x28\x0b\x7a\xca\xbb\x36\x78\x6c\x3b\x41\xde\xf0\x0c\x4c\xcc\x27\xd9\x27\x51\x14\x12\x77\xdc\xc6\x31\xf1\x5f\xae\x1f\x1c\x69\x64\x32\x46\x1e\xa8\x9b\xd3\x0b\xe5\xc9\x79\x35\x58\x87\x27\x18\xb7\x68\x0f\x83\x94\x54\x56\xd4\xac\xfb\xa1\x84\xfc\x7d\x36\xf4\x40\xd3\x23\x95\x8e\xde\x88\x9c\x2b\x1a\x50\xe8\x59\xc7\x82\x8e\xad\x0e\xa0\x48\x83\x38\xb8\xe0\x3e\x97\x75\x81\x45\x9c\x19\x57\x48\x10\x28\x66\x83\x45\x76\xc6\xc6\xfc\x88\x29\x85\xd9\x83\xe7\xbe\x76\xb3\x14\x66\x5f\x84\xda\xcc\x16\xac\x6f\x0f\x6f\x47\xfd\xe1\xb9\xf3\x30\xc1\x4a\x7a\x44\x93\x65\x99\xf3\x56\xa8\x4d\x48\x27\xa1\xba\xe5\x9b\x25\x54\xdc\x7c\x8b\x5b\xdf\x23\xfd\x4d\xdb\x84\x06\x79\x05\xff\xaf\x7d\x25\xc9\x6c\x94\x51\xb3\x1b\x68\xda\xb4\x3b\x3a\xea\x93\xa4\xc6\x0f\x82\x12\xd2\x37\xb8\x1c\x60\x35\xb1\x8d\xfc\x48\x41\x3f\xd1\x76\x0f\xec\x1b\xee\xbf\xbf\x83\x37\xfa\x29\xa6\xa2\xe1\x4a\xe4\xf3\x75\xe5\xb3\x07\x63\x85\xf2\xeb\xf9\xec\xb6\x77\x31\x28\x78\xf9\x17\x77\x09\x85\x46\x07\x4a\x7b\xc0\xbd\x70\xfe\x1d\x38\xc4\xb1\xf0\x43\xee\xb8\x6c\xa3\x67\x04\x6a\xb1\xe8\x9a\x54\x81\x12\x3d\xce\x7b\x04\x61\xef\x18\x80\x50\xf9\x11\xfe\x40\xdf\xeb\x89\x12\xeb\xe0\x62\xb9\x84\x13\xca\xba\x4a\x9b\x6c\xb5\xb0\x1c\x21\x9f\x4f\x9a\x2c\xfa\xd2\x3b\xa1\x3c\x96\xdd\x47\xbe\x42\x89\xc5\x31\x1b\xe2\x37\xd6\x06\x7d\x9f\xbb\xe3\xc1\x39\xa6\xf0\xae\x44\x35\xec\xea\x51\xba\x76\xce\x62\xd6\xa5\xb1\xca\xba\x42\xa8\xa3\x31\xc4\x0f\x37\x1e\xbf\xfe\x44\x4e\xf3\xa3\x15\x79\x1c\xec\x47\x18\x4a\x91\x97\xe1\xa8\x43\x35\x05\xa1\x1b\x9a\xba\xd3\xfd\xc8\xa8\x6d\x37\x32\x9d\x47\x15\xda\x6d\x4c\xe0\xf4\xbc\x33\x4f\x34\x6b\xf6\x5c\x5c\xbf\xbb\x37\x9d\x81\x4a\xbb\x38\x03\xe3\x16\x73\x14\xdb\x38\x5c\x87\x10\x7f\xf9\x66\xc9\xe0\x01\x71\x70\x33\xf2\x12\x36\xfa\x99\x7f\xa8\x7b\x02\x4a\x19\x59\xa0\xe7\x42\x86\xf9\xbc\x2f\xa7\xf0\xe9\xd7\x7d\x57\x70\x29\xfc\x21\x7b\xa9\x85\x9c\xc4\x3e\xee\x24\xbf\x99\xd1\x3f\xfb\xcc\x1f\xa7\xcf\x9c\x1c\x4b\xa7\x87\xc0\xe7\xda\xce\xff\x02\x00\x00\xff\xff\xea\x8c\xe5\x23\x01\x12\x00\x00")

func genEndpointsGotemplateBytes() ([]byte, error) {
	return bindataRead(
		_genEndpointsGotemplate,
		"gen/endpoints.gotemplate",
	)
}

func genEndpointsGotemplate() (*asset, error) {
	bytes, err := genEndpointsGotemplateBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "gen/endpoints.gotemplate", size: 4609, mode: os.FileMode(0644), modTime: time.Unix(1464111000, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xdd, 0x28, 0x78, 0xac, 0x5, 0x3c, 0x60, 0x1a, 0xeb, 0x12, 0x1b, 0x3e, 0x5a, 0x94, 0x6c, 0xee, 0x8a, 0x6b, 0xd7, 0xde, 0x58, 0xf8, 0xb3, 0xd9, 0xfa, 0xfa, 0xeb, 0x24, 0x8, 0x98, 0x2c, 0x98}}
	return a, nil
}

var _genTransport_grpcGotemplate = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x56\xc1\x6e\xdb\x38\x10\x3d\x8b\x5f\x31\x6b\x14\x0b\xab\x70\xa8\x3d\x07\xc8\xa5\x49\x9b\x16\xdd\x26\x86\xd7\xed\x1e\x8a\xa2\xa0\xa5\xb1\x44\x58\x22\x15\x92\x76\xe2\x25\xf4\xef\x8b\xa1\x24\x5b\x89\x1d\xc7\x87\x00\x31\xf9\x38\xf3\xe6\xbd\xe1\x88\x49\x02\xd7\x3a\x43\xc8\x51\xa1\x11\x0e\x33\x58\x6c\xc1\x99\xb5\xb5\x1c\x6e\xee\xe1\xee\x7e\x0e\x1f\x6f\xbe\xcc\x39\x4b\x12\x98\xa1\x59\x2b\x25\x55\xde\x02\xe0\x51\x96\x25\xe8\x0d\x9a\x47\x23\x1d\x82\x2b\xa4\x85\xa5\x2c\x31\x80\x7f\xa0\xb1\x52\xab\x4b\xf0\x9e\x77\xff\x37\xcd\x60\x03\x6e\x84\xc3\xe1\x2e\xfd\x6e\x1a\xc6\x6a\x91\xae\x44\x1e\x18\x31\xc2\xcf\xfb\xb0\x50\x1b\xbd\x91\x19\x5a\xb0\x68\x36\x68\x2e\xac\xcc\x10\x16\x52\x65\x52\xe5\x16\x96\xda\x80\x2b\x10\xf2\xd9\xf4\x1a\x9c\x11\xca\xd6\xda\xb8\xc0\xe5\x8b\x83\xb5\x93\xa5\xfc\x0f\x6d\x80\xec\x76\x93\xdc\xd4\x29\xff\x27\x84\xe3\x8c\xc9\x8a\x16\x61\xcc\xa2\x91\x42\x97\x14\xce\xd5\x23\x16\x8d\x52\xad\x1c\x3e\xb9\x11\x63\xd1\x28\xd7\x3a\x2f\x91\xe7\xba\x14\x2a\xe7\xda\xe4\x21\x44\x52\xa1\x13\x99\x70\x82\x30\xb4\xb0\xcb\x00\xa3\x5c\xba\x62\xbd\xe0\xa9\xae\x92\x5c\x5f\xac\xa4\x4b\xe8\xef\x39\x05\xca\x32\xc0\xdd\xe8\x3b\x7c\xb4\xe4\x4b\x62\x5d\x96\xd4\xab\x3c\x21\x0e\x46\xa4\x81\x44\x2f\x0a\xf1\x96\x29\xb2\xa8\x5e\xc0\xc8\x7b\x3e\xfd\xf0\x25\x14\x30\x15\xae\x80\x8b\xa6\x19\xb1\x38\x28\xf8\x4d\xac\xf0\x76\x36\xbd\x6e\xeb\x84\x4a\xac\xd0\x82\x00\x8b\x0e\xf4\x12\x50\x65\xb5\x96\xca\x59\x10\x1b\x21\x4b\xb1\x28\x11\x04\xed\x07\x21\xbd\xe7\x5d\x1a\x7e\x27\x2a\x6c\x9a\x5e\xac\xe5\x5a\xa5\x2f\x22\x8f\xf7\xa1\x3e\xf6\xff\x4d\x40\xd7\x4e\x6a\x65\x81\x73\xfe\x4c\x99\x4e\xf6\xfb\xb0\x1d\x43\xbd\xe0\xaf\xe4\x02\xcf\x22\x3b\xc0\x5a\xb8\xbc\x82\x9f\xbf\x5e\x0f\xe6\x59\x14\x1d\xdb\xfd\x80\x4b\x6d\x70\xdc\x7b\x35\xd7\xd7\xad\xb1\xf1\x84\x45\xcd\xcb\x1c\x57\x20\xea\x1a\x55\x36\x7e\xb6\xbc\x2b\x87\x73\x1e\xb3\xc8\xa0\x5b\x1b\x05\x7f\x52\xb6\x36\x87\x0f\xf6\x78\x0f\x73\xfd\xb7\x7e\x44\x03\xcf\x4a\x82\xa6\x61\x91\xf7\x46\xa8\x1c\xe1\x9d\xa4\x42\x76\xfb\xdf\xd0\x15\x3a\xb3\x84\x88\xbc\xef\x8f\xbf\x93\x9d\x16\x97\xf0\xbc\xa4\x3b\x7c\xec\x54\x67\x51\x14\xed\x94\xe7\xde\xef\x8e\xf4\x26\x4c\x08\x71\x83\xa9\xce\x82\x59\x03\xc4\x0c\x1f\xd6\x68\x5b\xc0\x47\x75\x14\x60\x6b\xad\x2c\x06\xc4\x33\x25\x38\xe7\xb4\x48\xda\x79\x7f\x41\x5d\x44\xcc\x1b\xd6\x84\x96\xdb\x0b\x02\xb2\xaa\x4b\xac\x90\xba\x82\xee\x9e\xf7\xb7\x3a\x48\x71\xdc\x6b\xa9\x1c\x9a\xa5\x48\x91\xb9\x6d\x8d\xc3\x38\xd6\x99\x75\xea\xc0\xb3\xb7\xf5\x3b\x22\x1f\xc0\x0b\xfd\x3e\x0b\x95\x95\x68\xd8\x9e\x7c\xcb\xbc\x0b\x13\xc6\xc9\x20\xbb\xd3\xfb\x42\xce\xaf\xe1\x4d\xaa\xe1\x16\x8d\x2d\xbc\xdf\xa7\x8a\xf7\xe1\x77\xec\xc7\xa9\x7b\x82\x6e\x0c\xf1\xae\x6b\x27\x60\xf0\x01\xde\x87\x7b\xb3\xc7\x77\x8e\xce\xb7\x75\x4f\x2a\x86\xf1\x21\xa8\x75\x75\x80\x9a\x00\x1a\xa3\x29\x39\x8b\x7e\x53\xe8\x3a\xac\x10\x6d\xea\xa9\x03\x3d\xdb\x2b\x45\xdd\x42\xdc\x02\x97\x98\x45\x72\x19\x0e\xfd\x71\x05\x4a\x96\x14\xaa\xbf\x21\x4a\x96\x21\x5e\xb8\x68\xdd\x9a\xc1\x9a\x9f\x43\x2d\x9e\xd0\x71\xd6\x30\xef\x5b\xa3\xc8\xa6\x4e\xea\xb6\xab\xdf\xd6\x39\x49\xe0\xd4\x05\x00\x49\x03\xef\xc5\x67\xa1\x3d\xd0\x21\x3e\x91\x51\xae\x10\x8e\x6c\xd8\xa0\xa1\x71\x19\x1a\xbd\x1d\x92\x87\xfd\x66\xba\xc8\x4e\x83\x80\xb5\x45\x73\x91\xe9\x4a\x48\x75\x0a\xcc\x61\x6a\x64\x25\x8c\x2c\xb7\x74\x64\xb9\x2e\x41\xaa\x30\xa9\x07\x33\xf7\x54\x1d\xe3\xdf\x87\x5d\x42\xb5\xcc\xf0\x61\xdf\x95\x9e\x5a\x62\xf0\x6b\x68\x3d\xb5\xd4\xe5\x55\x7f\xe6\x98\x3d\x07\xed\x35\xf0\xf3\xe1\x84\x53\xed\x78\x39\xcb\xa9\x93\x93\xe8\xa8\x55\xed\x89\x1e\xf2\x9a\x57\x6f\xbb\xd0\xa5\x08\x9e\x9d\x70\xb6\x2e\xb7\x67\x59\x75\xb2\x90\x63\x5e\xed\x18\x9c\x69\x96\xad\x49\xc5\xfe\xd4\x79\xb7\x69\xe0\x97\xad\x8f\x19\xf6\x19\xcb\x1a\x8d\x65\x6d\x0d\x07\x5f\xcb\xe3\xb3\xa8\xca\x76\x48\xfe\xed\x26\x7e\x09\x20\xba\x34\x51\x57\x13\xd8\x04\xca\xa1\x09\xaa\x2c\xcc\x08\xb9\x84\xcd\x70\x66\xb4\x0f\x1c\x84\x15\x6e\x83\xdb\x59\x46\xcf\x52\xed\x0a\x92\xb8\xcf\x42\x03\xba\x12\x0e\xc6\xab\x18\x1e\x0b\x99\x16\x01\x5a\x96\x50\x92\x5d\x5d\x14\xa1\xb2\xf0\xd1\xa1\x97\x1c\xbf\x16\x4a\x2b\x99\x8a\xf2\x33\x8a\x0c\xcd\x57\xdc\xd2\xf3\xc7\x75\x89\xac\x6e\x5b\x46\x3a\x48\x85\x82\x05\xf6\x21\xd2\x14\xad\xc5\x8c\x72\xa3\x74\x05\x9a\x2e\x33\xed\x93\x14\x57\xbb\x5a\xff\x95\xae\xf8\x21\xca\x35\xb6\x23\x91\x6a\xfd\xf9\xd7\xaf\xf8\x4d\xe0\x2b\xec\xc6\xab\x78\x1f\x21\x7c\x5b\x19\xc0\xc9\x40\xfd\x23\xb1\xbf\xa6\xdf\x4d\xf9\x15\xb7\x64\x0e\xbf\x45\x37\x1e\x5d\xd6\xc2\x15\xa3\x38\x84\x3c\x37\xd4\xbc\xbf\x69\x21\xd2\xe8\x76\x3a\xbb\x9e\xce\xee\xe7\xf7\x1f\xbe\x7f\x1a\xed\x7b\x29\x75\x4f\xac\x61\xff\x07\x00\x00\xff\xff\x2f\x63\xe5\xed\x4d\x0c\x00\x00")

func genTransport_grpcGotemplateBytes() ([]byte, error) {
	return bindataRead(
		_genTransport_grpcGotemplate,
		"gen/transport_grpc.gotemplate",
	)
}

func genTransport_grpcGotemplate() (*asset, error) {
	bytes, err := genTransport_grpcGotemplateBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "gen/transport_grpc.gotemplate", size: 3149, mode: os.FileMode(0644), modTime: time.Unix(1464111000, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x19, 0xd9, 0x54, 0x92, 0xdf, 0x4f, 0xf, 0x60, 0x9f, 0xe1, 0x20, 0x79, 0xaa, 0x1b, 0xc2, 0xcc, 0x4e, 0x8e, 0x29, 0xea, 0xdf, 0xde, 0x8d, 0x2b, 0x70, 0x10, 0x12, 0xc5, 0x2b, 0xd2, 0x18, 0x40}}
	return a, nil
}

var _genTransport_httpGotemplate = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2c\xcb\xb1\x0d\xc2\x30\x10\x05\xd0\x3e\x53\x5c\x1d\x09\xdf\x1a\x29\x91\xe2\x05\xac\xf0\x31\x08\x93\xb3\xce\x1f\x1a\xeb\x76\xa7\x61\x80\x37\xa7\xae\xb2\x03\x52\xed\x42\xff\x8c\xa1\x15\x67\xb5\xd7\x93\xfa\x20\x3b\xbd\x9c\xa3\x9b\x53\x89\x77\x6f\x85\x18\xa9\x9a\xdc\xcd\xe5\xb0\x1b\x64\xd5\x88\x65\xce\xa3\xb4\x26\x69\xcb\xf9\xba\xa1\x75\x78\xda\xe1\x5f\x78\xfe\x1b\x49\x11\xcb\x2f\x00\x00\xff\xff\xdd\x3a\x4a\x8f\x6a\x00\x00\x00")

func genTransport_httpGotemplateBytes() ([]byte, error) {
	return bindataRead(
		_genTransport_httpGotemplate,
		"gen/transport_http.gotemplate",
	)
}

func genTransport_httpGotemplate() (*asset, error) {
	bytes, err := genTransport_httpGotemplateBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "gen/transport_http.gotemplate", size: 106, mode: os.FileMode(0644), modTime: time.Unix(1464111000, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x52, 0x57, 0x56, 0xc6, 0xb4, 0xe5, 0x1f, 0xf4, 0x1d, 0xa5, 0xda, 0x23, 0xea, 0x8f, 0xfb, 0xff, 0xae, 0x4b, 0x12, 0xe4, 0xf6, 0xbf, 0x11, 0xa6, 0x4, 0x83, 0x53, 0xfd, 0xbf, 0xce, 0x4a, 0x47}}
	return a, nil
}

var _serviceGotemplate = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x04\xc0\xc1\x0d\x80\x20\x0c\x05\xd0\xbb\x53\xf4\x4c\xa2\x1d\xc6\x09\x88\x7c\xab\x11\xa9\x69\xeb\x89\xb0\x3b\xaf\x77\x4e\xb4\x03\x24\xba\x86\xfd\xee\x2c\x68\xa2\xcf\x1d\x7c\xe5\x56\x2a\x8c\x03\xef\x57\x73\xc0\x37\x51\x3a\xd5\xe8\xd0\x02\x4a\x3c\xc6\x32\x03\x00\x00\xff\xff\xd6\x21\xab\x2e\x3e\x00\x00\x00")

func serviceGotemplateBytes() ([]byte, error) {
	return bindataRead(
		_serviceGotemplate,
		"service.gotemplate",
	)
}

func serviceGotemplate() (*asset, error) {
	bytes, err := serviceGotemplateBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "service.gotemplate", size: 62, mode: os.FileMode(0644), modTime: time.Unix(1464111000, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x1e, 0xcb, 0xd5, 0x72, 0x80, 0xc6, 0xf9, 0x82, 0x4b, 0xe0, 0x8b, 0x90, 0xb8, 0x9b, 0xbc, 0x5d, 0x8d, 0x12, 0xd4, 0x8e, 0x54, 0xf6, 0x72, 0xcb, 0xef, 0xf5, 0x12, 0xd0, 0xe1, 0xb8, 0x41, 0xc8}}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// AssetString returns the asset contents as a string (instead of a []byte).
func AssetString(name string) (string, error) {
	data, err := Asset(name)
	return string(data), err
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

// MustAssetString is like AssetString but panics when Asset would return an
// error. It simplifies safe initialization of global variables.
func MustAssetString(name string) string {
	return string(MustAsset(name))
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetDigest returns the digest of the file with the given name. It returns an
// error if the asset could not be found or the digest could not be loaded.
func AssetDigest(name string) ([sha256.Size]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s can't read by error: %v", name, err)
		}
		return a.digest, nil
	}
	return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s not found", name)
}

// Digests returns a map of all known files and their checksums.
func Digests() (map[string][sha256.Size]byte, error) {
	mp := make(map[string][sha256.Size]byte, len(_bindata))
	for name := range _bindata {
		a, err := _bindata[name]()
		if err != nil {
			return nil, err
		}
		mp[name] = a.digest
	}
	return mp, nil
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
	"gen/client/grpc/client.gotemplate": genClientGrpcClientGotemplate,
	"gen/client/http/client.gotemplate": genClientHttpClientGotemplate,
	"gen/endpoints.gotemplate":          genEndpointsGotemplate,
	"gen/transport_grpc.gotemplate":     genTransport_grpcGotemplate,
	"gen/transport_http.gotemplate":     genTransport_httpGotemplate,
	"service.gotemplate":                serviceGotemplate,
}

// AssetDebug is true if the assets were built with the debug flag enabled.
const AssetDebug = false

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"},
// AssetDir("data/img") would return []string{"a.png", "b.png"},
// AssetDir("foo.txt") and AssetDir("notexist") would return an error, and
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		canonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(canonicalName, "/")
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
	"gen": &bintree{nil, map[string]*bintree{
		"client": &bintree{nil, map[string]*bintree{
			"grpc": &bintree{nil, map[string]*bintree{
				"client.gotemplate": &bintree{genClientGrpcClientGotemplate, map[string]*bintree{}},
			}},
			"http": &bintree{nil, map[string]*bintree{
				"client.gotemplate": &bintree{genClientHttpClientGotemplate, map[string]*bintree{}},
			}},
		}},
		"endpoints.gotemplate":      &bintree{genEndpointsGotemplate, map[string]*bintree{}},
		"transport_grpc.gotemplate": &bintree{genTransport_grpcGotemplate, map[string]*bintree{}},
		"transport_http.gotemplate": &bintree{genTransport_httpGotemplate, map[string]*bintree{}},
	}},
	"service.gotemplate": &bintree{serviceGotemplate, map[string]*bintree{}},
}}

// RestoreAsset restores an asset under the given directory.
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
	return os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
}

// RestoreAssets restores an asset under the given directory recursively.
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
	canonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(canonicalName, "/")...)...)
}
