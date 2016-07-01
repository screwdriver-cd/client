'use strict';

const clientFactory = require('./client');
const convert = require('./utils/convert');
const filter = require('./utils/filter');

module.exports = {
    /**
     * [description]
     * @method
     * @param  {[type]} args   [description]
     * @param  {[type]} logger [description]
     * @return {[type]}        [description]
     */
    list: (args, logger) => {
        const options = args.options;

        return clientFactory().then((client) => client.v3.getV3Jobs().then((allJobs) => {
            const jobs = filter(convert(allJobs), options);

            logger(JSON.stringify(jobs, null, 4));
        }).catch((err) => {
            logger('Oops', convert(err));
        }));
    }
};
