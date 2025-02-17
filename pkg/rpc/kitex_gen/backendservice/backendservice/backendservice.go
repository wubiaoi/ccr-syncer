// Code generated by Kitex v0.8.0. DO NOT EDIT.

package backendservice

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	agentservice "github.com/selectdb/ccr_syncer/pkg/rpc/kitex_gen/agentservice"
	backendservice "github.com/selectdb/ccr_syncer/pkg/rpc/kitex_gen/backendservice"
	dorisexternalservice "github.com/selectdb/ccr_syncer/pkg/rpc/kitex_gen/dorisexternalservice"
	palointernalservice "github.com/selectdb/ccr_syncer/pkg/rpc/kitex_gen/palointernalservice"
	status "github.com/selectdb/ccr_syncer/pkg/rpc/kitex_gen/status"
	types "github.com/selectdb/ccr_syncer/pkg/rpc/kitex_gen/types"
)

func serviceInfo() *kitex.ServiceInfo {
	return backendServiceServiceInfo
}

var backendServiceServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "BackendService"
	handlerType := (*backendservice.BackendService)(nil)
	methods := map[string]kitex.MethodInfo{
		"exec_plan_fragment":           kitex.NewMethodInfo(execPlanFragmentHandler, newBackendServiceExecPlanFragmentArgs, newBackendServiceExecPlanFragmentResult, false),
		"cancel_plan_fragment":         kitex.NewMethodInfo(cancelPlanFragmentHandler, newBackendServiceCancelPlanFragmentArgs, newBackendServiceCancelPlanFragmentResult, false),
		"transmit_data":                kitex.NewMethodInfo(transmitDataHandler, newBackendServiceTransmitDataArgs, newBackendServiceTransmitDataResult, false),
		"submit_tasks":                 kitex.NewMethodInfo(submitTasksHandler, newBackendServiceSubmitTasksArgs, newBackendServiceSubmitTasksResult, false),
		"make_snapshot":                kitex.NewMethodInfo(makeSnapshotHandler, newBackendServiceMakeSnapshotArgs, newBackendServiceMakeSnapshotResult, false),
		"release_snapshot":             kitex.NewMethodInfo(releaseSnapshotHandler, newBackendServiceReleaseSnapshotArgs, newBackendServiceReleaseSnapshotResult, false),
		"publish_cluster_state":        kitex.NewMethodInfo(publishClusterStateHandler, newBackendServicePublishClusterStateArgs, newBackendServicePublishClusterStateResult, false),
		"submit_export_task":           kitex.NewMethodInfo(submitExportTaskHandler, newBackendServiceSubmitExportTaskArgs, newBackendServiceSubmitExportTaskResult, false),
		"get_export_status":            kitex.NewMethodInfo(getExportStatusHandler, newBackendServiceGetExportStatusArgs, newBackendServiceGetExportStatusResult, false),
		"erase_export_task":            kitex.NewMethodInfo(eraseExportTaskHandler, newBackendServiceEraseExportTaskArgs, newBackendServiceEraseExportTaskResult, false),
		"get_tablet_stat":              kitex.NewMethodInfo(getTabletStatHandler, newBackendServiceGetTabletStatArgs, newBackendServiceGetTabletStatResult, false),
		"get_trash_used_capacity":      kitex.NewMethodInfo(getTrashUsedCapacityHandler, newBackendServiceGetTrashUsedCapacityArgs, newBackendServiceGetTrashUsedCapacityResult, false),
		"get_disk_trash_used_capacity": kitex.NewMethodInfo(getDiskTrashUsedCapacityHandler, newBackendServiceGetDiskTrashUsedCapacityArgs, newBackendServiceGetDiskTrashUsedCapacityResult, false),
		"submit_routine_load_task":     kitex.NewMethodInfo(submitRoutineLoadTaskHandler, newBackendServiceSubmitRoutineLoadTaskArgs, newBackendServiceSubmitRoutineLoadTaskResult, false),
		"open_scanner":                 kitex.NewMethodInfo(openScannerHandler, newBackendServiceOpenScannerArgs, newBackendServiceOpenScannerResult, false),
		"get_next":                     kitex.NewMethodInfo(getNextHandler, newBackendServiceGetNextArgs, newBackendServiceGetNextResult, false),
		"close_scanner":                kitex.NewMethodInfo(closeScannerHandler, newBackendServiceCloseScannerArgs, newBackendServiceCloseScannerResult, false),
		"get_stream_load_record":       kitex.NewMethodInfo(getStreamLoadRecordHandler, newBackendServiceGetStreamLoadRecordArgs, newBackendServiceGetStreamLoadRecordResult, false),
		"check_storage_format":         kitex.NewMethodInfo(checkStorageFormatHandler, newBackendServiceCheckStorageFormatArgs, newBackendServiceCheckStorageFormatResult, false),
		"warm_up_cache_async":          kitex.NewMethodInfo(warmUpCacheAsyncHandler, newBackendServiceWarmUpCacheAsyncArgs, newBackendServiceWarmUpCacheAsyncResult, false),
		"check_warm_up_cache_async":    kitex.NewMethodInfo(checkWarmUpCacheAsyncHandler, newBackendServiceCheckWarmUpCacheAsyncArgs, newBackendServiceCheckWarmUpCacheAsyncResult, false),
		"sync_load_for_tablets":        kitex.NewMethodInfo(syncLoadForTabletsHandler, newBackendServiceSyncLoadForTabletsArgs, newBackendServiceSyncLoadForTabletsResult, false),
		"get_top_n_hot_partitions":     kitex.NewMethodInfo(getTopNHotPartitionsHandler, newBackendServiceGetTopNHotPartitionsArgs, newBackendServiceGetTopNHotPartitionsResult, false),
		"warm_up_tablets":              kitex.NewMethodInfo(warmUpTabletsHandler, newBackendServiceWarmUpTabletsArgs, newBackendServiceWarmUpTabletsResult, false),
		"ingest_binlog":                kitex.NewMethodInfo(ingestBinlogHandler, newBackendServiceIngestBinlogArgs, newBackendServiceIngestBinlogResult, false),
		"query_ingest_binlog":          kitex.NewMethodInfo(queryIngestBinlogHandler, newBackendServiceQueryIngestBinlogArgs, newBackendServiceQueryIngestBinlogResult, false),
		"publish_topic_info":           kitex.NewMethodInfo(publishTopicInfoHandler, newBackendServicePublishTopicInfoArgs, newBackendServicePublishTopicInfoResult, false),
		"get_realtime_exec_status":     kitex.NewMethodInfo(getRealtimeExecStatusHandler, newBackendServiceGetRealtimeExecStatusArgs, newBackendServiceGetRealtimeExecStatusResult, false),
	}
	extra := map[string]interface{}{
		"PackageName":     "backendservice",
		"ServiceFilePath": `thrift/BackendService.thrift`,
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Thrift,
		KiteXGenVersion: "v0.8.0",
		Extra:           extra,
	}
	return svcInfo
}

func execPlanFragmentHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*backendservice.BackendServiceExecPlanFragmentArgs)
	realResult := result.(*backendservice.BackendServiceExecPlanFragmentResult)
	success, err := handler.(backendservice.BackendService).ExecPlanFragment(ctx, realArg.Params)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newBackendServiceExecPlanFragmentArgs() interface{} {
	return backendservice.NewBackendServiceExecPlanFragmentArgs()
}

func newBackendServiceExecPlanFragmentResult() interface{} {
	return backendservice.NewBackendServiceExecPlanFragmentResult()
}

func cancelPlanFragmentHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*backendservice.BackendServiceCancelPlanFragmentArgs)
	realResult := result.(*backendservice.BackendServiceCancelPlanFragmentResult)
	success, err := handler.(backendservice.BackendService).CancelPlanFragment(ctx, realArg.Params)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newBackendServiceCancelPlanFragmentArgs() interface{} {
	return backendservice.NewBackendServiceCancelPlanFragmentArgs()
}

func newBackendServiceCancelPlanFragmentResult() interface{} {
	return backendservice.NewBackendServiceCancelPlanFragmentResult()
}

func transmitDataHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*backendservice.BackendServiceTransmitDataArgs)
	realResult := result.(*backendservice.BackendServiceTransmitDataResult)
	success, err := handler.(backendservice.BackendService).TransmitData(ctx, realArg.Params)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newBackendServiceTransmitDataArgs() interface{} {
	return backendservice.NewBackendServiceTransmitDataArgs()
}

func newBackendServiceTransmitDataResult() interface{} {
	return backendservice.NewBackendServiceTransmitDataResult()
}

func submitTasksHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*backendservice.BackendServiceSubmitTasksArgs)
	realResult := result.(*backendservice.BackendServiceSubmitTasksResult)
	success, err := handler.(backendservice.BackendService).SubmitTasks(ctx, realArg.Tasks)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newBackendServiceSubmitTasksArgs() interface{} {
	return backendservice.NewBackendServiceSubmitTasksArgs()
}

func newBackendServiceSubmitTasksResult() interface{} {
	return backendservice.NewBackendServiceSubmitTasksResult()
}

func makeSnapshotHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*backendservice.BackendServiceMakeSnapshotArgs)
	realResult := result.(*backendservice.BackendServiceMakeSnapshotResult)
	success, err := handler.(backendservice.BackendService).MakeSnapshot(ctx, realArg.SnapshotRequest)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newBackendServiceMakeSnapshotArgs() interface{} {
	return backendservice.NewBackendServiceMakeSnapshotArgs()
}

func newBackendServiceMakeSnapshotResult() interface{} {
	return backendservice.NewBackendServiceMakeSnapshotResult()
}

func releaseSnapshotHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*backendservice.BackendServiceReleaseSnapshotArgs)
	realResult := result.(*backendservice.BackendServiceReleaseSnapshotResult)
	success, err := handler.(backendservice.BackendService).ReleaseSnapshot(ctx, realArg.SnapshotPath)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newBackendServiceReleaseSnapshotArgs() interface{} {
	return backendservice.NewBackendServiceReleaseSnapshotArgs()
}

func newBackendServiceReleaseSnapshotResult() interface{} {
	return backendservice.NewBackendServiceReleaseSnapshotResult()
}

func publishClusterStateHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*backendservice.BackendServicePublishClusterStateArgs)
	realResult := result.(*backendservice.BackendServicePublishClusterStateResult)
	success, err := handler.(backendservice.BackendService).PublishClusterState(ctx, realArg.Request)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newBackendServicePublishClusterStateArgs() interface{} {
	return backendservice.NewBackendServicePublishClusterStateArgs()
}

func newBackendServicePublishClusterStateResult() interface{} {
	return backendservice.NewBackendServicePublishClusterStateResult()
}

func submitExportTaskHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*backendservice.BackendServiceSubmitExportTaskArgs)
	realResult := result.(*backendservice.BackendServiceSubmitExportTaskResult)
	success, err := handler.(backendservice.BackendService).SubmitExportTask(ctx, realArg.Request)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newBackendServiceSubmitExportTaskArgs() interface{} {
	return backendservice.NewBackendServiceSubmitExportTaskArgs()
}

func newBackendServiceSubmitExportTaskResult() interface{} {
	return backendservice.NewBackendServiceSubmitExportTaskResult()
}

func getExportStatusHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*backendservice.BackendServiceGetExportStatusArgs)
	realResult := result.(*backendservice.BackendServiceGetExportStatusResult)
	success, err := handler.(backendservice.BackendService).GetExportStatus(ctx, realArg.TaskId)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newBackendServiceGetExportStatusArgs() interface{} {
	return backendservice.NewBackendServiceGetExportStatusArgs()
}

func newBackendServiceGetExportStatusResult() interface{} {
	return backendservice.NewBackendServiceGetExportStatusResult()
}

func eraseExportTaskHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*backendservice.BackendServiceEraseExportTaskArgs)
	realResult := result.(*backendservice.BackendServiceEraseExportTaskResult)
	success, err := handler.(backendservice.BackendService).EraseExportTask(ctx, realArg.TaskId)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newBackendServiceEraseExportTaskArgs() interface{} {
	return backendservice.NewBackendServiceEraseExportTaskArgs()
}

func newBackendServiceEraseExportTaskResult() interface{} {
	return backendservice.NewBackendServiceEraseExportTaskResult()
}

func getTabletStatHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {

	realResult := result.(*backendservice.BackendServiceGetTabletStatResult)
	success, err := handler.(backendservice.BackendService).GetTabletStat(ctx)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newBackendServiceGetTabletStatArgs() interface{} {
	return backendservice.NewBackendServiceGetTabletStatArgs()
}

func newBackendServiceGetTabletStatResult() interface{} {
	return backendservice.NewBackendServiceGetTabletStatResult()
}

func getTrashUsedCapacityHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {

	realResult := result.(*backendservice.BackendServiceGetTrashUsedCapacityResult)
	success, err := handler.(backendservice.BackendService).GetTrashUsedCapacity(ctx)
	if err != nil {
		return err
	}
	realResult.Success = &success
	return nil
}
func newBackendServiceGetTrashUsedCapacityArgs() interface{} {
	return backendservice.NewBackendServiceGetTrashUsedCapacityArgs()
}

func newBackendServiceGetTrashUsedCapacityResult() interface{} {
	return backendservice.NewBackendServiceGetTrashUsedCapacityResult()
}

func getDiskTrashUsedCapacityHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {

	realResult := result.(*backendservice.BackendServiceGetDiskTrashUsedCapacityResult)
	success, err := handler.(backendservice.BackendService).GetDiskTrashUsedCapacity(ctx)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newBackendServiceGetDiskTrashUsedCapacityArgs() interface{} {
	return backendservice.NewBackendServiceGetDiskTrashUsedCapacityArgs()
}

func newBackendServiceGetDiskTrashUsedCapacityResult() interface{} {
	return backendservice.NewBackendServiceGetDiskTrashUsedCapacityResult()
}

func submitRoutineLoadTaskHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*backendservice.BackendServiceSubmitRoutineLoadTaskArgs)
	realResult := result.(*backendservice.BackendServiceSubmitRoutineLoadTaskResult)
	success, err := handler.(backendservice.BackendService).SubmitRoutineLoadTask(ctx, realArg.Tasks)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newBackendServiceSubmitRoutineLoadTaskArgs() interface{} {
	return backendservice.NewBackendServiceSubmitRoutineLoadTaskArgs()
}

func newBackendServiceSubmitRoutineLoadTaskResult() interface{} {
	return backendservice.NewBackendServiceSubmitRoutineLoadTaskResult()
}

func openScannerHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*backendservice.BackendServiceOpenScannerArgs)
	realResult := result.(*backendservice.BackendServiceOpenScannerResult)
	success, err := handler.(backendservice.BackendService).OpenScanner(ctx, realArg.Params)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newBackendServiceOpenScannerArgs() interface{} {
	return backendservice.NewBackendServiceOpenScannerArgs()
}

func newBackendServiceOpenScannerResult() interface{} {
	return backendservice.NewBackendServiceOpenScannerResult()
}

func getNextHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*backendservice.BackendServiceGetNextArgs)
	realResult := result.(*backendservice.BackendServiceGetNextResult)
	success, err := handler.(backendservice.BackendService).GetNext(ctx, realArg.Params)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newBackendServiceGetNextArgs() interface{} {
	return backendservice.NewBackendServiceGetNextArgs()
}

func newBackendServiceGetNextResult() interface{} {
	return backendservice.NewBackendServiceGetNextResult()
}

func closeScannerHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*backendservice.BackendServiceCloseScannerArgs)
	realResult := result.(*backendservice.BackendServiceCloseScannerResult)
	success, err := handler.(backendservice.BackendService).CloseScanner(ctx, realArg.Params)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newBackendServiceCloseScannerArgs() interface{} {
	return backendservice.NewBackendServiceCloseScannerArgs()
}

func newBackendServiceCloseScannerResult() interface{} {
	return backendservice.NewBackendServiceCloseScannerResult()
}

func getStreamLoadRecordHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*backendservice.BackendServiceGetStreamLoadRecordArgs)
	realResult := result.(*backendservice.BackendServiceGetStreamLoadRecordResult)
	success, err := handler.(backendservice.BackendService).GetStreamLoadRecord(ctx, realArg.LastStreamRecordTime)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newBackendServiceGetStreamLoadRecordArgs() interface{} {
	return backendservice.NewBackendServiceGetStreamLoadRecordArgs()
}

func newBackendServiceGetStreamLoadRecordResult() interface{} {
	return backendservice.NewBackendServiceGetStreamLoadRecordResult()
}

func checkStorageFormatHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {

	realResult := result.(*backendservice.BackendServiceCheckStorageFormatResult)
	success, err := handler.(backendservice.BackendService).CheckStorageFormat(ctx)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newBackendServiceCheckStorageFormatArgs() interface{} {
	return backendservice.NewBackendServiceCheckStorageFormatArgs()
}

func newBackendServiceCheckStorageFormatResult() interface{} {
	return backendservice.NewBackendServiceCheckStorageFormatResult()
}

func warmUpCacheAsyncHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*backendservice.BackendServiceWarmUpCacheAsyncArgs)
	realResult := result.(*backendservice.BackendServiceWarmUpCacheAsyncResult)
	success, err := handler.(backendservice.BackendService).WarmUpCacheAsync(ctx, realArg.Request)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newBackendServiceWarmUpCacheAsyncArgs() interface{} {
	return backendservice.NewBackendServiceWarmUpCacheAsyncArgs()
}

func newBackendServiceWarmUpCacheAsyncResult() interface{} {
	return backendservice.NewBackendServiceWarmUpCacheAsyncResult()
}

func checkWarmUpCacheAsyncHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*backendservice.BackendServiceCheckWarmUpCacheAsyncArgs)
	realResult := result.(*backendservice.BackendServiceCheckWarmUpCacheAsyncResult)
	success, err := handler.(backendservice.BackendService).CheckWarmUpCacheAsync(ctx, realArg.Request)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newBackendServiceCheckWarmUpCacheAsyncArgs() interface{} {
	return backendservice.NewBackendServiceCheckWarmUpCacheAsyncArgs()
}

func newBackendServiceCheckWarmUpCacheAsyncResult() interface{} {
	return backendservice.NewBackendServiceCheckWarmUpCacheAsyncResult()
}

func syncLoadForTabletsHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*backendservice.BackendServiceSyncLoadForTabletsArgs)
	realResult := result.(*backendservice.BackendServiceSyncLoadForTabletsResult)
	success, err := handler.(backendservice.BackendService).SyncLoadForTablets(ctx, realArg.Request)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newBackendServiceSyncLoadForTabletsArgs() interface{} {
	return backendservice.NewBackendServiceSyncLoadForTabletsArgs()
}

func newBackendServiceSyncLoadForTabletsResult() interface{} {
	return backendservice.NewBackendServiceSyncLoadForTabletsResult()
}

func getTopNHotPartitionsHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*backendservice.BackendServiceGetTopNHotPartitionsArgs)
	realResult := result.(*backendservice.BackendServiceGetTopNHotPartitionsResult)
	success, err := handler.(backendservice.BackendService).GetTopNHotPartitions(ctx, realArg.Request)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newBackendServiceGetTopNHotPartitionsArgs() interface{} {
	return backendservice.NewBackendServiceGetTopNHotPartitionsArgs()
}

func newBackendServiceGetTopNHotPartitionsResult() interface{} {
	return backendservice.NewBackendServiceGetTopNHotPartitionsResult()
}

func warmUpTabletsHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*backendservice.BackendServiceWarmUpTabletsArgs)
	realResult := result.(*backendservice.BackendServiceWarmUpTabletsResult)
	success, err := handler.(backendservice.BackendService).WarmUpTablets(ctx, realArg.Request)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newBackendServiceWarmUpTabletsArgs() interface{} {
	return backendservice.NewBackendServiceWarmUpTabletsArgs()
}

func newBackendServiceWarmUpTabletsResult() interface{} {
	return backendservice.NewBackendServiceWarmUpTabletsResult()
}

func ingestBinlogHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*backendservice.BackendServiceIngestBinlogArgs)
	realResult := result.(*backendservice.BackendServiceIngestBinlogResult)
	success, err := handler.(backendservice.BackendService).IngestBinlog(ctx, realArg.IngestBinlogRequest)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newBackendServiceIngestBinlogArgs() interface{} {
	return backendservice.NewBackendServiceIngestBinlogArgs()
}

func newBackendServiceIngestBinlogResult() interface{} {
	return backendservice.NewBackendServiceIngestBinlogResult()
}

func queryIngestBinlogHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*backendservice.BackendServiceQueryIngestBinlogArgs)
	realResult := result.(*backendservice.BackendServiceQueryIngestBinlogResult)
	success, err := handler.(backendservice.BackendService).QueryIngestBinlog(ctx, realArg.QueryIngestBinlogRequest)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newBackendServiceQueryIngestBinlogArgs() interface{} {
	return backendservice.NewBackendServiceQueryIngestBinlogArgs()
}

func newBackendServiceQueryIngestBinlogResult() interface{} {
	return backendservice.NewBackendServiceQueryIngestBinlogResult()
}

func publishTopicInfoHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*backendservice.BackendServicePublishTopicInfoArgs)
	realResult := result.(*backendservice.BackendServicePublishTopicInfoResult)
	success, err := handler.(backendservice.BackendService).PublishTopicInfo(ctx, realArg.TopicRequest)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newBackendServicePublishTopicInfoArgs() interface{} {
	return backendservice.NewBackendServicePublishTopicInfoArgs()
}

func newBackendServicePublishTopicInfoResult() interface{} {
	return backendservice.NewBackendServicePublishTopicInfoResult()
}

func getRealtimeExecStatusHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*backendservice.BackendServiceGetRealtimeExecStatusArgs)
	realResult := result.(*backendservice.BackendServiceGetRealtimeExecStatusResult)
	success, err := handler.(backendservice.BackendService).GetRealtimeExecStatus(ctx, realArg.Request)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newBackendServiceGetRealtimeExecStatusArgs() interface{} {
	return backendservice.NewBackendServiceGetRealtimeExecStatusArgs()
}

func newBackendServiceGetRealtimeExecStatusResult() interface{} {
	return backendservice.NewBackendServiceGetRealtimeExecStatusResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) ExecPlanFragment(ctx context.Context, params *palointernalservice.TExecPlanFragmentParams) (r *palointernalservice.TExecPlanFragmentResult_, err error) {
	var _args backendservice.BackendServiceExecPlanFragmentArgs
	_args.Params = params
	var _result backendservice.BackendServiceExecPlanFragmentResult
	if err = p.c.Call(ctx, "exec_plan_fragment", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) CancelPlanFragment(ctx context.Context, params *palointernalservice.TCancelPlanFragmentParams) (r *palointernalservice.TCancelPlanFragmentResult_, err error) {
	var _args backendservice.BackendServiceCancelPlanFragmentArgs
	_args.Params = params
	var _result backendservice.BackendServiceCancelPlanFragmentResult
	if err = p.c.Call(ctx, "cancel_plan_fragment", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) TransmitData(ctx context.Context, params *palointernalservice.TTransmitDataParams) (r *palointernalservice.TTransmitDataResult_, err error) {
	var _args backendservice.BackendServiceTransmitDataArgs
	_args.Params = params
	var _result backendservice.BackendServiceTransmitDataResult
	if err = p.c.Call(ctx, "transmit_data", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) SubmitTasks(ctx context.Context, tasks []*agentservice.TAgentTaskRequest) (r *agentservice.TAgentResult_, err error) {
	var _args backendservice.BackendServiceSubmitTasksArgs
	_args.Tasks = tasks
	var _result backendservice.BackendServiceSubmitTasksResult
	if err = p.c.Call(ctx, "submit_tasks", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) MakeSnapshot(ctx context.Context, snapshotRequest *agentservice.TSnapshotRequest) (r *agentservice.TAgentResult_, err error) {
	var _args backendservice.BackendServiceMakeSnapshotArgs
	_args.SnapshotRequest = snapshotRequest
	var _result backendservice.BackendServiceMakeSnapshotResult
	if err = p.c.Call(ctx, "make_snapshot", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) ReleaseSnapshot(ctx context.Context, snapshotPath string) (r *agentservice.TAgentResult_, err error) {
	var _args backendservice.BackendServiceReleaseSnapshotArgs
	_args.SnapshotPath = snapshotPath
	var _result backendservice.BackendServiceReleaseSnapshotResult
	if err = p.c.Call(ctx, "release_snapshot", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) PublishClusterState(ctx context.Context, request *agentservice.TAgentPublishRequest) (r *agentservice.TAgentResult_, err error) {
	var _args backendservice.BackendServicePublishClusterStateArgs
	_args.Request = request
	var _result backendservice.BackendServicePublishClusterStateResult
	if err = p.c.Call(ctx, "publish_cluster_state", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) SubmitExportTask(ctx context.Context, request *backendservice.TExportTaskRequest) (r *status.TStatus, err error) {
	var _args backendservice.BackendServiceSubmitExportTaskArgs
	_args.Request = request
	var _result backendservice.BackendServiceSubmitExportTaskResult
	if err = p.c.Call(ctx, "submit_export_task", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetExportStatus(ctx context.Context, taskId *types.TUniqueId) (r *palointernalservice.TExportStatusResult_, err error) {
	var _args backendservice.BackendServiceGetExportStatusArgs
	_args.TaskId = taskId
	var _result backendservice.BackendServiceGetExportStatusResult
	if err = p.c.Call(ctx, "get_export_status", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) EraseExportTask(ctx context.Context, taskId *types.TUniqueId) (r *status.TStatus, err error) {
	var _args backendservice.BackendServiceEraseExportTaskArgs
	_args.TaskId = taskId
	var _result backendservice.BackendServiceEraseExportTaskResult
	if err = p.c.Call(ctx, "erase_export_task", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetTabletStat(ctx context.Context) (r *backendservice.TTabletStatResult_, err error) {
	var _args backendservice.BackendServiceGetTabletStatArgs
	var _result backendservice.BackendServiceGetTabletStatResult
	if err = p.c.Call(ctx, "get_tablet_stat", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetTrashUsedCapacity(ctx context.Context) (r int64, err error) {
	var _args backendservice.BackendServiceGetTrashUsedCapacityArgs
	var _result backendservice.BackendServiceGetTrashUsedCapacityResult
	if err = p.c.Call(ctx, "get_trash_used_capacity", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetDiskTrashUsedCapacity(ctx context.Context) (r []*backendservice.TDiskTrashInfo, err error) {
	var _args backendservice.BackendServiceGetDiskTrashUsedCapacityArgs
	var _result backendservice.BackendServiceGetDiskTrashUsedCapacityResult
	if err = p.c.Call(ctx, "get_disk_trash_used_capacity", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) SubmitRoutineLoadTask(ctx context.Context, tasks []*backendservice.TRoutineLoadTask) (r *status.TStatus, err error) {
	var _args backendservice.BackendServiceSubmitRoutineLoadTaskArgs
	_args.Tasks = tasks
	var _result backendservice.BackendServiceSubmitRoutineLoadTaskResult
	if err = p.c.Call(ctx, "submit_routine_load_task", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) OpenScanner(ctx context.Context, params *dorisexternalservice.TScanOpenParams) (r *dorisexternalservice.TScanOpenResult_, err error) {
	var _args backendservice.BackendServiceOpenScannerArgs
	_args.Params = params
	var _result backendservice.BackendServiceOpenScannerResult
	if err = p.c.Call(ctx, "open_scanner", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetNext(ctx context.Context, params *dorisexternalservice.TScanNextBatchParams) (r *dorisexternalservice.TScanBatchResult_, err error) {
	var _args backendservice.BackendServiceGetNextArgs
	_args.Params = params
	var _result backendservice.BackendServiceGetNextResult
	if err = p.c.Call(ctx, "get_next", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) CloseScanner(ctx context.Context, params *dorisexternalservice.TScanCloseParams) (r *dorisexternalservice.TScanCloseResult_, err error) {
	var _args backendservice.BackendServiceCloseScannerArgs
	_args.Params = params
	var _result backendservice.BackendServiceCloseScannerResult
	if err = p.c.Call(ctx, "close_scanner", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetStreamLoadRecord(ctx context.Context, lastStreamRecordTime int64) (r *backendservice.TStreamLoadRecordResult_, err error) {
	var _args backendservice.BackendServiceGetStreamLoadRecordArgs
	_args.LastStreamRecordTime = lastStreamRecordTime
	var _result backendservice.BackendServiceGetStreamLoadRecordResult
	if err = p.c.Call(ctx, "get_stream_load_record", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) CheckStorageFormat(ctx context.Context) (r *backendservice.TCheckStorageFormatResult_, err error) {
	var _args backendservice.BackendServiceCheckStorageFormatArgs
	var _result backendservice.BackendServiceCheckStorageFormatResult
	if err = p.c.Call(ctx, "check_storage_format", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) WarmUpCacheAsync(ctx context.Context, request *backendservice.TWarmUpCacheAsyncRequest) (r *backendservice.TWarmUpCacheAsyncResponse, err error) {
	var _args backendservice.BackendServiceWarmUpCacheAsyncArgs
	_args.Request = request
	var _result backendservice.BackendServiceWarmUpCacheAsyncResult
	if err = p.c.Call(ctx, "warm_up_cache_async", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) CheckWarmUpCacheAsync(ctx context.Context, request *backendservice.TCheckWarmUpCacheAsyncRequest) (r *backendservice.TCheckWarmUpCacheAsyncResponse, err error) {
	var _args backendservice.BackendServiceCheckWarmUpCacheAsyncArgs
	_args.Request = request
	var _result backendservice.BackendServiceCheckWarmUpCacheAsyncResult
	if err = p.c.Call(ctx, "check_warm_up_cache_async", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) SyncLoadForTablets(ctx context.Context, request *backendservice.TSyncLoadForTabletsRequest) (r *backendservice.TSyncLoadForTabletsResponse, err error) {
	var _args backendservice.BackendServiceSyncLoadForTabletsArgs
	_args.Request = request
	var _result backendservice.BackendServiceSyncLoadForTabletsResult
	if err = p.c.Call(ctx, "sync_load_for_tablets", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetTopNHotPartitions(ctx context.Context, request *backendservice.TGetTopNHotPartitionsRequest) (r *backendservice.TGetTopNHotPartitionsResponse, err error) {
	var _args backendservice.BackendServiceGetTopNHotPartitionsArgs
	_args.Request = request
	var _result backendservice.BackendServiceGetTopNHotPartitionsResult
	if err = p.c.Call(ctx, "get_top_n_hot_partitions", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) WarmUpTablets(ctx context.Context, request *backendservice.TWarmUpTabletsRequest) (r *backendservice.TWarmUpTabletsResponse, err error) {
	var _args backendservice.BackendServiceWarmUpTabletsArgs
	_args.Request = request
	var _result backendservice.BackendServiceWarmUpTabletsResult
	if err = p.c.Call(ctx, "warm_up_tablets", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) IngestBinlog(ctx context.Context, ingestBinlogRequest *backendservice.TIngestBinlogRequest) (r *backendservice.TIngestBinlogResult_, err error) {
	var _args backendservice.BackendServiceIngestBinlogArgs
	_args.IngestBinlogRequest = ingestBinlogRequest
	var _result backendservice.BackendServiceIngestBinlogResult
	if err = p.c.Call(ctx, "ingest_binlog", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) QueryIngestBinlog(ctx context.Context, queryIngestBinlogRequest *backendservice.TQueryIngestBinlogRequest) (r *backendservice.TQueryIngestBinlogResult_, err error) {
	var _args backendservice.BackendServiceQueryIngestBinlogArgs
	_args.QueryIngestBinlogRequest = queryIngestBinlogRequest
	var _result backendservice.BackendServiceQueryIngestBinlogResult
	if err = p.c.Call(ctx, "query_ingest_binlog", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) PublishTopicInfo(ctx context.Context, topicRequest *backendservice.TPublishTopicRequest) (r *backendservice.TPublishTopicResult_, err error) {
	var _args backendservice.BackendServicePublishTopicInfoArgs
	_args.TopicRequest = topicRequest
	var _result backendservice.BackendServicePublishTopicInfoResult
	if err = p.c.Call(ctx, "publish_topic_info", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetRealtimeExecStatus(ctx context.Context, request *backendservice.TGetRealtimeExecStatusRequest) (r *backendservice.TGetRealtimeExecStatusResponse, err error) {
	var _args backendservice.BackendServiceGetRealtimeExecStatusArgs
	_args.Request = request
	var _result backendservice.BackendServiceGetRealtimeExecStatusResult
	if err = p.c.Call(ctx, "get_realtime_exec_status", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
