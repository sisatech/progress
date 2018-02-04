package progress

import (
	pb "gopkg.in/cheggaaa/pb.v1"
)

// UpdatePB ..
func UpdatePB(b **pb.ProgressBar, pt ProgressTracker) {

	status := pt.Status()

	if *b == nil {
		(*b) = pb.New(int(status.Total))
		(*b).SetMaxWidth(80)
		(*b).Start()
	}

	(*b).Prefix(status.Stage)
	(*b).Set(int(status.Progress))
}
