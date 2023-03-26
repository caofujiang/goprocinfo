package linux

type Disk struct {
	All        uint64 `json:"all"`
	Used       uint64 `json:"used"`
	Free       uint64 `json:"free"`
	FreeInodes uint64 `json:"freeInodes"`
}

func ReadDisk(path string) (*Disk, error) {
	disk := Disk{}
	//if runtime.GOOS != "windows" {
	//	fs := syscall.Statfs_t{}
	//	err := syscall.Statfs(path, &fs)
	//	if err != nil {
	//		return nil, err
	//	}
	//	disk.All = fs.Blocks * uint64(fs.Bsize)
	//	disk.Free = fs.Bfree * uint64(fs.Bsize)
	//	disk.Used = disk.All - disk.Free
	//	disk.FreeInodes = fs.Ffree
	//}
	return &disk, nil
}
