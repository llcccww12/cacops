let db;
let openRequest;
const indexedDB =
    window.indexedDB ||
    window.webkitIndexedDb ||
    window.mozIndexed ||
    window.msIndexedDB;

// transaction 事务处理 意味着一系列操作步骤之中，只要有一步失败，整个事务就都取消，数据库回滚到事务发生之前的状态，不存在只改写一部分数据的情况。

// tableList [{tableName,keyPath,indexName}]
/**
 * 初始化数据库
 * @param {string} dbName 数据库名
 * @param {Array} tableList 数据表列表
 */

// 初始化
export function init({ dbName, tableList }) {
    // 存在就打开 不存在新建 第二个参数为db 版本
    openRequest = indexedDB.open(dbName);
    // 新的数据库创建 或者数据库的版本号被更改会被触发
    openRequest.onupgradeneeded = function (e) {
        // 表的创建在这个回调里执行
        const thisDb = e.target.result;
        if (tableList?.length) {
            tableList.forEach((table) => {
                if (!thisDb.objectStoreNames.contains(table.tableName)) {
                    // keyPath 主键 autoIncrement 是否自增
                    const objectStore = thisDb.createObjectStore(table.tableName, {
                        keyPath: table.keyPath,
                        ...table.attr,
                        // autoIncrement: true,
                    });
                    if (table.indexName) {
                        // 创建表的时候 可以去指定那些字段是可以被索引的字段
                        objectStore.createIndex(table.indexName, table.indexName, {
                            unique: table.unique || false,
                        });
                    }
                }
            });
        } else {
            console.error('请传入数据表参数');
        }
    };
    // 已经创建好的数据库创建成功的时候
    openRequest.onsuccess = function (e) {
        db = e.target.result;
        window.db = e.target.result;
        db.onerror = function (event) {
        };
    };
    // 打开失败时调用
    openRequest.onerror = function (e) {
        console.error('openRequest.onerror', e);
    };
}

/**
 * 添加一行
 * @param {string} tableName 表名
 * @param {object} data 数据
 * @returns {promise}
 */
export function add({ tableName, data }) {
    return new Promise((suc, fail) => {
        const request = window.db
            .transaction([tableName], 'readwrite')
            .objectStore(tableName)
            .add(data);

        request.onsuccess = function (event) {
            suc();
        };

        request.onerror = function (event) {
            // fail();
            update({tableName, data })
        };
    });
}

/**
 * 遍历所有数据 其实这里有点坑 真正readAll是下面readAllPC这个方法，因为低版本webview手机不兼容objectStore.getAll所以使用游标来实现
 * @param {string} tableName 表名
 * @returns {promise}
 */
export function readAll(tableName) {
    return new Promise((suc, fail) => {
        const objectStore = db
            .transaction([tableName], 'readwrite')
            .objectStore(tableName);
        // 我这里做的是把所有的结果全部收集起来 当然我们可以做其他事情此处拿到的value是每条数据的结果、还有primaryKey主键、key、与direction
        const result = [];
        objectStore.openCursor().onsuccess = function (event) {
            const cursor = event.target.result;
            if (cursor) {
                result.push({ ...cursor.value });
                cursor.continue();
            } else {
                suc(result);
            }
        };
        objectStore.openCursor().onerror = function (event) {
            console.dir(event);
            fail();
        };
    });
}

/**
 * 读取所有数据 仅在pc上或者版本高的手机浏览器中使用 readAllPC这个方法 这个在低版本weview的手机浏览器里面不兼容
 * @param {string} tableName
 * @returns {promise}
 */
export function readAllForHighVersion(tableName) {
    return new Promise((suc, fail) => {
        const objectStore = db
            .transaction([tableName], 'readwrite')
            .objectStore(tableName);
        const request = objectStore.getAll();

        request.onerror = function (event) {
            fail();
        };

        request.onsuccess = function () {
            suc(request.result || []);
        };
    });
}

/**
 * 根据主键查询对应数据
 * @param {string} tableName 表名
 * @param {string} key 主键
 * @returns {promise}
 */
export function readByMainKey({ tableName, key }) {
    return new Promise((suc, fail) => {
        if (!key) {
            fail();
            return;
        }
        const objectStore = db
            .transaction([tableName], 'readwrite')
            .objectStore(tableName);
        const request = objectStore.get(key);

        request.onerror = function (event) {
            console.error('根据主键查询对应数据', event);
            fail();
        };

        request.onsuccess = function () {
            suc(request.result || {});
        };
    });
}

/**
 * 根据主键删除对应数据
 * @param {string} tableName 表名
 * @param {string} key 主键
 * @returns {promise}
 */
export function remove({ tableName, key }) {
    return new Promise((suc, fail) => {
        const objectStore = db
            .transaction([tableName], 'readwrite')
            .objectStore(tableName);
        const request = objectStore.delete(key);

        request.onerror = function (event) {
            fail();
        };

        request.onsuccess = function (event) {
            suc();
        };
    });
}

/**
 * 根据主键更新对应数据
 * @param {object} data 对应的主键与值 和 数据
 * @param {string} tableName 表名
 * @returns {promise}
 */
export function update({ tableName, data }) {
    return new Promise((suc, fail) => {
        const objectStore = db
            .transaction([tableName], 'readwrite')
            .objectStore(tableName);
        const request = objectStore.put(data);

        request.onerror = function (event) {
            fail();
        };

        request.onsuccess = function (event) {
            suc();
        };
    });
}

/**
 * 通过索引查找对应数据
 * @param {string} indexName 索引名称
 * @param {any} indexVal index 索引值
 * @param {string} tableName 表名
 * @returns {promise}
 */
export function readByIndex({ tableName, indexName, indexVal }) {
    return new Promise((suc, fail) => {
        const objectStore = db
            .transaction([tableName], 'readwrite')
            .objectStore(tableName);
        // 假定新建表格的时候，对name字段建立了索引。
        // objectStore.createIndex('name', 'name', { unique: false });

        const index = objectStore.index(indexName);
        const request = index.get(indexVal);
        request.onerror = function (event) {
            fail();
        };

        request.onsuccess = function (event) {
            if (request.result) {
                suc(request.result);
            } else {
            }
        };
    });
}


/**
 * 删除数据库
 * @param {string} DB_NAME 数据库名称
 * @returns
 */
export async function deleteDB(DB_NAME) {
    return indexedDB.deleteDatabase(DB_NAME);
}
/**
 * 关闭数据库
 * @param {string} DB_NAME 数据库名称
 * @returns
 */
export function closeDB(DB_NAME) {
    return indexedDB.close(DB_NAME);
}
/**
 * 清除表
 * @param {string} tableName
 * @returns {promise}
 */
export function clearTable(tableName) {
    return new Promise((suc, fail) => {
        const objectStore = db
            .transaction([tableName], 'readwrite')
            .objectStore(tableName);
        const request = objectStore.clear();

        request.onerror = function (event) {
            fail();
        };

        request.onsuccess = function (event) {
            suc();
        };
    });
}

//   export default {
//       deleteDB,
//       updateDB,
//       getDataByKey,
//       addData
//   }
