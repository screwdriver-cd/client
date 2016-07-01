'use strict';

const clientFactory = require('./client');
const convert = require('./utils/convert');
const filter = require('./utils/filter');
const spawn = require('child_process').spawnSync;

module.exports = {
    /**
     * [description]
     * @method
     * @param  {[type]} args   [description]
     * @param  {[type]} logger [description]
     * @return {[type]}        [description]
     */
    create: (args, logger) => {
        const jobId = args.jobId;
        const container = 'node:4';

        return clientFactory().then((client) => client.v3.postV3Builds({
            body: { jobId, container }
        }).then((build) => {
            logger(JSON.stringify(convert(build), null, 4));
        }).catch((err) => {
            logger('Ooops', convert(err));
        }));
    },

    /**
     * [description]
     * @method
     * @param  {[type]} args   [description]
     * @param  {[type]} logger [description]
     * @return {[type]}        [description]
     */
    get: (args, logger) => {
        const id = args.id;

        return clientFactory().then((client) => client.v3.getV3BuildsId({ id }).then((build) => {
            logger(JSON.stringify(convert(build), null, 4));
        }).catch((err) => {
            logger('Oops', convert(err));
        }));
    },

    /**
     * [description]
     * @method
     * @param  {[type]} args   [description]
     * @param  {[type]} logger [description]
     * @return {[type]}        [description]
     */
    list: (args, logger) => {
        const options = args.options;

        return clientFactory().then((client) => client.v3.getV3Builds().then((allBuilds) => {
            const builds = filter(convert(allBuilds), options);

            logger(JSON.stringify(builds, null, 4));
        }).catch((err) => {
            logger('Oops', convert(err));
        }));
    },

    /**
     * [description]
     * @method
     * @param  {[type]} args   [description]
     * @param  {[type]} logger [description]
     * @return {[type]}        [description]
     */
    log: (args, logger) => {
        const id = args.id;
        const pod = `kubectl get pods --selector=sdbuild=${id} -a --output=jsonpath={.items..metadata.name}`;

        spawn('sh', ['-c', `kubectl logs -f \`${pod}\` build`], {
            stdio: 'inherit'
        });

        return Promise.resolve();
    }
};
