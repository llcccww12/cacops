const kvArray = [
    ["guangxi", "广西"],
    ["guangdong", "广东"],
    ["hainan", "海南"],
    ["jiangxi", "江西"],
    ["fujian", "福建"],
    ["zhejiang", "浙江"],
    ["anhui", "安徽"],
    ["jiangsu", "江苏"],
    ["shanghai", "上海"],
    ["shandong", "山东"],
    ["henan", "河南"],
    ["hebei", "河北"],
    ["shanxi", "山西"],
    ["neimenggu", "内蒙古"],
    ["liaoning", "辽宁"],
    ["heilongjiang", "黑龙江"],
    ["jilin", "吉林"],
    ["hubei", "湖北"],
    ["hunan", "湖南"],
    ["guizhou", "贵州"],
    ["sichuan", "四川"],
    ["yunnan", "云南"],
    ["xijiang", "新疆"],
    ["qinghai", "青海"],
    ["gansu", "甘肃"],
    ["ningxia", "宁夏"],
    ["taiwan", "台湾"],
    ["hongkong", "香港"],
    ["macau", "澳门"],
    ["beijing", "北京"],
    ["tianjin", "天津"],
    ["chongqing", "重庆"],
    ["xizang", "西藏"],
    ["shaanxi", "陕西"],    
]
const ProvinceKey = new Map(kvArray);

export const handelData = (r)=>{
    const v = [...r]
    const aMap = new Map()
    const arr = []
    v.forEach((e) => {
        const k = e.Province
        aMap.set(k, (aMap.get(k) || 0) + Number(e.ComputeScale))
    })
    aMap.forEach((e, k) => {
        if (e === 0) {
            e = e + 1
        }
        arr.push({ name: k, value: e })
    })
    return arr
}
export const sortData = (a, b) => {
    const aLower = a.toLowerCase();
    const bLower = b.toLowerCase();

    if (aLower < bLower) {
        return -1;
    }
    if (aLower > bLower) {
        return 1;
    }
    return 0;
}