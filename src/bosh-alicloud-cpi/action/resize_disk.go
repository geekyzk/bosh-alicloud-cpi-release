/*
 * Copyright (C) 2017-2017 Alibaba Group Holding Limited
 */
package action

import (
	"github.com/cppforlife/bosh-cpi-go/apiv1"
	"bosh-alicloud-cpi/alicloud"
)

type ResizeDiskMethod struct {
	CallContext
	disks alicloud.DiskManager
}

func NewResizeDiskMethod(cc CallContext, disks alicloud.DiskManager) ResizeDiskMethod {
	return ResizeDiskMethod{cc, disks}
}

func (a HasDiskMethod) ResizeDisk(diskCID apiv1.DiskCID, size int) (error) {
	sizeGB := ConvertToGB(float64(size))
	err := a.disks.ResizeDisk(diskCID.AsString(), sizeGB)
	if err != nil {
		return a.WrapErrorf(err, "ResizeDisk %s to %dMB failed", diskCID.AsString(), size)
	}
	return nil
}
