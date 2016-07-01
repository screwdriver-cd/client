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
    create: (args, logger) => {
        const scmUrl = args.scmUrl;
        const configUrl = args.configUrl;

        return clientFactory().then((client) => client.v3.postV3Pipelines({
            body: { scmUrl, configUrl }
        }).then((pipeline) => {
            logger(JSON.stringify(convert(pipeline), null, 4));
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
    list: (args, logger) => {
        const options = args.options;

        return clientFactory().then((client) => client.v3.getV3Pipelines().then((allPipelines) => {
            const pipelines = filter(convert(allPipelines), options);

            logger(JSON.stringify(pipelines, null, 4));
        }).catch((err) => {
            logger('Oops', convert(err));
        }));
    }
};
