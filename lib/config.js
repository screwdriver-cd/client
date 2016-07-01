'use strict';

const config = require('configrc').make('sd');
const url = require('url');

module.exports = {
    /**
     * [description]
     * @method
     * @param  {[type]} args   [description]
     * @param  {[type]} logger [description]
     * @return {[type]}        [description]
     */
    getSet: (args, logger) => {
        const field = args.field;
        const value = args.value;

        if (value) {
            config.set(field, value);
            logger(`${field} saved with "${value}"`);
        } else {
            logger(config.get(field));
        }

        return Promise.resolve();
    },

    /**
     * [description]
     * @method
     * @param  {[type]} args   [description]
     * @param  {[type]} logger [description]
     * @return {[type]}        [description]
     */
    login: (args, logger) => {
        const loginUrl = url.resolve(config.get('url'), '/v3/login');

        logger(`Please visit the following URL: ${loginUrl}`);
        logger('Save your token via `config token <your-token-here>`');

        return Promise.resolve();
    }
};
