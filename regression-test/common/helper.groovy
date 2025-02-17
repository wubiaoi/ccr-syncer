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

import com.google.common.collect.Maps

import java.util.Map
import java.util.List

class Describe {
    String index
    String field
    String type
    Boolean is_key

    Describe(String index, String field, String type, Boolean is_key) {
        this.index = index
        this.field = field
        this.type = type
        this.is_key = is_key
    }

    String toString() {
        return "index: ${index}, field: ${field}, type: ${type}, is_key: ${is_key}"
    }
}

class Helper {
    def suite
    def context
    def logger
    String alias = null

    // the configurations about ccr syncer.
    def sync_gap_time = 5000
    def syncerAddress = "127.0.0.1:9190"

    Helper(suite) {
        this.suite = suite
        this.context = suite.context
        this.logger = suite.logger
    }

    void set_alias(String alias) {
        this.alias = alias
    }

    String randomSuffix() {
        def hashCode = UUID.randomUUID().toString().replace("-", "").hashCode()
        if (hashCode < 0) {
            hashCode *= -1;
        }
        return Integer.toString(hashCode)
    }

    def get_backup_label_prefix(String table = "") {
        return "ccrs_" + get_ccr_job_name(table)
    }

    def get_ccr_job_name(String table = "") {
        def name = context.suiteName
        if (!table.equals("")) {
            name = name + "_" + table
        }
        return name
    }

    def get_ccr_body(String table, String db = null) {
        if (db == null) {
            db = context.dbName
        }

        def gson = new com.google.gson.Gson()

        Map<String, Object> srcSpec = context.getSrcSpec(db)
        srcSpec.put("table", table)

        Map<String, Object> destSpec = context.getDestSpec(db)
        if (alias != null) {
            destSpec.put("table", alias)
        } else {
            destSpec.put("table", table)
        }

        Map<String, Object> body = Maps.newHashMap()
        String name = context.suiteName
        if (!table.equals("")) {
            name = name + "_" + table
        }
        body.put("name", name)
        body.put("src", srcSpec)
        body.put("dest", destSpec)

        return gson.toJson(body)
    }

    void ccrJobDelete(table = "") {
        def bodyJson = get_ccr_body "${table}"
        suite.httpTest {
            uri "/delete"
            endpoint syncerAddress
            body "${bodyJson}"
            op "post"
        }
    }

    void ccrJobCreate(table = "") {
        def bodyJson = get_ccr_body "${table}"
        suite.httpTest {
            uri "/create_ccr"
            endpoint syncerAddress
            body "${bodyJson}"
            op "post"
            check { code, body ->
                if (!"${code}".toString().equals("200")) {
                    throw new Exception("request failed, code: ${code}, body: ${body}")
                }
                def jsonSlurper = new groovy.json.JsonSlurper()
                def object = jsonSlurper.parseText "${body}"
                if (!object.success) {
                    throw new Exception("request failed, error msg: ${object.error_msg}")
                }
            }
        }
    }

    void ccrJobCreateAllowTableExists(table = "") {
        def bodyJson = get_ccr_body "${table}"
        def jsonSlurper = new groovy.json.JsonSlurper()
        def object = jsonSlurper.parseText "${bodyJson}"
        object['allow_table_exists'] = true
        logger.info("json object ${object}")

        bodyJson = new groovy.json.JsonBuilder(object).toString()
        suite.httpTest {
            uri "/create_ccr"
            endpoint syncerAddress
            body "${bodyJson}"
            op "post"
        }
    }

    void ccrJobPause(table = "") {
        def bodyJson = get_ccr_body "${table}"
        suite.httpTest {
            uri "/pause"
            endpoint syncerAddress
            body "${bodyJson}"
            op "post"
        }
    }

    void ccrJobResume(table = "") {
        def bodyJson = get_ccr_body "${table}"
        suite.httpTest {
            uri "/resume"
            endpoint syncerAddress
            body "${bodyJson}"
            op "post"
        }
    }

    void ccrJobDesync(table = "") {
        def bodyJson = get_ccr_body "${table}"
        suite.httpTest {
            uri "/desync"
            endpoint syncerAddress
            body "${bodyJson}"
            op "post"
        }
    }

    void enableDbBinlog() {
        suite.sql """
            ALTER DATABASE ${context.dbName} SET properties ("binlog.enable" = "true")
            """
    }

    Boolean checkShowTimesOf(sqlString, myClosure, times, func = "sql") {
        Boolean ret = false
        List<List<Object>> res
        while (times > 0) {
            try {
                if (func == "sql") {
                    res = suite.sql "${sqlString}"
                } else {
                    res = suite.target_sql "${sqlString}"
                }
                if (myClosure.call(res)) {
                    ret = true
                }
            } catch (Exception e) {}

            if (ret) {
                break
            } else if (--times > 0) {
                sleep(sync_gap_time)
            }
        }

        if (!ret) {
            logger.info("last select result: ${res}")
        }

        return ret
    }

    // wait until all restore tasks of the dest cluster are finished.
    Boolean checkRestoreFinishTimesOf(checkTable, times) {
        Boolean ret = false
        while (times > 0) {
            def sqlInfo = suite.target_sql "SHOW RESTORE FROM TEST_${context.dbName}"
            for (List<Object> row : sqlInfo) {
                if ((row[10] as String).contains(checkTable)) {
                    logger.info("SHOW RESTORE result: ${row}")
                    ret = (row[4] as String) == "FINISHED"
                }
            }

            if (ret) {
                break
            } else if (--times > 0) {
                sleep(sync_gap_time)
            }
        }

        return ret
    }

    // Check N times whether the num of rows of the downstream data is expected.
    Boolean checkSelectTimesOf(sqlString, rowSize, times) {
        def tmpRes = suite.target_sql "${sqlString}"
        while (tmpRes.size() != rowSize) {
            sleep(sync_gap_time)
            if (--times > 0) {
                tmpRes = suite.target_sql "${sqlString}"
            } else {
                logger.info("last select result: ${tmpRes}")
                logger.info("expected row size: ${rowSize}, actual row size: ${tmpRes.size()}")
                break
            }
        }
        return tmpRes.size() == rowSize
    }

    Boolean checkSelectColTimesOf(sqlString, colSize, times) {
        def tmpRes = suite.target_sql "${sqlString}"
        while (tmpRes.size() == 0 || tmpRes[0].size() != colSize) {
            sleep(sync_gap_time)
            if (--times > 0) {
                tmpRes = suite.target_sql "${sqlString}"
            } else {
                break
            }
        }
        return tmpRes.size() > 0 && tmpRes[0].size() == colSize
    }

    Boolean checkData(data, beginCol, value) {
        if (data.size() < beginCol + value.size()) {
            return false
        }

        for (int i = 0; i < value.size(); ++i) {
            if ((data[beginCol + i]) as int != value[i]) {
                return false
            }
        }

        return true
    }

    Integer getRestoreRowSize(checkTable) {
        def result = suite.target_sql "SHOW RESTORE FROM TEST_${context.dbName}"
        def size = 0
        for (List<Object> row : result) {
            if ((row[10] as String).contains(checkTable)) {
                size += 1
            }
        }

        return size
    }

    Boolean checkRestoreNumAndFinishedTimesOf(checkTable, expectedRestoreRows, times) {
        while (times > 0) {
            def restore_size = getRestoreRowSize(checkTable)
            if (restore_size >= expectedRestoreRows) {
                return checkRestoreFinishTimesOf(checkTable, times)
            }
            if (--times > 0) {
                sleep(sync_gap_time)
            }
        }

        return false
    }

    void force_fullsync(tableName = "") {
        def bodyJson = get_ccr_body "${tableName}"
        suite.httpTest {
            uri "/force_fullsync"
            endpoint syncerAddress
            body "${bodyJson}"
            op "post"
        }
    }

    Object get_job_progress(tableName = "") {
        def request_body = get_ccr_body(tableName)
        def get_job_progress_uri = { check_func ->
            suite.httpTest {
                uri "/job_progress"
                endpoint syncerAddress
                body request_body
                op "post"
                check check_func
            }
        }

        def result = null
        get_job_progress_uri.call() { code, body ->
            if (!"${code}".toString().equals("200")) {
                throw "request failed, code: ${code}, body: ${body}"
            }
            def jsonSlurper = new groovy.json.JsonSlurper()
            def object = jsonSlurper.parseText "${body}"
            if (!object.success) {
                throw "request failed, error msg: ${object.error_msg}"
            }
            logger.info("job progress: ${object.job_progress}")
            result = object.job_progress
        }
        return result
    }

    // test whether the ccr syncer has set a feature flag?
    Boolean has_feature(name) {
        def features_uri = { check_func ->
            suite.httpTest {
                uri "/features"
                endpoint syncerAddress
                body ""
                op "get"
                check check_func
            }
        }

        def result = null
        features_uri.call() { code, body ->
            if (!"${code}".toString().equals("200")) {
                throw "request failed, code: ${code}, body: ${body}"
            }
            def jsonSlurper = new groovy.json.JsonSlurper()
            def object = jsonSlurper.parseText "${body}"
            if (!object.success) {
                throw "request failed, error msg: ${object.error_msg}"
            }
            logger.info("features: ${object.flags}")
            result = object.flags
        }

        for (def flag in result) {
            if (flag.feature == name && flag.value) {
                return true
            }
        }
        return false
    }

    String upstream_version() {
        def version_variables = suite.sql_return_maparray "show variables like 'version_comment'"
        return version_variables[0].Value
    }

    Boolean is_version_supported(versions) {
        def version_variables = suite.sql_return_maparray "show variables like 'version_comment'"
        def matcher = version_variables[0].Value =~ /doris-(\d+\.\d+\.\d+)/
        if (matcher.find()) {
            def parts = matcher.group(1).tokenize('.')
            def major = parts[0].toLong()
            def minor = parts[1].toLong()
            def patch = parts[2].toLong()
            def version = String.format("%d%02d%02d", major, minor, patch).toLong()
            for (long expect : versions) {
                logger.info("current version ${version}, expect version ${expect}")
                def expect_version_set = expect / 100
                def got_version_set = version / 100
                if (expect_version_set == got_version_set && version < expect) {
                    return false
                }
            }
        }
        return true
    }

    Map<String, List<Describe>> get_table_describe(String table, String source = "sql") {
        def res
        if (source == "sql") {
            res = suite.sql_return_maparray "DESC ${table} ALL"
        } else {
            res = suite.target_sql_return_maparray "DESC ${table} ALL"
        }

        def map = Maps.newHashMap()
        def index = ""
        for (def row : res) {
            if (row.IndexName != "") {
                index = row.IndexName
            }
            if (row.Field == "") {
                continue
            }

            if (!map.containsKey(index)) {
                map.put(index, [])
            }
            def is_key = false
            if (row.Key == "true" || row.Key == "YES") {
                is_key = true
            }
            map.get(index).add(new Describe(index, row.Field, row.Type, is_key))
        }
        return map
    }

    Boolean check_describes(Map<String, List<Describe>> expect, Map<String, List<Describe>> actual) {
        if (actual.size() != expect.size()) {
            return false
        }

        for (def key : expect.keySet()) {
            if (!actual.containsKey(key)) {
                return false
            }
            def expect_list = expect.get(key)
            def actual_list = actual.get(key)
            if (expect_list.size() != actual_list.size()) {
                return false
            }
            for (int i = 0; i < expect_list.size(); ++i) {
                if (expect_list[i].toString() != actual_list[i].toString()) {
                    return false
                }
            }
        }
        return true
    }

    Boolean check_table_describe_times(String table, times = 30) {
        while (times > 0) {
            def upstream_describe = get_table_describe(table)
            def downstream_describe = get_table_describe(table, "target")
            if (check_describes(upstream_describe, downstream_describe)) {
                return true
            }
            sleep(sync_gap_time)
            times--
        }

        def upstream_describe = get_table_describe(table)
        def downstream_describe = get_table_describe(table, "target")
        logger.info("upstream describe: ${upstream_describe}")
        logger.info("downstream describe: ${downstream_describe}")
        return false
    }

    Boolean check_table_exists(String table, times = 30) {
        while (times > 0) {
            def res = suite.target_sql "SHOW TABLES LIKE '${table}'"
            if (res.size() > 0) {
                return true
            }
            sleep(sync_gap_time)
            times--
        }
        return false
    }
}

new Helper(suite)
