import service from '../service';

export const getStorageSummary = (params) => {
    return service({
        url: '/api/v1/storage/summary',
        method: 'get',
        params: params,
    });
}

export const getStorageDataset = (params,type) => {
    return service({
        url: `/api/v1/storage/${type}`,
        method: 'get',
        params: params,
    });
}

export const batchDelStorageDataset = (params,type) => {
    return service({
        url: `/api/v1/${type}/batch_delete`,
        method: 'post',
        params: params,
    });
}