package progress

import (
	pb "gopkg.in/cheggaaa/pb.v2"
)

// UpdatePB ..
func UpdatePB(b **pb.ProgressBar, pt ProgressTracker) {

	status := pt.Status()

	if *b == nil {
		(*b) = pb.New(int(status.Total))
		(*b).SetWidth(80)
		(*b).Start()
	}

	// (*b).Prefix(status.Stage)
	(*b).SetCurrent(int64(status.Progress))
}
