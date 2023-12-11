package sstable

// SSTable the basic structure of SST which is the disk storage unit in LSM-true
type SSTable struct {
}

func NewSST(opt *Options) *SSTable {
	sst := new(SSTable)
	return sst
}
