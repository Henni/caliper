/**
* Copyright 2017 HUAWEI. All Rights Reserved.
*
* SPDX-License-Identifier: Apache-2.0
*
*/

'use strict';

module.exports.info  = 'doing nothing';


let bc, contx, iterations;
module.exports.init = function(blockchain, context, args) {
    bc         = blockchain;
    contx      = context;
    iterations = args.iterations;
    return Promise.resolve();
};

module.exports.run = function() {
    return bc.invokeSmartContract(contx, 'cpuHeavy', 'v0', {verb: 'cpuHeavy', iterations: iterations}, 30000);
};

module.exports.end = function(results) {
    // do nothing
    return Promise.resolve();
};
