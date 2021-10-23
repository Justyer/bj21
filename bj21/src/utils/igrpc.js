module.exports = {
    marshalGrpcBytes: (dic) => {
        return new Buffer.from(JSON.stringify(dic), 'utf8').toString("base64")
    },
    bizKey: (biz) => {
        return 'reply-' + biz
    }
}