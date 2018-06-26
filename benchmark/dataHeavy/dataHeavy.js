/**
* Copyright 2017 HUAWEI. All Rights Reserved.
*
* SPDX-License-Identifier: Apache-2.0
*
*/

'use strict';

module.exports.info  = 'doing nothing';


let bc, contx, payload;
module.exports.init = function(blockchain, context, args) {
    bc       = blockchain;
    contx    = context;
    payload  = new Int8Array(args.bytes);
    return Promise.resolve();
};

module.exports.run = function() {
    return bc.invokeSmartContract(contx, 'dataHeavy', 'v0', {verb: 'dataHeavy', payload: payload}, 30000);
};

module.exports.end = function(results) {
    // do nothing
    return Promise.resolve();
};
