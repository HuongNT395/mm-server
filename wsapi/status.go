// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

package wsapi

import (
	"github.com/mattermost/mattermost-server/v5/mlog"
	"github.com/mattermost/mattermost-server/v5/model"
)

func (api *API) InitStatus() {
	api.Router.Handle("get_statuses", api.ApiWebSocketHandler(api.getStatuses))
	api.Router.Handle("get_statuses_by_ids", api.ApiWebSocketHandler(api.getStatusesByIds))
}

func (api *API) getStatuses(req *model.WebSocketRequest) (map[string]interface{}, *model.AppError) {
	statusMap := api.App.GetAllStatuses()
	return model.StatusMapToInterfaceMap(statusMap), nil
}

func (api *API) getStatusesByIds(req *model.WebSocketRequest) (map[string]interface{}, *model.AppError) {
	var userIds []string
	if userIds = model.ArrayFromInterface(req.Data["user_ids"]); len(userIds) == 0 {
		mlog.Error(model.StringInterfaceToJson(req.Data))
		mlog.Error("da vao dc getStatusesByIds 25")
		return nil, NewInvalidWebSocketParamError(req.Action, "user_ids")
	}

	statusMap, err := api.App.GetStatusesByIds(userIds)
	if err != nil {
		return nil, err
	}

	return statusMap, nil
}
