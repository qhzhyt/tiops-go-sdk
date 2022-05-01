package record

import (
	"time"
	"tiops/common/models"
	"tiops/common/utils"
	"tiops/engine/types"
)

type ExecutionRecordManager struct {
	ExecutionRecordID string
	updateInterval    time.Duration
	ticker            *time.Ticker
	//*types.EngineContext
	processRecords    []*models.ProcessRecord
	oldProcessRecords []*models.ProcessRecord
	batchSizes        map[string]int64
	countFactor       int
}

const maxRecordCount = 3600

func (m *ExecutionRecordManager) Start() {
	go func() {

		for range m.ticker.C {
			//m.Error(m.processRecords)
			processRecords := m.processRecords

			processRecordMap := map[string]*models.ProcessRecord{}

			m.processRecords = nil

			for _, record := range processRecords {
				if processRecordMap[record.NodeId] == nil {
					record.DataIds = []string{record.RecordId}
					processRecordMap[record.NodeId] = record
				} else {
					oldRecord := processRecordMap[record.NodeId]
					oldRecord.EndTime = record.EndTime
					oldRecord.BatchCount = oldRecord.BatchCount + record.BatchCount
					oldRecord.ElapsedTime = oldRecord.ElapsedTime + record.ElapsedTime
					oldRecord.ItemCount = oldRecord.ItemCount + record.ItemCount
					oldRecord.ProcessRate = float32(oldRecord.ItemCount) / float32(oldRecord.ElapsedTime) * 1000
					oldRecord.BacklogBatches = record.BacklogBatches
					oldRecord.DataIds = append(oldRecord.DataIds, record.RecordId)
				}
			}

			for _, record := range processRecordMap {
				if m.batchSizes[record.NodeId] < 1 {
					m.batchSizes[record.NodeId] = record.ItemCount / record.BatchCount
				}
				record.BatchSize = m.batchSizes[record.NodeId]
				record.RecordTime = utils.CurrentTimeStampMS()

				m.oldProcessRecords = append(m.oldProcessRecords, record)

				/*_, err := m.AddProcessRecord(record)
				if err != nil {
					m.Error(err.Error())
				}*/
			}
		}
	}()
}

func (m *ExecutionRecordManager) Records() []*models.ProcessRecord {
	result := m.oldProcessRecords
	result = append(result, m.processRecords...)
	return result
}

func (m *ExecutionRecordManager) AddProcessRecord(record *models.ProcessRecord) {
	m.processRecords = append(m.processRecords, record)
}

func NewExecutionRecordManager(updateInterval time.Duration, ctx *types.EngineContext) *ExecutionRecordManager {
	return &ExecutionRecordManager{
		//EngineContext: ctx,
		updateInterval: updateInterval,
		ticker:         time.NewTicker(updateInterval),
		batchSizes:     map[string]int64{},
		countFactor:    1,
	}
}
