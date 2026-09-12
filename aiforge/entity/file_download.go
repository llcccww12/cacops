package entity

type FileDownloadInfo struct {
	Readers        []FileReader
	ResultType     FileType
	ResultFileName string
	DownloadUrl    string
}

func (f *FileDownloadInfo) IsEmpty() bool {
	return (f.Readers == nil || len(f.Readers) == 0) && f.DownloadUrl == ""
}

func (f *FileDownloadInfo) Close() {
	if f.Readers != nil && len(f.Readers) > 0 {
		for _, r := range f.Readers {
			if r.Reader != nil {
				r.Reader.Close()
			}
		}
	}
}
