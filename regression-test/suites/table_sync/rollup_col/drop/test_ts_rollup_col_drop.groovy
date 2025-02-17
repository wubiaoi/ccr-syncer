// Licensed to the Apache Software Foundation (ASF) under one
// or more contributor license agreements.  See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership.  The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License.  You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.
suite("test_ts_rollup_col_drop") {
    def helper = new GroovyShell(new Binding(['suite': delegate]))
            .evaluate(new File("${context.config.suitePath}/../common", "helper.groovy"))

    def tableName = "tbl_" + helper.randomSuffix()
    def test_num = 0
    def insert_num = 5

    def exist = { res -> Boolean
        return res.size() != 0
    }

    def has_count = { count ->
        return { res -> Boolean
            res.size() == count
        }
    }

    sql """
        CREATE TABLE if NOT EXISTS ${tableName}
        (
            `id` INT,
            `col1` INT,
            `col2` INT,
            `col3` INT,
            `col4` INT,
        )
        ENGINE=OLAP
        DUPLICATE KEY(`id`, `col1`, `col2`)
        DISTRIBUTED BY HASH(id) BUCKETS 1
        PROPERTIES (
            "replication_allocation" = "tag.location.default: 1",
            "binlog.enable" = "true"
        )
    """
    sql """
        ALTER TABLE ${tableName}
        ADD ROLLUP rollup_${tableName} (id, col2, col4)
        """

    assertTrue(helper.checkShowTimesOf("""
                                SHOW ALTER TABLE ROLLUP
                                FROM ${context.dbName}
                                WHERE TableName = "${tableName}" AND State = "FINISHED"
                                """,
                                has_count(1), 30))

    helper.ccrJobCreate(tableName)
    assertTrue(helper.checkRestoreFinishTimesOf("${tableName}", 30))
    assertTrue(helper.check_table_describe_times(tableName, 30))

    first_job_progress = helper.get_job_progress(tableName)

    logger.info("=== Test 1: drop key column ===")
    // {
    //   "type": "SCHEMA_CHANGE",
    //   "dbId": 10577,
    //   "tableId": 10640,
    //   "tableName": "tbl_1919050016",
    //   "jobId": 10682,
    //   "jobState": "FINISHED",
    //   "rawSql": "ALTER TABLE `regression_test_table_sync_rollup_col_drop`.`tbl_1919050016` DROP COLUMN `col2` IN `rollup_tbl_1919050016`",
    //   "iim": {
    //     "10683": 10647
    //   }
    // }
    sql """
        ALTER TABLE ${tableName}
        DROP COLUMN `col2`
        FROM rollup_${tableName}
        """
    sql "sync"

    assertTrue(helper.checkShowTimesOf("""
                                SHOW ALTER TABLE COLUMN
                                FROM ${context.dbName}
                                WHERE TableName = "${tableName}"
                                    AND IndexName = "rollup_${tableName}"
                                    AND State = "FINISHED"
                                """,
                                has_count(1), 30))

    assertTrue(helper.check_table_describe_times(tableName, 30))

    logger.info("=== Test 2: drop value column ===")
    // {
    //   "type": "SCHEMA_CHANGE",
    //   "dbId": 10577,
    //   "tableId": 10640,
    //   "tableName": "tbl_1919050016",
    //   "jobId": 10717,
    //   "jobState": "FINISHED",
    //   "rawSql": "ALTER TABLE `regression_test_table_sync_rollup_col_drop`.`tbl_1919050016` DROP COLUMN `col4` IN `rollup_tbl_1919050016`",
    //   "iim": {
    //     "10718": 10683
    //   }
    // }
    sql """
        ALTER TABLE ${tableName}
        DROP COLUMN `col4`
        FROM rollup_${tableName}
        """
    sql "sync"

    assertTrue(helper.checkShowTimesOf("""
                                SHOW ALTER TABLE COLUMN
                                FROM ${context.dbName}
                                WHERE TableName = "${tableName}"
                                    AND IndexName = "rollup_${tableName}"
                                    AND State = "FINISHED"
                                """,
                                has_count(2), 30))
    assertTrue(helper.check_table_describe_times(tableName, 30))

    // no full sync triggered.
    last_job_progress = helper.get_job_progress(tableName)
    assertTrue(last_job_progress.full_sync_start_at == first_job_progress.full_sync_start_at)
}
