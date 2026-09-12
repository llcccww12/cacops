import uuidv1 from 'uuidv1';
// retrieve a HTML string for given SVG icon name and size in pixels
export function svg(name, size) {
  return `<svg class="svg ${name}" width="${size}" height="${size}" aria-hidden="true"><use xlink:href="#${name}"/></svg>`;
}

// transform /path/to/file.ext to file.ext
export function basename(path = '') {
  return path ? path.replace(/^.*\//, '') : '';
}

// transform /path/to/file.ext to .ext
export function extname(path = '') {
  const [_, ext] = /.+(\.[^.]+)$/.exec(path) || [];
  return ext || '';
}

// test whether a variable is an object
export function isObject(obj) {
  return Object.prototype.toString.call(obj) === '[object Object]';
}

// returns whether a dark theme is enabled
export function isDarkTheme() {
  return document.documentElement.classList.contains('theme-arc-green');
}

export function matterTree(data) {
  for (let i = 0; i < data.length; i++) {
    data[i].label = data[i].name;
    data[i].filePath = data[i].path;
    data[i].isEdit = true;
    if (data[i].children && !data[i].children.length) data[i].isLeaf = true;
    if (data[i].children && data[i].children.length > 0) {
      // children若不为空数组，则继续 递归调用 本方法
      matterTree(data[i].children);
    }
  }
  return data;
}

export function disposeTreeName(data, target) {
  for (let i = 0; i < data.length; i++) {
    data[i].filePath = `${target}/${data[i].name}`;
    data[i].id = uuidv1().replaceAll('-', '');
    if (data[i].children && data[i].children.length > 0) {
      // children若不为空数组，则继续 递归调用 本方法
      matterTree(data[i].children);
    }
  }
  return data;
}

export function RecurveAddNode(treeRootData, treeD, FilePath) {
  if(FilePath === ""){
    treeRootData.push(treeD);
    return treeRootData;
  }
  treeRootData.forEach((item) => {
    if (item.filePath === FilePath) {
      return item.children.push(treeD);
    }
    if (item.children?.length) {
      RecurveAddNode(item.children, treeD, FilePath);
    }
  });
  return treeRootData;
}

// 递归单个循环操作 判断是否新增
export function RecurveQuery(currentData, filePath) {
  currentData.forEach((item, index) => {
    if (item.filePath === filePath && item.isLeaf) {
      currentData[index].operation = 'delete';
    } else if (item.children?.length > 0) {
      RecurveQuery(item.children, filePath);
    }
  });
  return currentData;
}

// 递归 对新增的数据进行删除
export function RecurveQueryDelete(data, filePath) {
  data.forEach((item, index) => {
    if (item.filePath === filePath && item.isLeaf) {
      data.splice(index, 1);
    } else if (item.children?.length) {
      RecurveQueryDelete(item.children, filePath);
    }
  });
  return data;
}

// 循环插入改变的数据
export function RecurveUpdateValue(currentData, fileInfoParams, base64Content) {
  currentData.forEach((item, index) => {
    if (item.filePath === fileInfoParams.filePath) {
      currentData[index].oldContent = base64Content;
      currentData[index].content = base64Content;
    } else if (item.children?.length > 0) {
      RecurveQuery(item.children, fileInfoParams.filePath);
    }
  });
  return currentData;
}



const flattenTree = (data) => {
  return data.reduce((arr, currentValue) => {
    if (currentValue.children && currentValue.children.length > 0) {
      return arr.concat(flattenTree(currentValue.children))
    } else {
      return arr.concat(currentValue)
    }
  }, [])
}

export const uniqueArray = (array1, array2) => {
  return array1
    .concat(array2)
    .reduce((prev, cur) => {
      const duplicate = prev.filter(x => x.name === cur.name)
      if (duplicate.length) {
        return prev
      } else {
        prev.push(cur)
        return prev
      }
    }, [])
}

export const isBase64 = (str) => {
  if(!str) return false;
  if (/^([0-9a-zA-Z+/]{4})*(([0-9a-zA-Z+/]{2}==)|([0-9a-zA-Z+/]{3}=))?$/.test(str)) {
    try {
      return window.btoa(window.atob(str)) == str;
    } catch (err) {
      return false;
    }
  }else{
    return false
  }

}




export const formatTime = (value,format) => {
  format = format || "Y-M-D h:m";
  const formatNumber = n => `0${n}`.slice(-2);
  const date = new Date(value);
  const formatList = ["Y", "M", "D", "h", "m", "s"];
  const resultList = [];
  resultList.push(date.getFullYear().toString());
  resultList.push(formatNumber(date.getMonth() + 1));
  resultList.push(formatNumber(date.getDate()));
  resultList.push(formatNumber(date.getHours()));
  resultList.push(formatNumber(date.getMinutes()));
  resultList.push(formatNumber(date.getSeconds()));
  for (let i = 0; i < resultList.length; i++) {
    format = format.replace(formatList[i], resultList[i]);
  }
  return format;
}

export function objectToQuery() {
  let obj = arguments[0];
  let prefix = arguments[1];
  if (typeof obj !== "object") return "";
  const attrs = Object.keys(obj);
  return attrs.reduce((query, attr, index) => {
    // 判断是否是第一层第一个循环
    if (index === 0 && !prefix) query += "?";
    if (typeof obj[attr] === "object") {
      const subPrefix = prefix ? `${prefix}[${attr}]` : attr;
      query += this.objectToQuery(obj[attr], subPrefix);
    } else {
      if (prefix) {
        query += `${prefix}[${attr}]=${obj[attr]}`;
      } else {
        query += `${attr}=${obj[attr]}`;
      }
    }
    // 判断是否是第一层最后一个循环
    if (index !== attrs.length - 1) query += "&";
    return query;
  }, "");
}

export const  getOnwerOrProject = () => {
  const obj = location.pathname.replace(/\//g," ").trim().split(" ");
  return {
    owner:obj[0],
    project:obj[1]
  }
}
