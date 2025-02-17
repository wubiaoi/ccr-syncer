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

suite("test_ts_prop_auto_increment") {
    def helper = new GroovyShell(new Binding(['suite': delegate]))
            .evaluate(new File("${context.config.suitePath}/../common", "helper.groovy"))

    def dbName = context.dbName
    def tableName = "tbl_" + helper.randomSuffix()
    def test_num = 0
    def insert_num = 5

    def exist = { res -> Boolean
        return res.size() != 0
    }

    sql "DROP TABLE IF EXISTS ${dbName}.${tableName}"
    target_sql "DROP TABLE IF EXISTS TEST_${dbName}.${tableName}"

    helper.enableDbBinlog()

    sql """
            CREATE TABLE ${tableName} (
                `id` BIGINT NOT NULL AUTO_INCREMENT,
                `value` int(11) NOT NULL
            ) ENGINE=OLAP
            DUPLICATE KEY(`id`)
            PROPERTIES (
            "replication_allocation" = "tag.location.default: 1"
            )
    """

    for (int index = 0; index < insert_num; index++) {
        sql "INSERT INTO ${tableName} (value) VALUES (${insert_num})"
    }
    sql "sync"

    helper.ccrJobDelete(tableName)
    helper.ccrJobCreate(tableName)

    assertTrue(helper.checkRestoreFinishTimesOf("${tableName}", 30))

    assertTrue(helper.checkShowTimesOf("SHOW TABLES LIKE \"${tableName}\"", exist, 60, "sql"))

    assertTrue(helper.checkShowTimesOf("SHOW TABLES LIKE \"${tableName}\"", exist, 60, "target"))

    def target_res = target_sql "SHOW CREATE TABLE ${tableName}"

    assertTrue(target_res[0][1].contains("`id` bigint NOT NULL AUTO_INCREMENT(1)"))

    res = sql "select * from ${tableName} order by id"

    target_res = target_sql "select * from ${tableName} order by id"

    assertEquals(target_res, res)
}