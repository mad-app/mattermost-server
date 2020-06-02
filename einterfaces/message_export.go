// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

package einterfaces

import (
	"context"

	"github.com/mad-app/mattermost-server/v5/model"
)

type MessageExportInterface interface {
	StartSynchronizeJob(ctx context.Context, exportFromTimestamp int64) (*model.Job, *model.AppError)
	RunExport(format string, since int64) *model.AppError
}
