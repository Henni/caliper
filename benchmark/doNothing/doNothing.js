/**
* Copyright 2017 HUAWEI. All Rights Reserved.
*
* SPDX-License-Identifier: Apache-2.0
*
*/

'use strict';

module.exports.info  = 'doing nothing';


let bc, contx;
module.exports.init = function(blockchain, context, args) {
    bc       = blockchain;
    contx    = context;
    return Promise.resolve();
};

module.exports.run = function() {
    return bc.invokeSmartContract(contx, 'doNothing', 'v0', {verb: 'doNothing'}, 30);
};

module.exports.end = function(results) {
    // do nothing
    return Promise.resolve();
};
