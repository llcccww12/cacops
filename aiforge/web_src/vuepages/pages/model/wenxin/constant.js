export const ERNIE_CASE = [
    '中国山水画，青山绿水，溪水长流，古风，青山相看两不厌，丹青水墨，中国风',
    '这是一幅中国画风格的画，画的是山上一个宁静的湖泊，头顶上有一片清澈的蓝天',
    '云上的繁华都市中，世界传说，仙境，插画作品，超现实主义',
    '一只带着酷酷眼镜的蒸汽朋克猫portait，在艺术界大行其道，蒸汽波艺术',
    '心脏，人类的器官，钢铁材质，机械感，精致，细节清晰，未来科技，未来主义',
    '天空中有颗巨大的心，光线柔和，心形光圈，写实风格',
    '泰迪熊穿着西服在时代广场跳舞，细节清晰，高清，8k',
    '钛合金，机甲兔子，赛博朋克，未来感，设计作品，cg感，摄影棚光照，未来主义',
    '青绿色草地连成一片，山间起伏，有倾泻的流水，星星点缀天空，蓝天白云，天色泛白，巨幅画面感，遥拍视角，4K，高清',
    '抹茶味奶茶，古香书桌，安静的环境，落地窗外天清云淡，写实主义',
    '蓝色星体漂浮在破碎的红色城堡之上，概念艺术，吉姆·马霍德风格',
    '金凤凰，背景绚烂，高饱和，古风，仙境，高清，4K，古风',
    '红色灯笼，长城，秋天，高清，奇幻风格，史诗感，艺术站，超现实主义',
    '海滩，落日，棕榈树，莫奈风格',
    '古长城，烽火台，山河无恙，国泰民安，超现实主义',
    '短发，高清，精致面容，细节清晰，少年，二次元，cg感',
    '戴贝雷帽的女孩，银色的长发，白色的斗篷，手拿气球，日本动漫风格',
    '穿着时尚潮牌的一只大熊猫，日出，大步向我们走来，赛博朋克，烟雾',
    '超级逼真的未来世界，真实照片，虚幻引擎',
]
export const ERNIE_CASE_IMG = [
    '桃花海，桃花源，二次元，古风，pixiv，浪漫，唯美，户外场景，厚涂',
    '机械姬，头像，赛博朋克，精致妆容，二次元',
    '神话般的水晶城堡，艺术站，虚拟现实，高清，手绘，厚涂，闪闪发光',
    '水粉，动漫，插画，古风建筑群，莲花池，府邸，明月夜，蓝色配色，国风国潮，体积照明，光影效果，特写镜头，细节丰富，超广角，超宽画幅'
]
/**
 * 轮询类封装
 * @param config 配置对象
 * @constructor
 */
export function Thread(config){
  this.params = this.init(config);
}

/**
* 初始化对象参数
* @param config 配置对象
*/
Thread.prototype.init = function (config) {
  const params = config ? config : {
      start: function () {},
      stop: function () {},
      number: 0,
      time: 300
  };
  params.start = params.start ? params.start : function (){};
  params.stop = params.stop ? params.stop : function (){};
  params.number = params.number ? Math.abs(parseInt(params.number)) : 0;
  params.time = params.time ? Math.abs(parseInt(params.time)) : 300;
  this.time = params.time
  return params;
}

/**
* 轮询实现
* @param start
* @param stop
* @param number
* @param time
*/
Thread.prototype.run = function () {
  this.timer = setTimeout(this.runTime.bind(this), this.time)
}

/**
* 执行过程函数
*/
Thread.prototype.runTime = function () {
  try {
      this.params.start()
  } finally {
      if (!this.params.number == 0) {
          if (this.total >= this.params.number) {
              this.params.stop()
              clearTimeout(this.timer)
              return
          }
          if (!this.total) {
              this.total = 1
          }
          this.total++
      }
      clearTimeout(this.timer)
  }
  this.timer = setTimeout(this.runTime.bind(this), this.time)
}

/**
* 停止轮询
*/
Thread.prototype.stop = function () {
  clearTimeout(this.timer)
  this.params.stop()
}
export default {
    /**
       * desc: base64对象转blob文件对象
       * @param urlData  ：数据的base64对象
       * @param type  ：类型 png,pdf,doc,mp3等;
       * @returns {Blob}：Blob文件对象
       */
    base64ToBlob (urlData, type) {
      let arr = urlData.split(',');
      let array = arr[0].match(/:(.*?);/);
      let mime = (array && array.length > 1 ? array[1] : type) || type;
      // 去掉url的头，并转化为byte
      let bytes = window.atob(arr[1]);
      // 处理异常,将ascii码小于0的转换为大于0
      let ab = new ArrayBuffer(bytes.length);
      // 生成视图（直接针对内存）：8位无符号整数，长度1个字节
      let ia = new Uint8Array(ab);
      for (let i = 0; i < bytes.length; i++) {
        ia[i] = bytes.charCodeAt(i);
      }
      return new Blob([ab], {
        type: mime
      });
    },
    /**
     * desc: 下载导出文件
     * @param blob  ：返回数据的blob对象或链接
     * @param fileName  ：下载后文件名标记
     * @param fileType  ：文件类 word(docx) excel(xlsx) ppt等
     */
    downloadExportFile (blob, fileName, fileType) {
      let downloadElement = document.createElement('a');
      let href = blob;
      if (typeof blob == 'string') {
        downloadElement.target = '_blank';
      } else {
        href = window.URL.createObjectURL(blob); //创建下载的链接
      }
      downloadElement.href = href;
      downloadElement.download = fileName + '.' + fileType; //下载后文件名
      document.body.appendChild(downloadElement);
      downloadElement.click(); //触发点击下载
      document.body.removeChild(downloadElement); //下载完成移除元素
      if (typeof blob != 'string') {
        window.URL.revokeObjectURL(href); //释放掉blob对象
      }
    },
  
    /**
     * desc: base64转文件并下载
     * @param base64 {String} : base64数据
     * @param fileType {String} : 要导出的文件类型png,pdf,doc,mp3等
     * @param fileName {String} : 文件名
     */
    downloadFile(base64, fileName, fileType) {
      let typeHeader = 'data:image/' + fileType + ';base64,' // 定义base64 头部文件类型
      let converedBase64 = typeHeader + base64;  // 拼接最终的base64
      let blob = this.base64ToBlob(converedBase64, fileType)  // 转成blob对象
      this.downloadExportFile(blob, fileName, fileType) // 下载文件
    },

    timestampToTime(times) { 
        let mdy = times[0]
        mdy = mdy.split('/')
        let month = parseInt(mdy[0])
        let day = parseInt(mdy[1])
        let year = parseInt(mdy[2])
        return year + '-' + month + '-' + day
    },
    
  
  }
  