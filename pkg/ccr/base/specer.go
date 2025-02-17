package base

import (
	"github.com/selectdb/ccr_syncer/pkg/ccr/record"
	"github.com/selectdb/ccr_syncer/pkg/utils"
)

type SpecEvent int

const (
	feNotMasterEvent  SpecEvent = 0
	httpNotFoundEvent SpecEvent = 1
)

// this interface is used to for spec operation, treat it as a mysql dao
type Specer interface {
	Valid() error
	IsDatabaseEnableBinlog() (bool, error)
	IsTableEnableBinlog() (bool, error)
	IsEnableRestoreSnapshotCompression() (bool, error)
	GetAllTables() ([]string, error)
	GetAllViewsFromTable(tableName string) ([]string, error)
	ClearDB() error
	CreateDatabase() error
	CreateTableOrView(createTable *record.CreateTable, srcDatabase string) error
	CheckDatabaseExists() (bool, error)
	CheckTableExists() (bool, error)
	CheckTableExistsByName(tableName string) (bool, error)
	GetValidBackupJob(snapshotNamePrefix string) (string, error)
	GetValidRestoreJob(snapshotNamePrefix string) (string, error)
	CancelRestoreIfExists(snapshotName string) error
	CreatePartialSnapshot(snapshotName, table string, partitions []string) error
	CreateSnapshot(snapshotName string, tables []string) error
	CheckBackupFinished(snapshotName string) (bool, error)
	CheckRestoreFinished(snapshotName string) (bool, error)
	GetRestoreSignatureNotMatchedTableOrView(snapshotName string) (string, bool, error)
	WaitTransactionDone(txnId int64) // busy wait

	LightningSchemaChange(srcDatabase string, tableAlias string, changes *record.ModifyTableAddOrDropColumns) error
	RenameColumn(destTableName string, renameColumn *record.RenameColumn) error
	RenameTable(destTableName string, renameTable *record.RenameTable) error
	RenameTableWithName(destTableName, newName string) error
	ModifyComment(destTableName string, modifyComment *record.ModifyComment) error
	TruncateTable(destTableName string, truncateTable *record.TruncateTable) error
	ReplaceTable(fromName, toName string, swap bool) error
	DropTable(tableName string, force bool) error
	DropView(viewName string) error
	AlterViewDef(srcDatabase, viewName string, alterView *record.AlterView) error

	AddPartition(destTableName string, addPartition *record.AddPartition) error
	DropPartition(destTableName string, dropPartition *record.DropPartition) error
	RenamePartition(destTableName, oldPartition, newPartition string) error

	LightningIndexChange(tableAlias string, changes *record.ModifyTableAddOrDropInvertedIndices) error
	BuildIndex(tableAlias string, buildIndex *record.IndexChangeJob) error

	RenameRollup(destTableName, oldRollup, newRollup string) error
	DropRollup(destTableName, rollupName string) error

	DesyncTables(tables ...string) error

	utils.Subject[SpecEvent]
}
