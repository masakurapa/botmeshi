package model

// File ファイル情報の構造体
type File struct {
	// ファイルパス
	path string
	// ファイル名
	name string
	// ファイルデータ
	data string
}

// NewFile returns File instance
func NewFile(path, name, data string) *File {
	return &File{
		path: path,
		name: name,
		data: data,
	}
}

// GetPath returns path string
func (f *File) GetPath() string {
	return f.path
}

// GetName returns file name string
func (f *File) GetName() string {
	return f.name
}

// GetData returns file data string
func (f *File) GetData() string {
	return f.data
}
