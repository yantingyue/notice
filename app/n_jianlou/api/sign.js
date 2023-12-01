let CryptoJS = require("crypto-js");


function sha1_hash(word) {
    return CryptoJS.SHA1(word).toString(CryptoJS.enc.Hex);
}
function md5_hash(word) {
    return CryptoJS.MD5(word).toString(CryptoJS.enc.Hex);
}

function get_random(len) {
    let r = Math.random().toString(36).substring(2);
    if (r.length < len) {
        r += get_random(len - r.length)
    }
    return r;

}

function get_did() {
    return md5_hash(get_random(16));
}

function get_sign_text(input) {
    let output = []
    let ptr = input;
    let output_index = 4;
    while (input.length >> 1 > output_index - 4) {
        const c1 = ptr.charCodeAt(0);
        const c2 = ptr.charCodeAt(1);
        const byte = ((c1 & 0xF0) | (c2 & 0xF));
        output[output_index++] = byte;
        ptr = ptr.slice(2);
    }
    return output.map(value => {
        return String.fromCharCode(value)
    }).join('')
}


function nice_sign_v3(JsonData, did, random) {
    let result = {
        data: "",
        random: random || get_random(),
        did: did || get_did()
    }
    let md5_list = [];
    md5_list.push(md5_hash(result.did.substring(16, 32) + result.did.substring(0, 16)));
    md5_list.push(md5_hash(result.random + md5_list[0] + "8a5f746c1c9c99c0b458e1ed510845e5"));
    md5_list[1] = md5_list[1].substring(16, 32) + md5_list[1].substring(0, 16)
    let obj = JsonData === "" ? {} : typeof JsonData === "string" ? JSON.parse(JsonData) : JsonData
    let SignText = get_sign_text(Object.keys(obj).sort().map(value => {
        return value + "=" + obj[value];
    }).join('&'))
    let Sign = sha1_hash(SignText + md5_list[1])
    result.data = `nice-sign-v1://${Sign.substring(24, 41) + Sign.substring(8, 24)}:${result.random}/${JSON.stringify(obj)}`
    return JSON.stringify(result)
}

console.log(nice_sign_v3('{"id":"851585","size_id":"240051","stock_id":"128","price":"1081","pay_type":"","address_id":"","unique_token":"VzIAYQxlXTENYAQwUTMGNg==-6a4567b99250965c330bcb7381338c13","sale_id":"","need_storage":"yes","coupon_id":"","stamp_id":"","discount_id":"","substitute_id":"","express_type":"","order_source":"","params":{},"price_list":[{"price":"1081","num":1}],"purchase_num":1}', "0c7aaccf025e5125373b37e9743147c1",'6xykg2fl5vhqpz0r'))



